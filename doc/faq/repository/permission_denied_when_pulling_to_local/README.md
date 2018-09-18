# 'Permission Denied' when storing app in local repository
If encountering this error, it is likely that the user does not have adequate read/write permissions on the filesystem location specified as the 'storage' destination. This field is labelled as the 'storage-directory' in the service configuration file.

```toml
[storage]
  type = "local"
  service-account-key = ""
  bucket-name = ""
  storage-directory = "/home/USER/.vorteild/storage"
```

## See Also
- [Where is the service configuration file?](../../host_configuration/where_is_the_config)