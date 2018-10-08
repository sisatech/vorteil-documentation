# Projects

Projects are a structured directory from which Vorteil images can be generated. A generic Helloworld project structure is shown below.

![Helloworld Project Structure](https://storage.googleapis.com/vorteil-dl/assets/documentation/project-structure.png "Helloworld Project Structure")

- etc - contains the certificates to enable ssl
- lib/lib64 - shared libraries that are required to run the binary
- .vorteilproject - project configuration file that tells the daemon server how to intepret this setup
- binary - the binary that is run
- icon - icon is mainly used in displaying it from a package. An example of this would be browsing the marketplace.
- vcfg - configuration file to tell vorteil how to run this application.

The way a project is structured will depend on the application you are running. As Vorteil Tools builds the base level directory as root for each virtual disk.

The '.vorteilproject' file allows us to build variants of our application. We can choose which files or an icon that we want to build onto the disk. The other option is to tell each build variant to build with different VCFG files so we could completely run the application with different settings. An example of a .vorteilproject contents is shown below.

```
ignore = [".vorteilproject", "app.png", "app.vcfg"]

[[target]]
  name = "default"
  vcfgs = ["app.vcfg"]
  icon = "app.png"

```

- ignore - tells Vorteil to not build these files on the virtual disk.
- target - is used to tell Vorteil how to run the application. We want to use the app.vcfg and the app.png during the build process. Another example shown below merges two vcfgs together. You can also have as many targets as you would want.

```
ignore = [".vorteilproject", "app.png", "app.vcfg"]

[[target]]
  name = "default"
  vcfgs = ["app.vcfg", "app2.vcfg"]
  icon = "app.png"
[[target]]
  name = "prod"
  vcfgs = ["papp.vcfg", "papp2.vcfg"]
  icon = "app.png"
```

The API provides project linting help if you need to check why your project isn't running. Can give you helpful tips like the binary isn't specified in the VCFG file. The portability of our packages allows us to easily transfer entire application setups to different users without needing to sort dependencies out.


## See Also
- [For a more in-depth project explanation](../../../apps/projects/introduction/)
- [Project Targets](../../../apps/projects/targets/)
- [Ignore Targets](../../../apps/projects/ignore/)