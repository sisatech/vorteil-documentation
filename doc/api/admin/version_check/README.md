# Version
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