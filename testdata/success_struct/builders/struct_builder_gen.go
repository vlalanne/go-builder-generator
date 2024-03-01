// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"context"

	"github.com/kilianpaquier/go-builder-generator/testdata/success_struct"
)

// StructBuilder is an alias of Struct to build Struct with builder-pattern.
type StructBuilder success_struct.Struct

// NewStructBuilder creates a new StructBuilder.
func NewStructBuilder() *StructBuilder {
	return &StructBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *StructBuilder) Copy() *StructBuilder {
	c := *b
	return &c
}

// Build returns built Struct.
func (b *StructBuilder) Build() *success_struct.Struct {
	c := (success_struct.Struct)(*b)
	return &c
}

// SetAnotherStruct sets Struct's AnotherStruct.
func (b *StructBuilder) SetAnotherStruct(anotherStruct struct {
	Nested struct {
		Field string
		Ctx   context.Context
	}
	NotNested int64
	Ctx       context.Context
	Alias     success_struct.Int64Alias
}) *StructBuilder {
	b.AnotherStruct = anotherStruct
	return b
}