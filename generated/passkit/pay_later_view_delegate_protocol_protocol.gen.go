// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PPayLaterViewDelegate is the PKPayLaterViewDelegate protocol interface.
//
// Methods the framework calls when the Apple Pay Later view’s size changes.
//
// Availability:
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.passkit/documentation/PassKit/PKPayLaterViewDelegate
type PPayLaterViewDelegate interface {
	// Required methods
	PayLaterViewDidUpdateHeight(view PayLaterView /* not a class type */)/* debug [protocol_interface/required_method]: PayLaterViewDidUpdateHeight */
}

// PayLaterViewDelegate is a delegate implementation builder for the PPayLaterViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PayLaterViewDelegate struct {
	_PayLaterViewDidUpdateHeight func(view PayLaterView /* not a class type */)
}

// SetPayLaterViewDidUpdateHeight sets the handler for the PayLaterViewDidUpdateHeight delegate method.
//
// Tells the delegate when the Apple Pay Later visual merchandising widget’s height changes.
func (d *PayLaterViewDelegate) SetPayLaterViewDidUpdateHeight(f func(view PayLaterView /* not a class type */)) {
	d._PayLaterViewDidUpdateHeight = f
}

// PayLaterViewDidUpdateHeight implements the PPayLaterViewDelegate interface.
func (d *PayLaterViewDelegate) PayLaterViewDidUpdateHeight(view PayLaterView /* not a class type */) {
	if d._PayLaterViewDidUpdateHeight != nil {
		d._PayLaterViewDidUpdateHeight(view)
	}
}

// HasPayLaterViewDidUpdateHeight returns true if a handler for PayLaterViewDidUpdateHeight has been set.
func (d *PayLaterViewDelegate) HasPayLaterViewDidUpdateHeight() bool {
	return d._PayLaterViewDidUpdateHeight != nil
}
