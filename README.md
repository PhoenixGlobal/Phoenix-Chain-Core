## Phoenix Chain Core

Golang implementation of the Phoenix Chain protocol.

Welcome to the Phoenix-Chain-Core source code repository! The Phoenix Chain protocol is based on Ethereum and Tendermint. It has a high-performance consensus mechanism,which is a hybrid of DPOS and PBFT.

## Building the source

- install golang 1.16+

```
wget https://dl.google.com/go/go1.17.7.linux-amd64.tar.gz
tar -zxf go1.17.7.linux-amd64.tar.gz -C /usr/local
vi /etc/profile
```

paste the following to /etc/profile:
```
#golang env config
export GO111MODULE=on
export GOROOT=/usr/local/go
export GOPATH=/home/gopath
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

then
```
cd /home 
mkdir gopath
source /etc/profile
```

- install libraries

for ubuntu:
```
sudo apt install libgmp-dev libssl-dev
```
for centos:
```
yum install -y gmp-devel openssl-devel
```

- build
```
make all
sudo cp -f ./build/bin/phoenixkey /usr/bin/
sudo cp -f ./build/bin/phoenixchain /usr/bin/
```

## Starting up your node to connect to the Phoenix Chain network

```
mkdir data
cd data
```
- generate keys
```
phoenixkey genkeypair | tee >(grep "PrivateKey" | awk '{print $2}' > ./nodekey) >(grep "PublicKey" | awk '{print $3}' > ./nodeid) && phoenixkey genblskeypair | tee >(grep "PrivateKey" | awk '{print $2}' > ./blskey) >(grep "PublicKey" | awk '{print $3}' > ./blspub)
```

- generate a wallet
```
phoenixchain --datadir ./ account new --name "{wallet-name}"
```

- run node
```
cd ..
nohup phoenixchain --identity "phoenix" --networkid 908 --datadir ./data --port 16789 --rpcaddr 0.0.0.0 --rpcport 6789 --rpcapi "phoenixchain,net,web3,admin,personal" --rpc --nodekey ./data/nodekey --pbft.blskey ./data/blskey > ./data/phoenix.log 2>&1 &
```

- check node
```
phoenixchain attach http://localhost:6789 --exec phoenixchain.blockNumber
```

## License
The Phoenix-Chain-Core library (i.e. all code outside of the commands directory) is licensed under the GNU Lesser General Public License v3.0, also included in our repository in the COPYING.LESSER file.

The Phoenix-Chain-Core binaries (i.e. all code inside of the commands directory) is licensed under the GNU General Public License v3.0, also included in our repository in the COPYING file.