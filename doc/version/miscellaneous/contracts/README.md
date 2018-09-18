# Learn about version compatibility guarantees (contracts)

The version information for both the CLI (client) and server can include
extended information about the sub-modules they provide and the contracts they
support. Indentation indicates that a sub-module or a contract is part of the
greater element. When displaying the version information using the "format" flag
modules and contracts are clearly labelled.

A contract is a pledge that updates up until a certain version will not break
backwards compatibility with existing tools. These contracts are generally
given as partial versions, and are not valid complete Semantic Versioning 2.0.0
version names. This is because a contract uses a lesser degree of precision.

Fully qualified semantic versions (i.e. '*.*.*') without extended information
such as the commit ID represent the versions of sub-modules provided by the
tool, which may or may not satisfy the broader contracts of other modules.

## Examples:

"api: 0.1.5" provided by a server satisfies "api: 0.1.x" required by a client

"kernel: 0.2.x" is satisfied by Vorteil kernel 0.2.24

"package: 1.x.x" is satisfied by version 1.1.12 Vorteil packages

"kernel: 0.3.x" is not satisfied by Vorteil kernel 0.2.24

"kernel: 0.3.x" is not satisfied by Vorteil kernel 0.4.1

# See Also

* [CLI version command](../../../cli/general/version)	 - View client and server version information
