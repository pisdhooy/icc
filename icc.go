package icc

import (
	"os"

	"github.com/pisdhooy/icc/header"
	"github.com/pisdhooy/icc/table"
	"github.com/pisdhooy/icc/tag"
	"github.com/pisdhooy/icc/tag/v4"
)

type ICCProfile struct {
	Header   *header.Header
	TagTable *table.TagTable
	TagData  []tag.DataTag
}

func (iccProfile *ICCProfile) GetTypeID() int {
	return 1039
}

func NewICCProfile() *ICCProfile {
	return new(ICCProfile)
}

func (iccProfile *ICCProfile) Parse(file *os.File, version int) {
	header := header.NewHeader()
	tagTable := table.NewTagList()

	header.Parse(file)
	iccProfile.Header = header
	tagTable.Parse(file)
	iccProfile.TagTable = tagTable
	for i := 0; i < int(iccProfile.TagTable.Count); i++ {
		switch version {
		case 2:
			//TODO: make a polyfill for this
		case 4:
			tag := v4.ParseTagData(file, iccProfile.TagTable.Tags[i].Sig)
			iccProfile.TagData = append(iccProfile.TagData, tag)
		default:
			//TODO: make this not panic and actually return an error
			panic("unknown icc version")
		}
	}
}
