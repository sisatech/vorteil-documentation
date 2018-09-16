# Learn how to use flags to modify app configuration (vorteil vcfgs flags)

In every context where packages or disk images are created we support the
ability to make configuration changes via flags. These flags are numerous and
ubiquitous, and so they are hidden from general help output in most cases unless
the verbose (-v, --verbose) flag is provided.

## Merging VCFG Files:

--vcfg strings	This is a repeatable flag that can be used to specify
the file paths to existing vcfg files for merging in the
order that they are listed. It is also possible provide
a URL instead of a file path, where the server can send
a GET request and retrieve the file. To interpret the
value as a URL the provided string must begin with the
protocol ("http://" or "https://").

See the "merging" additional help topic for more information about merging.

## Adjusting Configuration:

Every field within a VCFG has a sister flag that can be used to modify the
value from the CLI. The keys for each of these flags is closely related to the
structure of the VCFG itself to keep things simple and clear. For fields that
can be defined with a single string the flag key is the same as the field name
in the VCFG appended after the key(s) of its parent(s) field(s), separated by
periods.

For repeatable VCFG structures such as network configuration the index for the
applicable structure is added as another part of the flag key after the name
of the repeatable structure, separated by periods. Remember that these indices
begin the count at zero.

For maps, i.e. fields within the VCFG that are defined by the user, the values
given for the flag must be in the form of KEY=VALUE. The "env" and "redirect"
fields are two examples of this.

### Examples:

--vm.kernel=0.3.0		→	[vm]
kernel = "0.3.0"

--network.0.ip=dhcp 	→	[[network]]
ip = "dhcp"

--env="HOME=/users/root" 	→	[env]
HOME = "/users/root"


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 14-Sep-2018
