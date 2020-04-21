package main

/**
 * Shows how to use the GetQueryResultWithPagination
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


// GetDatesByPrice returns the data for the specified date
// Get the dates on which the prices of Cryptocurrency higher than the specified price
func GetDatesByPrice(stub shim.ChaincodeStubInterface,args []string) peer.Response {
	var pagesize int32 = 20
	bookmark := ""
	var counter = 0
	var pageCounter = 0
	var resultJSON = "["
	var hasMorePages = true

	// Query string - multi line strings use backward single quote
	query := `{
		"selector": {
		   "docType": "CryptocoinTransactions",
		   "usdPrice": {
			  "$gte": `
	query += args[0] 

	query += `}
		},
		"fields": [
		   "txnDate",
		   "usdPrice"
		]
	 }`

	// Print the received query on the console
	fmt.Printf("Query JSON=%s \n", query)

	 // variables to hold query iterator and metadata
	var qryIterator 	shim.StateQueryIteratorInterface
	var queryMetaData 	*peer.QueryResponseMetadata
	var err		error
	// start the pagination read loop
	lastBookmark := ""
	for hasMorePages {
		// execute the rich query
		qryIterator, queryMetaData, err = stub.GetQueryResultWithPagination(query, pagesize,bookmark)
		if err != nil {
			fmt.Printf("GetQueryResultWithPagination Error=" + err.Error())
			return shim.Error(err.Error())
		}
		var arr ="["
		var resultKV *queryresult.KV
		// Result read loop only if we received a different bookmark
		if lastBookmark != queryMetaData.Bookmark {
			for qryIterator.HasNext() {

				// Get the next element
				resultKV, err = qryIterator.Next()
				
				// Increment Counter
				counter++
				if arr != "[" {
					arr += ","
				}
				arr += "\"" + string(resultKV.GetValue()) + "\""
			}
			arr +="]"

			// Increment Page Counter
			pageCounter++

			if resultJSON != "[" {
				resultJSON += ","
			}
			
			fmt.Printf("Page: %d \n", pageCounter)

			resultJSON += "{\"page\":"+strconv.Itoa(pageCounter)+",\"DatesPrice\":"+arr+"}"
		} 

		// Get start key for the next page
		bookmark = queryMetaData.Bookmark

		// boomark = blank indicates no more records
		hasMorePages = (bookmark != "" && lastBookmark != bookmark)
		lastBookmark = bookmark

		// Close the iterator
		qryIterator.Close()
	}

	resultJSON += "]"
	resultJSON = "{\"count\":"+strconv.Itoa(counter)+",\"pages\":"+resultJSON+"}"

	return shim.Success([]byte(resultJSON))
}

