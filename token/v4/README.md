# Demonstrates the use of Argument functions on the stub

Test out v4/token
=================

# 1. Start the environment
<<Terminal#1>>
dev-init.sh

source set-env.sh acme

# 2.0 Setup the chaincode
<<Terminal#2>>
set-chain-env.sh   -n token   -v 1.0   -p token/v4

chain.sh install -p

chain.sh instantiate 

# 3.0 Follow the logs

<<Terminal#1>>
cc-logs.sh -f


# 4.0 Invoke the function
set-chain-env.sh   -i    '{"args":["func-name"]}'
chain.sh   invoke

Observe in Terminal #1

# 5
set-chain-env.sh   -i    '{"args":["func-name","param-1","param-2"]}'
chain.sh   invoke

Observe in Terminal #1


