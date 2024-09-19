package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func ProduceMessage() {
	topic := "your-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	defer conn.Close()

	msg := kafka.Message{
		Value: []byte("Hello, Kafka!"),
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(msg)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	log.Println("Message sent successfully!")
}

func ConsumeMessages() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "your-topic",
		Partition: 0,
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		log.Printf("received: %s", string(msg.Value))
	}
}

func main() {
	ProduceMessage()
	ConsumeMessages()
}
