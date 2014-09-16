This is attempt at creating Go bindings for OpenTransactions (opentxs).

Instructions
============

In order install the bindings, you need to link the necessary header files from
`opentxs` to the `include/` subdirectory.

````
cd gotxs/opentxs/
mkdir include/
ln -s $HOME/path/to/opentxs/deps/ include/
ln -s $HOME/path/to/opentxs/include/opentxs/ include/

cd ..
go install
````

