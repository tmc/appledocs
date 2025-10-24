// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
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
	AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition(attributes foundation.IDictionary, location objc.IObject, textContainer ITextContainer, proposedLineFragment corefoundation.CGRect, position corefoundation.CGPoint) corefoundation.CGRect
	ImageForBoundsAttributesLocationTextContainer(bounds corefoundation.CGRect, attributes foundation.IDictionary, location objc.IObject, textContainer ITextContainer) Image
	ViewProviderForParentViewLocationTextContainer(parentView IView, location objc.IObject, textContainer ITextContainer) TextAttachmentViewProvider
}
