# Pushing and Pulling Applications
Applications can easily be shared between users and repositories with Vorteil's push/pull capabities. This article details the following processes:

- pushing a local application to a repository
- pulling an application to the local repository from an external repository

## Pushing a Local Application
Internally, this process is identical to [uploading an application](../../apps/uploading_apps) (in fact, the ```upload``` command is just an alias of the ```push``` command). This can be achieved by running the following command, where 'BUILDABLE' is a filepath to a packaged Vorteil application, or an unpacked Vorteil Project.

```bash
$ vorteil push local BUCKET/APP BUILDABLE
```
*Note: 'BUILDABLE' can also be in the form of a repository URL. This means that users may 'push' from any repository (provided they have sufficient access).*

## Pulling an Application to the Local Repository
The 'pull' command syntax is very similar to the 'push' command syntax.

```bash
$ vorteil pull BUILDABLE BUCKET APP
```

For example:
```bash
$ vorteil pull marketplace:sisatech/helloworld myApps helloworld
```
*This will pull the 'helloworld' application from the marketplace and store it within the 'myApps' bucket, titled 'helloworld'.*

## See Also
- [Applying ACL Rules to an Object](../../acls/applying_acl_rules)
- [Running Your First App on Vorteil](../../../guides/quickstart/running_your_first_app)
- [What is 'BUILDABLE'?](../../../apps/general/buildable)