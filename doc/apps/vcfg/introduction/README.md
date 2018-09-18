# What is a VCFG?

The VCFG is a term we use to refer to both an app's configuration and the TOML
file that stores such information. Such files typically end in the ".vorteil"
file extension. It is an abbreviation and contraction of the words "Vorteil Configuration", and it is encompasses the entire range of configurable settings
for a Vorteil app, from descriptive metadata to all of the kernel's configurable
runtime settings. 

Vorteil will try its best to apply useful default values over any undefined 
fields. In many cases it is nearly possible to get away without worrying about 
the VCFG at all. Small VCFGs are preferred, and in general one shouldn't define 
a setting unless it needs to have a value that isn't being supplied by default.

## Sections

VCFGs are broken up into a number of sections to keep them clear and organized.
Rather than keep a complete and monolithic description of the VCFG structure 
here, we've broken down these sections into separate articles.

* [Binary and Arguments](../args)
* [Environment Variables (env)](../env)
* [Descriptive Metadata (info)](../info)
* [Network Settings (network)](../network)
* [Advanced Runtime Settings (system)](../system)
* [File-System Redirects (redirects)](../redirects)
* [NFS Settings (nfs)](../nfs)
* [Compiler and Deployer Virtual Machine Recommendations (vm)](../vm)

## See Also 

* [Example VCFG](../example)
* [Merging and Splitting Up VCFGs](../merging)
* [Tom's Obvious, Minimal Language (TOML)](https://github.com/toml-lang/toml)