package network

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const (
	AutoGeneratedEnvFileMessage = `##############################################################################################################################
#
# AUTO GENERATED FILE - please do not modify - modify node_conf.yaml instead
#
##############################################################################################################################
`
)

func GenerateEnvFile() error {
	fmt.Printf("Generating .env file from ")
	err := WriteEnvMap(GetEnvironmentConfig(), ".env")
	if err != nil {
		return err
	}
	return nil
}

func WriteEnvMap(envMap map[string]string, filename string) error {
	content, err := godotenv.Marshal(envMap)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(AutoGeneratedEnvFileMessage + "\n\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString(content + "\n")
	if err != nil {
		return err
	}
	file.Sync()
	return err
}

func GetDefaultNodeConfigsIfDoesntExist(chain Chain) *NodeConfigs {
	if !FileExists(NodeConfigLocation) {
		return GetDefaultNodeConfig(chain)
	}
	return nil
}

func SetupNetwork(chain Chain) error {
	fmt.Printf("Setting up node for chain %s.\n", chain.String())
	clientVersion := BeaconClientPrysm

	if !IsChainSupported(chain) {
		return fmt.Errorf("the network %s does not exist or is not supported\n", chain.String())
	}

	return NewResourceDownloader(chain, clientVersion).DownloadAll()
}

func SetupDevNetwork(devLocation string) error {
	fmt.Printf("Setting up node for dev chain %s.\n", devLocation)
	clientVersion := BeaconClientPrysm
	return NewDevResourceDownloader(devLocation, clientVersion).DownloadAll()
}
