# What is a Package?

A package is a convenient way to distribute apps for use. They are 
optimized for performance, and are the preferred format for just about every situation. They are [packable](../../general/packable), 
[buildable](../../general/buildable), 
[provisionable](../../general/provisionable), 
[runnable](../../general/runnable), and 
[unpackable](../../general/unpackable). A package is an immutable and 
reproducible app, capable of being built into identical disk images 
every time.

## Projects as Arguments

Wherever the CLI or API asks for a packable, buildable, provisionable, runnable, or unpackable argument represented by a string, a package may be provided. The string to use directs the tool to the location of the package on the local file-system. For some APIs it may be possible to 
upload a package directly as an argument instead, making packages the 
ideal method for getting something built or run on remote infrastructure.

For the API it is important to remember that a file-path should be an absolute path because the server cannot know which directory you're working from. The CLI on the other hand will convert relative paths to absolute ones automatically before sending a request to a server, so relative file paths are acceptable for use with the CLI.

## See Also 

* [Starting a new package with the developer studio](../../../studio/projects/new)
* [Creating a new package manually](../creation)