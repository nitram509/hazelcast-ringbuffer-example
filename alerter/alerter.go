package main

import (
	"context"
	"github.com/hazelcast/hazelcast-go-client"
	"log"
	"strings"
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

	log.Println("Processing. Press CTRL+C to exit.")
	for sequence := int64(0); true; {
		resultSet, err := rb.ReadMany(ctx, sequence, 5, 10, nil)
		if err != nil {
			panic(err)
		}
		for i := 0; i < resultSet.Size(); i++ {
			if msg, err := resultSet.Get(i); err == nil && strings.HasPrefix(msg.(string), "Alert") {
				msgSequence, _ := resultSet.GetSequence(i)
				log.Printf("Received alert sequence=%d msg=%s", msgSequence, msg)
			}
		}
		sequence = resultSet.GetNextSequenceToReadFrom()
	}

}
