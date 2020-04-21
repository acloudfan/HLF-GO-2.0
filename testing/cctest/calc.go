package main

/**
 * Implements a calculator for testing
 **/
import (
   "fmt"

	// April 2020, Updated to Fabric 2.0 Shim
	"github.com/hyperledger/fabric-chaincode-go/shim"

	peer "github.com/hyperledger/fabric-protos-go/peer"

   // Conversion functions
   "strconv"
)

// CalcChaincode Represents our chaincode object
type CalcChaincode struct {
}


// Init Implements the Init method
func (tokencc *CalcChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Initialize the value of a token with 100
	stub.PutState("token", []byte("100"))

	// Simulate an error - uncomment the following and comment the real return
	// return shim.Error("Simulate an error !!")

	return shim.Success([]byte("100"))
}

// Invoke method
func (tokencc *CalcChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()
	
	if funcName == "invoke" {
		
		// Check the length of the argument
		if len(args) < 2 {
			return shim.Error("Operation | Argument missing !!")
		} 

		// Get the operation from args array
		op := args[0]

		fmt.Printf("Function=%s   Op=%s  \n", funcName, op)

		if op != "add" && op != "subtract" {
			return shim.Error("Unsupported operation !!!")
		}

	} else if funcName == "query"{
		tokenBytes, _ := stub.GetState("token")
		return shim.Success(tokenBytes)
	}

	//Execute the local calculator method
	return tokencc.Calculator(stub, args)
}

// Calculator - takes args[0]=operator args[1]=operand
func (tokencc *CalcChaincode) Calculator(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 2 {
		return shim.Error("Operator & Number MUST be provided !!")
	}
	// convert to int
	number, err := strconv.ParseInt(string(args[1]),10,64)
	if err != nil {
		return shim.Error("Operand MUST be a number !!")
	}
	// get the current value of the token
	tokenBytes, _ := stub.GetState("token")
	token, _ := strconv.ParseInt(string(tokenBytes),10,64)

	// Add the number to token
	if args[0] == "add" {
		token += number
	} else if args[0] == "subtract" {
		token -= number
	}

	// Store it back in the state DB
	tokenBytes = []byte(strconv.FormatInt(token,10))
	stub.PutState("token", tokenBytes)

	// Return success with new value of the token
	return shim.Success(tokenBytes)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Println("Started....")
	err := shim.Start(new(CalcChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}