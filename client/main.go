package main

import (
	"example/lib"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	fmt.Println("Client started ...")

	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)

	reply := lib.Reply{Message: ""}
	err = client.Call("HelloService.Hello", "Stranger", &reply)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Reply: ", reply)
}
