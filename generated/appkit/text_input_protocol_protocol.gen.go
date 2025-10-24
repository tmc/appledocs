// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/vision"
)

// PTextInput is the NSTextInput protocol interface.
//
// A set of methods that text views need to implement to interact properly with the text input management system.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextInput
type PTextInput interface {
	// Required methods
	AttributedSubstringFromRange(range_ corefoundation.Range) foundation.AttributedString/* debug [protocol_interface/required_method]: AttributedSubstringFromRange */
	CharacterIndexForPoint(point vision.Point) uint/* debug [protocol_interface/required_method]: CharacterIndexForPoint */
	FirstRectForCharacterRange(range_ corefoundation.Range) Rect/* debug [protocol_interface/required_method]: FirstRectForCharacterRange */
	HasMarkedText() bool/* debug [protocol_interface/required_method]: HasMarkedText */
	MarkedRange() corefoundation.Range/* debug [protocol_interface/required_method]: MarkedRange */
	SelectedRange() corefoundation.Range/* debug [protocol_interface/required_method]: SelectedRange */
	SetMarkedTextSelectedRange(string_ objc.IObject, selRange corefoundation.Range)/* debug [protocol_interface/required_method]: SetMarkedTextSelectedRange */
	UnmarkText()/* debug [protocol_interface/required_method]: UnmarkText */
	ValidAttributesForMarkedText() foundation.Array/* debug [protocol_interface/required_method]: ValidAttributesForMarkedText */
}
