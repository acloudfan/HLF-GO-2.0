================================================
Chaincode Setup in Dev mode BROKEN on Fabric 2.0
================================================

=======================================
Environment Setup & Data Initialization
=======================================
# 1. Setup the vendor dependencies
cd $GOPATH/src/token/qry/v3
./govendor.sh

# 2. Launch the environment with CouchDB enabled
dev-init.sh  -s

# 3. Deploy the chaincode
. set-env.sh  acme

reset-chain-env.sh

set-chain-env.sh -n CryptocurrencyTxn -v 1.0 -p token/qry/v3 -c '{"Args":[]}' -C airlinechannel
set-chain-env.sh -I false  <<Init not needed>>

chain.sh install -p
chain.sh instantiate

# 4. Upload the data (It may take 15+ minutes)

cd $GOPATH/src/token/qry/v3/data

<<Terminal #1>> Optional 
cd $GOPATH/src/token/qry/v3/data
. set-env.sh  acme
cc-logs.sh -f

<<Termonal #2>>
cd $GOPATH/src/token/qry/v3/data
. set-env.sh  acme

./setup-data.sh

================
ExecuteRichQuery
================
set-chain-env.sh -q '{"Args":["ExecuteRichQuery","{\"selector\":{\"txnDate\": \"2009-12-12T00:00:00Z\"}}"]}'
chain.sh query

# Utility for executing queries:
PS: Remember query attributes limit & skip are ignored by Fabric

qry/v3/samples
./run-query.sh    #NUMBER#      Executes the query against chaincode
./run-query.sh    1.1           Executes the query in file:  query-1.1.json

===============
GetDatesByPrice
===============
- Get the dates on which the price of crypto was $19200 and above
set-chain-env.sh -q '{"Args":["GetDatesByPrice","19200"]}'
chain.sh query

- Get the dates on which the price of crypto was $15000 and above
set-chain-env.sh -q '{"Args":["GetDatesByPrice","15000"]}'
chain.sh query

==================================
GetAveragesBetweenDates (Exercise)
==================================

set-chain-env.sh -q '{"Args":["GetAveragesBetweenDates","2009-01-01T00:00:00Z","2019-02-15T00:00:00Z"]}'

chain.sh query

set-chain-env.sh -q '{"Args":["GetAveragesBetweenDates","2017-12-01T00:00:00Z","2018-06-31T00:00:00Z"]}'
chain.sh query

set-chain-env.sh -q '{"Args":["GetAveragesBetweenDates","2017-12-01T00:00:00Z","2018-01-31T00:00:00Z"]}'
chain.sh query

=====================================================
Chaincode setup in net mode - For Index related tests
=====================================================

NOTE: Do this ONLY for resetting the environment 

#1 Start the env in net mode
dev-init.sh -s
#2 Setup org context
. set-env.sh  acme
#3 Setup the chain env
set-chain-env.sh -n CryptocurrencyTxn -v 1.0 -p token/qry/v3 -c '{"Args":[]}' -C airlinechannel

#4 Install & Instantiate
chain.sh install -p
set-chain-env.sh -s Seq#   <<if needed>>
chain.sh instantiate

#5 Setup the data
Run the script under token/qry/v3/data
. set-env.sh acme
./setup-data.sh

=========================
Index Performance Testing
=========================
======
Part-1  Execute the function GetPriceByDate without index to measure the time
======
- Setup query arguments
. set-env.sh acme
set-chain-env.sh -q '{"Args":["GetDatesByPrice","19000"]}'
- Time the execution
time chain.sh query

PS: Note down the real execution time
======
Part-2  Install the price index
======
- Copy the index file samples/index-OnUsdPrice.json to META-INF/statedb/couchdb/indexes
- Upgrade the chaincode
chain.sh  upgrade-auto

Confirm on Futon that index is installed
======
Part-3  Examine the performance gain
======
- Setup query arguments
set-chain-env.sh -q '{"Args":["GetDatesByPrice","15000"]}'
- Time the execution
time chain.sh query

PS: You should see marked difference in performance


Testing the Sorting
===================
1. Test the function    GenerateVolumeReport

set-chain-env.sh -q '{"Args":["GenerateVolumeReport","2019-02-01T00:00:00Z","2019-02-07T00:00:00Z"]}'
chain.sh query

PS: This will fail as the sort requires the index
[It will succeed if the index file index-OnTxnDateVolume.json exist under the folder v3/META-INF/statedb/couchdb/indexes/]

2. Package the index JSON and upgrade

- Copy the index file samples/index-OnTxnDateVolume.json to META-INF/statedb/couchdb/indexes

3. Test again - this time it would work

set-chain-env.sh -q '{"Args":["GenerateVolumeReport","2019-02-01T00:00:00Z","2019-02-07T00:00:00Z"]}'
chain.sh query

Queries/Selectors
=================
http://docs.couchdb.org/en/latest/api/database/find.html#find-selectors
Query => A JSON object with standard "attribute" names for specifying the query criteria
"selector": Object specifying the document selection criteria
            Operators:
            Selectors may use the operators for defining the criteria
            $eq     Equal
            $gt     Greater than
            $lt     Less than

"fields":   An array of field names that need to be returned
"sort":     An array of field names for sorting result set


Data on Bitcoin
===============
#1 Downloaded the data from this site in csv format
https://coinmetrics.io/data-downloads/

#2 Remove the (USD) from headers - Replaced (USD)  with ""
Open in VSC editor and just replace

#3 Converted the data from csv to json format
http://www.convertcsv.com/csv-to-json.htm



CouchDB Indexes
===============
http://docs.couchdb.org/en/stable/api/database/find.html#db-index
