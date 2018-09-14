# List available repositories (vorteil repositories connections)

List repositories that have been added using the 'connect' command. This command
produces a complete list of every available repository, with the exception of
the local repository, which is not deletable and accessible via the reserved
name 'local'.

# Usage

```
vorteil repositories connections [flags]
```

## Options

```
  -h, --help   help for connections
```

## Hidden options

```
  -d, --debug           print debug information to stderr
      --peer string     route requests to this peer of the server
      --server string   server network address (default "127.0.0.1:7472")
  -v, --verbose         print additional information to stderr
```


###### Auto generated for CLI 2.0.0-2ffceded-dirty on 14-Sep-2018
