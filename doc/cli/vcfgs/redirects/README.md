# Learn about Vorteil's file-system redirects (vorteil vcfgs redirects)

Collecting the logs from an app can be an invaluable part of a deployment.
Vorteil virtual machines cannot be accessed via SSH, and with their disposable
nature it can be confusing to figure out what is happening live. Vorteil
supports something we call file-system redirects for this purpose.

File-system redirects allow the Vorteil kernel to take any data that would
normally be written to a file and instead send it out over the network to a
third-party logging server. Up to four file-system redirects can be defined for
an app.

The redirect configuration does not have a fixed skeleton like other sections.
It is a map where the app developer defines their own keys. Keys are treated as
file path prefixes and values are network addresses, either as IP:PORT or
DOMAIN:PORT. If the app would write data to a file path that begins with the
given prefix then it will instead be redirected over UDP. Prefix matching makes
it easy to support rotating logs.

"stdout" is a reserved key for all data written to stdout and stderr by the
app and the kernel.

Note that TOML allows field keys to appear in quotes, which is a useful safety
measure here since not every filepath is a valid TOML key.

Tip: when using file-system redirects it is often helpful to modify the output
so that logging servers can identify the source and filter accordingly.
See the additional help topic "vorteil vcfgs system" and the information
relating to the output-format field to learn how to have the kernel take
care of this for you.

## Example:

[redirects]
stdout = "10.0.0.2:8000"
"/logs" = "10.0.0.2:8001"

### which is equivalent to:

--redirects="stdout=10.0.0.2:8000" --redirects="/logs=10.0.0.2:8000"


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 14-Sep-2018
