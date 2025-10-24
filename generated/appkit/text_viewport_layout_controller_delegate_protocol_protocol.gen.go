// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PTextViewportLayoutControllerDelegate is the NSTextViewportLayoutControllerDelegate protocol interface.
//
// Optional methods that delegates implement to respond to viewport layout changes.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextViewportLayoutControllerDelegate
type PTextViewportLayoutControllerDelegate interface {
	// Required methods
	TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(textViewportLayoutController ITextViewportLayoutController, textLayoutFragment ITextLayoutFragment)/* debug [protocol_interface/required_method]: TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment */
	ViewportBoundsForTextViewportLayoutController(textViewportLayoutController ITextViewportLayoutController) corefoundation.CGRect/* debug [protocol_interface/required_method]: ViewportBoundsForTextViewportLayoutController */
	// Optional methods
	TextViewportLayoutControllerDidLayout(textViewportLayoutController ITextViewportLayoutController)
	HasTextViewportLayoutControllerDidLayout() bool
	TextViewportLayoutControllerWillLayout(textViewportLayoutController ITextViewportLayoutController)
	HasTextViewportLayoutControllerWillLayout() bool
}

// TextViewportLayoutControllerDelegate is a delegate implementation builder for the PTextViewportLayoutControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextViewportLayoutControllerDelegate struct {
	_TextViewportLayoutControllerDidLayout func(textViewportLayoutController ITextViewportLayoutController)
	_TextViewportLayoutControllerWillLayout func(textViewportLayoutController ITextViewportLayoutController)
	_TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment func(textViewportLayoutController ITextViewportLayoutController, textLayoutFragment ITextLayoutFragment)
	_ViewportBoundsForTextViewportLayoutController func(textViewportLayoutController ITextViewportLayoutController) corefoundation.CGRect
}

// SetTextViewportLayoutControllerDidLayout sets the handler for the TextViewportLayoutControllerDidLayout delegate method.
//
// The method the framework calls when the text viewport layout controller finishes its layout process.
func (d *TextViewportLayoutControllerDelegate) SetTextViewportLayoutControllerDidLayout(f func(textViewportLayoutController ITextViewportLayoutController)) {
	d._TextViewportLayoutControllerDidLayout = f
}

// SetTextViewportLayoutControllerWillLayout sets the handler for the TextViewportLayoutControllerWillLayout delegate method.
//
// The method the framework calls before the text viewport layout controller starts its layout process.
func (d *TextViewportLayoutControllerDelegate) SetTextViewportLayoutControllerWillLayout(f func(textViewportLayoutController ITextViewportLayoutController)) {
	d._TextViewportLayoutControllerWillLayout = f
}

// SetTextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment sets the handler for the TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment delegate method.
//
// The method the framework calls when the layout controller lays out a text layout fragment in the UI.
func (d *TextViewportLayoutControllerDelegate) SetTextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(f func(textViewportLayoutController ITextViewportLayoutController, textLayoutFragment ITextLayoutFragment)) {
	d._TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment = f
}

// SetViewportBoundsForTextViewportLayoutController sets the handler for the ViewportBoundsForTextViewportLayoutController delegate method.
//
// Returns the current viewport, which is the view visible bounds plus the overdraw area.
func (d *TextViewportLayoutControllerDelegate) SetViewportBoundsForTextViewportLayoutController(f func(textViewportLayoutController ITextViewportLayoutController) corefoundation.CGRect) {
	d._ViewportBoundsForTextViewportLayoutController = f
}

// TextViewportLayoutControllerDidLayout implements the PTextViewportLayoutControllerDelegate interface.
func (d *TextViewportLayoutControllerDelegate) TextViewportLayoutControllerDidLayout(textViewportLayoutController ITextViewportLayoutController) {
	if d._TextViewportLayoutControllerDidLayout != nil {
		d._TextViewportLayoutControllerDidLayout(textViewportLayoutController)
	}
}

// HasTextViewportLayoutControllerDidLayout returns true if a handler for TextViewportLayoutControllerDidLayout has been set.
func (d *TextViewportLayoutControllerDelegate) HasTextViewportLayoutControllerDidLayout() bool {
	return d._TextViewportLayoutControllerDidLayout != nil
}

// TextViewportLayoutControllerWillLayout implements the PTextViewportLayoutControllerDelegate interface.
func (d *TextViewportLayoutControllerDelegate) TextViewportLayoutControllerWillLayout(textViewportLayoutController ITextViewportLayoutController) {
	if d._TextViewportLayoutControllerWillLayout != nil {
		d._TextViewportLayoutControllerWillLayout(textViewportLayoutController)
	}
}

// HasTextViewportLayoutControllerWillLayout returns true if a handler for TextViewportLayoutControllerWillLayout has been set.
func (d *TextViewportLayoutControllerDelegate) HasTextViewportLayoutControllerWillLayout() bool {
	return d._TextViewportLayoutControllerWillLayout != nil
}

// TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment implements the PTextViewportLayoutControllerDelegate interface.
func (d *TextViewportLayoutControllerDelegate) TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(textViewportLayoutController ITextViewportLayoutController, textLayoutFragment ITextLayoutFragment) {
	if d._TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment != nil {
		d._TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(textViewportLayoutController, textLayoutFragment)
	}
}

// HasTextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment returns true if a handler for TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment has been set.
func (d *TextViewportLayoutControllerDelegate) HasTextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment() bool {
	return d._TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment != nil
}

// ViewportBoundsForTextViewportLayoutController implements the PTextViewportLayoutControllerDelegate interface.
func (d *TextViewportLayoutControllerDelegate) ViewportBoundsForTextViewportLayoutController(textViewportLayoutController ITextViewportLayoutController) corefoundation.CGRect {
	if d._ViewportBoundsForTextViewportLayoutController != nil {
		return d._ViewportBoundsForTextViewportLayoutController(textViewportLayoutController)
	}
	var zero corefoundation.CGRect
	return zero
}

// HasViewportBoundsForTextViewportLayoutController returns true if a handler for ViewportBoundsForTextViewportLayoutController has been set.
func (d *TextViewportLayoutControllerDelegate) HasViewportBoundsForTextViewportLayoutController() bool {
	return d._ViewportBoundsForTextViewportLayoutController != nil
}
