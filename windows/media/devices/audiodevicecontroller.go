// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package devices

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureAudioDeviceController string = "rc(Windows.Media.Devices.AudioDeviceController;{edd4a388-79c7-4f7c-90e8-ef934b21580a})"

type AudioDeviceController struct {
	ole.IUnknown
}

func (impl *AudioDeviceController) SetMuted(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioDeviceController))
	defer itf.Release()
	v := (*iAudioDeviceController)(unsafe.Pointer(itf))
	return v.SetMuted(value)
}

func (impl *AudioDeviceController) GetMuted() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioDeviceController))
	defer itf.Release()
	v := (*iAudioDeviceController)(unsafe.Pointer(itf))
	return v.GetMuted()
}

func (impl *AudioDeviceController) SetVolumePercent(value float32) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioDeviceController))
	defer itf.Release()
	v := (*iAudioDeviceController)(unsafe.Pointer(itf))
	return v.SetVolumePercent(value)
}

func (impl *AudioDeviceController) GetVolumePercent() (float32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioDeviceController))
	defer itf.Release()
	v := (*iAudioDeviceController)(unsafe.Pointer(itf))
	return v.GetVolumePercent()
}

func (impl *AudioDeviceController) GetAvailableMediaStreamProperties(mediaStreamType capture.MediaStreamType) (*collections.IVectorView, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIMediaDeviceController))
	defer itf.Release()
	v := (*IMediaDeviceController)(unsafe.Pointer(itf))
	return v.GetAvailableMediaStreamProperties(mediaStreamType)
}

func (impl *AudioDeviceController) GetMediaStreamProperties(mediaStreamType capture.MediaStreamType) (*mediaproperties.IMediaEncodingProperties, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIMediaDeviceController))
	defer itf.Release()
	v := (*IMediaDeviceController)(unsafe.Pointer(itf))
	return v.GetMediaStreamProperties(mediaStreamType)
}

func (impl *AudioDeviceController) SetMediaStreamPropertiesAsync(mediaStreamType capture.MediaStreamType, mediaEncodingProperties *mediaproperties.IMediaEncodingProperties) (*foundation.IAsyncAction, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIMediaDeviceController))
	defer itf.Release()
	v := (*IMediaDeviceController)(unsafe.Pointer(itf))
	return v.SetMediaStreamPropertiesAsync(mediaStreamType, mediaEncodingProperties)
}

const GUIDiAudioDeviceController string = "edd4a388-79c7-4f7c-90e8-ef934b21580a"
const SignatureiAudioDeviceController string = "{edd4a388-79c7-4f7c-90e8-ef934b21580a}"

type iAudioDeviceController struct {
	ole.IInspectable
}

type iAudioDeviceControllerVtbl struct {
	ole.IInspectableVtbl

	SetMuted         uintptr
	GetMuted         uintptr
	SetVolumePercent uintptr
	GetVolumePercent uintptr
}

func (v *iAudioDeviceController) VTable() *iAudioDeviceControllerVtbl {
	return (*iAudioDeviceControllerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iAudioDeviceController) SetMuted(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetMuted,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iAudioDeviceController) GetMuted() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetMuted,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iAudioDeviceController) SetVolumePercent(value float32) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetVolumePercent,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in float32
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iAudioDeviceController) GetVolumePercent() (float32, error) {
	var out float32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetVolumePercent,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out float32
	)

	if hr != 0 {
		return 0.0, ole.NewError(hr)
	}

	return out, nil
}
