Ethereal is a command line tool for managing common tasks in Ethereum.  It is designed to be

It is designed to allow integration in to batch scripts and to be called from other programs by providing concise information and through use of the quiet mode where results are translated in to exit codes.

## Installing Ethereal

Ethereal requires [Go](https://golang.org/).  With Go installed you can install Ethereal by running `go get github.com/wealdtech/ethereal`

## Using Ethereal
### Access to local wallets

Ethereal works with Geth, MIST and Parity wallets in the standard locations.  A simple way to check the addresses that can be seen by Ethereal is to run `ethereal account list` which will list all accounts that Ethereal can see.  If you expect an address to show up and it doesn't then please raise an issue with the relevant details.

### Access to Ethereum networks

Ethereal supports all main Ethereum networks  It auto-detects the network by querying the connected node for the network ID.  The connection should be geth-compatible, so either geth itself or parity with the `--geth` flag to enable geth compatibility mode.  The connectino could be a local node or a network service such as Infura.

## Examples

### Increase the gas price for transaction
You have submitted a transaction to the network but it's taking a long time to process because the gas price is too low.

```
ethereal transaction up --transaction=0x5219b09d629158c2759035c97b11b604f57d0c733515738aaae0d2dafb41ab98 --gasprice=20GWei --pasphrase=secret
```
where `transaction` is the hash of the pending transactions, `gasprice` is the price you want to set for gas, and `passphrase` is the passphrase for the account that sent the transaction.

### Cancel a transaction
You have submitted a transaction to the network by mistake and want to cancel it.
```
ethereal transaction cancel --transaction=0x5219b09d629158c2759035c97b11b604f57d0c733515738aaae0d2dafb41ab98 --pasphrase=secret
```
where `transaction` is the hash of the pending transactions and `passphrase` is the passphrase for the account that sent the transaction.

### Sweep Ether
You want to transfer all Ether in one account to another.
```
ethereal ether sweep --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --to=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d --pasphrase=secret
```
where `from` is the address from which the Ether will be transferred, `to` is the address to which the Ether will be transferred, and `passphrase` is the passphrase for the `from` account.

### Transfer a token
You want to transfer a token to another account.
```
ethereal token transfer --token=omg --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --to=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d --amount=10.2 --pasphrase=secret
```
where `token` is the token to transfer, `from` is the address from which the token will be transferred, `to` is the address to which the token will be transferred, `amount` is the amount of the token to transfer, `gasprice` is the price you want to set for gas, and `passphrase` is the passphrase for the `from` account.

*Please note that before using a token name such as 'omg' you should confirm that the contract address matches the expected contract address by using `ethereal info --token=omg` or similar.*
