
# /api/provision/(operationID)

#### POST
This API provides us the ability to provision an application on run time with added files, binary or VCFG. This grants us  ways to run an application from a repository but with a set of new files imported.

	mutation{
       provision(germ:"/home/sisatech/Documents/projects/helloworld" injections:["files"], platform:"kvm", kernelType:"", name:"kvm-prov-name", start:"true"){
            job{
                id
            }
            uri
        }
	}

The fields required are germ, platform, start.
    start    - tells the daemon to run this virtual machine when provisioned.
    germ     - tells the daemon what we are provisioning whether it be from a repository or project.
    platform - tells the daemon what platform we would like to provision to.

The optional fields for this Graphql request are injections, kernelType and name.
    injections - what you would like to inject into the operation. Binary, files and configuration
    kernelType - whether you want to run this with the debug or production kernel.
    name       - The ability to name the virtual machine the daemon deployed.

Upon sending the Graphql request, you should now be able to send a post request to the above API to tell the daemon these are the files you want to inject. The progress of the provision itself can be easily tracked by subscribing to the Job ID.

