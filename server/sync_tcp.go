package server

import (
	"diceDb/config"
	"io"
	"log"
	"net"
	"strconv"
)

func readCommand(c net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respond(cmd string, c net.Conn) error {
	_, err := c.Write([]byte(cmd))
	if err != nil {
		return err
	}

	return nil
}

func RunSyncTCPServer() {
	log.Println("Starting a synchronous TCP server...", config.Host, config.Port)
	var con_clients int = 0

	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}

	for {
		c, err := lsnr.Accept()
		if err != nil {
			panic(err)
		}

		// increment the number of concurrent clients
		con_clients++
		log.Println("Client connected with address:", c.RemoteAddr(),
			"concurrent_clients", con_clients)

		for {
			// over the socket connection, continuously read the command and print it out
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				con_clients--
				log.Println("Client disconnected", c.RemoteAddr(), "concurrent clients", con_clients)
				if err == io.EOF {
					break
				}
				log.Println("err", err)
			}

			err = respond(cmd, c)
			if err != nil {
				log.Println("err write:", err)
			}
			log.Println("command", cmd)
		}
	}
}
