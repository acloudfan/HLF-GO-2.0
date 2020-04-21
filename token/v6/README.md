InvokeChaincode
===============
Shows the use of Invocation of a chaincode from another chaincode


Testing:
=======
1. Set up the chaincode by executing the utility script

$GOPATH/src/token/v6/setup-cc.sh

2. Open terminal #1 & follow the logs for token

.  set-env.sh   acme
set-chain-env.sh  -n token
bash cc-logs.sh -f

3. Open terminal #2 & follow the logs for caller

.  set-env.sh   acme
set-chain-env.sh  -n caller
bash cc-logs.sh -f

4. Open terminal #3 and setup the chaincode parameters

  .  set-env.sh   acme
  set-chain-env.sh  -n caller
  set-chain-env.sh  -n caller -i '{"args":["setOnCaller"]}'  -q '{"args":["getOnCaller"]}'

5. Execute the chaincode

  chain.sh query
  chain.sh invoke



Testing in Dev mode
===================

#Terminal 1
- Set env context
  .  set-env.sh   acme
- Launch the environment
  dev-init.sh   dev
- Set the chaincode env for 'token'
  set-chain-env.sh  -n token   -v 1.0  -p token/v5   -c '{"args":[]}'
- Launch the 'token' chaincode
  cc-run.sh

#Terminal 2
- Set env context
  .  set-env.sh   acme
- Install & Instantiate the 'token' chaincode
  chain.sh install
  chain.sh instantiate

- Set the chaincode env for 'caller'
  set-chain-env.sh  -n caller   -v 1.0  -p token/v6   -c '{"args":[]}'
- Launch the 'caller' chaincode
  cc-run.sh

#Terminal 3
- Set env context
  .  set-env.sh   acme
- Install & Instantiate the 'caller' chaincode
  chain.sh install
  chain.sh instantiate

Now we are ready to test
- Setup the query & invoke parameters
  set-chain-env.sh   -i '{"args":["setOnCaller"]}'
  set-chain-env.sh   -q '{"args":["getOnCaller"]}'

- To test execute invoke | query on caller
  chain.sh query
  chain.sh invoke