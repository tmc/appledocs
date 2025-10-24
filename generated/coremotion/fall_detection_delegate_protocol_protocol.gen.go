// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PFallDetectionDelegate is the CMFallDetectionDelegate protocol interface.
//
// A delegate that receives information about fall detection events and authorization status changes.
//
// Availability:
//   - watchOS 7.2+
//
// See: doc://com.apple.coremotion/documentation/CoreMotion/CMFallDetectionDelegate
type PFallDetectionDelegate interface {
	// Optional methods
	FallDetectionManagerDidDetectEventCompletionHandler(fallDetectionManager ICMFallDetectionManager, event ICMFallDetectionEvent, handler unsafe.Pointer)
	HasFallDetectionManagerDidDetectEventCompletionHandler() bool
	FallDetectionManagerDidChangeAuthorization(fallDetectionManager ICMFallDetectionManager)
	HasFallDetectionManagerDidChangeAuthorization() bool
}

// FallDetectionDelegate is a delegate implementation builder for the PFallDetectionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type FallDetectionDelegate struct {
	_FallDetectionManagerDidDetectEventCompletionHandler func(fallDetectionManager ICMFallDetectionManager, event ICMFallDetectionEvent, handler unsafe.Pointer)
	_FallDetectionManagerDidChangeAuthorization func(fallDetectionManager ICMFallDetectionManager)
}

// SetFallDetectionManagerDidDetectEventCompletionHandler sets the handler for the FallDetectionManagerDidDetectEventCompletionHandler delegate method.
//
// Indicates a fall detection event occurred.
func (d *FallDetectionDelegate) SetFallDetectionManagerDidDetectEventCompletionHandler(f func(fallDetectionManager ICMFallDetectionManager, event ICMFallDetectionEvent, handler unsafe.Pointer)) {
	d._FallDetectionManagerDidDetectEventCompletionHandler = f
}

// SetFallDetectionManagerDidChangeAuthorization sets the handler for the FallDetectionManagerDidChangeAuthorization delegate method.
//
// Indicates the fall detection authorization status changed.
func (d *FallDetectionDelegate) SetFallDetectionManagerDidChangeAuthorization(f func(fallDetectionManager ICMFallDetectionManager)) {
	d._FallDetectionManagerDidChangeAuthorization = f
}

// FallDetectionManagerDidDetectEventCompletionHandler implements the PFallDetectionDelegate interface.
func (d *FallDetectionDelegate) FallDetectionManagerDidDetectEventCompletionHandler(fallDetectionManager ICMFallDetectionManager, event ICMFallDetectionEvent, handler unsafe.Pointer) {
	if d._FallDetectionManagerDidDetectEventCompletionHandler != nil {
		d._FallDetectionManagerDidDetectEventCompletionHandler(fallDetectionManager, event, handler)
	}
}

// HasFallDetectionManagerDidDetectEventCompletionHandler returns true if a handler for FallDetectionManagerDidDetectEventCompletionHandler has been set.
func (d *FallDetectionDelegate) HasFallDetectionManagerDidDetectEventCompletionHandler() bool {
	return d._FallDetectionManagerDidDetectEventCompletionHandler != nil
}

// FallDetectionManagerDidChangeAuthorization implements the PFallDetectionDelegate interface.
func (d *FallDetectionDelegate) FallDetectionManagerDidChangeAuthorization(fallDetectionManager ICMFallDetectionManager) {
	if d._FallDetectionManagerDidChangeAuthorization != nil {
		d._FallDetectionManagerDidChangeAuthorization(fallDetectionManager)
	}
}

// HasFallDetectionManagerDidChangeAuthorization returns true if a handler for FallDetectionManagerDidChangeAuthorization has been set.
func (d *FallDetectionDelegate) HasFallDetectionManagerDidChangeAuthorization() bool {
	return d._FallDetectionManagerDidChangeAuthorization != nil
}
