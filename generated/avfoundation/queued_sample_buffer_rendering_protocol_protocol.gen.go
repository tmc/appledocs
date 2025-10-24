// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PQueuedSampleBufferRendering is the AVQueuedSampleBufferRendering protocol interface.
//
// Methods you can implement to enqueue sample buffers for presentation.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVQueuedSampleBufferRendering
type PQueuedSampleBufferRendering interface {
	// Required methods
	EnqueueSampleBuffer(sampleBuffer SampleBufferRef /* not a class type */)/* debug [protocol_interface/required_method]: EnqueueSampleBuffer */
	Flush()/* debug [protocol_interface/required_method]: Flush */
	RequestMediaDataWhenReadyOnQueueUsingBlock(queue objectivec.IObject, block unsafe.Pointer)/* debug [protocol_interface/required_method]: RequestMediaDataWhenReadyOnQueueUsingBlock */
	StopRequestingMediaData()/* debug [protocol_interface/required_method]: StopRequestingMediaData */
}
