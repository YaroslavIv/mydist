package keeper

import (
	"context"
	"fmt"

	"example/x/mydist/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fmt.Println("\n\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n\n")
	fmt.Printf("GetProposerAddress HEX: %x\n", ctx.CometInfo().GetProposerAddress())
	fmt.Printf("GetLastCommit Round: %d\n", ctx.CometInfo().GetLastCommit().Round())
	fmt.Printf("GetLastCommit Len: %d\n", ctx.CometInfo().GetLastCommit().Votes().Len())
	fmt.Printf("GetLastCommit Get: %x\n", ctx.CometInfo().GetLastCommit().Votes().Get(0).Validator().Address())
	fmt.Println("\n\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n\n")

	fmt.Printf("ProposerAddress : %x\n", ctx.BlockHeader().ProposerAddress)
	fmt.Println("\n\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n\n")

	for i, vote := range ctx.VoteInfos() {
		fmt.Printf("VoteInfos[%d]: %x\n", i, vote.Validator.GetAddress())
	}
	fmt.Println("\n\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n\n")

	addr, _ := sdk.Bech32ifyAddressBytes(sdk.Bech32PrefixValAddr, ctx.CometInfo().GetProposerAddress())
	fmt.Printf("ConsensusAddress: %s", addr)
	addr, _ = sdk.Bech32ifyAddressBytes(sdk.Bech32PrefixAccAddr, ctx.CometInfo().GetProposerAddress())
	fmt.Println("\n\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n\n")

	addr, _ = k.stakingKeeper.ConsensusAddressCodec().BytesToString(ctx.CometInfo().GetProposerAddress())
	add, _ := k.stakingKeeper.ValidatorByConsAddr(ctx, ctx.CometInfo().GetProposerAddress())
	fmt.Printf("Operator: %s", add.GetOperator())
	fmt.Println("\n\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n\n")

	post := types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}
	id := k.AppendPost(ctx, post)

	return &types.MsgCreatePostResponse{Id: id}, nil
}
