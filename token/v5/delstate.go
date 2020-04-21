package		main

// Solution for the exercise
import (
	"fmt"

	// April 2020, Updated for Fabric 2.0
	// Video may have shim package import for Fabric 1.4 - please ignore

	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"

)

// DeleteToken deletes the token from the database
// V5
// Returns 0 if successful
func  DeleteToken(stub shim.ChaincodeStubInterface) peer.Response {

	// Check if the key exists - if not then return false
	value, _ := stub.GetState("MyToken")
	if value == nil {
		return shim.Success([]byte("false"))
	}
	// Delete the key
	if err := stub.DelState("MyToken"); err != nil {
		fmt.Println("Delete Failed!!! ", err.Error())
		return shim.Error(("Delete Failed!! "+err.Error()+"!!!"))
	}

	return shim.Success([]byte("true"))
}