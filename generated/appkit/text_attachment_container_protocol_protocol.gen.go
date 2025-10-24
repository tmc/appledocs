// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PTextAttachmentContainer is the NSTextAttachmentContainer protocol interface.
//
// A set of methods that defines the interface to text attachment objects from a layout manager.
//
// Availability:
//   - macOS 10.11+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextAttachmentContainer
type PTextAttachmentContainer interface {
	// Required methods
	AttachmentBoundsForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer ITextContainer, lineFrag corefoundation.CGRect, position corefoundation.CGPoint, charIndex uint) corefoundation.CGRect/* debug [protocol_interface/required_method]: AttachmentBoundsForTextContainerProposedLineFragmentGlyphPositionCharacterIndex */
	ImageForBoundsTextContainerCharacterIndex(imageBounds corefoundation.CGRect, textContainer ITextContainer, charIndex uint) Image/* debug [protocol_interface/required_method]: ImageForBoundsTextContainerCharacterIndex */
}
