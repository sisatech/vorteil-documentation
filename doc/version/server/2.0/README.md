# Server Changelogs

Changelogs, feature pipelines, and other information related to the server's 
version.

## Relationships
- [API: 0.1](../../api/0.1)
- [Packager: 2.0](../../packages/2.0)
- [Compiler: 2.0](../../compiler/2.0)
- 
## 2.0.7 (2018-11-19)
- Added more logging for provisioning to Google Cloud Platform and Amazon.
## 2.0.5 (2018-11-05)
- Changed default directory name for 'drop-path' field (within 'kernel-sources') from 'drop' to 'watch'.
- If an error occurs while attempting to fetch kernels from a remote source, up to 5 subsequent attempts will be made.
- (LINUX) Added 'network-online.target' as a service dependency. This ensures that the service will not attempt to perform actions such as retrieving kernels from remote sources before the network is available for use.

## 2.0.4 (2018-10-22)
- Can now build to 'xva' disks.
- Smarter automatic disk size and inode counts. The fields can be used the same way as before, but now support a plus sign in the field to indicate that the disk should be automatically sized to have at least some amount of spare disk space. e.g. --vm.disk-size="+100 MiB". Leaving the field blank will now be equivalent to "+0 MiB", so the automatically sized disks will be as small as they can be whilst still compiling.
- Improved 'kernel sources' logic. Now supports multiple remote kernel sources, and watching a local directory for kernels.
- 'importSharedObjects' API now correctly imports all shared objects; previously it had not acted recursively on the direct dependencies of an app within a project.
- Implemented a 'local' solution for database data. This means that it is no longer necessary to install Cockroach DB to run Vorteil.
- (WINDOWS) - Vorteil will no longer fork a process to run an instance of Cockroach DB.
- (WINDOWS) - Now uses Internet Explorer proxy settings.
- (WINDOWS) - Changed some default paths within the server config file.
- Installer now offers a 'minimal' installation, which will exclude the Developer Studio GUI.
- Installer can be run in 'unattended' mode by running with the "--mode unattended" flag.
- Extended functionality of server logging; now supports log rotation.
- Platform managers now check if the name exists before provisioning a new virtual machine
- VSphere will now re-authenicate if it times out
- GCP Platform Manager now restores the temp directory correctly
- Fixed a bug that occured when a logger wasn't set for a manager
- Improved readabilit of logs

## 2.0.3 (2018-10-12)
- Improved functionality of the 'importSharedObjects' API, which now acts recursively on any libraries that the target depends on.

## 2.0.2 (2018-10-10)
- Modified the Cockroach service file on OSX to include the "--host" flag. This reduces the risk of any issue occuring in the case that 'localhost' can not be resolved.

## 2.0.1 (2018-09-26)
- Hostnames are no longer trimmed 1 character shorter than expected when building a disk.