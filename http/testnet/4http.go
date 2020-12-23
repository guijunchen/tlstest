package main

import (
	"net"
	"os"
	"bytes"
	"fmt"
	//"io"
)

func main(){

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)


	b := []byte(`{"head":{"code":"1002","sys_id":"UI","app_id":"chengj","ip":"192.1.2.4","time":"20190426152943"},"body":{"sql":"delete from bt_hsm_group where hsm_group_id in('gtest')"}}`)
	
	fmt.Println("len b", len(b))
	fmt.Println("len(b)/256", len(b)/256)
	fmt.Println("len(b)%256", len(b)%256)
	 var h1 byte = byte(len(b)/256)
	 var h2 byte = byte(len(b)%256)
	lenhead := []byte{h1, h2}
	fmt.Println("lenhead", lenhead)
	data :=  append(lenhead, b...)
	fmt.Println("data", data)

	_, err = conn.Write(data)

	checkError(err)
	fmt.Println("11111")
	result, err := readFully(conn)
	checkError(err)
	fmt.Println("string(result)", string(result))
	//fmt.Println(string(result))

	os.Exit(0)
}


func checkError(err error){
	if err != nil{
		fmt.Fprintf(os.Stderr, "fatal error: %s", err.Error)
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error){
	defer conn.Close()
	fmt.Println("readFully")

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	//for {
		
		n, err := conn.Read(buf[0:])
		fmt.Println("buf[2:]", buf[2:])
		
		//fmt.Println("string(result)", string(buf[2:n]))
		result.Write(buf[2:n])
		if err != nil{
			// if err == io.EOF {
			// 	break
			// }
			return result.Bytes(), nil
		}
		//fmt.Println("3333")
	//}
	// fmt.Println("222")
	// fmt.Println("result.Bytes()", result.Bytes())
	return result.Bytes(), nil
}

