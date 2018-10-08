# Network

By default, all ports are closed for the virtual machine being run on Vorteil unless specified otherwise. Its very simple to open access to ports with a distinct protocol with modifying the VCFG that is being used to run the application.

For example to expose port 8080 to HTTP traffic and 443 to HTTPS traffic is shown below.

```toml
[[network]]
    ip="dhcp"
    http = ["8080"]
    https = ["443"]
```
The double brackets in this TOML struct also imply its repeatable so we're able to have more than one network if required.

If the ports the configuration is asking for are not available at run time (for example running on a NAT network), the Vorteil tools will automatically select an alternative available port.

Supported protocols to include within the 'network' area of an application config include:

- http
- https
- udp
- tcp

Additional settings include:

- gateway
- mask
- mtu (maximum transmission unit)
- disable-tso (boolean value: disable TCP segmentation offloading)

## See Also
- [In-depth Networking Details](../../../apps/vcfg/network)