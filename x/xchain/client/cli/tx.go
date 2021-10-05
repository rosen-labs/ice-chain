package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/rosen-labs/xchain/x/xchain/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	return cmd
}

func BridgeXToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge-to [chain_id] [amount] [fee] [reciever]",
		Short: "Bridge amounts of x token to specific chain",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsChainId := string(args[0])
			argsAmount := string(args[1])
			argsFee := string(args[2])
			argsReciever := string(args[3])

			amount, err := strconv.ParseFloat(argsAmount, 32)
			if err == nil {
				return fmt.Errorf("amount must be float")
			}

			fee, err := strconv.ParseFloat(argsFee, 32)
			if err == nil {
				return fmt.Errorf("fee must be float")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBridgeRequest(
				argsChainId,
				float32(amount),
				float32(fee),
				argsReciever,
				clientCtx.GetFromAddress().String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
