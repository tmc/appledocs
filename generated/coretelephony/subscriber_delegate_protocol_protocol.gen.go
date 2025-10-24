// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PSubscriberDelegate is the CTSubscriberDelegate protocol interface.
//
// A protocol to handle changes to subscriber information.
//
// Availability:
//   - iOS 12.1+
//   - iPadOS 12.1+
//
// See: doc://com.apple.coretelephony/documentation/CoreTelephony/CTSubscriberDelegate
type PSubscriberDelegate interface {
	// Required methods
	SubscriberTokenRefreshed(subscriber ICTSubscriber)/* debug [protocol_interface/required_method]: SubscriberTokenRefreshed */
}

// SubscriberDelegate is a delegate implementation builder for the PSubscriberDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SubscriberDelegate struct {
	_SubscriberTokenRefreshed func(subscriber ICTSubscriber)
}

// SetSubscriberTokenRefreshed sets the handler for the SubscriberTokenRefreshed delegate method.
//
// Tells the delegate the subscriber’s token refreshed.
func (d *SubscriberDelegate) SetSubscriberTokenRefreshed(f func(subscriber ICTSubscriber)) {
	d._SubscriberTokenRefreshed = f
}

// SubscriberTokenRefreshed implements the PSubscriberDelegate interface.
func (d *SubscriberDelegate) SubscriberTokenRefreshed(subscriber ICTSubscriber) {
	if d._SubscriberTokenRefreshed != nil {
		d._SubscriberTokenRefreshed(subscriber)
	}
}

// HasSubscriberTokenRefreshed returns true if a handler for SubscriberTokenRefreshed has been set.
func (d *SubscriberDelegate) HasSubscriberTokenRefreshed() bool {
	return d._SubscriberTokenRefreshed != nil
}
