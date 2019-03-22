![](images/wtc_logo.jpg)

# GO WTC
Waltonchain Mainnet User Manual

## The goal
To let testers experience mining, query, transfer, smart contract creation and other functions of the Waltonchain test mainnet.

## Steps

### 1. Run docker container
`# docker pull ko12/demo`  
`# docker run -it -p 10101:10101 ko12/demo bash`

### 2. Compile source code
`# source ~/.profile`  
`# cd /usr/local/src`  
`# git clone https://github.com/disy-yin/waltonchain-gwtc.git`  
`# cd waltonchain-gwtc`  
`# make gwtc`  
`# ./build/bin/gwtc version`  

### 3. Deploy
`# cd /usr/local/src/waltonchain-gwtc/gwtc_bin/`  
`# cp ../build/bin/gwtc ./bin/gwtc_v1.0.0stable`  
`# ./backend.sh`

### 4. Enter console
`# cd /usr/local/src/waltonchain-gwtc/gwtc_bin/`  
`# ./bin/gwtc_v1.0.0stable attach ./data/gwtc.ipc`

### 5. View information of the connected node
`# admin.peers`

### 6. Create account
`# personal.newAccount()`  
`# ******`  ---- Enter new account password  
`# ******`  ---- Confirm the new account password  

### 7. Mine
`# miner.start()`

### 8. Query
`# wtc.getBalance(wtc.coinbase)`

### 9. Unlock account
`# personal.unlockAccount(wtc.coinbase)`

### 10. Transfer
`# wtc.sendTransaction({from: wtc.accounts[0], to: wtc.accounts[1], value: web3.toWei(1)})`

### 11. Exit console
`# exit`

### 12. View log
`# cd /usr/local/src/waltonchain-gwtc/gwtc_bin/`  
`# tail -f gwtc.log`

### 13. Stop gwtc
`# cd /usr/local/src/waltonchain-gwtc/gwtc_bin/`  
`# ./stop.sh`
