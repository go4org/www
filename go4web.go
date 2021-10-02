// Copyright 2021 Brad Fitzpatrick. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

var (
	listen = flag.String("listen", defaultListen(), "port, ip:port to listen on")
)

func defaultListen() string {
	if p := os.Getenv("PORT"); p != "" {
		return ":" + p
	}
	return "localhost:8080"
}

func main() {
	flag.Parse()

	log.Printf("serving go4web on %v...", *listen)
	log.Fatal(http.ListenAndServe(*listen, http.HandlerFunc(mux)))
}

var hostHandlers = map[string]http.Handler{
	"go4.org":              http.HandlerFunc(serveGo4),
	"inet.af":              http.HandlerFunc(serveInetAf),
	"grpc.go4.org":         http.HandlerFunc(serveGo4GRPC),
	"grpc-codegen.go4.org": http.HandlerFunc(serveGo4GRPCCodegen),
}

func mux(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/debug/goroutines" {
		buf := make([]byte, 1<<20)
		buf = buf[:runtime.Stack(buf, true)]
		rw.Header().Set("Content-Type", "text/plain")
		rw.Write(buf)
		return
	}

	host := strings.ToLower(req.Host)
	if v := req.FormValue("behost"); v != "" {
		host = v
	}
	h, ok := hostHandlers[host]
	if !ok {
		h = hostHandlers["go4.org"]
	}
	h.ServeHTTP(rw, req)
}
