# Google Cloud Platform

Vorteil supports using a Google Cloud Platform environment as a target for provisioning images or deploying instances. Environments can be configured with the Command-Line Interface (CLI), Developer Studio, or directly through the GraphQL API.

## Behaviour
### Provisioning
Vorteil will upload a compatible disk to the bucket provided during setup. An image will be created from the uploaded file, before cleaning up the file from the bucket.

### Running
Vorteil will upload a compatible disk to the bucket provided during setup. An image will be created from the uploaded file, which will then be used to spawn a virtual machine instance. The image and the bucket file are then removed.

## Setup
For the setup process, users will require the following:

- bucket name (as listed on Google Cloud Platform)
- instance zone (created instances will be created in this zone)
- network (as listed on Google Cloud Platform; created instances will be connected to this)
- key file (.json service account key used to authorize/authenticate with Google Cloud Platform)

### Authentication / Authorization
The service account key provided during setup must have access to at least the following scopes:

- compute.disks.create
- compute.images.create
- compute.images.delete
- compute.images.get
- compute.images.list
- compute.images.update
- compute.images.useReadOnly
- compute.instanceGroupManagers.update
- compute.instanceTemplates.create
- compute.instanceTemplates.useReadOnly
- compute.instances.create
- compute.instances.setMetadata
- compute.instances.setTags
- compute.networks.get
- compute.subnetworks.use
- compute.subnetworks.useExternalIp
- storage.buckets.get
- storage.objects.create
- storage.objects.delete

## Command-Line Interface (CLI)
```bash
$ vorteil settings platforms new gcp MyGooglePlatform /PATH/TO/SERVICE/ACCOUNT/KEY.json --bucket mybucket --network default --zone australia-southeast1-a
```

## Developer Studio
With the Developer Studio client open, click the 'settings' button in the bottom-left of the window. In the 'Settings' tab, select 'google-cloud-platform' from the 'Add a Platform' dropdown list.

Users will then be prompted to input the require information into a form. Finish the process by clicking 'Add Platform'.

## GraphQL API
```
URL: http://localhost:7472/graphql
Method: POST

mutation {
    addPlatformGCP(name:"MyGooglePlatform",
    bucket: "mybucket",
    zone: "australia-southeast1-a",
    network: "default",
    key: "PATH/TO/SERVICE/ACCOUNT/KEY.json") {
        name
    }
}
```