# CLI Changelogs

Changelogs, feature pipelines, and other information related to the CLI's 
version.

## Relationships
- [API: 0.1](../../api/0.1)

## 2.0.5 (2018-11-05)
- Virtual machines being run/provisioned via the CLI will no longer had an empty 'id' field.
- Added the functionality to show a VCFG file of a package. Using the command vorteil packages vcfg PACKAGE [-FLAGS].
  
## 2.0.4 (2018-10-22)
- Resolved an issue which caused port mapping information not to be displayed when running a virtual machine, even when alternative ports were chosen due to configured ports being in use.

## 2.0.2 (2018-10-10)
- 'images build' no longer ouputs a disk named "." if no arg was provided (ie. running directly from project directory)

## 2.0.1 (2018-09-26)
- Now correctly handles relative paths as project arguments with build variants specified (eg. './myProject:dev').
- The 'vcfgs lint' command no longer erroneously returns a 'vcfg does not exist' error when provided a relative path.
- The 'pack' command no longer produces '..vorteil' when working from the current working directory (no arg provided).
- Improved readability of VM status logs (in the run command) by colouring all VM status log lines in blue and filtering severity out according to the use of debug/verbose flags.