# Fragments
Some objects on the Vorteil Repository have one or more 'fragments'. This simply refers to a single file within the larger object (eg. an icon within a bucket, or a .vcfg file within an application version).

The Command-Line Interface and Developer Studio both shield the user from this, using or displaying fragments when the situation requires it. The GraphQL API can be used for this purpose.

```
URL: http://localhost:7472/graphql
Method: POST

query {
    bucket(name:String){
        icon{
            downloadURL
        }
        app(name:String){
            file{
                downloadURL
            }
            icon{
                downloadURL
            }
        }
    }
}
```

Each fragment returns a 'downloadURL', which can be added to the repository host address to get a direct download link for that file via a GET request (eg. http://localhost:7472/api/signed/0123456789ABCDEF).

## See Also
- [Fragments (API)](../../../api/fragments/signeduris)
