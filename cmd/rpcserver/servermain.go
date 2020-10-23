package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"

	common "github.com/gochan-org/go-rpc-example/pkg/rpccommon"
)

var (
	changer      common.Main
	currentColor common.Color
)

func main() {
	var err error
	currentColor = common.Color{Red: 255}

	rpc.RegisterName("changer", &changer)
	rpc.HandleHTTP()
	if err = os.Remove("./rpcsocket"); err == os.ErrNotExist && err != nil {
		panic(err)
	}

	rpcListen, err := net.Listen("unix", "./rpcsocket")
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("Cleaning up")
		rpcListen.Close()
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	go func() {
		fmt.Println("Starting RPC server")
		http.Serve(rpcListen, nil)
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Current color: %#v", currentColor)
	})
	// go func() {
	// fmt.Println("Listening at http://127.0.0.1:8080")
	// 	err := http.ListenAndServe(":8080", nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }()
	<-sc

}
