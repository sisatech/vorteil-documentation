# Pick an Application
With the Vorteil Marketplace connection saved, a list of available applications can be retrieved. 

#### CLI
```bash
$ vorteil repositories list sisatech --repository marketplace
```

#### Studio
Upon adding the repository you should now be able to go to the repository via the sidebar. When clicked on a new list should appear which will be applications inside a bucket.

![Marketplace List Image](https://storage.googleapis.com/vorteil-dl/assets/documentation/marketplace-list.png "Marketplace List")
#### Graphql
```
POST
{
"query": "
    query{
        bucket(name:"sisatech"){
            appsList{
                edges{
                    node{
                      name
                    }
                }
            }
        }
    }
  ",
  "operationName": "bucket",
  "variables": {}
}
```
*Note: the above actions will list all apps contained within the 'sisatech' bucket on the marketplace. All applications hosted on the marketplace are listed within this bucket.*
