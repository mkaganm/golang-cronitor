package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	// Define a cron job with a cron-like schedule expression
	_, err := c.AddFunc("* * * * *", func() {
		fmt.Println("This is a cron job!")
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	c.Start()

	// Wait until the application is stopped
	select {}
}
