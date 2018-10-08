# Using the Studio
## Creating a Platform

In order to run an application, a platform must be configured. Vorteil will attempt to automatically configure itself to use any hypervisors detected on the local machine; but it is still possible to add platforms with alternative configuration (such as GUI/headless modes), and even add supported remote platforms (such as Google Cloud Platform, Amazon Web Services, or VMWare vCenter).

With the Developer Studio open, click on the 'Settings' icon on the bottom-left corner of the window.

![Settings](https://storage.googleapis.com/vorteil-dl/assets/documentation/settings.png "Settings")

At the top of the Settings page, current default settings are displayed. These settings will be used in various processes if no alternative is specifically defined. Other settings, such as 'Ignore Kernel' (instructing Vorteil to use the default kernel instead of the kernel defined in an application config), exist as 'quality of life' improvements.

Below this section is the 'Platforms' section.

![VirtualBox](https://storage.googleapis.com/vorteil-dl/assets/documentation/vbox.png "Virtualbox Form")

The above image shows the input form for a new 'VirtualBox' configuration. Provide a name, and complete the rest of the form as desired. 

- Platform Name: a user defined name for the new platform configuation
- Network Interface: if provided with valid input, applications running on this platform will run on a bridged network. Leave blank for NAT.
- Super Debug: when active, processes will include more debug information in the logs relating to this platform.
- Headless: if checked, virtual machines will start up without a GUI window opening.

## Get a Project

Projects can be manually created and built by users, or extracted from existing Vorteil application packages, such as those found on the official [Vorteil Marketplace](https://marketplace.vorteil.io).


To connect to the Vorteil Marketplace with the Developer Studio, left-click on the '+' icon within the application sidebar. This will open a modal which prompts the user for information:

![New Repo](https://storage.googleapis.com/vorteil-dl/assets/documentation/newrepo.png "New Repo Screenshot")

Leave the dropdown menu set as 'No Auth', and name the repository as desired (for example: 'marketplace'). In the second text input field, enter the address at which the repository is located: `https://marketplace.vorteil.io`.

To finish this process, click the 'Add' button.
A new icon will appear in the application sidebar; clicking on this will reveal a list of accessible applications hosted within the newly-added repository. Right-click on any application within the repository, and select 'Export as Project'.

Once the process is complete, the new project will appear in the projects list.

## Run the Project

Running a project is simple. Within the 'projects' view, right click on the desired application and click on 'Run' within the resulting context menu. 'Run As' can be select from the menu if a non-default platform should be used.

Upon completion the Developer Studio will open the appropriate 'Virtual Machine' page.

![Virtual Machine Information](https://storage.googleapis.com/vorteil-dl/assets/documentation/vm.png "Virtual Machine Information")