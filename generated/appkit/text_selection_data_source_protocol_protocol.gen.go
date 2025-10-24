// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PTextSelectionDataSource is the NSTextSelectionDataSource protocol interface.
//
// A set of methods that objects implement to provide data for, and manage text selections.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextSelectionDataSource
type PTextSelectionDataSource interface {
	// Required methods
	BaseWritingDirectionAtLocation(location unsafe.Pointer) TextSelectionNavigationWritingDirection/* debug [protocol_interface/required_method]: BaseWritingDirectionAtLocation */
	EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock(location unsafe.Pointer, block unsafe.Pointer)/* debug [protocol_interface/required_method]: EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock */
	EnumerateSubstringsFromLocationOptionsUsingBlock(location unsafe.Pointer, options StringEnumerationOptions /* not a class type */, block unsafe.Pointer)/* debug [protocol_interface/required_method]: EnumerateSubstringsFromLocationOptionsUsingBlock */
	LineFragmentRangeForPointInContainerAtLocation(point corefoundation.CGPoint, location unsafe.Pointer) TextRange/* debug [protocol_interface/required_method]: LineFragmentRangeForPointInContainerAtLocation */
	LocationFromLocationWithOffset(location unsafe.Pointer, offset int) unsafe.Pointer/* debug [protocol_interface/required_method]: LocationFromLocationWithOffset */
	OffsetFromLocationToLocation(from unsafe.Pointer, to unsafe.Pointer) int/* debug [protocol_interface/required_method]: OffsetFromLocationToLocation */
	TextRangeForSelectionGranularityEnclosingLocation(selectionGranularity TextSelectionGranularity, location unsafe.Pointer) TextRange/* debug [protocol_interface/required_method]: TextRangeForSelectionGranularityEnclosingLocation */
	// Optional methods
	EnumerateContainerBoundariesFromLocationReverseUsingBlock(location unsafe.Pointer, reverse bool, block unsafe.Pointer)
	HasEnumerateContainerBoundariesFromLocationReverseUsingBlock() bool
	TextLayoutOrientationAtLocation(location unsafe.Pointer) TextSelectionNavigationLayoutOrientation
	HasTextLayoutOrientationAtLocation() bool
}

// TextSelectionDataSource is a delegate implementation builder for the PTextSelectionDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextSelectionDataSource struct {
	_EnumerateContainerBoundariesFromLocationReverseUsingBlock func(location unsafe.Pointer, reverse bool, block unsafe.Pointer)
	_TextLayoutOrientationAtLocation func(location unsafe.Pointer) TextSelectionNavigationLayoutOrientation
	_BaseWritingDirectionAtLocation func(location unsafe.Pointer) TextSelectionNavigationWritingDirection
	_EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock func(location unsafe.Pointer, block unsafe.Pointer)
	_EnumerateSubstringsFromLocationOptionsUsingBlock func(location unsafe.Pointer, options StringEnumerationOptions /* not a class type */, block unsafe.Pointer)
	_LineFragmentRangeForPointInContainerAtLocation func(point corefoundation.CGPoint, location unsafe.Pointer) TextRange
	_LocationFromLocationWithOffset func(location unsafe.Pointer, offset int) unsafe.Pointer
	_OffsetFromLocationToLocation func(from unsafe.Pointer, to unsafe.Pointer) int
	_TextRangeForSelectionGranularityEnclosingLocation func(selectionGranularity TextSelectionGranularity, location unsafe.Pointer) TextRange
}

// SetEnumerateContainerBoundariesFromLocationReverseUsingBlock sets the handler for the EnumerateContainerBoundariesFromLocationReverseUsingBlock delegate method.
//
// Enumerates all the container boundaries starting from the location you specify.
func (d *TextSelectionDataSource) SetEnumerateContainerBoundariesFromLocationReverseUsingBlock(f func(location unsafe.Pointer, reverse bool, block unsafe.Pointer)) {
	d._EnumerateContainerBoundariesFromLocationReverseUsingBlock = f
}

// SetTextLayoutOrientationAtLocation sets the handler for the TextLayoutOrientationAtLocation delegate method.
//
// Returns the layout orientation at the location you specify.
func (d *TextSelectionDataSource) SetTextLayoutOrientationAtLocation(f func(location unsafe.Pointer) TextSelectionNavigationLayoutOrientation) {
	d._TextLayoutOrientationAtLocation = f
}

// SetBaseWritingDirectionAtLocation sets the handler for the BaseWritingDirectionAtLocation delegate method.
//
// Returns the base writing direction at the location you specify.
func (d *TextSelectionDataSource) SetBaseWritingDirectionAtLocation(f func(location unsafe.Pointer) TextSelectionNavigationWritingDirection) {
	d._BaseWritingDirectionAtLocation = f
}

// SetEnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock sets the handler for the EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock delegate method.
//
// Enumerates all the insertion point caret offsets from left to right in visual order.
func (d *TextSelectionDataSource) SetEnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock(f func(location unsafe.Pointer, block unsafe.Pointer)) {
	d._EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock = f
}

// SetEnumerateSubstringsFromLocationOptionsUsingBlock sets the handler for the EnumerateSubstringsFromLocationOptionsUsingBlock delegate method.
//
// Enumerates the textual segment boundaries starting at the location you specify.
func (d *TextSelectionDataSource) SetEnumerateSubstringsFromLocationOptionsUsingBlock(f func(location unsafe.Pointer, options StringEnumerationOptions /* not a class type */, block unsafe.Pointer)) {
	d._EnumerateSubstringsFromLocationOptionsUsingBlock = f
}

// SetLineFragmentRangeForPointInContainerAtLocation sets the handler for the LineFragmentRangeForPointInContainerAtLocation delegate method.
//
// Returns the range of the line fragment that contains the point you specify.
func (d *TextSelectionDataSource) SetLineFragmentRangeForPointInContainerAtLocation(f func(point corefoundation.CGPoint, location unsafe.Pointer) TextRange) {
	d._LineFragmentRangeForPointInContainerAtLocation = f
}

// SetLocationFromLocationWithOffset sets the handler for the LocationFromLocationWithOffset delegate method.
//
// Returns a new location using the location and offset you specify.
func (d *TextSelectionDataSource) SetLocationFromLocationWithOffset(f func(location unsafe.Pointer, offset int) unsafe.Pointer) {
	d._LocationFromLocationWithOffset = f
}

// SetOffsetFromLocationToLocation sets the handler for the OffsetFromLocationToLocation delegate method.
//
// Returns the offset between the two locations you specify.
func (d *TextSelectionDataSource) SetOffsetFromLocationToLocation(f func(from unsafe.Pointer, to unsafe.Pointer) int) {
	d._OffsetFromLocationToLocation = f
}

// SetTextRangeForSelectionGranularityEnclosingLocation sets the handler for the TextRangeForSelectionGranularityEnclosingLocation delegate method.
//
// Returns a text range that corresponds to selection granularity of the enclosing location.
func (d *TextSelectionDataSource) SetTextRangeForSelectionGranularityEnclosingLocation(f func(selectionGranularity TextSelectionGranularity, location unsafe.Pointer) TextRange) {
	d._TextRangeForSelectionGranularityEnclosingLocation = f
}

// EnumerateContainerBoundariesFromLocationReverseUsingBlock implements the PTextSelectionDataSource interface.
func (d *TextSelectionDataSource) EnumerateContainerBoundariesFromLocationReverseUsingBlock(location unsafe.Pointer, reverse bool, block unsafe.Pointer) {
	if d._EnumerateContainerBoundariesFromLocationReverseUsingBlock != nil {
		d._EnumerateContainerBoundariesFromLocationReverseUsingBlock(location, reverse, block)
	}
}

// HasEnumerateContainerBoundariesFromLocationReverseUsingBlock returns true if a handler for EnumerateContainerBoundariesFromLocationReverseUsingBlock has been set.
func (d *TextSelectionDataSource) HasEnumerateContainerBoundariesFromLocationReverseUsingBlock() bool {
	return d._EnumerateContainerBoundariesFromLocationReverseUsingBlock != nil
}

// TextLayoutOrientationAtLocation implements the PTextSelectionDataSource interface.
func (d *TextSelectionDataSource) TextLayoutOrientationAtLocation(location unsafe.Pointer) TextSelectionNavigationLayoutOrientation {
	if d._TextLayoutOrientationAtLocation != nil {
		return d._TextLayoutOrientationAtLocation(location)
	}
	var zero TextSelectionNavigationLayoutOrientation
	return zero
}

// HasTextLayoutOrientationAtLocation returns true if a handler for TextLayoutOrientationAtLocation has been set.
func (d *TextSelectionDataSource) HasTextLayoutOrientationAtLocation() bool {
	return d._TextLayoutOrientationAtLocation != nil
}

// BaseWritingDirectionAtLocation implements the PTextSelectionDataSource interface.
func (d *TextSelectionDataSource) BaseWritingDirectionAtLocation(location unsafe.Pointer) TextSelectionNavigationWritingDirection {
	if d._BaseWritingDirectionAtLocation != nil {
		return d._BaseWritingDirectionAtLocation(location)
	}
	var zero TextSelectionNavigationWritingDirection
	return zero
}

