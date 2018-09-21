# VMWare vCenter

Vorteil supports using a VMWare vCenter environment as a target for provisioning images or deploying instances. Environments can be configured with the Command-Line Interface (CLI), Developer Studio, or directly through the GraphQL API.

## Behaviour
### Provisioning
Vorteil will upload an OVA to the target environment, and mark it as a virtual machine template.

### Running
Vorteil will upload an OVA to the target environment, and mark it as a virtual machine.

## Setup
For the setup process, users will require the following:

- API Endpoint Address (ie. ```https://10.0.0.15```)
- Username
- Password
- Datacenter name
- Network name (virtual machines will connect to this network)
- Datastore name
- Cluster name

## Command-Line Interface (CLI)
```bash
$ vorteil settings platforms new vcenter myVCenter --address https://example.com --cluster exampleCluster --datacenter exampleDatacenter --datastore exampleDatastore --network exampleNetwork --password examplePassword --username exampleUsername
```

## Developer Studio
With the Developer Studio client open, click the 'settings' button in the bottom-left of the window. In the 'Settings' tab, select 'vmware-vcenter' from the 'Add a Platform' dropdown list.

Users will then be prompted to input the require information into a form. Finish the process by clicking 'Add Platform'.

## GraphQL API
```
URL: http://localhost:7472/graphql
Method: POST

mutation {
    addPlatformVcenter(name:"myVCenter",
    address: "https://example.com",
    cluster: "exampleCluster",
    network: "exampleNetwork",
    username: "exampleUsername",
    password: "examplePassword",
    datacenter: "exampleDatacenter") {
        name
    }
}
```