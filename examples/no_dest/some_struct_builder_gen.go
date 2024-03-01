// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package no_dest

// SomeStructBuilder is an alias of SomeStruct to build SomeStruct with builder-pattern.
type SomeStructBuilder SomeStruct

// NewSomeStructBuilder creates a new SomeStructBuilder.
func NewSomeStructBuilder() *SomeStructBuilder {
	return &SomeStructBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *SomeStructBuilder) Copy() *SomeStructBuilder {
	c := *b
	return &c
}

// Build returns built SomeStruct.
func (b *SomeStructBuilder) Build() *SomeStruct {
	return (*SomeStruct)(b)
}

// SetSomeChan sets SomeStruct's SomeChan.
func (b *SomeStructBuilder) SetSomeChan(someChan chan<- SomeAlias) *SomeStructBuilder {
	b.SomeChan = someChan
	return b
}

// SetSomeSlice sets SomeStruct's SomeSlice.
func (b *SomeStructBuilder) SetSomeSlice(someSlice []string) *SomeStructBuilder {
	b.SomeSlice = someSlice
	return b
}