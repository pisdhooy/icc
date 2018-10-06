package test

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/pisdhooy/icc"
)

func TestICCParser(t *testing.T) {
	profile := icc.NewICCProfile()
	file, err := os.Open("files/Probev1_ICCv2.icc")
	if err != nil {
		t.Error(err)
	}

	profile.Parse(file)
	spew.Dump(profile.TagData[0])
}
