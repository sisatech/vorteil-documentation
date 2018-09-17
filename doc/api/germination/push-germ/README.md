# /api/push/(operationID)

#### POST
The ability to push your project into a repository for others to access. Alternatively with this API you could push from repository A to Z. Provision and this request are different than the other germination strings as you don't get anything from the response if you were to GET it.

	mutation{
        push(germ:"/home/sisatech/Documents/projects/helloworld" injections:["abc"], bucket:"sisatech", app:"helloworld" node:"marketplace"){
                job{
                    id
                }
                uri
            }
	}

This allows us to have cool functionality like I want to update this binary on a repository. I can push to itself with a different binary remotely and it should update after uploading the binary. After the job progress has been completed the Application should now appear to where you pushed to.

Below is the appropriate headers for the binary. The body of this request should be the binary file by itself.

        Injection-Id   | abc
        Injection-Type | binary

