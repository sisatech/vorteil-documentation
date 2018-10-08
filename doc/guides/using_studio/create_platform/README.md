# Creating a Platform

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