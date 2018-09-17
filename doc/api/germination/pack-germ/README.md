# /api/pack/(operationID)
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
