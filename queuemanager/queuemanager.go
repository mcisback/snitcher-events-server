package queue

import (
	"fmt"
	"time"

	"github.com/mcisback/snitcher-events-server/queue"
	qs "github.com/mcisback/snitcher-events-server/queuemanager/qsubscriber"
)

type QMap map[string]*QManagerNode

type QManagerNode struct {
	name        string
	queue       *queue.Queue
	subscribers []*qs.QSubscriber
}

type QManager struct {
	Qmap QMap
}

func New() *QManager {
	q := &QManager{}

	q.Qmap = make(QMap)

	return q
}

func (q *QManager) InitNotifier() {
	for key := range q.Qmap {
		node := q.Qmap[key]

		go node.LoopNotify()
	}
}

func (q *QManager) AddQ(name string) {
	node := &QManagerNode{}

	node.queue = queue.New()
	node.subscribers = make([]*qs.QSubscriber, 0)

	node.name = name

	q.Qmap[name] = node

	go node.LoopNotify()
}

// TODO: Add to Queue

func (q *QManager) Snitch(qName string, message string, data any) {
	q.GetNode(qName).queue.Add(message, data)

	fmt.Printf("[%s] Snitching %s", qName, message)
	fmt.Println(data)
}

func (q *QManager) GetNode(qName string) *QManagerNode {
	return q.Qmap[qName]
}

func (q *QManager) DelQ(name string) {
	delete(q.Qmap, name)
}

func (q *QManager) AddSubscriber(qName string, sub *qs.QSubscriber) {
	q.Qmap[qName].subscribers = append(q.Qmap[qName].subscribers, sub)

	fmt.Printf("[%s] Added subscriber %s\n", qName, sub.Name)
}

// Check for new events in queue and notify all subscribers
func (n *QManagerNode) LoopNotify() {
	for {
		message := <-n.queue.Channel

		if n.queue.Size() <= 0 {
			fmt.Println("No Messages in queue: ", n.name)

			time.Sleep(time.Second * 1)

			continue
		}

		for _, sub := range n.subscribers {
			go sub.Notify(message)
		}

		// for {
		// 	message := n.queue.Next()

		// 	if message == nil {
		// 		break
		// 	}

		// 	for _, sub := range n.subscribers {
		// 		sub.Notify(message)
		// 	}
		// }

	}
}
