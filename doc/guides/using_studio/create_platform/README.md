# Create a platform

Now to run this application using the vorteil kernel. We will need to create a platform. The platform we will be creating will be for the Virtualbox hypervisor.

To get to the area to add a new platform look at the bottom left for the Settings menu. Upon clicking the settings menu you should see a tab window open like below.

![Settings](https://storage.googleapis.com/vorteil-dl/assets/documentation/settings.png "Settings")

The first area is the Default Settings. These settings will be used when the VCFG doesn't provide the field they want to use. Some of the other settings will stop the VCFG details for example Ignore Kernel will always use the default one set than the one in the VCFG.

Right now we are after the Platforms section to add a platform. Upon choosing virtualbox a new form should appear like the one below.

![VirtualBox](https://storage.googleapis.com/vorteil-dl/assets/documentation/vbox.png "Virtualbox Form")

- Platform Name - what the daemon will refer to for provisioning
- Network Interface - used to work out whether its Bridged or NAT. Leave default for NAT or choose the device you want to run this on for bridged.
- Super Debug - Provides more provisioning logs to help debug why your virtual machine isn't running properly as it may be happening in a stage thats harder to correct.
- Headless - You can run most hypervisors in GUI or non GUI mode. The default is set to non GUI as the Developer Studio outputs the information on the virtual machine page.
