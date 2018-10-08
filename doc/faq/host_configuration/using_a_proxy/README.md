# Configuring Vorteil to Use a Proxy 
## Linux
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

## Windows
1. Select the 'Edit the system environment variables' from the Start Menu.

![Edit the system environment variables](https://downloads.vorteil.io/assets/documentation/windows-proxy-01.png)

2. Click on 'Environment Variables..."

![Click on 'Environment Variables...'](https://downloads.vorteil.io/assets/documentation/windows-proxy-02.png)

3. Add 'http_proxy' and 'https_proxy' as environment variables, using the correct proxy server URL as the value.

![Add the environment variables](https://downloads.vorteil.io/assets/documentation/windows-proxy-03.png)

4. Restart your machine for the changes to take effect.