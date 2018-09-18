# NFS Settings (nfs)

The Vorteil kernel includes functionality to mount a network file-system (NFS)
using NFS4. The repeatable 'nfs' section of a VCFG is dedicated to configuring
NFS mounts.

## Fields

Unlike most other sections in the VCFG, the line beginning a nfs section is
enclosed with two sets of box brackets, instead of the usual one set. This is
an indication in TOML that the section may appear repeatedly. 

For a VCFG, each nfs section configures a single nfs mount. Each of the 
following fields must appear within a nfs section, which can be identified by 
the appearance of the following line. Everything from that line up until the
appearance of another section is part of the nfs section.

```
[[nfs]]
```

### mount

**Field type: String**

```
  mount = "/nfs"
```

The mount field determines the path on the file-system where the NFS will be
mounted.

### server

**Field type: String**

```
  server = "192.168.1.10"
```

The server field must provide a resolvable address to a NFS server.

## Notes

Although the example above does not demonstrate this, the server field can be
a domain name if it is resolvable via the app's DNS servers.

## See Also

* [What is a VCFG?](../introduction)
* [Example VCFG](../example)