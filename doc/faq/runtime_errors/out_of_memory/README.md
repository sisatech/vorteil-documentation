# Panic: out of memory. no pages on stack

# Description

Your Vorteil application terminates unexpectedly. In the serial logs shortly before the application stopped, appears a panic level message out of memory. no pages on stack?.

# What is going wrong

Vorteil VMs commonly run with substantially lesser resources than the servers running fully-featured operating systems they might have run on otherwise. It is not unreasonable for some applications to need no more than 128 MiB of ram or less. Sometimes these memory allocations are set too frugally, causing the VM to run out. 

_Version 0.3 of the Vorteil kernel does not support memory swapping._

# Solution

Increase the amount of memory provided to the virtual machine. If you didn't configure the virtual machine manually, the problem can usually be resolved by [adjusting the recommended memory](https://vorteil.kayako.com/article/28-memory).

  