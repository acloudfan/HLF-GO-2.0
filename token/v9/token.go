package main
/**
 * token v9
 * Its a modified version of v5 - manages 2 tokens instead of 1
 * Demostrates the use of Endorsement policies & Validation API
 **/
import (
	"fmt"
	// April 2020, Updated for Fabric 2.0
	// Video may have shim package import for Fabric 1.4 - please ignore

	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"
	// KeyEndorsementPolicy interface for create key EP
	// https://godoc.org/github.com/hyperledger/fabric/core/chaincode/shim/ext/statebased#KeyEndorsementPolicy
	// "github.com/hyperledger/fabric/core/chaincode/shim/ext/statebased"
)

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}


// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed")

	// Lets initialize the Key-Value pairs
	stub.PutState("UnProtectedToken", []byte("Inited-UnProtected"))
	stub.PutState("ProtectedToken", []byte("Inited-Protected"))

	// Key policy may be defined as part of initialization
	// BudgetMSP.member is the endorser for the "ProtectedToken"
	ep := "AND('BudgetMSP.peer','AcmeMSP.peer')"
	err := stub.SetStateValidationParameter("ProtectedToken", []byte(ep))

	if err != nil {
		return shim.Error(("err="+err.Error()))
	}

	// Return success
	return shim.Success([]byte("true"))
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	funcName, args := stub.GetFunctionAndParameters()
	fmt.Println("Function=", funcName)


	if(funcName == "set"){
		// Sets the value
		return SetToken(stub, args)

	} else if(funcName == "get"){

		// Gets the value for both tokens
		return GetToken(stub)

	} else if(funcName == "getEP"){

		// Updates the EP for key=MyToken
		return GetEPProtected(stub)

	} else if(funcName == "setEP"){

		// Updates the EP for key=MyToken
		return SetEPProtected(stub, args)
	}
	
	
	// This is not good
	return shim.Error(("Bad Function Name = "+funcName+"!!!"))
}


// SetToken gets the token name and value for token
// args[0] = "unprotected" | "protected" 
// args[1] = Value for the token
func SetToken(stub shim.ChaincodeStubInterface, args []string) peer.Response{
	
	if len(args) < 2 {
		return shim.Error("Failed - incorrect number of parameters!! ")
	}

	// Token name & value
	tokenName := args[0]
	tokenValue:= args[1]

	// Execute PutState - overwrites the current value
	err := stub.PutState(tokenName, []byte(tokenValue))
	if err != nil {
		fmt.Print("Error=", err.Error())
		return shim.Error("Failed - error in PutState!! "+err.Error())
	}

	jsonString := "{ \"Token\":\""+tokenName+"\","
	jsonString += "   \"Value\":\":"+tokenValue+"\"}"

	stub.SetEvent("SetToken", []byte(jsonString))

	return shim.Success([]byte(jsonString))
}

// GetToken reads the value of the tokens 
func  GetToken(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the current value
	valueUnProtectedToken, err1 := stub.GetState("UnProtectedToken")
	valueProtectedToken, err2 := stub.GetState("ProtectedToken")

	// If there is error in retrieve send back an error response
	if(err1 != nil || err2 != nil){
		return  shim.Error(err1.Error()+"**"+err2.Error())
	}

	// Holds a string for the response
	jsonString := "{ \"UnProtected\":\""+string(valueUnProtectedToken)+"\","
	jsonString += "   \"Protected\":\""+string(valueProtectedToken)+"\"}"

	return shim.Success([]byte(jsonString))
}

// SetEPProtected updates the EP for the protected token
// Args => All orgs
// https://godoc.org/github.com/hyperledger/fabric/core/chaincode/shim/ext/statebased#KeyEndorsementPolicy
func SetEPProtected(stub shim.ChaincodeStubInterface,args []string) peer.Response {
	
	ep := "AND("
	for i:=0; i < len(args); i++ {
		if i > 0 {
			ep += ","
		}
		ep += "'"+args[i]+"'"
	}
	ep += ")"

	// Set the EP for the ProtectedToken
	err := stub.SetStateValidationParameter("ProtectedToken", []byte(ep))

	if err != nil {
		return shim.Error(err.Error())
	}

	return  shim.Success([]byte(ep))
}



// GetEPProtected returns the EP for the protected key
func GetEPProtected(stub shim.ChaincodeStubInterface)  peer.Response  {
	 ep, err := stub.GetStateValidationParameter("ProtectedToken")
	 if err != nil {
		 return shim.Error(err.Error())
	 } 

	 
	 return shim.Success(ep)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. token /v9\n")
	err := shim.Start(new(TokenChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

/**
 * SetEPProtected_1 Shows how to use 
 **/
// func SetEPProtected_1(stub shim.ChaincodeStubInterface,args []string) peer.Response {
	
// 	ep, err  := statebased.NewStateEP(nil)
// 	for i:=0; i < len(args); i++ {
// 		ep.AddOrgs(statebased.RoleTypeMember , args[i])
// 	}
	
// 	epBytes,_ := ep.Policy()
// 	err = stub.SetStateValidationParameter("ProtectedToken", epBytes)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	return  shim.Success(epBytes)
// }