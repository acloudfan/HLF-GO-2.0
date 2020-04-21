How to use?
===========
The queries need to have the " escaped as \"

This makes things complicated to test. The script in this folder will help in simplifying the task for learning the generic query that is built in the sample chaincode.

Create a Query & place it in a file with name      query-NUMBER.json
To execute the query simply execute                ./run-query.sh   NUMBER

For sample queries refer to query-1.json, query-2.json & sample-queries.json

Sample Index File
=================
index-1.json        Index on txnDate & txnVolume
index-2.json        


{
		"selector": {
		   "docType": "CryptocoinTransactions",
		   "$and": [
			  {
				 "txnDate": {
					"$gte": "2009-01-01T00:00:00Z"
                 }
              },
              {
				 "txnDate": {
					"$lte": "2019-02-15T00:00:00Z"
                 }
              }
        }
}