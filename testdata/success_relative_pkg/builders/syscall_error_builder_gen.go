// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"os"
)

// SyscallErrorBuilder is an alias of SyscallError to build SyscallError with builder-pattern.
type SyscallErrorBuilder os.SyscallError

// NewSyscallErrorBuilder creates a new SyscallErrorBuilder.
func NewSyscallErrorBuilder() *SyscallErrorBuilder {
	return &SyscallErrorBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *SyscallErrorBuilder) Copy() *SyscallErrorBuilder {
	c := *b
	return &c
}

// Build returns built SyscallError.
func (b *SyscallErrorBuilder) Build() *os.SyscallError {
	c := (os.SyscallError)(*b)
	return &c
}

// SetErr sets SyscallError's Err.
func (b *SyscallErrorBuilder) SetErr(err error) *SyscallErrorBuilder {
	b.Err = err
	return b
}

// SetSyscall sets SyscallError's Syscall.
func (b *SyscallErrorBuilder) SetSyscall(syscall string) *SyscallErrorBuilder {
	b.Syscall = syscall
	return b
}