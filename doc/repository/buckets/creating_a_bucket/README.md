# Creating a Bucket
Buckets exist as a storage container concept within a Vorteil repository. They can be created and maintained via the Command-Line Interface, the Developer Studio, or the GraphQL API.

By default, any object (including buckets) are instantiated with read, write, and execute permission available only to members of the creator's primary group. To grant access to additional users, see '[Applying ACL Rules to an Object](../../acls/applying_acl_rules)'.

## Command-Line Interface
The following command will create the 'myBucket' bucket in the local Vorteil repository (alternative repositories can be specified by the ```--repository``` flag).

```bash
$ vorteil repositories new-bucket myBucket
```

## Developer Studio
The Vorteil Developer Studio provides a simplified approach to creating buckets within repositories. Open the Studio, and move your mouse cursor to the sidebar icon corresponding with the repository to act upon. Right click to open the context menu, and select 'Add Bucket'. 

Enter the desired bucket name, and press 'enter' to finish creating the bucket.

## GraphQL API
The CLI and Developer Studio approaches both wrap the GraphQL API in order to achieve the desired effect. These APIs can be programmatically triggered to reach the same outcome.

```bash
URL: http://localhost:7472/graphql
Method: POST

mutation {
    newBucket(name:"myBucket"){
        authorization {
            id
        }
    }
}
```

## See Also
- [Applying ACL Rules to an Object](../../acls/applying_acl_rules)
- [Granting Access to a Singleton Object](../../acls/granting_access_to_singleton)