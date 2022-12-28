package gov

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/shentufoundation/shentu/v2/x/gov/keeper"
	"github.com/shentufoundation/shentu/v2/x/gov/types"
	shieldtypes "github.com/shentufoundation/shentu/v2/x/shield/types"
)

func removeInactiveProposals(ctx sdk.Context, k keeper.Keeper) {
	logger := k.Logger(ctx)

	k.IterateInactiveProposalsQueue(ctx, ctx.BlockHeader().Time, func(proposal types.Proposal) bool {
		k.DeleteProposalByProposalID(ctx, proposal.ProposalId)
		k.RefundDepositsByProposalID(ctx, proposal.ProposalId)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				govTypes.EventTypeInactiveProposal,
				sdk.NewAttribute(govTypes.AttributeKeyProposalID, fmt.Sprintf("%d", proposal.ProposalId)),
				sdk.NewAttribute(govTypes.AttributeKeyProposalResult, govTypes.AttributeValueProposalDropped),
			),
		)

		logger.Info(
			"proposal did not meet minimum deposit; deleted",
			"proposal", proposal.ProposalId,
			"title", proposal.GetTitle(),
			"min_deposit", k.GetDepositParams(ctx).MinDeposit.String(),
			"total_deposit", proposal.TotalDeposit.String(),
		)

		return false
	})
}

func updateVeto(ctx sdk.Context, k keeper.Keeper, proposal types.Proposal) {
	if proposal.ProposalType() == shieldtypes.ProposalTypeShieldClaim {
		c := proposal.GetContent().(*shieldtypes.ShieldClaimProposal)
		k.ShieldKeeper.ClaimEnd(ctx, c.ProposalId, c.PoolId, c.Loss)
	}
}

func updateAbstain(ctx sdk.Context, k keeper.Keeper, proposal types.Proposal) {
	if proposal.ProposalType() == shieldtypes.ProposalTypeShieldClaim {
		c := proposal.GetContent().(*shieldtypes.ShieldClaimProposal)
		proposer, err := sdk.AccAddressFromBech32(proposal.ProposerAddress)
		if err != nil {
			panic(err)
		}
		k.ShieldKeeper.RestoreShield(ctx, c.PoolId, proposer, c.PurchaseId, c.Loss)
		k.ShieldKeeper.ClaimEnd(ctx, c.ProposalId, c.PoolId, c.Loss)
	}
}

func processActiveProposal(ctx sdk.Context, k keeper.Keeper, proposal types.Proposal) bool {
	var (
		tagValue, logMsg string
		pass, veto       bool
		tallyResults     govTypes.TallyResult
	)
	logger := k.Logger(ctx)

	if proposal.HasSecurityVoting() && !k.IsCertifierVoted(ctx, proposal.ProposalId) {
		//if proposal.Status == govTypes.StatusCertifierVotingPeriod {
		var endVoting bool
		pass, endVoting, tallyResults = keeper.SecurityTally(ctx, k, proposal)
		if !endVoting {
			// Skip the rest of this iteration, because the proposal needs to go
			// through the validator voting period now.
			k.SetCertifierVoted(ctx, proposal.ProposalId)
			k.DeleteAllVotes(ctx, proposal.ProposalId)
			k.ActivateVotingPeriod(ctx, proposal)
			return false
		}
	} else {
		pass, veto, tallyResults = keeper.Tally(ctx, k, proposal)
	}

	if veto {
		k.DeleteDepositsByProposalID(ctx, proposal.ProposalId)
		updateVeto(ctx, k, proposal)
	} else {
		k.RefundDepositsByProposalID(ctx, proposal.ProposalId)
		if !pass {
			updateAbstain(ctx, k, proposal)
		}
	}

	if pass {
		handler := k.Router().GetRoute(proposal.ProposalRoute())
		cacheCtx, writeCache := ctx.CacheContext()
		// The proposal handler may execute state mutating logic depending on the
		// proposal content. If the handler fails, no state mutation is written and
		// the error message is logged.
		err := handler(cacheCtx, proposal.GetContent())
		if err == nil {
			proposal.Status = govTypes.StatusPassed
			tagValue = govTypes.AttributeValueProposalPassed
			logMsg = "passed"

			// write state to the underlying multi-store
			writeCache()
		} else {
			proposal.Status = govTypes.StatusFailed
			tagValue = govTypes.AttributeValueProposalFailed
			logMsg = fmt.Sprintf("passed, but failed on execution: %s", err)
		}
	} else {
		proposal.Status = govTypes.StatusRejected
		tagValue = govTypes.AttributeValueProposalRejected
		logMsg = "rejected"
	}

	proposal.FinalTallyResult = tallyResults

	k.SetProposal(ctx, proposal)
	k.RemoveFromActiveProposalQueue(ctx, proposal.ProposalId, proposal.VotingEndTime)

	logger.Info(
		"proposal tallied",
		"proposal", proposal.ProposalId,
		"title", proposal.GetTitle(),
		"result", logMsg,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			govTypes.EventTypeActiveProposal,
			sdk.NewAttribute(govTypes.AttributeKeyProposalID, fmt.Sprintf("%d", proposal.ProposalId)),
			sdk.NewAttribute(govTypes.AttributeKeyProposalResult, tagValue),
		),
	)
	return false
}

