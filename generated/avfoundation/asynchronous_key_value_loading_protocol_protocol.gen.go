// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	LoadValuesAsynchronouslyForKeysCompletionHandler(keys []string, handler unsafe.Pointer)/* debug [protocol_interface/required_method]: LoadValuesAsynchronouslyForKeysCompletionHandler */
	StatusOfValueForKeyError(key objc.IObject /* cross-framework: NSString */, outError objectivec.IObject) KeyValueStatus/* debug [protocol_interface/required_method]: StatusOfValueForKeyError */
}
