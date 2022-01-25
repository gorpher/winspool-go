//go:build windows
// +build windows

package printer

import (
	"reflect"
	"syscall"
	"unsafe"
)

const utf16StringMaxBytes = 1024

func utf16PtrToStringSize(s *uint16, bytes uint32) string {
	if s == nil {
		return ""
	}

	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(s)),
		Len:  int(bytes / 2),
		Cap:  int(bytes / 2),
	}
	c := *(*[]uint16)(unsafe.Pointer(&hdr))

	return syscall.UTF16ToString(c)
}

func utf16PtrToString(s *uint16) string {
	return utf16PtrToStringSize(s, utf16StringMaxBytes)
}

func newBuffer(buf []byte) uintptr {
	return uintptr(unsafe.Pointer(&buf[0])) //nolint
}

func newBufferSize(bufsize *int) uintptr {
	return uintptr(unsafe.Pointer(bufsize)) //nolint
}

func IntPtr(n int32) uintptr {
	return uintptr(n)
}

func utf16PtrFromString(s string) *uint16 {
	u, _ := syscall.UTF16PtrFromString(s)
	return u
}

func utf16FromString(s string) []uint16 {
	u, _ := syscall.UTF16FromString(s)
	return u
}
