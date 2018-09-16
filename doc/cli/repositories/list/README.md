# List repositories, buckets, apps, or app versions (vorteil repositories list)

List objects on a repository. Without any provided arguments the list will be of
all available buckets. If a bucket is specified without also specifying an app
then the list will be of all apps within that bucket. If both bucket and app are
specified, the list will be of all versions for that app.

# Usage

```
vorteil repositories list BUCKET[/APP] [flags]
```

## Options

```
  -h, --help                help for list
      --repository string   choose a repository to perform the operation on
```

## Hidden options

```
  -d, --debug           print debug information to stderr
      --peer string     route requests to this peer of the server
  -p, --plain           
      --server string   server network address (default "127.0.0.1:7472")
  -v, --verbose         print additional information to stderr
```


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 14-Sep-2018
