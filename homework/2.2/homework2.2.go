package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")

}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

}

func os_env_GOPATH(w http.ResponseWriter, req *http.Request) {
	//  con := os.Getenv("GOPATH")
	fmt.Fprintf(w, "OS_ENV:%s\n", os.Getenv("GOPATH"))

}

func get_client_info(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RemoteAddr)
	//	fmt.Println(req.Response.StatusCode)
}

func health(w http.ResponseWriter, req *http.Request) {
	//设置 http请求状态
	w.WriteHeader(http.StatusOK)
	//写入页面数据
	w.Write([]byte("200"))
}

func main() {
	http.HandleFunc("/", get_client_info)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/version", os_env_GOPATH)
	http.HandleFunc("/healthz", health)
	http.ListenAndServe(":80", nil)

}

// Header contains the request header fields either received
// by the server or to be sent by the client.
//
// If a server received a request with header lines,
//
//	Host: example.com
//	accept-encoding: gzip, deflate
//	Accept-Language: en-us
//	fOO: Bar
//	foo: two
//
// then
//
//	Header = map[string][]string{
//		"Accept-Encoding": {"gzip, deflate"},
//		"Accept-Language": {"en-us"},
//		"Foo": {"Bar", "two"},
//	}
//
// For incoming requests, the Host header is promoted to the
// Request.Host field and removed from the Header map.
//
// HTTP defines that header names are case-insensitive. The
// request parser implements this by using CanonicalHeaderKey,
// making the first character and any characters following a
// hyphen uppercase and the rest lowercase.
//
// For client requests, certain headers such as Content-Length
// and Connection are automatically written when needed and
// values in Header may be ignored. See the documentation
// for the Request.Write method.
