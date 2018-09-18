# Binary and Arguments

The single most important thing Vorteil needs to know in order to launch an app
is what it should supply to the app as its arguments. 

## Fields

Because of their importance these two VCFG fields are not a part of any other 
subsections. They are the only two fields in the VCFG that aren't a part of 
another section, due to the rules of TOML this means that these fields must 
appear at the top of a VCFG file before any subsections.

### binary 

**Field type: String**

```
binary = "myapp"
```

The binary field serves two purposes. First and foremost, it it passed to the 
application as-is as its zeroth argument. If an application is picky about its
own name or location on the file-system, this is the field that needs to be 
carefully changed to make it behave correctly. 

The other purpose of the field is to locate the binary itself within a 
[project](../../projects/introduction) tree. This is only important if the app 
is currently in the form of a project.

### args

**Field type: String**

```
args = "arg1 arg2 arg3"
```

## Notes

The values defined within these two fields should be combinable in such a way 
that the app could be executed from a Linux shell if they were typed one after 
the other, assuming that the app can be found at the correct path. The example
values above would be equivalent to the following in a shell.

```sh
$ myapp arg1 arg2 arg3
```

## See Also 

* [What is a VCFG?](../introduction)
* [Example VCFG](../example)
* [CLI's Binary and Arguments Documentation](../../../cli/vcfgs/invocation)