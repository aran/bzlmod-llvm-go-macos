package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/aran/llvm-ar-repro/lib"

	"github.com/aran/llvm-ar-repro/schemas/echo"
)

func main() {
	fmt.Println(lib.Hello())

	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &echo.Server{})

	lis, _ := net.Listen("tcp", ":50051")
	s.Serve(lis)
}
