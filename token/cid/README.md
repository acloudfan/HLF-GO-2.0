Client Identity Library
=======================
https://github.com/hyperledger/fabric-chaincode-go/tree/master/pkg/cid

# Vendoring
cd $GOPATH/src/token/cid/
./govendor.sh

=============
Users in acme
=============
Please check out the attributes associated with the various users.
These identities have been discused in the lecture.

#  mary
"app.accounting.role=tradeapprover:ecert","department=accounting:ecert"
#  john
"app.accounting.role=accountant:ecert","department=accounting:ecert"
# anil
"department=logistics:ecert"

======================
Set up the environment
======================
# Make sure the context is admin otherwise the instantiate will Fail
# with endrosement error
. set-env.sh acme
. set-ca-msp.sh admin

ca-dev-init.sh

====================
Set up the chaincode 
====================
set-chain-env.sh -n cid -v 1.0  -p token/cid -c '{"Args":["init"]}' -I false

# Generate the package & install
chain.sh install -p

# Instantiate the chaincode
chain.sh instantiate





==========================================
Query - Get callers certificate attributes
==========================================
# Check out the attributes of the callers
set-chain-env.sh  -q '{"Args": ["ReadAttributesOfCaller"]}'
chain.sh query

. set-ca-msp.sh  mary
chain.sh query

. set-ca-msp.sh  john
chain.sh query

. set-ca-msp.sh  anil
chain.sh query

======================================
Invoke - Assert on caller's department
======================================
set-chain-env.sh  -i '{"Args": ["AsssertOnCallersDepartment"]}'

. set-ca-msp.sh  mary
chain.sh invoke

. set-ca-msp.sh  anil
chain.sh invoke

. set-ca-msp.sh  john
chain.sh invoke


======================================================
Exercise - Enhance the cid.go to manage trade approval
======================================================

+ Copy the cid/solution/ApproveTrade.go to folder cid/
+ Uncomment the call to ApproveTrade in the invoke function

If you already have the chaincode installed then upgrade:
=========================================================
. set-env.sh acme
. set-ca-msp.sh admin
chain.sh   upgrade-auto

If you do not have the chaincode installed then install/instantiate
====================================================================
. set-env.sh acme
. set-ca-msp.sh admin

set-chain-env.sh -n cid -v 1.0  -p token/cid -I false

chain.sh install -p
chain.sh instantiate

Business Rules
==============
In our test identity setup only mary & john are from dept accounting
so they are the only ones who can approve trade.

If trade is < 100K both john/mary can approve
If trade is >= 100K only mary can approve

Test Case: Trade value < $100K
===============================
set-chain-env.sh  -i '{"Args": ["ApproveTrade", "50000"]}'

# Anil Should NOT be able to approve any trade
. set-ca-msp.sh anil
chain.sh invoke   

# John should be able to approve as trade value is < 100,000
. set-ca-msp.sh john
chain.sh invoke   

. set-ca-msp.sh mary
chain.sh invoke 

Test Case: Trade value >= $100K
===============================

set-chain-env.sh  -i '{"Args": ["ApproveTrade", "2000000"]}'

# Anil is not from accounting so this will fail
. set-ca-msp.sh anil
chain.sh invoke 

# John should not be able to approve as trade value > $100K
. set-ca-msp.sh john
chain.sh invoke 

# Only Mary can approve
. set-ca-msp.sh mary
chain.sh invoke 


============
Common Issue
============
Error on Install | Approve | Commit
"Error: proposal failed with status: 500 - failed to invoke backing implementation of 'CommitChaincodeDefinition': chaincode definition not agreed to by this org "

Solution: Set the MSP for admin & try again
=========
. set-ca-msp.sh admin