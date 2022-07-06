package keeper

import (
	"context"
	epochingtypes "github.com/babylonchain/babylon/x/epoching/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/babylonchain/babylon/x/checkpointing/types"
)

type msgServer struct {
	k Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) AddBlsSig(goCtx context.Context, header *types.MsgAddBlsSig) (*types.MsgAddBlsSigResponse, error) {
	panic("TODO: implement me")
}

// WrappedCreateValidator stores validator's BLS public key as well as corresponding MsgCreateValidator message
func (m msgServer) WrappedCreateValidator(goCtx context.Context, msg *types.MsgWrappedCreateValidator) (*types.MsgWrappedCreateValidatorResponse, error) {
	// TODO: verify pop
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.k.CreateRegistration(ctx, *msg.Key.Pubkey, types.ValidatorAddress(msg.MsgCreateValidator.ValidatorAddress))
	if err != nil {
		return nil, err
	}

	// enqueue the msg into the epoching module
	queueMsg := epochingtypes.QueuedMessage{
		Msg: &epochingtypes.QueuedMessage_MsgCreateValidator{MsgCreateValidator: msg.MsgCreateValidator},
	}

	err = m.k.epochingKeeper.EnqueueMsg(ctx, queueMsg)
	if err != nil {
		return nil, err
	}

	return &types.MsgWrappedCreateValidatorResponse{}, err
}
