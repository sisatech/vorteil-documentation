# Tagging an Application Version
The Vorteil Repository provides easy and simply version control for applications. Versions of an application can be tagged for a better user experience. The following command will tag the latest version of 'helloworld' (within the 'sisatech' bucket) as 'testTag'.

```bash
$ vorteil repositories tag sisatech/helloworld testTag
```

This command can be used to target a specific version of an app. To tag the 'testTag' version as 'retagged':

```bash
$ vorteil repositories tag sisatech/helloworld/testTag retagged
```
