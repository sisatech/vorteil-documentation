# Probe an image for information (vorteil images analyze)

Probe a VMDK or a RAW disk image for detailed diagnostic information. Currently
this command is useful only for identifying redundant package contents.

The common behaviour of the probe is to output a comprehensive list of every
item on the image's file-system. Items printed in white text have been accessed
since the image was built, while items printed in blue text have never been
accessed. If an image has never been run all items will appear blue, but if it
has been run and the program's logic branches have been tested extensively then
any files that remain blue are probably redundant and OK to remove from the app
source.

# Usage

```
vorteil images analyze [flags] IMAGE
```

## Aliases

```
analyse
```

## Options

```
      --cull string   Cut redundant files from a project
  -h, --help          help for analyze
```

## Hidden options

```
  -d, --debug           print debug information to stderr
      --peer string     route requests to this peer of the server
      --server string   server network address (default "127.0.0.1:7472")
  -v, --verbose         print additional information to stderr
```


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 17-Sep-2018
