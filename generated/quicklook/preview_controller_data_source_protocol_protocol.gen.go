// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PPreviewControllerDataSource is the QLPreviewControllerDataSource protocol interface.
//
// The protocol that a data source for a preview controller needs to adopt to provide preview items to the controller.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.quicklook/documentation/QuickLook/QLPreviewControllerDataSource
type PPreviewControllerDataSource interface {
	// Required methods
	NumberOfPreviewItemsInPreviewController(controller IQLPreviewController) int/* debug [protocol_interface/required_method]: NumberOfPreviewItemsInPreviewController */
	PreviewControllerPreviewItemAtIndex(controller IQLPreviewController, index int) unsafe.Pointer/* debug [protocol_interface/required_method]: PreviewControllerPreviewItemAtIndex */
}

// PreviewControllerDataSource is a delegate implementation builder for the PPreviewControllerDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PreviewControllerDataSource struct {
	_NumberOfPreviewItemsInPreviewController func(controller IQLPreviewController) int
	_PreviewControllerPreviewItemAtIndex func(controller IQLPreviewController, index int) unsafe.Pointer
}

// SetNumberOfPreviewItemsInPreviewController sets the handler for the NumberOfPreviewItemsInPreviewController delegate method.
//
// Returns the number of preview items to include in the preview navigation list.
func (d *PreviewControllerDataSource) SetNumberOfPreviewItemsInPreviewController(f func(controller IQLPreviewController) int) {
	d._NumberOfPreviewItemsInPreviewController = f
}

// SetPreviewControllerPreviewItemAtIndex sets the handler for the PreviewControllerPreviewItemAtIndex delegate method.
//
// Returns the preview item that the controller displays for the specified index.
func (d *PreviewControllerDataSource) SetPreviewControllerPreviewItemAtIndex(f func(controller IQLPreviewController, index int) unsafe.Pointer) {
	d._PreviewControllerPreviewItemAtIndex = f
}

// NumberOfPreviewItemsInPreviewController implements the PPreviewControllerDataSource interface.
func (d *PreviewControllerDataSource) NumberOfPreviewItemsInPreviewController(controller IQLPreviewController) int {
	if d._NumberOfPreviewItemsInPreviewController != nil {
		return d._NumberOfPreviewItemsInPreviewController(controller)
	}
	var zero int
	return zero
}

// HasNumberOfPreviewItemsInPreviewController returns true if a handler for NumberOfPreviewItemsInPreviewController has been set.
func (d *PreviewControllerDataSource) HasNumberOfPreviewItemsInPreviewController() bool {
	return d._NumberOfPreviewItemsInPreviewController != nil
}

// PreviewControllerPreviewItemAtIndex implements the PPreviewControllerDataSource interface.
func (d *PreviewControllerDataSource) PreviewControllerPreviewItemAtIndex(controller IQLPreviewController, index int) unsafe.Pointer {
	if d._PreviewControllerPreviewItemAtIndex != nil {
		return d._PreviewControllerPreviewItemAtIndex(controller, index)
	}
	var zero unsafe.Pointer
	return zero
}

// HasPreviewControllerPreviewItemAtIndex returns true if a handler for PreviewControllerPreviewItemAtIndex has been set.
func (d *PreviewControllerDataSource) HasPreviewControllerPreviewItemAtIndex() bool {
	return d._PreviewControllerPreviewItemAtIndex != nil
}
