# Connect to the Vorteil Marketplace
The Vorteil Marketplace is a repository (just like the local repository included in a default Vorteil installation) which has been configured to be accessible by the public. 

To begin interacting with the Marketplace, it is required that the connection is added to the local Vorteil environment:

#### CLI
```bash
$ vorteil repositories connect marketplace https://marketplace.vorteil.io
```
#### Studio
By clicking the blue '+' icon. A modal will open, allowing you to add a repository to connect to.

![Connect Repository Image](https://storage.googleapis.com/vorteil-dl/assets/documentation/marketplace-repo.png "Connect Repository")

#### Graphql
```
POST
{
"query": "
    mutation{
        newNode(name:"marketplace", host:"https://marketplace.vorteil.io", insecureSkipVerify:false){
            name
            host
            type
        }
    }
  ",
  "operationName": "newNode",
  "variables": {}
}
```
