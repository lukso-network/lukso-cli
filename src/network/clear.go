package network

import "os"

func Clear() error {
	execVol, err := getExecutionDataVolume()
	if err != nil {
		return err
	}
	err = os.RemoveAll(execVol)
	if err != nil {
		return err
	}
	consVol, err := getConsensusDataVolume()
	if err != nil {
		return err
	}
	err = os.RemoveAll(consVol)
	if err != nil {
		return err
	}
	validatorVol, err := getValidatorDataVolume()
	if err != nil {
		return err
	}
	err = os.RemoveAll(validatorVol)
	if err != nil {
		return err
	}
	nodeconf, err := GetLoadedNodeConfigs()
	if err != nil {
		return err
	}
	configLocation, err := nodeconf.getConfigPath()
	if err != nil {
		return err
	}
	err = os.RemoveAll(configLocation)
	if err != nil {
		return err
	}
	return nil
}
