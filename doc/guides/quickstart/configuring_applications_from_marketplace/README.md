# Configuring Applications From The Marketplace
Running applications from an external repository such as the Vorteil Marketplace is a simple and easy process, but eventually users may wish to configure applications beyond the default settings of each application.

The focus of this guide will be to pull an application from the Vorteil Marketplace, export it as a 'project, and edit it locally.

It is assumed that users will have already added the Vorteil Marketplace repository to their environment.

## Pulling From a Remote Source to the Local Repository
To pull 'helloworld' from the Marketplace, run the following command:

```bash
$ vorteil repositories pull marketplace:sisatech/helloworld sisatech helloworld
```
*This is directing Vorteil to pull the 'sisatech/helloworld' application from the connection saved as 'marketplace', and store the resulting application in the local 'sisatech' bucket, as 'helloworld'.*

## Unpacking an Application
The 'helloworld' application from the Marketplace is now located within the local Vorteil Repository. It is now possible unpack this application:

```bash
$ vorteil packages unpack local:sisatech/helloworld
```
*An additional option argument can provide a destination for the unpacked files.*

The CLI will print the location of the unpacked files to stdout.

## Reconfiguring an Application
Vorteil uses 'Vorteil Configuration' (.vcfg) files to configure applications at build/runtime. Within the directory containing the unpacked files, open 'app.vcfg'.

There is a lot of information within this file, but don't be overwhelmed. The 'hellworld' application is configured to host a basic webpage, with the background colour set to the value provided by the 'BACKGROUND' environment variable. Add the following fields beneath the ```binary = "/app"``` line:

```toml
[env]
  BACKGROUND="0x123456"
```
Save the file, and run the application (this time directly from the directory of unpacked files)!

```bash
$ vorteil run
```

If your default/chosen hypervisor is bridged, you may observe the IP address assigned to the application within the Virtual Machine logs that will have printed to stdout. Otherwise, visit ```http://localhost:8888``` to view your 'helloworld' website, complete with the background colour set by the 'BACKGROUND' environment variable.

## Finishing Up
If you are satisfied with the changes that you have made, you can push this to your local repository as a new application version:

```bash
$ vorteil repositories push local sisatech/helloworld ./
```

This will compile the files located within the directory, and import the resulting package directly into the local repository. 

## See Also
- [What is a VCFG?](../../../apps/vcfg/introduction)
- [Tagging an Application Version](../tagging_an_application_version)
