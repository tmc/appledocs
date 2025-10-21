// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [PreviewReply] class.
var (
	PreviewReplyClass     _PreviewReplyClass
	PreviewReplyClassOnce sync.Once
)

func getPreviewReplyClass() _PreviewReplyClass {
	PreviewReplyClassOnce.Do(func() {
		PreviewReplyClass = _PreviewReplyClass{objc.GetClass("QLPreviewReply")}
	})
	return PreviewReplyClass
}

type _PreviewReplyClass struct {
	class objc.Class
}

// An interface definition for the [PreviewReply] class.
type IPreviewReply interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply
type PreviewReply struct {
	objectivec.Object
}

// PreviewReplyFrom constructs a [PreviewReply] from an unsafe.Pointer.
func PreviewReplyFrom(ptr unsafe.Pointer) PreviewReply {
	return PreviewReply{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PreviewReplyClass) Alloc() PreviewReply {
	rv := objc.Send[PreviewReply](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PreviewReplyClass) New() PreviewReply {
	rv := objc.Send[PreviewReply](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewReply) Init() PreviewReply {
	rv := objc.Send[PreviewReply](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewReply) Autorelease() PreviewReply {
	rv := objc.Send[PreviewReply](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewReply creates a new PreviewReply instance.
func NewPreviewReply() PreviewReply {
	return getPreviewReplyClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/initWithDataOfContentType:contentSize:dataCreationBlock:
func NewPreviewReplyWithDataOfContentTypeContentSizeDataCreationBlock(contentType unsafe.Pointer, contentSize coregraphics.CGSize, dataCreationBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithDataOfContentType:contentSize:dataCreationBlock:"), contentType, contentSize, dataCreationBlock)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/init(fileURL:)
func NewPreviewReplyWithFileURL(fileURL unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithFileURL:"), fileURL)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/initForPDFWithPageSize:documentCreationBlock:
func NewPreviewReplyForPDFWithPageSizeDocumentCreationBlock(defaultPageSize coregraphics.CGSize, documentCreationBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initForPDFWithPageSize:documentCreationBlock:"), defaultPageSize, documentCreationBlock)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/initWithContextSize:isBitmap:drawingBlock:
func NewPreviewReplyWithContextSizeIsBitmapDrawingBlock(contextSize coregraphics.CGSize, isBitmap bool, drawingBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithContextSize:isBitmap:drawingBlock:"), contextSize, isBitmap, drawingBlock)
	rv.Autorelease()
	return rv
}


// Attachments for HTML data previews. The keys of the dictionary are the attachment identifiers (eg foo) that can be referenced with the cid:id URL (eg cid:foo).
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/attachments
func (p_ PreviewReply) Attachments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("attachments"))
	return rv
}


// SetAttachments sets the value of the attachments property.
// Attachments for HTML data previews. The keys of the dictionary are the attachment identifiers (eg foo) that can be referenced with the cid:id URL (eg cid:foo).

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/attachments
func (p_ PreviewReply) SetAttachments(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAttachments:"), value)
}
// String encoding for text or html based previews. Defaults to NSUTF8StringEncoding.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/stringEncoding-1k9kb
func (p_ PreviewReply) StringEncoding() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("stringEncoding"))
	return rv
}


// SetStringEncoding sets the value of the stringEncoding property.
// String encoding for text or html based previews. Defaults to NSUTF8StringEncoding.

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/stringEncoding-1k9kb
func (p_ PreviewReply) SetStringEncoding(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStringEncoding:"), value)
}
// Custom display title for the preview. If left as the empty string, QuickLook will use the file name.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/title
func (p_ PreviewReply) Title() string {
	rv := objc.Send[string](p_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// Custom display title for the preview. If left as the empty string, QuickLook will use the file name.

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/title
func (p_ PreviewReply) SetTitle(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), objc.String(value))
}

