package tcp

import (
	"fmt"
	"io"
	"log"
	"net"
	"steam-lancache/internal/config"
)

func handleConnection(clientConnection net.Conn, targetAddress string) {
	defer clientConnection.Close()

	serverConnection, err := net.Dial("tcp", targetAddress)
	if err != nil {
		log.Printf("[TCP] Failed to connect to SteamDL on %s: %v", targetAddress, err)
		return
	}
	defer serverConnection.Close()

	go io.Copy(serverConnection, clientConnection)
	io.Copy(clientConnection, serverConnection)
}

func Start(upstreamIP string) {
	listener, err := net.Listen("tcp", ":"+config.HTTPSPort)
	if err != nil {
		log.Fatalf("[TCP] Failed to start TCP listener: %v", err)
	}
	defer listener.Close()

	target := fmt.Sprintf("%s:%s", upstreamIP, config.HTTPSPort)
	log.Printf("[TCP] Proxy listening on %s -> %s", config.HTTPSPort, target)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn, target)
	}
}
