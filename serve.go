package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func name() string {
	name, err := os.Hostname()
	if err != nil {
		return ""
	}
	return name
}

func ip() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func n(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HostName: ", name())
}

func i(w http.ResponseWriter, r *http.Request) {
	i := ip()
	fmt.Fprintln(w, "IP Address: ", i)
}

func hello(w http.ResponseWriter, r *http.Request) {
	s := "Hello from App Runner. HostName: " + name() + " IP Address: " + ip()
	fmt.Fprintln(w, s)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
