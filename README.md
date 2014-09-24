This is attempt at creating Go bindings for OpenTransactions (opentxs).

#Instructions

In order install the bindings, you need to link the necessary header files from
`opentxs` to the `include/` subdirectory.

````
cd gotxs/opentxs/
ln -s $HOME/path/to/opentxs/deps/containers include/
ln -s $HOME/path/to/opentxs/include/opentxs include/
# fix the paths in Makefile
make install
````

Then you can install the `gotxs` package

````
go install
````

## Mac OSX

There are some addional changes required to the Makefile to run this on a Mac:

1. LIBPATH should point to the lib/ dir of your opentxs install.
2. GOINCLUDE points to /usr/local/Cellar/go/1.3.1/libexec/pkg/darwin_amd64 (or what ever version you have).
3. LIBNAME is libopentxs-golang.dylib.
4. Add the magic '-flat_namespace' and '-undefined suppress' to the g++ line.

Then run 'make'. Ignore the ldconfig error, it's a linux thing.
Then export an environment variable to point to the installed dylib:

```
$ set -xg DYLD_LIBRARY_PATH ~/opentxs/lib/
```

Possibly DYLD_FALLBACK_LIBRARY_PATH is better.
