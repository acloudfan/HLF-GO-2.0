/**
 * Unit test file for calc.go
 * Demonstrates the use of MockStub for creating units tests
 */
 package main

 import (
	"fmt"

	"strings"

	"strconv"

	// go lang standard unit test package
	"testing"

	// April 2020, Updated to Fabric 2.0 Shim
	"github.com/hyperledger/fabric-chaincode-go/shim"

	peer "github.com/hyperledger/fabric-protos-go/peer"

	// Need to use the MockStub
	// 2.0 change - the MockStub was moved from package shim to shimtest !!!
	// Package may need to be installed >  go get github.com/hyperledger/fabric-chaincode-go/shimtest
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
)

// Utilitity functions
// func   SetupArgsArray(...)
// func	dumpResponse(...)

func TestCalculator(t *testing.T) {
	
	// 1. Test the initialization Logic & create the mock stub
	stub := InitChaincode(t)

	// 2. Invoke the transaction with op=add, number=10

	// Call utility function to create the [][]byte type for args
	ccArgs := SetupArgsArray("invoke","add", "10")

	// Execute the MockInvoke
	response := stub.MockInvoke("TxAdd", ccArgs)

	// Convert the received number in Payload to int
	result, _ := strconv.ParseInt(string(response.Payload),10,64)

	// Log the received value
	t.Logf("Add Received Result = %d", result)

	// To simulate failure change the result to any number other than 110
	// The function with op=subtract will not be executed due to use of FailNow()
	if result != 110 {
		// No point in going on with the next test
		t.FailNow()
	}

	// Output dumped only if the test failed
	dumpResponse(ccArgs, response, t.Failed())


	// 3. Invoke the transaction with op=subtract, number=20
	ccArgs = SetupArgsArray("invoke","subtract", "20")
	response = stub.MockInvoke("TxSubtract", ccArgs)
	result, _ = strconv.ParseInt(string(response.Payload),10,64)
	t.Logf("Subtract Received Result = %d", result)

	if result != 90 {
		// No point in going on with the next test
		t.FailNow()
	}
	
}



// InitChaicode creates the mock stub & initializes the chaincode
// 2.0 changed from shim to shimtest
func InitChaincode(t *testing.T) *shimtest.MockStub {

	// Create an instance of the MockStub
	// 2.0 changed shim to shimtest
	stub := shimtest.NewMockStub("CalcTestStub", new(CalcChaincode))

	// Execute the init
    response := stub.MockInit("mockTxId", nil)

	// Get the status
	status := response.GetStatus()

	// Log the status
	t.Logf("Received status = %d", status)

	// This is a check that indicates if there is an initialization failure
    if response.GetStatus() != shim.OK {
       t.FailNow()
	}
	
	// Return the stub instance to be used from MockInvoke
    return stub
}	

// This is a dummy test function - left as a exercise
func TestMultiply(t *testing.T) {
	// Please code it as an Exercise
}


// SetupArgsArray sets up the args arrays based on passed args
func   SetupArgsArray(funcName string, args ...string) [][]byte {
	// Create an args array with 1 additional element for the funcName
	ccArgs := make([][]byte, 1+len(args))

	// Setup the function name
	ccArgs[0] = []byte(funcName)

	// Set up the args array
	for i, arg := range args {
		ccArgs[i + 1] = []byte(arg)
	}

	return ccArgs
}

// Prints the content of the Peer Response
func	dumpResponse(args [][]byte, response peer.Response, printFlag bool) {
	if !printFlag {
		return
	}

	// Holds arg strings
	argsArray := make([]string, len(args))
	for i, arg := range args {
		argsArray[i] = string(arg)
	}
	fmt.Println("Call:    ", strings.Join(argsArray,","))
    fmt.Println("RetCode: ", response.Status)
    fmt.Println("RetMsg:  ", response.Message)
    fmt.Println("Payload: ", string(response.Payload))
}

