// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package control

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/foundation/collections"
	"github.com/saltosystems/winrt-go/windows/storage/streams"
)

const SignatureGlobalSystemMediaTransportControlsSessionMediaProperties string = "rc(Windows.Media.Control.GlobalSystemMediaTransportControlsSessionMediaProperties;{68856cf6-adb4-54b2-ac16-05837907acb6})"

type GlobalSystemMediaTransportControlsSessionMediaProperties struct {
	ole.IUnknown
}

func (impl *GlobalSystemMediaTransportControlsSessionMediaProperties) GetTitle() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionMediaProperties))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionMediaProperties)(unsafe.Pointer(itf))
	return v.GetTitle()
}

func (impl *GlobalSystemMediaTransportControlsSessionMediaProperties) GetSubtitle() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionMediaProperties))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionMediaProperties)(unsafe.Pointer(itf))
	return v.GetSubtitle()
}

func (impl *GlobalSystemMediaTransportControlsSessionMediaProperties) GetAlbumArtist() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionMediaProperties))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionMediaProperties)(unsafe.Pointer(itf))
	return v.GetAlbumArtist()
}

func (impl *GlobalSystemMediaTransportControlsSessionMediaProperties) GetArtist() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionMediaProperties))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionMediaProperties)(unsafe.Pointer(itf))
	return v.GetArtist()
}

func (impl *GlobalSystemMediaTransportControlsSessionMediaProperties) GetAlbumTitle() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionMediaProperties))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionMediaProperties)(unsafe.Pointer(itf))
	return v.GetAlbumTitle()
}

func (impl *GlobalSystemMediaTransportControlsSessionMediaProperties) GetTrackNumber() (int32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionMediaProperties))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionMediaProperties)(unsafe.Pointer(itf))
	return v.GetTrackNumber()
}

func (impl *GlobalSystemMediaTransportControlsSessionMediaProperties) GetGenres() (*collections.IVectorView, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionMediaProperties))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionMediaProperties)(unsafe.Pointer(itf))
	return v.GetGenres()
}

func (impl *GlobalSystemMediaTransportControlsSessionMediaProperties) GetAlbumTrackCount() (int32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionMediaProperties))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionMediaProperties)(unsafe.Pointer(itf))
	return v.GetAlbumTrackCount()
}

func (impl *GlobalSystemMediaTransportControlsSessionMediaProperties) GetPlaybackType() (*foundation.IReference, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionMediaProperties))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionMediaProperties)(unsafe.Pointer(itf))
	return v.GetPlaybackType()
}

func (impl *GlobalSystemMediaTransportControlsSessionMediaProperties) GetThumbnail() (*streams.IRandomAccessStreamReference, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGlobalSystemMediaTransportControlsSessionMediaProperties))
	defer itf.Release()
	v := (*iGlobalSystemMediaTransportControlsSessionMediaProperties)(unsafe.Pointer(itf))
	return v.GetThumbnail()
}

const GUIDiGlobalSystemMediaTransportControlsSessionMediaProperties string = "68856cf6-adb4-54b2-ac16-05837907acb6"
const SignatureiGlobalSystemMediaTransportControlsSessionMediaProperties string = "{68856cf6-adb4-54b2-ac16-05837907acb6}"

type iGlobalSystemMediaTransportControlsSessionMediaProperties struct {
	ole.IInspectable
}

type iGlobalSystemMediaTransportControlsSessionMediaPropertiesVtbl struct {
	ole.IInspectableVtbl

	GetTitle           uintptr
	GetSubtitle        uintptr
	GetAlbumArtist     uintptr
	GetArtist          uintptr
	GetAlbumTitle      uintptr
	GetTrackNumber     uintptr
	GetGenres          uintptr
	GetAlbumTrackCount uintptr
	GetPlaybackType    uintptr
	GetThumbnail       uintptr
}

func (v *iGlobalSystemMediaTransportControlsSessionMediaProperties) VTable() *iGlobalSystemMediaTransportControlsSessionMediaPropertiesVtbl {
	return (*iGlobalSystemMediaTransportControlsSessionMediaPropertiesVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iGlobalSystemMediaTransportControlsSessionMediaProperties) GetTitle() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetTitle,
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

func (v *iGlobalSystemMediaTransportControlsSessionMediaProperties) GetSubtitle() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSubtitle,
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

func (v *iGlobalSystemMediaTransportControlsSessionMediaProperties) GetAlbumArtist() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetAlbumArtist,
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

func (v *iGlobalSystemMediaTransportControlsSessionMediaProperties) GetArtist() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetArtist,
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

func (v *iGlobalSystemMediaTransportControlsSessionMediaProperties) GetAlbumTitle() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetAlbumTitle,
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

func (v *iGlobalSystemMediaTransportControlsSessionMediaProperties) GetTrackNumber() (int32, error) {
	var out int32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetTrackNumber,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out int32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionMediaProperties) GetGenres() (*collections.IVectorView, error) {
	var out *collections.IVectorView
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetGenres,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IVectorView
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionMediaProperties) GetAlbumTrackCount() (int32, error) {
	var out int32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetAlbumTrackCount,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out int32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionMediaProperties) GetPlaybackType() (*foundation.IReference, error) {
	var out *foundation.IReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetPlaybackType,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGlobalSystemMediaTransportControlsSessionMediaProperties) GetThumbnail() (*streams.IRandomAccessStreamReference, error) {
	var out *streams.IRandomAccessStreamReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetThumbnail,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out streams.IRandomAccessStreamReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
