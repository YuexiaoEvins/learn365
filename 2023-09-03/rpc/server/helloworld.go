package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloWorldService struct{}

const HelloWorldServiceName = "HelloWorldService"

func (s *HelloWorldService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}

func main() {
	_ = rpc.RegisterName(HelloWorldServiceName, new(HelloWorldService))

	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("listen :12345 failed:", err)
		panic(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept connections failed:", err)
	}

	rpc.ServeConn(conn)
}
