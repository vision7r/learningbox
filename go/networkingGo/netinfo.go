package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Printf("Lets See The Network Information.....\n")

	ifaces, err := net.Interfaces()

	if err != nil {
		fmt.Print(err)
		return
	}

	for _, iface := range ifaces {
		address, _ := iface.Addrs()

		fmt.Printf("Interface:%s, Mac:%s, MTU:%d, Address:%v\n", iface.Name, iface.HardwareAddr, iface.MTU, address)
	}
}
