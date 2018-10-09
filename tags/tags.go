package tags

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type Tag struct {
	Sig    string
	Offset uint32
	Size   uint32
}

func NewTag() *Tag {
	return new(Tag)
}

func (tag *Tag) Parse(file *os.File) {
	tag.Sig = fmtbytes.ReadBytesString(file, 4)
	tag.Offset = fmtbytes.ReadBytesLong(file)
	tag.Size = fmtbytes.ReadBytesLong(file)
}
