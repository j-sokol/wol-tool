package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

func sendWoLPacket(macAddress string, broadcastAddress string) error {
	macAddr, err := net.ParseMAC(macAddress)
	if err != nil {
		return err
	}

	// Create magic packet
	magicPacket := bytes.Repeat([]byte{0xFF}, 6)
	for i := 0; i < 16; i++ {
		magicPacket = append(magicPacket, macAddr...)
	}

	// Send UDP broadcast
	conn, err := net.Dial("udp", broadcastAddress+":9")
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(magicPacket)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [MAC address] [IP address]")
		os.Exit(1)
	}

	macAddress := os.Args[1]
	ipAddress := os.Args[2]

	err := sendWoLPacket(macAddress, ipAddress)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
