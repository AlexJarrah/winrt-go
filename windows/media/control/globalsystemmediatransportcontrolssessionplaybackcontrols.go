// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package control

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureGlobalSystemMediaTransportControlsSessionPlaybackControls string = "rc(Windows.Media.Control.GlobalSystemMediaTransportControlsSessionPlaybackControls;{6501a3e6-bc7a-503a-bb1b-68f158f3fb03})"

type GlobalSystemMediaTransportControlsSessionPlaybackControls struct {
	ole.IUnknown
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPlayEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsPlayEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPauseEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsPauseEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsStopEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsStopEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsRecordEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsRecordEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsFastForwardEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsFastForwardEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsRewindEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsRewindEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsNextEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsNextEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPreviousEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsPreviousEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsChannelUpEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsChannelUpEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsChannelDownEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsChannelDownEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPlayPauseToggleEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsPlayPauseToggleEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsShuffleEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsShuffleEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsRepeatEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsRepeatEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPlaybackRateEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsPlaybackRateEnabled()
}

func (impl *GlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPlaybackPositionEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionPlaybackControls)(unsafe.Pointer(itf))
	return v.GetIsPlaybackPositionEnabled()
}

const GUIDiGlobalSystemMediaTransportControlsSessionPlaybackControls string = "6501a3e6-bc7a-503a-bb1b-68f158f3fb03"
const SignatureiGlobalSystemMediaTransportControlsSessionPlaybackControls string = "{6501a3e6-bc7a-503a-bb1b-68f158f3fb03}"

type iGlobalSystemMediaTransportControlsSessionPlaybackControls struct {
	ole.IInspectable
}

type iGlobalSystemMediaTransportControlsSessionPlaybackControlsVtbl struct {
	ole.IInspectableVtbl

	GetIsPlayEnabled             uintptr
	GetIsPauseEnabled            uintptr
	GetIsStopEnabled             uintptr
	GetIsRecordEnabled           uintptr
	GetIsFastForwardEnabled      uintptr
	GetIsRewindEnabled           uintptr
	GetIsNextEnabled             uintptr
	GetIsPreviousEnabled         uintptr
	GetIsChannelUpEnabled        uintptr
	GetIsChannelDownEnabled      uintptr
	GetIsPlayPauseToggleEnabled  uintptr
	GetIsShuffleEnabled          uintptr
	GetIsRepeatEnabled           uintptr
	GetIsPlaybackRateEnabled     uintptr
	GetIsPlaybackPositionEnabled uintptr
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) VTable() *iGlobalSystemMediaTransportControlsSessionPlaybackControlsVtbl {
	return (*iGlobalSystemMediaTransportControlsSessionPlaybackControlsVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPlayEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsPlayEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPauseEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsPauseEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsStopEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsStopEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsRecordEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsRecordEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsFastForwardEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsFastForwardEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsRewindEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsRewindEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsNextEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsNextEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPreviousEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsPreviousEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsChannelUpEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsChannelUpEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsChannelDownEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsChannelDownEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPlayPauseToggleEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsPlayPauseToggleEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsShuffleEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsShuffleEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsRepeatEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsRepeatEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPlaybackRateEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsPlaybackRateEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionPlaybackControls) GetIsPlaybackPositionEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsPlaybackPositionEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}
