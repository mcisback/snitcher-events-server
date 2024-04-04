package qsubscriber

import (
	"fmt"
	"net"

	"github.com/mcisback/snitcher-events-server/queue"
)

type QSubscriber struct {
	Name       string
	Connection net.Conn
	RemoteAddr string
}

func (sub *QSubscriber) Notify(msg *queue.Message) {

	// jsonData, err := json.Marshal(*msg)
	// if err != nil {
	// 	fmt.Printf("Error marshaling JSON: %v", err)
	// }

	// TODO: Send to socket
	fmt.Printf("Notifing %s (%s) of %s", sub.Name, sub.RemoteAddr, msg.Message)
}
