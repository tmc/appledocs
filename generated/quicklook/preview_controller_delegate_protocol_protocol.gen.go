// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/appkit"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"
)

// PPreviewControllerDelegate is the QLPreviewControllerDelegate protocol interface.
//
// The protocol that a delegate of a preview controller needs to adopt to handle Quick Look previews.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.quicklook/documentation/QuickLook/QLPreviewControllerDelegate
type PPreviewControllerDelegate interface {
	// Optional methods
	PreviewControllerDidSaveEditedCopyOfPreviewItemAtURL(controller IQLPreviewController, previewItem unsafe.Pointer, modifiedContentsURL objc.IObject /* cross-framework: NSURL */)
	HasPreviewControllerDidSaveEditedCopyOfPreviewItemAtURL() bool
	PreviewControllerDidUpdateContentsOfPreviewItem(controller IQLPreviewController, previewItem unsafe.Pointer)
	HasPreviewControllerDidUpdateContentsOfPreviewItem() bool
	PreviewControllerEditingModeForPreviewItem(controller IQLPreviewController, previewItem unsafe.Pointer) PreviewItemEditingMode
	HasPreviewControllerEditingModeForPreviewItem() bool
	PreviewControllerFrameForPreviewItemInSourceView(controller IQLPreviewController, item unsafe.Pointer, view unsafe.Pointer) corefoundation.CGRect
	HasPreviewControllerFrameForPreviewItemInSourceView() bool
	PreviewControllerShouldOpenURLForPreviewItem(controller IQLPreviewController, url objc.IObject /* cross-framework: NSURL */, item unsafe.Pointer) bool
	HasPreviewControllerShouldOpenURLForPreviewItem() bool
	PreviewControllerTransitionImageForPreviewItemContentRect(controller IQLPreviewController, item unsafe.Pointer, contentRect corefoundation.CGRect) appkit.Image
	HasPreviewControllerTransitionImageForPreviewItemContentRect() bool
	PreviewControllerTransitionViewForPreviewItem(controller IQLPreviewController, item unsafe.Pointer) appkit.View
	HasPreviewControllerTransitionViewForPreviewItem() bool
	PreviewControllerDidDismiss(controller IQLPreviewController)
	HasPreviewControllerDidDismiss() bool
	PreviewControllerWillDismiss(controller IQLPreviewController)
	HasPreviewControllerWillDismiss() bool
}

// PreviewControllerDelegate is a delegate implementation builder for the PPreviewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PreviewControllerDelegate struct {
	_PreviewControllerDidSaveEditedCopyOfPreviewItemAtURL func(controller IQLPreviewController, previewItem unsafe.Pointer, modifiedContentsURL objc.IObject /* cross-framework: NSURL */)
	_PreviewControllerDidUpdateContentsOfPreviewItem func(controller IQLPreviewController, previewItem unsafe.Pointer)
	_PreviewControllerEditingModeForPreviewItem func(controller IQLPreviewController, previewItem unsafe.Pointer) PreviewItemEditingMode
	_PreviewControllerFrameForPreviewItemInSourceView func(controller IQLPreviewController, item unsafe.Pointer, view unsafe.Pointer) corefoundation.CGRect
	_PreviewControllerShouldOpenURLForPreviewItem func(controller IQLPreviewController, url objc.IObject /* cross-framework: NSURL */, item unsafe.Pointer) bool
	_PreviewControllerTransitionImageForPreviewItemContentRect func(controller IQLPreviewController, item unsafe.Pointer, contentRect corefoundation.CGRect) appkit.Image
	_PreviewControllerTransitionViewForPreviewItem func(controller IQLPreviewController, item unsafe.Pointer) appkit.View
	_PreviewControllerDidDismiss func(controller IQLPreviewController)
	_PreviewControllerWillDismiss func(controller IQLPreviewController)
}

