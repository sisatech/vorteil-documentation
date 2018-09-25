# /api/unpack/(operationID)
#### GET
The ability to fetch a tar archive that contains the contents of a project. After untarring you will be able to add this project to the GUI or run straight from the command line.
#### POST
The ability to unpack is started by sending a Graphql request to start the operation.

    mutation{
    unpack(germ:"marketplace:sisatech/helloworld"){
            job{
                id
            }
            uri
        }
    }

With the job id you would be able to subscribe to it to find out when the operation is completed. Being that we didn't have any injections with this request we would be able to GET the tar straight away as a stream.