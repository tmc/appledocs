//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureMetadataInput


// Provides metadata to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataInput/append(_:)
func (c_ CaptureMetadataInput) AppendTimedMetadataGroupError(metadata IAVTimedMetadataGroup, outError objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("appendTimedMetadataGroup:error:"), metadata, outError)
	return rv
}

// iOS-only properties




