package description

import (
	"fmt"
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type Description struct {
	Signature      string
	Reserved       []byte
	NumberOfNames  uint32
	NameRecordSize uint32
	Names          []*Name
}

type Name struct {
	LanguageCode         uint16
	FirstNameCountryCode uint16
	FirstNameLength      uint32
	FirstNameOffset      uint32
	Value                string
}

func (Description *Description) GetTypeName() string {
	return "desc"
}

func NewDescription() *Description {
	return new(Description)
}

func NewName() *Name {
	return new(Name)
}

func (name *Name) Parse(file *os.File, tagStartPos int) {
	name.LanguageCode = fmtbytes.ReadBytesShort(file)
	name.FirstNameCountryCode = fmtbytes.ReadBytesShort(file)
	name.FirstNameLength = fmtbytes.ReadBytesLong(file)
	name.FirstNameOffset = fmtbytes.ReadBytesLong(file)
	tmpOldFilePointerPos, _ := file.Seek(0, 1)
	file.Seek(int64(int(name.FirstNameOffset)+tagStartPos), 1)
	name.Value = fmtbytes.ReadBytesString(file, int(name.FirstNameLength))
	file.Seek(tmpOldFilePointerPos, 1)
}

func (description *Description) Parse(file *os.File) error {
	description.Signature = fmtbytes.ReadBytesString(file, 4)
	if description.Signature != "mluc" {
		return fmt.Errorf("invalid description signature")
	}

	description.Reserved = fmtbytes.ReadRawBytes(file, 4)

	description.NumberOfNames = fmtbytes.ReadBytesLong(file)
	description.NameRecordSize = fmtbytes.ReadBytesLong(file)

	for i := 0; i < int(description.NumberOfNames); i++ {
		name := NewName()
		currentFilePointerPos, _ := file.Seek(0, 1)
		name.Parse(file, int(currentFilePointerPos))
		description.Names = append(description.Names, name)
	}
	return nil
}
