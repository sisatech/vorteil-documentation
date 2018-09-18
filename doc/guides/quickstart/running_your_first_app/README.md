# Running Your First App on Vorteil
This focus of this guide is to walk users through the process of establishing a connection to the Vorteil Marketplace and running one of the listed applications. This is a simple process, that can be carried out with the Command-Line Interface (CLI), the Developer Studio application, or directly via the GraphQL APIs.

For demonstrative purposes, steps displayed will be from the CLI approach.

## Connect to the Vorteil Marketplace
The Vorteil Marketplace is a repository (just like the local repository included in a default Vorteil installation) which has been configured to be accessible by the public. 

To begin interacting with the Marketplace, it is required that the connection is added to the local Vorteil environment:

```bash
$ vorteil repositories connect marketplace https://marketplace.vorteil.io
```

## Pick an Application
With the Vorteil Marketplace connection saved, a list of available applications can be retrieved. 

```bash
$ vorteil repositories list sisatech --repository marketplace
```
*Note: this action will list all apps contained within the 'sisatech' bucket on the marketplace. All applications hosted on the marketplace are listed within this bucket.*

## Run an Application
Running a Vorteil application is easy; in fact, you're not even required to download the package files to your local machine before running! 

Execute the following command to run the 'helloworld' application from the marketplace:

```bash
$ vorteil run marketplace:sisatech/helloworld
```
Congratulations! You've run run your first application with Vorteil!

## See Also
- [Configuring Applications From The Marketplace]()