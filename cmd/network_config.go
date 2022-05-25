/*
Copyright © 2022 The LUKSO authors

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// initCmd represents the setup command
var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Creates node_config.yaml based on the chain provided.",
	Long:    `This command will create the default config for a given chain`,
	Example: "lukso network [--chain l16beta]",
	RunE: func(cmd *cobra.Command, args []string) error {
		chain := network.GetChainByString(viper.GetString(network.CommandOptionChain))
		isGenerated, err := network.GenerateDefaultNodeConfigsIfDoesntExist(chain)
		if err != nil {
			cobra.CompErrorln(err.Error())
			os.Exit(1)
		}

		if isGenerated {
			fmt.Println("Successfully created config file node_config.yaml...")
		} else {
			fmt.Println("./node_config.yaml already exists")
		}

		return nil
	},
}

func init() {
	//networkCmd.AddCommand(configCmd)
}