// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PreviewReplyAttachment] class.
var (
	PreviewReplyAttachmentClass     _PreviewReplyAttachmentClass
	PreviewReplyAttachmentClassOnce sync.Once
)

func getPreviewReplyAttachmentClass() _PreviewReplyAttachmentClass {
	PreviewReplyAttachmentClassOnce.Do(func() {
		PreviewReplyAttachmentClass = _PreviewReplyAttachmentClass{objc.GetClass("QLPreviewReplyAttachment")}
	})
	return PreviewReplyAttachmentClass
}

type _PreviewReplyAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [PreviewReplyAttachment] class.
type IPreviewReplyAttachment interface {
	objectivec.IObject
}

// An attachment for a Quick Look preview reply that provides additional content for the system to display a preview.
//
// When providing a data-based Quick Look preview with HTML, use to include images, CSS, and other linked content in the HTML of the preview. Reference content in your html using the notation for the reference. For instance, if your HTML preview response includes an image, create a with the image, add it the reply’s with an associated string, and reference the image with the associated string, prefixed by . The following example illustrates returning HTML as a preview reply with an image as an attachment:
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReplyAttachment
type PreviewReplyAttachment struct {
	objectivec.Object
}

// PreviewReplyAttachmentFrom constructs a [PreviewReplyAttachment] from an unsafe.Pointer.
//
// An attachment for a Quick Look preview reply that provides additional content for the system to display a preview.
func PreviewReplyAttachmentFrom(ptr unsafe.Pointer) PreviewReplyAttachment {
	return PreviewReplyAttachment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PreviewReplyAttachmentClass) Alloc() PreviewReplyAttachment {
	rv := objc.Send[PreviewReplyAttachment](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PreviewReplyAttachmentClass) New() PreviewReplyAttachment {
	rv := objc.Send[PreviewReplyAttachment](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewReplyAttachment) Init() PreviewReplyAttachment {
	rv := objc.Send[PreviewReplyAttachment](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewReplyAttachment) Autorelease() PreviewReplyAttachment {
	rv := objc.Send[PreviewReplyAttachment](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewReplyAttachment creates a new PreviewReplyAttachment instance.
func NewPreviewReplyAttachment() PreviewReplyAttachment {
	return getPreviewReplyAttachmentClass().New()
}




// Creates a preview reply attachment with the specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReplyAttachment/init(data:contentType:)
func NewPreviewReplyAttachmentWithDataContentType(data unsafe.Pointer, contentType unsafe.Pointer) PreviewReplyAttachment {
	instance := getPreviewReplyAttachmentClass().Alloc()
	rv := objc.Send[PreviewReplyAttachment](instance.ID, objc.Sel("initWithData:contentType:"), data, contentType)
	rv.Autorelease()
	return rv
}


// The content type of the preview attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReplyAttachment/contentType
func (p_ PreviewReplyAttachment) ContentType() UTType {
	rv := objc.Send[UTType](p_.ID, objc.Sel("contentType"))
	return rv
}

// The data of the preview attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReplyAttachment/data
func (p_ PreviewReplyAttachment) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("data"))
	return rv
}

// The attachments for a preview reply that provide additional data for the system to display the preview.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewreply/attachments
func (p_ PreviewReplyAttachment) Attachments() string {
	rv := objc.Send[string](p_.ID, objc.Sel("attachments"))
	return rv
}


// SetAttachments sets the value of the attachments property.
// The attachments for a preview reply that provide additional data for the system to display the preview.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewreply/attachments
func (p_ PreviewReplyAttachment) SetAttachments(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAttachments:"), objc.String(value))
}


