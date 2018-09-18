# Advanced Runtime Settings (system)

Some settings are rarely meant to be altered, but still need to be alterable.
The 'system' section of a VCFG is dedicated to advanced settings effecting the 
app's runtime that should rarely be needed.

## Fields

Each of the following fields must appear within the system section, which can be
identified by the appearance of the following line. Everything from that line 
up until the appearance of another section is part of the system section.

```
[system]
```

### delay 

**Field type: String**

```
  delay = "3s"
```

The delay field specifies a length of time that the kernel should pause for 
before attempting to connect to the network(s). Since networks are set up before
the app is launched, this also delays the execution of the app. This setting 
exists because some environments aren't ready to process network traffic from a 
virtual machine immediately after it is launched. Vorteil boots up much more 
quickly than these environments expect, but this setting can be used to slow it
down.

### disk-cache 

**Field type: String**

```
  disk-cache = "2 MiB"
```

The disk-cache field defines the amount of RAM that should be dedicated to a 
disk cache. A larger disk cache allows the kernel to provide more perfomant 
access to the app's file-system through a combination of pre-loading and less 
frequent flushing. The field accepts plain numbers, which will be interpreted as
bytes, or strings which can include a number followed by a unit, as per the 
above example.

### dns 

**Field type: String Array**

```
  dns = ["1.1.1.1", "1.0.0.1"]
```

The dns field defines the list of DNS servers the kernel should use to resolve
domain names. Each entry in the list must be a valid IPv4 address. 

### hostname 

**Field type: String**

```
  hostname = "helloworld"
```

The hostname field determines the app's hostname on the local network. 

### max-fds

**Field type: Unsigned Integer**

```
  max-fds = 1024
```

The max-fds field determines the maximum number of file-descriptors available to
the app.

### output-mode

**Field type: String**

```
  output-mode = "standard"
```

The output-mode field makes it possible to alter the way the kernel displays or
outputs information. By default, anything written to stdout or stderr, or by the
kernel is printed to the virtual machine's serial port and displayed on its 
screen. Other options allow the developer to disable either or both of those 
outputs, potentially improving performance. 

Acceptable values:

* standard
* screen
* serial
* disabled

### output-format

**Field type: String**

```
  output-format = "standard"
```

The output-format field can instruct the kernel to alter the way it formats and
attaches metadata to each line of output that would be sent out through 
[redirects](../redirects), stdout, stderr, or kernel logs. Right now there are
only two modes, the plain unmodified format, and a custom format we've created
to attach metadata that a third-party logging server could process to filter and
organize logs. 

Acceptable values:

* standard
* vlistener

## Notes 

Unless you know why you're doing it, all of these fields should generally be 
omitted from your VCFGs.

## See Also 

* [What is a VCFG?](../introduction)
* [Example VCFG](../example)
* [CLI's Advanced System Config Documentation](../../../cli/vcfgs/system)