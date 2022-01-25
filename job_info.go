package printer

import (
	"strings"
	"syscall"
)

// JOB_INFO_1 status values.
const (
	JOB_STATUS_PAUSED            uint32 = 0x00000001
	JOB_STATUS_ERROR             uint32 = 0x00000002
	JOB_STATUS_DELETING          uint32 = 0x00000004
	JOB_STATUS_SPOOLING          uint32 = 0x00000008
	JOB_STATUS_PRINTING          uint32 = 0x00000010
	JOB_STATUS_OFFLINE           uint32 = 0x00000020
	JOB_STATUS_PAPEROUT          uint32 = 0x00000040
	JOB_STATUS_PRINTED           uint32 = 0x00000080
	JOB_STATUS_DELETED           uint32 = 0x00000100
	JOB_STATUS_BLOCKED_DEVQ      uint32 = 0x00000200
	JOB_STATUS_USER_INTERVENTION uint32 = 0x00000400
	JOB_STATUS_RESTART           uint32 = 0x00000800
	JOB_STATUS_COMPLETE          uint32 = 0x00001000
	JOB_STATUS_RETAINED          uint32 = 0x00002000
	JOB_STATUS_RENDERING_LOCALLY uint32 = 0x00004000
)

// SetJob command values.
const (
	JOB_CONTROL_PAUSE             uint32 = 1
	JOB_CONTROL_RESUME            uint32 = 2
	JOB_CONTROL_CANCEL            uint32 = 3
	JOB_CONTROL_RESTART           uint32 = 4
	JOB_CONTROL_DELETE            uint32 = 5
	JOB_CONTROL_SENT_TO_PRINTER   uint32 = 6
	JOB_CONTROL_LAST_PAGE_EJECTED uint32 = 7
	JOB_CONTROL_RETAIN            uint32 = 8
	JOB_CONTROL_RELEASE           uint32 = 9
)

//go:generate go run github.com/gorpher/winspool-go/cmd/mkwinsyscall_struct -struct job

// JOB_INFO_1
//typedef struct _JOB_INFO_1 {
//  DWORD      JobId;
//  LPTSTR     pPrinterName;
//  LPTSTR     pMachineName;
//  LPTSTR     pUserName;
//  LPTSTR     pDocument;
//  LPTSTR     pDatatype;
//  LPTSTR     pStatus;
//  DWORD      Status;
//  DWORD      Priority;
//  DWORD      Position;
//  DWORD      TotalPages;
//  DWORD      PagesPrinted;
//  SYSTEMTIME Submitted;
//} JOB_INFO_1, *PJOB_INFO_1;
type JOB_INFO_1 struct {
	JobID        uint32
	pPrinterName *uint16
	pMachineName *uint16
	pUserName    *uint16
	pDocument    *uint16
	pDatatype    *uint16
	pStatus      uintptr
	Status       uint32
	Priority     uint32
	Position     uint32
	TotalPages   uint32
	PagesPrinted uint32
	Submitted    syscall.Systemtime
}

// JOB_INFO_2
//typedef struct _JOB_INFO_2 {
//  DWORD                JobId;
//  LPTSTR               pPrinterName;
//  LPTSTR               pMachineName;
//  LPTSTR               pUserName;
//  LPTSTR               pDocument;
//  LPTSTR               pNotifyName;
//  LPTSTR               pDatatype;
//  LPTSTR               pPrintProcessor;
//  LPTSTR               pParameters;
//  LPTSTR               pDriverName;
//  LPDEVMODE            pDevMode;
//  LPTSTR               pStatus;
//  PSECURITY_DESCRIPTOR pSecurityDescriptor;
//  DWORD                Status;
//  DWORD                Priority;
//  DWORD                Position;
//  DWORD                StartTime;
//  DWORD                UntilTime;
//  DWORD                TotalPages;
//  DWORD                Size;
//  SYSTEMTIME           Submitted;
//  DWORD                Time;
//  DWORD                PagesPrinted;
//} JOB_INFO_2, *PJOB_INFO_2;
type JOB_INFO_2 struct {
	JobID               uint32
	pPrinterName        *uint16
	pMachineName        *uint16
	pUserName           *uint16
	pDocument           *uint16
	pNotifyName         *uint16
	pDatatype           *uint16
	pPrintProcessor     *uint16
	pParameters         *uint16
	pDriverName         *uint16
	pDevMode            *DevMode
	pSecurityDescriptor uint32
	Status              uint32
	Position            uint32
	Priority            uint32
	StartTime           uint32
	UntilTime           uint32
	Size                uint32
	Submitted           syscall.Systemtime
	Time                uint32
	PagesPrinted        uint32
}

// JOB_INFO_3
//typedef struct _JOB_INFO_3 {
//  DWORD JobId;
//  DWORD NextJobId;
//  DWORD Reserved;
//} JOB_INFO_3, *PJOB_INFO_3;
type JOB_INFO_3 struct {
	JobID     uint32
	NextJobId uint32
	Reserved  uint32
}

func (p *Job) GetStatus() string {
	var status string
	if p.Status == 0 {
		status += "Queue Paused,"
	}
	if p.Status&JOB_STATUS_PRINTING != 0 {
		status += "Printing,"
	}
	if p.Status&JOB_STATUS_PAUSED != 0 {
		status += "Paused,"
	}
	if p.Status&JOB_STATUS_ERROR != 0 {
		status += "Error,"
	}
	if p.Status&JOB_STATUS_DELETING != 0 {
		status += "Deleting,"
	}
	if p.Status&JOB_STATUS_SPOOLING != 0 {
		status += "Spooling,"
	}
	if p.Status&JOB_STATUS_OFFLINE != 0 {
		status += "Printer Offline,"
	}
	if p.Status&JOB_STATUS_PAPEROUT != 0 {
		status += "Out of Paper,"
	}
	if p.Status&JOB_STATUS_PRINTED != 0 {
		status += "Printed,"
	}
	if p.Status&JOB_STATUS_DELETED != 0 {
		status += "Deleted,"
	}
	if p.Status&JOB_STATUS_BLOCKED_DEVQ != 0 {
		status += "Driver Error,"
	}
	if p.Status&JOB_STATUS_USER_INTERVENTION != 0 {
		status += "User Action Required,"
	}
	if p.Status&JOB_STATUS_RESTART != 0 {
		status += "Restarted,"
	}
	if p.Status&JOB_STATUS_COMPLETE != 0 {
		status += "Sent to Printer,"
	}
	if p.Status&JOB_STATUS_RETAINED != 0 {
		status += "Retained,"
	}
	if p.Status&JOB_STATUS_RENDERING_LOCALLY != 0 {
		status += "Rendering on Client,"
	}
	return strings.TrimRight(status, ",")
}
