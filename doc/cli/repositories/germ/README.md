# Output a valid germ string for an app within a repository (vorteil repositories germ)

This simple helper command outputs an equivalent 'germ string' for the provided
arguments, usable as a target germ wherever applicable.

# Usage

```
vorteil repositories germ BUCKET/APP[/REF] [flags]
```

## Options

```
  -h, --help                help for germ
      --repository string   choose a repository to perform the operation on
```

## Hidden options

```
  -d, --debug           print debug information to stderr
      --peer string     route requests to this peer of the server
      --server string   server network address (default "127.0.0.1:7472")
  -v, --verbose         print additional information to stderr
```


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 17-Sep-2018
