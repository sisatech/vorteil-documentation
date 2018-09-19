# Starting a germination
The germination can be started by a sending Graphql request that resembles the request name. An example provided below for packing a project.

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