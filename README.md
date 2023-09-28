# ethereal

[![Tag](https://img.shields.io/github/tag/wealdtech/ethereal.svg)](https://github.com/wealdtech/ethereal/releases/)
[![License](https://img.shields.io/github/license/wealdtech/ethereal.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/wealdtech/ethereal?status.svg)](https://godoc.org/github.com/wealdtech/ethereal)
[![Travis CI](https://img.shields.io/travis/wealdtech/ethereal.svg)](https://travis-ci.org/wealdtech/ethereal)

A command-line tool for managing common tasks in Ethereum.

## Table of Contents

- [Install](#install)
  - [Binaries](#binaries)
  - [Docker](#docker)
  - [Source](#source)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

### Binaries

Binaries for the latest version of `ethereal` can be obtained from [the releases page](https://github.com/wealdtech/ethereal/releases).

### Docker

You can obtain the latest version of `ethereal` using docker with:

```
docker pull wealdtech/ethereal
```

### Source
`ethereal` is a standard Go program which can be installed with:

```sh
go install github.com/wealdtech/ethereal/v2@latest
```

Note that `ethereal` requires at least version 1.13 of go to operate.  The version of go can be found with `go version`.

The docker image can be built locally with:

```sh
docker build -t ethereal .
```

You can run `ethereal` using docker after that. Example:

```sh
docker run -it ethereal --help
```

Note that that many `ethereal` commands connect to an Ethereum node to obtain information or send transactions.  If the Ethereum 1 node is running directly on the server this requires the `--network=host` command, for example:

```sh
docker run --network=host etheral chain status
```

Alternatively, if the Ethereum node is running in a separate docker container a shared network can be created with `docker network create eth2` and accessed by adding `--network=eth` added to both the beacon node and `ethereal` containers.

## Usage

Ethereal contains a large number of features that are useful for day-to-day interactions with the Ethereum blockchain.

### Access to local wallets

Ethereal works with Geth, MIST and Parity wallets in the standard locations.  A simple way to check the addresses that can be seen by Ethereal is to run `ethereal account list` which will list all accounts that Ethereal can see.  If you expect an address to show up and it doesn't then please raise an issue with the relevant details.

If you use Parity and want to import a private key or a wallet from another system please see https://github.com/paritytech/parity/wiki/Backing-up-&-Restoring#restoring-options

If you use Geth and want to import a private key or a wallet from another system please see https://github.com/ethereum/go-ethereum/wiki/Managing-your-accounts

When accessing local wallets a `--passphrase` option is required to unlock the account.  Note that this is not shown in the examples

Alternatively you can use a private key directly with the `--privatekey` option, although be aware that this can leave your private key in command history.

### Access to Ethereum networks

Ethereal supports all main Ethereum networks  It auto-detects the network by querying the connected node for the network ID.  The connection should be geth-compatible, so either geth itself or parity with the `--geth` flag to enable geth compatibility mode.  The connection could be a local node or a network service such as Infura.

Ethereal contains default connections via Infura to most major networks that can be defined by the `--network` argument.  Supported neworks are mainnet, goerli, sepolia and holesky.  Alternatively a connection to a custom node can be created using the `--connection` argument.  For example a local IPC node might use `--connection=/home/ethereum/.ethereum/geth.ipc` or `--connection=http://localhost:8545/`

**The Infura key for Ethereal is shared among all users.  If you are going to carry out a lot of queries of chain data please either use a local node or your own Infura account.**

### Configuration file

Ethereal supports a configuration file; by default in the user's home directory but changeable with the `--config` argument on the command line.  The configuration file provides values that override the defaults but themselves can be overridden with command-line arguments.

The default file name is `.ethereal.json` or `.ethereal.yml` depending on the encoding used (JSON or YAML, respectively).  An example `.ethereal.json` file is shown below:

```json
{
  "timeout": "20s",
  "verbose": true,
  "network": "ropsten",
  "passphrase": "my secret passphrase"
}
```

### Output and exit status

If set, the `--quiet` argument will suppress all output.

If set, the `--verbose` argument will output additional information related to the command.  Details of the additional information is command-specific and explained in the command help below.

If set, the `--debug` argument will output additional information about the operation of Ethereal as it carries out its work.

Commands will have an exit status of 0 on success and 1 on failure.  The specific definition of success is specified in the help for each command.  For commands that generate transactions and wait for them to be mined there is an additional exit status of 2 which means the transaction has been submitted but not mined within the requested time limit.

### Transactions

Many Ethereal commands generate Ethereum transactions.  These commands have a number of settings.

The `--priority-fee-per-gas` argument sets the tip for the transaction, for example `--priority-fee-per-gas="2 gwei"`.  If not supplied it defaults to 1.5 Gwei.

The `--max-fee-per-gas` argument sets the maximum combined fee plus priority fee for the transaction, for example `--max-fee-per-gas=100gwei`.  If not supplied it defaults to 200 Gwei.

The `--gaslimit` argument hardcodes the maximum gas for the transaction, for example `--gas=100000"`.  If not supplied the gas price will be automatically calculated.

The `--nonce` argument hardcodes the nonce for the transaction, for example `--nonce=123"`.  If not supplied the nonce will be retrieved automatically from the blockchain.

The `--passphrase` argument supplies the passphrase to unlock the submitting account, for example `--passphrase="my secret passphrase"`.

The `--privatekey` argument supplies the private key to obtain and submitting account, for example `--privatekey=0x0000000000000000000000000000000000000000000000000000000000000001`.

Note that information such as the passphrase and private key might be stored in your command line history.  If this is an issue the values can be provided in the Ethereal configuration file as described above.

By default Ethereal will return once the transaction has been submitted.  The `--wait` argument makes the command wait for the transaction to be mined as well.  If waiting should be limited this can be specified with the `--limit` argument, for example `--wait --limit=60s`.

### Logging

Any time Ethereal broadcasts a transaction it logs the details in a file.  By default the file is `ethereal.log` in the user's home directory, with each line being a JSON object with the relevant fields.  The log file location can be changed with the `--log` argument.

### ENS

Ethereal fully supports ENS.  Wherever an address is seen in the examples below an ENS name can be used instead.

Ethereal will always return addresses as ENS names if ENS reverse resolution is configured.

### `account` commands

Account commands focus on information about local accounts, generally those used by Geth and Parity but also those from hardware devices.

#### `checksum`

`ethereal account checksum` generates or verifies the [EIP-55](https://eips.ethereum.org/EIPS/eip-55) checksum for a provided account address.  With the `--check` flag it checks if the supplied address is correctly checksummed, otherwise it generates a correctly checksummed version of the supplied address.  For example:

```sh
$ ethereal account checksum --address=0x7e5f4552091a69125d5dfcb7b8c2659029395bdf --check
Checksum is incorrect
$ ethereal account checksum --address=0x7e5f4552091a69125d5dfcb7b8c2659029395bdf
0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
$ ethereal account checksum --address=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf --check
Checksum is correct
```

#### `keys`

`ethereal account keys` shows the private key, public key and Ethereum address for a given account or private key.  For example:

```sh
$ ethereal account keys --privatekey=0x0000000000000000000000000000000000000000000000000000000000000001
Private key:            0x00000000000000000000000000000001
Public key:             0x0479be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8
Ethereum address:       0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
```

#### `list`

`ethereal account list` shows the Ethereum addresses of known accounts on the local computer.  For example:

```sh
$ ethereal account list
0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF
0x6813Eb9362372EEF6200f3b1dbC3f819671cBA69
0x1efF47bc3a10a45D4B230B5d10E37751FE6AA718
0x003F53E95e293D08dc34C69ABcAbF5b577E50Cf5
```

With the `--verbose` flag this will provide the location of the keystore, current Ether funds and next nonce.  For example:

```sh
$ ethereal account list --verbose
Location:       keystore:///home/ethereum/.ethereum/keystore/UTC--2019-03-12T10-12-47.585144239Z--7e5f4552091a69125d5dfcb7b8c2659029395bdf
Address:        0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
Balance:        0
Next nonce:     243
...
```

#### `nonce`

`ethereal account nonce` shows the next nonce of an Ethereum address.  For example:

```sh
$ ethereal account nonce --address=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
243
```

### `beacon` commands

Beacon commands focus on interactions with the Ethereum 2 beacon deposit contract.

### `deposit`

`ethereal beacon deposit` creates and sends an Ethereum 2 beacon deposit contract transaction.  For example:

```sh
$ ethereal beacon deposit --data=deposit.json --eth2network=mainnet --from=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
```

Note that `ethereal` obtains the information about the amount of Ether to send with deposits from the supplied deposit data.

`ethereal beacon deposit` has a number of options to control deposits.  It carries out as many checks as possible given the information to ensure the deposit is valid, correct and unique, and as such in non-standard deposit situations these options may be required to ensure the deposit is processed.

### `block` commands

Block commands focus on information about specific blocks.

#### `info`

`ethereal block info` provides information about a block.  For example:

```sh
$ ethereal block info --block=5188504
Number:                 5188504
Hash:                   0x01262b8549472c95714993135f9aa1cb09685bd33076541522e3db0481f63fe7
Block time:             1552386092 (2019-03-12 10:21:32 +0000 GMT)
Gas limit:              8000000
Gas used:               7983831 (99.80%)
Uncles:                 0
Transactions:           66
```

With the `--verbose` flag this will provide information about the miner.  For example:

```sh
$ ethereal block info --block=5188504 --verbose
Number:                 5188504
Hash:                   0x01262b8549472c95714993135f9aa1cb09685bd33076541522e3db0481f63fe7
Block time:             1552386092 (2019-03-12 10:21:32 +0000 GMT)
Mined by:               0x6212Dd88f890FefE0Af24D1404d96aDF488e4E3B
Extra:                  ؃geth�go1.10.4�linux
Difficulty:             2232415661
Gas limit:              8000000
Gas used:               7983831 (99.80%)
Transactions:           66
```

Note that `--block=latest` will provide information on the latest mined block.

#### `overview`

`ethereal block overview` provides high-level statistics about the last few mined blocks.  For example:

```sh
$ ethereal block overview
Block    Gas used/Gas limit     Block time              Gap     Coinbase
5188514   7882080/  8000000     19/03/12 10:24:00               0xCd626bc764E1d553e0D75a42f5c4156B91a63F23
5188513   7994026/  8000000     19/03/12 10:23:58       2s      0xCd626bc764E1d553e0D75a42f5c4156B91a63F23
5188512   7964126/  8000029     19/03/12 10:23:48       10s     0x635B4764D1939DfAcD3a8014726159abC277BecC
5188511   7981224/  8000000     19/03/12 10:23:43       5s      0x6212Dd88f890FefE0Af24D1404d96aDF488e4E3B
5188510   7958922/  8000000     19/03/12 10:23:35       8s      0x6212Dd88f890FefE0Af24D1404d96aDF488e4E3B
```

With the `--verbose` flag this will provide column headers.  For example:

```sh
$ ethereal block overview --verbose
Block    Gas used/Gas limit     Block time              Gap     Coinbase
5188514   7882080/  8000000     19/03/12 10:24:00               0xCd626bc764E1d553e0D75a42f5c4156B91a63F23
5188513   7994026/  8000000     19/03/12 10:23:58       2s      0xCd626bc764E1d553e0D75a42f5c4156B91a63F23
5188512   7964126/  8000029     19/03/12 10:23:48       10s     0x635B4764D1939DfAcD3a8014726159abC277BecC
5188511   7981224/  8000000     19/03/12 10:23:43       5s      0x6212Dd88f890FefE0Af24D1404d96aDF488e4E3B
5188510   7958922/  8000000     19/03/12 10:23:35       8s      0x6212Dd88f890FefE0Af24D1404d96aDF488e4E3B
```

The number of blocks displayed in the overview can be altered using the `--blocks` parameter.

### `contract` commands

Contract commands focus on deploying and interacting with Ethereum smart contracts.

The examples of the commands below use the following contract at `SampleContract.sol`:

```solidity
pragma solidity ^0.5.0;
  
contract SampleContract {
    uint256 private value;

    constructor(uint256 _value) public {
        value = _value;
    }

    function getValue() public view returns (uint256) {
        return value;
    }

    function setValue(uint256 _value) public {
        value = _value;
    }
}
```

which is compiled using the command line:

```sh
$ solc --optimize --combined-json=bin,abi SampleContract.sol >SampleContract.json
```

Note that best results the names of the files should be the same as the name of the contract (ignoring the suffix), as per the example above.

#### `call`

`ethereal contract call` calls a contract function locally on the connected node.  For example:

```sh
$ ethereal contract call --contract=0x3c24F71e826D3762f5145f6a27d41545A7dfc8cF --json=SampleContract.json --call='getValue()' --from=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
5
```

#### `deploy`

`ethereal contract deploy` deploys a contract to the Ethereum blockchain.

The binary to deploy can be supplied in two different ways.  The simplest is to compile the contract using the `--combined-json=abi,json` option of `solc` to provide a JSON file containing both the binary data and the contract's ABI and deploy using that.  For example:

```sh
$ ethereal contract deploy --json=SampleContract.json --constructor='constructor(5)' --from=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
```

Alternatively the binary data and constructor arguments can be supplied directly on the command-line.  For example:

```sh
$ BIN=`solc --optimize --bin SampleContract.sol | egrep -A 2 SampleContract.sol:SampleContract | tail -1`
$ CONSTRUCTORARGS=`0000000000000000000000000000000000000000000000000000000000000005`
$ ethereal contract deploy --data="${BIN}${CONSTRUCTORARGS}" --from=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
```

#### `send`

`ethereal contract send` sends a contract transaction to the Ethereum blockchain.  For example:

```sh
$ ethereal contract send --contract=0x3c24F71e826D3762f5145f6a27d41545A7dfc8cF --json=SampleContract.json --call='setValue(6)' --from=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
```

#### `storage`

`ethereal contract storage` accesses contract storage directly.  Key values depend on the value stored; for more details see [this article](https://medium.com/aigang-network/how-to-read-ethereum-contract-storage-44252c8af925).

```sh
$ ethereal contract storage --contract=0x3c24F71e826D3762f5145f6a27d41545A7dfc8cF --key=0x00
0x0000000000000000000000000000000000000000000000000000000000000006
```

### `dns` commands

DNS commands focus on interacting with the [EthDNS](https://www.wealdtech.com/articles/ethdns-an-ethereum-backend-for-the-domain-name-system/) system to allow DNS records to be stored on Ethereum.

Getting and setting DNS records works on the basis of a DNS resource record set.  A resource record set is defined by the tuple (domain,name,resource record type) for example (ehdns.xyz,www.ethdns.xyz,A) would return all 'A' (address) records help for www.ethdns.xyz in the domain ethdns.xyz.

#### `clear`

`ethereal dns clear` clears all resource records for a DNS zone.

#### `get`

`ethereal dns get` obtains a single resource record set for the (domain,name,resource record type) tuple.  For example:

```sh
$ ethereal dns get --domain=ethdns.xyz --name=www --resource=CNAME
www.ethdns.xyz. 21600   IN      CNAME   ethdns.xyz.
```

Resource record sets on the root domain can be fetched by omitting the `name` argument. For example:

```sh
$ ethereal dns get --domain=ethdns.xyz --resource=NS
ethdns.xyz.     43200   IN      NS      ns1.ethdns.xyz.
ethdns.xyz.     43200   IN      NS      ns2.ethdns.xyz.
```

#### `set`

`ethereal dns set` sets a single resource record set for the (domain,name,resource record type) tuple.  For example:

```sh
$ ethereal dns set --domain=ethdns.xyz --name=www --resource=CNAME --record="ethdns.xyz."
```

Resource record sets with multiple values can be supplied by separating them with "&&".  For example:

```sh
$ ethereal dns set --domain=ethdns.xyz --resource=NS --record="ns1.ethdns.xyz&&ns2.ethdns.xyz"
```

### `ens` commands

ENS commands focus on interacting with the [Ethereum Name Service](https://ens.domains/) contracts that address resources using human-readable names.

#### `address clear`

`ethereal ens address clear` removes an address associated with an ENS domain.  For example:

```sh
$ ethereal ens address clear --domain=mydomain.eth
```

#### `address get`

`ethereal ens address get` gets the address associated with an ENS domain.  For example:

```sh
$ ethereal ens address get --domain=mydomain.eth
0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
```

#### `address set`

`ethereal ens address set` sets the address associated with an ENS domain.  For example:

```sh
$ ethereal ens address set --domain=mydomain.eth --address=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
```

#### `contenthash clear`

`ethereal ens contenthash clear` clears the contenthash associated with an ENS domain.  For example:

```sh
$ ethereal ens contenthash clear --domain=mydomain.eth
```

#### `contenthash get`

`ethereal ens contenthash get` gets the contenthash associated with an ENS domain.  For example:

```sh
$ ethereal ens contenthash get --domain=mydomain.eth
/swarm/d1de9994b4d039f6548d191eb26786769f580809256b4685ef316805265ea162
```

#### `contenthash set`

`ethereal ens contenthash set` sets the contenthash associated with an ENS domain.  For example:

```sh
$ ethereal ens contenthash set --domain=mydomain.eth --content=/swarm/d1de9994b4d039f6548d191eb26786769f580809256b4685ef316805265ea162
```

Valid content hash codecs are "ipfs" and "swarm".

#### `controller get`

`ethereal ens controller get` gets the controller of the domain.  For example:

```sh
$ ethereal ens controller get --domain=mydomain.eth
0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF
```

#### `controller set`

`ethereal ens controller set` sets the controller of the domain.  For example:

```sh
$ ethereal ens controller set --domain=mydomain.eth --owner=0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF
```

#### `domain clear`

`ethereal ens domain clear` clears the ENS reverse resolution domain of an address.  For example:

```sh
$ ethereal ens domain clear --address=0x6813Eb9362372EEF6200f3b1dbC3f819671cBA69
```

#### `domain get`

`ethereal ens domain get` gets the ENS reverse resolution domain of an address.  For example:

```sh
$ ethereal ens domain get --address=0x6813Eb9362372EEF6200f3b1dbC3f819671cBA69
mydomain.eth
```

#### `domain set`

`ethereal ens domain set` sets the ENS reverse resolution domain of an address.  For example:

```sh
$ ethereal ens domain set --address=0x6813Eb9362372EEF6200f3b1dbC3f819671cBA69 --domain=mydomain.eth
```

#### `expiry`

`etheral ens expiry` obtains the date at which a domain expires.  For example:

```sh
$ ethereal ens expire --domain=mydomain.eth
```

#### `info`

`ethereal ens info` obtains various information about a domain.  For example:

```sh
$ ethereal ens info --domain=mydomain.eth
Registrant is mydomain.eth (0x388Ea662EF2c223eC0B047D41Bf3c0f362142ad5)
Registration expires at 2020-03-30 19:04:48 +0100 BST
Controller is mydomain.eth (0x388Ea662EF2c223eC0B047D41Bf3c0f362142ad5)
Resolver is 0x4C641FB9BAd9b60EF180c31F56051cE826d21A9A
Domain resolves to 0xe8E98228Ca36591952Efdf6F645C5B229E6Cf688
Address resolves to mydomain.eth
```

With the `--verbose` flag this will provide more information about the domain.  For example:

```sh
$ ethereal ens info --domain=mydomain.eth --verbose
Normalised domain is mydomain.eth
Top-level domain is eth
Domain level is 1
Name hash is 0xf6180603ce45d5470887aff0a135e31c00b5676ac13e1095d394b378df2fe532
Label is mydomain
Label hash is 0x53759ad0a707437a18aaaf314dda5a4f9bbd6dabd605c777ebaf354ac934f3c3
Domain registered on permanent registrar
Registrant is mydomain.eth (0x388Ea662EF2c223eC0B047D41Bf3c0f362142ad5)
Registration expires at 2020-03-30 19:04:48 +0100 BST
Controller is mydomain.eth (0x388Ea662EF2c223eC0B047D41Bf3c0f362142ad5)
Resolver is 0x4C641FB9BAd9b60EF180c31F56051cE826d21A9A
Domain resolves to 0xe8E98228Ca36591952Efdf6F645C5B229E6Cf688
Address resolves to mydomain.eth
```

#### `migrate`

`ethereal ens migrate` migrates a domain from the temporary registrar to the permanent registrar.  For example:

```sh
$ ethereal ens migrate --domain=mydomain.eth
```

#### `pubkey get`

`ethereal ens pubkey get` gets the public key associated with an ENS domain.  For example:

```sh
$ ethereal ens pubkey get --domain=mydomain.eth
(0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f,0x1f1e1d1c1b1a191817161514131211100f0e0d0b0c0a09080706050403020100)
```

#### `pubkey set`

`ethereal ens pubkey set` sets the public key associated with an ENS domain.  For example:

```sh
$ ethereal ens pubkey set --domain=mydomain.eth --key='(0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f,0x1f1e1d1c1b1a191817161514131211100f0e0d0b0c0a09080706050403020100)'
```


#### `register`

`ethereal ens register` registers a new ENS domain.  For example:

```sh
$ ethereal ens register --domain=mydomain.eth
```

Registration is a two-stage process.  The first stage sends a transaction committing to claim the domain, and the second stage sends a transaction revealing the commitment and obtaining the domain.  To avoid frontrunning there needs to be a delay between these two transactions of at least 10 minutes, and by default the command will send the first transaction, wait for the required time period, then send the second transaction.

`--period` is the amount of time for which the registration will be rented; use `ethereal ens rent` to find out how much it will cost to rent the domain.

#### `release`

`ethereal ens release` releases a domain, returning the name to the available pool.  If the domain is registered with the temporary registrary then any funds locked in the registration deed will be returned.  For example:

```sh
$ ethereal ens release --domain=mydomain.eth
```

#### `resolver clear`

`ethereal ens resolver clear` clears the resolver contract for the domain.  For example:

```sh
$ ethereal ens resolver clear --domain=mydomain.eth
```

#### `resolver get`

`ethereal ens resolver get` gets the address of the resolver contract for the domain.  For example:

```sh
$ ethereal ens resolver get --domain=mydomain.eth
0x5FfC014343cd971B7eb70732021E26C35B744cc4
```

#### `resolver set`

`ethereal ens resolver set` sets the resolver contract for the domain.  If the standard public resolver (found at `resolver.eth`) is required then just the domain is required to set it.  For example:

```sh
$ ethereal ens resolver set --domain=mydomain.eth
```

If a non-standard resolver is required it can be supplied with the `--resolver` argument.  For example:

```sh
$ ethereal ens resolver set --domain=mydomain.eth --resolver=0x4d9b7D10e3a42E81659A90fDbaB51Bf19DD9bba7
```

#### `subdomain create`

`ethereal ens subdomain create` creates a subdomain of an existing ENS domain.  For example:

```sh
$ ethereal ens subdomain create --domain=mydomain.eth --subdomain=mysub
```

The subdomain will be owned by the domain owner.

#### `text clear`

`ethereal ens text clear` clears the text for a given key for the domain.  For example:

```sh
$ ethereal ens text clear --domain=mydomain.eth --key="My info"
```

#### `text get`

`ethereal ens text get` gets the text for a given key for the domain.  For example:

```sh
$ ethereal ens text get --domain=mydomain.eth --key="My info"
Information goes here
```

#### `text set`

`ethereal ens text set` sets the text for a given key for the domain.  For example:

```sh
$ ethereal ens text set --domain=mydomain.eth --key="My info" --text="Information goes here"
```

#### `transfer`

`ethereal ens transfer` transfers registration of a name to another address.  For example:

```sh
$ ethereal ens transfer --domain=mydomain.eth --newregistrant=0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF
```

### `ether` commands

Ether commands focus on information about and movement of Ether.

#### `balance`

`ethereal ether balance` provides the Ether balance of an address.  For example:

```sh
$ ethereal ether balance --address=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
5189.916425903288395771 Ether
```

If required the balance can be supplied in Wei with the `--wei` option.  For example:

```sh
$ ethereal ether balance --address=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf --wei
5189916425903288395771
```

#### `sweep`

`ethereal ether sweep` sweeps all Ether from one address to another, leaving 0 behind.  For example:

```sh
$ ethereal ether sweep --from=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf --to=0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF
```

#### `transfer`

`ethereal ether transfer` transfers a set amount of Ether from one address to another.  For example:

```sh
$ ethereal ether transfer --from=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf --to=0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF --amount="1.2 Ether"
```

### `gas` commands

#### `price`

`ethereal gas price` calaculates a gas price from historical information that should allow a transaction to be included within a certain number of blocks.  For example:

```sh
$ ethereal gas price
5.229829545 GWei
```

The value is the average of the values of the 9th decile of transactions in each block when each block's transactions are ordered by gas price.  If the absolute lowest value is required instead the `--lowest` argument can be used.  The number of blocks over which to take the average can be supplied with the `--blocks` argument.

By default this command does not consider gas used when calculating the price.  Commonly the gas price for high gas transactions is higher due to them needing to be included in a block earlier to fit.  The `--gas` argument can supply an amount of gas, in which case the value returned will be the average of the gas price required to fit a transaction with the supplied gas in to the blocks.

### `hd` commands

### `keys`

`ethereal hd keys` shows the private key, public key and Ethereum address for a given hierarchical deterministic seed and path.  For example:

```sh
$ ethereal hd keys --seed="yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow" --path="m/44'/60'/0'/0/0"
Private key:            0x1b48e04041e23c72cacdaa9b0775d31515fc74d6a6d3c8804172f7e7d1248529
Public key:             0x04c4755e0a7a0f7082749bf46cdae4fcddb784e11428446a01478d656f588f94c17d02f3312b43364a0c480d628483c4fb4e3e9f687ac064717d90fdc42cfb6e0e
Ethereum address:       0xA27DF20E6579aC472481F0Ea918165d24bFb713b
```

### `network` commands

#### `blocktime`

`ethereal network blocktime` calculates the average blocktime over a number of blocks.  For example:

```sh
$ ethereal network blocktime --blocks=30
11.76s
```

Instead of using `--blocks` it is possible to specify the time over which to calculate the blocktime.  For example:

```sh
$ ethereal network blocktime --time=24h
11.52s
```

With the `--verbose` flag this will provide more information about the start and end block for the calculation.  For example:

```sh
$ ethereal network blocktime --time=2h --verbose
Block 7370164 mined at 2019-03-14 23:48:53 +0000 GMT
Block 7369614 mined at 2019-03-14 21:49:28 +0000 GMT
13.02s
```
#### `gps`

`ethereal network gps` provides a gas-per-second metric for the Ethereum network over a number of blocks.  For example:

```sh
$ ethereal network gps --blocks=20
339950
```

With the `--verbose` flag this will provide more information about each block.  For example:

```sh
$ ethereal network gps --blocks=5 --verbose
Block 7370156 used 7407093 gas in 2 seconds
Block 7370155 used 6831267 gas in 9 seconds
Block 7370154 used 7787059 gas in 2 seconds
Block 7370153 used 219751 gas in 4 seconds
Block 7370152 used 1113156 gas in 6 seconds
1015579
```

#### `id`

`ethereal network id` provides the ID of the Ethereum network.  For example:

```sh
$ ethereal network id
1
```

#### `tps`

`ethereal network tps` provides a transactions-per-second metric for the Ethereum network over a number of blocks.  For example:

```sh
$ ethereal network tps --blocks=20
4.72
```
With the `--verbose` flag this will provide more information about each block.  For example:

```sh
$ ethereal network tps --blocks=5 --verbose
Block 7373054 processed 143 transactions in 13 seconds
Block 7373053 processed 149 transactions in 36 seconds
Block 7373052 processed 172 transactions in 7 seconds
Block 7373051 processed 92 transactions in 6 seconds
Block 7373050 processed 50 transactions in 5 seconds
9.04
```

#### `usage`

`ethereal network usage` provides a % usage metric for the Ethereum network over a number of blocks in terms of `gas used/gas limit`.  For example:

```sh
$ ethereal network usage --blocks=20
97.37%
```
With the `--verbose` flag this will provide more information about each block.  For example:

```sh
$ ethereal network usage --blocks=5 --verbose
Block 7495042 used 70.37% of gas limit (5629676/8000000)
Block 7495041 used 99.82% of gas limit (7985401/8000000)
Block 7495040 used 99.81% of gas limit (7992790/8007811)
Block 7495039 used 99.92% of gas limit (7993641/8000000)
Block 7495038 used 99.76% of gas limit (7980625/8000029)
93.94%
```

### `node` commands

Node commands focus on the state of the Ethereum nodes as specified in the connection.

#### `sync`

`ethereal node sync` obtains the synchronisation state of the node as defined by the `connection` option.  For example:

```sh
$ ethereal node sync --connection=/home/ethereum/.ethereum/goerli/geth.ipc
Node is at block 1157120, syncing to block 1165105
```


### `registry` commands

Ether commands focus on use of the [ERC-1820](https://eips.ethereum.org/EIPS/eip-1820) registry.

#### `implementer get`

`ethereal registry implementer get` gets the contract that implements a specified interface for a specified address.  For example:

```sh
$ ethereal registry implementer get --interface="ERC777Token" --address=0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF
0x3c24F71e826D3762f5145f6a27d41545A7dfc8cF
```

#### `implementer set`

`ethereal registry implementer set` sets the contract that implements a specified interface for a specified address.  For example:

```sh
$ ethereal registry implementer set --interface="ERC777TokensSender" --address=0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF --implementer=0x3c24F71e826D3762f5145f6a27d41545A7dfc8cF
```

#### `implements`

`ethereal registry implements` checks if a contract implements a specified interface.  For example:

```sh
$ ethereal registry implements --interface="ERC777TokensSender" --address=0x62284ed69b907af90ecba2feef0bf12a99563563
Yes
```

#### `manager get`

`ethereal registry manager get` gets the manager for a specified address.  For example:

```sh
$ ethereal registry manager get --address=0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF
0x6813Eb9362372EEF6200f3b1dbC3f819671cBA69
```

Note that if there is no manager set for the address then this will return the provided address.  For example:

```sh
$ ethereal registry manager get --address=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
```

### `manager set`

`ethereal registry manager set` sets the manager for a specified address.  For example:

```sh
$ ethereal registry manager set --address=0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF --manager=0x6813Eb9362372EEF6200f3b1dbC3f819671cBA69
```

### `signature` commands

Signature commands focus on generation and verification of signatures within Ethereum.

### `signature sign`

`ethereal signature sign` signs provided data.  For example:

```sh
$ ethereal signature sign --data="false,2,0x5FfC014343cd971B7eb70732021E26C35B744cc4" --types="bool,uint256,address" --signer=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
08140077a94642919041503caf5cc1795b23ecf256578655de186858540a45ba44fddebfb97ba6f74d12611263a97174f5ac1ee9db30a79fe16c9a2346ef23b301
```

There are two types of information that can be signed: text and data.  A text string is a simple value for data, for example:

```sh
$ ethereal signature sign --data="Hello, world" --signer=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
fdb006b0359c64152f36022662b3ecd2c315e88c937444f337dabf18208cc111063b100ada7dc86647a8337d50c819cac7e04f90f1b2ea509ccd3a0ae82e7de700
```

Data is a set of comma-separated values with types supplied in the `--types` argument.  In this situation the data is turned in to an [ABI-encoded](https://solidity.readthedocs.io/en/develop/abi-spec.html) value; by default the data is encoded in full but can be encoded packed with the `--packed` argument.

By default the data is hashed prior to being signed; this can be overridden by supplying the `--nohash` argument.  For example:

```sh
$ ethereal signature sign --data="Hello, world" --nohash --signer=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
16f7a3ddacbffa12a1c416c72b4d3aed54fe605490c48bbca1cdf6ff2b3c3b122102a40374587a963968fae5cf75dda47f97f8f2e5992144edd59b1e7827821500
```

After hashing but before being signed the data has the standard Ethereum header added to it.  This is the data prepended with the standard Ethereum signing message of "\\x19Ethereum Signed Message:\n" followed by the number of bytes in the data and finally the data itself, for example in the prior example this would be "\\x19Ethereum Signed Message:\n12Hello, world".

### `signature signer`

`ethereal signature signer` obtains the address of the signer given a signature and the related data.  For example:

```sh
$ ethereal signature signer --data="Hello, world" --nohash --signature=16f7a3ddacbffa12a1c416c72b4d3aed54fe605490c48bbca1cdf6ff2b3c3b122102a40374587a963968fae5cf75dda47f97f8f2e5992144edd59b1e7827821500
0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
```

The same rules apply to `ethereal signature signer` as those in `ethereal signature sign` above.

### `signature verify`

`ethereal signature verify` verifies the address of the signer given an address, signature and the related data.  For example:

```sh
$ ethereal signature verify --data="false,2,0x5FfC014343cd971B7eb70732021E26C35B744cc4" --types="bool,uint256,address" --signature=08140077a94642919041503caf5cc1795b23ecf256578655de186858540a45ba44fddebfb97ba6f74d12611263a97174f5ac1ee9db30a79fe16c9a2346ef23b301 --signer=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf
```

The same rules apply to `ethereal signature verify` as those in `ethereal signature sign` above.

### `token` commands

Token commands focus on information and management of ERC-20 and ERC-777 tokens.

### `transaction` commands

Transaction commands focus on information and management of Ethereum transactions.

#### `cancel`

`ethereal transaction cancel` cancels a pending transaction.  For example:

```sh
$ ethereal transaction cancel --transaction="0x19f5c8369d49bf96e82941e938d3978f4ce9bf20e09598cf16e7c55018006f4e"
0x4dada7ddd3841d9e754fa1caa4232155d1a6d976fef610da5c3c0d00025bd0c1
```

Note that in reality Ethereum has no notion of cancelling transactions so instead the transaction is replaced with a new transaction that does nothing.  For this command to succeed transaction's maximum base fee and priority fee must both be increased by 10% over that of the existing transaction; this will happen automatically.

#### `info`

`ethereal transaction info` provides information about an Ethereum transaction.  For example:

```sh
$ ethereal transaction info --transaction=0x581560df6b07612293996772a40966e8b85f70af2d53eee624513324fad8a99a
Type:                   Mined transaction
Result:                 Succeeded
Block:                  7380609
From:                   0x2B5634C42055806a59e9107ED44D43c426E58258
To:                     0xf3db7560E820834658B590C96234c333Cd3D5E5e
Gas used:               37081
Gas price:              15.176 GWei
Value:                  0
Data:                   transfer(0x7755B69903BcbCc419260dBb65772412E0C4ad2b,3903811515500000000000)
```

With the `--verbose` flag this will provide more information about the transaction.  For example:

```sh
$ ethereal transaction info --transaction=0x581560df6b07612293996772a40966e8b85f70af2d53eee624513324fad8a99a --verbose
Type:                   Mined transaction
Result:                 Succeeded
Block:                  7380609
From:                   0x2B5634C42055806a59e9107ED44D43c426E58258
To:                     0xf3db7560E820834658B590C96234c333Cd3D5E5e
Nonce:                  1382943
Gas limit:              76351
Gas used:               37081
Gas price:              15.176 GWei
Value:                  0
Data:                   transfer(0x7755B69903BcbCc419260dBb65772412E0C4ad2b,3903811515500000000000)
Logs:
        0:
                From:   0xf3db7560E820834658B590C96234c333Cd3D5E5e
                Event:  Transfer(0x2B5634C42055806a59e9107ED44D43c426E58258,0x7755B69903BcbCc419260dBb65772412E0C4ad2b,3903811515500000000000)
```

#### `send`

`ethereal transaction send` sends a transaction.  For example:

```sh
$ ethereal transaction send --from=0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf --to=0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF  --amount="1 Ether" --data=0x010203
```

#### `up`

`ethereal transaction up` increases the gas price of an existing pending transaction.  For example:

```sh
$ ethereal transaction up --transaction=0x581560df6b07612293996772a40966e8b85f70af2d53eee624513324fad8a99a
```

For this command to succeed transaction's maximum base fee and priority fee must both be increased by 10% over that of the existing transaction; this will happen automatically.

#### `wait`

`ethereal transaction wait` waits for a pending transaction to be mined.  For example:

```sh
$ ethereal transaction wait --transaction=0x581560df6b07612293996772a40966e8b85f70af2d53eee624513324fad8a99a
```

By default this waits forever; if a timeout is required it can be supplied with the `--limit` argument.

### `version`

`ethereal version` provides the current version of Ethereal.  For example:

```sh
$ ethereal version
2.0.889
```

## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/wealdtech/ethereal/issues).

## License

[Apache-2.0](LICENSE) © 2017-2019 Weald Technology Trading Ltd

