# Structure of a Project

At its most basic, a project need only be a directory containing an application
binary and a minimal ".vorteilproject" file identifying it. But the project file
can be expanded to do some very helpful and powerful things.

Generally speaking, the project directory itself will mirror the entire contents
of an app's file-system as if it was the root directory. In addition to being
intuitive, this has the side-benefit of allowing the developer to test their
ELF64 binary in their Linux shell and expect in many cases to see the same
behaviour from their app when it runs within a virtual machine.

## The /etc/hosts file

The vorteil kernel generates /etc/hosts on the fly based on the .vcfg file.
In some instances it is required to provide a /etc/hosts file to the application.
Therefore the kernel checks if a /etc/hosts file exists and uses it instead if
it has been manually provided.

## The Project File

Every project must have a project file, which is a file named ".vorteilproject".
A project file is a TOML file that contains settings and rules governing the use
of files and directories within the project directory. The basic template for a
project file is as follows.

```
ignore = []

[[target]]
  name = ""
  icon = ""
  vcfgs = []
  files = []
```

The ignore field is optional, as are most sub-fields on a target. The minimal
working project file will generally look something like this:

```
[[target]]
  vcfgs = [".default.vcfg"]
```

## See Also

* [What is a Project?](../introduction)
* [Excluding Items From a Project](../ignore)
* [Project Targets](../targets)
