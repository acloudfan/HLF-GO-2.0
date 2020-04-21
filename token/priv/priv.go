package main

/**
 * Demonstrates the use of Private Data Collections
 * Path:   token/priv
 * Requires the creation of the PDC definition using the pcollection.json 
 **/
import (
	"fmt"

	// April 2020, Updated to Fabric 2.0 Shim
	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"
	
)

// PrivChaincode Represents our chaincode object
type PrivChaincode struct {
}

// Init Implements the Init method
func (privCode *PrivChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed")

	// Return success
	return shim.Success([]byte("true"))
}

// Invoke method
func (privCode *PrivChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the func name and parameters
	funcName, params := stub.GetFunctionAndParameters()

	fmt.Printf("funcName=%s  Params=%s \n", funcName, params)

	if funcName == "Set" {

		return privCode.Set(stub, params)

	} else if funcName == "Get" {

		return privCode.Get(stub)

	} else if funcName == "Del" {

		return privCode.Del(stub, params)

	} else if funcName == "GetFromImplicitCollection" {

		return privCode.GetFromImplicitCollection(stub, params)

	} else if funcName == "MemberOnlyTest" {

		return privCode.Del(stub, params)
	}

	return shim.Error("Invalid Function Name: " + funcName)
}

// Set function
func (privCode *PrivChaincode) Set(stub shim.ChaincodeStubInterface, params []string) peer.Response {

	// Minimum of 2 args is needed - skipping the check for clarity
	// params[0]=Collection name
	// params[1]=Value for the token

	CollectionName := params[0]
	TokenValue := params[1]

	err := stub.PutPrivateData(CollectionName, "token", []byte(TokenValue))
	if err != nil {
		return shim.Error("Error=" + err.Error())
	}

	return shim.Success([]byte("true"))
}

// Get gets the value of token from both collections
func (privCode *PrivChaincode) Get(stub shim.ChaincodeStubInterface) peer.Response {

	// This is returned
	resultString := "{}"

	// Read the open data
	dataOpen, err1 := stub.GetPrivateData("AcmeBudgetOpen", "token")
	if err1 != nil {
		return shim.Error("Error1=" + err1.Error())
	}

	// Read the acme private data
	dataSecret, err2 := stub.GetPrivateData("AcmePrivate", "token")

	accessError := "N.A."
	if err2 != nil {
		//return shim.Error("Error="+err1.Error())
		fmt.Println("Error2=" + err2.Error())
		accessError = err2.Error()
		dataSecret = []byte("**** Not Allowed ***")
	}

	// Returns the token value from 2 PDC + error in retrieving from AcmePrivate collection
	resultString = "{open:\"" + string(dataOpen) + "\", secret:\"" + string(dataSecret) + "\" , error:\"" + accessError + "\"}"

	return shim.Success([]byte(resultString))
}

// Get gets the value of token from both collections
func (privCode *PrivChaincode) GetFromImplicitCollection(stub shim.ChaincodeStubInterface, params []string) peer.Response {
	CollectionName := params[0]

	dataSecret, err := stub.GetPrivateData(CollectionName, "token")
	if err != nil {
		return shim.Error("Error=" + err.Error())
	}

	return shim.Success([]byte(dataSecret))
}

// Exercise solution
// Del gets the value of token from both collections
func (privCode *PrivChaincode) Del(stub shim.ChaincodeStubInterface, params []string) peer.Response {
	// Check for args count MUST be done - not being done here for clarity
	CollectionName := params[0]

	err := stub.DelPrivateData(CollectionName, "token")
	if err != nil {
		return shim.Error("Error=" + err.Error())
	}

	return shim.Success([]byte("true"))
}

// Experimental only :) Do not use
// MemberOnlyTest reads the private data & sets it in a chaincode state key tokenOpen, tokenSecret
func (privCode *PrivChaincode) MemberOnlyTest(stub shim.ChaincodeStubInterface)  peer.Response {
	// Read the open data
	dataOpen, err1 := stub.GetPrivateData("AcmeBudgetOpen", "token")
	if err1 != nil {
		dataOpen = []byte(err1.Error())
	}

	// Read the acme private data
	dataSecret, err2 := stub.GetPrivateData("AcmePrivate", "token")
	if err2 != nil {
		dataSecret = []byte(err2.Error())
	}

	resultString := "{open:\"" + string(dataOpen) + "\", secret:\"" + string(dataSecret) + "\"}"

	return shim.Success([]byte(resultString))
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. priv\n")
	err := shim.Start(new(PrivChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
