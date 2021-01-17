package main

import (
	"net"
	"fmt"
	"time"
	"github.com/powergee/judge-server-communication-layer/utils"
)

func main() {
	server, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("net.Listen:", err)
	}
	defer server.Close()

	for {
		fmt.Println("Listening:9999...")
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("server.Accept:", err)
			continue
		}
		conn.SetReadDeadline(time.Time{})
		defer conn.Close()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Connected with judge server!")

	if info := utils.HandShake(conn); info != nil {
		fmt.Println("Handshaking is succeeded!")
	} else {
		fmt.Println("Handshaking is failed...")
		return
	}

	for conn != nil {
		fmt.Println()
		fmt.Println("1. ping")
		fmt.Println("2. get-current-submission")
		fmt.Println("3. submission-request")
		fmt.Println("4. terminate-submission")
		fmt.Println("5. disconnect")
		fmt.Print("Comm > ")

		var sel int
		fmt.Scan(&sel)

		switch sel {
		case 1:
			conn = utils.SendPing(conn)
		case 2:
			conn = utils.GetCurrentSubmission(conn)
		case 3:
			conn = utils.RequestSubmission(conn)
		case 4:

		case 5:
			conn = utils.Disconnect(conn)
		default:
			fmt.Println("Wrong command. Try again.")
		}
	}
}