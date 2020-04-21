package main

/**
 * Shows how to use the "GetQueryResult" function
 **/

import (
	// For printing messages on console
	"fmt"

	// April 2020, Updated to Fabric 2.0 Shim
	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"

	// KV Interface
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"

	// Conversion functions
	"strconv"
)

// ExecuteRichQuery executes the passed query on the data
// Result count restricted by totalQueryLimit - DO NOT use for lage result sets
func ExecuteRichQuery(stub shim.ChaincodeStubInterface,args []string) peer.Response {

	// Query JSON received as argument
	qry:=args[0]

	// Print the received query on the console
	fmt.Printf("Query JSON=%s \n\n", qry)

	// GetQueryResult
	QryIterator, err := stub.GetQueryResult(qry)

	// Return if there is an error
	if err != nil {
		fmt.Println(err.Error())
		return shim.Success([]byte("Error: "+err.Error()))
	}

	// Iterate through the result set
	counter := 0
	for QryIterator.HasNext() {
		// Hold pointer to the query result
		var resultKV *queryresult.KV
		var err error

		// Get the next element
		resultKV, err = QryIterator.Next()

		// Return if there is an error
		if err != nil {
			fmt.Println("Err=" + err.Error())
			return shim.Success([]byte("Error in parse: "+err.Error()))
		}

		// Increment the counter
		counter++
		key := resultKV.GetKey()
		value := string(resultKV.GetValue())

		// Print the receieved result on the console
		fmt.Printf("Result# %d   %s   %s \n\n", counter, key, value)
		
	}

	// Close the iterator
	QryIterator.Close()

	// Return the count
	total := "Count="+strconv.Itoa(counter)
	return shim.Success([]byte(total))
}