func processSecurityVote(ctx sdk.Context, k keeper.Keeper, proposal types.Proposal) bool {
	var (
		tagValue     string
		pass         bool
		tallyResults govTypes.TallyResult
	)

	// Only process proposals in the security voting period.
	if proposal.HasSecurityVoting() && k.IsCertifierVoted(ctx, proposal.ProposalId) {
		//if proposal.Status != govTypes.StatusCertifierVotingPeriod {
		return false
	}

	var endVoting bool
	pass, endVoting, tallyResults = keeper.SecurityTally(ctx, k, proposal)
	if !pass {
		// Do nothing, because the proposal still has time before the voting period
		// ends.
		return false
	}
	// Else: the proposal passed the certifier voting period.

	if endVoting {
		handler := k.Router().GetRoute(proposal.ProposalRoute())
		cacheCtx, writeCache := ctx.CacheContext()

		// The proposal handler may execute state mutating logic depending on the
		// proposal content. If the handler fails, no state mutation is written and
		// the error message is logged.
		err := handler(cacheCtx, proposal.GetContent())
		if err == nil {
			proposal.Status = govTypes.StatusPassed
			tagValue = govTypes.AttributeValueProposalPassed

			// write state to the underlying multi-store
			writeCache()
		} else {
			proposal.Status = govTypes.StatusFailed
			tagValue = govTypes.AttributeValueProposalFailed
		}

		proposal.FinalTallyResult = tallyResults

		k.SetProposal(ctx, proposal)
		k.RemoveFromActiveProposalQueue(ctx, proposal.ProposalId, proposal.VotingEndTime)

		// TODO log tallying result

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				govTypes.EventTypeActiveProposal,
				sdk.NewAttribute(govTypes.AttributeKeyProposalID, fmt.Sprintf("%d", proposal.ProposalId)),
				sdk.NewAttribute(govTypes.AttributeKeyProposalResult, tagValue),
			),
		)
	} else {
		// Activate validator voting period
		k.DeleteAllVotes(ctx, proposal.ProposalId)
		k.ActivateVotingPeriod(ctx, proposal)
		k.SetCertifierVoted(ctx, proposal.ProposalId)
	}
	return false
}

// EndBlocker is called every block, removes inactive proposals, tallies active
// proposals and deletes/refunds deposits.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	// delete inactive proposal from store and its deposits
	removeInactiveProposals(ctx, k)

	// fetch active proposals whose voting periods have ended (are passed the
	// block time)
	k.IterateActiveProposalsQueue(ctx, ctx.BlockHeader().Time, func(proposal types.Proposal) bool {
		return processActiveProposal(ctx, k, proposal)
	})

	// Iterate over all active proposals, regardless of end time, so that
	// security voting can end as soon as a passing threshold is met.
	k.IterateActiveProposalsQueue(ctx, ctx.BlockHeader().Time, func(proposal types.Proposal) bool {
		return processSecurityVote(ctx, k, proposal)
	})
}
