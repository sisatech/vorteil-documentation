# File-System Redirects (redirects)

The Vorteil kernel includes functionality to support redirecting data that would
be written to a file to instead be sent out to a third-party server via UDP.
This can be an extremely helpful approach to gathering logging information from
Vorteil virtual machines, since it's not possible to SSH onto them and 
preserving the virtual disk images for mounting and manual inspection can be 
very tedious. 

## Fields

Unlike most other sections of a VCFG, the redirects section is a user-defined 
map. This looks just like any other section with string-valued fields, but there
are no pre-defined fields at all. The developer defines their own fields, and 
the field name will be used as a path prefix, whilst its value will be the 
address of a UDP server listening for the data.

Example:

```
[redirects]
  "stdout" = "192.168.1.10:7000"
  "/logs" = "192.168.1.10:7000"
```

The reserved "stdout" can be used to apply these redirects from the app's 
stdout, stderr, and kernel logs instead of a file.

## Notes

In this section it is worth noting that although not normally required, it is 
acceptable to use quotation marks around field names in TOML. This can be a safe
practice to follow to prevent accidental bugs caused by creating fields with 
characters that are not legal in TOML without quotes.

Redirecting files will not typically be useful unless metadata is attached to 
the data so that the receiving server can process and organize it. See the 
output-format field in the [system](../system) section for how to change this
behaviour.

## See Also

* [What is a VCFG?](../introduction)
* [Example VCFG](../example)
* [CLI's Redirects Documentation](../../../cli/vcfgs/redirects)