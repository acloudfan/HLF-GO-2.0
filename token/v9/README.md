Dependency management
govendor init
https://github.com/golang/go/wiki/PackageManagementTools


Test Steps #1
=============
1. Setup the env
    . set-env.sh   acme
    set-chain-env.sh  -n token -v 1.0 -p token/v9 -c '{"Args": ["init"]}'
    set-chain-env.sh -q   '{"args":["get"]}'

2. Install the chaincode
   chain.sh install -p

3. Setup the signature policy
   set-chain-env.sh -g "OR('BudgetMSP.member')"

4. As Acme Approve for the defined signature policy
   chain.sh approve
   chain.sh check

5. As Acme Commit the chaincode
   chain.sh commit
   <<Commit will be successful but CC will not be inited>>

6. Setup the event listener
   events.sh -t chaincode -n token -e SetToken -c airlinechannel 

7. Open a new terminal <<Terminal #2>> Invoke the chaincode
   source set-env.sh acme
   set-chain-env.sh -i   '{"args":["set", "UnProtectedToken","Acme"]}' 
   chain.sh invoke
   <<Error...token has not been initialized for this version, must call as init first>>
   chain.sh init
   <<Endorsement Policy Failure>>

   << Observe in Terminal #1>>

8. Since the EP requires endorsement from Budget - let's install CC on Budget

   . set-env.sh budget
   chain.sh install            [-p not needed as the package is already created]
   chain.sh approveformyorg
   chain.sh list

   # Do we need to Commit?

   
8. Now do a init and invoke
   . set-env.sh acme
   set-chain-env.sh -e budget

   chain.sh init
   
   set-chain-env.sh -i   '{"args":["set", "UnProtectedToken","Budget"]}' 
   chain.sh invoke

Exercise Solution - Update the Endorsement Policy
=================================================
Update the chaincode EP
Assumption: The chaincode token/v9 is already installed & committed
In version 2.0 you do not need to re-install the chaincode

Update the EP to be "AND('BudgetMSP.member','AcmeMSP.member')" 

1. Set the EP and the new sequence number
   set-chain-env.sh -g "AND('BudgetMSP.member','AcmeMSP.member')" 
   set-chain-env.sh -s 2

2. Acme approve with a new Sequence#
   source set-env.sh acme
   chain.sh approve
   chain.sh check

3. Budget approves 
   source set-env.sh budget
   chain.sh approve
   chain.sh check

   chain.sh commit

4. Invoke the chaincode

   set-chain-env.sh -e both

   chain.sh invoke







# Get the package for 
go get github.com/hyperledger/fabric/core/chaincode/shim/ext/statebased


Demostrates the use of Endorsement policies (v1.4)
==================================================
1. Setup the env
    . set-env.sh   acme
    set-chain-env.sh  -n token -v 1.0 -p token/v9 -c '{"Args": ["init"]}'
    set-chain-env.sh  -i   '{"args":["set", "UnProtectedToken","By Acme"]}' 

2. Install & Instantiate
    chain.sh install 
    set-chain-env.sh -P   "OR('BudgetMSP.member')"
    chain.sh instantiate

3. Setup the event listener
    events.sh -t chaincode -n token -e SetToken -c airlinechannel 

4. Query <Teminal#2>
    
    set-chain-env.sh -q   '{"args":["get"]}'
    chain.sh   query     Will work for both Acme & Budget

5. <Terminal#2> Invoke Will work for Budget only unless Acme sends Txn proposal to Budgetx
    . set-env.sh   acme
    set-chain-env.sh  -i   '{"args":["set", "UnProtectedToken","By Acme"]}' 
    chain.sh invoke

6.  Install on Budget
    . set-env.sh budget
    chain.sh install

7.  Invoke on Acme again
. set-env.sh  acme
   .   cc.env.sh

   peer chaincode invoke -o "$ORDERER_ADDRESS" -C "$CC_CHANNEL_ID" -n "$CC_NAME"  -c "$CC_INVOKE_ARGS" --peerAddresses budget-peer1.budget.com:8051

   Execute invoke with Budget context
   . set-env.sh budget
   chain.sh install

   set-chain-env.sh -i   '{"args":["set", "UnProtectedToken","Budget"]}' 
   chain.sh invoke

Exercise Upgrade EP
===================
Launch the environment with explorer
1. Setup the chaincode environment
   set-chain-env.sh -v 2.0   -P   "AND('AcmeMSP.member', 'BudgetMSP.member')"

2. Upgrade the Chaincode with new policy 
   . set-env.sh acme
   chain.sh install
   chain.sh upgrade

 3. Invoke the chain code

   . cc.env.sh

   peer chaincode invoke -o "$ORDERER_ADDRESS" -C "$CC_CHANNEL_ID" -n "$CC_NAME"  -c "$CC_INVOKE_ARGS" --peerAddresses acme-peer1.acme.com:7051  --peerAddresses budget-peer1.budget.com:8051

   Check in the Explorer

4. Kill the container for one of the peers
   docker kill acme-peer1.acme.com

