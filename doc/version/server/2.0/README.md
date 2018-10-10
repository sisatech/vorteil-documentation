# Server Changelogs

Changelogs, feature pipelines, and other information related to the server's 
version.

## Relationships
- [API: 0.1](../../api/0.1)
- [Packager: 2.0](../../packages/2.0)
- [Compiler: 2.0](../../compiler/2.0)

## 2.0.2 (2018-10-10)
- Modified the Cockroach service file on OSX to include the "--host" flag. This reduces the risk of any issue occuring in the case that 'localhost' can not be resolved.

## 2.0.1 (2018-09-26)
- Hostnames are no longer trimmed 1 character shorter than expected when building a disk.