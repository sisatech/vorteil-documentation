# Granting Access to a Singleton Object
Singletons act as gatekeepers for broader ranges of functionality within the Vorteil environment. A user/group that does not have any permissions on the repository singleton object, for example, will be unable to retrieve any information from the repository (such as buckets, apps, etc.).

This process can be done through the admin web UI (hosted by default on ```http://localhost:7471```), or programmatically via the GraphQL API.

```bash
URL: http://localhost:7472/graphql
Method: POST

query {
    getSingletonID(type:repository)
}
```
*Acceptable values for 'type' include: repository, local_ops, nodes, projects.*

Now that the singleton ID has been retrieved, follow the steps in [this article](../applying_acl_rules).