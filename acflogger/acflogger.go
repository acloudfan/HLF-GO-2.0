package acflogger

import (
	"fmt"
	"os"
	"strings"
)

//DEBUG Logging level
const DEBUG = 0
//INFO Logging level
const INFO = 1
//NOTICE Logging Level
const NOTICE = 2
//WARNING Logging Level
const WARNING = 3
//ERROR Logging Level
const ERROR = 4
//FATAL logging Level
const FATAL = 5

//AcfLogger structure
type AcfLogger struct {
	logLevel   int
}

//NewLogger creates an instance 
func NewLogger() (*AcfLogger) {

	// Read the Log level string
	levelString := os.Getenv("CORE_CHAINCODE_LOGGING_LEVEL")
	levelString = strings.ToUpper(levelString)
	
	// Setup the log level as an integer
	level := 0
	switch levelString{
		case "DEBU": level=DEBUG; break;
		case "DEBUG": level=DEBUG; break;

		case "INFO": level=INFO; break;

		case "NOTICE": level=NOTICE; break;

		case "WARN": level=WARNING; break;
		case "WARNING": level=WARNING; break;

		case "ERROR": level=ERROR; break;
		
		case "FATAL": level=FATAL; break;
		default: level = DEBUG
	}
	x := AcfLogger{level}

	fmt.Printf("AcfLogger: Logging level=%s %d\n", levelString, level)

	// Return the instance of logger to the caller
	return &x
}

//Debug message printing 
func (logger *AcfLogger) Debug(s string) {
	if logger.logLevel <= DEBUG {
		fmt.Printf("DEBUG: %s\n", s)
	}
}

//Info message printing
func (logger *AcfLogger) Info(s string) {
	if logger.logLevel <= INFO {
		fmt.Printf("INFO: %s\n", s)
	}
}

//Notice message printing
func (logger *AcfLogger) Notice(s string) {
	if logger.logLevel <= NOTICE {
		fmt.Printf("NOTICE: %s\n", s)
	}
}

//Warning message printing
func (logger *AcfLogger) Warning(s string) {
	if logger.logLevel <= WARNING {
		fmt.Printf("WARNING: %s\n", s)
	}
}

//Error message printing
func (logger *AcfLogger) Error(s string) {
	if logger.logLevel <= ERROR {
		fmt.Printf("ERROR: %s\n", s)
	}
}

//Fatal message printing
func (logger *AcfLogger) Fatal(s string) {
	if logger.logLevel <= FATAL {
		fmt.Printf("FATAL: %s\n", s)
	}
}

