# Private data
Fabric 2.0
https://hyperledger-fabric.readthedocs.io/en/release-2.0/private-data/private-data.html



Print the commands
==================
set-chain-env.sh   -p  token/priv
set-chain-env.sh   -R  pcollection.json
chain.sh  approveformyorg   -o
chain.sh  commit -o
chain.sh  checkcommitreadiness -o


Install, Instantiate & Test  token/priv
=======================================
dev-init.sh

# 1. Setup the chaincode parameters
set-chain-env.sh  -n priv -v 1.0 -p token/priv  -c '{"Args": ["init"]}' -C airlinechannel
set-chain-env.sh -q '{"Args": ["Get"]}'
set-chain-env.sh -R pcollection.0.json
set-chain-env.sh -I false  <<Chaincode has no need for initialization, -c will be ignored>>

# 2. Acme : Install | Approve the chaincode
. set-env.sh acme

chain.sh install -p

chain.sh instantiate

[Commands]
Generate the peer commands / copy paste commands to execute
OR remove the -o flag to execute

chain.sh package -o
chain.sh install -o
chain.sh approveformyorg -o   

# 3. Budget : Install | Approve | Commit the chaincode
. set-env.sh budget

chain.sh install <<-p not needed as the pacage is already created>>
chain.sh approveformyorg 

[Commands]
chain.sh approveformyorg  -o   

# 4. Invoke as Acme

. set-env.sh acme

set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Acme has set the OPEN data"]}'
chain.sh invoke

set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Acme has set the SECRET data"]}'
chain.sh invoke

chain.sh query

# 5. Invoke as Budget

. set-env.sh budget

set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Budget has set the OPEN data"]}'
chain.sh invoke

set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Budget has set the SECRET data"]}'
chain.sh invoke

chain.sh query

# 6. Check if Acme can get the data 

. set-env.sh acme

chain.sh query


Exercise : Update PDC (Reqires use of memberOnlyWrite)
====================================================
The PDC Attribute memberOnlyWrite is new in Fabric 2.x
It allows flexibility to define whether unauthorzed members can write to a PDC or not

# Validate 
get-cc-info.sh -e

# Checkout current collections setup
chain.sh querycommitted -j  > temp.json


# 1. Acme : Update the policy to pcollection.2.json
. set-env.sh acme
set-chain-env.sh -R pcollection.2.json
set-chain-env.sh -s 2
chain.sh approve


# 2. Budget : Approve and Commit
. set-env.sh budget
chain.sh approve
chain.sh commit

# 4. Acme : Invoke the Chaincode
Invoke for both will succeed for Acme

. set-env.sh acme

set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Acme has set the OPEN data"]}'
chain.sh invoke

set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Acme has set the SECRET data"]}'
chain.sh invoke


# 5. Budget : Invoke the Chaincode

. set-env.sh budget

* This invoke will succeed because memberOnlyWrite=false

set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Budget has set the OPEN data"]}'
chain.sh invoke

* This invoke will Fail because memberOnlyWrite=true
set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Budget has set the SECRET data"]}'
chain.sh invoke

Walkthrough : Implicit private collections
==========================================

# Start by initializing the environment
dev-init.sh

# 1. Setup the chaincode parameters - DO NOT set the PDC
reset-chain-env.sh 

set-chain-env.sh  -n priv -v 1.0 -p token/priv  -I false
set-chain-env.sh -q '{"Args": ["GetFromImplicitCollection","_implicit_org_AcmeMSP"]}'

# Confirm that Collection is NOT set
show-chain-env.sh


# 2. Acme : Install |  Approve & Commit the chaincode
. set-env.sh acme

chain.sh install -p
chain.sh instantiate

# 3. Budget : Install | Approve the chaincode
. set-env.sh budget

chain.sh install 
chain.sh approveformyorg 

# 4. Budget : Invoke the Chaincode & Query
. set-env.sh budget

set-chain-env.sh -i '{"Args": ["Set","_implicit_org_AcmeMSP", "Budget setting the data in Implicit collection on Acme"]}'

# Set Acme as the endorser - otherwise call will FAIL
set-chain-env.sh -e acme
chain.sh invoke

chain.sh query

# 5. Acme : Query

. set-env.sh acme

chain.sh query



