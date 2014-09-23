This is attempt at creating Go bindings for OpenTransactions (opentxs).

# Instructions

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
# in gotxs/
go install
````


# Development

The bindings consist of three separate packages:

## Package *gotxs/opentxs*

This cointains a SWIG-generated module generated from an interface file. It
exports a few hundred functions that are translated from C++, which can be seen
in [this text file](opentxs/opentxs.go).

Low-level API functions start with the prefix `OTAPI_Wrap` and are static
methods of the C++ class with the same.

A higher-level API that provides network synchronization is available through
instances of the class `OT_ME` (*OpenTransactions Made Easy*).

The general rule is that the `OT_ME` bindings should be used when possible.

Unfortunately, the exported methods are difficult to use directly:

* The methods have no documentation
* There are initialization steps that need to be executed
* Some of the called methods unexpectedly crash on invalid arguments (process
  hangs indefinitely due to the `OT_FAIL` macro)
* The methods sometimes use special return values, like the empty string or `-1`
  to signal error conditions. This is contrary to the Go way of signaling
  errors via multiple return values.

## Packages *gotxs* and *gotxs/easy*

In order to fix this issues, we wrap the SWIG-generated methods by hand in our
own packages. All methods exported by `OTAPI_Wrap` go to the `gotxs` packages.
The higher-level methods exposed by the `OT_ME` go to `gotxs/easy` where we
instantiate a single class instance.

### Naming convention

For Go-exported methods, we simply copy the name and translate it to CamelCase
with a leading capital.

### Documentation

Each method should provide documentation that should go beyond *calls underlying
method foobar*. When the documentation isn't available from the wrapped C++
methods, the implementation needs to be examined.

### Input sanitation

This is difficult to do upfront. When we discover that something causes opentxs
to crash hard unexpectedly, we should fix it on the Go side.

### Return values

The C++ methods often encode error conditions in the return value. The
Go-wrappers translate these into multiple return values.  For now, we use simple
conditionals (`if retval == ""`) to check the C++ return values. The error
should be descriptive and explain what went wrong.

See [Error handling and Go](http://blog.golang.org/error-handling-and-go) for
more information.

### Code Style

We use `go fmt` to format the source.

### Navigating the C++ code

I recommend Doxygen with the call and caller graphs for browsing the opentxs C++
code

```
# in opentxs/
# install doxygen and graphviz

cmake .. -DDOC_FULLGRAPHS=YES
make doc
```
