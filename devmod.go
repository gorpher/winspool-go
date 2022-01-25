package printer

import (
	"fmt"
	"strings"
)

// DEVMODE constants.
const (
	CCHDEVICENAME = 32
	CCHFORMNAME   = 32

	DM_SPECVERSION uint16 = 0x0401
	DM_COPY        uint32 = 2
	DM_MODIFY      uint32 = 8

	DM_ORIENTATION        = 0x00000001
	DM_PAPERSIZE          = 0x00000002
	DM_PAPERLENGTH        = 0x00000004
	DM_PAPERWIDTH         = 0x00000008
	DM_SCALE              = 0x00000010
	DM_POSITION           = 0x00000020
	DM_NUP                = 0x00000040
	DM_DISPLAYORIENTATION = 0x00000080
	DM_COPIES             = 0x00000100
	DM_DEFAULTSOURCE      = 0x00000200
	DM_PRINTQUALITY       = 0x00000400
	DM_COLOR              = 0x00000800
	DM_DUPLEX             = 0x00001000
	DM_YRESOLUTION        = 0x00002000
	DM_TTOPTION           = 0x00004000
	DM_COLLATE            = 0x00008000
	DM_FORMNAME           = 0x00010000
	DM_LOGPIXELS          = 0x00020000
	DM_BITSPERPEL         = 0x00040000
	DM_PELSWIDTH          = 0x00080000
	DM_PELSHEIGHT         = 0x00100000
	DM_DISPLAYFLAGS       = 0x00200000
	DM_DISPLAYFREQUENCY   = 0x00400000
	DM_ICMMETHOD          = 0x00800000
	DM_ICMINTENT          = 0x01000000
	DM_MEDIATYPE          = 0x02000000
	DM_DITHERTYPE         = 0x04000000
	DM_PANNINGWIDTH       = 0x08000000
	DM_PANNINGHEIGHT      = 0x10000000
	DM_DISPLAYFIXEDOUTPUT = 0x20000000

	DMORIENT_PORTRAIT  int16 = 1
	DMORIENT_LANDSCAPE int16 = 2

	DMCOLOR_MONOCHROME int16 = 1
	DMCOLOR_COLOR      int16 = 2

	DMDUP_SIMPLEX    int16 = 1
	DMDUP_VERTICAL   int16 = 2
	DMDUP_HORIZONTAL int16 = 3

	DMCOLLATE_FALSE int16 = 0
	DMCOLLATE_TRUE  int16 = 1

	DMNUP_SYSTEM uint32 = 1
	DMNUP_ONEUP  uint32 = 2
)

// DEVMODE struct.
type DevMode struct {
	// WCHAR dmDeviceName[CCHDEVICENAME]
	dmDeviceName, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ uint16

	dmSpecVersion   uint16
	dmDriverVersion uint16
	dmSize          uint16
	dmDriverExtra   uint16
	dmFields        uint32

	dmOrientation   int16
	dmPaperSize     int16
	dmPaperLength   int16
	dmPaperWidth    int16
	dmScale         int16
	dmCopies        int16
	dmDefaultSource int16
	dmPrintQuality  int16
	dmColor         int16
	dmDuplex        int16
	dmYResolution   int16
	dmTTOption      int16
	dmCollate       int16
	// WCHAR dmFormName[CCHFORMNAME]
	dmFormName, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ uint16

	dmLogPixels        int16
	dmBitsPerPel       uint16
	dmPelsWidth        uint16
	dmPelsHeight       uint16
	dmNup              uint32
	dmDisplayFrequency uint32
	dmICMMethod        uint32
	dmICMIntent        uint32
	dmMediaType        uint32
	dmDitherType       uint32
	dmReserved1        uint32
	dmReserved2        uint32
	dmPanningWidth     uint32
	dmPanningHeight    uint32
}

