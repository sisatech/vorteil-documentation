# Platforms

Platforms are a concept Vorteil uses to name their platforms. A platform could be any hypervisor we support which is Google Cloud Platform, Amazon Web Services, KVM, Virtualbox,  VMware Fusion/Player/Workstation and VSphere.

Vorteil Tools (such as the cli) are able to export to multiple disk formats like  VMDK, RAW, OVA and Google Cloud Platform.

Platforms can be run on most local hypervisors, remote cloud platforms and inhouse datacenters!

Vorteil Tools supports simple platform management. It provides a way to deploy to a Google Cloud Environment just as easy as running locally to KVM.

They provide the user a way to setup their hypervisor to run with certain settings. For an example I want to provision to KVM without the GUI; we could easily set the platform to use headless mode.

## See Also
- [Studio Platform Details](../../../studio/platforms)
- [CLI Run Using Platforms](../../../cli/general/run)
- [Graphql API Help](../../../api/graphql/graphql)