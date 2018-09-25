# Introduction to the Vorteil Repository (vRepo)

The Vorteil Repository (vRepo) is included with a standard Vorteil installation, and provides a means of app storage and version control. Access to the vRepo is governed by ACLs which can be set and adjusted by the appropriate user, creating a safe environment to selectively share information with other users.

It is possible to add external Vorteil Repositories as connections within a local Vorteil environment for the purposes of pushing/pulling applications to/from a remote environment.As an example, the [official Vorteil Marketplace](https://marketplace.vorteil.io) is a vRepo!

## Default Settings

A default Vorteil Repository configuration is not hosted on an exposed address, and contains no pre-existing buckets or applications. The only access point is via an internal bind address (it is important not to expose this specific endpoint, as it assumes full access to all repository functionality).

Configuring the vRepo to be accessible via an exposed, secure endpoint is easy. Within the Vorteil [service configuration file](../../../faq/host_configuration/where_is_the_config), enable a network-exposed endpoint with the following:

```toml
[[network]]
  disabled = false
  bind = ADDRESS:PORT
  disable-tls = true
  website-files = "/opt/vorteil/web/vrepo"
```

Examples of acceptable bind address syntax:
- eth01:7474
- 0.0.0.0:7474
- 192.168.1.32:7474
- :7474

With this configuration, users would now be able to connect to the Vorteil Repository; but they will not have access to any resources by default. Detailed instructions for granting access to users and groups can be found [here](../../acls/applying_acl_rules).

### Inclusions

By default, a Vorteil installation includes a web interface for the exposed 'vRepo' functionality, as well as for the protected 'admin' functionality. The directory from which web files are being served can be found by observing the ```website-files``` field of the ```[[network]]```  and ```[admin]``` objects in the Vorteil server configuration file.

*Note that [[network]] is encased by 2 sets of square brackets. This is toml syntax for an array; multiple exposed endpoints can be configured.*

## Features

Any user with access to a vRepo has the freedom to run, unpack, or pull directly from the repository. This is done by providing what is referred to as a 'repository url' for the PULLABLE/BUILDABLE/etc. argument of an operation.

*eg. "marketplace:sisatech/helloworld" would refer to the "helloworld" application within the "sisatech" bucket of the "marketplace" repository.*

## See Also
- [Applying ACL Rules to an Object](../../acls/applying_acl_rules)