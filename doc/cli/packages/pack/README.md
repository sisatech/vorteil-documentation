# Create a Vorteil package (vorteil packages pack)

Create a new Vorteil package from any valid source germ including projects,
elf64 binaries, other packages, and remote repositories. Vorteil packages are
compressed and optimized archives containing all of the information needed to
construct a Vorteil virtual disk image. They generally represent an immutable
application that can be expected to operate identically on all supported
hypervisors, and they can include metadata and information that helps to
identify it and explain its purpose and its use.

Vorteil packages themselves are valid source germs for all commands that require
one, such as the run and build commands. Since packages contain all of their
dependencies and configuration information they are a convenient way to create a
snapshot of an application for backup or distribution.

## Application Germs:

A germ is an initial stage from which something may develop, and we use the
term to refer to the various types of input that can be developed into a
Vorteil application.  There are four different types of Vorteil germs:
packages, repositories, projects, and binaries.

### Packages:

The Vorteil compiler has been optimized to produce virtual disk images from
Vorteil package files ('.vorteil' files). A package germ's germ-string is a
filepath to a valid Vorteil package available on the filesystem.

### Repositories:

Repositories are vRepo servers managing a version-controlled archive of
Vorteil application packages. Since they are capable of providing access to
packages which are themselves valid germs, applications on vRepo servers
have been made into acceptable germs. A valid repository germ-string must
identify the source's vRepo server and provide a URI to a Vorteil
application stored therein. The format of this germ-string is
SERVER::BUCKET/APP[/REF]. The REF URI element is optional, and if left
unspecified is assumed to be the latest version.

### Projects:

Projects are organized collections of loose files that are destined to
become a Vorteil package in the future. A project germ's germ-string is a
filepath to an existing directory containing a valid Vorteil project.

### Binaries:

Compared to a project, the more ad hoc binary germ (sometimes referred to as
an elf64 germ) allows for disorganized collections of loose files to be used
as a valid germ. A binary germ's germ-string is a filepath to an existing
64-bit ELF binary file. Since binary germs are usually disorganized it is
the responsibility of the user to explicitly specify every necessary
configuration and file-system inclusion.

As a holdover from the first generation implementation of VCLI the modern
CLI will automatically check for a valid configuration file and file-system
contents directory adjacent to the binary itself and with identical file
names but for the addition of the '.vcfg' and '.files' suffix, respectively.
If it finds either of these it will automatically include them as a part of
the germ without additional instruction.

## Just-In-Time Germ Injection:

Part of Vorteil's power and flexibility comes from its ability to inject
changes to a germ immediately before its germination begins. Typical uses for
this capability might require adding to or overwriting the various
configuration files on the application's file-system, or overwriting the
germ's configuration settings as found in VCFG files such as its hostname or
disk size.

### Injecting objects onto the file-system:

Wherever germs are used an opportunity to provide a list of file-system
inclusions is available. Using the CLI this comes from an optional and
repeatable 'files' flag. This flag is generally hidden from basic help
information due to its ubiquity, but should be available for every command
where a germ string is used as an argument. The value(s) for the flag should
be filepaths to existing objects on the local file-sysem. If the target is a
file it will be included onto the root directory of the germ's file-system.
If the target is a directory all of its contents will be merged onto the
germ's file-system recursively, anchored onto the root directory instead:
i.e. the target's top-level contents will become top-level contents of the
germ's root directory, but the target itself will not be added to the germ's
file-system.

File-system injections are smart enough to recursively produce parent
directories as required. Any type of collision other than injecting one
directory over an existing directory will result in the germ's original
conflicting contents being completely displaced by the injected content.
Adding a directory object to a path on the germ's filesystem that already
contains a directory will result in a recursive merge operation whereby the
contents of the injected directory are each themselves injected into
expected locations, resulting in further merges and displacements as
necessary.

### Injecting configuration adjustments:

Wherever germs are used an opportunity to provide a list of overriding
configuration files is available. Using the CLI this comes from an optional
and repeatable 'vcfg' flag, as well as a complete set of optional
fine-tuning flags that mirror the syntax of the VCFG data structure. These
flags are generally hidden from basic help information due to their ubiquity
and their quantity, but should be available for every command where a germ
string is used as an argument. The value(s) for the 'vcfg' flag should be
filepaths to existing VCFG files on the local file-system. The values for
the various fine-tuning flags mirror values that would be acceptable within
VCFG files. The fine-tuning flags have names that reflect their flattened
name within the VCFG file, e.g. to adjust the 'kernel' field, which is a
subfield of the 'vm' structure, the flag is 'vm.kernel'. A comprehensive
list of these flags can be read from the help information on "vcfg new".

# Usage

```
vorteil packages pack [flags] [PACKABLE]
```

## Aliases

```
package, new, create, build, make
```

## Options

```
      --compression-level uint   compression level (0-9) (default 1)
  -f, --force                    force overwrite of existing files
  -h, --help                     help for pack
      --ignore-kernel            ignore-kernel
  -o, --output string            path to put output file
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


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 17-Sep-2018
