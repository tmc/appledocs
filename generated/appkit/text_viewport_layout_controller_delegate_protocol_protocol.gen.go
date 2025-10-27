// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(textViewportLayoutController ITextViewportLayoutController, textLayoutFragment ITextLayoutFragment)
	ViewportBoundsForTextViewportLayoutController(textViewportLayoutController ITextViewportLayoutController) corefoundation.CGRect
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

// TextViewportLayoutControllerDelegateObject wraps an existing Objective-C object that conforms to the PTextViewportLayoutControllerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TextViewportLayoutControllerDelegateObject struct {
	objectivec.Object
}

// NewTextViewportLayoutControllerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTextViewportLayoutControllerDelegate protocol.
func NewTextViewportLayoutControllerDelegateObject(obj objectivec.Object) *TextViewportLayoutControllerDelegateObject {
	return &TextViewportLayoutControllerDelegateObject{obj}
}

// Make sure TextViewportLayoutControllerDelegateObject implements PTextViewportLayoutControllerDelegate.
var _ PTextViewportLayoutControllerDelegate = (*TextViewportLayoutControllerDelegateObject)(nil)

// TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment implements the PTextViewportLayoutControllerDelegate interface.
// This required method is always available on objects conforming to TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment.
func (o *TextViewportLayoutControllerDelegateObject) TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(textViewportLayoutController ITextViewportLayoutController, textLayoutFragment ITextLayoutFragment) {
	objc.Send[objc.ID](o.ID, objc.Sel("textViewportLayoutController:configureRenderingSurfaceForTextLayoutFragment:"), textViewportLayoutController, textLayoutFragment)
}

// ViewportBoundsForTextViewportLayoutController implements the PTextViewportLayoutControllerDelegate interface.
// This required method is always available on objects conforming to ViewportBoundsForTextViewportLayoutController.
func (o *TextViewportLayoutControllerDelegateObject) ViewportBoundsForTextViewportLayoutController(textViewportLayoutController ITextViewportLayoutController) corefoundation.CGRect {
	return objc.Send[corefoundation.CGRect](o.ID, objc.Sel("viewportBoundsForTextViewportLayoutController:"), textViewportLayoutController)
}

// TextViewportLayoutControllerDidLayout implements the PTextViewportLayoutControllerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewportLayoutControllerDelegateObject) TextViewportLayoutControllerDidLayout(textViewportLayoutController ITextViewportLayoutController) {
	objc.Send[objc.ID](o.ID, objc.Sel("textViewportLayoutControllerDidLayout:"), textViewportLayoutController)
}

// HasTextViewportLayoutControllerDidLayout returns true; this is a placeholder for optional method checks.
func (o *TextViewportLayoutControllerDelegateObject) HasTextViewportLayoutControllerDidLayout() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewportLayoutControllerWillLayout implements the PTextViewportLayoutControllerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewportLayoutControllerDelegateObject) TextViewportLayoutControllerWillLayout(textViewportLayoutController ITextViewportLayoutController) {
	objc.Send[objc.ID](o.ID, objc.Sel("textViewportLayoutControllerWillLayout:"), textViewportLayoutController)
}

// HasTextViewportLayoutControllerWillLayout returns true; this is a placeholder for optional method checks.
func (o *TextViewportLayoutControllerDelegateObject) HasTextViewportLayoutControllerWillLayout() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
