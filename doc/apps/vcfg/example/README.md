# Example VCFG 

This page presents an illustrative and somewhat exhaustive example of what a 
VCFG can be. To cover all of the bases it touches on almost every possible 
field. Most apps do not need such extensive configuration, and indeed this 
example is contrived. The values are not suitable defaults for any application.

## Example

```
binary = "/usr/bin/path"
args = "-jar myapp.jar"

[env]
  HOME = "/"
  USER = "root"

[[network]]
  ip = "dhcp"
  mask = "255.255.255.255"
  gateway = "192.168.1.1"
  udp = ["1000"]
  tcp = ["8080"]
  http = ["80"]
  https = ["443", "8000"]
  mtu = 1500

[[nfs]]
  mount = "/nfs"
  server = "192.168.1.10"

[redirects]
  "/logs" = "192.168.1.11"

[system]
  disk-cache = "64 MiB"
  dns = ["8.8.8.8", "8.8.4.4"]
  hostname = "localhost"
  max-fds = 1024
  output-mode = "standard"
  output-format = "standard"

[info]
  name = "My App"
  author = "Mr Smith"
  summary = "MyApp that does stuff"
  description = """# MyApp

that has a markdown description!"""
  url = "vorteil.io"
  date = "18 Sep 18 11:22 AEST"
  version = "1.0.0"

[vm]
  cpus = 1
  ram = "128 MiB"
  inodes = 1024
  kernel = "0.3.0"
  disk-size = "64 MiB"

```

## See Also

* [What is a VCFG?](../introduction)
* [Binary and Arguments](../args)
* [Environment Variables (env)](../env)
* [Descriptive Metadata (info)](../info)
* [Network Settings (network)](../network)
* [Advanced Runtime Settings (system)](../system)
* [File-System Redirects (redirects)](../redirects)
* [NFS Settings (nfs)](../nfs)
* [Compiler and Deployer Virtual Machine Recommendations (vm)](../vm)