package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// package main

// main
func main() {
	fmt.Println("1. set server")
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8887,
		//Zone: "",
	})
	if err != nil {
		return
	}
	defer con.Close()

	fmt.Println("2. new reader")
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("3. input message :")
		line, err := input.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println("3. send message to server")
		n, err := con.Write([]byte(line))
		if err != nil {
			return
		}

		fmt.Println("4. receive message from server")
		buf := make([]byte, 1024)
		n, addr, err := con.ReadFromUDP(buf)
		if err != nil {
			return
		}
		fmt.Printf("read from %v, msg : %v\n", addr, string(buf[:n]))
	}

}
