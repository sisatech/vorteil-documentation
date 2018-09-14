# Learn about an app's network configuration (vorteil vcfgs network)

The network configuration is often an important part of a Vorteil application.

Vorteil supports up to four network interfaces at once, and will default to
automatically configuring them using DHCP unless specific instructions to do
otherwise are found. A VCFG can include up to four "network" sections, with each
corresponding to a single network interface. The Vorteil kernel applies network
configurations to interfaces in the same order that they were defined, which
makes the behaviour predictable but also requires the deployer to know exactly
what network to connect to each interface.

Here is the skeleton of a network configuration, as well as the associated flags
to change these values directly from the commandline. Assume that this skeleton
represents the 0th network configuration:

[[network]]
ip = ""			→	--network.0.ip=""
mask = ""			→	--network.0.mask=""
gateway = ""		→	--network.0.gateway=""
mtu = 0 			→	--network.0.mtu=0
disable-tso = false 	→ 	--network.0.disable-tso=false
udp = [] 			→	--network.0.udp="" 	(repeatable)
tcp = []			→	--network.0.tcp="" 	(repeatable)
http = []			→	--network.0.http="" 	(repeatable)
https = []			→	--network.0.https="" 	(repeatable)

The most important parts of the network configuration are the ip, mask, and
gateway fields. If the ip is left as an empty string or explicitly set to "dhcp"
then the network interface will configure itself using DHCP, and ignore the mask
and gateway fields. For any other values the ip will be parsed as a static IPv4
address, and the mask and gateway fields will be applied. IPv6 is not currently
supported.

## Port Use:

The upd, tcp, http, and https fields are arrays that should list every port
the app intends to listen on using each respective protocol. There is no need
to double up ports from one protocol to another, i.e. don't bother adding tcp
port 80 if you've already added http port 80.

These port lists do not directly effect compiled Vorteil disk images and so
they are technically optional, but there is still a lot of value in populating
them. Firstly, they are helpfully descriptive about the behaviour of your app,
which can help customers to understand what they do. Secondly, provisioning
and deployment stacks can use this information to configure firewalls or NAT
rules dynamically for the app without human intervention.

## Advanced Settings:

The mtu field can be used to configure the maximum transmission unit for the
interface. This is an advanced setting that should not be tampered with unless
you know exactly what number to use. Using a number too low will reduce
network performance, using a number too high will usually result in fatal
packet loss.

The "disable-tso" field accepts a boolean value and disables the TCP
Segmentation Offload behaviour of the network interface if true. TCP
segmentation offloading improves the VM's performance but depending upon the
infrastructure environment it can in some cases add a significant performance
cost elsewhere.

## Example:

[[network]]
ip = "10.0.0.2"
mask = "255.255.255.0"
gateway = "10.0.0.1"
mtu = 1500
disable-tcp-offload = false
udp = []
tcp = ["8000", "8080"]
http = ["80"]
https = []

### which is equivalent to:

--network.0.ip=10.0.0.2 --network.0.mask=255.255.255.0 \
--network.0.gateway=10.0.0.1 --network.0.mtu=1500 --network.0.udp="" \
--network.0.tcp="8000" --network.0.tcp="8080" --network.0.http="80" \
--network.0.https="" --network.0.disable-tcp-offload=false


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 14-Sep-2018
