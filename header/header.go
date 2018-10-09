package header

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type Header struct {
	Size            uint32
	Cmmid           string
	Version         uint32
	DeviceClass     string
	ColorSpace      string
	Pcs             string
	DateTime        *DateTime
	Magic           string
	Platform        string
	Flags           uint32
	Manufacturer    string
	Model           uint32
	Attributes      uint64
	RenderingIntent uint32
	Illuminant      *Vec3_32
	Creator         string
	ProfileId       []byte
	Reserved        []byte
}

type DateTime struct {
	Year    uint16
	Month   uint16
	Day     uint16
	Hours   uint16
	Minutes uint16
	Seconds uint16
}

type Vec3_32 struct {
	X uint32
	Y uint32
	Z uint32
}

func NewHeader() *Header {
	return new(Header)
}

func NewDateTime() *DateTime {
	return new(DateTime)
}

func NewVec3_32() *Vec3_32 {
	return new(Vec3_32)
}

func (dateTime *DateTime) Parse(file *os.File) {
	dateTime.Year = fmtbytes.ReadBytesShort(file)
	dateTime.Month = fmtbytes.ReadBytesShort(file)
	dateTime.Day = fmtbytes.ReadBytesShort(file)
	dateTime.Hours = fmtbytes.ReadBytesShort(file)
	dateTime.Minutes = fmtbytes.ReadBytesShort(file)
	dateTime.Seconds = fmtbytes.ReadBytesShort(file)
}

func (header *Header) Parse(file *os.File) {
	dateTimeObject := NewDateTime()
	illuminantObject := NewVec3_32()

	header.Size = fmtbytes.ReadBytesLong(file)
	header.Cmmid = fmtbytes.ReadBytesString(file, 4)
	header.Version = fmtbytes.ReadBytesLong(file)
	header.DeviceClass = fmtbytes.ReadBytesString(file, 4)
	header.ColorSpace = fmtbytes.ReadBytesString(file, 4)
	header.Pcs = fmtbytes.ReadBytesString(file, 4)

	dateTimeObject.Parse(file)

	header.DateTime = dateTimeObject

	header.Magic = fmtbytes.ReadBytesString(file, 4)
	header.Platform = fmtbytes.ReadBytesString(file, 4)
	header.Flags = fmtbytes.ReadBytesLong(file)

	header.Manufacturer = fmtbytes.ReadBytesString(file, 4)

	header.Model = fmtbytes.ReadBytesLong(file)
	header.Attributes = fmtbytes.ReadBytesLongLong(file)

	header.RenderingIntent = fmtbytes.ReadBytesLong(file)

	illuminantObject.Parse(file)

	header.Illuminant = illuminantObject
	header.Creator = fmtbytes.ReadBytesString(file, 4)

	header.ProfileId = fmtbytes.ReadBytesNInt(file, 16)
	header.Reserved = fmtbytes.ReadBytesNInt(file, 28)
}

func (header *Header) GetFullname(field string) string {
	nameMap := map[string]string{
		"ADBE": "Adobe Systems Inc.",
		"ACMS": "Agfa Graphics N.V. ",
		"appl": "Apple Computer",
		"CCMS": "Canon",
		"UCCM": "Canon",
		"UCMS": "Canon",
		"EFI":  "EFI",
		"FF ":  "Fuji Film Electronic Imaging",
		"EXAC": "ExactCODE GmbH",
		"HCMM": "Global Graphics Software Inc",
		"argl": "Graeme Gill",
		"LgoS": "GretagMacbeth",
		"HDM ": "Heidelberger Druckmaschinen AG",
		"lcms": "Hewlett Packard ",
		"RMIX": "ICC",
		"KCMS": "Kodak",
		"MCML": "Konica Minolta",
		"WCS":  "Microsoft",
		"SIGN": "Mutoh",
		"ONYX": "Onyx Graphics",
		"RGMS": "Rolf Gierling Multitools",
		"SICC": "SampleICC",
		"TCMM": "Toshiba TEC Corporation",
		"32BT": "the imaging factory",
		"vivo": "Vivo Mobile Communication",
		"WTG ": "Ware To Go",
		"zc00": "Zoran Corporation",
	}

	if val, ok := nameMap[field]; ok {
		return val
	}
	return "unknown manufacturer"
}

func (vec3_32 *Vec3_32) Parse(file *os.File) {
	vec3_32.X = fmtbytes.ReadBytesLong(file)
	vec3_32.Y = fmtbytes.ReadBytesLong(file)
	vec3_32.Z = fmtbytes.ReadBytesLong(file)
}
