package main

import (
	"fmt"
	"net"
)

// package main

// process
func process(con *net.UDPConn, addr *net.UDPAddr, data []byte) error {
	fmt.Printf("client : %s", string(data))

	fmt.Println("3. write to client")
	n, err := con.WriteToUDP(data, addr)
	if err != nil {
		return err
	}
	fmt.Printf("n is %v\n\n", n)

	return nil
}

// main
func main() {
	fmt.Println("1. listen port ...")
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8887,
		//Zone: "",
	})
	if err != nil {
		return
	}
	defer con.Close()

	for {
		fmt.Println("2. get data from client")
		buf := make([]byte, 1024)
		n, addr, err := con.ReadFromUDP(buf)
		if err != nil {
			return
		}

		data := make([]byte, n)
		copy(data, buf[:n])
		err = process(con, addr, data)
		if err != nil {
			return
		}
	}
}
