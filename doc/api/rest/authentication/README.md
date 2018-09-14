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