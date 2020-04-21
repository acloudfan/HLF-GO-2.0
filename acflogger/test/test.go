package main

import (
	"acflogger"
	"os"
)


/** main function to test the acflogger functionality **/
func main() {

	// This simulates the setting of the env var for the Peer
	os.Setenv("CORE_CHAINCODE_LOGGING_LEVEL", "INFO")

	// Create an instance of the logger
	logger := acflogger.NewLogger()

	// Log a message at each of the 6 levels
	logger.Debug("test Debug")
	logger.Info("test Info")
	logger.Notice("test Notice")
	logger.Warning("test Warning")
	logger.Error("test Error")
	logger.Fatal("test Fatal")
}
