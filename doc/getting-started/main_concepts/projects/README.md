# Projects

The concept of a 'project' within the Vorteil ecosystem refers to a structured directory from which a package or disk image can be built.

For example:

![Helloworld Project Structure](https://storage.googleapis.com/vorteil-dl/assets/documentation/project-structure.png "Helloworld Project Structure")

All projects require a '.vorteilproject' file, as depicted in the above image. The '.vorteilproject' file plays an important role, defining which files/folders within the project directory will be included on built disk image filesystems, as well as allowing multiple built 'targets' (or 'variants'.

```toml
ignore = [".vorteilproject", "app.png", "app.vcfg"]

[[target]]
  name = "default"
  vcfgs = ["app.vcfg"]
  icon = "app.png"

[[target]]
  name = "without-icon"
  vcfgs = ["app.vcfg"]
  icon = ""
```
The above example defines two build variants: "default" and "without-icon".
All variants defined within the .vorteilproject file will adhere to the ignore patterns defined within the top-level `ignore` field.

Here's an example of running this application from the command-line:
```bash
$ vorteil run ./helloworld:without-icon
```

Tools within the Vorteil ecosystem also support linting existing '.vcfg' files within a project directory, providing useful feedback for troubleshooting configuration errors.

## See Also
- [For a more in-depth project explanation](../../../apps/projects/introduction/)
- [Project Targets](../../../apps/projects/targets/)
- [Ignore Targets](../../../apps/projects/ignore/)