func (dm *DevMode) String() string {
	s := []string{
		fmt.Sprintf("device name: %s", dm.GetDeviceName()),
		fmt.Sprintf("spec version: %d", dm.dmSpecVersion),
	}
	if dm.dmFields&DM_ORIENTATION != 0 {
		s = append(s, fmt.Sprintf("orientation: %d", dm.dmOrientation))
	}
	if dm.dmFields&DM_PAPERSIZE != 0 {
		s = append(s, fmt.Sprintf("paper size: %d", dm.dmPaperSize))
	}
	if dm.dmFields&DM_PAPERLENGTH != 0 {
		s = append(s, fmt.Sprintf("paper length: %d", dm.dmPaperLength))
	}
	if dm.dmFields&DM_PAPERWIDTH != 0 {
		s = append(s, fmt.Sprintf("paper width: %d", dm.dmPaperWidth))
	}
	if dm.dmFields&DM_SCALE != 0 {
		s = append(s, fmt.Sprintf("scale: %d", dm.dmScale))
	}
	if dm.dmFields&DM_COPIES != 0 {
		s = append(s, fmt.Sprintf("copies: %d", dm.dmCopies))
	}
	if dm.dmFields&DM_DEFAULTSOURCE != 0 {
		s = append(s, fmt.Sprintf("default source: %d", dm.dmDefaultSource))
	}
	if dm.dmFields&DM_PRINTQUALITY != 0 {
		s = append(s, fmt.Sprintf("print quality: %d", dm.dmPrintQuality))
	}
	if dm.dmFields&DM_COLOR != 0 {
		s = append(s, fmt.Sprintf("color: %d", dm.dmColor))
	}
	if dm.dmFields&DM_DUPLEX != 0 {
		s = append(s, fmt.Sprintf("duplex: %d", dm.dmDuplex))
	}
	if dm.dmFields&DM_YRESOLUTION != 0 {
		s = append(s, fmt.Sprintf("y-resolution: %d", dm.dmYResolution))
	}
	if dm.dmFields&DM_TTOPTION != 0 {
		s = append(s, fmt.Sprintf("TT option: %d", dm.dmTTOption))
	}
	if dm.dmFields&DM_COLLATE != 0 {
		s = append(s, fmt.Sprintf("collate: %d", dm.dmCollate))
	}
	if dm.dmFields&DM_FORMNAME != 0 {
		s = append(s, fmt.Sprintf("formname: %s", utf16PtrToString(&dm.dmFormName)))
	}
	if dm.dmFields&DM_LOGPIXELS != 0 {
		s = append(s, fmt.Sprintf("log pixels: %d", dm.dmLogPixels))
	}
	if dm.dmFields&DM_BITSPERPEL != 0 {
		s = append(s, fmt.Sprintf("bits per pel: %d", dm.dmBitsPerPel))
	}
	if dm.dmFields&DM_PELSWIDTH != 0 {
		s = append(s, fmt.Sprintf("pels width: %d", dm.dmPelsWidth))
	}
	if dm.dmFields&DM_PELSHEIGHT != 0 {
		s = append(s, fmt.Sprintf("pels height: %d", dm.dmPelsHeight))
	}
	if dm.dmFields&DM_NUP != 0 {
		s = append(s, fmt.Sprintf("display flags: %d", dm.dmNup))
	}
	if dm.dmFields&DM_DISPLAYFREQUENCY != 0 {
		s = append(s, fmt.Sprintf("display frequency: %d", dm.dmDisplayFrequency))
	}
	if dm.dmFields&DM_ICMMETHOD != 0 {
		s = append(s, fmt.Sprintf("ICM method: %d", dm.dmICMMethod))
	}
	if dm.dmFields&DM_ICMINTENT != 0 {
		s = append(s, fmt.Sprintf("ICM intent: %d", dm.dmICMIntent))
	}
	if dm.dmFields&DM_DITHERTYPE != 0 {
		s = append(s, fmt.Sprintf("dither type: %d", dm.dmDitherType))
	}
	if dm.dmFields&DM_PANNINGWIDTH != 0 {
		s = append(s, fmt.Sprintf("panning width: %d", dm.dmPanningWidth))
	}
	if dm.dmFields&DM_PANNINGHEIGHT != 0 {
		s = append(s, fmt.Sprintf("panning height: %d", dm.dmPanningHeight))
	}
	return strings.Join(s, ", ")
}

func (dm *DevMode) GetDeviceName() string {
	return utf16PtrToStringSize(&dm.dmDeviceName, CCHDEVICENAME*2)
}

func (dm *DevMode) GetOrientation() (int16, bool) {
	return dm.dmOrientation, dm.dmFields&DM_ORIENTATION != 0
}

func (dm *DevMode) SetOrientation(orientation int16) {
	dm.dmOrientation = orientation
	dm.dmFields |= DM_ORIENTATION
}

func (dm *DevMode) GetPaperSize() (int16, bool) {
	return dm.dmPaperSize, dm.dmFields&DM_PAPERSIZE != 0
}

func (dm *DevMode) SetPaperSize(paperSize int16) {
	dm.dmPaperSize = paperSize
	dm.dmFields |= DM_PAPERSIZE
}

func (dm *DevMode) ClearPaperSize() {
	dm.dmFields &^= DM_PAPERSIZE
}

func (dm *DevMode) GetPaperLength() (int16, bool) {
	return dm.dmPaperLength, dm.dmFields&DM_PAPERLENGTH != 0
}

func (dm *DevMode) SetPaperLength(length int16) {
	dm.dmPaperLength = length
	dm.dmFields |= DM_PAPERLENGTH
}

func (dm *DevMode) ClearPaperLength() {
	dm.dmFields &^= DM_PAPERLENGTH
}

func (dm *DevMode) GetPaperWidth() (int16, bool) {
	return dm.dmPaperWidth, dm.dmFields&DM_PAPERWIDTH != 0
}

func (dm *DevMode) SetPaperWidth(width int16) {
	dm.dmPaperWidth = width
	dm.dmFields |= DM_PAPERWIDTH
}

func (dm *DevMode) ClearPaperWidth() {
	dm.dmFields &^= DM_PAPERWIDTH
}

func (dm *DevMode) GetCopies() (int16, bool) {
	return dm.dmCopies, dm.dmFields&DM_COPIES != 0
}

func (dm *DevMode) SetCopies(copies int16) {
	dm.dmCopies = copies
	dm.dmFields |= DM_COPIES
}

func (dm *DevMode) GetColor() (int16, bool) {
	return dm.dmColor, dm.dmFields&DM_COLOR != 0
}

func (dm *DevMode) SetColor(color int16) {
	dm.dmColor = color
	dm.dmFields |= DM_COLOR
}

func (dm *DevMode) GetDuplex() (int16, bool) {
	return dm.dmDuplex, dm.dmFields&DM_DUPLEX != 0
}

func (dm *DevMode) SetDuplex(duplex int16) {
	dm.dmDuplex = duplex
	dm.dmFields |= DM_DUPLEX
}

func (dm *DevMode) GetCollate() (int16, bool) {
	return dm.dmCollate, dm.dmFields&DM_COLLATE != 0
}

func (dm *DevMode) SetCollate(collate int16) {
	dm.dmCollate = collate
	dm.dmFields |= DM_COLLATE
}
