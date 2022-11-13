package main

import (
	"context"
	"github.com/hazelcast/hazelcast-go-client"
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

	println(rb.Name())
}
