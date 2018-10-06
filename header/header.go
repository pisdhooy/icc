package header

import (
	"os"

	"github.com/pisdhooy/fsutil"
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
	dateTime.Year = fsutil.ReadBytesShort(file)
	dateTime.Month = fsutil.ReadBytesShort(file)
	dateTime.Day = fsutil.ReadBytesShort(file)
	dateTime.Hours = fsutil.ReadBytesShort(file)
	dateTime.Minutes = fsutil.ReadBytesShort(file)
	dateTime.Seconds = fsutil.ReadBytesShort(file)
}

func (header *Header) Parse(file *os.File) {
	dateTimeObject := NewDateTime()
	illuminantObject := NewVec3_32()

	header.Size = fsutil.ReadBytesLong(file)
	header.Cmmid = fsutil.ReadBytesString(file, 4)
	header.Version = fsutil.ReadBytesLong(file)
	header.DeviceClass = fsutil.ReadBytesString(file, 4)
	header.ColorSpace = fsutil.ReadBytesString(file, 4)
	header.Pcs = fsutil.ReadBytesString(file, 4)

	dateTimeObject.Parse(file)

	header.DateTime = dateTimeObject

	header.Magic = fsutil.ReadBytesString(file, 4)
	header.Platform = fsutil.ReadBytesString(file, 4)
	header.Flags = fsutil.ReadBytesLong(file)

	header.Manufacturer = fsutil.ReadBytesString(file, 4)

	header.Model = fsutil.ReadBytesLong(file)
	header.Attributes = fsutil.ReadBytesLongLong(file)

	header.RenderingIntent = fsutil.ReadBytesLong(file)

	illuminantObject.Parse(file)

	header.Illuminant = illuminantObject
	header.Creator = fsutil.ReadBytesString(file, 4)

	header.ProfileId = fsutil.ReadBytesNInt(file, 16)
	header.Reserved = fsutil.ReadBytesNInt(file, 28)
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
	vec3_32.X = fsutil.ReadBytesLong(file)
	vec3_32.Y = fsutil.ReadBytesLong(file)
	vec3_32.Z = fsutil.ReadBytesLong(file)
}
