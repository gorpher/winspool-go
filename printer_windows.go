package printer

import (
	"syscall"
	"unsafe"
)

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output winspool_windows.go printer_windows.go
//sys	getDefaultPrinter(buf *uint16, bufN *uint32) (err error)   = winspool.GetDefaultPrinterW
//sys	closePrinter(h HANDLE) (err error) = winspool.ClosePrinter
//sys	openPrinter(name *uint16, h *HANDLE, defaults uintptr) (err error) = winspool.OpenPrinterW
//sys	startDocPrinter(h HANDLE, level uint32, docinfo *DOC_INFO_1) (err error) = winspool.StartDocPrinterW
//sys	endDocPrinter(h HANDLE) (err error) = winspool.EndDocPrinter
//sys	writePrinter(h HANDLE, buf *byte, bufN uint32, written *uint32) (err error) = winspool.WritePrinter
//sys	startPagePrinter(h HANDLE) (err error) = winspool.StartPagePrinter
//sys	endPagePrinter(h HANDLE) (err error) = winspool.EndPagePrinter
//sys	enumPrinters(flags uint32, name *uint16, level uint32, buf *byte, bufN uint32, needed *uint32, returned *uint32) (err error) = winspool.EnumPrintersW
//sys	getPrinterDriver(h HANDLE, env *uint16, level uint32, di *byte, n uint32, needed *uint32) (err error) = winspool.GetPrinterDriverW
//sys	enumJobs(h HANDLE, firstJob uint32, noJobs uint32, level uint32, buf *byte, bufN uint32, bytesNeeded *uint32, jobsReturned *uint32) (err error) = winspool.EnumJobsW
//sys   setJob(h HANDLE, jobId uint32, level uint32, buf uintptr, command uint32) (err error)  = winspool.SetJobW
//sys   getJob(h HANDLE, jobId uint32, level uint32 ,buf *byte, cbBuf uint32, pcbNeeded *uint32) (err error) = winspool.GetJobW

const (
	PRINTER_ENUM_DEFAULT     = 0x00000001
	PRINTER_ENUM_LOCAL       = 0x00000002
	PRINTER_ENUM_CONNECTIONS = 0x00000004
	PRINTER_ENUM_FAVORITE    = 0x00000004
	PRINTER_ENUM_NAME        = 0x00000008
	PRINTER_ENUM_REMOTE      = 0x00000010
	PRINTER_ENUM_SHARED      = 0x00000020
	PRINTER_ENUM_NETWORK     = 0x00000040
	PRINTER_ENUM_EXPAND      = 0x00004000
	PRINTER_ENUM_CONTAINER   = 0x00008000
	PRINTER_ENUM_ICONMASK    = 0x00ff0000
	PRINTER_ENUM_ICON1       = 0x00010000
	PRINTER_ENUM_ICON2       = 0x00020000
	PRINTER_ENUM_ICON3       = 0x00040000
	PRINTER_ENUM_ICON4       = 0x00080000
	PRINTER_ENUM_ICON5       = 0x00100000
	PRINTER_ENUM_ICON6       = 0x00200000
	PRINTER_ENUM_ICON7       = 0x00400000
	PRINTER_ENUM_ICON8       = 0x00800000
	PRINTER_ENUM_HIDE        = 0x01000000
)
const (
	PRINTER_DRIVER_XPS = 0x00000002
)

type HANDLE syscall.Handle

func OpenPrinter(printerName string) (HANDLE, error) {
	var pPrinterName *uint16
	pPrinterName, err := syscall.UTF16PtrFromString(printerName)
	if err != nil {
		return 0, err
	}
	var hPrinter HANDLE
	err = openPrinter(pPrinterName, &hPrinter, 0)
	if err != nil {
		return 0, err
	}
	return hPrinter, nil
}

func EnumPrinters2() ([]Printer, error) {
	const level = 2
	const flags = PRINTER_ENUM_LOCAL | PRINTER_ENUM_CONNECTIONS
	var needed, returned uint32
	buf := make([]byte, 1)
	err := enumPrinters(flags, nil, level, &buf[0], uint32(len(buf)), &needed, &returned)
	if err != nil {
		if err != syscall.ERROR_INSUFFICIENT_BUFFER {
			return nil, err
		}
		buf = make([]byte, needed)
		err = enumPrinters(flags, nil, level, &buf[0], uint32(len(buf)), &needed, &returned)
		if err != nil {
			return nil, err
		}
	}
	ps := (*[1024]PRINTER_INFO_2)(unsafe.Pointer(&buf[0]))[:returned:returned]
	printers := make([]Printer, returned)
	for i, p := range ps {
		printers[i] = p.GetPrinter()
	}
	return printers, nil
}

func EnumPrinters5() ([]Printer, error) {
	const level = 5
	const flags = PRINTER_ENUM_LOCAL | PRINTER_ENUM_CONNECTIONS
	var needed, returned uint32
	buf := make([]byte, 1)
	err := enumPrinters(flags, nil, level, &buf[0], uint32(len(buf)), &needed, &returned)
	if err != nil {
		if err != syscall.ERROR_INSUFFICIENT_BUFFER {
			return nil, err
		}
		buf = make([]byte, needed)
		err = enumPrinters(flags, nil, level, &buf[0], uint32(len(buf)), &needed, &returned)
		if err != nil {
			return nil, err
		}
	}
	ps := (*[1024]PRINTER_INFO_5)(unsafe.Pointer(&buf[0]))[:returned:returned]
	printers := make([]Printer, returned)
	for i, p := range ps {
		printers[i] = p.GetPrinter()
	}
	return printers, nil
}

