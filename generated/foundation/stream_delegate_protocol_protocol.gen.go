// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PStreamDelegate is the NSStreamDelegate protocol interface.
//
// An interface that delegates of a stream instance use to handle events on the stream.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/StreamDelegate
type PStreamDelegate interface {
}

// StreamDelegate is a delegate implementation builder for the PStreamDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type StreamDelegate struct {
}

// StreamDelegateObject wraps an existing Objective-C object that conforms to the PStreamDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type StreamDelegateObject struct {
	objectivec.Object
}

// NewStreamDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSStreamDelegate protocol.
func NewStreamDelegateObject(obj objectivec.Object) *StreamDelegateObject {
	return &StreamDelegateObject{obj}
}

// Make sure StreamDelegateObject implements PStreamDelegate.
var _ PStreamDelegate = (*StreamDelegateObject)(nil)
