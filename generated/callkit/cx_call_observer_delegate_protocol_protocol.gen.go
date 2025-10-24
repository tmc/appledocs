// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCXCallObserverDelegate is the CXCallObserverDelegate protocol interface.
//
// A collection of methods the system calls when a call changes state.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - visionOS 1.0+
//   - watchOS 9.0+
//
// See: doc://com.apple.callkit/documentation/CallKit/CXCallObserverDelegate
type PCXCallObserverDelegate interface {
	// Required methods
	CallObserverCallChanged(callObserver ICXCallObserver, call ICXCall)/* debug [protocol_interface/required_method]: CallObserverCallChanged */
}

// CXCallObserverDelegate is a delegate implementation builder for the PCXCallObserverDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CXCallObserverDelegate struct {
	_CallObserverCallChanged func(callObserver ICXCallObserver, call ICXCall)
}

// SetCallObserverCallChanged sets the handler for the CallObserverCallChanged delegate method.
//
// Called when a call is changed.
func (d *CXCallObserverDelegate) SetCallObserverCallChanged(f func(callObserver ICXCallObserver, call ICXCall)) {
	d._CallObserverCallChanged = f
}

// CallObserverCallChanged implements the PCXCallObserverDelegate interface.
func (d *CXCallObserverDelegate) CallObserverCallChanged(callObserver ICXCallObserver, call ICXCall) {
	if d._CallObserverCallChanged != nil {
		d._CallObserverCallChanged(callObserver, call)
	}
}

// HasCallObserverCallChanged returns true if a handler for CallObserverCallChanged has been set.
func (d *CXCallObserverDelegate) HasCallObserverCallChanged() bool {
	return d._CallObserverCallChanged != nil
}
