// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package foundation

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const GUIDIStringable string = "96369f54-8eb6-48f0-abce-c1b211e627c3"
const SignatureIStringable string = "{96369f54-8eb6-48f0-abce-c1b211e627c3}"

type IStringable struct {
	ole.IInspectable
}

type IStringableVtbl struct {
	ole.IInspectableVtbl

	ToString uintptr
}

func (v *IStringable) VTable() *IStringableVtbl {
	return (*IStringableVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IStringable) ToString() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().ToString,
		uintptr(unsafe.Pointer(v)),        // this
		uintptr(unsafe.Pointer(&outHStr)), // out string
	)

	if hr != 0 {
		return "", ole.NewError(hr)
	}

	out := outHStr.String()
	ole.DeleteHString(outHStr)
	return out, nil
}
