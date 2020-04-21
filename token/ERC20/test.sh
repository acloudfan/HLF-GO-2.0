#!/bin/bash

# Updated April 2020 - for Fabric 2.0

# Include the Chaincode environment properties
source cc.env.sh

# Include the unit test driver
source  utest.sh

# Override the CC properties
# CC_PATH=exercise/ERC20
CC_PATH=token/ERC20
CC_NAME=erc20
CC_VERSION=1.0
CC_CHANNEL_ID=airlinechannel

# Setup the logging level for peer binary
# export CORE_LOGGING_LEVEL='INFO'
export FABRIC_LOGGING_SPEC='ERROR'

# If you would like to generate a unique CC_NAME everytime
# DO NOT USE THIS in 'dev' mode
CC_ORIGINAL_NAME=$CC_NAME
generate_unique_cc_name
set-chain-env.sh -n $CC_NAME

# Set the Organization Context to acme
set_org_context  acme

# Install
chain_install 

# Instantiate
CC_CONSTRUCTOR='{"Args":["init","ACFT","1000", "A Cloud Fan Token!!!","raj"]}'
chain_instantiate
# if needed sleep for additinal time in sec using e.g.,    txn_sleep   3s

############################### Test Case#1 #######################################
set_test_case   'Chaincode Should be initialized with "raj" as owner of 1000 tokens'
export CC_QUERY_ARGS='{"Args":["balanceOf","raj"]}'
chain_query 
assert_json_equal "$QUERY_RESULT" '.response.balance' "1000"

############################## Test Case#2 & Case#3 ################################
set_test_case   'Transfer 10 token from raj to sam - balance for raj should be 990'
# Invoke the transfer
export CC_INVOKE_ARGS='{"Args":["transfer", "raj", "sam", "10"]}'
chain_invoke 

# Get the balance for raj
CC_QUERY_ARGS='{"Args":["balanceOf","raj"]}' 
chain_query
assert_json_equal "$QUERY_RESULT"  '.response.balance'  "990"

# Case#3
set_test_case   'Sam balance should be 10'
CC_QUERY_ARGS='{"Args":["balanceOf","sam"]}' 
chain_query
assert_json_equal "$QUERY_RESULT"  '.response.balance'  "10"
# Save Sam's balance in a variable
extract_json "$QUERY_RESULT" '.response.balance'
SAM_BALANCE=$EXTRACT_RESULT

############################## Test Case#4 ##########################################
# Invoke again in the context of Budget org
set_org_context     budget

# Install the chaincode on Budget peer
chain_install
# Approve for the org
chain_approveformyorg

set_test_case   'From Budget Peer also Sam balance should be 10'
CC_QUERY_ARGS='{"Args":["balanceOf","sam"]}' 
chain_query
assert_json_equal "$QUERY_RESULT"  '.response.balance'  "$SAM_BALANCE"


# 7. Set the name to original chaincode name 
set-chain-env.sh -n $CC_ORIGINAL_NAME

