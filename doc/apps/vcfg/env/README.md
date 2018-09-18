# Environment Variables (env)

Environment variables are used by many apps as an alternative to arguments. 
A Vorteil app typically has far fewer environment variables than the average 
Linux shell. With the exception of the variables 
[provided by the kernel](../../../runtime/environment/cloud), only variables
explicitly added to the VCFG will be available to an app. 

## Fields

Unlike most other sections of a VCFG, the env section is a user-defined map. 
This looks just like any other section with string-valued fields, but there are
no pre-defined fields at all. The developer defines their own fields, and the
field name will be used as the environment variable key, and its value will 
become the variable's value.

Example:

```
[env]
  HOME = "/home"
  "USER" = "root"
```

## Notes

In this section it is worth noting that although not normally required, it is 
acceptable to use quotation marks around field names in TOML. This can be a safe
practice to follow to prevent accidental bugs caused by creating fields with 
characters that are not legal in TOML without quotes.

## See Also

* [What is a VCFG?](../introduction)
* [Example VCFG](../example)
* [CLI's Binary and Arguments Documentation](../../../cli/vcfgs/invocation)