// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"
)

// PTextAttachmentLayout is the NSTextAttachmentLayout protocol interface.
//
// A set of methods that defines the interface to attachment objects from a text layout manager.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextAttachmentLayout
type PTextAttachmentLayout interface {
	// Required methods
	AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition(attributes foundation.IDictionary, location unsafe.Pointer, textContainer ITextContainer, proposedLineFragment corefoundation.CGRect, position corefoundation.CGPoint) corefoundation.CGRect/* debug [protocol_interface/required_method]: AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition */
	ImageForBoundsAttributesLocationTextContainer(bounds corefoundation.CGRect, attributes foundation.IDictionary, location unsafe.Pointer, textContainer ITextContainer) Image/* debug [protocol_interface/required_method]: ImageForBoundsAttributesLocationTextContainer */
	ViewProviderForParentViewLocationTextContainer(parentView IView, location unsafe.Pointer, textContainer ITextContainer) TextAttachmentViewProvider/* debug [protocol_interface/required_method]: ViewProviderForParentViewLocationTextContainer */
}
