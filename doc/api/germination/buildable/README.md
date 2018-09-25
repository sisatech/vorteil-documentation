# /api/build/(operationID)
#### GET
The ability to build virtual disks on command whether it is from a repository or local project setup. The current disk formats Vorteil support are OVA, RAW, GCP, VMDK and Stream-opt VMDK. The disk will be streamed back on the request.
#### POST

    mutation{
        build(germ:"marketplace:sisatech/helloworld" injections:["binarytest"]){
            job{
                id
            }
            uri
        }
    }

The job id can be used to track the progress of the operation. The progress will only begin when all injections have been provided as you can have multiple. Notice we are using a binary injection as we want to use the application on the marketplace but build it with a different binary to run. Attaching the binary to the POST request with the headers below will result in building "Helloworld" with a separate binary.
    
    Injection-Id   | binarytest
    Injection-Type | binary
    
Upon sending this request off the response back should be of status code 200. Then you will be able to GET the same address for the file that you have just created.

