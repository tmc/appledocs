// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/corefoundation"
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
	CellBaselineOffset() corefoundation.CGPoint
	CellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer ITextContainer, lineFrag corefoundation.CGRect, position corefoundation.CGPoint, charIndex uint) corefoundation.CGRect
	CellSize() corefoundation.CGSize
	DrawWithFrameInView(cellFrame corefoundation.CGRect, controlView IView)
	DrawWithFrameInViewCharacterIndex(cellFrame corefoundation.CGRect, controlView IView, charIndex uint)
	DrawWithFrameInViewCharacterIndexLayoutManager(cellFrame corefoundation.CGRect, controlView IView, charIndex uint, layoutManager ILayoutManager)
	HighlightWithFrameInView(flag bool, cellFrame corefoundation.CGRect, controlView IView)
	TrackMouseInRectOfViewAtCharacterIndexUntilMouseUp(theEvent IEvent, cellFrame corefoundation.CGRect, controlView IView, charIndex uint, flag bool) bool
	TrackMouseInRectOfViewUntilMouseUp(theEvent IEvent, cellFrame corefoundation.CGRect, controlView IView, flag bool) bool
	WantsToTrackMouse() bool
	WantsToTrackMouseForEventInRectOfViewAtCharacterIndex(theEvent IEvent, cellFrame corefoundation.CGRect, controlView IView, charIndex uint) bool
}
