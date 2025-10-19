// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextAttachment] class.
var (
	textAttachmentClass     _TextAttachmentClass
	textAttachmentClassOnce sync.Once
)

func getTextAttachmentClass() _TextAttachmentClass {
	textAttachmentClassOnce.Do(func() {
		textAttachmentClass = _TextAttachmentClass{objc.GetClass("NSTextAttachment")}
	})
	return textAttachmentClass
}

type _TextAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [TextAttachment] class.
type ITextAttachment interface {
	objectivec.IObject
}

// The values for the attachment characteristics of attributed strings and related objects.
//
// The class uses text attachment objects as the values for attachment attributes (stored in the attributed string under the key). A text attachment object contains either an object or an object, which in turn holds the contents of the attached file. The properties of this class configure the appearance of the text attachment in your interface. In macOS, the text attachment also uses a cell object that conforms to the protocol to draw the image that represents the text and handles mouse events. For more information about text attachments, see the and . In macOS 12 and iOS 15 and later, and provide additional capabilities to represent document locations in terms of an  or an  , and provide support for view-based text attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment
type TextAttachment struct {
	objectivec.Object
}

// TextAttachmentFrom constructs a [TextAttachment] from an unsafe.Pointer.
//
// The values for the attachment characteristics of attributed strings and related objects.
func TextAttachmentFrom(ptr unsafe.Pointer) TextAttachment {
	return TextAttachment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextAttachmentClass) Alloc() TextAttachment {
	rv := objc.Send[TextAttachment](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextAttachmentClass) New() TextAttachment {
	rv := objc.Send[TextAttachment](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextAttachment) Init() TextAttachment {
	rv := objc.Send[TextAttachment](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextAttachment) Autorelease() TextAttachment {
	rv := objc.Send[TextAttachment](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextAttachment creates a new TextAttachment instance.
func NewTextAttachment() TextAttachment {
	return getTextAttachmentClass().New()
}




