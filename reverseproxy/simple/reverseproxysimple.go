package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var (
	hosts = []string{
		"http://localhost:8080",
		"http://localhost:8081",
	}

	proxys = []*httputil.ReverseProxy{}
)

type baseHandle struct{}

func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostNum := time.Now().UnixMilli() % int64(len(hosts))
	p := proxys[hostNum]
	fmt.Println("throwing request to", hosts[hostNum])

	p.ServeHTTP(w, r)
}

func main() {
	for _, u := range hosts {
		remoteUrl, err := url.Parse(u)
		if err != nil {
			log.Fatal(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
		proxys = append(proxys, proxy)
	}

	h := &baseHandle{}
	http.Handle("/", h)

	server := &http.Server{
		Addr:    ":8082",
		Handler: h,
	}

	log.Fatal(server.ListenAndServe())
}
