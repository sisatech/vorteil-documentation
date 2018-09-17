# Delete a bucket, app, or app version (vorteil repositories delete)

Delete objects on a repository. If a bucket is specified without also specifying
an app then the entire bucket will be deleted. If both bucket and app are
specified, the entire app will be deleted. If a specific app version, tag, or
reference is also specified, then only that specific app ref will be deleted.

# Usage

```
vorteil repositories delete BUCKET[/APP[/REF]] [flags]
```

## Options

```
  -h, --help                help for delete
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
