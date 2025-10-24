// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PInputPickerInteractionDelegate is the AVInputPickerInteractionDelegate protocol interface.
//
// The   protocol defines methods you use to receive notifications about transitions in an   object.
//
// Availability:
//   - iOS 26.0+
//   - iPadOS 26.0+
//
// See: doc://com.apple.avkit/documentation/AVKit/AVInputPickerInteraction/Delegate-swift.protocol
type PInputPickerInteractionDelegate interface {
	// Optional methods
	InputPickerInteractionDidEndDismissing(inputPickerInteraction IAVInputPickerInteraction)
	HasInputPickerInteractionDidEndDismissing() bool
	InputPickerInteractionDidEndPresenting(inputPickerInteraction IAVInputPickerInteraction)
	HasInputPickerInteractionDidEndPresenting() bool
	InputPickerInteractionWillBeginDismissing(inputPickerInteraction IAVInputPickerInteraction)
	HasInputPickerInteractionWillBeginDismissing() bool
	InputPickerInteractionWillBeginPresenting(inputPickerInteraction IAVInputPickerInteraction)
	HasInputPickerInteractionWillBeginPresenting() bool
}

// InputPickerInteractionDelegate is a delegate implementation builder for the PInputPickerInteractionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type InputPickerInteractionDelegate struct {
	_InputPickerInteractionDidEndDismissing func(inputPickerInteraction IAVInputPickerInteraction)
	_InputPickerInteractionDidEndPresenting func(inputPickerInteraction IAVInputPickerInteraction)
	_InputPickerInteractionWillBeginDismissing func(inputPickerInteraction IAVInputPickerInteraction)
	_InputPickerInteractionWillBeginPresenting func(inputPickerInteraction IAVInputPickerInteraction)
}

// SetInputPickerInteractionDidEndDismissing sets the handler for the InputPickerInteractionDidEndDismissing delegate method.
//
// Tells the delegate that the input picker view has finished dismissing devices.
func (d *InputPickerInteractionDelegate) SetInputPickerInteractionDidEndDismissing(f func(inputPickerInteraction IAVInputPickerInteraction)) {
	d._InputPickerInteractionDidEndDismissing = f
}

// SetInputPickerInteractionDidEndPresenting sets the handler for the InputPickerInteractionDidEndPresenting delegate method.
//
// Tells the delegate that the input picker view has finished presenting devices
func (d *InputPickerInteractionDelegate) SetInputPickerInteractionDidEndPresenting(f func(inputPickerInteraction IAVInputPickerInteraction)) {
	d._InputPickerInteractionDidEndPresenting = f
}

// SetInputPickerInteractionWillBeginDismissing sets the handler for the InputPickerInteractionWillBeginDismissing delegate method.
//
// Tells the delegate that the input picker view is about to dismiss devices.
func (d *InputPickerInteractionDelegate) SetInputPickerInteractionWillBeginDismissing(f func(inputPickerInteraction IAVInputPickerInteraction)) {
	d._InputPickerInteractionWillBeginDismissing = f
}

// SetInputPickerInteractionWillBeginPresenting sets the handler for the InputPickerInteractionWillBeginPresenting delegate method.
//
// Tells the delegate that the input picker view is about to present devices.
func (d *InputPickerInteractionDelegate) SetInputPickerInteractionWillBeginPresenting(f func(inputPickerInteraction IAVInputPickerInteraction)) {
	d._InputPickerInteractionWillBeginPresenting = f
}

// InputPickerInteractionDidEndDismissing implements the PInputPickerInteractionDelegate interface.
func (d *InputPickerInteractionDelegate) InputPickerInteractionDidEndDismissing(inputPickerInteraction IAVInputPickerInteraction) {
	if d._InputPickerInteractionDidEndDismissing != nil {
		d._InputPickerInteractionDidEndDismissing(inputPickerInteraction)
	}
}

// HasInputPickerInteractionDidEndDismissing returns true if a handler for InputPickerInteractionDidEndDismissing has been set.
func (d *InputPickerInteractionDelegate) HasInputPickerInteractionDidEndDismissing() bool {
	return d._InputPickerInteractionDidEndDismissing != nil
}

// InputPickerInteractionDidEndPresenting implements the PInputPickerInteractionDelegate interface.
func (d *InputPickerInteractionDelegate) InputPickerInteractionDidEndPresenting(inputPickerInteraction IAVInputPickerInteraction) {
	if d._InputPickerInteractionDidEndPresenting != nil {
		d._InputPickerInteractionDidEndPresenting(inputPickerInteraction)
	}
}

// HasInputPickerInteractionDidEndPresenting returns true if a handler for InputPickerInteractionDidEndPresenting has been set.
func (d *InputPickerInteractionDelegate) HasInputPickerInteractionDidEndPresenting() bool {
	return d._InputPickerInteractionDidEndPresenting != nil
}

// InputPickerInteractionWillBeginDismissing implements the PInputPickerInteractionDelegate interface.
func (d *InputPickerInteractionDelegate) InputPickerInteractionWillBeginDismissing(inputPickerInteraction IAVInputPickerInteraction) {
	if d._InputPickerInteractionWillBeginDismissing != nil {
		d._InputPickerInteractionWillBeginDismissing(inputPickerInteraction)
	}
}

// HasInputPickerInteractionWillBeginDismissing returns true if a handler for InputPickerInteractionWillBeginDismissing has been set.
func (d *InputPickerInteractionDelegate) HasInputPickerInteractionWillBeginDismissing() bool {
	return d._InputPickerInteractionWillBeginDismissing != nil
}

// InputPickerInteractionWillBeginPresenting implements the PInputPickerInteractionDelegate interface.
func (d *InputPickerInteractionDelegate) InputPickerInteractionWillBeginPresenting(inputPickerInteraction IAVInputPickerInteraction) {
	if d._InputPickerInteractionWillBeginPresenting != nil {
		d._InputPickerInteractionWillBeginPresenting(inputPickerInteraction)
	}
}

// HasInputPickerInteractionWillBeginPresenting returns true if a handler for InputPickerInteractionWillBeginPresenting has been set.
func (d *InputPickerInteractionDelegate) HasInputPickerInteractionWillBeginPresenting() bool {
	return d._InputPickerInteractionWillBeginPresenting != nil
}
