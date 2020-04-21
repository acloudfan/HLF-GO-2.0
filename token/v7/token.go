/**
 * Allows the managing of multiple tokens
 *
 * Token1    x       0
 * Token1    x       10
 * Token2    y       5
 *
 **/
 package main

 import (
	"fmt"

	// April 2020, Updated for Fabric 2.0
	// Video may have shim package import for Fabric 1.4 - please ignore

	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"

	// Conversion functions
	"strconv"
)

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}

// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// V7
	// Nothing done in init
	fmt.Println("init()")

	return shim.Success([]byte("true"))
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()
	fmt.Println("Function=", funcName)

	if funcName == "add" {
		return AddToken(stub, args)
	} else if funcName == "delete" {
		return DeleteToken(stub, args)
	} else if funcName == "get" {
		return GetToken(stub, args)
	} else if funcName == "addNumber" {
		return AddNumber(stub, args)
	} else if funcName == "exists" {
		return TokenExists(stub, args)
	}

	// This is not good
	return shim.Error(("Bad Function Name = "+funcName+"!!!"))
}

// AddToken adds a token to the state
// Otherwise returns -1
// Takes 2 argumets = tokenName & a positive number
// Sets an event AddTokenEvent
func	AddToken(stub shim.ChaincodeStubInterface,args []string) peer.Response {
	// Check if correct number of args were passed
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments - MUST provide token name & initial value!!!")
	}

	// Check if the token already exist? if it exists then return false
	value, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if(value != nil){
		fmt.Println("Token already exist!!!")
		return shim.Error("Token name MUST be unique!!!")
	}
	// Convert the passed value to number
	_, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(args[0], []byte(args[1]))

	return shim.Success([]byte("true"))
}

// GetToken returns the value of the token
// Otherwise returns -1
// Takes an argumet = tokenName
func	GetToken(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// Check if correct number of args were passed
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments - MUST provide token name!!!")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if(value == nil){
		fmt.Println("Token Not found!!!")
		return shim.Error("Token Not Found!!")
	}

	fmt.Println(args[0])
	fmt.Println(value)

	return shim.Success([]byte(value))
}
// DeleteToken deletes the token
// Returns true | false
// Takes an argumet = tokenName
func	DeleteToken(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// Check if correct number of args were passed
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments - MUST provide token name!!!")
	}

	value, err := stub.GetState(args[0])
	if(err != nil){
		return shim.Error(err.Error())
	}

	// Check if token exists - if it doesnt then return false
	if value != nil {
		err := stub.DelState(args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte("true"))
	}
	// format boolean to string
	return shim.Success([]byte("false"))
}

// AddNumber adds the passed unsigned number to the token
// Takes 2 argumets = tokenName & a positive number
// Returns   true | false
func    AddNumber(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	// Check if correct number of args were passed
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments!!!")
	}

	// Convert the number in arg to int
	number, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if token exists - if it doesnt return false
	value, err := stub.GetState(args[0])
	if(err != nil){
		return shim.Error(err.Error())
	}

	valueInt, err := strconv.Atoi(string(value))
	
	// Add the number
	valueInt += number

	// Update the value in state database
	stub.PutState(args[0], []byte(strconv.FormatInt(int64(valueInt),10)))

	return shim.Success([]byte("true"))
}

// TokenExists returns true if the token exists
func  TokenExists(stub shim.ChaincodeStubInterface, args []string) peer.Response  {
	// Check if correct number of args were passed
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments!!!")
	}

	value, err := stub.GetState(args[0])
	if(err != nil){
		return shim.Error(err.Error())
	}
	if value != nil {
		return shim.Success([]byte("true"))
	} 
	
	return shim.Success([]byte("false"))
	
}



// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode.\n")
	err := shim.Start(new(TokenChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
