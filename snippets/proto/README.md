# Shows the basics of protocol buffers

Read the details @   https://developers.google.com/protocol-buffers/

Try it out

* Install the protocol buffer compiler
network/setup/install-protoc.sh

* Install the package
go get github.com/golang/protobuf/proto


cd $GOPATH/src/token/snippets/proto

1. Generate the go representation of the buffer
$GOPATH/bin/protoc-gen-go   --go_out=*.proto

2. 