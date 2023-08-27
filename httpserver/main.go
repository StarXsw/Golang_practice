package main

import (
	"fmt"
	"github.com/golang/glog"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 1.接收客户端 request，并将 request 中带的 header 写入 response header
		for name, values := range r.Header {
			for _, value := range values {
				w.Header().Add(name, value)
			}
		}

		// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
		version := os.Getenv("VERSION")
		if version != "" {
			w.Header().Add("VERSION", version)
		}

		// 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
		log.Printf("client IP: %s, HTTP status code: %d, version: %s", r.RemoteAddr, http.StatusOK, version)
		glog.V(2).Info("client IP: %s, HTTP status code: %d, version: %s", r.RemoteAddr, http.StatusOK, version)
		// 4.当访问 localhost/healthz 时，应返回200
		if r.URL.Path == "/healthz" {
			w.WriteHeader(http.StatusOK)
			fmt.Println(w, "200")
			return
		}

		fmt.Println(w, "Hello World!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
