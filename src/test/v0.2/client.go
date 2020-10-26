package main

import (
	"time"
	"net"
	"fmt"
)

func main() {
	fmt.Println("client start...")

	conn, err := net.Dial("tcp4", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("Dial err", err);
		return
	}

	time.Sleep(1 * time.Second)

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		fmt.Println("Write err",err)
	}

	buf := make([]byte, 512)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Read err",err)
	}
	
	fmt.Println("server replay: ", buf)
}