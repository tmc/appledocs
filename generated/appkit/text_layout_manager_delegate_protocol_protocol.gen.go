// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PTextLayoutManagerDelegate is the NSTextLayoutManagerDelegate protocol interface.
//
// Optional methods that delegates implement to respond to layout changes.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextLayoutManagerDelegate
type PTextLayoutManagerDelegate interface {
	// Optional methods
	TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes(textLayoutManager ITextLayoutManager, link objectivec.IObject, location unsafe.Pointer, renderingAttributes foundation.IDictionary) foundation.IDictionary
	HasTextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes() bool
	TextLayoutManagerShouldBreakLineBeforeLocationHyphenating(textLayoutManager ITextLayoutManager, location unsafe.Pointer, hyphenating bool) bool
	HasTextLayoutManagerShouldBreakLineBeforeLocationHyphenating() bool
	TextLayoutManagerTextLayoutFragmentForLocationInTextElement(textLayoutManager ITextLayoutManager, location unsafe.Pointer, textElement ITextElement) ITextLayoutFragment
	HasTextLayoutManagerTextLayoutFragmentForLocationInTextElement() bool
}

// TextLayoutManagerDelegate is a delegate implementation builder for the PTextLayoutManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextLayoutManagerDelegate struct {
	_TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes func(textLayoutManager ITextLayoutManager, link objectivec.IObject, location unsafe.Pointer, renderingAttributes foundation.IDictionary) foundation.IDictionary
	_TextLayoutManagerShouldBreakLineBeforeLocationHyphenating func(textLayoutManager ITextLayoutManager, location unsafe.Pointer, hyphenating bool) bool
	_TextLayoutManagerTextLayoutFragmentForLocationInTextElement func(textLayoutManager ITextLayoutManager, location unsafe.Pointer, textElement ITextElement) ITextLayoutFragment
}

// SetTextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes sets the handler for the TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes delegate method.
//
// The method the framework calls to return a dictionary of attributes for rendering a link attribute name.
func (d *TextLayoutManagerDelegate) SetTextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes(f func(textLayoutManager ITextLayoutManager, link objectivec.IObject, location unsafe.Pointer, renderingAttributes foundation.IDictionary) foundation.IDictionary) {
	d._TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes = f
}

// SetTextLayoutManagerShouldBreakLineBeforeLocationHyphenating sets the handler for the TextLayoutManagerShouldBreakLineBeforeLocationHyphenating delegate method.
//
// The method the framework calls to determine the soft line break point.
func (d *TextLayoutManagerDelegate) SetTextLayoutManagerShouldBreakLineBeforeLocationHyphenating(f func(textLayoutManager ITextLayoutManager, location unsafe.Pointer, hyphenating bool) bool) {
	d._TextLayoutManagerShouldBreakLineBeforeLocationHyphenating = f
}

// SetTextLayoutManagerTextLayoutFragmentForLocationInTextElement sets the handler for the TextLayoutManagerTextLayoutFragmentForLocationInTextElement delegate method.
//
// The method the framework calls to give the delegate an opportunity to return a custom text layout fragment.
func (d *TextLayoutManagerDelegate) SetTextLayoutManagerTextLayoutFragmentForLocationInTextElement(f func(textLayoutManager ITextLayoutManager, location unsafe.Pointer, textElement ITextElement) ITextLayoutFragment) {
	d._TextLayoutManagerTextLayoutFragmentForLocationInTextElement = f
}

// TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes implements the PTextLayoutManagerDelegate interface.
func (d *TextLayoutManagerDelegate) TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes(textLayoutManager ITextLayoutManager, link objectivec.IObject, location unsafe.Pointer, renderingAttributes foundation.IDictionary) foundation.IDictionary {
	if d._TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes != nil {
		return d._TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes(textLayoutManager, link, location, renderingAttributes)
	}
	var zero foundation.IDictionary
	return zero
}

// HasTextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes returns true if a handler for TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes has been set.
func (d *TextLayoutManagerDelegate) HasTextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes() bool {
	return d._TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes != nil
}

// TextLayoutManagerShouldBreakLineBeforeLocationHyphenating implements the PTextLayoutManagerDelegate interface.
func (d *TextLayoutManagerDelegate) TextLayoutManagerShouldBreakLineBeforeLocationHyphenating(textLayoutManager ITextLayoutManager, location unsafe.Pointer, hyphenating bool) bool {
	if d._TextLayoutManagerShouldBreakLineBeforeLocationHyphenating != nil {
		return d._TextLayoutManagerShouldBreakLineBeforeLocationHyphenating(textLayoutManager, location, hyphenating)
	}
	var zero bool
	return zero
}

