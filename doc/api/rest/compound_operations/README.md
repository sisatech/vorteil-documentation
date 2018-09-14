## Compount Operations
The operationID can be retrieved by sending Graphql requests that resemble the request name. An example provided below for packing a project.

	mutation{
		pack(germ:"/home/sisatech/Documents/projects/helloworld"){
            job{
                id
            }
            uri
	    }
	}

The response below

    {
        "data": {
            "pack": {
                "job": {
                    "id": "job-townwv"
                },
                "uri": "rMhbkikzqQzDQGgwFuzorBbnFVbqEyFHuYakgxwJcQZqQYLGQLVqgSGXOuvdJpEk"
            }
        }
    }

The uri section of this response is the operationID. The job can be subscribed to track the progress of the operation. When the job has completed you may fetch the result using the APIs listed below.

The operationID section handles post requests for when you add injections to the Graphql request. For example I want to push an object to a repository with a new vcfg than the one it contains. I can send a post request to `/api/push/(operationID)` with the body of a VCFG that will replace it and the operationID being the uri of the Graphql request.

### /api/pack/(operationID)
#### GET
The ability to fetch the package that you created using the Graphql request. It will be streamed out to you.
#### POST
An example of POST request injecting a VCFG field.

	mutation{
		pack(germ:"/home/sisatech/Documents/projects/helloworld" injections:["abc"]){
            job{
                id
            }
            uri
	    }
	}

To inject the vcfg you send the request to this address. You need to apply two headers for it to succeed. Notice we used `injections ["abc"]` means were referring to whatever we upload to this request as abc. Below is the appropriate headers for the vcfg input. The body of this request should be the VCFG file by itself.

        Injection-Id   | abc
        Injection-Type | configuration

Upon sending this request off the response back should be of status code 200. Then you will be able to GET the same address for the file that you have just created.

### /api/unpack/(operationID)
#### GET
The ability to fetch a tar archive that contains the contents of a project. After untarring you will be able to add this project to the GUI or run straight from the command line.
#### POST
### /api/push/(operationID)
#### POST
### /api/provision/(operationID)
#### POST
### /api/build/(operationID)
#### GET
The ability to build virtual disks on command. The current disk formats Vorteil support are
#### POST