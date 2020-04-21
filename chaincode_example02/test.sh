#!/bin/bash

echo "Sample test script"
echo "Assumes the chaincode is already installed & instantiated"

# Set the logging level
export FABRIC_LOGGING_SPEC=WARN

# set invoke & query args
set-chain-env.sh   -i  '{"Args":["invoke","a","b","10"]}'
set-chain-env.sh   -q  '{"Args":["query","a"]}'

echo "Value of a:"
chain.sh  query

chain.sh  invoke

# give time for transaction to be committed
sleep 3s

echo "Value of a:"
chain.sh query