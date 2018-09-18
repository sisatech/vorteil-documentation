# Connect to a Vorteil repository (vorteil repositories connect)

Connect to a remote repository so that the repository subcommands can interact
with them. The name argument should be a short and unique name that you will use
to specify the repository using the 'repository' flag in future commands. The
url argument should be a URL, domain name, or IP address for the machine where
the repository is hosted.

Credentials is an optional argument that can be used to pass in an encrypted
access token that will be used to authenticate and authorize access to the
repository. If no credentials are provided the interactions will be attempted
anonymously, which will have varying degrees of success depending on the
administration of the repository.

# Usage

```
vorteil repositories connect NAME ADDRESS [CREDENTIALS] [flags]
```

## Options

```
  -h, --help                   help for connect
  -k, --insecure-skip-verify   disables client security check for this connection
```

## Hidden options

```
  -d, --debug           print debug information to stderr
      --peer string     route requests to this peer of the server
      --server string   server network address (default "127.0.0.1:7472")
  -v, --verbose         print additional information to stderr
```


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 17-Sep-2018
