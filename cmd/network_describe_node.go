/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/api"
	"github.com/lukso-network/lukso-cli/api/gethrpc"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "gets a description of peers",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := cmd.Flags().GetString("ip")
		if ip == "" {
			fmt.Println("Error: ip is empty")
		}

		i := api.NewIP(ip)

		c := gethrpc.NewRPCClient(i.RPCAddress())

		nodeInfo, err := api.AdminNodeInfoRequest(c)

		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		peers, err := api.AdminPeersRequest(c)
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		blocknumber, err := api.BlockNumber(c)
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		fmt.Println("Execution:")
		fmt.Println(".........................................")
		fmt.Println("Enode: ", nodeInfo.Enode)
		fmt.Println("Peers: ", len(peers))
		fmt.Println("Latest Block:", blocknumber)

		//
		//fmt.Println("Consensus:")
		//fmt.Println(".........................................")
		//fmt.Println("Enode: ", nodeInfo.Enode)
		//fmt.Println("Peers: ", len(peers))
	},
}

func init() {
	describeCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringP("ip", "i", "", "set ip")
	_ = nodeCmd.MarkFlagRequired("ip")

}
