package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Math int

func (t *Math) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Math) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	math := new(Math)
	rpc.Register(math)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}

}
