# Shows the use of shim functions
- token.go      V3 of token
- proposal.go   Shows extraction of info from SignedProposal


Part-1
======
Demonstrates the use of stub.GetTxID(), stub.GetChannelID() and stub.GetTxTimestamp()

# 1. Init the environment
dev-init.sh

source set-env.sh acme

# 2. Install & Instantiate the chaincode

set-chain-env.sh -n token -p token/v3

chain.sh install -p

chain.sh instantiate 

# 3. Start following the log

cc-logs.sh -f

# 4. Open a new terminal

.  set-env.sh acme

chain.sh invoke

# 5. Check out the output in Terminal #1

Part-2
======

# 1. Install the latest version of the chaincode

chain.sh install-auto

set-chain-env.sh -s 2

chain.sh instantiate

# 2. Follow the logs

cc-logs.sh -f

# 3. Invoke the chaincode in terminal #2

chain.sh invoke

# 3. Observe the output in terminal #1



# GET THE Package otherwise you *MAY* get an error :)
go get github.com/golang/protobuf/proto

# Note: 
- In order to follow the code in proposal.go you MUST understand ProtoBuffers
- Refer to the documentation/links on the details of the various buffer structures
- Peer Proto Definition 
https://godoc.org/github.com/hyperledger/fabric/protos/peer
- Common Proto Definitions
https://godoc.org/github.com/hyperledger/fabric/protos/common
- Proposal proto buffer definitions
https://github.com/hyperledger/fabric-protos-go/blob/master/peer/proposal.pb.go


go get -u github.com/golang/protobuf/protoc-gen-go
