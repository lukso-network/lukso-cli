# Change Log



## [0.1.3-rc] - 2022-05-23

This update prepares the cli for multi chains.

### Added

- command **lukso network block -n N** that will return the N-th execution block
- command **lukso network validator start** that will start a validator
- command **lukso network validator stop** that will stop a validator
- command **lukso network balance -a 0x...** to call balance of validator
- new chain env **local**. A network can be setup with **lukso network init --chain local**
- command **lukso network update** to update the bootnodes 
- command **lukso network config** to only create a config file

### Changed

- command **lukso network start validator** became deprecated
- 2 new scripts installer script **install.sh** & **download.sh** to give more control on how to get the CLI binary

### Fixed

- added Api entity to attach to the load balancer
- changed the **lukso network init** to receive configs from different directory

## [0.0.4] - 2022-05-13

This update mainly improves existing commands 

### Added

- command **lukso network describe** that will give information on the state of the network
- command **lukso network validator deposit** that will deposit new validators

### Changed

- command **lukso network validator describe** that will show additional information
- the installer script **cli_downloader.sh** will copy the cli binary to /usr/local/bin

### Fixed

- command **lukso network validator setup** it will create a local transaction wallet.