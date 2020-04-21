package main

/**
 * Multiple Testing functions in use
 **/

 import (
	"fmt"

	"testing"
)

//TestDummy shows the various functions available in the Testing.T type
func TestDummy(t *testing.T)  {

	//prints no matter what
	fmt.Println("THIS IS First Message FROM DUMMY !!!")

	// The following will show up in case test case FAILs
	t.Log("This log from dummy")
	// Formatted log message
	t.Logf("This is a fromatted log %s \n","string")

	// Force the test case to fail but call continue
	t.Fail()

	// FailNow() will skip rest of the function call
	// t.FailNow()

	// Checks the status of the test
	t.Logf("Failed flag=%t", t.Failed())

	// Error = Log followed Fail
	t.Error("This is a message from Error - indicates failure")
	// Errorf = Logf + Fail
	t.Errorf("This is a %s message from Error - indicates failure", "formatted")

	// Fatal = Log followed by FailNow.
	// Comment this to run the next statement
	t.Fatal("This is a message from Fatal() call; messages after this will NOT be printed !!")
	// Fatal = Logf followed by FailNow
	t.Fatalf("This is a %s message from Fatal() call; messages after this will NOT be printed !!", "formatted")

	// The message below will NOT be printed because of Fatal()
	fmt.Println("THIS IS Last Message FROM DUMMY !!!")
}
