/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
)

// depositCmd represents the deposit command
var depositAltCmd = &cobra.Command{
	Use:   "deposit_alt",
	Short: "Send Deposit transactions to activate validator",
	Long: `After preparing wallets and deposit data, this command prepares deposit transactions to the deposit contract
address. Remember it will need your wallet address and private keys. Thus it will deduct balance from your wallet.

This tool is necessary to activate new validators`,
	Example: "lukso network validator deposit_alt",
	Run: func(cmd *cobra.Command, args []string) {
		nodeConf, err := network.GetLoadedNodeConfigs()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		valSecrets := nodeConf.GetValSecrets()
		if valSecrets == nil {
			cobra.CompErrorln("no validator credential is presented")
			return
		}

		if err := network.Deposit(valSecrets.Deposit.DepositFileLocation, valSecrets.Deposit.ContractAddress, valSecrets.Eth1Data.WalletPrivKey, nodeConf.ApiEndpoints.ExecutionApi); err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		fmt.Println("You successfully deposited your key(s)! Your keys need to be activated which takes around 8h. You can check the status by calling:")
		fmt.Println(utils.ConsoleInBlue("        lukso network validator describe"))
	},
}

func init() {
	validatorCmd.AddCommand(depositAltCmd)
}
