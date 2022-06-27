package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	hostTarget = map[string]string{
		"app1.domain.com": "http://192.168.1.2:8080",
		"app2.domain.com": "http://192.168.1.2:8888",
	}
	hostProxy = map[string]*httputil.ReverseProxy{}
)

type baseHandle struct{}

func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host

	if fn, ok := hostProxy[host]; ok {
		fn.ServeHTTP(w, r)
		return
	}

	if target, ok := hostTarget[host]; ok {
		remoteUrl, err := url.Parse(target)
		if err != nil {
			log.Println("target parse fail:", err)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
		hostProxy[host] = proxy
		proxy.ServeHTTP(w, r)
		return
	}
	w.Write([]byte("403: Host forbidden " + host))
}

func main() {
	h := &baseHandle{}
	http.Handle("/", h)

	server := &http.Server{
		Addr:    ":8082",
		Handler: h,
	}
	log.Fatal(server.ListenAndServe())
}
