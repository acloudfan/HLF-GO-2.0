Demonstrates Query by Range
===========================

# Setup the vendor dependencies
cd $GOPATH/src/token/qry/v1   [Change the version as needed]
./govendor.sh

This may take a few minutes

================================================
Testing Get State Range functions (token/qry/v1)
================================================
1. Launch dev environment in either mode
dev-init.sh

2. Setup the  env
. set-env.sh acme

3. Setup chaincode env. The init function takes args for setting sample data
reset-chain-env.sh
set-chain-env.sh   -n qry  -v 1.0 -p token/qry/v1 -I true
set-chain-env.sh   -c '{"Args": ["init","1","50"]}'

4. Install and instantiate
chain.sh install -p 

<< Depending on env - you may be prompted to change sequence >>

chain.sh instantiate
*OR*
chain.sh approveformyorg
chain.sh commit
chain.sh init

5. Try out the queries now
set-chain-env.sh -q '{"Args": ["GetTokenByRange", "key10", "key12"]}' 
chain.sh query

set-chain-env.sh -q '{"Args": ["GetTokenByRange", "key1", "key2"]}' 
chain.sh query

6. Get all rows
set-chain-env.sh -q '{"Args": ["GetTokenByRange", "", ""]}'
chain.sh query

============================================================
Exercise - extend invoke to support pagintion (token/qry/v1)
============================================================
Extend the token/qry/v1 such that the funcname = "GetTokenByRangeWithPagination"
will invoke the function stub.GetStateByRangeWithPagination

arg[0]="startKey"   arg[1]="endKey"  arg[3]="pageSize" i.e., number of records/page

Solution=>Solution is in solution/pagination.go 

Testing=>
1. Copy the pagination.go to token/qry/v1

2. Uncomment code in invoke function to invoke "GetTokenByRangeWithPagination"

3. Reset the environment
dev-init.sh

reset-chain-env.sh 
set-chain-env.sh   -n qry  -v 1.0 -p token/qry/v1 
set-chain-env.sh   -c '{"Args": ["init","1","50"]}'  -I true

4. Install & Instantiate the chaincode

chain.sh install -p
chain.sh instantiate

5. Get the data with startKey=10  endKey=20 pageSize=5

set-chain-env.sh -q '{"Args": ["GetTokenByRangeWithPagination", "key10", "key20","5"]}' 
chain.sh query

6. Get all of the data with pagSize=10

set-chain-env.sh -q '{"Args": ["GetTokenByRangeWithPagination", "", "","10"]}' 
chain.sh query

# All data in chunks of 5 / page
set-chain-env.sh -q '{"Args": ["GetTokenByRangeWithPagination", "", "","5"]}' 
chain.sh query


----------------------- Fabric v 1.4 ------------------

Testing Get State Range functions (token/qry/v1)
================================================
1. Launch dev environment in either mode
dev-init.sh

2. Setup the  env
. set-env.sh acme

3. Setup chaincode env. The init function takes args for setting sample data
reset-chain-env.sh
set-chain-env.sh   -n qry  -v 1.0 -p token/qry/v1 
set-chain-env.sh   -c '{"Args": ["init","1","50"]}'

4. Install and instantiate
chain.sh install
chain.sh instantiate

5. Try out the queries now
set-chain-env.sh -q '{"Args": ["GetTokenByRange", "key10", "key12"]}' 
chain.sh query

set-chain-env.sh -q '{"Args": ["GetTokenByRange", "key1", "key2"]}' 

6. Get all rows
set-chain-env.sh -q '{"Args": ["GetTokenByRange", "", ""]}'

Exercise - extend invoke to support pagintion
=============================================
Extend the token/qry/v1 such that the funcname = "GetTokenByRangeWithPagination"
will invoke the function stub.GetStateByRangeWithPagination
arg[0]="startKey"   arg[1]="endKey"  arg[3]="pageSize" i.e., number of records/page

Solution=>Solution is in solution/pagination.go 

Testing=>
1. Copy the pagination.go to token/qry/v1
2. Uncomment code in ivoke function
3. Reset the environment
dev-init.sh
set-chain-env.sh   -n qry  -v 1.0 -p token/qry/v1 
set-chain-env.sh   -c '{"Args": ["init","1","50"]}'
4. Get the data with startKey=10  endKey=20 pageSize=5
set-chain-env.sh -q '{"Args": ["GetTokenByRangeWithPagination", "key10", "key20","5"]}' 
chain.sh query
5. Get all of the data with pagSize=10
set-chain-env.sh -q '{"Args": ["GetTokenByRangeWithPagination", "", "","10"]}' 
chain.sh query

# All data in chunks of 5 / page
set-chain-env.sh -q '{"Args": ["GetTokenByRangeWithPagination", "", "","5"]}' 


Query Result Interface
======================
https://godoc.org/github.com/hyperledger/fabric/protos/ledger/queryresult

