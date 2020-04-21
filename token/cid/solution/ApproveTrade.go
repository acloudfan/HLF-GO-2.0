package main

/**
 * Scenario is an example of how programatic access control works
 * Business Rule:
 * 1. Caller MUST be from the accounting department
 * 2. If tradeValue < 100K it may be approved by anyone from accounting dept
 * 3. If tradeValue >= 100K it MUST be approved by someone from accounting dept with a role = manager
 **/
 import (

	// April 2020, Updated to Fabric 2.0 Shim
	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"

	// Client Identity Library
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"

	"strconv"
 )

 // ApproveTrade - checks the amount of trade passed in args[0] - applies the business rules
 // Rule#1  Caller MUST be from accounting dept
 // Rule#2 If trade < 100K it can be approved by anyone from accounting
 // Rule#3 If trade >= 100K the caller MUST have a role = manager
 func (clientdid *CidChaincode) ApproveTrade(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	if len(args) < 1 {
		return shim.Error("Must provide Trade Value in args[0] !!!")
	}

	enrollID, _, _ := cid.GetAttributeValue(stub, "hf.EnrollmentID")

	// Rule#1  Caller MUST be from accounting dept
	attrValue, flag, _ := cid.GetAttributeValue(stub, "department")
	if !flag || attrValue != "accounting" {
		return shim.Error("REJECTED - Caller MUST be from accounting department to Approve the trade !!!")
	}

	// Convert trade value to number - ignoring error
	tradeValue, _ := strconv.ParseUint(string(args[0]),10,64)

	// Rule#2 If trade < 100K it can be approved by anyone from accounting
	if tradeValue < 100000 {

		return clientdid.ProcessTheTrade(stub, args[0], enrollID)
	}

	// Rule#3 If trade >= 100K the caller MUST have a role = manager
	attrValue, flag, _ = cid.GetAttributeValue(stub, "app.accounting.role")

	if !flag || attrValue != "manager" {
		return shim.Error("REJECTED - Caller has role='"+attrValue+"' but since Tradevalue="+args[0]+" it requires role='manager'")
	}

	// All rules fulfilled
	return clientdid.ProcessTheTrade(stub, args[0], enrollID)
 }

 // ProcessTheTrade - dummy function - in real sceanrio the state will change for asset in the trade
 func (clientdid *CidChaincode) ProcessTheTrade(stub shim.ChaincodeStubInterface, tradeValue, enrollID string) peer.Response {

	// Result string sent to caller as part of the endorsement
	result := "APPROVED - Trade value="+tradeValue+" by "+enrollID

	return shim.Success([]byte(result))
 }