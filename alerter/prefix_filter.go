package main

import (
	"fmt"
	"github.com/hazelcast/hazelcast-go-client/serialization"
)

type PrefixFilter struct {
	Prefix string
}

func (p *PrefixFilter) FactoryID() int32 {
	return 666
}

func (p *PrefixFilter) ClassID() int32 {
	return 14
}

func (p *PrefixFilter) WriteData(output serialization.DataOutput) {
	output.WriteString(p.Prefix)
}

func (p *PrefixFilter) ReadData(input serialization.DataInput) {
	p.Prefix = input.ReadString()
}

type IdentifiedFactory struct{}

func (f IdentifiedFactory) Create(id int32) serialization.IdentifiedDataSerializable {
	if id == 14 {
		return &PrefixFilter{}
	}
	panic(fmt.Sprintf("unknown class ID: %d", id))
}

func (f IdentifiedFactory) FactoryID() int32 {
	return 666
}
