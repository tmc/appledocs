// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corevideo"
)

// PMETrackReader is the METrackReader protocol interface.
//
// A protocol that defines the information to provide about a track within a media asset.
//
// Availability:
//   - macOS 14.0+
//
// See: doc://com.apple.mediaextension/documentation/MediaExtension/METrackReader
type PMETrackReader interface {
	// Required methods
	GenerateSampleCursorAtPresentationTimeStampCompletionHandler(presentationTimeStamp objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: GenerateSampleCursorAtPresentationTimeStampCompletionHandler */
	GenerateSampleCursorAtFirstSampleInDecodeOrderWithCompletionHandler(completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: GenerateSampleCursorAtFirstSampleInDecodeOrderWithCompletionHandler */
	GenerateSampleCursorAtLastSampleInDecodeOrderWithCompletionHandler(completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: GenerateSampleCursorAtLastSampleInDecodeOrderWithCompletionHandler */
	LoadTrackInfoWithCompletionHandler(completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: LoadTrackInfoWithCompletionHandler */
	// Optional methods
	LoadEstimatedDataRateWithCompletionHandler(completionHandler unsafe.Pointer)
	HasLoadEstimatedDataRateWithCompletionHandler() bool
	LoadMetadataWithCompletionHandler(completionHandler unsafe.Pointer)
	HasLoadMetadataWithCompletionHandler() bool
	LoadTotalSampleDataLengthWithCompletionHandler(completionHandler unsafe.Pointer)
	HasLoadTotalSampleDataLengthWithCompletionHandler() bool
	LoadUneditedDurationWithCompletionHandler(completionHandler unsafe.Pointer)
	HasLoadUneditedDurationWithCompletionHandler() bool
}
