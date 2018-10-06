package icc

import (
	"os"

	"github.com/pisdhooy/fsutil"
	"github.com/pisdhooy/icc/header"
	"github.com/pisdhooy/icc/tags"
)

type ICCProfile struct {
	Header   *header.Header
	TagTable *TagTable
	TagData  [][]byte
}

type TagTable struct {
	Count uint32
	Tags  []*tags.Tag
}

func (iccProfile *ICCProfile) GetTypeID() int {
	return 1039
}

func NewICCProfile() *ICCProfile {
	return new(ICCProfile)
}

func (iccProfile *ICCProfile) Parse(file *os.File) {
	header := header.NewHeader()
	tagTable := NewTagList()

	header.Parse(file)
	iccProfile.Header = header
	tagTable.Parse(file)
	iccProfile.TagTable = tagTable
	for i := 0; i < int(iccProfile.TagTable.Count); i++ {
		offset := int64(iccProfile.TagTable.Tags[i].Offset)
		file.Seek(offset, 0)
		buffer := fsutil.ReadBytesNInt(file, iccProfile.TagTable.Tags[i].Size)
		iccProfile.TagData = append(iccProfile.TagData, buffer)
	}
}

func NewTagList() *TagTable {
	return new(TagTable)
}

func (tagTable *TagTable) Parse(file *os.File) {
	tagTable.Count = fsutil.ReadBytesLong(file)
	for i := 0; i < int(tagTable.Count); i++ {
		tag := tags.NewTag()
		tag.Parse(file)
		tagTable.Tags = append(tagTable.Tags, tag)
	}
}
