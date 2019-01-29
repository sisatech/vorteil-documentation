# Network Routing

By default Vorteil will route all traffic via the default gateway. Although
routing is not required for most of the applications, there might be
cases where traffic needs to be routed via a different gateway. The repeatable
'route' section of a VCFG is dedicated to such information. Up to four routes
can be configured.


## Fields

Unlike most other sections in the VCFG, the line beginning a 'route' section is
enclosed with two sets of box brackets, instead of the usual one set. This is
an indication in TOML that the section may appear repeatedly.

For a VCFG, each routing section configures a single routing entry. Up to four
routes can be configured.

Each of the following fields must appear within a 'route' section, which can be
identified by the appearance of the following line. Everything from that line
up until the appearance of another section is part of the routing section.

```
[[route]]
```

### interface

**Field type: String**

```
  interface = "eth0"
```

The interface field defines the network card which this routing entry applies
to. The network card has to be defined in the network section where up to four
network cards can be configured. Valid values for this field are eth0, eth1,
eth2 and eth3.

### destination

**Field type: String**

```
  destination = "10.0.10.0/24"
```

The mask field defines the destination network for this route. The value needs
to be in "slash notation".

### gateway

**Field type: String**

```
  gateway = "192.168.1.1"
```

The gateway field defines the gateway the kernel should set for a route.


## See Also

* [What is a VCFG?](../introduction)
* [Example VCFG](../example)
* [CLI's Network Config Documentation](../../../cli/vcfgs/network)
