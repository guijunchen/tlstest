package main

import (
	"net/http"
	"net/url"
)

func main() {
    resp, err := http.PostForm("http://192.1.5.63:22506/E160",url.Values{"keyName": {"cloud01.test3.zek"}, "data": {"cloudtest3123456cloudtest3123456"}})
		
	//resp, err := http.PostForm("http://example.com/posts", url.Values{"title":{"article title"}, "content": {"article body"}})
	// if err != nil {
	// // 处理错误
	// return
	// }
    if err != nil {
        // handle error
    }

    defer resp.Body.Close()
    //body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    //fmt.Println(string(body))

}