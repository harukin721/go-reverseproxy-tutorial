package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	targetURL, err := url.Parse("http://www.google.com")
	if err != nil {
		log.Fatalf("URLの解析に失敗しました: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
