// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PCustomRoutingControllerDelegate is the AVCustomRoutingControllerDelegate protocol interface.
//
// A protocol for delegates of a custom routing controller.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.avrouting/documentation/AVRouting/AVCustomRoutingControllerDelegate
type PCustomRoutingControllerDelegate interface {
	// Required methods
	CustomRoutingControllerHandleEventCompletionHandler(controller IAVCustomRoutingController, event IAVCustomRoutingEvent, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: CustomRoutingControllerHandleEventCompletionHandler */
	// Optional methods
	CustomRoutingControllerDidSelectItem(controller IAVCustomRoutingController, customActionItem IAVCustomRoutingActionItem)
	HasCustomRoutingControllerDidSelectItem() bool
	CustomRoutingControllerEventDidTimeOut(controller IAVCustomRoutingController, event IAVCustomRoutingEvent)
	HasCustomRoutingControllerEventDidTimeOut() bool
}

// CustomRoutingControllerDelegate is a delegate implementation builder for the PCustomRoutingControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CustomRoutingControllerDelegate struct {
	_CustomRoutingControllerDidSelectItem func(controller IAVCustomRoutingController, customActionItem IAVCustomRoutingActionItem)
	_CustomRoutingControllerEventDidTimeOut func(controller IAVCustomRoutingController, event IAVCustomRoutingEvent)
	_CustomRoutingControllerHandleEventCompletionHandler func(controller IAVCustomRoutingController, event IAVCustomRoutingEvent, completionHandler unsafe.Pointer)
}

// SetCustomRoutingControllerDidSelectItem sets the handler for the CustomRoutingControllerDidSelectItem delegate method.
//
// Tells the delegate when a user selects a custom item in the route picker.
func (d *CustomRoutingControllerDelegate) SetCustomRoutingControllerDidSelectItem(f func(controller IAVCustomRoutingController, customActionItem IAVCustomRoutingActionItem)) {
	d._CustomRoutingControllerDidSelectItem = f
}

// SetCustomRoutingControllerEventDidTimeOut sets the handler for the CustomRoutingControllerEventDidTimeOut delegate method.
//
// Tells the delegate when a routing event times out.
func (d *CustomRoutingControllerDelegate) SetCustomRoutingControllerEventDidTimeOut(f func(controller IAVCustomRoutingController, event IAVCustomRoutingEvent)) {
	d._CustomRoutingControllerEventDidTimeOut = f
}

// SetCustomRoutingControllerHandleEventCompletionHandler sets the handler for the CustomRoutingControllerHandleEventCompletionHandler delegate method.
//
// Connects to, or disconnects from, a device when a user requests it in the picker.
func (d *CustomRoutingControllerDelegate) SetCustomRoutingControllerHandleEventCompletionHandler(f func(controller IAVCustomRoutingController, event IAVCustomRoutingEvent, completionHandler unsafe.Pointer)) {
	d._CustomRoutingControllerHandleEventCompletionHandler = f
}

// CustomRoutingControllerDidSelectItem implements the PCustomRoutingControllerDelegate interface.
func (d *CustomRoutingControllerDelegate) CustomRoutingControllerDidSelectItem(controller IAVCustomRoutingController, customActionItem IAVCustomRoutingActionItem) {
	if d._CustomRoutingControllerDidSelectItem != nil {
		d._CustomRoutingControllerDidSelectItem(controller, customActionItem)
	}
}

// HasCustomRoutingControllerDidSelectItem returns true if a handler for CustomRoutingControllerDidSelectItem has been set.
func (d *CustomRoutingControllerDelegate) HasCustomRoutingControllerDidSelectItem() bool {
	return d._CustomRoutingControllerDidSelectItem != nil
}

// CustomRoutingControllerEventDidTimeOut implements the PCustomRoutingControllerDelegate interface.
func (d *CustomRoutingControllerDelegate) CustomRoutingControllerEventDidTimeOut(controller IAVCustomRoutingController, event IAVCustomRoutingEvent) {
	if d._CustomRoutingControllerEventDidTimeOut != nil {
		d._CustomRoutingControllerEventDidTimeOut(controller, event)
	}
}

// HasCustomRoutingControllerEventDidTimeOut returns true if a handler for CustomRoutingControllerEventDidTimeOut has been set.
func (d *CustomRoutingControllerDelegate) HasCustomRoutingControllerEventDidTimeOut() bool {
	return d._CustomRoutingControllerEventDidTimeOut != nil
}

// CustomRoutingControllerHandleEventCompletionHandler implements the PCustomRoutingControllerDelegate interface.
func (d *CustomRoutingControllerDelegate) CustomRoutingControllerHandleEventCompletionHandler(controller IAVCustomRoutingController, event IAVCustomRoutingEvent, completionHandler unsafe.Pointer) {
	if d._CustomRoutingControllerHandleEventCompletionHandler != nil {
		d._CustomRoutingControllerHandleEventCompletionHandler(controller, event, completionHandler)
	}
}

// HasCustomRoutingControllerHandleEventCompletionHandler returns true if a handler for CustomRoutingControllerHandleEventCompletionHandler has been set.
func (d *CustomRoutingControllerDelegate) HasCustomRoutingControllerHandleEventCompletionHandler() bool {
	return d._CustomRoutingControllerHandleEventCompletionHandler != nil
}
