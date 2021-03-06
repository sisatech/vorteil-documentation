# Cloud Specific Environment Variables 

To simplify things when it comes to running applications on the cloud, the 
kernel probes its networks to gather some extra information. Information it 
finds is made available to the app via environment variables, and these 
environment variables are also substituted into appropriate places within the
app's arguments, where they're present.

It is not necessary to define these environment variables in an app's VCFG to 
access them, but it can be a good practice to do so anyway if access is required
for reasons of transparency and documentation. The correct way to do this is to 
duplicate their keys and their keys as their values in the VCFG. Take the 
HOSTNAME variable as an example. The best practices way to use this in the VCFG
is like so.

```
[env]
  HOSTNAME = $HOSTNAME
```

If you do document the use of a variable in this way it is important to follow
that guideline to the letter. This is because any defined environment variables
that clash with variables the kernel provides will override the values the 
kernel would have set. This can be a useful feature at some times, but it can 
also cause bugs if you're not careful. 

## Kernel-Provided Variables

The following is a list of currently supported environment variables added by 
the kernel. 

* HYPERVISOR
* CLOUD_PROVIDER
* ETH_COUNT
* HOSTNAME
* IP0
* IP1
* IP2
* IP3
* EXT_HOSTNAME
* EXT_IP0
* EXT_IP1
* EXT_IP2
* EXT_IP3

### HYPERVISOR

This variable will have one of the following values, depending upon which 
hypervisor the kernel thinks it is running on.

* KVM
* VMWARE
* XEN
* AZURE
* UNKNOWN

### CLOUD_PROVIDER

This variable will have one of the following values, depending upon which 
cloud provider it believes is hosting it. 

* KVM
* VMWARE
* XEN
* AZURE 
* UNKNOWN
* GCP
* AWS
* HYPERV

### ETH_COUNT

This variable will have the number of connected network interfaces as its value.
You can use this number to determine how many of the IP and EXT_IP variables to
process.

### HOSTNAME and EXT_HOSTNAME

These variables will have the hostname of the virtual machine. The HOSTNAME will
have the value the kernel has set, whilst the EXT_HOSTNAME will have the value
the cloud provider is using, for in the event where these values do not match. 
Both variables will be set to useful values on any environment.

### IP and EXT_IP 

These variables will have the IPv4 address (as a string) for their respective 
network interfaces on the virtual machine. IP will have the value the kernel has
set, whilst the EXT_HOSTNAME will have the value the cloud provider is allowing 
people to connect to from the internet. Both variables will be set to useful 
values on any environment, but only as many as there are network interfaces. 