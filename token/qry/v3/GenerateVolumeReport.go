/**
 * Shows how the aggregattion functions may be implemented
 **/
 package main

 import (
	// For printing messages on console
	"fmt"

	// April 2020, Updated to Fabric 2.0 Shim
	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"

	// KV Interface
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"


)

// GenerateVolumeReport generates a report for prices between given date range
// Price data is sent back in the Desc order
func GenerateVolumeReport(stub shim.ChaincodeStubInterface,args []string) peer.Response {
	var pagesize int32 = 20
	bookmark := ""
	var counter uint64
	var pageCounter = 0
	var hasMorePages = true

	// Query string - multi line strings use backward single quote
	query := `{
		"selector": {
		   "docType": "CryptocoinTransactions",
		   "$and": [
			  {
				 "txnDate": {
					"$gte": `
					
	query += "\""+args[0]+"\""
	query += `	 }
			  },
			  {
				 "txnDate": {
					"$lte": `
	query += "\""+args[1]+"\""
	query +=  `}
			  }
		   ],
		   "txnVolume": {
			  "$gte": 0
		   }
		},
		"fields": [
		   "txnVolume",
		   "txnDate"
		],
		"sort": [
		   {
			  "txnVolume": "desc"
		   },
		   {
			  "txnDate": "desc"
		   }
		]
	 }`
	 
	// Print the received query on the console
	fmt.Printf("Query JSON=%s \n", query)

	 // variables to hold query iterator and metadata
	var qryIterator 	shim.StateQueryIteratorInterface
	var queryMetaData 	*peer.QueryResponseMetadata
	var err		error
	resultJSON := ""
	// start the pagination read loop
	lastBookmark := ""
	for hasMorePages {
		// execute the rich query
		qryIterator, queryMetaData, err = stub.GetQueryResultWithPagination(query, pagesize,bookmark)
		if err != nil {
			fmt.Printf("GetQueryResultWithPagination Error=" + err.Error())
			return shim.Error(err.Error())
		}
		var resultKV *queryresult.KV
		// Result read loop only if we received a different bookmark
		
		if lastBookmark != queryMetaData.Bookmark {
			
			for qryIterator.HasNext() {

				// Get the next element
				resultKV, err = qryIterator.Next()
				
				// Increment Counter
				counter++

				resultJSON += string(resultKV.GetValue())+"\n"
			}

						// Increment Page Counter
			pageCounter++

			fmt.Printf("Processed Page: %d \n", pageCounter)
		} 

		
		// Get start key for the next page
		bookmark = queryMetaData.Bookmark

		// boomark = blank indicates no more records
		hasMorePages = (bookmark != "" && lastBookmark != bookmark && bookmark != "nil")
		lastBookmark = bookmark

		// Close the iterator
		qryIterator.Close()
	}

	// Total processed documents
	fmt.Printf("Processed  Documents: %d \n", counter)

	// Return the result JSON
	return shim.Success([]byte(resultJSON))
}