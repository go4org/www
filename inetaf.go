// Copyright 2021 Brad Fitzpatrick. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Serve inet.af

package main

import (
	"io"
	"net/http"
	"strings"
)

func serveInetAf(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if r.Method != "GET" && r.Method != "HEAD" {
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}
	isGoGet := strings.Contains(r.RequestURI, "go-get=1") && r.URL.Query().Get("go-get") == "1"

	if isGoGet {
		if strings.HasPrefix(path, "/tcpproxy") {
			io.WriteString(w, `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="inet.af/tcpproxy git https://github.com/inetaf/tcpproxy">
<meta name="go-source" content="inet.af/tcpproxy https://github.com/inetaf/tcpproxy/ https://github.com/inetaf/tcpproxy/tree/master{/dir} https://github.com/inetaf/tcpproxy/blob/master{/dir}/{file}#L{line}">
</head>
</html>
`)
			return
		}
		if strings.HasPrefix(path, "/netaddr") {
			io.WriteString(w, `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="inet.af/netaddr git https://github.com/inetaf/netaddr">
<meta name="go-source" content="inet.af/netaddr https://github.com/inetaf/netaddr/ https://github.com/inetaf/netaddr/tree/main{/dir} https://github.com/inetaf/netaddr/blob/main{/dir}/{file}#L{line}">
</head>
</html>
`)
			return
		}
		if strings.HasPrefix(path, "/peercred") {
			io.WriteString(w, `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="inet.af/peercred git https://github.com/inetaf/peercred">
<meta name="go-source" content="inet.af/peercred https://github.com/inetaf/peercred/ https://github.com/inetaf/peercred/tree/main{/dir} https://github.com/inetaf/peercred/blob/main{/dir}/{file}#L{line}">
</head>
</html>
`)
			return
		}
		if strings.HasPrefix(path, "/wf") {
			io.WriteString(w, `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="inet.af/wf git https://github.com/inetaf/wf">
<meta name="go-source" content="inet.af/wf https://github.com/inetaf/wf/ https://github.com/inetaf/wf/tree/main{/dir} https://github.com/inetaf/wf/blob/main{/dir}/{file}#L{line}">
</head>
</html>
`)
			return
		}
		if strings.HasPrefix(path, "/netstack") {
			io.WriteString(w, `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="inet.af/netstack git https://github.com/inetaf/netstack">
<meta name="go-source" content="inet.af/netstack https://github.com/inetaf/netstack/ https://github.com/inetaf/netstack/tree/main{/dir} https://github.com/inetaf/netstack/blob/main{/dir}/{file}#L{line}">
</head>
</html>
`)
			return
		}
		if path == "/" || strings.Contains(path, "/http") {
			io.WriteString(w, `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="inet.af git https://github.com/bradfitz/exp-httpclient">
<meta name="go-source" content="inet.af https://github.com/bradfitz/exp-httpclient/ https://github.com/bradfitz/exp-httpclient/tree/master{/dir} https://github.com/bradfitz/exp-httpclient/blob/master{/dir}/{file}#L{line}">
</head>
</html>
`)
		}
		return
	}
	if path == "/http" {
		http.Redirect(w, r, "https://github.com/bradfitz/exp-httpclient", http.StatusFound)
		return
	}
	if strings.HasPrefix(path, "/tcpproxy") ||
		strings.HasPrefix(path, "/netaddr") ||
		strings.HasPrefix(path, "/netstack") ||
		strings.HasPrefix(path, "/peercred") ||
		strings.HasPrefix(path, "/wf") {
		http.Redirect(w, r, "https://pkg.go.dev/inet.af"+path, http.StatusFound)
		return
	}
	if strings.Contains(path, "/http") {
		target := "https://pkg.go.dev/inet.af" + path
		http.Redirect(w, r, target, http.StatusFound)
		return
	}
	const target = "http://pubs.opengroup.org/onlinepubs/009696699/basedefs/sys/socket.h.html"
	http.Redirect(w, r, target, http.StatusFound)
}
