## cli network stop

Stops running docker containers

### Synopsis

This command stops consensus engine, execution engine, validator client and eth2-stats.
It uses docker-compose file to stop these containers

```
cli network stop [flags]
```

### Examples

```
lukso network stop
```

### Options

```
  -h, --help   help for stop
```

### Options inherited from parent commands

```
      --chain string   provide chain you want to target [l16,...] (default "mainnet")
```

### SEE ALSO

* [cli network](cli_network.md)	 - subcommand "network" for LUKSO network related things

###### Auto generated by spf13/cobra on 21-Jun-2022
