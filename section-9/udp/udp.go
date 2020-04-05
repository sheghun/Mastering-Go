package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8000", "address? host:port")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		_ = runServer(*address)
	case "C":
		_ = runClient(*address)
	}
}

func runClient(address string) error {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println("What message would like to send?")
		_, err = conn.Write(scanner.Bytes())

		if err != nil {
			log.Fatal(err)
		}

		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(buffer))
		fmt.Println("What message would like to send?")
	}

	return nil
}

func runServer(address string) error {

	return nil
}
