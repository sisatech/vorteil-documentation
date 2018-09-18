# View or change the default platform (vorteil settings platforms default)

The default platform will be used whenever provisioning virtual machines unless
explicitly commanded otherwise. This command has two different modes: providing
no argument will print the current default platform's ID to stdout but will make
no changes. If a PLATFORM argument is provided the CLI will attempt to change
the server's default platform instead.

# Usage

```
vorteil settings platforms default [PLATFORM] [flags]
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

* [vorteil settings platforms](../platforms)	 - Define platforms for hosting virtual machines

###### Auto generated for CLI 2.0.0-7169db7d-dirty on 17-Sep-2018
