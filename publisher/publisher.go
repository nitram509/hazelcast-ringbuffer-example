package main

import (
	"context"
	"fmt"
	"github.com/hazelcast/hazelcast-go-client"
	"time"
)

func main() {
	ctx := context.TODO()

	client, err := hazelcast.StartNewClient(ctx)
	if err != nil {
		panic(err)
	}

	rb, err := client.GetRingbuffer(ctx, "my-ringbuffer")
	if err != nil {
		panic(err)
	}

	msg := fmt.Sprintf("Hello World, sender's time is %s", time.Now())
	rb.Add(ctx, msg, hazelcast.OverflowPolicyOverwrite)
}
