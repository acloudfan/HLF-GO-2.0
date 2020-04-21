#!/bin/bash

. set-env.sh acme



# jq -n --slurpfile arr ./convertcsv.json '$arr' 
# length=$( cat convertcsv.json | jq length)
# arr=$( cat convertcsv.json | jq '.[0]')

# Read the elements in an array
arr=$(cat convertcsv.json | jq -c '.[]')

# Document type for the data
docType="\"CryptocoinTransactions\""

# echo $arr
COUNTER=1
for item in $arr;
do

  txnDate=$(echo $item | jq .date)
  txnVolume=$(echo $item | jq .txVolume)
  # remove decimal part
  txnVolume=${txnVolume%.*}
  generatedCoins=$(echo $item | jq .generatedCoins)
  txnCount=$(echo $item | jq .txCount)
  paymentCount=$(echo $item | jq .paymentCount)
  activeAddresses=$(echo $item | jq .activeAddresses)

  usdPrice=$(echo $item | jq .price)
  # remove decimal part
  usdPrice=${usdPrice%.*}

  echo "$COUNTER  $txnDate"
  COUNTER=$((COUNTER+1))

  args="{\"Args\":[\"AddData\",$docType,$txnDate, \"$txnVolume\",\"$txnCount\",\"$paymentCount\",\"$generatedCoins\",\"$activeAddresses\",\"$usdPrice\"]}"
  set-chain-env.sh -i  "$args"
  chain.sh invoke

  echo $args

done

