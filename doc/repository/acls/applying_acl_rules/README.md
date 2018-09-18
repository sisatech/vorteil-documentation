# Applying ACL Rules to an Object
Objects within the Vorteil repository are protected by access control lists (ACLs). By default, objects are instantiated with read, write, and execute permission granted only to the creator's primary group. In order to share objects with other users or groups, one must add ACLs to an object.

This can be achieved through the Developer Studio application, or with the GraphQL API.

## Developer Studio
With the Vorteil Developer Studio open, navigate to the repository containing the bucket or app that requires modified ACLs. Right click on the object, and select 'Change ACLs' from the resulting context menu. 

A modal will open, displaying the current ACL rules for the selected object. To add rules for an entity that does not already exist within the populated list, enter the entity name in the input field provided. Check the box for each access level the entity will be granted. Press 'Add'.

Users may also modify the existing ACLs by ticking/unticking the boxes within the populated table. When satisfied, press 'Save'.

## GraphQL API
In order to apply ACL rules to an object, the object's ID must be obtained. Objects that support ACL actions contain an 'authorization' section within the GraphQL API, from which this can be retrieved.

*For example:*
```bash
URL: http://localhost:7472/graphql
Method: POST

query {
    bucket(name:"myBucket"){
        authorization {
            id
        }
    }
}
```

With the identifier available, ACLs may now be applied.

```bash
URL: http://localhost:7472/graphql
Method: POST

mutation {
    applyACLRule(id: String, action: String, group: String) {
        id
    }
}
```
*Appropriate values for 'action' field include: "READ", "WRITE", "EXEC".*

## See Also
- [Granting Access to a Singleton Object](../granting_access_to_singleton)