package main

/**
 * v2 shows the use of Composite Key functions
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

	// Used for converting []string to string
	"strings"
)

// QueryChaincode Represents our chaincode object
type QueryChaincode struct {
}

// Object type constant
const objectType = "owner~country~symbol"

// TokenBalance represent balances of different type of tokens
type TokenBalance struct {
	Symbol  string `json:"symbol"`  // symbol of the token
	Owner   string `json:"owner"`   // unique identity of the owner
	Country string `json:"country"` // country of residence
	Balance uint   `json:"balance"` // balance of the token
}


// Init Implements the Init method
func (token *QueryChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Add the data
	token.SetupSampleData(stub)

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *QueryChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the function name and parameters
	funcName, args := stub.GetFunctionAndParameters()

	if funcName == "getStateOnKey" {
		// Query with GetState
		return token.GetTokenByCompositeKey(stub, args)
	} else if funcName == "getStateRangeOnKey" {
		// Query with GetStateByPartialCompositeKey
		return token.GetTokensByPartialCompositeKey(stub, args)
	}

	// This is not good
	return shim.Error(("Bad Function Name = !!!"))
}

// GetTokenByCompositeKey executes the Range query on state data
// getStateOnKey
func (token *QueryChaincode) GetTokenByCompositeKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 3 {
		//return shim.Error("MUST provide start & end Key!!")
		fmt.Println("Partial key will return 0 results with GetState()!!!")
	}

	// Create the query key
	qryKey, errkey := stub.CreateCompositeKey(objectType, args)
	if errkey != nil {
		fmt.Printf("Error in creating key =" + errkey.Error())
		return shim.Error(errkey.Error())
	}

	fmt.Printf("Composite Key=%s\n", qryKey)

	var resultJSON = "["
	// Get the data
	dat, _ := stub.GetState(qryKey)

	// Set the data in result string
	resultJSON += string(dat)
	resultJSON += "]"

	return shim.Success([]byte(resultJSON))
}

// GetTokensByPartialCompositeKey gets the result set with the composite key
// getStateRangeOnKey
func (token *QueryChaincode) GetTokensByPartialCompositeKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	// Print statements used in dev mode
	fmt.Printf("==== Exec qry with:  ")
	fmt.Println(args)

	// Gets the state by partial query key
	QryIterator, err := stub.GetStateByPartialCompositeKey(objectType, args)
	if err != nil {
		fmt.Printf("Error in getting by range=" + err.Error())
		return shim.Error(err.Error())
	}
	var resultJSON = "["
	counter := 0
	// Iterate to read the keys returned
	for QryIterator.HasNext() {
		// Hold pointer to the query result
		var resultKV *queryresult.KV
		var err error

		// Get the next element
		resultKV, err = QryIterator.Next()
		if err != nil {
			fmt.Println("Err=" + err.Error())
			return shim.Error(err.Error())
		}

		// Split the composite key and send it as part of the result set
		key, arr, _ := stub.SplitCompositeKey(resultKV.GetKey())
		fmt.Println(key)
		resultJSON += " [" + strings.Join(arr, "~") + "] "
		counter++

	}
	// Closing
	QryIterator.Close()

	resultJSON += "]"
	resultJSON = "Counter=" + strconv.Itoa(counter) + "  " + resultJSON
	fmt.Println("Done.")
	return shim.Success([]byte(resultJSON))
}

// SetupSampleData creates multiple instances of the ERC20Token
func (token *QueryChaincode) SetupSampleData(stub shim.ChaincodeStubInterface) {

	value := []byte{0x00}
	stub.PutState(objectType, value)

	addData(stub, "BTC", "john", "USA", 26)
	addData(stub, "ETH", "john", "USA", 4)
	addData(stub, "ETH", "john", "UK", 21)

	addData(stub, "BTC", "sam", "USA", 42)
	addData(stub, "ETH", "sam", "USA", 1)

	addData(stub, "BTC", "kim", "UK", 12)

	fmt.Println("Initialized with the sample data!!")
}

func addData(stub shim.ChaincodeStubInterface, symbol, owner, country string, balance uint) {
	tokbal := TokenBalance{Symbol: symbol, Owner: owner, Country: country, Balance: balance}
	jsonTokbal, _ := json.Marshal(tokbal)
	balanceIndexKey, _ := stub.CreateCompositeKey(objectType, []string{tokbal.Owner, tokbal.Country, tokbal.Symbol})
	stub.PutState(balanceIndexKey, jsonTokbal)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. token/qry/v2\n")
	err := shim.Start(new(QueryChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
