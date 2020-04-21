
# Setup the vendor dependencies
cd $GOPATH/src/token/qry/v2
./govendor.sh

This may take a few minutes

================================================
Testing Get State Range functions (token/qry/v2)
================================================
1. Launch dev environment in either mode
dev-init.sh

2. Setup the  env
. set-env.sh acme

3. Setup chaincode env. The init function takes args for setting sample data
reset-chain-env.sh
set-chain-env.sh   -n qry  -v 1.0 -p token/qry/v2 -c '{"Args": ["init"]}' -I true

4. Install and instantiate
chain.sh install  -p
chain.sh instantiate

5. Try out the Composite Key with GetState(..)
# This will execute the GetState() - with all 3 attributes of composite key
set-chain-env.sh -q '{"Args": ["getStateOnKey", "john", "USA","BTC"]}' 
chain.sh query

This will return at most 1 record matching the attributes

# This will execute the GetState() - with all 2 attributes of composite key
set-chain-env.sh -q '{"Args": ["getStateOnKey", "john", "USA"]}' 
chain.sh query

This will return at most 0 record since partial key is not supported by GetState

PS: Based on data you may still receieve a row

5. Try out the Composite Key with GetStateByPartialCompositeKey(..)

set-chain-env.sh -q '{"Args": ["getStateRangeOnKey", "john"]}' 
chain.sh query
 
</Returns all records with name=john>

# This will execute the getStateRangeOnKey
set-chain-env.sh -q '{"Args": ["getStateRangeOnKey", "john", "USA"]}' 
chain.sh query

</Returns all records with name=john & country=USA>

set-chain-env.sh -q '{"Args": ["getStateRangeOnKey", "john", "USA","BTC"]}' 
chain.sh query

</Returns all records with name=john & country=USA & symbol=BTC>

set-chain-env.sh -q '{"Args": ["getStateRangeOnKey"]}' 
chain.sh query

</Returns all records constrained by totalQueryLimit>