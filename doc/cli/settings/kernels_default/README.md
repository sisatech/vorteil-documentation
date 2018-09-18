# View or change the default kernel (vorteil settings kernels default)

The default kernel will be used whenever building or packaging apps unless
explicitly commanded otherwise. This command has two different modes: providing
no argument will print the current default kernel to stdout but will make no
changes. If a KERNEL argument is provided the CLI will attempt to change the
server's default kernel instead.

# Usage

```
vorteil settings kernels default [KERNEL] [flags]
```

## Options

```
  -h, --help   help for default
```

## Hidden options

```
  -d, --debug           print debug information to stderr
      --peer string     route requests to this peer of the server
      --server string   server network address (default "127.0.0.1:7472")
  -v, --verbose         print additional information to stderr
```

# See Also

* [vorteil settings kernels](../kernels)	 - Manage a library of kernel versions

###### Auto generated for CLI 2.0.0-7169db7d-dirty on 17-Sep-2018
