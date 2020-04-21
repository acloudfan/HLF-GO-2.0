package main

/**
 * v1 shows the use of Range functions
 **/

import (
	// For printing messages on console
	"fmt"

	// April 2020, Updated to Fabric 2.0 Shim
	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"

	// KV Interface
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"

	// JSON Encoding
	"encoding/json"

	// Conversion functions
	"strconv"
)

// QueryChaincode Represents our chaincode object
type QueryChaincode struct {
}

// SimpleToken represents a standard token implementation
type SimpleToken struct {
	Symbol      string `json:"symbol"`
	TotalSupply uint64 `json:"totalSupply"`
}

// Init Implements the Init method
// args[0]=Start index, args[1]=Count of records
// If args not provided then Initializes the state with 10 records (key1 to key10)
func (token *QueryChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the function name and parameters
	_, args := stub.GetFunctionAndParameters()

	// Add the data
	token.SetupSampleData(stub,args)

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *QueryChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	
	// Get the function name and parameters
	funcName, args := stub.GetFunctionAndParameters()

	if funcName == "GetTokenByRange" {
		return token.GetTokenByRange(stub, args)
	} else if funcName == "GetTokenByRangeWithPagination" {
		// To be implemented in the exercise
		return token.GetTokenByRangeWithPagination(stub, args)
	}

	// This is not good
	return shim.Error(("Bad Function Name = " + funcName + "!!!"))
}

// GetTokenByRange executes the Range query on state data
// Returns the data in the specified key range
// args[0]=startKey   args[1]=endKey
func (token *QueryChaincode) GetTokenByRange(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	// Check the number of args
	if len(args) < 2 {
		return shim.Error("MUST provide start & end Key!!")
	}

	// Get the data by range
	QryIterator, err := stub.GetStateByRange(args[0], args[1])
	if err != nil {
		fmt.Printf("Error=" + err.Error())
		return shim.Error(err.Error())
	}

	var counter = 0
	var resultJSON = "["
	// Read the returned results in a loop
	for QryIterator.HasNext() {

		// Hold pointer to the query result
		var resultKV *queryresult.KV
		var err error

		// Get the next element
		resultKV, err = QryIterator.Next()

		if err != nil {
			fmt.Println("Err=" + err.Error())
		} else {
			// Create the data string
			var tokenData string
			tokenData = "{\"key\":\"" + resultKV.GetKey() + "\",\"token\":" + string(resultKV.GetValue()) + "}"
			if counter > 0 {
				resultJSON += "," + "\n "
			}
			resultJSON += tokenData
		}
		// Increment counter
		counter++
	}
	resultJSON += "]"
	resultJSON = "{ \"count\":" + strconv.Itoa(counter) + ",\"queryResult\":" + resultJSON + "}"

	// Close the query iterator instance
	QryIterator.Close()

	return shim.Success([]byte(resultJSON))
}




// SetupSampleData creates multiple instances of the SimpleToken
// Requires the startIndex & the count of records to be created in the State DB
// Once data is setup, we will use it for querying using range function
func (token *QueryChaincode) SetupSampleData(stub shim.ChaincodeStubInterface,args []string) {

	var startIndex, recordCount int64
	if len(args) < 2 {
		startIndex = 1
		recordCount = 10
	} else {
		// error not being checked (to keep code simple)
		startIndex, _ = strconv.ParseInt(args[0], 10, 64)
		recordCount, _ = strconv.ParseInt(args[1], 10, 64)
	}

	// Simply print a message
	fmt.Printf("Init executed in qry   startIndex=%d   recordCount=%d \n", startIndex, recordCount)

	var simple SimpleToken
	for i := startIndex; i < (startIndex+recordCount); i++ {
		// Convert the int to string
		var iStr = strconv.Itoa(int(i))
		// Create an instance of the Simple token
		simple = SimpleToken{Symbol: "TOK" + iStr, TotalSupply: 1000}
		// Marshall the token struct to JSON
		jsonSimple, _ := json.Marshal(simple)
		// Add it to the state DB
		stub.PutState("key"+iStr, jsonSimple)
	}

	fmt.Printf("Initialized Chaincode with %d Tokens ", recordCount)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. token/qry/v1\n")
	err := shim.Start(new(QueryChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
