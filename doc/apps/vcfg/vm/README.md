# Compiler and Deployer Virtual Machine Recommendations (vm)

When it comes time to compile a virtual disk image or provision a virtual 
machine, the tools taking care of these processes need to know what resources
should go into it. The 'vm' section of a VCFG is dedicated to information that
comes into play at these times.

## Fields 

Each of the following fields must appear within the vm section, which can be
identified by the appearance of the following line. Everything from that line 
up until the appearance of another section is part of the vm section.

```
[vm]
```

### cpus

**Field type: Unsigned Integer**

```
  cpus = 1
```

The cpus field provides a recommendation as to the minimum number of virtual 
CPUs that should be allocated to a virtual machine running this app.

### ram

**Field type: String**

```
  ram = "64 MiB"
```

The ram field provides a recommendation as to the minimum amount of virtual RAM
that should be allocated to a virtual machine running this app. The field 
accepts plain numbers, which will be interpreted as bytes, or strings which can
include a number followed by a unit, as per the above example.

### inodes

**Field type: Unsigned Integer**

```
  inodes = 1024
```

The inodes field provides a recommendation as to the minimum number of inodes 
that should be added to the file-system during the virtual disk image 
compilation process. This number will very rarely match precisely the number of 
inodes on the image because the compiler performs optimizations to squeeze out 
as many 'free' inodes as possible (extra inodes that add no overhead). This can 
be safely treated as an enforced minimum, however.

### kernel 

**Field type: String**

```
  kernel = "0.3.0"
```

The kernel field tells the compiler which kernel should be used when building 
the virtual disk image. It is up to the compiler to make the final decision as 
to which kernel to use, but the standard compiler requires an exact match.

### disk-size 

**Field type: String**

```
  disk-size = "128 MiB"
```

The disk-size field provides a recommendation as to the minimum size of the 
virtual disk image to create during the disk-building process. The field accepts
plain numbers, which will be interpreted as bytes, or strings which can include
a number followed by a unit, as per the above example.

## Notes

Everything in this section should only ever be treated as a minimum 
recommendation, never a rule. Whilst various compilers or provisioners may 
choose to adhere to these recommendations in varying degrees, they are never 
required to. They couldn't be, because some cloud platforms will not afford you
such fine-grained allocation of resources as others will, and others may 
require you to pad the disk size to be aligned a certain way. 

## See Also

* [What is a VCFG?](../introduction)
* [Example VCFG](../example)
* [CLI's VM Recommendations Documentation](../../../cli/vcfgs/vm)