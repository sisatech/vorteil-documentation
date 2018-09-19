# Application Version Control
The Vorteil Repository automatically handles most aspects of application version control; it will generate a hash for the packaged application based on it's contents and use this as the package indentifier. If an application is subsequently uploaded to the same location, one of the following two scenarios will occur:

- The package identifiers will conflict (identical package), and the upload will not complete.
- The package identifiers will not conflict, and a new version will be added alongside the existing version(s).

In some cases, it may be prudent to 'tag' application versions for increased usability / clarity in the future. 

```bash
$ vorteil repositories tag BUCKET/APP[/REF] TAG
```
*If 'REF' is not provided, the latest version of the application will be targeted.*

A version 'REF' can be obtained by running the 'list' subcommand.

```bash
$ vorteil repositories list BUCKET/APP
```