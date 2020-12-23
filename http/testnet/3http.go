package main

import (
    "fmt"
    "net/http"
)

func main() {
    resp, err := http.Get("http://192.1.5.86:8082/login?redirect=%2Flogin")
    if err != nil {
        // print err
        return
    } 
    
    defer resp.Body.Close()
    
    // print resp.Status, resp.StatusCode, resp.Header, resp.Body
    
    // body 需要 io 的方式读取
    //var tmp string
    buf := make([]byte, 4 * 1024)
    for {
        n, _ := resp.Body.Read(buf)
        if n == 0 {
            // print err
            break
        }
        fmt.Println(n)
        //tmp += n
    }
    fmt.Println(buf)
    // print tmp 网页内容
}