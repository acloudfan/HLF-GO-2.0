package main

/**
 * tokenv2
 * Shows the
 *    A) Use of Logger
 **/
import (
	"fmt"

	// April 2020, Updated for Fabric 2.0
	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"

	// acloudFan custom Logger
	"acflogger"
)

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}

// V2

// ChaincodeName - Create an instance of the Logger
const ChaincodeName = "tokenv2"

// var logger = shim.NewLogger(ChaincodeName) // v1.4x not supported in 2.x
var logger = acflogger.NewLogger()

// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	// fmt.Println("Init executed")

	logger.Debug("Init executed v2 - DEBUG")

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	//log a warning message
	logger.Debug("Invoke executed v2   logger.Debug()")
	//log a info message
	logger.Info("Invoke executed v2    logger.Info()")
	//logger notice message
	logger.Notice("Invoke executed v2   logger.Notice()")
	//log a warning message
	logger.Warning("Invoke executed v2  logger.Warning()")
	//log an error message
	logger.Error("Invoke executed v2  logger.Error()")
	//log a fatal message
	logger.Fatal("Invoke executed v2  logger.Fatal()")

	return shim.Success(nil)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. token/v2 \n\n")
	err := shim.Start(new(TokenChaincode))
	if err != nil {
		//fmt.Printf("Error starting chaincode: %s", err)
		// V2   https://godoc.org/builtin#error
		logger.Error("Error starting chaincode: " + err.Error())
	}
}
