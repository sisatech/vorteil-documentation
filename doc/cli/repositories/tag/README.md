# Tag an app (vorteil repositories tag)

Tag a version of an app with a more memorable or meaningful string than its
reference ID. If ref is left unspecified the tag will apply to the most recent
version of the app.

# Usage

```
vorteil repositories tag BUCKET/APP[/REF] TAG [flags]
```

## Options

```
  -h, --help                help for tag
      --repository string   choose a repository to perform the operation on
```

## Hidden options

```
  -d, --debug           print debug information to stderr
      --peer string     route requests to this peer of the server
      --server string   server network address (default "127.0.0.1:7472")
  -v, --verbose         print additional information to stderr
```


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 14-Sep-2018
