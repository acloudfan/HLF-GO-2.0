Custom Logger
=============
Starting Shim version 2.0 the shim.NewLogger is NOT supported.
Developer of the chaincode can use any standard logger of their
choice.

ACloudFan custom logger
=======================
This is a minimilistic logger implementation that can be used
from the chaincode. Please note this is NOT production grade
and built for demostrating the logging setup for chaincode.

Refer to the chaincode at   $GOPATH/src/acflogger

Standalone Testing
==================


Testing
=======
1. Refer to the code v2/token.go

2. Terminal#1
dev-init.sh

3. Terminal#2
. set-chain.sh acme
chain.sh -p token/v2
chain.sh install
chain.sh approveformmyorg
chain.sh commit

4. Switch to Terminal#1 & start logging for the chaincode container
cc-logs.sh -t 10 -f

5. Execute the invoke in Terminal#2
chain.sh init
chain.sh invoke
<<Check messages in Terminal#1>>

6. Change chaincode Logging level in network/config/core.yaml

7. Stop & Start the network

8. Check the logs
cc-logs.sh -t 10 -f

9.  Execute the invoke in Terminal#2
chain.sh init
chain.sh invoke
<<Check messages in Terminal#1>>