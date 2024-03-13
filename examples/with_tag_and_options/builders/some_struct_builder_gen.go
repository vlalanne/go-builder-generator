// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"fmt"

	"github.com/kilianpaquier/go-builder-generator/examples/with_tag_and_options"
)

// SomeStructBuilder is an alias of SomeStruct to build SomeStruct with builder-pattern.
type SomeStructBuilder with_tag_and_options.SomeStruct

// NewSomeStructBuilder creates a new SomeStructBuilder.
func NewSomeStructBuilder() *SomeStructBuilder {
	return &SomeStructBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *SomeStructBuilder) Copy() *SomeStructBuilder {
	result := *b
	return &result
}

// Build returns built SomeStruct.
func (b *SomeStructBuilder) Build() (*with_tag_and_options.SomeStruct, error) {
	b = b.SetTheChan()

	result := (with_tag_and_options.SomeStruct)(*b)
	if err := result.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate 'SomeStruct' struct: %w", err)
	}
	return &result, nil
}

// SetSomeSlice sets SomeStruct's SomeSlice.
func (b *SomeStructBuilder) SetSomeSlice(someSlice ...string) *SomeStructBuilder {
	b.SomeSlice = append(b.SomeSlice, someSlice...)
	return b
}