# CLI Changelogs

Changelogs, feature pipelines, and other information related to the CLI's 
version.

## Relationships
- [API: 0.1](../../api/0.1)

## 2.0.1 (2018-09-26)
- Now correctly handles relative paths as project arguments with build variants specified (eg. './myProject:dev').
- The 'vcfgs lint' command no longer erroneously returns a 'vcfg does not exist' error when provided a relative path.
- The 'pack' command no longer produces '..vorteil' when working from the current working directory (no arg provided).
- Improved readability of VM status logs (in the run command) by colouring all VM status log lines in blue and filtering severity out according to the use of debug/verbose flags.