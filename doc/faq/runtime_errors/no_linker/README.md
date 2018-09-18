# Panic: can not find linker

##**Description**

If your application is compiled against dynamically-linked libraries, you may see this error almost immediately on the serial output of the Vorteil VM, and your application will fail to run.

##**What is going wrong**

Programs compiled against dynamically-linked libraries are not complete programs. On their own, they do not work. When you launch them they ask the kernel to assemble the complete runtime binary in memory by "linking" relevant parts of dynamically-linked libraries. To do this, the kernel needs to find the "linker" on the file-system, which is not included there by default.

##**Solution**

There are two common solutions to this problem:

Solution 1: statically compile the application. With this approach no linking is done, so the error will disappear.  

  Solution 2: copy the missing linker onto the filesystem for the Vorteil app. If the error message says "can not find linker: /lib64/ld-linux-x86-64.so.2";, then copy that file from your local filesystem into a folder that is being built onto the virtual disk (!!! see disk contents link !!!). Don't forget that you'll also need to copy the libraries that will be dynamically linked there as well, otherwise you will encounter a different error: error while loading shared libraries (!!! panic: error loading shared libraries link !!!).