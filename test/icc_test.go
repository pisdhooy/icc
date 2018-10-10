package test

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pisdhooy/icc"
)

// func TestICCParserV2(t *testing.T) {
// 	profile := icc.NewICCProfile()
// 	file, err := os.Open("files/Probev1_ICCv2.icc")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	profile.Parse(file, 2)
// 	// spew.Dump(profile.TagTable)
// }

func TestICCParserV4(t *testing.T) {
	profile := icc.NewICCProfile()
	file, err := os.Open("files/Probev1_ICCv4.icc")
	if err != nil {
		t.Error(err)
	}

	profile.Parse(file, 4)
	spew.Dump(profile.TagData)
}
