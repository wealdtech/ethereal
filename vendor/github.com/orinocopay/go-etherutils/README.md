This package contains a number of command-line utilities and libraries useful when working with Ethereum and Go.

# ens

ens is a command-line utility that allows users to manage resolvers and addreses for ENS names.

To build ens from source run `go build` from the `cmd/ens` subdirectory.

## Sample usage

### Obtain the resolver for a name

`ens resolver myname.eth`

### Set the resolver for a name

`ens resolver set myname.eth --passphrase="my secret passphrase"`

### Obtain the address for a name

`ens address myname.eth`

### Set the address for a name

`ens address set myname.eth --passphrase="my secret passphrase" --address=0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1`

Further details about ens usage can be obtained with `ens help`
