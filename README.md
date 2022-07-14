
# common-functions

This common go module is used to store all the functions and models that will be used across personal applications.

## Programming Standards

If you are contributing to this you need to follow the following standards:
 - create files per set of models
 - code coverage must be maintained over 90% preferably at 100%
 - do not commit: binaries, videos, IDE, or OS specific files
 - when updating version follow major.minor.patch standard (ie 1.2.3)

## Usage

This module is currently intended for private use, use of this module outside of any Benjamin Abrams application is strictly prohibited

### Makefile Use

Makefiles currently are only setup to get dependencies of the common module.

```shell
make get
```

### local use

Run the following command to use a local version of this module

```shell
go mod edit -replace github.com/geeksheik9/common-functions=RELATIVE_PATH_TO_LOCAL_MODULE
```

When done testing run the following command to return to hosted version

```shel
go mod edit -dropreplace github.com/geeksheik9/common-functions
```


