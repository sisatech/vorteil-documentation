# Compiler Changelogs

Changelogs, feature pipelines, and other information related to the compiler's 
version.

## Relationships
- [Packager: 2.0](../../packages/2.0)
- [Kernel: 0.3](../../kernel/0.3)
## 2.0.2 (2018-11-19)
- Added NTP functionality to the VCFG.
## 2.0.1 (2018-09-26)
- Will now automatically try to guess sensible values for the number of inodes and the size of the disk if the respective fields are left undefined in the VCFG. 
- Will now force in optimal MTU values for Google Cloud Platform when compiling GCP image archives, overriding and disregarding any value defined in the VCFG.
- Fixed a bug that could sometimes cause the compiler to fail when trying to build GCP image archives because the forced 1 GiB disk size alignment wasn't always sticking.