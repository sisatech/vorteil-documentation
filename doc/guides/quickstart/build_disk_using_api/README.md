# Build a Disk using the APIs
Building a disk using the APIs follows a simple two step procedure. The first using Graphql to tell the daemon to start the operation of build on a germination string. Following that we will send a GET request to retrieve the disk that was built. 
## Graphql

	mutation{
		build(germ:"rootmarketplace:sisatech/helloworld" injections:["binaryid"]){
            job{
                id
            }
            uri
	    }
	}

* The injections is way for users to change the current application's setup and alter it on run time. An example being changing the binary, configuration or filesystem.
* The job id is a way for us to subscribe to the progress of this operation we have just started.
* The URI is the path in which we will send a REST request to. Being that this is a build request the URI will only be valid at `localhost:7472/api/build/{uri}`.

The response of the Graphql request will be similiar to the one below.

    {
        "data": {
            "build": {
                "job": {
                    "id": "job-gydins"
                },
                "uri": "LndCUKYkynSgYXrgHBQcVtAmIsMUoYGgcXNPChZtMEQpIShHZqDLHbczPlsHdfgg"
            }
        }
    }

You won't notice any changes to the job until the REST request happen. But to check the job we can simply use Graphql.

    query {
        job(id:"job-gydins"){
            id
            progress{
                status
                started
                finished
                error
                total
                units
                progress
            }
        }
    }
The result of this request should be similar to below.

    {
        "data": {
            "job": {
                "id": "job-gydins",
                "progress": {
                    "error": "",
                    "finished": null,
                    "progress": 0,
                    "started": 1537305242,
                    "status": "running",
                    "total": 100,
                    "units": "%"
                }
            }
        }
    }

## REST
In our build request above we want to inject a binary when we run this operation. So we have to tell the daemon what binary we want to use instead. This can easily be done by posting to the url of        

    http://localhost:7472/api/build/LndCUKYkynSgYXrgHBQcVtAmIsMUoYGgcXNPChZtMEQpIShHZqDLHbczPlsHdfgg

We also need to tell the daemon what we are injecting as of right now it won't know. This is done by parsing two headers for the request on being a Type and the other being the ID that we called the injection.

    Injection-Type - binary
    Injection-Id   - binaryid

After we've setup the POST request attach the binary file to the request and send it.

With the same URL we used to POST the binary we also can GET the disk. You should be able to get the disk near instantly as it is sent out in a stream. No progress will be made if the injections aren't resolved however. Upon being successful the disk should be now downloaded to where the GET request ended up.