// SetPreviewControllerDidSaveEditedCopyOfPreviewItemAtURL sets the handler for the PreviewControllerDidSaveEditedCopyOfPreviewItemAtURL delegate method.
//
// Tells the delegate that the preview item’s edited content was successfully saved to a copy at the given URL.
func (d *PreviewControllerDelegate) SetPreviewControllerDidSaveEditedCopyOfPreviewItemAtURL(f func(controller IQLPreviewController, previewItem unsafe.Pointer, modifiedContentsURL objc.IObject /* cross-framework: NSURL */)) {
	d._PreviewControllerDidSaveEditedCopyOfPreviewItemAtURL = f
}

// SetPreviewControllerDidUpdateContentsOfPreviewItem sets the handler for the PreviewControllerDidUpdateContentsOfPreviewItem delegate method.
//
// Tells the delegate that the content of a preview was updated successfully.
func (d *PreviewControllerDelegate) SetPreviewControllerDidUpdateContentsOfPreviewItem(f func(controller IQLPreviewController, previewItem unsafe.Pointer)) {
	d._PreviewControllerDidUpdateContentsOfPreviewItem = f
}

// SetPreviewControllerEditingModeForPreviewItem sets the handler for the PreviewControllerEditingModeForPreviewItem delegate method.
//
// Returns a value that indicates how the preview controller handles edits to the content of the previewed file.
func (d *PreviewControllerDelegate) SetPreviewControllerEditingModeForPreviewItem(f func(controller IQLPreviewController, previewItem unsafe.Pointer) PreviewItemEditingMode) {
	d._PreviewControllerEditingModeForPreviewItem = f
}

// SetPreviewControllerFrameForPreviewItemInSourceView sets the handler for the PreviewControllerFrameForPreviewItemInSourceView delegate method.
//
// Tells the delegate that the system is about to present the preview full screen or dismiss it, and asks for information to provide a zoom effect.
func (d *PreviewControllerDelegate) SetPreviewControllerFrameForPreviewItemInSourceView(f func(controller IQLPreviewController, item unsafe.Pointer, view unsafe.Pointer) corefoundation.CGRect) {
	d._PreviewControllerFrameForPreviewItemInSourceView = f
}

// SetPreviewControllerShouldOpenURLForPreviewItem sets the handler for the PreviewControllerShouldOpenURLForPreviewItem delegate method.
//
// Tells the delegate that the preview controller is trying to open a URL.
func (d *PreviewControllerDelegate) SetPreviewControllerShouldOpenURLForPreviewItem(f func(controller IQLPreviewController, url objc.IObject /* cross-framework: NSURL */, item unsafe.Pointer) bool) {
	d._PreviewControllerShouldOpenURLForPreviewItem = f
}

// SetPreviewControllerTransitionImageForPreviewItemContentRect sets the handler for the PreviewControllerTransitionImageForPreviewItemContentRect delegate method.
//
// Tells the delegate that the system is about to present the preview full screen or dismiss it, and asks for information to provide a smooth transition when zooming.
func (d *PreviewControllerDelegate) SetPreviewControllerTransitionImageForPreviewItemContentRect(f func(controller IQLPreviewController, item unsafe.Pointer, contentRect corefoundation.CGRect) appkit.Image) {
	d._PreviewControllerTransitionImageForPreviewItemContentRect = f
}

// SetPreviewControllerTransitionViewForPreviewItem sets the handler for the PreviewControllerTransitionViewForPreviewItem delegate method.
//
// Tells the delegate that the system is about to present the preview full screen or dismiss it, and asks for information to provide a smooth transition when zooming.
func (d *PreviewControllerDelegate) SetPreviewControllerTransitionViewForPreviewItem(f func(controller IQLPreviewController, item unsafe.Pointer) appkit.View) {
	d._PreviewControllerTransitionViewForPreviewItem = f
}

