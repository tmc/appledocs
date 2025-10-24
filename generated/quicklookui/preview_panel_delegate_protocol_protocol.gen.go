// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/appkit"
)

// PPreviewPanelDelegate is the QLPreviewPanelDelegate protocol interface.
//
// A protocol for the delegate of the Quick Look preview panel.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.quicklookui/documentation/QuickLookUI/QLPreviewPanelDelegate
type PPreviewPanelDelegate interface {
	// Optional methods
	PreviewPanelHandleEvent(panel IQLPreviewPanel, event appkit.Event) bool
	HasPreviewPanelHandleEvent() bool
	PreviewPanelSourceFrameOnScreenForPreviewItem(panel IQLPreviewPanel, item unsafe.Pointer) Rect
	HasPreviewPanelSourceFrameOnScreenForPreviewItem() bool
	PreviewPanelTransitionImageForPreviewItemContentRect(panel IQLPreviewPanel, item unsafe.Pointer, contentRect Rect /* not a class type */) objc.ID
	HasPreviewPanelTransitionImageForPreviewItemContentRect() bool
}

// PreviewPanelDelegate is a delegate implementation builder for the PPreviewPanelDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PreviewPanelDelegate struct {
	_PreviewPanelHandleEvent func(panel IQLPreviewPanel, event appkit.Event) bool
	_PreviewPanelSourceFrameOnScreenForPreviewItem func(panel IQLPreviewPanel, item unsafe.Pointer) Rect
	_PreviewPanelTransitionImageForPreviewItemContentRect func(panel IQLPreviewPanel, item unsafe.Pointer, contentRect Rect /* not a class type */) objc.ID
}

// SetPreviewPanelHandleEvent sets the handler for the PreviewPanelHandleEvent delegate method.
//
// Handles an event that the preview panel receives, but doesn’t handle.
func (d *PreviewPanelDelegate) SetPreviewPanelHandleEvent(f func(panel IQLPreviewPanel, event appkit.Event) bool) {
	d._PreviewPanelHandleEvent = f
}

// SetPreviewPanelSourceFrameOnScreenForPreviewItem sets the handler for the PreviewPanelSourceFrameOnScreenForPreviewItem delegate method.
//
// Returns the screen rectangle for a given preview item.
func (d *PreviewPanelDelegate) SetPreviewPanelSourceFrameOnScreenForPreviewItem(f func(panel IQLPreviewPanel, item unsafe.Pointer) Rect) {
	d._PreviewPanelSourceFrameOnScreenForPreviewItem = f
}

// SetPreviewPanelTransitionImageForPreviewItemContentRect sets the handler for the PreviewPanelTransitionImageForPreviewItemContentRect delegate method.
//
// Returns the image to use for the transition zoom effect for a given item.
func (d *PreviewPanelDelegate) SetPreviewPanelTransitionImageForPreviewItemContentRect(f func(panel IQLPreviewPanel, item unsafe.Pointer, contentRect Rect /* not a class type */) objc.ID) {
	d._PreviewPanelTransitionImageForPreviewItemContentRect = f
}

// PreviewPanelHandleEvent implements the PPreviewPanelDelegate interface.
func (d *PreviewPanelDelegate) PreviewPanelHandleEvent(panel IQLPreviewPanel, event appkit.Event) bool {
	if d._PreviewPanelHandleEvent != nil {
		return d._PreviewPanelHandleEvent(panel, event)
	}
	var zero bool
	return zero
}

// HasPreviewPanelHandleEvent returns true if a handler for PreviewPanelHandleEvent has been set.
func (d *PreviewPanelDelegate) HasPreviewPanelHandleEvent() bool {
	return d._PreviewPanelHandleEvent != nil
}

// PreviewPanelSourceFrameOnScreenForPreviewItem implements the PPreviewPanelDelegate interface.
func (d *PreviewPanelDelegate) PreviewPanelSourceFrameOnScreenForPreviewItem(panel IQLPreviewPanel, item unsafe.Pointer) Rect {
	if d._PreviewPanelSourceFrameOnScreenForPreviewItem != nil {
		return d._PreviewPanelSourceFrameOnScreenForPreviewItem(panel, item)
	}
	var zero Rect
	return zero
}

// HasPreviewPanelSourceFrameOnScreenForPreviewItem returns true if a handler for PreviewPanelSourceFrameOnScreenForPreviewItem has been set.
func (d *PreviewPanelDelegate) HasPreviewPanelSourceFrameOnScreenForPreviewItem() bool {
	return d._PreviewPanelSourceFrameOnScreenForPreviewItem != nil
}

// PreviewPanelTransitionImageForPreviewItemContentRect implements the PPreviewPanelDelegate interface.
func (d *PreviewPanelDelegate) PreviewPanelTransitionImageForPreviewItemContentRect(panel IQLPreviewPanel, item unsafe.Pointer, contentRect Rect /* not a class type */) objc.ID {
	if d._PreviewPanelTransitionImageForPreviewItemContentRect != nil {
		return d._PreviewPanelTransitionImageForPreviewItemContentRect(panel, item, contentRect)
	}
	var zero objc.ID
	return zero
}

// HasPreviewPanelTransitionImageForPreviewItemContentRect returns true if a handler for PreviewPanelTransitionImageForPreviewItemContentRect has been set.
func (d *PreviewPanelDelegate) HasPreviewPanelTransitionImageForPreviewItemContentRect() bool {
	return d._PreviewPanelTransitionImageForPreviewItemContentRect != nil
}
