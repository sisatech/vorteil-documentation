# Build a Disk using the APIs

Building a disk image with Vorteil is simple process that can be completed in just 2-3 steps. The first step is to instruct Vorteil to build the disk from a [valid source](../../../../apps/general/buildable). Users can optionally define 'injections', for which the build process will wait until files for injection have been provided, or the request times out. After this, sending a GET request to the URI provided will yield the disk image:

```rest
mutation {
    build(germ: "marketplace:sisatech/helloworld", injections: ["replacementBinary"]) {
        job {
            id
        }
        uri
    } 
}
```
*Note that the 'injections' argument has been provided to the above mutation. Injections defined here instruct Vorteil to wait until a corresponding request has been made; this is typically used to files or configuration settings to the application at build-time. Don't include this argument if no injections are required.*

The injection defined above is labelled as "replacementBinary". To inject a binary file (which will overwrite the existing binary within the BUILDABLE source), a POST request can be sent to the URI provided in the response from the first request.

```rest
URL: http://localhost:7472/URI_FROM_RESPONSE
Method: POST

Headers:
    Injection-Type: binary
    Injection-Id: replacementBinary

Body:
    Binary file.
```

When satisfied, users may request the disk image by sending a GET request to the same URL used in the previous step.
