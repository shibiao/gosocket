package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"

)

var host = flag.String("host", "", "host")
var port = flag.String("port", "8000", "port")

func main() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error listening: ", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + *host + ":" + *port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		go handleRequest(conn)
		go handleReponse(conn)
	}
}
func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		io.Copy(os.Stdout, conn)
	}
}
func handleReponse(conn net.Conn) {
	defer conn.Close()
	//一直发送消息
	//for {
	//	conn.Write([]byte("shibia\n"))
	//}
	conn.Write([]byte("shibiao1"))
}