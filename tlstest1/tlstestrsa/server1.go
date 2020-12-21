package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)

func main() {
	//cert1 := "../asserts/rsacert/server.pem"
	//key1 := "../asserts/rsacert/server.key"
	cert2 := "../asserts/rsacert2/server.crt"
	key2 := "../asserts/rsacert2/server.key"

	cert, err := tls.LoadX509KeyPair(cert2, key2)
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", "127.0.0.1:5443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
