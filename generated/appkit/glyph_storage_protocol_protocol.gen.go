// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PGlyphStorage is the NSGlyphStorage protocol interface.
//
// A set of methods that a glyph storage object must implement to interact properly with  .
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSGlyphStorage
type PGlyphStorage interface {
	// Required methods
	AttributedString() foundation.AttributedString
	InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex(glyphs objc.IObject /* cross-framework: Glyph */, length uint, glyphIndex uint, charIndex uint)
	LayoutOptions() uint
	SetIntAttributeValueForGlyphAtIndex(attributeTag int, val int, glyphIndex uint)
}
