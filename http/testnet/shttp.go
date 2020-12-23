package main

import (
	"net"
	"os"
	"bytes"
	"fmt"
	"io"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	// b := []byte(`POST /E160 HTTP/1.1 \r\n
	// 	sysID: TE \r\n
	// 	appID: TE  \r\n
	// 	Content-Type: application/json \r\n
	// 	User-Agent: PostmanRuntime/7.11.0 \r\n
	// 	Accept: */* \r\n
	// 	Cache-Control: no-cache \r\n
	// 	Postman-Token: 93f71acb-f7db-4ddc-bc7b-9661d04cbf29 \r\n
	// 	Host: 192.1.5.63:22506 \r\n
	// 	accept-encoding: gzip, \r\n
	// 	content-length: 80 \r\n
	// 	Connection: keep-alive \r\n

	// 	{
	// 			"keyName":      "000.000.zek",
	// 			"data": cloudtest3123456cloudtest3123456
	// 	}`)

	// b := []byte(`POST / HTTP/1.0
	// content-length: 80
	// {
	// 	"head": {
	// 		"code": "1002",
	// 		"sys_id":       "UI",
	// 		"app_id":       "chengj",
	// 		"ip":   "192.1.2.4",
	// 		"time": "20190426152943"
	// 	},
	// 	"body": {
	// 		"sql":  "delete from bt_hsm_group where hsm_group_id in('gtest')"
	// 	}
	// }`)

	//b := []byte(`POST / HTTP/1.0\r\ncontent-length:165\r\n\r\n{"head":{"code":"1002","sys_id":"UI","app_id":"chengj","ip":"192.1.2.4","time":"20190426152943"},"body":{"sql":"deletefrombt_hsm_groupwherehsm_group_idin('gtest')"}}`)
	
	//a ：= `{"head":{"code":"1002","sys_id":"UI","app_id":"chengj","ip":"192.1.2.4","time":"20190426152943"},"body":{"sql":"delete from bt_hsm_group where hsm_group_id in('gtest')"}}`

	b := []byte(`{"head":{"code":"1002","sys_id":"UI","app_id":"chengj","ip":"192.1.2.4","time":"20190426152943"},"body":{"sql":"delete from bt_hsm_group where hsm_group_id in('gtest')"}}`)
	fmt.Println("len b", len(b))
	//var msg [2]byte = {len(b)/256,len(b) % 256}
	var h1 byte = len(b)/256
	var h2 byte = len(b) % 256
	mySlice3 := []byte{h1, h2}
	


	// c = "";
	// len(c);
	// d1 = len(c)//256
	// d2 = len(c) % 256;
	// e = d1+d2+c;

	//b := []byte(`HEAD / HTTP/1.0\r\n\r\n`)

	///mySlice1 := make([]byte, 200)
	//var parmas = []byte("POST / HTTP/1.0\r\n content-length: 80 \r\n ")
	// s := make([][]byte, 300)
	// bytes.Join(s,parmas)
	// bytes.Join(s,b)
	//var 
	//mySlice1 = append(mySlice1,b)
	//mySlice1 = append(mySlice1,parmas)

	_, err = conn.Write(b)

	//b := []byte(`HEAD / HTTP/1.0\r\n\r\n`)

	//_, err = conn.Write([]byte("POST / HTTP/1.1\n sysID: TE\n appID: TE\n Connection: keep-alive\n Content-Type: application/json\n User-Agent: PostmanRuntime/7.11.0\n Accept: */*\n Connection: keep-alive\n Cache-Control: no-cache\n Postman-Token: 0c80b7a7-5aad-4eb2-a037-063453ca8060\n Host: 192.1.5.63:22506\n accept-encoding: gzip,\n content-length: 80\n  {\"keyName\": \"000.000.zek\",\"data\": \"cloudtest3123456cloudtest3123456\"}"))
	
	//_, err = conn.Write([]byte("POST / HTTP/1.1\n sysID: TE\n appID: TE\n Connection: keep-alive\n Content-Type: application/json\n User-Agent: PostmanRuntime/7.11.0\n Accept: */*\n Connection: keep-alive\n Cache-Control: no-cache\n Postman-Token: 0c80b7a7-5aad-4eb2-a037-063453ca8060\n Host: 192.1.5.63:22506\n accept-encoding: gzip,\n content-length: 80\n  {\"keyName\": \"000.000.zek\",\"data\": \"********************************\"}"))
	
	//_, err = conn.Write([]byte("POST / HTTP/1.1\n sysID: TE\n appID: TE\n Connection: keep-alive\n Content-Type: application/json\n User-Agent: PostmanRuntime/7.11.0\n Accept: */*\n Connection: keep-alive\n Cache-Control: no-cache\n Postman-Token: 0c80b7a7-5aad-4eb2-a037-063453ca8060\n Host: 192.1.5.63:22506\n accept-encoding: gzip,\n content-length: 80\n  {'head':{'code':'1001','sys_id':'UI','app_id':'chengj','ip':'192.1.2.4','time':'20190426152944'},'body':{'sql':'select * from bt_hsm_group order by hsm_group_id asc ','page_no':'1','page_size':'20'}}"))
	
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))

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

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil{
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

