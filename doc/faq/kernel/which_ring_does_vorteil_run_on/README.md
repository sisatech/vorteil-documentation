# Which ring does Vorteil run in?

Before answering the question, it is critical to understand the history of rings. The information cited by the Wikipedia [Protection Rings](https://en.wikipedia.org/wiki/Protection_ring) article contains enough information for the reader to familiarise himself/herself with the concept.

Important to note is the information on [Hypervisor Mode](https://en.wikipedia.org/wiki/Protection_ring#Hypervisor_mode). Because Vorteil only runs on hypervisors (and not bare metal), the intermediate ring-0 created by the hypervisor is critical to the security constraints applied to Vorteil on cloud platforms.

So how does Vorteil treat rings:

_The app and the kernel are running in the same ring, which is ring 0. That sounds insecure at first, but keep the following in mind:_

_1. Ring-0 on a virtual machine is not ring-0 on bare metal (guest kernel space). There is still the hypervisor in between.___

_2. Vorteil only allows a single process to run, and it is using only your provided libraries (if any). The user knows exactly what is running on the box with no random shared objects, files etc._

_3. Within Vorteil, the kernel space (ring-0) contains basically nothing, because there are no other processes which could leak and no critical files. It is an app in user space._

_4. Vorteil still randomises memory, so even if a malicious attacker manages to control the virtual machine, the attacker would still need to guess where the stacks are!_

_The question to ask though,__what means controlling the virtual machine? How do you get to the app or libraries?_

_With the Vorteil approach the developer controls few aspects of the virtual machine (except for the application residing on it). Its worth noting that a fundamentally insecure application is vulnerable to attacks, regardless of the platform it is running on. However, if an attacker can buffer overflow the application, it might crash, but that is pretty much everything ..._

___With extended support for 4kb pages,  this will be increasingly difficult to achieve.”_

More information from [StackExchange](https://security.stackexchange.com/questions/129098/what-is-protection-ring-1)

  The rings nomenclature (0-3) you usually see these days started with the requested privilege level field in segment selectors as part of the design of x86 protected mode.

Back in the day, it was possible to make exclusive sections of the memory space called segments. In real mode it was necessary since you only had 20-bit addressable memory. When protected mode came along it still offered segmentation, but also privilege levels. Levels 0-2 are supervisor level and can do most things. Rings 1-2 cannot run privileged instructions but this is the only real limit; otherwise they are as privileged as ring 0.

Ring 3 meanwhile is user mode. If you have your segment selector set to point to this ring, you require the help of the kernel via some system call interface in order to do anything requiring privileged CPU or memory access.

These days, its pretty much required in 64-bit x86 to not use segmentation. However, segment selectors are still there - all segments just overlap and cover the entire address space.

So, the original purpose of ring 0-3 was to isolate privilege between user mode code and the kernel and stop user mode code walking all over system control structures.

Then virtualization became a thing on x86 and Intel/AMD decided to add hardware support for it. This requires a piece of supervisor (hypervisor) code to set up some control structures (called VMCS) defining the virtual machines and then call `vmenter` and handle `vmexit` i.e. conditions on which the virtual machine needs help from the hypervisor.

This piece of code is referred to as ring -1. There is no such actual privilege level, but since it can host multiple kernels all of which believe they have ring 0 access to the system, it makes sense.

System Management Mode is another beast with special instructions. Firmware (your BIOS) sets up a SMM handler to handle System Management Interrupts - configurable depending on what the firmware wants to be notified of. When these events are triggered, the OS (or even hypervisor) is suspended and a special address space is entered. This area is supposed to be invisible to the OS itself, while executing on the same processor. Hence ring -2, since it is more privileged than a hypervisor would be.

Youll also hear ring -3 mentioned here and there in reference to Intel ME or AMDs PSP. This is a second processor running a separate firmware (Intel I believe uses ARC SoC processors) capable of doing anything it likes to the primary system. Ostensibly this is to provide IPMI/remote management of hardware type functionality. It can run whenever there is power to the hardware regardless of whether the main system is powered on or not - its purpose, as I say, would be to power on the main system.

From a security perspective, the lower ring you can get yourself into, the more undetectable you can make yourself. The bluepill research was about hiding from an OS the fact it was truly running in a VM. Later research has been done on SMM persistence. SMM persistence for example would potentially allow you to reinstall your malware even on a complete wipe of the hard disk and reinstall. Intel ME potentially opens up an always on persistent networked chip to install malware on the main target.

Ive stuck to Intel chips here but you should be aware other platforms work differently. For example, ARM chips have supervisorand user modes, amongst others.