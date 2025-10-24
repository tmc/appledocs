// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

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
	CellBaselineOffset() corefoundation.Point
	CellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer ITextContainer, lineFrag objc.IObject /* cross-framework: Rect */, position objc.IObject /* cross-framework: Point */, charIndex uint) corefoundation.Rect
	CellSize() corefoundation.Size
	DrawWithFrameInView(cellFrame objc.IObject /* cross-framework: Rect */, controlView IView)
	DrawWithFrameInViewCharacterIndex(cellFrame objc.IObject /* cross-framework: Rect */, controlView IView, charIndex uint)
	DrawWithFrameInViewCharacterIndexLayoutManager(cellFrame objc.IObject /* cross-framework: Rect */, controlView IView, charIndex uint, layoutManager ILayoutManager)
	HighlightWithFrameInView(flag bool, cellFrame objc.IObject /* cross-framework: Rect */, controlView IView)
	TrackMouseInRectOfViewAtCharacterIndexUntilMouseUp(theEvent IEvent, cellFrame objc.IObject /* cross-framework: Rect */, controlView IView, charIndex uint, flag bool) bool
	TrackMouseInRectOfViewUntilMouseUp(theEvent IEvent, cellFrame objc.IObject /* cross-framework: Rect */, controlView IView, flag bool) bool
	WantsToTrackMouse() bool
	WantsToTrackMouseForEventInRectOfViewAtCharacterIndex(theEvent IEvent, cellFrame objc.IObject /* cross-framework: Rect */, controlView IView, charIndex uint) bool
}
