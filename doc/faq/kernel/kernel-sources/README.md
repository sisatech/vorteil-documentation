# Kernel Sources

A default Vorteil installation is configured to observe all publicly released kernel versions. It is possible to  configure Vorteil to watch multiple additional kernel sources via the settings defined in the server configuration file.

```toml
[kernel-sources]
  directory = "/home/exampleuser/.vorteild/kernels"
  drop-path = "/home/exampleuser/.vorteild/kernels/watch"
  remote-repositories = ["https://downloads.vorteil.io/system"]
  upgrade-strategy = "none"
```

- **directory**
  
  The directory within which Vorteil will store kernels.

- **drop-path**
  
  A directory which Vorteil will observe for any additional kernels.

- **remote-repositories**

  An array of URLs pointing to remote kernel sources. Public-release kernels are sourced from 'https://downloads.vorteil.io/system'. Release-candidate kernels are sourced from 'https://downloads.vorteil.io/system-release-candidates'.

- **upgrade-strategy**
  
  Upgrade the kernel specified within an app, to the latest version of kernel as specified by the upgrade-strategy field.

  - Acceptable values: 
    - major
    - minor
    - patch 
    - none

## See Also
- [Where is the service configuration file?](../../host_configuration/where_is_the_config)