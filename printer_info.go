package printer

// PRINTER_INFO_2 attribute values
const (
	PRINTER_ATTRIBUTE_QUEUED            uint32 = 0x00000001
	PRINTER_ATTRIBUTE_DIRECT            uint32 = 0x00000002
	PRINTER_ATTRIBUTE_DEFAULT           uint32 = 0x00000004
	PRINTER_ATTRIBUTE_SHARED            uint32 = 0x00000008
	PRINTER_ATTRIBUTE_NETWORK           uint32 = 0x00000010
	PRINTER_ATTRIBUTE_HIDDEN            uint32 = 0x00000020
	PRINTER_ATTRIBUTE_LOCAL             uint32 = 0x00000040
	PRINTER_ATTRIBUTE_ENABLE_DEVQ       uint32 = 0x00000080
	PRINTER_ATTRIBUTE_KEEPPRINTEDJOBS   uint32 = 0x00000100
	PRINTER_ATTRIBUTE_DO_COMPLETE_FIRST uint32 = 0x00000200
	PRINTER_ATTRIBUTE_WORK_OFFLINE      uint32 = 0x00000400
	PRINTER_ATTRIBUTE_ENABLE_BIDI       uint32 = 0x00000800
	PRINTER_ATTRIBUTE_RAW_ONLY          uint32 = 0x00001000
	PRINTER_ATTRIBUTE_PUBLISHED         uint32 = 0x00002000
)

// PRINTER_INFO_2 status values.
const (
	PRINTER_STATUS_PAUSED               uint32 = 0x00000001
	PRINTER_STATUS_ERROR                uint32 = 0x00000002
	PRINTER_STATUS_PENDING_DELETION     uint32 = 0x00000004
	PRINTER_STATUS_PAPER_JAM            uint32 = 0x00000008
	PRINTER_STATUS_PAPER_OUT            uint32 = 0x00000010
	PRINTER_STATUS_MANUAL_FEED          uint32 = 0x00000020
	PRINTER_STATUS_PAPER_PROBLEM        uint32 = 0x00000040
	PRINTER_STATUS_OFFLINE              uint32 = 0x00000080
	PRINTER_STATUS_IO_ACTIVE            uint32 = 0x00000100
	PRINTER_STATUS_BUSY                 uint32 = 0x00000200
	PRINTER_STATUS_PRINTING             uint32 = 0x00000400
	PRINTER_STATUS_OUTPUT_BIN_FULL      uint32 = 0x00000800
	PRINTER_STATUS_NOT_AVAILABLE        uint32 = 0x00001000
	PRINTER_STATUS_WAITING              uint32 = 0x00002000
	PRINTER_STATUS_PROCESSING           uint32 = 0x00004000
	PRINTER_STATUS_INITIALIZING         uint32 = 0x00008000
	PRINTER_STATUS_WARMING_UP           uint32 = 0x00010000
	PRINTER_STATUS_TONER_LOW            uint32 = 0x00020000
	PRINTER_STATUS_NO_TONER             uint32 = 0x00040000
	PRINTER_STATUS_PAGE_PUNT            uint32 = 0x00080000
	PRINTER_STATUS_USER_INTERVENTION    uint32 = 0x00100000
	PRINTER_STATUS_OUT_OF_MEMORY        uint32 = 0x00200000
	PRINTER_STATUS_DOOR_OPEN            uint32 = 0x00400000
	PRINTER_STATUS_SERVER_UNKNOWN       uint32 = 0x00800000
	PRINTER_STATUS_POWER_SAVE           uint32 = 0x01000000
	PRINTER_STATUS_SERVER_OFFLINE       uint32 = 0x02000000
	PRINTER_STATUS_DRIVER_UPDATE_NEEDED uint32 = 0x04000000
)

//go:generate go run github.com/gorpher/winspool-go/cmd/mkwinsyscall_struct -struct printer

// PRINTER_INFO_2 struct.
// sys_struct: PRINTER
type PRINTER_INFO_2 struct {
	pServerName         *uint16
	pPrinterName        *uint16
	pShareName          *uint16
	pPortName           *uint16
	pDriverName         *uint16
	pComment            *uint16
	pLocation           *uint16
	pDevMode            *DevMode
	pSepFile            *uint16
	pPrintProcessor     *uint16
	pDatatype           *uint16
	pParameters         *uint16
	pSecurityDescriptor uintptr
	attributes          uint32
	priority            uint32
	defaultPriority     uint32
	startTime           uint32
	untilTime           uint32
	status              uint32
	cJobs               uint32
	averagePPM          uint32
}

type PRINTER_INFO_5 struct {
	pPrinterName             *uint16
	pPortName                *uint16
	attributes               uint32
	deviceNotSelectedTimeout uint32
	transmissionRetryTimeout uint32
}

type PrinterStatusType string

const (
	PrinterStatusIdle       PrinterStatusType = "IDLE"
	PrinterStatusProcessing PrinterStatusType = "PROCESSING"
	PrinterStatusStopped    PrinterStatusType = "STOPPED"
)

