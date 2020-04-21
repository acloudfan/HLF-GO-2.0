package main

/**
 * Shows how to use the history
 **/

import (
	// For printing messages on console
	"fmt"


	// April 2020, Updated to Fabric 2.0 Shim
	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"

	// KV Interface
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"

	// JSON Encoding
	"encoding/json"

	"strconv"
)

// VehicleChaincode Represents our chaincode object
type VehicleChaincode struct {
}

// Vehicle Represents our car asset
type Vehicle struct {
	DocType			string  `json:"docType"`
	VIN				string  `json:"vin"`
	Year 			uint    `json:"year"`
	Make			string  `json:"make"`
	Model			string  `json:"model"`
	Owner			string  `json:"owner"`
	Transfer		string  `json:"transfer"`	
}
// DocType Represents the object type
const	DocType	= "VehicleAsset"

// Init Implements the Init method
func (history *VehicleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed in history")

	// Setup the sample data
	history.SetupSampleData(stub)

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (history *VehicleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

		// Get the function name and parameters
		funcName, args := stub.GetFunctionAndParameters()

		if funcName == "GetVehicleByVin" {
			// Returns the vehicle's current state
			return history.GetVehicleByVin(stub, args)

		} else if funcName == "TransferOwnership" {
			// Invoke this function to transfer ownership of vehicle
			return history.TransferOwnership(stub, args)

		} else if funcName == "GetVehicleHistory" {
			// Query this function to get txn history for specific vehicle
			return history.GetVehicleHistory(stub, args)

		} else if funcName == "GetVehiclesByYear" {
			
			// Get all vehicle by year - just another example of query
			return history.GetVehiclesByYear(stub, args)

		} else if funcName == "GetVehiclesOwners" {
			// To be implemented in the exercise
			// return history.GetVehiclesOwners(stub, args)
		} 

		// This is not good
		return shim.Error(("Bad Function Name = !!!"))
}

// GetVehicleHistory gets the history of the vehicle by VIN
// args[0] = VIN
func (history *VehicleChaincode) GetVehicleHistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	// Check the number of args
	if len(args) < 1 {
		return shim.Error("MUST provide VIN !!!")
	}

	// Get the history for the key i.e., VIN#
	historyQueryIterator, err := stub.GetHistoryForKey(args[0])
	
	// In case of error - return error
	if err != nil {
		return shim.Error("Error in fetching history !!!"+err.Error())
	}

	// Local variable to hold the history record
	var resultModification *queryresult.KeyModification
	counter := 0
	resultJSON := "["

	// Start a loop with check for more rows
	for historyQueryIterator.HasNext() {

		// Get the next record
		resultModification, err = historyQueryIterator.Next()

		if err != nil {
			return shim.Error("Error in reading history record!!!"+err.Error())
		}
		
		// Append the data to local variable
		data :="{\"txn\":" + resultModification.GetTxId()
		data +=" , \"value\": "+ string(resultModification.GetValue()) + "}  "
		if counter > 0 {
			data = ", "+data
		}
		resultJSON += data

		counter++
	}

	// Close the iterator
	historyQueryIterator.Close()

	// finalize the return string
	resultJSON += "]"
	resultJSON = "{ \"counter\": " + strconv.Itoa(counter) + ", \"txns\":" + resultJSON  + "}"

	// return success
	return shim.Success([]byte(resultJSON))
}

// TransferOwnership gets the asset information
// Transfer the ownership of the vehicle from owner1 to owner2
// args[0]=vin   args[1]=current owner  args[2]=new owner args[3]=transfer date
// args[1] used for validation
func (history *VehicleChaincode) TransferOwnership(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 4 {
		return shim.Error("MUST provide VIN, Current owner, New owner, Transfer date !!!")
	}
	// get the state information
	bytes, _ := stub.GetState(args[0])
	if bytes == nil {
		return shim.Error("Provided VIN not found!!!")
	}
	// unmarshall the data
	// Read the JSON and Initialize the struct
	var vehicle  Vehicle
	_ = json.Unmarshal(bytes, &vehicle)

	// Business rule - the current owner MUST match
	if vehicle.Owner != args[1]{
		return shim.Error("Current owner MUST match !!!")
	}

	// check should be made if string is NOT blank
	vehicle.Owner=args[2]
	vehicle.Transfer=args[3]
	jsonVehicle, _ := json.Marshal(vehicle)

	stub.PutState(vehicle.VIN, jsonVehicle)

	return shim.Success([]byte("Vehicle Record Updated!!! "+ string(jsonVehicle)))
}

// GetVehicleByVin gets the asset information
func (history *VehicleChaincode) GetVehicleByVin(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// Check on args
	if len(args) < 1 {
		return shim.Error("MUST provide Vin Number in args[0] !!!")
	}

	// Get the data
	vehicle, _ := stub.GetState(args[0])

	return shim.Success([]byte(vehicle))
}

// GetVehiclesByYear gets all the vehicles >= year
// Another sample to show the use of Rich Queries 
// To make this work you need to create an index :)
// Not using Pagination - so results restricted to a max of totalQueryLimit
func (history *VehicleChaincode) GetVehiclesByYear(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args[0]) < 1 {
		return shim.Error("Please provide a valid year !!!")
	}
	qry := `{
		"selector": {
		   "year": {
			  "$gte": `

	qry += args[0]
	qry += `		  
		   }
		},
		"sort": [{"year": "desc"}]
	 }`

	// GetQueryResult
	QryIterator, err := stub.GetQueryResult(qry)
	if err != nil {
		return shim.Error("Error in executing rich query !!!! "+err.Error())
	}
	// hold the result json
	resultJSON := "["
	counter := 0
	for QryIterator.HasNext() {
		// Hold pointer to the query result
		var resultKV *queryresult.KV

		// Get the next element
		resultKV, _ = QryIterator.Next()

		value := string(resultKV.GetValue())
		if counter > 0 {
			resultJSON += ", "
		}
		resultJSON += value
		counter++
	}
	resultJSON += "]"

	return shim.Success([]byte(resultJSON))
}


// SetupSampleData creates multiple instances of the ERC20history
func (history *VehicleChaincode) SetupSampleData(stub shim.ChaincodeStubInterface) {
	
	// This the car data for testing
	AddData(stub, "100","toyota","corolla",2011,"J Smith","2015-12-20")
	AddData(stub,"200","honda","civic",2012,"G Roger","2016-01-15")
	AddData(stub,"300","audi","a5",2015,"S Ripple","2018-07-22")
	AddData(stub,"400","bmw","x5",2013,"M Jane","2019-02-19")
	AddData(stub,"500","toyota","camry",2018,"J Hoover","2019-01-15")

	fmt.Println("Initialized with the sample data!!")
}

//AddData adds a car asset to the chaincode asset database
//Structure is created and initialized then it is marshalled to JSON for storage using PutState
func AddData(stub shim.ChaincodeStubInterface,vin string, make, model string, year uint, owner, transfer string) {
	vehicle := Vehicle{DocType: DocType, VIN: vin, Year: year, Make: make, Model: model, Owner: owner, Transfer: transfer}
	jsonVehicle, _ := json.Marshal(vehicle)
	// Key = VIN#, Value = Car's JSON representation
	stub.PutState(vin, jsonVehicle)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. token/history\n")
	err := shim.Start(new(VehicleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}