// HasTextLayoutManagerShouldBreakLineBeforeLocationHyphenating returns true if a handler for TextLayoutManagerShouldBreakLineBeforeLocationHyphenating has been set.
func (d *TextLayoutManagerDelegate) HasTextLayoutManagerShouldBreakLineBeforeLocationHyphenating() bool {
	return d._TextLayoutManagerShouldBreakLineBeforeLocationHyphenating != nil
}

// TextLayoutManagerTextLayoutFragmentForLocationInTextElement implements the PTextLayoutManagerDelegate interface.
func (d *TextLayoutManagerDelegate) TextLayoutManagerTextLayoutFragmentForLocationInTextElement(textLayoutManager ITextLayoutManager, location unsafe.Pointer, textElement ITextElement) ITextLayoutFragment {
	if d._TextLayoutManagerTextLayoutFragmentForLocationInTextElement != nil {
		return d._TextLayoutManagerTextLayoutFragmentForLocationInTextElement(textLayoutManager, location, textElement)
	}
	var zero ITextLayoutFragment
	return zero
}

// HasTextLayoutManagerTextLayoutFragmentForLocationInTextElement returns true if a handler for TextLayoutManagerTextLayoutFragmentForLocationInTextElement has been set.
func (d *TextLayoutManagerDelegate) HasTextLayoutManagerTextLayoutFragmentForLocationInTextElement() bool {
	return d._TextLayoutManagerTextLayoutFragmentForLocationInTextElement != nil
}

// TextLayoutManagerDelegateObject wraps an existing Objective-C object that conforms to the PTextLayoutManagerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TextLayoutManagerDelegateObject struct {
	objectivec.Object
}

// NewTextLayoutManagerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTextLayoutManagerDelegate protocol.
func NewTextLayoutManagerDelegateObject(obj objectivec.Object) *TextLayoutManagerDelegateObject {
	return &TextLayoutManagerDelegateObject{obj}
}

// Make sure TextLayoutManagerDelegateObject implements PTextLayoutManagerDelegate.
var _ PTextLayoutManagerDelegate = (*TextLayoutManagerDelegateObject)(nil)

// TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes implements the PTextLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextLayoutManagerDelegateObject) TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes(textLayoutManager ITextLayoutManager, link objectivec.IObject, location unsafe.Pointer, renderingAttributes foundation.IDictionary) foundation.IDictionary {
	return objc.Send[foundation.IDictionary](o.ID, objc.Sel("textLayoutManager:renderingAttributesForLink:atLocation:defaultAttributes:"), textLayoutManager, link, location, renderingAttributes)
}

// HasTextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes returns true; this is a placeholder for optional method checks.
func (o *TextLayoutManagerDelegateObject) HasTextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextLayoutManagerShouldBreakLineBeforeLocationHyphenating implements the PTextLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextLayoutManagerDelegateObject) TextLayoutManagerShouldBreakLineBeforeLocationHyphenating(textLayoutManager ITextLayoutManager, location unsafe.Pointer, hyphenating bool) bool {
	return objc.Send[bool](o.ID, objc.Sel("textLayoutManager:shouldBreakLineBeforeLocation:hyphenating:"), textLayoutManager, location, hyphenating)
}

// HasTextLayoutManagerShouldBreakLineBeforeLocationHyphenating returns true; this is a placeholder for optional method checks.
func (o *TextLayoutManagerDelegateObject) HasTextLayoutManagerShouldBreakLineBeforeLocationHyphenating() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextLayoutManagerTextLayoutFragmentForLocationInTextElement implements the PTextLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextLayoutManagerDelegateObject) TextLayoutManagerTextLayoutFragmentForLocationInTextElement(textLayoutManager ITextLayoutManager, location unsafe.Pointer, textElement ITextElement) ITextLayoutFragment {
	return objc.Send[ITextLayoutFragment](o.ID, objc.Sel("textLayoutManager:textLayoutFragmentForLocation:inTextElement:"), textLayoutManager, location, textElement)
}

// HasTextLayoutManagerTextLayoutFragmentForLocationInTextElement returns true; this is a placeholder for optional method checks.
func (o *TextLayoutManagerDelegateObject) HasTextLayoutManagerTextLayoutFragmentForLocationInTextElement() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
