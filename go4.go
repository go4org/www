// Copyright 2021 Brad Fitzpatrick. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Serve go4.org

package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const header = `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="go4.org git https://github.com/go4org/go4">
<meta name="go-source" content="go4.org https://github.com/go4org/go4/ https://github.com/go4org/go4/tree/master{/dir} https://github.com/go4org/go4/blob/master{/dir}/{file}#L{line}">
`

const grpcHTML = `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="go4.org/grpc git https://github.com/go4org/grpc">
<meta name="go-source" content="go4.org/grpc https://github.com/go4org/grpc https://github.com/go4org/grpc/tree/master{/dir} https://github.com/go4org/grpc/blob/master{/dir}/{file}#L{line}">
`

const grpcGenHTML = `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="go4.org/grpc-codegen git https://github.com/go4org/grpc-codegen">
<meta name="go-source" content="go4.org/grpc-codegen https://github.com/go4org/grpc-codegen https://github.com/go4org/grpc-codegen/tree/master{/dir} https://github.com/go4org/grpc-codegen/blob/master{/dir}/{file}#L{line}">
`

const memHTML = `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="go4.org/mem git https://github.com/go4org/mem">
<meta name="go-source" content="go4.org/mem https://github.com/go4org/mem https://github.com/go4org/mem/tree/master{/dir} https://github.com/go4org/mem/blob/master{/dir}/{file}#L{line}">
`

const unsafeHTML = `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="go4.org/unsafe/assume-no-moving-gc git https://github.com/go4org/unsafe-assume-no-moving-gc">
`

const internHTML = `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="go4.org/intern git https://github.com/go4org/intern">
`

func serveGo4(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/golang/go-blockchain-support-whitepaper.pdf" {
		http.Redirect(w, req, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusFound)
		return
	}
	if req.URL.Path == "/" {
		if req.TLS == nil {
			http.Redirect(w, req, "https://go4.org/", http.StatusFound)
			return
		}
		io.WriteString(w, header)
		io.WriteString(w, `
</head>
<h1>go4.org</h1>
<p>Misc <a href="https://golang.org/">Go</a> packages.</p>
<ul>
  <li><a href="https://godoc.org/?q=go4.org">Browse</a></li>
  <li><a href="https://github.com/go4org/go4/">About</a></li>
</ul>
<p>
<i>... go4, go four, gopher... get it?</i> Jokes are funnier explained!
</p>
`)
		return
	}

	target := "https://godoc.org/go4.org" + req.URL.Path
	if strings.IndexAny(target, " \t\n\r'\"<>&") != -1 {
		w.WriteHeader(400)
		return
	}

	switch {
	case strings.HasPrefix(req.URL.Path, "/grpc-codegen"):
		io.WriteString(w, grpcGenHTML)
	case strings.HasPrefix(req.URL.Path, "/grpc"):
		io.WriteString(w, grpcHTML)
	case strings.HasPrefix(req.URL.Path, "/mem"):
		io.WriteString(w, memHTML)
	case strings.HasPrefix(req.URL.Path, "/intern"):
		io.WriteString(w, internHTML)
	case strings.HasPrefix(req.URL.Path, "/unsafe"):
		if req.Method == "GET" && req.FormValue("go-get") == "1" {
			io.WriteString(w, unsafeHTML)
		} else if req.URL.Path == "/unsafe/assume-no-moving-gc" {
			io.WriteString(w, "<html><head>\n")
			writeRedirectHTML(w, "https://pkg.go.dev/go4.org/unsafe/assume-no-moving-gc")
			return
		}
	default:
		io.WriteString(w, header)
	}
	writeRedirectHTML(w, target)
}

func serveGo4GRPC(w http.ResponseWriter, req *http.Request) {
	if req.TLS == nil {
		http.Redirect(w, req, "https://grpc.go4.org"+req.URL.Path, http.StatusFound)
		return
	}
	target := "https://godoc.org/grpc.go4.org" + req.URL.Path
	if strings.IndexAny(target, " \t\n\r'\"<>&") != -1 {
		w.WriteHeader(400)
		return
	}
	io.WriteString(w, `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="grpc.go4.org git https://github.com/go4org/grpc">
<meta name="go-source" content="grpc.go4.org https://github.com/go4org/grpc https://github.com/go4org/grpc/tree/master{/dir} https://github.com/go4org/grpc/blob/master{/dir}/{file}#L{line}">
`)
	writeRedirectHTML(w, target)
}

func serveGo4GRPCCodegen(w http.ResponseWriter, req *http.Request) {
	if req.TLS == nil {
		http.Redirect(w, req, "https://grpc-codegen.go4.org"+req.URL.Path, http.StatusFound)
		return
	}
	target := "https://godoc.org/grpc-codegen.go4.org" + req.URL.Path
	if strings.IndexAny(target, " \t\n\r'\"<>&") != -1 {
		w.WriteHeader(400)
		return
	}
	io.WriteString(w, `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="grpc-codegen.go4.org git https://github.com/go4org/grpc-codegen">
<meta name="go-source" content="grpc-codegen.go4.org https://github.com/go4org/grpc-codegen https://github.com/go4org/grpc-codegen/tree/master{/dir} https://github.com/go4org/grpc-codegen/blob/master{/dir}/{file}#L{line}">
`)
	writeRedirectHTML(w, target)
}

func writeRedirectHTML(w io.Writer, target string) {
	fmt.Fprintf(w, `<meta http-equiv="refresh" content="0; url=%s">
</head>
<body>
See <a href="%s">docs</a>.
</body>
</html>
`, target, target)
}
