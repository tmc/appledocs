// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
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
	TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes(textLayoutManager ITextLayoutManager, link objc.IObject, location unsafe.Pointer, renderingAttributes foundation.IDictionary) foundation.IDictionary
	HasTextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes() bool
	TextLayoutManagerShouldBreakLineBeforeLocationHyphenating(textLayoutManager ITextLayoutManager, location unsafe.Pointer, hyphenating bool) bool
	HasTextLayoutManagerShouldBreakLineBeforeLocationHyphenating() bool
	TextLayoutManagerTextLayoutFragmentForLocationInTextElement(textLayoutManager ITextLayoutManager, location unsafe.Pointer, textElement ITextElement) TextLayoutFragment
	HasTextLayoutManagerTextLayoutFragmentForLocationInTextElement() bool
}

// TextLayoutManagerDelegate is a delegate implementation builder for the PTextLayoutManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextLayoutManagerDelegate struct {
	_TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes func(textLayoutManager ITextLayoutManager, link objc.IObject, location unsafe.Pointer, renderingAttributes foundation.IDictionary) foundation.IDictionary
	_TextLayoutManagerShouldBreakLineBeforeLocationHyphenating func(textLayoutManager ITextLayoutManager, location unsafe.Pointer, hyphenating bool) bool
	_TextLayoutManagerTextLayoutFragmentForLocationInTextElement func(textLayoutManager ITextLayoutManager, location unsafe.Pointer, textElement ITextElement) TextLayoutFragment
}

// SetTextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes sets the handler for the TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes delegate method.
//
// The method the framework calls to return a dictionary of attributes for rendering a link attribute name.
func (d *TextLayoutManagerDelegate) SetTextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes(f func(textLayoutManager ITextLayoutManager, link objc.IObject, location unsafe.Pointer, renderingAttributes foundation.IDictionary) foundation.IDictionary) {
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
func (d *TextLayoutManagerDelegate) SetTextLayoutManagerTextLayoutFragmentForLocationInTextElement(f func(textLayoutManager ITextLayoutManager, location unsafe.Pointer, textElement ITextElement) TextLayoutFragment) {
	d._TextLayoutManagerTextLayoutFragmentForLocationInTextElement = f
}

// TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes implements the PTextLayoutManagerDelegate interface.
func (d *TextLayoutManagerDelegate) TextLayoutManagerRenderingAttributesForLinkAtLocationDefaultAttributes(textLayoutManager ITextLayoutManager, link objc.IObject, location unsafe.Pointer, renderingAttributes foundation.IDictionary) foundation.IDictionary {
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
func (d *TextLayoutManagerDelegate) TextLayoutManagerTextLayoutFragmentForLocationInTextElement(textLayoutManager ITextLayoutManager, location unsafe.Pointer, textElement ITextElement) TextLayoutFragment {
	if d._TextLayoutManagerTextLayoutFragmentForLocationInTextElement != nil {
		return d._TextLayoutManagerTextLayoutFragmentForLocationInTextElement(textLayoutManager, location, textElement)
	}
	var zero TextLayoutFragment
	return zero
}

// HasTextLayoutManagerTextLayoutFragmentForLocationInTextElement returns true if a handler for TextLayoutManagerTextLayoutFragmentForLocationInTextElement has been set.
func (d *TextLayoutManagerDelegate) HasTextLayoutManagerTextLayoutFragmentForLocationInTextElement() bool {
	return d._TextLayoutManagerTextLayoutFragmentForLocationInTextElement != nil
}
