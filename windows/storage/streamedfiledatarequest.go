// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package storage

import (
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/storage/streams"
)

const SignatureStreamedFileDataRequest string = "rc(Windows.Storage.StreamedFileDataRequest;{905a0fe6-bc53-11df-8c49-001e4fc686da})"

type StreamedFileDataRequest struct {
	ole.IUnknown
}

func (impl *StreamedFileDataRequest) WriteAsync(buffer *IBuffer) (*foundation.IAsyncOperationWithProgress, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(streams.GUIDIOutputStream))
	defer itf.Release()
	v := (*streams.IOutputStream)(unsafe.Pointer(itf))
	return v.WriteAsync(buffer)
}

func (impl *StreamedFileDataRequest) FlushAsync() (*foundation.IAsyncOperation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(streams.GUIDIOutputStream))
	defer itf.Release()
	v := (*streams.IOutputStream)(unsafe.Pointer(itf))
	return v.FlushAsync()
}

func (impl *StreamedFileDataRequest) Close() error {
	itf := impl.MustQueryInterface(ole.NewGUID(foundation.GUIDIClosable))
	defer itf.Release()
	v := (*foundation.IClosable)(unsafe.Pointer(itf))
	return v.Close()
}

func (impl *StreamedFileDataRequest) FailAndClose(failureMode StreamedFileFailureMode) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIStreamedFileDataRequest))
	defer itf.Release()
	v := (*IStreamedFileDataRequest)(unsafe.Pointer(itf))
	return v.FailAndClose(failureMode)
}
