package table

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type TableTag struct {
	Sig    string
	Offset uint32
	Size   uint32
}

type TagTable struct {
	Count uint32
	Tags  []*TableTag
}

func NewTableTag() *TableTag {
	return new(TableTag)
}

func (tag *TableTag) Parse(file *os.File) {
	tag.Sig = fmtbytes.ReadBytesString(file, 4)
	tag.Offset = fmtbytes.ReadBytesLong(file)
	tag.Size = fmtbytes.ReadBytesLong(file)
	if tag.Size%4 != 0 {
		tag.Size = ((tag.Size + 4 - 1) / 4) * 4
	}
}

func NewTagList() *TagTable {
	return new(TagTable)
}

func (tagTable *TagTable) Parse(file *os.File) {
	tagTable.Count = fmtbytes.ReadBytesLong(file)
	for i := 0; i < int(tagTable.Count); i++ {
		tag := NewTableTag()
		tag.Parse(file)
		tagTable.Tags = append(tagTable.Tags, tag)
	}
}
