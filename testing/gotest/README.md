# Demonstrates the use of the 'go test' command

cd $GOPATH/src/testing/gotest

1. Execute the following to see all tests getting executed

go test

go test -v      The above command does not show the test that PASS
                With -v you will see the tests that passed

go test -cover  Shows the coverage analysis

2. Make changes to target_test.go & dummy_test.go to understand the behavior