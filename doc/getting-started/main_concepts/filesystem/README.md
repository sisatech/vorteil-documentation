# Filesystem
The filesystem of a Vorteil application can comprise of one or many files or directories. At runtime, if an application is configured to use multiple sources, they will be merged in to one single filesystem. 

The merging of two filesystems is similiar to copying a directory of files on Windows into another directory. Which means it will merge if possible if not it will replace that file currently on the filesystem. Below is a picture of the expected outcome when merging two filesystems together.

![Merging Filesystems](https://storage.googleapis.com/vorteil-dl/assets/documentation/mergedscreenshot.png "Two Filesystem Merge")

A quick example of merging two file systems together is shown below

```
vcli run marketplace:sisatech/helloworld --files ~/go/src/etc/ -- files ~/go/src/lib --files ~/go/src/lib64
```
The example above shows a CLI command merging certificate and shared library files on runtime with an application that is currently being hosted on the marketplace. Isn't this awesome!

Shared objects that the application depends on can all be included within the directory specified to become the virtual machine's filesystem! The shared objects are stored in a lib and lib64 directory depending on what the object is for.

Vorteil tools support specifying files/folders to merge in to the filesystem at runtime (when building/running/pushing/pulling/provisioning an app).

## See Also
- [Filesystem Redirects](../../../apps/vcfg/redirects/)
- [Network File System](../../../apps/vcfg/nfs/)
- [Merging VCFGs](../../../apps/vcfg/merging/)