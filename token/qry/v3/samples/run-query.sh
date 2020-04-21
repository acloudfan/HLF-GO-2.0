#!/bin/bash
#Queries the chaincode. 
#Usage:  ./run-query.sh  number 
#        All JSON queries in the files with name  query-<number>.json
#        This script is a utility that reads the query file and replaces the " with \"
#        Sets the query args and then executes query

if [ "$1" == "" ]; then
    echo "Please provide query file number  query-<number>.json e.g.,   ./run-query.sh  1"
    exit
fi

query=$(cat query-$1.json)
# remove the \n
query="$(echo "$query"|tr -d '\n')"
# remove the ' '
query="$(echo "$query"|tr -d ' ')"
# replace " with \"
query=$(echo $query | sed -e 's/\"/\\\"/g')

# execute the generic query function
query={\"Args\":[\"ExecuteRichQuery\",\"$query\"]}

# set the chaincode argument
set-chain-env.sh -q $query


chain.sh query 
