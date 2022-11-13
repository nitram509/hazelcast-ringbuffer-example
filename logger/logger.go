package main

import (
	"context"
	"github.com/hazelcast/hazelcast-go-client"
	"log"
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

	for sequence := int64(0); true; sequence++ {
		item, err := rb.ReadOne(ctx, sequence)
		if err != nil {
			panic(err)
		}
		log.Printf("Received item idx=%d: %s", sequence, item)
	}

}
