/**
 * This implements the ERC20 standard token
 * https://theethereum.wiki/w/index.php/ERC20_Token_Standard
 * The ERC20 standard is used for defining tokens in Ethereum. This is an implementation
 * of the same standard on Fabric.
 **/
 package main

 import (
	"fmt"

	// April 2020, Updated to Fabric 2.0 Shim
	"github.com/hyperledger/fabric-chaincode-go/shim"

	peer "github.com/hyperledger/fabric-protos-go/peer"

	// Conversion functions
	"strconv"

	// JSON Encoding
	"encoding/json"
)

// ERC20TokenChaincode Represents our chaincode object
type ERC20TokenChaincode struct {
}

// ERC20Token structure manages the state 
type ERC20Token struct {
	Symbol   		string   `json:"symbol"`
	TotalSupply     uint64   `json:"totalSupply"`
	Description		string   `json:"description"`
	Creator			string   `json:"creator"`
}

// OwnerPrefix is used for creating the key for balances
const   OwnerPrefix="owner."

// Init Implements the Init method
// Receives 4 parameters =  [0] Symbol [1] TotalSupply   [2] Description  [3] Owner
func (token *ERC20TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed")
	_, args := stub.GetFunctionAndParameters()

	// Check if we received the right number of arguments
	if len(args) < 4 {
		return shim.Error("Failed - incorrect number of parameters!! ")
	}
	symbol := string(args[0])
	// Get total supply & check if it is > 0
	totalSupply, err := strconv.ParseUint(string(args[1]),10,64)

	if err != nil || totalSupply == 0 {
		return shim.Error("Total Supply MUST be a number > 0 !! ")
	}

	// Creator name cannot be zeo length
	if len(args[3]) == 0 {
		return errorResponse("Creator identity cannot be 0 length!!!", 3)
	}
	creator := string(args[3])

	// Create an instance of the token struct
	var erc20 = ERC20Token{Symbol: symbol, TotalSupply: totalSupply, Description: string(args[2]), Creator: creator}

	// Convert to JSON and store token in the state
	jsonERC20, _ := json.Marshal(erc20)
	stub.PutState("token", []byte(jsonERC20))

	// Maintain the balances in the state db
	// In the begining all tokens are owned by the creator of the token
	key := OwnerPrefix+creator
	fmt.Println("Key=",key)
	err=stub.PutState(key,[]byte(args[1]))
	if err != nil {
		return errorResponse(err.Error(), 4)
	}

	return shim.Success([]byte(jsonERC20))
}

// Invoke method
func (token *ERC20TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the function name and parameters
	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Invoke executed : ", function, " args=", args)

	switch {

	// Query function
	case	function == "totalSupply":
			return totalSupply(stub)
	case	function == "balanceOf":
			return balanceOf(stub, args)
	case	function == "transfer":
			return transfer(stub, args)
	}

	return errorResponse("Invalid function",1)
}

/**
 * Getter function 
 * function totalSupply() public view returns (uint);
 * Returns the totalSupply for the token
 **/
func totalSupply(stub shim.ChaincodeStubInterface) peer.Response {

	bytes, err := stub.GetState("token")
	if err != nil {
		return errorResponse(err.Error(), 5)
	}

	// Read the JSON and Initialize the struct
	var erc20  ERC20Token
	_ = json.Unmarshal(bytes, &erc20)

	
	// Create the JSON Response with totalSupply
	return successResponse(strconv.FormatUint(erc20.TotalSupply,10))
}

/**
 * Getter function
 * function balanceOf(address tokenOwner) public view returns (uint balance);
 * Returns the balance for the specified owner
 * {"Args":["balanceOf","owner-id"]}
 **/
 func balanceOf(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// Check if owner id is in the arguments
	if len(args) < 1   {
		return errorResponse("Needs OwnerID!!!", 6)
	}
	OwnerID := args[0]
	bytes, err := stub.GetState(OwnerPrefix+OwnerID)
	if err != nil {
		return errorResponse(err.Error(), 7)
	}

	response := balanceJSON(OwnerID, string(bytes))

	return successResponse(response)
 }

 /**
  * Setter function
  * function transfer(address to, uint tokens) public returns (bool success);
  * Transfer tokens 
  * {"Args":["from","to","amount"]}
  **/
  func transfer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// Check if owner id is in the arguments
	if len(args) < 3   {
		return errorResponse("Needs to, from & amount!!!", 700)
	}

	from := string(args[0])
	to := string(args[1])
	amount, err := strconv.Atoi(string(args[2]))
	if err != nil {
		return errorResponse(err.Error(), 701)
	}
	if(amount <= 0){
		return errorResponse("Amount MUST be > 0!!!", 702)
	}

	// Get the Balance for from
	bytes, _ := stub.GetState(OwnerPrefix+from)
	if len(bytes) == 0 {
		// That means 0 token balance
		return errorResponse("Balance MUST be > 0!!!", 703)
	}
	fromBalance, _ := strconv.Atoi(string(bytes))
	if fromBalance < amount {
		return errorResponse("Insufficient balance to cover transfer!!!", 704)
	}
	// Reduce the tokens in from account
	fromBalance = fromBalance - amount
	
	// Get the balance in to account
	bytes, _ = stub.GetState(OwnerPrefix+to)
	toBalance := 0
	if len(bytes) > 0 {
		toBalance, _ = strconv.Atoi(string(bytes))
	}
	toBalance += amount

	// Update the balance
	bytes = []byte(strconv.FormatInt(int64(fromBalance), 10))
	err = stub.PutState(OwnerPrefix+from, bytes)

	bytes = []byte(strconv.FormatInt(int64(toBalance), 10))
	err = stub.PutState(OwnerPrefix+to, bytes)

	// Emit Transfer Event
	eventPayload := "{\"from\":\""+from+"\", \"to\":\""+to+"\",\"amount\":"+strconv.FormatInt(int64(amount),10)+"}"
	stub.SetEvent("transfer", []byte(eventPayload))
	return successResponse("Transfer Successful!!!")
  }

 // balanceJSON creates a JSON for representing the balance
 func balanceJSON(OwnerID, balance string) (string) {
	 return "{\"owner\":\""+OwnerID+"\", \"balance\":"+balance+ "}"
 }


func errorResponse(err string, code  uint ) peer.Response {
	codeStr := strconv.FormatUint(uint64(code), 10)
	// errorString := "{\"error\": \"" + err +"\", \"code\":"+codeStr+" \" }"
	errorString := "{\"error\":" + err +", \"code\":"+codeStr+" \" }"
	return shim.Error(errorString)
}

func successResponse(dat string) peer.Response {
	success := "{\"response\": " + dat +", \"code\": 0 }"
	return shim.Success([]byte(success))
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Println("Started....")
	err := shim.Start(new(ERC20TokenChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