EXPT for EP with Implicit
=========================
set-chain-env.sh -g "OR('BudgetMSP.member')" 
chain.sh upgrade-auto

Test for EP Override
====================










Test the setup
==============
1. Start the Environment

# Start the environment in net mode
dev-init.sh -e

reset-chain-env.sh

set-chain-env.sh  -n priv -v 1.0 -p token/priv -c '{"Args": ["init"]}' -C airlinechannel
# Use the -R option to set the PDC
# At instantiation chain.sh will specify the full path to PDC collection
set-chain-env.sh -R pcollection.0.json

Install & Instantiate
. set-env.sh acme
chain.sh install
chain.sh instantiate

. set-env.sh budget
chain.sh install

2. Invoke the Set as ACME & Query

<Terminal#1>

# Invoke to set the value for 2 tokens
. set-env.sh acme
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Acme has set the OPEN data"]}'
chain.sh invoke
set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Acme has set the SECRET data"]}'
chain.sh invoke
# Get the value for 2 tokens
set-chain-env.sh -q '{"Args": ["Get"]}'
chain.sh query

3. Invoke the Set as BUDGET & Query

<Terminal#2>
. set-env.sh budget

set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Budget has set the OPEN data"]}'
chain.sh invoke

set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Budget has set the SECRET data"]}'
chain.sh invoke

# Get the value for 2 tokens - Budget will NOT seet the value for protected token
chain.sh query         

4. Query as ACME
<Terminal#1>
. set-env.sh acme
chain.sh query  

. set-env.sh acme

Exercise
========
Extend the priv chaincode - add a function to delete the key in specific collection

. set-env.sh acme
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Acme has set the OPEN data"]}'
chain.sh invoke
set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Acme has set the SECRET data"]}'
chain.sh invoke


chain.sh query
set-chain-env.sh -i '{"Args": ["Del", "AcmeBudgetOpen"]}'
chain.sh invoke

chain.sh query


Experimental
============
set-chain-env.sh -i '{"Args": ["Del", "MemberOnlyTest"]}'

Testing in Dev Mode
====================
Use the instructions below to test the PDC in in DEV mode

Install & Instantiate
======================
Regular install process
Instantiate requires the collection.json to be specified
--collections-config

# Start the environment in Dev mode
dev-init.sh dev
set-chain-env.sh  -n priv -v 1.0 -p token/priv -c '{"Args": ["init"]}' 
# Use the -R option to set the PDC
# At instantiation chain.sh will specify the full path to PDC collection
set-chain-env.sh -R pcollection.json

# Launch the chaincode instance on Acme Peer
<Terminal#1>
. set-env.sh acme
cc-run.sh

# Launch the chaincode instance on Budget Peer
<Terminal#1>
. set-env.sh budget
cc-run.sh

<Terminal#3>
# Install the chaincode on Acme & Budget peers
. set-env.sh acme
chain.sh install

chain.sh instantiate-priv

. set-env.sh budget
chain.sh install

Test
====
Invalid collection name will lead to error


1. Acme can set both the public & private data
. set-env.sh acme
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Acme has set the OPEN data"]}'
chain.sh invoke
set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Acme has set the SECRET data"]}'
chain.sh invoke

2. Acme can get both public and secret data
set-chain-env.sh -q '{"Args": ["Get"]}'
chain.sh query

3. Budget can the public & the private data
# Switch context to budget
. set-env.sh budget

# Change the parameters for invoke
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Budget has set the OPEN data"]}'
chain.sh invoke

set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Budget has set the SECRET data"]}'
chain.sh invoke

4. Budget can get only the public data
chain.sh query







============================================ Fabric 1.4 ===========================
Install & Instantiate
=====================
dev-init.sh -e

. set-env.sh acme

reset-chain-env.sh
set-chain-env.sh  -n priv -v 1.0 -p token/priv -c '{"Args": ["init"]}' -C airlinechannel
set-chain-env.sh -R pcollection.json

Exercise (v.1.4)
================
Install & Instantiate the token/priv chaincode using "peer chaincode instantiate .." command

<Solution>
Setup the environment variables
chain.sh install
. cc.env.sh
peer chaincode instantiate -C "$CC_CHANNEL_ID" -n "$CC_NAME"  -v "$CC_VERSION" -c "$CC_CONSTRUCTOR" -o "$ORDERER_ADDRESS"  --collections-config "$GOPATH/src/token/priv/pcollection.json"







