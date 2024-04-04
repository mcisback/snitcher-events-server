package queue

type Message struct {
	message string
	data    any
}

type Queue struct {
	items   []Message
	size    int
	channel chan *Message
}

func New() *Queue {
	q := &Queue{}

	q.channel = make(chan *Message)

	return q
}

func (q *Queue) Add(message string, data any) {
	newMessage := &Message{
		message: message,
		data:    data,
	}

	q.items = append(q.items, *newMessage)

	q.size++

	q.channel <- newMessage
}

func (q *Queue) Next() *Message {
	if q.size == 0 {
		return nil
	}

	dequeued := q.items[0]
	q.items = q.items[1:]

	q.size--

	return &dequeued
}

func (q *Queue) Size() int {
	return q.size
}
