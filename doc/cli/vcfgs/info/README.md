# Learn about an app's informational metadata (vorteil vcfgs info)

Being able to identify a package is important if you ever plan on distributing
it to other users. Even without any such plans, it can be very useful for
organizing packages and remembering what they are at a later date. With Vorteil
it is possible to store helpful application metadata within the configuration
for the app itself.

Here is the skeleton of the "info" configuration, as well as the associated
flags to change these values directly from the commandline.

[info]
name = ""   		→	--info.name=""
author = ""	  	        →	--info.author=""
summary = ""		→	--info.summary=""
description = "" 		→	--info.description=""
url = ""			→	--info.url=""
date = ""			→	--info.date=""
version = ""		→	--info.version=""

By convention, if the application is a powerful general purpose app like a
database server then fields in the information section should describe the
underlying application. This means that the author should be the author of the
original application rather than the package creator, the version should be the
version of the underlying application rather than the package version, and so
on.

On the other hand, if the application is a sort of skeleton application that
only becomes useful based on its configuration or the contents of its
file-system the package itself should be the focus of this section. An example
of this would be a generic web server customized to serve a particular website.
Packages built using programs run on interpreted languages should avoid
describing the runtime itself as the app.

Tip: the description field should often contain large amounts of information
and be sensitive to formatting (markdown is common), so remember that in
TOML it is possible to provide large multi-line strings by using triples of
double-quotes. i.e. """Like this"""

## Example:

[info]
name = "My App"
author = "Sisa-Tech Pty Ltd"
summary = "My short blurb description."
description = "My much longer documentation description."
url = "vorteil.io"
date = "2018-05-16T08:28:06.801064-04:00"
version = "1.0.0"

### which is equivalent to:

--info.name="My App" --info.author="Sisa-Tech Pty Ltd" \
--info.summary="My short blurb description." \
--info.description="My much longer documentation description." \
--info.url=vorteil.io --info.date="2018-05-16T08:28:06.801064-04:00" \
--info.version=1.0.0


###### Auto generated for CLI 2.0.0-7169db7d-dirty on 14-Sep-2018