func GetDefaultPrinter() (string, error) {
	b := make([]uint16, 3)
	n := uint32(len(b))
	err := getDefaultPrinter(&b[0], &n)
	if err != nil {
		if err != syscall.ERROR_INSUFFICIENT_BUFFER {
			return "", err
		}
		b = make([]uint16, n)
		err = getDefaultPrinter(&b[0], &n)
		if err != nil {
			return "", err
		}
	}
	return syscall.UTF16ToString(b), nil

}

func (h HANDLE) ClosePrinter() error {
	return closePrinter(h)
}

// GetJob 获取打印机作业信息
//BOOL GetJob(
//  _In_  HANDLE  hPrinter,
//  _In_  DWORD   JobId,
//  _In_  DWORD   Level,
//  _Out_ LPBYTE  pJob,
//  _In_  DWORD   cbBuf,
//  _Out_ LPDWORD pcbNeeded
//);
func (h HANDLE) GetJob(jobId uint32) (*Job, error) {
	var returned uint32
	buf := make([]byte, 1)
	err := getJob(h, jobId, 1, &buf[0], uint32(len(buf)), &returned)
	if err != syscall.ERROR_INSUFFICIENT_BUFFER {
		return nil, err
	}
	buf = make([]byte, returned)
	err = getJob(h, jobId, 1, &buf[0], uint32(len(buf)), &returned)
	if err != nil {
		return nil, err
	}
	job := *(*JOB_INFO_1)(unsafe.Pointer(&buf[0]))
	j := job.GetJob()
	return &j, nil
}

// EnumJobs 获取打印机作业列表
//BOOL EnumJobs(
//  _In_  HANDLE  hPrinter,
//  _In_  DWORD   FirstJob,
//  _In_  DWORD   NoJobs,
//  _In_  DWORD   Level,
//  _Out_ LPBYTE  pJob,
//  _In_  DWORD   cbBuf,
//  _Out_ LPDWORD pcbNeeded,
//  _Out_ LPDWORD pcReturned
//);
func (h *HANDLE) EnumJobs1() ([]Job, error) {
	var bytesNeeded, jobsReturned uint32
	buf := make([]byte, 1)
	for {
		err := enumJobs(*h, 0, 255, 1, &buf[0], uint32(len(buf)), &bytesNeeded, &jobsReturned)
		if err == nil {
			break
		}
		if err != syscall.ERROR_INSUFFICIENT_BUFFER {
			return nil, err
		}
		if bytesNeeded <= uint32(len(buf)) {
			return nil, err
		}
		buf = make([]byte, bytesNeeded)
	}
	if jobsReturned <= 0 {
		return nil, nil
	}
	jobs := make([]Job, jobsReturned)
	ji1 := (*[4096]JOB_INFO_1)(unsafe.Pointer(&buf[0]))[:jobsReturned:jobsReturned]
	for i, j := range ji1 {
		jobs[i] = j.GetJob()
	}
	return jobs, nil
}

func (h HANDLE) Driver() (Driver, error) {
	const level = 8
	var needed uint32
	b := make([]byte, 1024*10)
	var driver = Driver{}
	err := getPrinterDriver(h, nil, level, &b[0], uint32(len(b)), &needed)
	if err != syscall.ERROR_INSUFFICIENT_BUFFER {
		return driver, err
	}
	if needed <= uint32(len(b)) {
		return driver, err
	}
	b = make([]byte, needed)
	err = getPrinterDriver(h, nil, level, &b[0], uint32(len(b)), &needed)
	if err != nil {
		return driver, err
	}
	di := (*DRIVER_INFO_8)(unsafe.Pointer(&b[0]))
	driver = di.GetDriver()
	return driver, nil
}

func (h HANDLE) Print(name, datatype string, b []byte) error {
	di, err := h.Driver()
	if err != nil {
		return err
	}
	// See https://support.microsoft.com/en-us/help/2779300/v4-print-drivers-using-raw-mode-to-send-pcl-postscript-directly-to-the
	// for details.
	if datatype == "" {
		datatype = "RAW"
	}
	if di.PrinterDriverAttributes&PRINTER_DRIVER_XPS != 0 {
		datatype = "XPS_PASS"
	}
	err = h.startDocument(name, datatype)
	if err != nil {
		return err
	}
	defer endDocPrinter(h)
	err = startPagePrinter(h)
	if err != nil {
		return err
	}
	defer endPagePrinter(h)
	var written uint32
	return writePrinter(h, &b[0], uint32(len(b)), &written)
}

func (h HANDLE) startDocument(name, datatype string) error {
	d := DOC_INFO_1{
		DocName:    utf16PtrFromString(name),
		OutputFile: nil,
		Datatype:   utf16PtrFromString(datatype),
	}
	return startDocPrinter(h, 1, &d)
}

// ReleaseJob 释放作业
func (h HANDLE) ReleaseJob(jobId uint32) error {
	job, err := h.GetJob(jobId)
	if err != nil {
		return err
	}
	if job.Status&JOB_STATUS_RETAINED != 0 {
		err = setJob(h, jobId, 0, 0, JOB_CONTROL_RELEASE)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h HANDLE) SetJob1(jobId uint32, ji1 *JOB_INFO_1) error {
	return setJob(h, jobId, 1, uintptr(unsafe.Pointer(ji1)), 0)
}
