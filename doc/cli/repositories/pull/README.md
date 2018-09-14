# Upload an app from the local repository into a remote repository (vorteil repositories pull)

Upload an app from the local repository into a remote repository.

# Usage

```
vorteil repositories pull PULLABLE BUCKET APP [flags]
```

## Options

```
      --compression-level uint   compression level (0-9) (default 1)
  -h, --help                     help for pull
      --ignore-kernel            ignore-kernel
```

## Hidden options

```
      --args string                   arguments used on the app
      --binary string                 path to app's binary on app's filesystem
  -d, --debug                         print debug information to stderr
      --env strings                   add environment variable to app
      --files strings                 
      --icon string                   path to package icon
      --info.author string            name the author of the app
      --info.date string              date of app's release (YYYY-MM-DD)
      --info.description string       provide a description for the app
      --info.name string              name the app
      --info.summary string           provide a short summary for the app
      --info.url string               URL for more information about the app
      --info.version string           identify the app's version
      --mkdir strings                 
      --network.0.gateway string      configure app's network gateway
      --network.0.http strings        expose http port
      --network.0.https strings       expose https port
      --network.0.ip string           configure app's network IP address
      --network.0.mask string         configure app's subnet mask
      --network.0.mtu string          configure app's network interface MTU
      --network.0.tcp strings         expose tcp port
      --network.0.udp strings         expose udp port
      --network.1.gateway string      configure app's network gateway
      --network.1.http strings        expose http port
      --network.1.https strings       expose https port
      --network.1.ip string           configure app's network IP address
      --network.1.mask string         configure app's subnet mask
      --network.1.mtu string          configure app's network interface MTU
      --network.1.tcp strings         expose tcp port
      --network.1.udp strings         expose udp port
      --network.2.gateway string      configure app's network gateway
      --network.2.http strings        expose http port
      --network.2.https strings       expose https port
      --network.2.ip string           configure app's network IP address
      --network.2.mask string         configure app's subnet mask
      --network.2.mtu string          configure app's network interface MTU
      --network.2.tcp strings         expose tcp port
      --network.2.udp strings         expose udp port
      --network.3.gateway string      configure app's network gateway
      --network.3.http strings        expose http port
      --network.3.https strings       expose https port
      --network.3.ip string           configure app's network IP address
      --network.3.mask string         configure app's subnet mask
      --network.3.mtu string          configure app's network interface MTU
      --network.3.tcp strings         expose tcp port
      --network.3.udp strings         expose udp port
      --network.N.gateway string      configure app's network gateway
      --network.N.http strings        expose http port
      --network.N.https strings       expose https port
      --network.N.ip string           configure app's network IP address
      --network.N.mask string         configure app's subnet mask
      --network.N.mtu string          configure app's network interface MTU
      --network.N.tcp strings         expose tcp port
      --network.N.udp strings         expose udp port
      --nfs.0.mount string            configure app's nfs mounts
      --nfs.0.server string           configure app's nfs servers
      --nfs.1.mount string            configure app's nfs mounts
      --nfs.1.server string           configure app's nfs servers
      --nfs.2.mount string            configure app's nfs mounts
      --nfs.2.server string           configure app's nfs servers
      --nfs.3.mount string            configure app's nfs mounts
      --nfs.3.server string           configure app's nfs servers
      --nfs.N.mount string            configure app's nfs mounts
      --nfs.N.server string           configure app's nfs servers
      --peer string                   route requests to this peer of the server
      --redirect strings              redirect file data to a server (on write)
      --server string                 server network address (default "127.0.0.1:7472")
      --system.delay string           seconds to pause before starting app
      --system.disk-cache string      set the size of the app's disk cache
      --system.dns strings            set the DNS server list for the system
      --system.hostname string        set the hostname for the system
      --system.max-fds uint           maximum file descriptors available to app
      --system.output-format string   specify transmission format for redirects
      --system.output-mode string     specify vm output behaviour mode
      --vcfg strings                  
  -v, --verbose                       print additional information to stderr
      --vm.cpus uint                  number of cpus to allocate to app
      --vm.disk-size string           disk image capacity to allocate to app
      --vm.inodes uint                number of inodes to build on disk image
      --vm.kernel string              kernel to build app on
      --vm.ram string                 memory in MiB to allocate to app
```


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 14-Sep-2018
