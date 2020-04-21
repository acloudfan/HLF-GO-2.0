# Shows the use of Logging Interface from chaincode

1. Setup the dev Environment
============================
# 1. Launch env in net mode
dev-init.sh

# 2. Set up chaincode environment
source set-env.sh  acme

2. Setup the Chaincode
======================

# 1. setup the env
set-chain-env.sh -n token  -v 1.0  -p token/v2 -c '{"Args":["init"]}'

# 2. Package, Install & Instantiate the chaincode

chain.sh install -p

chain.sh instantiate

# 3. Start following the log messages for the chaincode

cc-logs.sh -f

3. Invoke the chaincode
=======================
<Terminal#2>
# 1. Open a new terminal #2 window and setup the environment

source set-env.sh  acme

# 2. Execute the invoke & observe in terminal # 1

chain.sh invoke

4. Change the Logging level to ERRROR
=====================================

Change the chainncode logging level in network/config/core.yaml

Set it to level = error

5. Stop | Start the dev environment
===================================
# 1. In terminal #1 Kill the log script using ^C

# 2. In terminal #1 Restart the containers

dev-stop.sh
dev-start.sh

# 3. In terminal #1 Start the log script

cc-logs.sh -f

6. Execute the invoke on chaincode
==================================
# 1. In terminal #2 Invoke the chain code

chain.sh invoke

In terminal#1 you will see only the Error & Fatal level messages

7. IMPORTANT: Revert the log level to info in network/config/core.yaml
======================================================================
Change the log level to info otherwise you would not see the details for any Chaincode !!!


