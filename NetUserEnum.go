package main

import (
	"fmt"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

type USER_INFO_1 struct {
	Usri1_name         *uint16
	Usri1_password     *uint16
	Usri1_password_age uint32
	Usri1_priv         uint32
	Usri1_home_dir     *uint16
	Usri1_comment      *uint16
	Usri1_flags        uint32
	Usri1_script_path  *uint16
}

func ByteToStr(a []uint16) string {
	for i, v := range a {
		if v == 0 {
			a = a[0:i]
			break
		}
	}
	return string(utf16.Decode(a))
}

var (
	dataPointer  uintptr
	entriesRead  uint32
	entriesTotal uint32
	resumeHandle uintptr
)

const (
	NET_API_STATUS_NERR_Success = 0
	MAX_PREFERRED_LENGTH        = 0xFFFFFFFF
)

func main() {
	var (
		netapi32         = syscall.NewLazyDLL("netapi32.dll")
		netuserenum      = netapi32.NewProc("NetUserEnum")
		netapibufferfree = netapi32.NewProc("NetApiBufferFree")
	)
	defer netapibufferfree.Call(dataPointer)
	r1, _, _ := syscall.Syscall9(netuserenum.Addr(), 8,
		uintptr(0), uintptr(1), uintptr(0),
		uintptr(unsafe.Pointer(&dataPointer)),
		uintptr(uint32(MAX_PREFERRED_LENGTH)),
		uintptr(unsafe.Pointer(&entriesRead)),
		uintptr(unsafe.Pointer(&entriesTotal)),
		uintptr(unsafe.Pointer(&resumeHandle)), 0)

	if r1 != NET_API_STATUS_NERR_Success {
		panic(fmt.Errorf("error fetching user entry"))
	} else if dataPointer == uintptr(0) {
		panic(fmt.Errorf("null pointer while fetching entry"))
	}

	var iter = dataPointer
	for i := uint32(0); i < entriesRead; i++ {
		var data = (*USER_INFO_1)(unsafe.Pointer(iter))
		fmt.Printf("%d: %v\n", i+1, ByteToStr((*[4096]uint16)(unsafe.Pointer(data.Usri1_name))[:]))
		iter = uintptr(iter + unsafe.Sizeof(USER_INFO_1{}))
	}
}
