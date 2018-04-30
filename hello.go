package main

import (
	"net"
	"fmt"
)

func main() {

	myIp := getMyIp()
	fmt.Println("myIp:", myIp)

	listener, _ := net.Listen("tcp", "0.0.0.0:9999")
	addr := listener.Addr().String()

	fmt.Println("listener.Addr()", addr)
	host, port, _ := net.SplitHostPort(addr)
	fmt.Println("host:", host)
	fmt.Println("port:", port)

}

func getMyIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, a := range addrs {
		ipnet, ok := a.(*net.IPNet)
		if !ok {
			continue
		}
		v4 := ipnet.IP.To4()
		if v4 == nil || v4[0] == 127 {
			continue
		} // loopback
		return ipnet.IP.String()
	}
	return ""
}
