# Vorteil Developer Studio Not Opening

## Symptoms

- User is running on a Linux operating system.
- The Vorteil Developer Studio application does nothing when attempting to open (zero feedback, no splashscreen).

## Cause 

The Vorteil Developer Studio is an electron app, and depends on certain shared libraries to run. As of version 18.04, Ubuntu seems to ship without the 'libgconf-2.so.4' library. This can be confirmed by attempting to manually launch the Studio client through the terminal, and analysing the output:

```bash 
$ /opt/vorteil/bin/gui/client
```

## Solution

### Find and install missing shared libraries

Users will need to locate and install 'libgconf-2.so.4' on their system to resolve this issue. On systems with 'apt', this can be installed with the following command:

```bash
$ apt -y install libgconf2-4
```