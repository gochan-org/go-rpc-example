package main

import (
	"net/rpc"

	common "github.com/gochan-org/go-rpc-example/pkg/rpccommon"
)

func main() {
	client, err := rpc.DialHTTP("unix", "./rpcsocket")
	if err != nil {
		panic(err)
	}
	from := common.Color{}
	to := common.Color{Red: 32, Green: 32, Blue: 32}
	if err = client.Call("changer.ChangeColor", []common.Color{from, to}, nil); err != nil {
		panic(err)
	}

	if err = client.Call("changer.HelloWorld", new(int), new(int)); err != nil {
		panic(err)
	}
}