5. Invoke the chain code
   peer chaincode invoke -o "$ORDERER_ADDRESS" -C "$CC_CHANNEL_ID" -n "$CC_NAME"  -c "$CC_INVOKE_ARGS"   --peerAddresses budget-peer1.budget.com:8051

   Check in the explorer


Key Level EP
============
1. Initialize the env
    dev-init.sh  -e

2. Install & Instantiate without chaincode EP
   .  set-env.sh  acme
   chain.sh install

   . set-env.sh   budget
   chain.sh install
   # Ensure chaincode EP is nil
   set-chain-env.sh   -P  ""
   chain.sh instantiate


4. Set the EP for "ProtectedToken"
   .  set-env.sh  acme
   set-chain-env.sh -i   '{"args":["setEP", "BudgetMSP.member"]}' 
   chain.sh invoke


3. Check the current EP for "ProtectedToken"
   set-chain-env.sh -q   '{"args":["getEP"]}'
   chain.sh query

6. Set the value of the "UnProtectedToken" value as "Acme"
   . set-env.sh acme
   set-chain-env.sh -i   '{"args":["set", "UnProtectedToken","Acme setting it"]}' 
   chain.sh invoke

6. Set the value of the "ProtectedToken" value as "Acme"
   . set-env.sh acme
   set-chain-env.sh -i   '{"args":["set", "ProtectedToken","Acme setting it"]}' 

   # Negative test - this will fail
   chain.sh invoke

   # Positive test - execute against the budget peer
   . cc.env.sh
   peer chaincode invoke -o "$ORDERER_ADDRESS" -C "$CC_CHANNEL_ID" -n "$CC_NAME"  -c "$CC_INVOKE_ARGS"    --peerAddresses acme-peer1.acme.com:7051 --peerAddresses budget-peer1.budget.com:8051

    Checkout the last 2 transactions

8.  Set the value of the "UnProtectedToken" as "Acme"


9.  chain.sh install
10. chain.sh invoke




Testing:

1. Setup the chaincode without EP
> Each org will use their own default peers 
> In dev setup this policy is like "OR('AcmeMSP', 'BudgetMSP')"

2. Setup the chaincode with EP "OR('BudgetMSP.member')"
> This says that BudgetMSP peer must endorse the transactions
set-chain-env.sh -P "OR('BudgetMSP.member')"

docker kill budget-peer1.budget.com
To restart budget peer =>  dev-start.sh


Endorsement Policy Testing
==========================
1. Start the environment in net mode
> dev-init.sh

2. Set the chaincode environment
> set-chain-env.sh -n token -v 1.0 -p token/v9  -c  '{"args":[]}'

3. Set the endorsement policy for the chaincode
>  set-chain-env.sh -P "OR('BudgetMSP.member')"

4. Set the environment context to acme
> . set-env.sh acme
> chain.sh install
> chain.sh instantiate
# This is equivalent to below:
peer chaincode instantiate -c  '{"args":[]}' -C airlinechannel -n token -v 1.0 -P "OR('BudgetMSP.member')" -o orderer.acme.com:7050

5. Now invoke the chaincode
> set-chain-env.sh   -i   '{"args":["set"]}' -q   '{"args":["get"]}'
> chain.sh invoke
> Checkout the explorer - you will see a transaction with "EP Failure"
  # Invoke need to be sent to the EP
> export CORE_PEER_ADDRESS=budget-peer1.budget.com:8051
> chain.sh invoke
> Checkout the explorer - you will see a transaction with "EP Failure"

6. Now install the chaincode on budget
> . set-env.sh budget
> chain.sh install

7.  Now invoke the chaincode as acme
> . set-env.sh acme
> export CORE_PEER_ADDRESS=budget-peer1.budget.com:8051
> chain.sh invoke
> Checkout the explorer - you will see a VALID transaction


Peer chaincode invoke on multiple Endorsers




. set-env.sh  budget
chain.sh install

* 5.3 Now invoke the chaincode as acme

Check the explorer - you will find a VALID txn from Acme

6. Test with EP down
6.1 To kill just use the docker kill
> docker kill budget-peer1.budget.com

Exercise
========
set-chain-env.sh -P "AND('AcmeMSP.member','BudgetMSP.member')"

chain.sh install
chain.sh instantiate
peer invoke --peerAddresses acme-peer1.acme.com:7051 --peerAddresses budget-peer1.budget.com.com:7051

export CORE_PEER_ADDRESS=budget-peer1.budget.com:8051,budget-peer1.budget.com.com:7051

peer chaincode invoke -o "$ORDERER_ADDRESS" -C "$CC_CHANNEL_ID" -n "$CC_NAME"  -c "$CC_INVOKE_ARGS" --peerAddresses acme-peer1.acme.com:7051 --peerAddresses budget-peer1.budget.com:8051

. set-env.sh acme
chain.sh install

6. Test with one of the EP down
6.1 To kill just use the docker kill
> docker kill budget-peer1.budget.com