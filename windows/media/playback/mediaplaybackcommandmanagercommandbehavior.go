// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package playback

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const SignatureMediaPlaybackCommandManagerCommandBehavior string = "rc(Windows.Media.Playback.MediaPlaybackCommandManagerCommandBehavior;{786c1e78-ce78-4a10-afd6-843fcbb90c2e})"

type MediaPlaybackCommandManagerCommandBehavior struct {
	ole.IUnknown
}

func (impl *MediaPlaybackCommandManagerCommandBehavior) GetCommandManager() (*MediaPlaybackCommandManager, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerCommandBehavior))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerCommandBehavior)(unsafe.Pointer(itf))
	return v.GetCommandManager()
}

func (impl *MediaPlaybackCommandManagerCommandBehavior) GetIsEnabled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerCommandBehavior))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerCommandBehavior)(unsafe.Pointer(itf))
	return v.GetIsEnabled()
}

func (impl *MediaPlaybackCommandManagerCommandBehavior) GetEnablingRule() (MediaCommandEnablingRule, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerCommandBehavior))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerCommandBehavior)(unsafe.Pointer(itf))
	return v.GetEnablingRule()
}

func (impl *MediaPlaybackCommandManagerCommandBehavior) SetEnablingRule(value MediaCommandEnablingRule) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerCommandBehavior))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerCommandBehavior)(unsafe.Pointer(itf))
	return v.SetEnablingRule(value)
}

func (impl *MediaPlaybackCommandManagerCommandBehavior) AddIsEnabledChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerCommandBehavior))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerCommandBehavior)(unsafe.Pointer(itf))
	return v.AddIsEnabledChanged(handler)
}

func (impl *MediaPlaybackCommandManagerCommandBehavior) RemoveIsEnabledChanged(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerCommandBehavior))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerCommandBehavior)(unsafe.Pointer(itf))
	return v.RemoveIsEnabledChanged(token)
}

const GUIDiMediaPlaybackCommandManagerCommandBehavior string = "786c1e78-ce78-4a10-afd6-843fcbb90c2e"
const SignatureiMediaPlaybackCommandManagerCommandBehavior string = "{786c1e78-ce78-4a10-afd6-843fcbb90c2e}"

type iMediaPlaybackCommandManagerCommandBehavior struct {
	ole.IInspectable
}

type iMediaPlaybackCommandManagerCommandBehaviorVtbl struct {
	ole.IInspectableVtbl

	GetCommandManager      uintptr
	GetIsEnabled           uintptr
	GetEnablingRule        uintptr
	SetEnablingRule        uintptr
	AddIsEnabledChanged    uintptr
	RemoveIsEnabledChanged uintptr
}

func (v *iMediaPlaybackCommandManagerCommandBehavior) VTable() *iMediaPlaybackCommandManagerCommandBehaviorVtbl {
	return (*iMediaPlaybackCommandManagerCommandBehaviorVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMediaPlaybackCommandManagerCommandBehavior) GetCommandManager() (*MediaPlaybackCommandManager, error) {
	var out *MediaPlaybackCommandManager
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetCommandManager,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MediaPlaybackCommandManager
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackCommandManagerCommandBehavior) GetIsEnabled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsEnabled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackCommandManagerCommandBehavior) GetEnablingRule() (MediaCommandEnablingRule, error) {
	var out MediaCommandEnablingRule
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetEnablingRule,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MediaCommandEnablingRule
	)

	if hr != 0 {
		return MediaCommandEnablingRuleAuto, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackCommandManagerCommandBehavior) SetEnablingRule(value MediaCommandEnablingRule) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetEnablingRule,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in MediaCommandEnablingRule
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMediaPlaybackCommandManagerCommandBehavior) AddIsEnabledChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddIsEnabledChanged,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackCommandManagerCommandBehavior) RemoveIsEnabledChanged(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveIsEnabledChanged,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