Exercise (v.1.4)
================
Install & Instantiate the token/priv chaincode using "peer chaincode instantiate .." command

<Solution>
Setup the environment variables
chain.sh install
. cc.env.sh
peer chaincode instantiate -C "$CC_CHANNEL_ID" -n "$CC_NAME"  -v "$CC_VERSION" -c "$CC_CONSTRUCTOR" -o "$ORDERER_ADDRESS"  --collections-config "$GOPATH/src/token/priv/pcollection.json"


Test the setup
==============
1. Start the Environment

# Start the environment in net mode
dev-init.sh -e

reset-chain-env.sh

set-chain-env.sh  -n priv -v 1.0 -p token/priv -c '{"Args": ["init"]}' -C airlinechannel
# Use the -R option to set the PDC
# At instantiation chain.sh will specify the full path to PDC collection
set-chain-env.sh -R pcollection.0.json

Install & Instantiate
. set-env.sh acme
chain.sh install
chain.sh instantiate

. set-env.sh budget
chain.sh install

2. Invoke the Set as ACME & Query

<Terminal#1>

# Invoke to set the value for 2 tokens
. set-env.sh acme
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Acme has set the OPEN data"]}'
chain.sh invoke
set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Acme has set the SECRET data"]}'
chain.sh invoke
# Get the value for 2 tokens
set-chain-env.sh -q '{"Args": ["Get"]}'
chain.sh query

3. Invoke the Set as BUDGET & Query

<Terminal#2>
. set-env.sh budget

set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Budget has set the OPEN data"]}'
chain.sh invoke

set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Budget has set the SECRET data"]}'
chain.sh invoke

# Get the value for 2 tokens - Budget will NOT seet the value for protected token
chain.sh query         

4. Query as ACME
<Terminal#1>
. set-env.sh acme
chain.sh query  

. set-env.sh acme

Exercise
========
Extend the priv chaincode - add a function to delete the key in specific collection

. set-env.sh acme
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Acme has set the OPEN data"]}'
chain.sh invoke
set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Acme has set the SECRET data"]}'
chain.sh invoke


chain.sh query
set-chain-env.sh -i '{"Args": ["Del", "AcmeBudgetOpen"]}'
chain.sh invoke

chain.sh query


Experimental
============
set-chain-env.sh -i '{"Args": ["Del", "MemberOnlyTest"]}'

Testing in Dev Mode
====================
Use the instructions below to test the PDC in in DEV mode

Install & Instantiate
======================
Regular install process
Instantiate requires the collection.json to be specified
--collections-config

# Start the environment in Dev mode
dev-init.sh dev
set-chain-env.sh  -n priv -v 1.0 -p token/priv -c '{"Args": ["init"]}' 
# Use the -R option to set the PDC
# At instantiation chain.sh will specify the full path to PDC collection
set-chain-env.sh -R pcollection.json

# Launch the chaincode instance on Acme Peer
<Terminal#1>
. set-env.sh acme
cc-run.sh

# Launch the chaincode instance on Budget Peer
<Terminal#1>
. set-env.sh budget
cc-run.sh

<Terminal#3>
# Install the chaincode on Acme & Budget peers
. set-env.sh acme
chain.sh install

chain.sh instantiate-priv

. set-env.sh budget
chain.sh install

Test
====
Invalid collection name will lead to error


1. Acme can set both the public & private data
. set-env.sh acme
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Acme has set the OPEN data"]}'
chain.sh invoke
set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Acme has set the SECRET data"]}'
chain.sh invoke

2. Acme can get both public and secret data
set-chain-env.sh -q '{"Args": ["Get"]}'
chain.sh query

3. Budget can the public & the private data
# Switch context to budget
. set-env.sh budget

# Change the parameters for invoke
set-chain-env.sh -i '{"Args": ["Set","AcmeBudgetOpen", "Budget has set the OPEN data"]}'
chain.sh invoke

set-chain-env.sh -i '{"Args": ["Set","AcmePrivate", "Budget has set the SECRET data"]}'
chain.sh invoke

4. Budget can get only the public data
chain.sh query
