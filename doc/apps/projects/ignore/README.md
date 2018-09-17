# Excluding Items From a Project

To avoid confusing situations for the developer, by default a project will not 
exclude any object from within its project tree. This includes the project file 
itself, as well as any other files the developer may want to keep within the 
project tree for the purposes of organization. The 'ignore' field on a project 
file makes it possible for the develoeper to manually exclude items from being 
used on the app's file-system.

The ignore field takes a list of any number of 'glob' patterns and checks each
of them against every item in the project tree. If it finds a match against an 
item, that item is excluded.

Objects are only excluded during the first pass, which occurs before 
target-specific inclusions. This means that items added to the app's file-system
via the 'files' field are not compared against any of the ignore patterns. This
behaviour can be exploited to keep things simple and clean when developing a 
project with multiple [targets](../targets). For excluding a directory called 
'test' and another called 'prod' and then re-including them using their targets 
respective files field.

## Standard Patterns

Whilst no ignore patterns are necessary, it is generally a good idea to exclude
everything that needs to be within the project directory but which does not need
to appear on the app's file-system. Whenever building or packing a project you
may want to check the list of items included in the app. This is not just good 
practice for keeping apps simple and minimizing their size, it's also important 
to prevent easy mistakes such as packing sensitive data into a distributable 
package.

As a starting point, it is generally recommended to begin with the following
ignore settings, and grow the field as necessary.

```
ignore = [".vorteilproject", "*.vcfg"]
```

## See Also

* [What is a Project?](../introduction)