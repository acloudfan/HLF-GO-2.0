package main

/**
 * v1 shows the use of Range functions
 **/

import (
	// For printing messages on console
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"

	// KV Interface
	"github.com/hyperledger/fabric/protos/ledger/queryresult"

	// JSON Encoding
	"encoding/json"
)

// QueryChaincode Represents our chaincode object
type QueryChaincode struct {
}

// ERC20Token represents a standard token implementation
type ERC20Token struct {
	Symbol      string `json:"symbol"`
	TotalSupply uint64 `json:"totalSupply"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
}

// Init Implements the Init method
func (token *QueryChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed in qry")

	// Add the data
	token.SetupSampleData(stub)

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *QueryChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Invoke executed")

	return token.GetTokenByRange(stub)
}

// GetTokenByRange executes the Range query on state data
func (token *QueryChaincode) GetTokenByRange(stub shim.ChaincodeStubInterface) peer.Response {

	QryIterator, err := stub.GetStateByRange("key1", "key3")
	if err!= nil {
		fmt.Printf("Error="+err.Error())
		return   shim.Error(err.Error())
	}

	var  counter = 0
	var resultJSON = "["
	for QryIterator.HasNext() {
		
		// Hold pointer to the query result
		var result  *queryresult.KV
		var err     error

		// Get the next element
		result, err = QryIterator.Next()

		if err != nil {
			fmt.Println("Err="+err.Error())
		} else {
			// fmt.Println(result.GetKey())
			// fmt.Println(string(result.GetValue()))
			var tokenData string
			tokenData = "{\"key\":\""+result.GetKey()+"\",\"token\":"+string(result.GetValue())+"}"
			if counter > 0 {
				resultJSON += ","
			}
			resultJSON += tokenData
		}
		// Increment counter
		counter++
	}
	resultJSON += "]"

	return shim.Success([]byte(resultJSON))
}

// SetupSampleData creates multiple instances of the ERC20Token
func (token *QueryChaincode) SetupSampleData(stub shim.ChaincodeStubInterface)  {
	erc20 := ERC20Token{Symbol: "ETHV", TotalSupply: 1000, Description: "Ethereum Variant Token", Creator: "John Doe"}
	jsonERC20, _ := json.Marshal(erc20)
	stub.PutState("key1", jsonERC20)

	erc20 = ERC20Token{Symbol: "BTCL", TotalSupply: 10000, Description: "Bitcoin Light Token", Creator: "Jane Doe"}
	jsonERC20, _ = json.Marshal(erc20)
	stub.PutState("key2", jsonERC20)

	erc20 = ERC20Token{Symbol: "XLML", TotalSupply: 5000, Description: "Stellar Like Token", Creator: "Sam Adam"}
	jsonERC20, _ = json.Marshal(erc20)
	stub.PutState("key3", jsonERC20)

	erc20 = ERC20Token{Symbol: "LTCH", TotalSupply: 100000, Description: "Litecoin Heavy Token", Creator: "Anil Kapoor"}
	jsonERC20, _ = json.Marshal(erc20)
	stub.PutState("key4", jsonERC20)

	erc20 = ERC20Token{Symbol: "EOST", TotalSupply: 100, Description: "EOS True Token", Creator: "Rob Cheng"}
	jsonERC20, _ = json.Marshal(erc20)
	stub.PutState("key5", jsonERC20)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. token/qry/v1\n")
	err := shim.Start(new(QueryChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}