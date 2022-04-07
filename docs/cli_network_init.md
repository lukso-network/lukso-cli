## cli network init

Initializes node by downloading configs and scripts

### Synopsis

This command downloads network starter scripts and config files
from the github repository. It also updates node name and IP address in the .env file

```
cli network init [flags]
```

### Examples

```
lukso-cli network init --config ./node_config.yaml --chainId l16 --nodeName my_node --docker
```

### Options

```
      --chainId string    provide chain ID for LUKSO network (default "l16")
      --docker            use docker or not (default true)
  -h, --help              help for init
      --nodeName string   set node name
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/node_config.yaml)
```

### SEE ALSO

* [cli network](cli_network.md)	 - Parent command of LUKSO network

###### Auto generated by spf13/cobra on 7-Apr-2022