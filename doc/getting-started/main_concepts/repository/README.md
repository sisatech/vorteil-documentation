# Repository (vRepo)

The Vorteil daemon allows us to connect and host repositories. This grants us access to easily share what we're currently working or grant the benefit of being able to run external applications locally.

The diagram below will try and explain what we are trying to achieve here.

![Simple Explanation of Repositories](https://storage.googleapis.com/vorteil-dl/assets/documentation/simple-repo-explanation.png "Simple Explanation of Repositories")

Basically the User will have a direct connection with his local Daemon server. The local daemon server will then forward requests to different repositories/daemons. The cool functionality that this provides is the ability to push packages from one external repository to another external repository. 

All of the access to a repository can be maintained by Access Control Lists. This provides the benefits of stopping users from accessing packages you may not want public.

The repository comes with built in version control where we hash each package that gets uploaded. To further improve the version control we allow the use of tagging applications to grant the users a way to name each version. 

The daemon provides peer sharing. The best example of peer sharing is connecting your daemon to the marketplace daemon server. This process doesn't make the user connect to the marketplace server it tells the daemon they are running locally to forward the request on.

With the repositories another component called Listeners come with them. A listener is an object which you can attach to a package on a repository. Whenever that package receives a new version update the listener will then trigger and do a certain action like a post request.

## See Also
- [Repository Introduction](../../../repository/introduction/introduction_to_vrepo)
- [Sharing / Connecting Repository](../../../repository/sharing/pushing_pulling_apps)
- [Buckets](../../../repository/buckets/create_a_bucket)
- [Apply ACL Rules](../../../repository/acls/applying_acl_rules)
- [Access to Singleton](../../../repository/acls/granting_access_to_singleton)