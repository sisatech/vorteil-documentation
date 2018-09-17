# Learn about how vcfgs are divided and merged (vorteil vcfgs merging)

In every context where VCFG files are used we support the use of multiple VCFGs
at once. VCFGs apply their adjustments to a process in a predictable order:

1.	Server defaults
2. 	Source configuration
3. 	First VCFG file
...
N-1. 	Last VCFG file
N. 	VCFG flags

Server defaults are generally sensible default values that are useful for most
applications. The source configuration refers to any configuration that already
exists within the source files, such as within a Vorteil package. Under the
hood, the VCFG flags are merely another VCFG file that is discarded after use.

When merging VCFGs, fields that are given in the incoming VCFG completely
overwrite data in the same fields of the incumbent VCFG. Even if the incoming
field has them empty. To prevent a VCFG from overwriting settings in other VCFGs
unintentionally, completely remove all fields that aren't required from it.

It may be desirable to split the application metadata from the rest of the
configuration to prevent often very long descriptions from cluttering up the
main configuration file. And some applications are distributed best as templates
that do not work without some final bit of configuration that varies from
customer to customer or even from deployment to deployment. The ability to
combine VCFGs enables this useful practice of splitting up configuration
settings.


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 17-Sep-2018
