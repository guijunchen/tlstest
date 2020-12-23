package main
import (
    "crypto/tls"
    "crypto/x509"
    "encoding/pem"
    "io/ioutil"
    "net/http"
    "fmt"
)

func InitHttpsClient(keyPem, certPem, pemPass string){
    // 读取私钥文件
    keyBytes, err := ioutil.ReadFile(keyPem)
    if err != nil {
        panic("Unable to read keyPem")
    }
    fmt.Println("keyBytes:",string(keyBytes))
    // 把字节流转成PEM结构
    block, rest := pem.Decode(keyBytes)
    if len(rest) > 0 {
        panic("Unable to decode keyBytes")
    }
    fmt.Println("block:",block)
    // 解密PEM
    der, err := x509.DecryptPEMBlock(block, []byte(pemPass))
    if err != nil {
        panic("Unable to decrypt pem block")
    }
    fmt.Println("der:",der)
    // 解析出其中的RSA 私钥
    key, err := x509.ParsePKCS1PrivateKey(der)
    if err != nil {
        panic("Unable to parse pem block")
    }
    fmt.Println("key:",key)
    // 编码成新的PEM 结构
    keyPEMBlock := pem.EncodeToMemory(
        &pem.Block{
            Type:  "RSA PRIVATE KEY",
            Bytes: x509.MarshalPKCS1PrivateKey(key),
        },
    )
    fmt.Println("keyPEMBlock:",keyPEMBlock)
    // 读取证书文件
    certPEMBlock, err := ioutil.ReadFile(certPem)
    if err != nil {
        panic("Unable to read certPem")
    }
    fmt.Println("certPEMBlock:",certPEMBlock)
    // 生成密钥对
    cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
    if err != nil {
        panic("Unable to read privateKey")
    }
    fmt.Println("cert:",cert)
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                Certificates:       []tls.Certificate{cert},
                InsecureSkipVerify: true,
            },
        },
    }
    _ = client
    //return client
}

func main(){
    InitHttpsClient("csspService.key","csspService.cer","MTIzNDU2")
}

// func InitHttpsClient(keyPem, certPem, pemPass string) *http.Client {
//     // 读取私钥文件
//     keyBytes, err := ioutil.ReadFile(keyPem)
//     if err != nil {
//         panic("Unable to read keyPem")
//     }
//     // 把字节流转成PEM结构
//     block, rest := pem.Decode(keyBytes)
//     if len(rest) > 0 {
//         panic("Unable to decode keyBytes")
//     }
//     // 解密PEM
//     der, err := x509.DecryptPEMBlock(block, []byte(pemPass))
//     if err != nil {
//         panic("Unable to decrypt pem block")
//     }
//     // 解析出其中的RSA 私钥
//     key, err := x509.ParsePKCS1PrivateKey(der)
//     if err != nil {
//         panic("Unable to parse pem block")
//     }
//     // 编码成新的PEM 结构
//     keyPEMBlock := pem.EncodeToMemory(
//         &pem.Block{
//             Type:  "RSA PRIVATE KEY",
//             Bytes: x509.MarshalPKCS1PrivateKey(key),
//         },
//     )
//     // 读取证书文件
//     certPEMBlock, err := ioutil.ReadFile(certPem)
//     if err != nil {
//         panic("Unable to read certPem")
//     }
//     // 生成密钥对
//     cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
//     if err != nil {
//         panic("Unable to read privateKey")
//     }

//     client := &http.Client{
//         Transport: &http.Transport{
//             TLSClientConfig: &tls.Config{
//                 Certificates:       []tls.Certificate{cert},
//                 InsecureSkipVerify: true,
//             },
//         },
//     }
//     return client
// }