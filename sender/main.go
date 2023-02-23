package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Failed Initializing Broker Connection")
		panic(1)
	}
	defer conn.Close() // defer to close connection

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close() // defer to close channel

	// declare queue in channel
	q, err := ch.QueueDeclare("TestQueue", false, false, false, false, nil)
	log.Println(q) // print status
	if err != nil {
		log.Println(err)
	}

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello rabbitmq"),
		},
	)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Published Message to Queue")

}
