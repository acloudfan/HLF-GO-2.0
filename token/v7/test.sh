#!/bin/bash

# Unit Testing for v8/token
# Assumes that chaincode is installed & instantiated outside of this script
# OK to use it in dev & net mode

# Include the unit test driver
source  utest.sh

# Include the Chaincode environment properties
source cc.env.sh

# Setup the logging level for peer binary
export CORE_LOGGING_LEVEL='ERROR'

# Set the Organization Context to acme
set_org_context  acme

############################### Install & Instantiate #############################
# Ignored in 'dev' mode
generate_unique_cc_name

# Set the Organization Context to acme
set_org_context  acme

# Install
chain_install 

# Instantiate
CC_CONSTRUCTOR='{"Args":["init"]}'
chain_instantiate

############################### Test Case#1 #######################################
set_test_case   'Add Token_1=100 & Token_2=200'
export CC_INVOKE_ARGS='{"Args":["add", "token-1", "100"]}'
chain_invoke 
export CC_INVOKE_ARGS='{"Args":["add", "token-2", "200"]}'
chain_invoke 

export CC_QUERY_ARGS='{"Args":["get", "token-1"]}'
chain_query
MY_TOKEN_VALUE_1=$QUERY_RESULT

export CC_QUERY_ARGS='{"Args":["get", "token-2"]}'
chain_query 
MY_TOKEN_VALUE_2=$QUERY_RESULT
echo $QUERY_RESULT
# Print the value of tokerns
print_info  "MY_TOKEN_VALUE_1=$MY_TOKEN_VALUE_1 MY_TOKEN_VALUE_2=$MY_TOKEN_VALUE_2 "

# Complex expressions
if (((MY_TOKEN_VALUE_1 == 100) && (MY_TOKEN_VALUE_2 == 200)))  ; then
    assert_boolean  "true" 
else
    assert_boolean  "false"
fi

############################### Test Case#2 #######################################
set_test_case   'Add 10 to both the token-1 and token-2'
export CC_INVOKE_ARGS='{"Args":["addNumber", "token-1", "10"]}'
chain_invoke 
export CC_INVOKE_ARGS='{"Args":["addNumber", "token-2", "10"]}'
chain_invoke 

export CC_QUERY_ARGS='{"Args":["get", "token-1"]}'
chain_query
MY_TOKEN_VALUE_1=$QUERY_RESULT

export CC_QUERY_ARGS='{"Args":["get", "token-2"]}'
chain_query 
MY_TOKEN_VALUE_2=$QUERY_RESULT
echo $QUERY_RESULT
# Print the value of tokerns
print_info  "MY_TOKEN_VALUE_1=$MY_TOKEN_VALUE_1 MY_TOKEN_VALUE_2=$MY_TOKEN_VALUE_2 "

# Complex expressions
if (((MY_TOKEN_VALUE_1 == 110) && (MY_TOKEN_VALUE_2 == 210)))  ; then
    assert_boolean  "true" 
else
    assert_boolean  "false"
fi