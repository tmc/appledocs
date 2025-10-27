// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
)

// PMetricEventStreamSubscriber is the AVMetricEventStreamSubscriber protocol interface.
//
// A type for objects that receive metric events.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVMetricEventStreamSubscriber
type PMetricEventStreamSubscriber interface {
	// Required methods
	PublisherDidReceiveEvent(publisher unsafe.Pointer, event IAVMetricEvent)
}
