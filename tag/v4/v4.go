package v4

import (
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/pisdhooy/icc/tag/v4/types/description"
)

//this is horrible, but unfortunately the only solution afaik.
type TagV4 struct {
	Value DataTag
}

type DataTag interface {
	GetTypeName() string
}

func (tagV4 *TagV4) GetTagVersion() int {
	return 4
}

func NewTagV4() *TagV4 {
	return new(TagV4)
}

func ParseTagData(file *os.File, tagName string) *TagV4 {

	tag := NewTagV4()

	switch tagName {
	case "desc":
		descObject := description.NewDescription()
		descObject.Parse(file)
		tag.Value = descObject
	case "A2B0":

	}
	spew.Dump(tag.Value)

	return tag
}
