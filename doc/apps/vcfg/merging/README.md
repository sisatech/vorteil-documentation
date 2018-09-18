# Merging and Splitting Up VCFGs

Vorteil's tools have been designed to support fragmented VCFGs. Generally 
speaking, anywhere a VCFG file is used, it is possible to supply multiple VCFG
files. The tools will merge each VCFG file in turn over the top of each other to
construct a complete VCFG in memory. 

It may be desirable to split the application metadata from the rest of the
configuration to prevent often very long descriptions from cluttering up the
main configuration file. And some applications are distributed best as templates
that do not work without some final bit of configuration that varies from
customer to customer or even from deployment to deployment. The ability to
combine VCFGs enables this useful practice of splitting up configuration
settings.

## See Also 

* [What is a VCFG?](../introduction)
* [CLI's Documentation about Merging VCFGs](../../../cli/vcfgs/merging)
