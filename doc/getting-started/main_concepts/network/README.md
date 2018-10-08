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

If the ports the configuration is asking for are not available at run time (for example running on a NATed network). The Vorteil Tools will automatically select an available port instead which will then be displayed to the users that the certain port they could not get has been forwarded to the new one.

Optional fields for the network object are shown below.
- http/https: both of these objects can also be optional.
- mask: defines the subnet mask the kernel should set for the network interface.
- gateway: defines the gateway the kernel should set for a network interface.
- udp: a list of ports that the UDP traffic will listen on.
- tcp: a list of ports that the TCP traffic will listen on.
- mtu: the maximum transmission unit for a network interface.
- disable-tso: a boolean to determine whether to use segmentation offloading.

## See Also
- [In-depth Networking Details](../../../apps/vcfg/network)