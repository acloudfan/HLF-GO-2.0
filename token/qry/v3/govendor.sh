#!/bin/bash

# Initialize
echo "Initializing ..."
govendor init
govendor add +external

# Add the dependencies
echo "Adding dependencies ... please wait"
# Fabric 1.4
# govendor add github.com/hyperledger/fabric/core/chaincode/shim
# govendor add github.com/hyperledger/fabric/protos/peer
# govendor add github.com/hyperledger/fabric/protos/ledger/queryresult

# April 2020 Fabric 2.0
govendor add github.com/hyperledger/fabric-chaincode-go/shim
govendor add github.com/hyperledger/fabric-chaincode-go/shim/internal
govendor add github.com/hyperledger/fabric-protos-go/peer
govendor add github.com/hyperledger/fabric-protos-go/ledger/queryresult