// HasBaseWritingDirectionAtLocation returns true if a handler for BaseWritingDirectionAtLocation has been set.
func (d *TextSelectionDataSource) HasBaseWritingDirectionAtLocation() bool {
	return d._BaseWritingDirectionAtLocation != nil
}

// EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock implements the PTextSelectionDataSource interface.
func (d *TextSelectionDataSource) EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock(location unsafe.Pointer, block unsafe.Pointer) {
	if d._EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock != nil {
		d._EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock(location, block)
	}
}

// HasEnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock returns true if a handler for EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock has been set.
func (d *TextSelectionDataSource) HasEnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock() bool {
	return d._EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock != nil
}

// EnumerateSubstringsFromLocationOptionsUsingBlock implements the PTextSelectionDataSource interface.
func (d *TextSelectionDataSource) EnumerateSubstringsFromLocationOptionsUsingBlock(location unsafe.Pointer, options StringEnumerationOptions /* not a class type */, block unsafe.Pointer) {
	if d._EnumerateSubstringsFromLocationOptionsUsingBlock != nil {
		d._EnumerateSubstringsFromLocationOptionsUsingBlock(location, options, block)
	}
}

// HasEnumerateSubstringsFromLocationOptionsUsingBlock returns true if a handler for EnumerateSubstringsFromLocationOptionsUsingBlock has been set.
func (d *TextSelectionDataSource) HasEnumerateSubstringsFromLocationOptionsUsingBlock() bool {
	return d._EnumerateSubstringsFromLocationOptionsUsingBlock != nil
}

// LineFragmentRangeForPointInContainerAtLocation implements the PTextSelectionDataSource interface.
func (d *TextSelectionDataSource) LineFragmentRangeForPointInContainerAtLocation(point corefoundation.CGPoint, location unsafe.Pointer) TextRange {
	if d._LineFragmentRangeForPointInContainerAtLocation != nil {
		return d._LineFragmentRangeForPointInContainerAtLocation(point, location)
	}
	var zero TextRange
	return zero
}

// HasLineFragmentRangeForPointInContainerAtLocation returns true if a handler for LineFragmentRangeForPointInContainerAtLocation has been set.
func (d *TextSelectionDataSource) HasLineFragmentRangeForPointInContainerAtLocation() bool {
	return d._LineFragmentRangeForPointInContainerAtLocation != nil
}

// LocationFromLocationWithOffset implements the PTextSelectionDataSource interface.
func (d *TextSelectionDataSource) LocationFromLocationWithOffset(location unsafe.Pointer, offset int) unsafe.Pointer {
	if d._LocationFromLocationWithOffset != nil {
		return d._LocationFromLocationWithOffset(location, offset)
	}
	var zero unsafe.Pointer
	return zero
}

// HasLocationFromLocationWithOffset returns true if a handler for LocationFromLocationWithOffset has been set.
func (d *TextSelectionDataSource) HasLocationFromLocationWithOffset() bool {
	return d._LocationFromLocationWithOffset != nil
}

// OffsetFromLocationToLocation implements the PTextSelectionDataSource interface.
func (d *TextSelectionDataSource) OffsetFromLocationToLocation(from unsafe.Pointer, to unsafe.Pointer) int {
	if d._OffsetFromLocationToLocation != nil {
		return d._OffsetFromLocationToLocation(from, to)
	}
	var zero int
	return zero
}

// HasOffsetFromLocationToLocation returns true if a handler for OffsetFromLocationToLocation has been set.
func (d *TextSelectionDataSource) HasOffsetFromLocationToLocation() bool {
	return d._OffsetFromLocationToLocation != nil
}

// TextRangeForSelectionGranularityEnclosingLocation implements the PTextSelectionDataSource interface.
func (d *TextSelectionDataSource) TextRangeForSelectionGranularityEnclosingLocation(selectionGranularity TextSelectionGranularity, location unsafe.Pointer) TextRange {
	if d._TextRangeForSelectionGranularityEnclosingLocation != nil {
		return d._TextRangeForSelectionGranularityEnclosingLocation(selectionGranularity, location)
	}
	var zero TextRange
	return zero
}

// HasTextRangeForSelectionGranularityEnclosingLocation returns true if a handler for TextRangeForSelectionGranularityEnclosingLocation has been set.
func (d *TextSelectionDataSource) HasTextRangeForSelectionGranularityEnclosingLocation() bool {
	return d._TextRangeForSelectionGranularityEnclosingLocation != nil
}
