# How to Configure the Vorteil Runtime

Configuration of Vorteil's runtime is handled entirely by modifying an app's 
configuration file [(VCFG)](../../../apps/vcfg/introduction). 

Some tools may hide this process from you, such as when the CLI makes one-time 
changes to the VCFG when requested via [flags](../../../cli/vcfgs/flags). At 
other times some VCFG values will be overwritten by the compiler at its 
discretion. The final VCFG will always be displayed in the verbose build logs, 
and so the values in these logs are what should be checked when things aren't 
behaving as expected, since they may not match the values found in your VCFG 
file exactly.

Just about every field in a VCFG is relevant to the runtime, with the exception 
of fields under the [info](../../../apps/vcfg/info) and 
[vm](../../../apps/vcfg/vm) sections, and the port fields under the 
[network](../../../apps/vcfg/network) sections, all of which are instead used 
by tools to manage or describe apps from without.

