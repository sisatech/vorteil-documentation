# Untag an app (vorteil repositories untag)

Remove a tag from an app. If ref is left unspecified the most recent version of
the app will be untagged. This command cannot be used to remove or change an
app's reference ID.

# Usage

```
vorteil repositories untag BUCKET/APP[/REF] [flags]
```

## Options

```
  -h, --help                help for untag
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
