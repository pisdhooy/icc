package a2b0

import "os"

type A2B0 struct {
	Signature         string
	reserved          uint32
	NumInputChannels  byte
	NumOutputChannels byte
	CLUT              byte
	padding           byte
	EE0               uint32
	EE1               uint32
	EE2               uint32
	EE3               uint32
	EE4               uint32
	EE5               uint32
	EE6               uint32
	EE7               uint32
	EE8               uint32
	InputTableSize    uint32
	OutputTableSize   uint32
}

func (a2b0 *A2B0) GetTypeName() string {
	return "A2B0"
}

func NewA2B0() *A2B0 {
	return new(A2B0)
}

func (obj *A2B0) Parse(file *os.File) {

}
