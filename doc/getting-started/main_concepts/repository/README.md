# Repository (vRepo)

Users may take advantage of the built-in Vorteil Repository's version control and sharing capabilities in order to streamline application development and collaboration. A single user may configure as many remote-repository connections as desired. Vorteil Repository access is governed by ACLs (Access Control Lists) which may be configured by the appropriate user.

![Vorteil Diagram](https://downloads.vorteil.io/assets/vorteil-diagram.png)

When communicating with external repositories, all requests pass through the local Vorteil server, applying appropriate authentication details to the request behind the scenes. This enables Vorteil to streamline many actions that a user may which to execute (examples of this include the ability to 'run', 'build', or 'pull' an app directly from the official [Vorteil Marketplace](https://marketplace.vorteil.io) repository).

## See Also
- [Introduction to the Vorteil Repository (vRepo)](../../../repository/introduction/introduction_to_vrepo)
- [Pushing and Pulling Applications](../../../repository/sharing/pushing_pulling_apps)
- [Creating a Bucket](../../../repository/buckets/creating_a_bucket)
- [Applying ACL Rules to an Object](../../../repository/acls/applying_acl_rules)
- [Granting Access to a Singleton Object](../../../repository/acls/granting_access_to_singleton) 