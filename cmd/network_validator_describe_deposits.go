/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// validatorDescribeDepositsCmd represents the check command
var validatorDescribeDepositsCmd = &cobra.Command{
	Use:     "deposits",
	Short:   "Show detailed status of the validators as deposited in the DepositDetails Contract",
	Long:    `Show detailed status of the validators as deposited in the DepositDetails Contract.`,
	Example: "lukso network validator check",
	Run: func(cmd *cobra.Command, args []string) {
		// get node conf from --chain param or get default chain
		nodeConf := network.GetDefaultNodeConfigByOptionParam(viper.GetString(network.CommandOptionChain))

		consensusApi, err := readConsensusApiEndpoint(cmd, nodeConf)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		executionApi, err := readExecutionApiEndpoint(cmd, nodeConf)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		address, err := readDepositAddress(cmd, nodeConf)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		events, err := network.NewDepositEvents(address, nodeConf.ApiEndpoints.ExecutionApi)
		if err != nil {
			cobra.CompErrorln(fmt.Sprintf("couldn't load deposit data from contract, reason: %s", err.Error()))
			return
		}

		err = network.DescribeValidatorKey(events.GetUniqueKeys(), address, executionApi, consensusApi, &events)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		utils.ColoredPrintln("Total Amount Of Deposits", utils.GWeiToString(events.TotalAmount(), true))
		utils.ColoredPrintln("Total Unique Validators", len(events.GetUniqueKeys()))
		utils.ColoredPrintln("Total Deposit Events", len(events.Events))

	},
}

func init() {
	validatorDescribeCmd.AddCommand(validatorDescribeDepositsCmd)
	validatorDescribeDepositsCmd.Flags().StringP(CommandOptionExecutionApi, CommandOptionExecutionApiShort, "", "execution api endpoint")
	validatorDescribeDepositsCmd.Flags().StringP(CommandOptionConsensusApi, CommandOptionConsensusApiShort, "", "consensus api endpoint")
	validatorDescribeDepositsCmd.Flags().StringP(CommandOptionDepositAddress, CommandOptionDepositAddressShort, "", "the address of the deposit contract")
}
