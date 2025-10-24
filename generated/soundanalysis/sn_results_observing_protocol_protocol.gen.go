// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PSNResultsObserving is the SNResultsObserving protocol interface.
//
// The interface your app implements to receive the results of an analysis request.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+
//
// See: doc://com.apple.soundanalysis/documentation/SoundAnalysis/SNResultsObserving
type PSNResultsObserving interface {
	// Required methods
	RequestDidProduceResult(request unsafe.Pointer, result unsafe.Pointer)/* debug [protocol_interface/required_method]: RequestDidProduceResult */
	// Optional methods
	RequestDidFailWithError(request unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */)
	HasRequestDidFailWithError() bool
	RequestDidComplete(request unsafe.Pointer)
	HasRequestDidComplete() bool
}
