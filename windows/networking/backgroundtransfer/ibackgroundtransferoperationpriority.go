// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package backgroundtransfer

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const GUIDIBackgroundTransferOperationPriority string = "04854327-5254-4b3a-915e-0aa49275c0f9"
const SignatureIBackgroundTransferOperationPriority string = "{04854327-5254-4b3a-915e-0aa49275c0f9}"

type IBackgroundTransferOperationPriority struct {
	ole.IInspectable
}

type IBackgroundTransferOperationPriorityVtbl struct {
	ole.IInspectableVtbl

	GetPriority uintptr
	SetPriority uintptr
}

func (v *IBackgroundTransferOperationPriority) VTable() *IBackgroundTransferOperationPriorityVtbl {
	return (*IBackgroundTransferOperationPriorityVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IBackgroundTransferOperationPriority) GetPriority() (BackgroundTransferPriority, error) {
	var out BackgroundTransferPriority
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetPriority,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out BackgroundTransferPriority
	)

	if hr != 0 {
		return BackgroundTransferPriorityDefault, ole.NewError(hr)
	}

	return out, nil
}

func (v *IBackgroundTransferOperationPriority) SetPriority(value BackgroundTransferPriority) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetPriority,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in BackgroundTransferPriority
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
