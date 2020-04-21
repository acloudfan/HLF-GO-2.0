package main
/**
 * tokenv3
 * Shows the use of ChaincodeStub API
 *    A) Transaction information
 *    B) Transaction timestamp
 *    C) Channel ID
 **/
import (
	"fmt"

	// April 2020, Updated for Fabric 2.0
	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"

	// Used for formatting the timestamp
	"time"

)

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}


// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed")

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Invoke executed ")

	// V3   Print the transaction ID
	fmt.Printf("GetTxID() => %s\n", stub.GetTxID())

	// V3   Print the channel ID
	fmt.Println("GetChannelID() =>", stub.GetChannelID())

	// V3   Print the transaction Timestamp
	TxTimestamp, _ := stub.GetTxTimestamp()
	timeStr := time.Unix(TxTimestamp.GetSeconds(),0)
	fmt.Printf("GetTxTimestamp() => %s\n", timeStr)

	
	// V3
	// Extract the information from proposal 
	// PrintSignedProposalInfo(stub)

	// V3
	// Will receieve empty map unless client set the transient data in Tx Proposal
	// transientData, _ := stub.GetTransient()
	// fmt.Println("GetTransient() =>", transientData)

	// PrintCreatorInfo(stub)

	return shim.Success(nil)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Println("Started Chaincode. token/v3")
	err := shim.Start(new(TokenChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}