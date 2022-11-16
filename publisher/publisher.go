package main

import (
	"context"
	"fmt"
	"github.com/hazelcast/hazelcast-go-client"
	"log"
	"math/rand"
	"time"
)

func main() {
	ctx := context.TODO()
	rand.Seed(time.Now().UnixNano())

	client, err := hazelcast.StartNewClient(ctx)
	if err != nil {
		panic(err)
	}

	rb, err := client.GetRingbuffer(ctx, "my-ringbuffer")
	if err != nil {
		panic(err)
	}

	log.Println("Processing. Press CTRL+C to exit.")
	for true {
		msg := fmt.Sprintf("Hello World, sender's time is %s", time.Now())
		if rand.Intn(100) < 10 {
			// a 10% chance of sending an alert
			msg = fmt.Sprintf("Alert at %s", time.Now())
		}
		sequence, err := rb.Add(ctx, msg, hazelcast.OverflowPolicyOverwrite)
		log.Printf("Published with sequence=%d, msg=%s", sequence, msg)
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second) // just be graceful in our little demo
	}
}
