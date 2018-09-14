## Signed URIs
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