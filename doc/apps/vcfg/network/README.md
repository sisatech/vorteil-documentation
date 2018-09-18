# Network Settings (network)

Customizing an app's network configuration is perhaps the most common and most
critical thing to do in Vorteil. By default an app will generally be customized
for a single network interface using DHCP, but fine-grained control is very
commonly required. The repeatable 'network' section of a VCFG is dedicated to
such information.

## Fields

Unlike most other sections in the VCFG, the line beginning a network section is
enclosed with two sets of box brackets, instead of the usual one set. This is
an indication in TOML that the section may appear repeatedly. 

For a VCFG, each network section configures a single network interface. The 
first to appear in the file configured the zeroth network interface. The next 
configures the first, and so on. At most four network interfaces are supported 
by Vorteil at this time.

Each of the following fields must appear within a network section, which can be
identified by the appearance of the following line. Everything from that line 
up until the appearance of another section is part of the network section.

```
[[network]]
```

### ip 

**Field type: String**

```
  ip = "10.0.0.1"
```

The ip field defines what IPv4 address the kernel should set for a network 
interface. The field may be set to "dhcp" to instruct the kernel to use DHCP 
when connecting to the network.

### mask 

**Field type: String**

```
  mask = "255.255.255.0"
```

The mask field defines the subnet mask the kernel should set for a network 
interface. If the ip field is set to "dhcp" this field is ignored.

### gateway 

**Field type: String**

```
  gateway = "192.168.1.1"
```

The gateway field defines the gateway the kernel should set for a network 
interface. If the ip field is set to "dhcp" this field is ignored.

### udp

**Field type: String Array**

```
  udp = ["1000", "1002"]
```

The udp field lists every port the app will listen for UDP traffic on. This 
field has no direct impact on the app's runtime, but it is instructive for users
to understand the app, and provisioning tools may use this field to make 
external changes such as adjusting firewall rules or configuring NAT forwarding
to the virtual machine.

### tcp 

**Field type: String Array**

```
  tcp = ["1000", "1002"]
```

The tcp field lists every port the app will listen for TCP traffic on. This 
field has no direct impact on the app's runtime, but it is instructive for users
to understand the app, and provisioning tools may use this field to make 
external changes such as adjusting firewall rules or configuring NAT forwarding
to the virtual machine.

### http 

**Field type: String Array**

```
  http = ["80", "8080"]
```

The http field lists every port the app will listen for HTTP traffic on. This 
field has no direct impact on the app's runtime, but it is instructive for users
to understand the app, and provisioning tools may use this field to make 
external changes such as adjusting firewall rules or configuring NAT forwarding
to the virtual machine.

### https 

**Field type: String Array**

```
  https = ["80", "8080"]
```

The https field lists every port the app will listen for HTTPS traffic on. This 
field has no direct impact on the app's runtime, but it is instructive for users
to understand the app, and provisioning tools may use this field to make 
external changes such as adjusting firewall rules or configuring NAT forwarding
to the virtual machine.

### mtu

**Field type: Unsigned Integer**

```
  mtu = 1500
```

The mtu field allows the app to define the maximum transmission unit for a 
network interface. This is an advanced setting that should only be used if you 
know what you're doing. Vorteil's compilers have no way of safety checking the
value here, so it must be appropriate for the infrastructure it runs on. The 
default value should be appropriate for most environments.

### disable-tso

**Field type: Boolean**

```
  disable-tso = true
```

The disable-tso field allows the app to define whether it may try to negotiate
TCP segmentation offloading. This is an advanced setting that may have 
signficant performance implications in some environments.

## Notes

Although HTTP traffic operates over TCP, and although HTTPS traffic operates 
over HTTP, there is no need to duplicate the values for tcp, http, and https 
fields across each other. Tools that make use of this information should be 
sophisticated enough to understand that the fields may need to be combined.
Developers should list a port in only one list, and it should be in the most 
instructive list possible. The distinction between TCP, HTTP, and HTTPS matters 
in some contexts. 

## See Also

* [What is a VCFG?](../introduction)
* [Example VCFG](../example)
* [CLI's Network Config Documentation](../../../cli/vcfgs/network)