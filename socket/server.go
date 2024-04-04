package socket

import (
	"encoding/json"
	"fmt"
	"net"

	qm "github.com/mcisback/snitcher-events-server/queuemanager"
	qs "github.com/mcisback/snitcher-events-server/queuemanager/qsubscriber"
)

type TCPServer struct {
	QManager *qm.QManager
}

type TCPCommand struct {
	Cmd     string         `json:"cmd"`
	Payload map[string]any `json:"payload"`
}

func New() *TCPServer {
	t := &TCPServer{}
	t.QManager = qm.New()

	return t
}

func (t *TCPServer) Start(port string) {
	fmt.Printf("TCP Server is running on port %s\n", port)

	// listen on a port
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// accept a connection
		connection, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go t.HandleNewConnection(connection)
	}
}

func (t *TCPServer) HandleNewConnection(connection net.Conn) {
	remoteAddr := connection.RemoteAddr().String()

	fmt.Println("Received new connection from: ", remoteAddr)

	d := json.NewDecoder(connection)

	var tcpCommand TCPCommand

	d.Decode(&tcpCommand)
	// fmt.Println(tcpCommand, err)

	//defer connection.Close()

	fmt.Println("Received CMD", tcpCommand.Cmd)

	if tcpCommand.Cmd == "SUBSCRIBE" {

		qName := tcpCommand.Payload["queue"].(string)
		clientName := tcpCommand.Payload["clientName"].(string)

		// TODO: check if qname exists
		subscriber := qs.QSubscriber{
			Name:       clientName,
			Connection: connection,
			RemoteAddr: remoteAddr,
		}

		t.QManager.AddSubscriber(qName, &subscriber)

	} else if tcpCommand.Cmd == "SNITCH" {

		qName := tcpCommand.Payload["queue"].(string)
		message := tcpCommand.Payload["message"].(string)
		data := tcpCommand.Payload["data"]

		t.QManager.Snitch(qName, message, data)

	}
}
