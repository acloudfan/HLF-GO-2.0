#!/bin/bash

# Requires the Fabric environment to be setup in net mode

# 1. Setup the chaincode env
set-chain-env.sh  -n   cctest
set-chain-env.sh  -p   testing/cctest   -v  1.0
source cc.env.sh

# 2. Include the unit test driver
source  utest.sh

# OPTIONAL - override the wait time after instantiate/invoke
# PS: Low value may lead to endorsement errors
# export TXN_WAIT_TIME=3s

# 3. Setup the logging spec - error is suggested
export FABRIC_LOGGING_SPEC='ERROR'

# 4. Generates a unique name everytime the script is executed
#    Comment this if you would like to use the same instance of
#    the chaincode *but* keep in mind that the state may change with every run

# Retain the original chaincode name otherwise it will be replaced by the random name !!
# Use this to set it back at the end of the test case implementation
# set-chain-env.sh -n $CC_ORIGINAL_NAME
CC_ORIGINAL_NAME=$CC_NAME
generate_unique_cc_name
set-chain-env.sh -n $CC_NAME

######## Install & Instantiate the chaincode ######
# 1. Set the org context
set_org_context  acme

# 2. Install
chain_install 

# 3. Setup init arguments

# Setup arguments
CC_CONSTRUCTOR='{"Args":[]}'

# 4. Instantiate
chain_instantiate

######## Test Case #1 ##################################
# 1. Start the test case with a description
set_test_case 'On adding 10 Chaincode should return 110'

# 2. Set the invoke args
export CC_INVOKE_ARGS='{"Args":["invoke", "add", "10"]}'

# 3. Invoke the chaincode
chain_invoke

# 4. Setup the Query args
export CC_QUERY_ARGS='{"Args":["query"]}'

# 5. Execute the query
chain_query
print_info  "Query Result = $QUERY_RESULT"

# 6. Assert on equal = 110
assert_equal  "$QUERY_RESULT"  "110"

######### Test Case #2 ##################################
# 1. Start the test case with a description
set_test_case 'On subtracting 20 Chaincode should return 90'

# 2. Set the invoke args
export CC_INVOKE_ARGS='{"Args":["invoke", "subtract", "20"]}'

# 3. Invoke the chaincode
chain_invoke

# 4. Setup the Query args
export CC_QUERY_ARGS='{"Args":["query"]}'

# 5. Execute the query
chain_query
print_info  "Query Result = $QUERY_RESULT"

# 6. Assert on equal = 110
assert_equal  "$QUERY_RESULT"  "90"

# 7. Set the name to original chaincode name 
set-chain-env.sh -n $CC_ORIGINAL_NAME