# CLI Changelogs

Changelogs, feature pipelines, and other information related to the CLI's 
version.

## 2.0.1
- Now correctly handles relative paths with build variants specified (eg. './myProject:dev').
- The 'vcfgs lint' command no longer erroneously returns a 'vcfg does not exist' error when provided a relative path.
- The 'pack' command no longer produces '..vorteil' when working from the current working directory (no arg provided).