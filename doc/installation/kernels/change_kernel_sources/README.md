# Experimental and Beta Kernels
Vorteil's default behaviour is to use the public release source to retrieve kernels from. In some cases, you may find yourself with the need to access 'release candidate' kernels.

To do this, open the Vorteil service configuration file for editing:

- OSX: /Applications/Vorteil.app/Contents/Resources/conf.toml
- Windows: C:\Program Files\Vorteil\conf\conf.toml
- Linux: /opt/vorteil/conf/conf.toml



```toml
[kernel-sources]
  remote-repositories = ["https://downloads.vorteil.io/system"]
```
for the default behaviour, or

```toml
[kernel-sources]
  remote-repositories = ["https://downloads.vorteil.io/system", "https://downloads.vorteil.io/system-release-candidates"]
```
for the latest release candidates.

After editing and saving your configuration file, restart your machine for the changes to take effect.
