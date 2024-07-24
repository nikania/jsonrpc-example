package main

import (
	"example/lib"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(name string, reply *lib.Reply) error {
	(*reply).Message = fmt.Sprintf("Hello, %s!", name)
	return nil
}

func main() {

	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server listening ...")

	rpc.RegisterName("HelloService", new(HelloService))

	for {
		Conn, err := listener.Accept()
		fmt.Println("Connection accepted ...")
		if err != nil {
			fmt.Println(err)
			continue
		}
		go jsonrpc.ServeConn(Conn)

	}
}
