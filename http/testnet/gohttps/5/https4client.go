package main
import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    cmds := os.Args
    if(len(cmds) != 2){
        fmt.Println("use: ./programbin  https://ip:port")
        os.Exit(1)
    }
    fmt.Println("cmdsss:",cmds[0],cmds[1])
    httpaddr := cmds[1]
    pool := x509.NewCertPool()
    caCertPath := "ca.crt"

    caCrt, err := ioutil.ReadFile(caCertPath)
    if err != nil {
        fmt.Println("ReadFile err:", err)
        return
    }
    pool.AppendCertsFromPEM(caCrt)

    cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
    if err != nil {
        fmt.Println("Loadx509keypair err:", err)
        return
    }

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{
            RootCAs:      pool,
            Certificates: []tls.Certificate{cliCrt},
        },
    }
    client := &http.Client{Transport: tr}
    // resp, err := client.Get("https://localhost:8081")
    resp, err := client.Get(httpaddr)
    
    if err != nil {
        fmt.Println("Get error:", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}