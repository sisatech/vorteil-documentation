# Virtual Machine Information

Below you will find the Virtual machine tab open.

![Virtual Machine Information](https://storage.googleapis.com/vorteil-dl/assets/documentation/highlightedvm.png "Virtual Machine Information")

This interface uses three tabs to separate a large amount of information provided. 
```
    App            | #bd1213    | Red
    Button Toolbar | #80c6bf    | Teal
    Tabs           | #6e1e66    | Purple
    Info           | #afbb2e    | Yellow
```

## Button Toolbar

### Status Change Section

This section provides you actions to perform on the virtual machine. The ability to Start, Pause, Stop and Delete a virtual machine. These actions are easily traceable by the provision logs.

### Other Actions

The first button on the toolbar is 'Rebuild and Run'. It provides you a way to rebuild the virtual machine with the same setup you had before but apply any changes you made from where it was built from.

The eye icon in the button toolbar is for the ability to tail logs. When its currently set to tail both the Virtual Machine Logs tab and Provisioning Logs tab will tail the file.

The details icon in the toolbar provides a user to see the corresponding job logs that was attached to the provisioning process of this virtual machine. 

The last icon in the button toolbar gives you the ability to download the disk the virtual machine is running. The only way to download the disk is to have the virtual machine in a Ready or Stopped state.

## Virtual Machine Logs

The first tab that opens after displaying the virtual machine page are the logs of the virtual machine. You will find stopping or starting this virtual machine repeatedly will keep adding to this log file. 

## Provisioning Logs

Provisioning logs provides you with helpful information on how the provision process of the virtual machine went. It provides a way to check the state changes of the virtual machine for example if it ended up in a 'Broken' state you would be able to check where it happened. Did it get to Started then break?

## Details

The most important information Details provides is the port forwarding table. Basically when a virtual machine is told to use an ip address it'll attempt to reserve it. If it ends up failing it will bind the address to a random port and let you know.

The rest of the details provides you with short information about the virtual machine. Was it run with Arguments? What size is the disk built? Where was this virtual machine built from?