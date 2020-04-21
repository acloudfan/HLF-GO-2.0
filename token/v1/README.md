Launch the environment
======================
Launch Dev Environment:     dev-init.sh  -e

Try out v1/token (Fabric version 2.x)
=====================================

# 1 Install the chaincode 

. set-env.sh acme

- Setup the chaincode environment
set-chain-env.sh   -n  token   -p token/v1    -c '{"Args": ["init"]}'  

Option-1
--------
chain.sh package
chain.sh install

Option-2
--------
# -p will lead to call to package before install
chain.sh install -p

# 2 Approve and commit the chaincode

Option-1
--------
chain.sh approveformyorg

(Optional)  chain.sh checkcommitreadiness

chain.sh commit

(Optional)  chain.sh querycommitted

Option-2
--------
chain.sh instantiate

# 3. Check out the transactions in the explorer

http://localhost:8080/#/transactions

# 4. Follow logs for chaincode container

cc-logs.sh -t 10 -f

# 5. Open terminal # 2
You will run the commands in terminal #2 and see the chaincode logs in Terminal#1
source set-env.sh  acme

# 6. Initialize the chaincode terminal #2
chain.sh  init

# 7. Query the chaincode  terminal #2
Observe the messages in terminal #1

chain.sh  query
chain.sh invoke
chain.sh query




Try out v1/token (1.4)
======================
#1  

<<Terminal #1>>  


- Setup the organization context to acme
. set-env.sh acme

- Setup the chaincode environment
set-chain-env.sh   -n  token   -p token/v1    -c '{"Args": ["init"]}'  

#2

<<Terminal #2>>

- Start the chaincode 
. set-env.sh acme
cc-run.sh

#3

<<Terminal #1>>

- Install & Instantiate the chaincode
chain.sh    install 
chain.sh    instantiate                             <<Observe terminal#2>>

Checkout explorer - you should see 1 transaction against the chaicode 'token'

#4

<<Terminal #1>>

- Setup and execute the invoke & query API
set-chain-env.sh  -i '{"Args": ["invoke"] }'
chain.sh  invoke                                    <<Observe terminal#2>>

set-chain-env.sh  -q '{"Args": ["invoke"] }' 
chain.sh  query

Checkout explorer - you should see 2 transaction against the chaicode 'token'
