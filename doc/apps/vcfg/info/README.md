# Descriptive Metadata (info)

When it comes time to archive or publish an app, it is important that the app
has a sufficient amount of descriptive metadata so that it can be identified and
understood without other context, which may get lost or forgotten. The 'info' 
section of a VCFG is dedicated to such information. 

## Fields 

None of these fields can have any direct impact on the runtime for app. 
Indirectly, the Vorteil compiler may choose to use the app name defined here 
when generating a hostname for the app, if no hostname is defined in the 
[system](../system) section. 

Each of the following fields must appear within the info section, which can be
identified by the appearance of the following line. Everything from that line 
up until the appearance of another section is part of the info section.

```
[info]
```

### name

**Field type: String**

```
  name = "Hello, world!"
```

The name field provides a name for the app. This name can include spaces, 
punctuation, and special characters without any issues. It should not generally 
include version numbers as they can be set in the version field.

The name of an app should reflect the function of the app, not the name of the 
binary itself. In many cases the two are one-and-the-same, but in many other 
cases it is the contents of the app's file-system that determine its behaviour
and differentiate it from other apps that may well use exactly the same binary.

### author

**Field type: String**

```
  author = "Sisa-Tech Pty Ltd"
```

The author field identifies the author for the app. Like the name field, the 
author field should reflect the author of the main function of the app, not 
necessarily the binary itself. 

### summary

**Field type: String**

```
  summary = "A blurb or brief description without formatting."
```

The summary field provides a very brief blurb or summary of the app. It should 
not use any special characters or formatting, and it should be limited to a 
maximum of 280 characters (the maximum size of a Tweet). 

### description

**Field type: String**

```
  description = """# Hello, world!

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla tincidunt quam ut tempus tempus. Phasellus porta scelerisque tellus ut pharetra. Donec hendrerit dignissim nisl, in sollicitudin arcu congue id. Donec ac sagittis lacus, a dignissim libero. Donec varius sodales hendrerit. Aliquam tellus est, blandit et ante non, dictum viverra sem. Ut euismod dictum neque, vel fringilla nisl sollicitudin id.

Aenean ut nisl vel odio pellentesque lobortis et eu metus. Vestibulum at ligula vehicula, iaculis orci eu, dictum purus. Praesent sagittis velit vel porta volutpat. Phasellus et vulputate ligula. Cras euismod magna eu luctus laoreet. Morbi ut diam placerat, tincidunt quam et, tristique eros. Etiam luctus ultricies dictum. Nam facilisis mollis dolor, eu maximus est elementum bibendum.
"""
```

The description field can contain any amount of information you want. This may 
be detailed descriptions about its purpose, usage information, examples, 
marketing fluff, or all of the above. Its length is not limited like the summary
field is, and the use of formatting is encouraged. Some of our tools will 
attempt to render the description as Markdown, which can be exploited to make 
long descriptions more readable.

For this field in particular it is worth noting that TOML supports large, 
multiline strings with preserved formatting and whitespace by using triple
quotes instead of normal quotes, as per the example above.

### url 

**Field type: String**

```
  url = "github.com/sisatech/helloworld"
```

The url field can contain a URL to a page with more information about the app. 

### date 

**Field type: String**

```
  date = "18 Sep 18 11:22 AEST"
```

The date field can contain a timestamp identifying the release date. It is up to
the app creator to decide whether the date should reflect the release date of 
the app itself or of the underlying program. The choice can be confusing when 
a creator is merely converting existing software for use with Vorteil.

Vorteil tries its best to be understanding when it comes to parsing the value 
of this field. It will attempt to parse with many timestamp format standards as
well as several simple non-standard formats. Try writing what feels natural, and
if Vorteil returns an error then fallback to a common standard instead. 

### version

**Field type: String**

```
  version = "v1.0.0"
```

The version field can contain any type of version identifying string required.
None of Vorteil's tools attempt to parse or process this value, it exists 
purely as a human-readable version. 

## Notes

Because the description field can sometimes be very long, it can clutter up the
rest of the VCFG and make it difficult to peruse. If you have an app with a 
significant description field, it may be worth splitting the information section
out into a separate VCFG file and relying on [VCFG merging](../merging) to 
recombine them as required.

## See Also

* [What is a VCFG?](../introduction)
* [Example VCFG](../example)
* [CLI's Info Metadata Documentation](../../../cli/vcfgs/info)