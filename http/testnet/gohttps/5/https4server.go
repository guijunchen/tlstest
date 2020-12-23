package main

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter,
                   r *http.Request) {
    fmt.Fprintf(w,
        "Hi, This is an example of http service in golang!\n")
}

func main() {
    cmds := os.Args
    if(len(cmds) != 2){
        fmt.Println("use: ./programbin  ip:port")
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

    s := &http.Server{
        Addr:    httpaddr,
        Handler: &myhandler{},
        TLSConfig: &tls.Config{
            ClientCAs:  pool,
            ClientAuth: tls.RequireAndVerifyClientCert,
        },
    }

    err = s.ListenAndServeTLS("server.crt", "server.key")
    if err != nil {
        fmt.Println("ListenAndServeTLS err:", err)
    }
}