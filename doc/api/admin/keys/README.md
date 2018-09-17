# Admin
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