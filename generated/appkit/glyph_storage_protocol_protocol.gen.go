// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/foundation"
)

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
	AttributedString() foundation.AttributedString/* debug [protocol_interface/required_method]: AttributedString */
	InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex(glyphs Glyph /* typedef */, length uint, glyphIndex uint, charIndex uint)/* debug [protocol_interface/required_method]: InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex */
	LayoutOptions() uint/* debug [protocol_interface/required_method]: LayoutOptions */
	SetIntAttributeValueForGlyphAtIndex(attributeTag int, val int, glyphIndex uint)/* debug [protocol_interface/required_method]: SetIntAttributeValueForGlyphAtIndex */
}
