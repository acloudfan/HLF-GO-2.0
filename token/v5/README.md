# Shows the use of PutState | GetState | DelState

# Test the original token.go 
1. Launch the env in dev mode             dev-init.sh
2. Setup the env context                  .  set-env.sh acme
3. Setup the chain env                    set-chain-env.sh  -n token -v 1.0 -p token/v5  -c '{"args":[]}' 
4. In a new terminal set the env context     .  set-env.sh acme
5. Install                                chain.sh install -p
6. Instantiate                            chain.sh instantiate
7. Follow logs in terminal #1             cc-logs.sh -f
8. Setup invoke & query args              set-chain-env.sh   -i   '{"args":["set"]}' -q   '{"args":["get"]}' 

To Test execute invoke/query .... everytime invoke is called value will increment by 10
9. Execute invoke                   chain.sh  invoke
10. Execute query                   chain.sh  query


Solution to exercise
====================
1. Solution is in the delstate.go
2. Add the following to the invoke function
	// Solution to the exercise
	   else if(funcName == "delete"){


		// Delete the token
		return DeleteToken(stub)
	}  

Testing the Solution
====================
PS: You may use it with original version of the token/v5 but keep in mind that the last step will FAIL which is fine as the "delete" function is not supported by the original token/v5 implementation

# Using the test.sh file
The test may be run net & dev mode. 
Prior to running the test install & instantiate the chaincode

1. Launch the setup in net mode
   dev-init.sh 

2. Set the org context
   . set-env.sh acme

3. Set the chaincode env
   set-chain-env.sh -n token -v 1.0 -p token/v5
   set-chain-env.sh   -i   '{"args":["set"]}' -q   '{"args":["get"]}' 

4. Install & Instantiate
   chain.sh  install -p

   set-chain-env.sh  -c  '{"args":[]}'
   chain.sh  instantiate

5. Launch the test
   chain.sh query
   chain.sh invoke

6. Delete & Test again
   chain.sh query




