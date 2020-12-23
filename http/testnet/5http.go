package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	//"io"
)

func main() {

	//service := "192.1.5.61:6001"
	conn, err := net.Dial("tcp", "192.1.5.61:6001")
	checkError(err)
	b := []byte(`{"head":{"code":"1002","sys_id":"UI","app_id":"chengj","ip":"192.1.2.4","time":"20190426152943"},"body":{"sql":"delete from bt_hsm_group where hsm_group_id in('gtest')"}}`)

	var h1 byte = byte(len(b) / 256)
	var h2 byte = byte(len(b) % 256)
	lenhead := []byte{h1, h2}
	data := append(lenhead, b...)
	_, err = conn.Write(data)
	checkError(err)
	result, err := readFully(conn)
	checkError(err)
	fmt.Println("string(result)", string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %s", err.Error)
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	result.Write(buf[2:n])
	return result.Bytes(), err
}
