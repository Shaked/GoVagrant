package main

import (
 "net"
 "net/http"
 "net/http/fcgi"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
 resp.Write([]byte("<h1>Hello, 世界</h1>\n<p>Behold my Go web app.</p>"))
}

func main() {
 listener, _ := net.Listen("tcp", "127.0.0.1:9001")
 srv := new(FastCGIServer)
 fcgi.Serve(listener, srv)
}
