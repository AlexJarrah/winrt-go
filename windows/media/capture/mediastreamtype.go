// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package capture

type MediaStreamType int32

const SignatureMediaStreamType string = "enum(Windows.Media.Capture.MediaStreamType;i4)"

const (
	MediaStreamTypeVideoPreview MediaStreamType = 0
	MediaStreamTypeVideoRecord  MediaStreamType = 1
	MediaStreamTypeAudio        MediaStreamType = 2
	MediaStreamTypePhoto        MediaStreamType = 3
)
