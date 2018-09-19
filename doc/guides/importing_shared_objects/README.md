# Importing Shared Objects
Applications depend on various shared objects (libraries) in order to run; because of this, they will need to be included on the Vorteil application filesystem.

Users working on Linux will benefit from the use of the ```import-shared-objects``` tool. This tool analyses the project ELF binary and copy the required shared objects into the project directory. Copied files will be placed in either './lib' or './lib64'.

Note that some applications may depend on libraries not detailed within the ELF binary header. In this case, the ```import-shared-objects``` tool will not retrieve all libraries, and users will need to locate remaining files manually.

## Command-Line Interface
```bash
$ vorteil projects import-shared-objects
```

## Developer Studio
Right-click on the project in the projects list panel, and select 'Import Shared Objects' from the resulting context menu.

## GraphQL API
```
URL: http://localhost:7472/graphql
Method: POST

mutation {
    importSharedObjects(project:String)
}
```

