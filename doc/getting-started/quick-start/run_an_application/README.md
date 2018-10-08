# Run an Application
Running a Vorteil application is easy; in fact, you're not even required to download the package files to your local machine before running! 
#### CLI
Execute the following command to run the 'helloworld' application from the marketplace:

```bash
$ vorteil run marketplace:sisatech/helloworld
```

#### Studio
With the list retrieved previously we should now be able to right-click on the application and run the application.

![Marketplace Run Image](https://storage.googleapis.com/vorteil-dl/assets/documentation/marketplace-run.png "Marketplace Run")

#### Graphql
```
{
"query": "
   mutation{
    provision(germ:"marketplace:sisatech/helloworld", platform:"kvm", start:true){
        job{
         id
        }
        uri
      }
    }
  ",
  "operationName": "provision",
  "variables": {}
}
```
Congratulations! You've run run your first application with Vorteil!
