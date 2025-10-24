// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PRPPreviewViewControllerDelegate is the RPPreviewViewControllerDelegate protocol interface.
//
// The protocol you implement to respond to changes to a screen-recording user interface.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.replaykit/documentation/ReplayKit/RPPreviewViewControllerDelegate
type PRPPreviewViewControllerDelegate interface {
	// Optional methods
	PreviewControllerDidFinishWithActivityTypes(previewController IRPPreviewViewController, activityTypes unsafe.Pointer)
	HasPreviewControllerDidFinishWithActivityTypes() bool
	PreviewControllerDidFinish(previewController IRPPreviewViewController)
	HasPreviewControllerDidFinish() bool
}

// RPPreviewViewControllerDelegate is a delegate implementation builder for the PRPPreviewViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type RPPreviewViewControllerDelegate struct {
	_PreviewControllerDidFinishWithActivityTypes func(previewController IRPPreviewViewController, activityTypes unsafe.Pointer)
	_PreviewControllerDidFinish func(previewController IRPPreviewViewController)
}

// SetPreviewControllerDidFinishWithActivityTypes sets the handler for the PreviewControllerDidFinishWithActivityTypes delegate method.
//
// Indicates that the preview view controller is ready to be dismissed with associated activity types.
func (d *RPPreviewViewControllerDelegate) SetPreviewControllerDidFinishWithActivityTypes(f func(previewController IRPPreviewViewController, activityTypes unsafe.Pointer)) {
	d._PreviewControllerDidFinishWithActivityTypes = f
}

// SetPreviewControllerDidFinish sets the handler for the PreviewControllerDidFinish delegate method.
//
// Indicates that the preview view controller is ready to be dismissed.
func (d *RPPreviewViewControllerDelegate) SetPreviewControllerDidFinish(f func(previewController IRPPreviewViewController)) {
	d._PreviewControllerDidFinish = f
}

// PreviewControllerDidFinishWithActivityTypes implements the PRPPreviewViewControllerDelegate interface.
func (d *RPPreviewViewControllerDelegate) PreviewControllerDidFinishWithActivityTypes(previewController IRPPreviewViewController, activityTypes unsafe.Pointer) {
	if d._PreviewControllerDidFinishWithActivityTypes != nil {
		d._PreviewControllerDidFinishWithActivityTypes(previewController, activityTypes)
	}
}

// HasPreviewControllerDidFinishWithActivityTypes returns true if a handler for PreviewControllerDidFinishWithActivityTypes has been set.
func (d *RPPreviewViewControllerDelegate) HasPreviewControllerDidFinishWithActivityTypes() bool {
	return d._PreviewControllerDidFinishWithActivityTypes != nil
}

// PreviewControllerDidFinish implements the PRPPreviewViewControllerDelegate interface.
func (d *RPPreviewViewControllerDelegate) PreviewControllerDidFinish(previewController IRPPreviewViewController) {
	if d._PreviewControllerDidFinish != nil {
		d._PreviewControllerDidFinish(previewController)
	}
}

// HasPreviewControllerDidFinish returns true if a handler for PreviewControllerDidFinish has been set.
func (d *RPPreviewViewControllerDelegate) HasPreviewControllerDidFinish() bool {
	return d._PreviewControllerDidFinish != nil
}