// SetPreviewControllerDidDismiss sets the handler for the PreviewControllerDidDismiss delegate method.
//
// Tells the delegate that the preview was closed.
func (d *PreviewControllerDelegate) SetPreviewControllerDidDismiss(f func(controller IQLPreviewController)) {
	d._PreviewControllerDidDismiss = f
}

// SetPreviewControllerWillDismiss sets the handler for the PreviewControllerWillDismiss delegate method.
//
// Tells the delegate that the preview is about to close.
func (d *PreviewControllerDelegate) SetPreviewControllerWillDismiss(f func(controller IQLPreviewController)) {
	d._PreviewControllerWillDismiss = f
}

// PreviewControllerDidSaveEditedCopyOfPreviewItemAtURL implements the PPreviewControllerDelegate interface.
func (d *PreviewControllerDelegate) PreviewControllerDidSaveEditedCopyOfPreviewItemAtURL(controller IQLPreviewController, previewItem unsafe.Pointer, modifiedContentsURL objc.IObject /* cross-framework: NSURL */) {
	if d._PreviewControllerDidSaveEditedCopyOfPreviewItemAtURL != nil {
		d._PreviewControllerDidSaveEditedCopyOfPreviewItemAtURL(controller, previewItem, modifiedContentsURL)
	}
}

// HasPreviewControllerDidSaveEditedCopyOfPreviewItemAtURL returns true if a handler for PreviewControllerDidSaveEditedCopyOfPreviewItemAtURL has been set.
func (d *PreviewControllerDelegate) HasPreviewControllerDidSaveEditedCopyOfPreviewItemAtURL() bool {
	return d._PreviewControllerDidSaveEditedCopyOfPreviewItemAtURL != nil
}

// PreviewControllerDidUpdateContentsOfPreviewItem implements the PPreviewControllerDelegate interface.
func (d *PreviewControllerDelegate) PreviewControllerDidUpdateContentsOfPreviewItem(controller IQLPreviewController, previewItem unsafe.Pointer) {
	if d._PreviewControllerDidUpdateContentsOfPreviewItem != nil {
		d._PreviewControllerDidUpdateContentsOfPreviewItem(controller, previewItem)
	}
}

// HasPreviewControllerDidUpdateContentsOfPreviewItem returns true if a handler for PreviewControllerDidUpdateContentsOfPreviewItem has been set.
func (d *PreviewControllerDelegate) HasPreviewControllerDidUpdateContentsOfPreviewItem() bool {
	return d._PreviewControllerDidUpdateContentsOfPreviewItem != nil
}

// PreviewControllerEditingModeForPreviewItem implements the PPreviewControllerDelegate interface.
func (d *PreviewControllerDelegate) PreviewControllerEditingModeForPreviewItem(controller IQLPreviewController, previewItem unsafe.Pointer) PreviewItemEditingMode {
	if d._PreviewControllerEditingModeForPreviewItem != nil {
		return d._PreviewControllerEditingModeForPreviewItem(controller, previewItem)
	}
	var zero PreviewItemEditingMode
	return zero
}

// HasPreviewControllerEditingModeForPreviewItem returns true if a handler for PreviewControllerEditingModeForPreviewItem has been set.
func (d *PreviewControllerDelegate) HasPreviewControllerEditingModeForPreviewItem() bool {
	return d._PreviewControllerEditingModeForPreviewItem != nil
}

// PreviewControllerFrameForPreviewItemInSourceView implements the PPreviewControllerDelegate interface.
func (d *PreviewControllerDelegate) PreviewControllerFrameForPreviewItemInSourceView(controller IQLPreviewController, item unsafe.Pointer, view unsafe.Pointer) corefoundation.CGRect {
	if d._PreviewControllerFrameForPreviewItemInSourceView != nil {
		return d._PreviewControllerFrameForPreviewItemInSourceView(controller, item, view)
	}
	var zero corefoundation.CGRect
	return zero
}

