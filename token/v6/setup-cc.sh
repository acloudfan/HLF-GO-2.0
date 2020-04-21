#!/bin/bash
#Utility script to install the chaincodes

# 1. Start the dev environment
echo "--------------------1. Starting the environment-------------"
dev-init.sh

.  set-env.sh   acme

# 2. Install the token/v5 chaincode
echo "--------------------2. Setting up token/v5-------------"
set-chain-env.sh  -n token   -v 1.0  -p token/v5   -c '{"args":[]}'
chain.sh   install   -p
chain.sh   instantiate

# 3. Install & Instantiate the caller chaincode
echo "--------------------3. Setting up caller-------------"
set-chain-env.sh  -n caller   -v 1.0  -p token/v6   -c '{"args":[]}'
chain.sh   install   -p
chain.sh   instantiate

echo "If there is no error then you are all set :-)"