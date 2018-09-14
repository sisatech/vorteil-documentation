# PANIC: can not find linker 

## Symptoms

The kernel panics almost immediately, and you see the following log written to the application's output:

        can not find linker: /lib64/ld-linux-x86-64.so.2

## Cause 

Programs compiled against dynamically linked libraries are not complete programs. On their own, they do not work. When you launch them the kernel is tasked with  assembling the complete runtime binary in memory by "linking" relevant parts of dynamically linked libraries. To do this, the kernel needs to find the "linker" on the file-system, which is not included there by default. This panic occurs when the application needs linking but the linker cannot be found on the file-system.

## Solution

### Add the missing linker

Copy the missing linker onto the filesystem for the Vorteil app. If the error message says "can not find linker: /lib64/ld-linux-x86-64.so.2", then copy that file and [add it to your Vorteil app](../../../apps/structure/filesystem). Don't forget that you'll also need to copy the libraries that will be dynamically linked there as well, otherwise you will encounter [a different error](../../runtime_errors/missing_shared_object).

### Statically compile the program

The other approach to resolving this issue is to recompile the program "statically", which means that everything it needs from shared objects is built directly into the binary itself, making linking completely unnecessary. The method for statically compiling an application varies from project to project and is beyond the scope of this article. 