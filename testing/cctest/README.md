
Fabric 2.0 Update
=================
MockStub moved from shim package to shimtest package
The package needs to be installed before it can be used.

go get github.com/hyperledger/fabric-chaincode-go/shimtest

# Demostrates the use of MockStub
1. To run the test

go test -v 

2. To understand the behavior add new test functions to calc_test.go
3. Add new files with test functions


# Using the test script
set-chain-env.sh  -n   cctest
set-chain-env.sh  -p   testing/cctest
./test.sh