# Learn about an app's advanced system settings (vorteil vcfgs system)

The "system" section of a VCFG contains fields that configure a number of
advanced system settings that usually don't need to be modified. Sometimes it
can be necessary to fiddle with some of these settings, so this section will
outline the options available here.

Here is the skeleton of the system configuration, as well as the associated
flags to change these values directly from the commandline.

[system]
delay = ""			→	--system.delay=""
disk-cache = ""		→	--system.disk-cache=""
dns = []			→	--system.dns="" 	(repeatable)
hostname = "" 		→	--system.hostname=""
max-fds = 0 		→	--system.max-fds=0
output-mode = ""		→	--system.output-mode=""
output-format = ""		→	--system.output-format=""

## Delay:

There are some environments where it may be desirable to pause the kernel
during its initialization process before it attempts to reach out over the
network. The system delay field can include a machine parsable string
specifying a duration to wait before proceeding with normal execution.

## Disk Cache:

The Vorteil kernel speeds up file I/O by reserving a pool of memory to use as
a cache for the app's file-system. This is not currently a dynamically scaling
value, so it is helpful to specify a good value here. Increasing the disk
cache size may improve performance up until it exceeds the capacity of the
app's file-system. The app's file-system capacity is not the same as the disk
size as there's always some amount of space on the image reserved for Vorteil
internals. The system disk cache field can include a number of bytes or a
string containing a number and units (e.g. "10 MiB").

## DNS:

The IP addresses to use for the domain name system can be configured here.
Currently only IPv4 is supported.

## Hostname:

The system's hostname can be specified here. If left unspecified the Vorteil
image compiler will attempt to produce a meaningful hostname based on the
app's name metadata (if specified), or fall back on "vorteil" if necessary.

## Maximum File Descriptors:

The maximum number of file descriptors available to the application can be
configured here. Provide a positive integer value to overrid the defaults.

## Output Mode:

The system's output mode toggles whether the kernel should print stdout and
stderr to the serial port and the display. Disabling either of these outputs
can improve performance, but they are both left enabled by default for their
usefulness when it comes to debugging.

"standard": the default prints stdout and stderr to both the display and to
the serial port.

"screen": prints stdout and stderr to the screen only (serial port disabled).

"serial": prints stdout and stderr to the serial port only (screen disabled).

"disabled": prints stdout and stderr to neither screen nor serial port.

## Output Format:

The system's output format defines which logic the kernel should use to
process and format anything written to stdout, stderr, and to any file-system
network redirects.

"standard": the default applies no formatting or transformations to the data,
sending it as-is.

"vlistener": the custom "vlistener" format prepends each line with the
virtual machine's hostname and a number identifying which file-system
redirect the line came from. These adjustments have a low performance
impact and can help logging servers to sort and filter the data.

Format pattern: "HOSTNAME#INDEX#FILENAME: MESSAGE"

Currently there are no other supported output formats, but this list will
expand in the future.

## Example:

[system]
delay = "0s"
disk-cache = "20mb"
dns = ["1.1.1.1", "1.0.0.1"]
hostname = "vapp"
max-fds = 1024
output-mode = "standard"
output-format = "standard"

### which is equivalent to:

--system.delay=0s --system.disk-cache=20mb --system.dns=1.1.1.1 \
--system.dns=1.0.0.1 --system.hostname=vapp --system.max-fds=1024 \
--system.output-mode="both" --system.output-format="default"


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 14-Sep-2018