func (p *Printer) GetStatus() map[string]bool {
	var status = map[string]bool{}
	if p.Status&PRINTER_STATUS_PAUSED > 0 {
		status["is_paused"] = true
	}
	if p.Status&PRINTER_STATUS_ERROR > 0 {
		status["is_error"] = true
	}
	if p.Status&PRINTER_STATUS_PENDING_DELETION > 0 {
		status["is_pending_deletion"] = true
	}

	if p.Status&PRINTER_STATUS_PAPER_JAM > 0 {
		status["is_paper_jam"] = true
	}

	if p.Status&PRINTER_STATUS_PAPER_OUT > 0 {
		status["js_paper_out"] = true
	}

	if p.Status&PRINTER_STATUS_MANUAL_FEED > 0 {
		status["is_manual_feed"] = true
	}

	if p.Status&PRINTER_STATUS_PAPER_PROBLEM > 0 {
		status["is_paper_problem"] = true
	}

	if p.Status&PRINTER_STATUS_OFFLINE > 0 {
		status["is_offline"] = true
	}

	if p.Status&PRINTER_STATUS_IO_ACTIVE > 0 {
		status["is_io_active"] = true
	}

	if p.Status&PRINTER_STATUS_BUSY > 0 {
		status["is_busy"] = true
	}

	if p.Status&PRINTER_STATUS_PRINTING > 0 {
		status["is_printing"] = true
	}

	if p.Status&PRINTER_STATUS_OUTPUT_BIN_FULL > 0 {
		status["is_output_bin_full"] = true
	}

	if p.Status&PRINTER_STATUS_NOT_AVAILABLE > 0 {
		status["is_out_available"] = true
	}

	if p.Status&PRINTER_STATUS_WAITING > 0 {
		status["is_waiting"] = true
	}

	if p.Status&PRINTER_STATUS_PROCESSING > 0 {
		status["is_processing"] = true
	}

	if p.Status&PRINTER_STATUS_INITIALIZING > 0 {
		status["is_initializing"] = true
	}

	if p.Status&PRINTER_STATUS_WARMING_UP > 0 {
		status["is_warming_up"] = true
	}

	if p.Status&PRINTER_STATUS_TONER_LOW > 0 {
		status["is_toner_low"] = true
	}

	if p.Status&PRINTER_STATUS_NO_TONER > 0 {
		status["is_no_toner"] = true
	}

	if p.Status&PRINTER_STATUS_PAGE_PUNT > 0 {
		status["is_page_punt"] = true
	}

	if p.Status&PRINTER_STATUS_USER_INTERVENTION > 0 {
		status["is_user_intervention"] = true
	}

	if p.Status&PRINTER_STATUS_OUT_OF_MEMORY > 0 {
		status["is_out_of_memory"] = true
	}

	if p.Status&PRINTER_STATUS_DOOR_OPEN > 0 {
		status["is_door_open"] = true
	}

	if p.Status&PRINTER_STATUS_SERVER_UNKNOWN > 0 {
		status["is_server_unknown"] = true
	}

	if p.Status&PRINTER_STATUS_POWER_SAVE > 0 {
		status["is_power_save"] = true
	}

	if p.Status&PRINTER_STATUS_SERVER_OFFLINE > 0 {
		status["is_server_offline"] = true
	}

	if p.Status&PRINTER_STATUS_DRIVER_UPDATE_NEEDED > 0 {
		status["is_driver_update_needed"] = true
	}
	return status
}

func (p *Printer) GetPrinterStatus() PrinterStatusType {
	if p.Status&(PRINTER_STATUS_PRINTING|PRINTER_STATUS_PROCESSING) != 0 {
		return PrinterStatusProcessing
	}
	if p.Status&PRINTER_STATUS_PAUSED != 0 {
		return PrinterStatusStopped
	}
	if p.Status&PRINTER_STATUS_ERROR != 0 {
		return PrinterStatusStopped
	}
	if p.Status&PRINTER_STATUS_PENDING_DELETION != 0 {
		return PrinterStatusStopped
	}
	if p.Status&PRINTER_STATUS_PAPER_JAM != 0 {
		return PrinterStatusStopped
	}
	if p.Status&PRINTER_STATUS_PAPER_OUT != 0 {
		return PrinterStatusStopped
	}
	if p.Status&PRINTER_STATUS_MANUAL_FEED != 0 {
		return PrinterStatusStopped
	}
	if p.Status&PRINTER_STATUS_PAPER_PROBLEM != 0 {
		return PrinterStatusStopped
	}
	return PrinterStatusIdle
}

func (p *Printer) GetAttributes() map[string]bool {
	var attributes = map[string]bool{}
	if p.Attributes&PRINTER_ATTRIBUTE_DIRECT > 0 {
		attributes["is_direct"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_DO_COMPLETE_FIRST > 0 {
		attributes["is_do_complete_first"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_ENABLE_DEVQ > 0 {
		attributes["is_enable_devq"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_HIDDEN > 0 {
		attributes["is_hidden"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_DEFAULT > 0 {
		attributes["is_default"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_WORK_OFFLINE > 0 {
		attributes["is_work_offline"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_KEEPPRINTEDJOBS > 0 {
		attributes["is_keep_printed_jobs"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_LOCAL > 0 {
		attributes["is_local"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_ENABLE_BIDI > 0 {
		attributes["is_enable_bidi"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_NETWORK > 0 {
		attributes["is_network"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_PUBLISHED > 0 {
		attributes["is_published"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_QUEUED > 0 {
		attributes["is_queued"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_RAW_ONLY > 0 {
		attributes["is_raw_only"] = true
	}
	if p.Attributes&PRINTER_ATTRIBUTE_SHARED > 0 {
		attributes["is_shared"] = true
	}
	return attributes
}
