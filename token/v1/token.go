package main

/**
 * v1\token
 * Shows the
 *    A) implementation of the Chaincode interface
 *    B) use of Start function in main method
 *    C) use or Error | Success for Response creation
 * 
 * PLEASE NOTE: Video may have slightly different code as the code evolves :) please ignore
 **/

import (
	// For printing messages on console
	"fmt"

	// April 2020, Updated to Fabric 2.0 Shim
	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}

// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed in v1")

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Invoke executed in v1")

	payload := []byte("This is the payload.")
	return shim.Success(payload)

	// Below statement represents an error response as status != 200
	// return peer.Response{Status:401, Message: "Unauthorized", Payload: payload}

}

// Chaincode registers with the Shim on startup
func main() {

	// Prints a message on console
	fmt.Println("Started Chaincode. token/v1")

	// Registers the chaincode with fabric runtime


	err := shim.Start(new(TokenChaincode))


	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

