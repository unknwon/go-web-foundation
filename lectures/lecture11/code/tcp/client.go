package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}
	service := os.Args[1]

	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Math.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Math error:", err)
	}
	fmt.Printf("Math: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Math.Divide", args, &quot)
	if err != nil {
		log.Fatal("Math error:", err)
	}
	fmt.Printf("Math: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
