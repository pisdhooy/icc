package tags

import (
	"os"

	"github.com/pisdhooy/fsutil"
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
	tag.Sig = fsutil.ReadBytesString(file, 4)
	tag.Offset = fsutil.ReadBytesLong(file)
	tag.Size = fsutil.ReadBytesLong(file)
}
