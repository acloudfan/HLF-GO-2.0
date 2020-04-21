Demonstrates the use of history
===============================
History.go chaincode manages Cars as assets on the chain.
Each Car managed on the chain is identified by a key = VIN (Vehicle Identification Number)
Test data will be setup in the init() function

TransferOwnership  = Transaction to Transfer the ownership of the car
GetVehicleHistory  = History of all ownership transfer transactions can be retrieved

Setup chaincode vendoring
=========================
Switch to the chaincode folder. 
cd $GOPATH/src/token/history
./govendor.sh

If vendoring is not done the "instantiate" command will fail in net mode
This script may take upto 10+ min sometime :) PLEASE be patient
(ignore warning/error)

Setup the chaincode
===================
. set-env.sh acme
set-chain-env.sh -n history -v 1.0  -p token/history -c '{"Args":["init"]}'

Validate chaincode
==================
chain.sh install
chain.sh instantiate

set-chain-env.sh  -q '{"Args": ["GetVehicleByVin", "100"]}'
chain.sh query

set-chain-env.sh  -i '{"Args": ["TransferOwnership", "100","J Smith","H Koolaid","2019-01-01"]}'
chain.sh invoke

set-chain-env.sh  -i '{"Args": ["TransferOwnership", "100","H Koolaid","M Rainbow","2019-02-01"]}'
chain.sh invoke

set-chain-env.sh  -q '{"Args": ["GetVehicleHistory", "100"]}'
chain.sh query

Assets:
======
VIN,Make,Model,Year,Owner
100,toyota,corolla,2001,J Smith
200,honda,civic,199,G Roger
300,audi,a5,1999,S Ripple
400,bmw,x5,2013,M Jane
500,toyota,camry,2018,J Hoover

KeyModification
===============
https://godoc.org/github.com/hyperledger/fabric/protos/ledger/queryresult

Additional function
===================
set-chain-env.sh  -q '{"Args": ["GetVehiclesByYear", "2012"]}'
chain.sh query

Index JSON:
{
    "index": {
       "fields": [
          "year"
       ]
    },
    "name": "index-on-year",
    "ddoc": "index-on-year",
    "type": "json"
 }