package tasks

import (
	"fmt"
	"time"
)

type EmailTask struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (e EmailTask) Execute() error {
	// Simulating Email
	fmt.Printf("Sending email to %s with subject %s\n", e.To, e.Subject)
	time.Sleep(5 * time.Second)
	fmt.Println("Email Sent.")

	return nil
}

func (e EmailTask) GetType() string {
	return "email"
}

func (e EmailTask) GetPayload() map[string]interface{} {
	return map[string]interface{}{
		"to":      e.To,
		"subject": e.Subject,
		"body":    e.Body,
	}
}