// HasPreviewControllerFrameForPreviewItemInSourceView returns true if a handler for PreviewControllerFrameForPreviewItemInSourceView has been set.
func (d *PreviewControllerDelegate) HasPreviewControllerFrameForPreviewItemInSourceView() bool {
	return d._PreviewControllerFrameForPreviewItemInSourceView != nil
}

// PreviewControllerShouldOpenURLForPreviewItem implements the PPreviewControllerDelegate interface.
func (d *PreviewControllerDelegate) PreviewControllerShouldOpenURLForPreviewItem(controller IQLPreviewController, url objc.IObject /* cross-framework: NSURL */, item unsafe.Pointer) bool {
	if d._PreviewControllerShouldOpenURLForPreviewItem != nil {
		return d._PreviewControllerShouldOpenURLForPreviewItem(controller, url, item)
	}
	var zero bool
	return zero
}

// HasPreviewControllerShouldOpenURLForPreviewItem returns true if a handler for PreviewControllerShouldOpenURLForPreviewItem has been set.
func (d *PreviewControllerDelegate) HasPreviewControllerShouldOpenURLForPreviewItem() bool {
	return d._PreviewControllerShouldOpenURLForPreviewItem != nil
}

// PreviewControllerTransitionImageForPreviewItemContentRect implements the PPreviewControllerDelegate interface.
func (d *PreviewControllerDelegate) PreviewControllerTransitionImageForPreviewItemContentRect(controller IQLPreviewController, item unsafe.Pointer, contentRect corefoundation.CGRect) appkit.Image {
	if d._PreviewControllerTransitionImageForPreviewItemContentRect != nil {
		return d._PreviewControllerTransitionImageForPreviewItemContentRect(controller, item, contentRect)
	}
	var zero appkit.Image
	return zero
}

// HasPreviewControllerTransitionImageForPreviewItemContentRect returns true if a handler for PreviewControllerTransitionImageForPreviewItemContentRect has been set.
func (d *PreviewControllerDelegate) HasPreviewControllerTransitionImageForPreviewItemContentRect() bool {
	return d._PreviewControllerTransitionImageForPreviewItemContentRect != nil
}

// PreviewControllerTransitionViewForPreviewItem implements the PPreviewControllerDelegate interface.
func (d *PreviewControllerDelegate) PreviewControllerTransitionViewForPreviewItem(controller IQLPreviewController, item unsafe.Pointer) appkit.View {
	if d._PreviewControllerTransitionViewForPreviewItem != nil {
		return d._PreviewControllerTransitionViewForPreviewItem(controller, item)
	}
	var zero appkit.View
	return zero
}

// HasPreviewControllerTransitionViewForPreviewItem returns true if a handler for PreviewControllerTransitionViewForPreviewItem has been set.
func (d *PreviewControllerDelegate) HasPreviewControllerTransitionViewForPreviewItem() bool {
	return d._PreviewControllerTransitionViewForPreviewItem != nil
}

// PreviewControllerDidDismiss implements the PPreviewControllerDelegate interface.
func (d *PreviewControllerDelegate) PreviewControllerDidDismiss(controller IQLPreviewController) {
	if d._PreviewControllerDidDismiss != nil {
		d._PreviewControllerDidDismiss(controller)
	}
}

// HasPreviewControllerDidDismiss returns true if a handler for PreviewControllerDidDismiss has been set.
func (d *PreviewControllerDelegate) HasPreviewControllerDidDismiss() bool {
	return d._PreviewControllerDidDismiss != nil
}

// PreviewControllerWillDismiss implements the PPreviewControllerDelegate interface.
func (d *PreviewControllerDelegate) PreviewControllerWillDismiss(controller IQLPreviewController) {
	if d._PreviewControllerWillDismiss != nil {
		d._PreviewControllerWillDismiss(controller)
	}
}

// HasPreviewControllerWillDismiss returns true if a handler for PreviewControllerWillDismiss has been set.
func (d *PreviewControllerDelegate) HasPreviewControllerWillDismiss() bool {
	return d._PreviewControllerWillDismiss != nil
}
