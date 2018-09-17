# Configuring Vorteil to Use a Proxy (Linux)
The Vorteil daemon/service can be configured to use a proxy by making a small adjustment to the 'vorteil-daemon' service file:

1. Locate the 'vorteil-daemon' service file (on a standard installation, this should be located at '/etc/systemd/system/vorteil-daemon.service').

2. Edit the service file, adding the following entries beneath the [Service] header, where appropriate:

```systemd
[Service]
Environment=HTTP_PROXY='http://user:password@prox-server:3128'
Environment=HTTPS_PROXY='http://user:password@prox-server:3128'
```

3. Restart the service to reflect these changes.

```bash
$ systemctl daemon-reload
$ service vorteil-daemon restart
```

The Vorteil daemon is now configured to use the proxy information provided in the 'HTTP_PROXY' and/or 'HTTPS_PROXY' environment variables.