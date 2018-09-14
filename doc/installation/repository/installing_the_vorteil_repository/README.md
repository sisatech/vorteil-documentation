# Installing the Vorteil Repository
This article will guide users through the repository setup and configuration process.

## General Installation
A standard Vorteil installation includes access to the Vorteil Repository, and can be configured on Windows, OSX, and most Linux distributions. Download an operatin system appropriate installer from the [Vorteil website](http://www.vorteil.io).

Run the installer and follow the prompts to complete the initial installation process.

## Configuring the Repository
Having run the installer, you will already have a basic repository installed. Now it's time to configure the repository for other users to access. This will require the use of the [Admin API](//TODO). The Admin API endpoint is set to localhost:7471 by default; alternatively, [a lightweight web interface is provided](http://localhost:7471).

Let's create an authentication key for a new user/group.

```rest
URL: http://localhost:7471/api/keys
Method: POST

Headers {
        "Content-Type": "application/json"
}

{
        "username":"testUser",
        "primarygroup":"testUser",
        "groups":["testUser", "testGroup"],
        "comment":"This is our test user, from the installation guide!"
}
```

The reponse will include the newly generated key. **Save this key somewhere: you will use it further into this guide.** It is worth noting that the new user and groups that we have just created do not yet have access to any repository functionality; this can be observed by visiting the [repository web interface](http://localhost:7472) and attempting to 'Sign in' with your new key.

Our new user or groups will need to be granted, at a minimum, READ permission on the 'Repository Singleton' object. The admin web interface provides a simplified approach to this: simply click on the 'Singletons' tab, and then the 'Add ACLs' next to the appropriate singleton.

The REST-equivalent of this process is as follows:

```rest
URL: http://localhost:7472/graphql
Method: POST

query {
        getSingletonID(type:repository)
}
```

Now we have the identifier of the repository singleton object. We can use this to grant our new user READ permission.

```rest
URL: http://localhost:7472/graphql
Method: POST

mutation {
        applyACLRule(id:"<RepositorySingletonID>", action:"READ", group:"testGroup"){
                id
        }
}
```

The 'testGroup' group now has READ access to the repository singleton!

## Verifying Expected Behaviour
Now that we've established an appropriate access level for our new group, let's verify that it is working! Using your preferred method (graphql API, CLI, or Vorteil Developer Studio, connect to the Vorteil Marketplace if you haven't already.

The Developer Studio provides a simplified approach to the following steps, so the process will be demonstrated using the CLI and GraphQL API.

```sh
vorteil repositories connect marketplace http://marketplace.vorteil.io
```

Create a bucket in your local repository to store an app in.

```sh
vorteil repositories new-bucket testBucket
```

Pull an application from the marketplace into your local repository.

```sh
vorteil repositories pull marketplace:sisatech/helloworld testBucket helloworld
```

Now, similar to the steps taken to grant 'testGroup' access to the repository singleton, let's grant access to the 'testBucket' bucket, and the application inside of it.

```rest
URL: http://localhost:7472/graphql
Method: POST

query {
  bucket(name:"testBucket"){
    authorization{
      id
    }
    app(name:"helloworld"){
      authorization{
        id
      }
    }
  }
}
```

The response will contain the identifiers for the bucket, and the application. For each identifier, perform the following:

```rest
URL: http://localhost:7472/graphql
Method: POST

mutation {
        applyACLRule(id:"<IDENTIFIER>", action:"READ", group:"testGroup"){
                id
        }
}
```

Revisit the repository web interface, and Sign in again with the key, if required. You should now see the 'testBucket' bucket. Click on the right-chevron listed beside the bucket listing to view the bucket contents, where you should now see your app!

## Exposing the Repository to Other Users
A default Vorteil installation will not be exposed on any external ports, for security reasons. To expose an authentication-protected port, edit the server configuration file.
This file can be found at the following location:

- OSX: /Applications/Vorteil.app/Contents/Resources/conf.toml
- Windows: C:\Program Files\Vorteil\conf\conf.toml
- Linux: /opt/vorteil/conf/conf.toml

```toml
[[network]]
  disabled = false
  bind = "wlp4s0:7472"
  disable-tls = true
  tls-certificate = ""
  tls-key = ""
  website-files = "/opt/vorteil/web/vrepo"
```

In this example, the bind address is set to port 7472 on the 'wlp4s0' network interface controller (NIC). You may also set this is a specific address, if specifying a NIC name does not appeal to you. For example:

```toml
  bind = "10.0.0.15:7472"
```


## Additional Notes
In this guide, we have created a key for previously unknown user/groups, and granted then access to some of the resources hosted on the Vorteil repository. Users will require an authentication key, like the one generated at the start of this guide, in order to establish an authenticated connection to the repository.

A 'default' key can also be established by modifying the Vorteil server's configuration file.

The following fields in the configuration file will result in all unauthenticated requests arriving through the protected endpoint being assigned a key for the "public" group.

```toml
[access.default-key]
enabled = true
primarygroup = "public"
groups = ["public"]
comment = "DefaultKey"
```

Assigning permissions to this user is identical to the process previously demontrated within this article.