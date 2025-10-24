//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for CaptureSynchronizedMetadataObjectData


// iOS-only properties

// The list of metadata objects captured at this synchronization timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedMetadataObjectData/metadataObjects
func (c_ CaptureSynchronizedMetadataObjectData) MetadataObjects() []MetadataObject {
	rv := objc.Send[[]MetadataObject](c_.ID, objc.Sel("metadataObjects"))
	return rv
}





