# User Interface

![Colour Coded view of the interface](https://storage.googleapis.com/vorteil-dl/assets/documentation/highlightedview.png "User Interface Screenshot")

    Title        | uncoloured | none
    Context Menu | #be4600    | Orange
    Actions      | #0d07a8    | Dark Blue
    Explorer     | #680a09    | Dark Red
    Tab Section  | #0d4612    | Dark Green

### Actions

The actions panel is the main navigation unit throughout the Vorteil Developer Studio. It allows us to navigate through Projects, Virtual Machines, Jobs and Repostiories.

### Explorer

#### Projects

The projects explorer is a detailed treeview showing each project that is currently added to the daemon. You can add, remove or edit files inside the project. The project list contains a context menu allowing you to Push, Package, Build, Provision and Run. 

#### Virtual Machines

The explorer section for Virtual Machines provides a simple list that is colour coded based on the status of the virtual machine. 

The list is ordered by time and the format for each object is 

`HH:MM (platform) virtual machine name`

    Red      - Stopped (The virtual machine is currently stopped.)
    Dark Red - Broken  (Something occured during the provisioning process which is why it ended up as broken status.)
    Yellow   - Paused  (The virtual machine is currently paused.)
    Green    - Alive   (The virtual machine is currently alive.)
    Purple   - Mystery (The daemon does not know what the current state of the virtual machine is.)

The difference between a Paused and Stopped status virtual machine is Paused holds onto the IP and wont release it. Where as Stopped releases everything related to the virtual machine.

The virtual machines list is provided a context menu on all the items which allows you to easily do group related actions like 'Stop All VMS' or 'Remove All Broken VMs'. Alternatively you can use the context menu on a single virtual machine to Start, Stop, Pause and Delete.

#### Jobs

The jobs section of the GUI allows us to view the operations we perform on the daemon. Basically any sort of Build, Provision, Run, Package or Push will show up as a Job log. Basically two line mentioning what the job is doing and how long ago did the job happen. 

The jobs list is ordered by time and formatted like below

    Provisioning VM
        a minute ago.

If you want a more detailed explanation of each job you are able to open up the job by clicking or hovering over the job to get the description.

#### Repositories 

The tree for a repository view

    Bucket|
           App|
               Version|
    
    An example
    sisatech|
             helloworld|
                        3e56789|

All three of the above objects have context menu's attached that do different operations. The main one that I recommend using is 'App' as it will always go with the latest version of that application.

The local repository is clearly shown in the actions panel bar as that database-like icon. Below the icon however you may find one letter characters which are external repositories. Alternatively you may only see a '+' icon which is a button used to create a new external repository.

### Tab Section

The main area of the client where most information will be displayed. This could range from the editor, virtual machine information or the job page. 

Tabs will only open if they dont exist in the group to begin with. If the tab exists in the group it will reopen and display that tab instead. The tabs also contain a context menu for your favourite actions 'Close', 'Close All' and 'Close Others'.

Tabs appear at the top and you are able to navigate through each tab with 'Ctrl+TAB' or if you simply want to close the tab middle mouse button works.

### Context Menu

A general application menu that allows us to complete actions like Create a project, add a repository or  find more help on how to use the application.