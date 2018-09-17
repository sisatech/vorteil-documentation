# Learn about an app's VM configuration (vorteil vcfgs vm)

The "vm" configuration section contains settings that provide recommendations to
the Vorteil image compiler as well as any provisioning and deploying stacks.
It is up to the tool to decide what it should do with these recommendations, and
the developer cannot necessarily be certain that the values provided in this
section were used as-is.

Here is the skeleton of a vm configuration, as well as the associated flags to
change these values directly from the commandline.

```
  [vm]
    cpus = 0			→	--vm.cpus=0
    disk-size = ""		→	--vm.disk-size=""
    inodes = ""			→	--vm.inodes=""
    kernel = "" 		→	--vm.kernel=""
    ram = ""			→	--vm.ram=""
```

## Virtual Hardware:

The number of CPUs and quantity of RAM to assign to the virtual machine can be
specified here. These fields are mere recommendations only because not all
cloud environments provide fine-grained customization of the virtual hardware
available. It is generally reasonable to assume that the deployer will take
these recommendations as a bare minimum.

The cpus field takes any positive integer, whilst the ram field can also take
a positive integer representing the number of bytes of memory, but more
commonly should take a string containing a number and units, e.g. "128 MiB".

## Virtual Disk:

The Vorteil image compiler takes two recommendations from this section with
regards to the virtual disk: inodes and disk-size. The disk-size can take a
positive integer representing the number of bytes, but more commonly should
take a string containing a number and units, e.g. "128 MiB". The reason
disk-size is merely a recommendation and not a guarantee is because various
virtual disk image formats require different disk size alignments. It is
always safe to assume that the compiler will produce an image at least as
large as the recommendation here.

The number of inodes tells the Vorteil image compiler how many inodes the
image must be capable of housing as a bare minimum. To save the developer the
hassle and minutia of optimizing this value, the compiler will automatically
scale the given number up to the biggest value that costs no additional disk
overhead. The inodes field can take a positive integer or a string containing
a number and units, e.g. "4K".

## Kernel:

This field specifies which version of the Vorteil kernel the image compiler
should use. It is baked into Vorteil packages to facilitate immutable apps
and to keep a historical record of which kernel an app was (hopefully) tested
on.

By default VCLI will always attempt to use exactly the kernel version
specified in this field, but other tools may employ different policies in the
future, such as guaranteeing the same major version but always forcing an
upgrade to the latest minor version and patch.

## Example:

```
  [vm]
    cpus = 1
    disk-size = "64mb"
    inodes = "1k"
    kernel = "0.3.0"
    ram = "128m"
```

### which is equivalent to:

```
  --vm.cpus=1 --vm.disk-size=64mb --vm.inodes=1k --vm.kernel=0.3.0 --vm.ram=128m
```



###### Auto generated for CLI 2.0.0-7169db7d-dirty on 17-Sep-2018
