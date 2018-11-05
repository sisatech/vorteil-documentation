# Kernel Changelogs

Changelogs, feature pipelines, and other information related to the kernel's
version.

## 0.3.3 (2018-11-05)
- Bug fixed where SIGTERM was not handled correctly
- Bug fixed where memory above 3GB could not be addressed
- Bug where virtual filesystem syscall open would modify input path
- Bug fixed where signals could cause 100% cpu utilisation
- Fixed several issues in log redirects

## 0.3.2 (2018-10-22)
- Bug fixed where POLLOUT would be unset
- Bug fixed where poll/select system calls would delete each others notifications
- Bug fixed where ARP logging would print random characters
- Added "poweroff" status to Xen tools
- Bug fixed where sockets would not be deleted in time_wait
- Added "crash" status to Xen tools if kernel throws a panic
- Bug fixed where sockets would not be drained and cause memory leak

## 0.3.1 (2018-10-12)
- Ported FreeBSD network stack
- Xen/EC2 support

## Relationships
- [Compiler: 2.0](../../compiler/2.0)
