# Project Targets 

Projects may define multiple 'targets'. If you are familiar with Makefiles, you 
can think of these targets as similar in purpose to a a make target. Basically,
they make it possible to define multiple variations on a project without needing
to duplicate the entirety of the project tree. This can be particularly useful 
if the app needs to be tested with different configurations or files depending 
upon the environment or the stage in its testing. 

## Fields

### Name 

A target may optionally have a name, which can be used to identify it. Targets 
should always have names in multi-target projects otherwise they may be 
impossible to reference. When resolving the target within a project, if no 
specific target is requested the first one to appear in the project file will 
be used. If a specific target is requested it will be matched to the first 
target in the project file that has the required name. Failing to give all 
targets names or defining multiple targets with the same name is never a 
critical error, but it may result in unexpected behaviour if it is done 
accidentally.

Names should be simple strings without spaces, or any other character that isn't
alphanumeric, an underscore, or a hyphen.

### Icon 

A path to an icon picture file to represent the app when it finally gets 
packaged. 

### VCFGs 

A list of paths to configuration (VCFG) files that should be used for the 
target. The VCFGs will be merged in-order as required. Every target requires at 
least one VCFG path here because every target needs enough VCFG information to 
resolve the location of the app binary from within the project tree.

### Files

A list of paths to additional files or directories that should be included on 
the app's file-system. This list will be merged in-order as required, if any 
conflicts occur the latest values in the list take precedence. This field would 
generally be used to either include files from outside of the project tree or in
combination with the [ignore](../ignore) settings to improve the organization of
the project tree.

## Requirements

Every project must define at least one valid target. The only thing a target
needs to be valid is a resolvable app binary. Tools attempt to resolve an app 
binary by processing the configuration for a target from its 'vcfgs' field and 
then confirming that the VCFG's 'binary' points towards a valid ELF64 binary 
within the project tree.

Note that since the binary field within a VCFG is also used as the zeroth 
argument to the app itself and since some apps are picky about where they're 
executed from, it is acceptable to set the binary field within the VCFG to 
appear as if it is an absolute path. For the purposes of resolving the binary 
absolute paths are treated as relative paths, relative to the project directory
itself. Tools will never attempt to resolve a binary from outside of the project
tree. This is a unique behaviour that does not apply to any other file-paths 
provided in a project file: for example, an absolute path to an icon will cause
the project to grab the icon from outside of the project tree when packing.

The double box brackets around the 'target' in the [project file](../structure)
signify that the entire target section is repeatable. Multiple targets may be 
defined for a single project. The name and icon fields take string values, and 
the vcfgs, and files fields may take any number of string values as an array.

## See Also 

* [What is a Project?](../introduction)