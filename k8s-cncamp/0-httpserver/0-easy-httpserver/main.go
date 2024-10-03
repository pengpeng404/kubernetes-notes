package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/golang/glog"
)

func main() {
	// 启用日志级别 4 及以下的日志输出
	err := flag.Set("v", "4")
	if err != nil {
		fmt.Printf("Error setting flag: %v\n", err)
		os.Exit(1)
	}
	glog.V(2).Info("Starting http server...")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)

	err = http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "httpserver is healthy\n")
	if err != nil {
		return
	}
	fmt.Println("httpserver is healthy")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// 只处理根路径 "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r) // 对于非根路径，返回 404
		return
	}
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	if user != "" {
		_, err := io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
		if err != nil {
			return
		}
	} else {
		_, err := io.WriteString(w, "hello [stranger]\n")
		if err != nil {
			return
		}
	}
	_, err := io.WriteString(w, "====Details of the http request header:====\n")
	if err != nil {
		return
	}
	for k, v := range r.Header {
		_, err2 := io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		if err2 != nil {
			return
		}
	}
}
