package client

import (
	"github.com/Stride-Labs/stride/v8/x/stakeibc/client/cli"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	AddValidatorsProposalHandler = govclient.NewProposalHandler(cli.CmdAddValidatorsProposal)
)
