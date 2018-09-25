# Amazon Web Services

Vorteil supports using a Amazon Web Services environment as a target for provisioning images or deploying instances. ENvironments can be configured with Command-Line Interface (CLI), Developer Studio, or directly through the GraphQL API.

## Behaviour
### Provisioning
Vorteil will upload a compatible disk to the bucket provided during setup. An image will be created from the uploaded file, before cleaning up the file from the bucket.

### Running
Vorteil will upload a compatible disk to the bucket provided during setup. An image will be created from the uploaded file, which will then be used to spawn a virtual machine instance. The image and the bucket file are then removed.

## Setup
For the setup process, users will require the following:

- bucket name (as listed on Amazon Web Services platform)
- region zone (instances will be created in this zone)
- key id (used to identify who is accessing the service)
- key secret (used to authenticate who is accessing the service)

## Command-Line Interface (CLI)
```bash
$ vorteil settings platforms new aws MyAWS --bucket mybucket --region US-EAST-1 --key-id ID OF AWS KEY --secret SECRET KEY STRING
```

## Developer Studio
With the Developer Studio client open, click the 'settings' button in the bottom-left of the window. In the 'Settings' tab, select 'amazon-web-services' from the 'Add a Platform' dropdown list.

Users will then be prompted to input the require information into a form. Finish the process by clicking 'Add Platform'.

## GraphQL API
```
URL: http://localhost:7472/graphql
Method: POST

mutation {
    addPlatformAWS(name:"MyAws",
    bucket: "mybucket",
    region: "US-EAST-1",
    key-id: "ID OF AWS KEY",
    secret: "SECRET KEY STRING
}
```