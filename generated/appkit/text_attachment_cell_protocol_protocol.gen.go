// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/vision"
)

// PTextAttachmentCell is the NSTextAttachmentCell protocol interface.
//
// A set of methods that declares the interface for objects that draw text attachment icons and handle mouse events on their icons.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextAttachmentCellProtocol
type PTextAttachmentCell interface {
	// Required methods
	CellBaselineOffset() vision.Point/* debug [protocol_interface/required_method]: CellBaselineOffset */
	CellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer ITextContainer, lineFrag Rect /* not a class type */, position vision.Point, charIndex uint) Rect/* debug [protocol_interface/required_method]: CellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex */
	CellSize() Size/* debug [protocol_interface/required_method]: CellSize */
	DrawWithFrameInView(cellFrame Rect /* not a class type */, controlView IView)/* debug [protocol_interface/required_method]: DrawWithFrameInView */
	DrawWithFrameInViewCharacterIndex(cellFrame Rect /* not a class type */, controlView IView, charIndex uint)/* debug [protocol_interface/required_method]: DrawWithFrameInViewCharacterIndex */
	DrawWithFrameInViewCharacterIndexLayoutManager(cellFrame Rect /* not a class type */, controlView IView, charIndex uint, layoutManager ILayoutManager)/* debug [protocol_interface/required_method]: DrawWithFrameInViewCharacterIndexLayoutManager */
	HighlightWithFrameInView(flag bool, cellFrame Rect /* not a class type */, controlView IView)/* debug [protocol_interface/required_method]: HighlightWithFrameInView */
	TrackMouseInRectOfViewAtCharacterIndexUntilMouseUp(theEvent IEvent, cellFrame Rect /* not a class type */, controlView IView, charIndex uint, flag bool) bool/* debug [protocol_interface/required_method]: TrackMouseInRectOfViewAtCharacterIndexUntilMouseUp */
	TrackMouseInRectOfViewUntilMouseUp(theEvent IEvent, cellFrame Rect /* not a class type */, controlView IView, flag bool) bool/* debug [protocol_interface/required_method]: TrackMouseInRectOfViewUntilMouseUp */
	WantsToTrackMouse() bool/* debug [protocol_interface/required_method]: WantsToTrackMouse */
	WantsToTrackMouseForEventInRectOfViewAtCharacterIndex(theEvent IEvent, cellFrame Rect /* not a class type */, controlView IView, charIndex uint) bool/* debug [protocol_interface/required_method]: WantsToTrackMouseForEventInRectOfViewAtCharacterIndex */
}
