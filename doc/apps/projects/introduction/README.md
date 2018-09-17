# What is a Project?

A project is a convenient way to structure files when developing and testing 
apps before they are ready to be packaged. Projects can be used in the place of 
packages in just about every situation, being 
[packable](../../general/packable), [buildable](../../general/buildable), 
[provisionable](../../general/provisionable), 
[runnable](../../general/runnable), and even 
[unpackable](../../general/unpackable). They are effectively a directory that 
can act like transparent, uncompressed packages which are modifiable by users.

Despite these obvious advantages, the ultimate goal of every project is always 
to [become a package](../../packages/creation). Packages are easier to 
distribute, transport, and archive, they're locked down and reproducible, 
they're compressed, and they're optimized to get the best possible performance 
when building. Projects should not be used as a final product, only ever as an
intermediate step on the way to a [package](../../packages/introduction).

## Projects as Arguments

Wherever the CLI or API asks for a packable, buildable, provisionable, runnable,
or unpackable argument represented by a string, a project may be provided. The 
string to use directs the tool to the location of the project on the local 
file-system. Optionally, this file-path may include a [target](../targets) 
directive by appending a colon followed by the target name after the path. 

For the API it is important to remember that this path should be an absolute 
path because the server cannot know which directory you're working from. The CLI
on the other hand will convert relative paths to absolute ones automatically 
before sending a request to a server, so relative file paths are acceptable for 
use with the CLI.

## See Also 

* [Starting a new project with the developer studio](../../../studio/projects/new)
* [Creating a new project manually](../structure)