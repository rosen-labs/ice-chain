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

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdSendMsgMintRequest())

	cmd.AddCommand(GetBridgeXTokenCmd())

	return cmd
}

func GetBridgeXTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge-to [chain_id] [amount] [fee] [reciever]",
		Short: "Bridge amounts of x token to specific chain",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAmount, err := argsToUint64("chain_id", args[1])
			if err != nil {
				return err
			}
			argsFee, err := argsToUint64("chain_id", args[2])
			if err != nil {
				return err
			}
			argsReciever := string(args[3])
			argsChainId, err := argsToUint32("chain_id", args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			if from == nil {
				return fmt.Errorf("must pass from flag")
			}

			msg := types.NewMsgBridgeRequest(
				argsChainId,
				argsAmount,
				argsFee,
				argsReciever,
				from,
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

func argsToUint32(argsName string, value string) (uint32, error) {
	temp, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("%s must be number (uint32)", argsName)
	} else {
		return uint32(temp), nil
	}
}

func argsToUint64(argsName string, value string) (uint64, error) {
	temp, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s must be number (uint32)", argsName)
	} else {
		return temp, nil
	}
}
