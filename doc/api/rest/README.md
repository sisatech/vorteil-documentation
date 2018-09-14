# Rest

Information about every REST api our daemon has. The REST apis are used for certain operations that Graphql is unable to do.

## Admin
The default port for the admin server is normally 7471.
### /api/keys
#### GET
Sending a get request to this api will respond with a list of keys currently active on the daemon.
##### Response
The response will contain a list key objects.
    
        [
            {
                "id": "ckmxjzemmwsd",
                "username": "macConnect",
                "created": 1536270890,
                "primarygroup": "root",
                "groups": [
                    "root",
                    "public"
                ],
                "comment": "Key for mac"
             }
        ]
    
#### POST
Sending this request with the body for user details will create a new key.
##### Body
    {
        "username":"John",
        "primarygroup":"root",
        "groups":["root", "test-john"],
        "comment":"description about this key access."
    }
##### Response
    {
        "key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoie1wiaWRcIjpcImhkc25jaHppeGx3Z1wiLFwidXNlcm5hbWVcIjpcIkpvaG5cIixcImNyZWF0ZWRcIjoxNTM2NzkwNzc3LFwicHJpbWFyeWdyb3VwXCI6XCJyb290XCIsXCJncm91cHNcIjpbXCJyb290XCIsXCJ0ZXN0LWpvaG5cIixcInB1YmxpY1wiXSxcImNvbW1lbnRcIjpcImRlc2NyaXB0aW9uIGFib3V0IHRoaXMga2V5IGFjY2Vzcy5cIn0iLCJzYWx0IjoiUnJBdVliaDA2OXpjR0ZyN2dvSXVhV0RKcjdxamRGQ3MyVk1zUU9YOU5kQT0ifQ.h-I1PEV1i35hmqJOpX1yNHuhf-YCcoKXkq2ij4waSlU"
    }
### /api/keys/(key)
Append the ID of a key object to the end of this API.
#### GET
Gets all the information about the key object.
##### Response
      {
                "id": "ckmxjzemmwsd",
                "username": "macConnect",
                "created": 1536270890,
                "primarygroup": "root",
                "groups": [
                    "root",
                    "public"
                ],
                "comment": "Key for mac"
      }
#### DELETE
Deletes the key
##### Response
No response is returned. If it responds with a 200 the key object has been successfully deleted.

## Authentication
Authentication with the repository mainly using the port 7472.
### /api/auth
#### GET
Sending a get request here checks if you are authenticated with the repository. Basically checking if the key for authentication exists and is not corrupted.
##### Response
The response to check if you are authenticated is empty. The status code of the request lets you know instead.
### /api/logout
#### POST
Sending a post request here will clear the cookie and prevent any other access using the apis till you have logged in.
##### Response
A 200 status code tells you that you have successfully logged out of the repository.
### /api/login
#### POST
Sending a post request with a JSON body containing a key.
##### Body
        {
            "key":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoie1wiaWRcIjpcImhkc25jaHppeGx3Z1wiLFwidXNlcm5hbWVcIjpcIkpvaG5cIixcImNyZWF0ZWRcIjoxNTM2NzkwNzc3LFwicHJpbWFyeWdyb3VwXCI6XCJyb290XCIsXCJncm91cHNcIjpbXCJyb290XCIsXCJ0ZXN0LWpvaG5cIixcInB1YmxpY1wiXSxcImNvbW1lbnRcIjpcImRlc2NyaXB0aW9uIGFib3V0IHRoaXMga2V5IGFjY2Vzcy5cIn0iLCJzYWx0IjoiUnJBdVliaDA2OXpjR0ZyN2dvSXVhV0RKcjdxamRGQ3MyVk1zUU9YOU5kQT0ifQ.h-I1PEV1i35hmqJOpX1yNHuhf-YCcoKXkq2ij4waSlU"
        }
##### Response
No response is returned based on status codes. If it was successful a 200 status code will be returned.

## Version
A rest API that allows you to get the version of the daemon. This request is normally sent to the port 7472.
### /version
#### GET 
Gets the version information for the server that receives this request.
#### Response
    {
        "Name": "server",
        "Major": 2,
        "Minor": 0,
        "Patch": 0,
        "Date": "2018-09-12T13:40:40+10:00",
        "ID": "2bedd2ec",
        "Dirty": true,
        "Modules": [
            {
                "Name": "api",
                "Major": 0,
                "Minor": 1,
                "Patch": 0,
                "Date": "0001-01-01T00:00:00Z",
                "ID": "",
                "Dirty": false,
                "Modules": null,
                "Contracts": null
            },
            {
                "Name": "compiler",
                "Major": 2,
                "Minor": 0,
                "Patch": 0,
                "Date": "0001-01-01T00:00:00Z",
                "ID": "",
                "Dirty": false,
                "Modules": null,
                "Contracts": [
                    {
                        "Name": "kernel",
                        "Major": 0,
                        "Minor": 3,
                        "Patch": null
                    },
                    {
                        "Name": "package",
                        "Major": 2,
                        "Minor": null,
                        "Patch": null
                    }
                ]
            }
        ],
        "Contracts": null
    }
## Daemon Objects
#### /api/signed/(signedString)
Anything related to a Fragment object will contain a signed URL. These allow us to download any file. An easy way to find a signed url would be to send this Graphql request.

    {
        listBuckets{
            edges{
                node{
                    name
                    icon{
                        downloadURL
                        id
                    }
                }
            }
        }
    }
This will return an object like below.

    {
	"data": {
		"listBuckets": {
			"edges": [
				{
					"node": {
						"icon": {
							"downloadURL": "/api/signed/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdHJpbmciOiJ7XCJLZXlcIjpcImRvd25sb2FkXCIsXCJNZXRob2RcIjpcIkdFVFwiLFwiRXhwaXJ5XCI6XCIyMDE4LTA5LTE0VDA4OjU0OjA5LjA3NDYyOTgxOCsxMDowMFwiLFwiQXJnc1wiOlwiZXlKR2NtRm5iV1Z1ZEVsRUlqb2lNemd3TWpJd056VTRPRE00TkRRME1ETXpJaXdpVlhObGNrbHVabThpT25zaVZYTmxjbTVoYldVaU9pSnliMjkwSWl3aVVISnBiV0Z5ZVVkeWIzVndJam9pY205dmRDSXNJa2R5YjNWd2N5STZXeUp5YjI5MElpd2ljSFZpYkdsaklsMTlmUT09XCJ9In0.Ymtnbn4mKJO5t1JXAC5CaQ8L4nXDrLLafAr94dBhB-Q",
							"id": "380220758838444033"
						},
						"name": "sisatech"
					}
				},
        }
    }

The downloadURL of that field is a 'signedString'. Entering localhost:7472/api/signed/signedString in your browser should display what that object would be.
## Process Operations
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