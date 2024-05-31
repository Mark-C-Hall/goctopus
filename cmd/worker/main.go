package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/mark-c-hall/goctopus/internal/queue"
	"github.com/mark-c-hall/goctopus/internal/tasks"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	rabbitMQ, err := queue.NewRabbitMQ(os.Getenv("RABBITMQ_URL"), os.Getenv("RABBITMQ_QUEUE"))
	if err != nil {
		log.Fatalln(err)
	}
	defer rabbitMQ.Connection.Close()
	defer rabbitMQ.Channel.Close()

	task := tasks.EmailTask{
		To:      "jdoe@mail.com",
		Subject: "Hello, World!",
		Body:    "This is a test email.",
	}

	rabbitMQ.Push(task)
	fmt.Println("Task pushed to RabbitMQ")

	msgs, err := rabbitMQ.Pull()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for msg := range msgs {
			fmt.Println("Received a message: ", string(msg.Body))
			msg.Ack(false)
		}
	}()

	// Wait for messages
	select {}
}
