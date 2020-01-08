package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	fmt.Printf("Starting server. Listening on port 8080\n")
	fmt.Printf("request to http://<ip_addr>:8080/<name> returns \"Hello, <name>\"\n")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
	//	fmt.Fprintf(w, "ip %s", getIpAddrs())
}

func getIpAddrs() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}
	fmt.Printf("%v\n", addrs)
}
