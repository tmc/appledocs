// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PAsynchronousKeyValueLoading is the AVAsynchronousKeyValueLoading protocol interface.
//
// A protocol that defines the interface to load media data asynchronously.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVAsynchronousKeyValueLoading
type PAsynchronousKeyValueLoading interface {
	// Required methods
	LoadValuesAsynchronouslyForKeysCompletionHandler(keys []string, handler unsafe.Pointer)
	StatusOfValueForKeyError(key foundation.foundation.INSString, outError foundation.foundation.INSError) KeyValueStatus
}
