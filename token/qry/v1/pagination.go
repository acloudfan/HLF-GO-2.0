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

	// April 2020 KV Interface
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"

	// JSON Encoding
	// "encoding/json"

	// Conversion functions
	"strconv"
)

// GetTokenByRangeWithPagination executes the stub funtion GetStateByRangeWithPagination
func (token *QueryChaincode) GetTokenByRangeWithPagination(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	// Check the number of arguments 
	// startKey = arg[0]  endKey = arg[1]   pagesize = arg[2]
	if len(args) < 3 {
		return shim.Error("MUST provide start, end Key & Page size!!")
	}

	// The pagesize will be int64
	pagesize, _ := strconv.ParseInt(string(args[2]),10,32)
	bookmark := ""
	var counter = 0
	var pageCounter = 0
	var resultJSON = "["
	var hasMorePages = true

	// variables to hold query iterator and metadata
	var qryIterator 	shim.StateQueryIteratorInterface
	var queryMetaData 	*peer.QueryResponseMetadata

	var err		error

	for hasMorePages {
		// Execute stub API to get the range with pagination
		qryIterator, queryMetaData, err = stub.GetStateByRangeWithPagination(args[0], args[1], int32(pagesize), bookmark)
		if err != nil {
			fmt.Printf("Error=" + err.Error())
			return shim.Error(err.Error())
		}

		var arr ="["
		var resultKV *queryresult.KV
		// Check if there are ny more records
		for qryIterator.HasNext() {

			// Get the next element
			resultKV, err = qryIterator.Next()
			
			// Increment Counter
			counter++
			if arr != "[" {
				arr += ","
			}
			arr += "\"" + resultKV.GetKey() + "\""
		}
		arr +="]"
		// Increment Page Counter
		pageCounter++

		if resultJSON != "[" {
			resultJSON += ","
		}

		resultJSON += "{\"page\":"+strconv.Itoa(pageCounter)+",\"keys\":"+arr+"}"

		// Get start key for the next page
		bookmark = queryMetaData.Bookmark

		// boomark = bland indicates not more records
		hasMorePages = (bookmark != "")

		fmt.Printf("Page: %d   Bookmark: %s \n", pageCounter, bookmark)

		// Close the iterator
		qryIterator.Close()
	}

	resultJSON += "]"

	resultJSON = "{\"count\":"+strconv.Itoa(counter)+",\"pages\":"+resultJSON+"}"

	return shim.Success([]byte(resultJSON))
}