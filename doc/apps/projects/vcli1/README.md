# Changes from VCLI 1 

If you used VCLI back before it installed as a service, when we called it VCLI
instead of VCLI 2, you may be used to a very different project structure. To 
reduce ambiguity, for the remainder of this article we'll refer to the older 
VCLI as VCLI 1. 

VCLI 1 didn't have an explicitly defined concept of a project, but we did 
recommend some 'best practices' directory layout that we were able to use and 
attempt to automatically resolve for you. This was done with simple naming 
conventions: if your app was called 'myapp', then VCLI 1 would automatically 
detect objects in the same directory begining with the same name and ending with
several suffixes, including 'myapp.files', 'myapp.vcfg', and 'myapp.icon'. If 
you're not familiar with this layout then you should skip this article and 
instead refer to the general information about how projects work today.

## Converting From VCLI 1 to VCLI 2

Since VCLI 1's projects didn't support the more advanced features VCLI 2's does,
converting to bring a project up-to-date is quite simple. Just follow these 
steps. As an illustrative example, for the remainder of this article we'll 
assume you have a VCLI 1 project with the following objects.

```
myapp
myapp.vcfg
myapp.icon
myapp.files/
myapp.files/myapp.conf
myapp.files/database/
myapp.files/database/sqlite.db
```

### Create the Project File

The first thing a VCLI 2 project needs is a project file. Create a new file 
called '.vorteilproject' with the following settings.

```
ignore = [".vorteilproject", "*.vcfg", "myapp.icon", "myapp"]

[[target]]
  name = "default"
  icon = "myapp.icon"
  vcfgs = ["myapp.vcfg"]
  files = []
```

A note about the last ignore string: generally speaking, with Vorteil it is not
necessary to include the app's binary on the app's file-system itself. The 
app binary is stored somewhere else on disk for the purpose of optimization. 
This was true even in VCLI 1, so excluding the app binary from the file-system 
using the ignore field in this example is important for getting an equivalent
project.

### Modernize the VCFG File

VCLI 2's VCFG files are conceptually similar to how they were in VCLI 1, but new
features and improvements to their user-friendliness have required us to make 
backwards incompatible changes. There are no short and simple instructions for 
converting from the old one to the new here, you'll need to understand what the
old one was trying to achieve and learn about how the new VCFGs work so you can
create a VCLI 2 VCFG that achieves your goals.  

Try the CLI's extensive [VCFG documentation](../../../cli/vcfgs/new) to learn 
more.

### Elevate the File-System Files

In VCLI 1 the file-system was treated very separately from the app binary. In 
VCLI 2 we like to keep them in the same file-tree so that the application can be
run effortlessly from a Linux shell with all of the files that would be on their
file-system still in the same relative location. 

To solve the final issue in the coversion process all of the files within the 
'.files' directory need to be elevated one level, and the directory itself must
then be deleted. In our example, this means moving 'myapp.files/myapp.conf' to 
'myapp.conf' and 'myapp.files/database/sqlite.db' to 'database/sqlite.db'.

### Finished Example

After each of the above steps has been undertaken, the project from the original
example should be changed to the following structure.

```
.vorteilproject
myapp
myapp.vcfg
myapp.icon
myapp.conf
database/
database/sqlite.db
```

## See Also 

* [What is a Project?](../introduction)