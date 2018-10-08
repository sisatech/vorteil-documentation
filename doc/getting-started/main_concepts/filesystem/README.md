# Filesystem
The filesystem of a Vorteil application can comprise of one or many files or directories. At runtime, if an application is configured to use multiple sources, they will be merged in to one single filesystem. 

Shared objects that the application depends on can all be included within the directory specified to become the virtual machine's filesystem!

Vorteil tools support specifying files/folders to merge in to the filesystem at runtime (when building/running/pushing/pulling/provisioning an app).