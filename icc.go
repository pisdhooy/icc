package icc

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/pisdhooy/fmtbytes"
	"github.com/pisdhooy/icc/header"
	"github.com/pisdhooy/icc/table"
)

type ICCProfile struct {
	Header   *header.Header
	TagTable *table.TagTable
	TagData  [][]byte
}

func (iccProfile *ICCProfile) GetTypeID() int {
	return 1039
}

func NewICCProfile() *ICCProfile {
	return new(ICCProfile)
}

func (iccProfile *ICCProfile) Parse(file *os.File) {
	tmpFP, _ := file.Seek(0, 1)
	header := header.NewHeader()
	tagTable := table.NewTagList()

	header.Parse(file)
	iccProfile.Header = header
	tagTable.Parse(file)
	iccProfile.TagTable = tagTable

	for i := 0; i < int(iccProfile.TagTable.Count); i++ {
		offset := int64(iccProfile.TagTable.Tags[i].Offset)
		file.Seek(tmpFP+offset, 0)
		buffer := fmtbytes.ReadBytesNInt(file, iccProfile.TagTable.Tags[i].Size)
		iccProfile.TagData = append(iccProfile.TagData, buffer)
	}
	fmt.Println("FP AFTER ICC PARSE")
	spew.Dump(file.Seek(0, 1))
}
