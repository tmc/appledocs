//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureSynchronizedData


// iOS-only properties

// The time at which this synchronized data was captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedData/timestamp
func (c_ CaptureSynchronizedData) Timestamp() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("timestamp"))
	return rv
}





