// Code generated by "stringer -struct printer"; DO NOT EDIT.

package printer

func (p *PRINTER_INFO_2) GetSepFile() string {
	return utf16PtrToString(p.pSepFile)
}
func (p *PRINTER_INFO_2) GetServerName() string {
	return utf16PtrToString(p.pServerName)
}
func (p *PRINTER_INFO_2) GetPrinterName() string {
	return utf16PtrToString(p.pPrinterName)
}
func (p *PRINTER_INFO_2) GetPortName() string {
	return utf16PtrToString(p.pPortName)
}
func (p *PRINTER_INFO_2) GetDriverName() string {
	return utf16PtrToString(p.pDriverName)
}
func (p *PRINTER_INFO_2) GetComment() string {
	return utf16PtrToString(p.pComment)
}
func (p *PRINTER_INFO_2) GetPrintProcessor() string {
	return utf16PtrToString(p.pPrintProcessor)
}
func (p *PRINTER_INFO_2) GetShareName() string {
	return utf16PtrToString(p.pShareName)
}
func (p *PRINTER_INFO_2) GetLocation() string {
	return utf16PtrToString(p.pLocation)
}
func (p *PRINTER_INFO_2) GetDatatype() string {
	return utf16PtrToString(p.pDatatype)
}
func (p *PRINTER_INFO_2) GetParameters() string {
	return utf16PtrToString(p.pParameters)
}
func (p *PRINTER_INFO_2) GetPrinter() Printer {
	return Printer{
		SepFile:         p.GetSepFile(),
		Priority:        p.priority,
		ServerName:      p.GetServerName(),
		PrinterName:     p.GetPrinterName(),
		PortName:        p.GetPortName(),
		Attributes:      p.attributes,
		AveragePPM:      p.averagePPM,
		DriverName:      p.GetDriverName(),
		Comment:         p.GetComment(),
		PrintProcessor:  p.GetPrintProcessor(),
		Status:          p.status,
		CJobs:           p.cJobs,
		StartTime:       p.startTime,
		UntilTime:       p.untilTime,
		ShareName:       p.GetShareName(),
		Location:        p.GetLocation(),
		Datatype:        p.GetDatatype(),
		Parameters:      p.GetParameters(),
		DefaultPriority: p.defaultPriority,
	}
}
func (p *PRINTER_INFO_5) GetPrinterName() string {
	return utf16PtrToString(p.pPrinterName)
}
func (p *PRINTER_INFO_5) GetPortName() string {
	return utf16PtrToString(p.pPortName)
}
func (p *PRINTER_INFO_5) GetPrinter() Printer {
	return Printer{
		PrinterName:              p.GetPrinterName(),
		PortName:                 p.GetPortName(),
		Attributes:               p.attributes,
		DeviceNotSelectedTimeout: p.deviceNotSelectedTimeout,
		TransmissionRetryTimeout: p.transmissionRetryTimeout,
	}
}

type Printer struct {
	SepFile                  string
	PrinterName              string
	Datatype                 string
	ShareName                string
	DriverName               string
	UntilTime                uint32
	DefaultPriority          uint32
	StartTime                uint32
	TransmissionRetryTimeout uint32
	DeviceNotSelectedTimeout uint32
	PortName                 string
	Comment                  string
	PrintProcessor           string
	CJobs                    uint32
	AveragePPM               uint32
	Location                 string
	Parameters               string
	Priority                 uint32
	Attributes               uint32
	ServerName               string
	Status                   uint32
}
