# ioeX.Client

## Summary

ioeX leveraged Elastos functions to create its owned features and also business cases.

This is the client program of the IOEX node, which is a command line tool to control node and see node info etc.
Also, this project including a light implementation of IOEX wallet, which can create your IOEX account, do receive, create, sign or send transactions.
You can run a node locally and set the miner address to your wallet account, then run the node to get your own IOEXs and do whatever you like.

## Wiki
For more details, please check on this [Wiki](https://github.com/ioeXNetwork/ioeX.Client/wiki) page.

## Build on Mac

### Check OS version

Make sure the OSX version is 16.7+

```shell
$ uname -srm
Darwin 16.7.0 x86_64
```

### Install Go distribution 1.9

Use Homebrew to install Golang 1.9.
```shell
$ brew install go@1.9
```
> If you install older version, such as v1.8, you may get missing math/bits package error when build.

### Setup basic workspace
In this instruction we use ~/dev/src as our working directory. If you clone the source code to a different directory, please make sure you change other environment variables accordingly (not recommended).

```shell
$ mkdir ~/dev/bin
$ mkdir ~/dev/src
```

### Set correct environment variables.

```shell
export GOROOT=/usr/local/opt/go@1.9/libexec
export GOPATH=$HOME/dev
export GOBIN=$GOPATH/bin
export PATH=$GOROOT/bin:$PATH
export PATH=$GOBIN:$PATH
```

### Install Glide

Glide is a package manager for Golang. We use Glide to install dependent packages.

```shell
$ brew install --ignore-dependencies glide
```

### Check Go version and glide version
Check the golang and glider version. Make sure they are the following version number or above.
```shell
$ go version
go version go1.9.2 darwin/amd64

$ glide --version
glide version 0.13.1
```
If you cannot see the version number, there must be something wrong when install.

### Clone source code to $GOPATH/src/github.com/ioeXNetwork/ folder
Make sure you are in the folder of `$GOPATH/src/github.com/ioeXNetwork/`
```shell
$ git clone https://github.com/ioeXNetwork/ioeX.Client.git
```

If clone works successfully, you should see folder structure like $GOPATH/src/github.com/ioeXNetwork/ioeX.Client/Makefile

### Glide install

Run `glide update && glide install` to download project dependencies.

### Install sqlite database
This will make the `make` progress far more faster.
```shell
go install github.com/ioeXNetwork/ioeX.Client/vendor/github.com/mattn/go-sqlite3
```

### Make

Run `make` to build the executable files `ioex-cli`

## Run on Mac
## Build on Ubuntu

### Check OS version
Make sure your ubuntu version is 16.04+
```shell
$ cat /etc/issue
Ubuntu 16.04.3 LTS \n \l
```

### Install basic tools
```shell
$ sudo apt-get install -y git
```

### Install Go distribution 1.9
```shell
$ sudo apt-get install -y software-properties-common
$ sudo add-apt-repository -y ppa:gophers/archive
$ sudo apt update
$ sudo apt-get install -y golang-1.9-go
```
> If you install older version, such as v1.8, you may get missing math/bits package error when build.

### Setup basic workspace
In this instruction we use ~/dev/src as our working directory. If you clone the source code to a different directory, please make sure you change other environment variables accordingly (not recommended).

```shell
$ mkdir ~/dev/bin
$ mkdir ~/dev/src
```

### Set correct environment variables.

```shell
export GOROOT=/usr/lib/go-1.9
export GOPATH=$HOME/dev
export GOBIN=$GOPATH/bin
export PATH=$GOROOT/bin:$PATH
export PATH=$GOBIN:$PATH
```

### Install Glide

Glide is a package manager for Golang. We use Glide to install dependent packages.

```shell
$ cd ~/dev
$ curl https://glide.sh/get | sh
```

### Check Go version and glide version
Check the golang and glider version. Make sure they are the following version number or above.
```shell
$ go version
go version go1.9.2 linux/amd64

$ glide --version
glide version v0.13.1
```
If you cannot see the version number, there must be something wrong when install.

### Clone source code to $GOPATH/src/github.com/ioeXNetwork/ folder
Make sure you are in the folder of `$GOPATH/src/github.com/ioeXNetwork/`
```shell
$ git clone https://github.com/ioeXNetwork/ioeX.Client.git
```

If clone works successfully, you should see folder structure like $GOPATH/src/github.com/ioeXNetwork/ioeX.Client/Makefile

### Glide install

Run `glide update && glide install` to install depandencies.

### Install sqlite database
This will make the `make` progress far more fester.
```shell
go install github.com/ioeXNetwork/ioeX.Client/vendor/github.com/mattn/go-sqlite3
```

### Make

Run `make` to build the executable files `ioex-cli`


## Run on Mac/Ubuntu

### Set up configuration file
A file named `cli-config.json` should be placed in the same folder with `ioex-cli` with the parameters as below.
```
{
    "Host": "127.0.0.1:20336"
}
```
> `Host` is the IP and Port witch this client is communicate with. Usually `ioex-cli` is working with `node` together on the same machine，
so mostly IP is set to `localhost` and `Port` value is according to the `HttpJsonPort` value set in the node `config.json` file.

### See node info
As the node is running, you can ge information from it by using `info` commands.
```shell
$ ./ioex-cli info
NAME:
   ioex-cli info - show node information

USAGE:
   ioex-cli info [command options] [args]

DESCRIPTION:
   With ioex-cli info, you could look up node status, query blocks, transactions, etc.

OPTIONS:
   --connections                         see how many peers are connected with current node
   --neighbor, --nbr                     show neighbor nodes information
   --state                               show current node status
   --currentheight, --height             show blockchain height on current not
   --getblockhash value, --blockh value  query a block's hash with it's height (default: -1)
   --getblock value, --block value       query a block with height or it's hash
   --gettransaction value, --tx value    query a transaction with it's hash
   --showtxpool, --txpool                show the transactions in node's transaction pool
```

### Mine
With `mine` command, you can toggle the node CPU mining, and when testing or try something else, waiting for the CPU mining is a waste of time, manual mine command can solve this problem.
To use manual mine, `ActiveNet` parameter in node config file should be set to `RegNet` and `AutoMining` must be set to `false`.
```shell
$ ./ioex-cli mine
NAME:
   ioex-cli mine - toggle cpu mining or manual mine

USAGE:
   ioex-cli mine [command options] [args]

DESCRIPTION:
   With ioex-cli mine, you can toggle cpu mining, or manual mine blocks.

OPTIONS:
   --toggle value, -t value  use --toggle [start, stop] to toggle cpu mining
   --number value, -n value  user --number [number] to mine the given number of blocks
```

### Log
This is the command to control node log print level, levels are from 0~6, the lower level the more logs will be print out, 0 means print out everything.
```shell
$ ./ioex-cli log
NAME:
   ioex-cli log - set node log output level

USAGE:
   ioex-cli log [command options] [args]

DESCRIPTION:
   With ioex-cli log, you could change node log output level.
   levels are 0~6, the lower level the more logs will be print out, 0 means print out everything

OPTIONS:
   --level value, -l value  --log level (default: -1)
```

### Wallet
For test purpose, this command line client implemented a simplified wallet program. You can use it to create your IOEX account, check account balance and build, sign or send transactions.
#### Tips
> for some reason, when using multiple command options, the option with no arguments must be put at the last, for example
`$ ./ioex-cli wallet --name my_wallet.dat --password ioeXNetwork --create`, in this case, `--create` must be put at the last, or the command will no work correctly.
```shell
$ ./ioex-cli wallet
NAME:
   ioex-cli wallet - wallet operations

USAGE:
   ioex-cli wallet [command options] [args]

DESCRIPTION:
   With ioex-cli wallet, you can create an account, check account balance or build, sign and send transactions.

OPTIONS:
   --password value, -p value     arguments to pass the password value
   --name value, -n value         to specify the created keystore file name or the keystore file path to open (default: "keystore.dat")
   --import value                 create your wallet using an existed private key
   --export                       export your private key from this wallet
   --create, -c                   create wallet, this will generate a keystore file within you account information
   --account, -a                  show account address, public key and program hash
   --changepassword               change the password to access this wallet, must do not forget it
   --reset                        clear the UTXOs stored in the local database
   --addaccount value             add a standard account with a public key, or add a multi-sign account with multiple public keys
                                  use -m to specify how many signatures are needed to create a valid transaction
                                  by default M is public keys / 2 + 1, witch means greater than half
   -m value                       the M value to specify how many signatures are needed to create a valid transaction (default: 0)
   --delaccount value             delete an account from database using it's address
   --list, -l                     list accounts information, including address, public key, balance and account type.
   --transaction value, -t value  use [create, sign, send], to create, sign or send a transaction
                                  create:
                                    use --to --amount --fee [--lock], or --file --fee [--lock]
                                    to create a standard transaction, or multi output transaction
                                  sign, send:
                                    use --file or --hex to specify the transaction file path or content
   --from value                   the spend address of the transaction
   --to value                     the receive address of the transaction
   --amount value                 the transfer amount of the transaction
   --fee value                    the transfer fee of the transaction
   --lock value                   the lock time to specify when the received asset can be spent
   --hex value                    the transaction content in hex string format to be sign or send
   --file value, -f value         the file path to specify a CSV file path with [address,amount] format as multi output content,
                                  or the transaction file path with the hex string content to be sign or send
```

### Examples
Create the wallet

`$ ./ioex-cli wallet --create` or `$ ./ioex-cli wallet -c`

Create a wallet with password arguments

`$ ./ioex-cli wallet --password ioeXNetwork --create`

Show account information

`$ ./ioex-cli wallet --account` or `$ ./ioex-cli wallet -a`
```shell
Password:
--------------------------------------------------------------------------------
Address:      EXiCyZBdvguJU5upFGZwUQMJFB53TBb6km
Public Key:   0242d5fd3dd847ca7a12cf01d67ffefc21b5ea739a31b9bd9b6289227527c88199
ProgramHash:  7721066f3791c6df687300c9706236544baaad9f21
--------------------------------------------------------------------------------
```

Show account balance

`$ ./ioex-cli wallet --list` or `$ ./ioex-cli wallet -l`
```shell
201 / 201 [=========================================================] 100.00% 0s
--------------------------------------------------------------------------------
Address:      EXiCyZBdvguJU5upFGZwUQMJFB53TBb6km
ProgramHash:  7721066f3791c6df687300c9706236544baaad9f21
Balance:      2
--------------------------------------------------------------------------------
Address:      8FQZdRrN8bSJuzSJh4im2teMqZoenmeJ4u
ProgramHash:  433572eda403ac6372cc0004b2df2dc602a2890312
Balance:      0
--------------------------------------------------------------------------------
```

Create a transaction

`$ ./ioex-cli wallet -t create --from EXiCyZBdvguJU5upFGZwUQMJFB53TBb6km --to EXYPqZpQQk4muDrdXoRNJhCpoQtFBQetYg --amount 10000 --fee 0.00001`

Create a multi output transaction

`$ ./ioex-cli wallet -t create --from 8JiMvfWKDwEeFNY3KN38PBif19ZhGGF9MH --file addresses.csv --fee 0.00001`

Sign a transaction

`$ ./ioex-cli wallet -t sign --file to_be_signed.txn`

Send a transaction

`$ ./ioex-cli wallet -t send --file ready_to_send.txn`

## License
ioeXNetwork client source code files are made available under the MIT License, located in the LICENSE file.
