// This chaincode calls the token/v5
// This is to JUST demonstrate the invoke mechanism
// This cc will act as a proxy
package main

import (
	"fmt"

	// April 2020, Updated for Fabric 2.0
	// Video may have shim package import for Fabric 1.4 - please ignore

	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

// CallerChaincode Represents our chaincode object
type CallerChaincode struct {
}

// Channel Name
const    Channel = "airlinechannel"
// Chaincode to be invoked
const    TargetChaincode = "token"

// Init func will do nothing
func (token *CallerChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init executed.")
	// Return success
	return shim.Success([]byte("Init Done."))
}

// Invoke method
func (token *CallerChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
		
	funcName, _ := stub.GetFunctionAndParameters()
	fmt.Println("Function=", funcName)

	if(funcName == "setOnCaller"){
		// Setup the args
		args := make([][]byte, 1)
		args[0] = []byte("set")

		// Sets the value of MyToken in token chaincode (V5)
		response := stub.InvokeChaincode(TargetChaincode, args, Channel)

		// Print on console
		fmt.Println("Receieved SET response from 'token' : "+response.String())

		return response

	} else if(funcName == "getOnCaller"){
		// Setup the args
		args := make([][]byte, 1)
		args[0] = []byte("get")

		// Gets the value of MyToken in token chaincode (V5)
		response := stub.InvokeChaincode(TargetChaincode, args, Channel)

		// Print on console
		fmt.Println("Receieved GET response from 'token' : "+response.String())

		return response
	} 

	// This is not good
	return shim.Error(("Bad Function Name from caller = "+funcName+"!!!"))
}


// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. caller/v6\n")
	err := shim.Start(new(CallerChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}