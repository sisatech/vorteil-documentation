# Learn about binary, args, and environment variables (vorteil vcfgs invocation)

Most applications require some arguments or environment variables to run. In
addition, Vorteil needs to be able to determine which file is the application
binary that should be executed within a project, and it does this by probing the
VCFGs for the "binary".

The binary field will also automatically be used as the zeroth argument to the
application, but is separated from the rest of the arguments in the vcfg due to
its significance.

Here is the skeleton of these fields, as well as the associated
flags to change these values directly from the commandline.

binary = ""
args = ""

[env]
KEY=VALUE

If the binary is provided as an absolute path then the zeroth argument to the
application will be an absolute path, but otherwise the path will still be
treated as relative. This means that a project for which the binary is "/app"
will look for the application binary on the top level of the project directory,
not at the root directory of the developer's file-system.

The environment variables are provided as a map with any number of KEY=VALUE
items defined.

Note: the binary and args fields are not sub-fields of any other section, and
must therefore appear at the beginning of a VCFG before any other
sections.

## Example:

binary = "/jdk1.8.0_172/bin/java"
args = "arg1 arg2 arg3"

[env]
HOME = "/root"
X = "x"

### which is equivalent to:

--binary="/jdk1.8.0_172/bin/java" --args="arg1 arg2 arg3" \
--env="HOME=/root" --env="X=x"


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 14-Sep-2018
