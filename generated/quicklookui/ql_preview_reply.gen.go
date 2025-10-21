// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
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

// The class you create when providing a data-based Quick Look preview extension.
//
// Create an instance of from the method in your subclass of . Create an instance to return data; for example, an image, PDF, or HTML; that the system displays as the preview for the content that the system indicates with .
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply
type PreviewReply struct {
	objectivec.Object
}

// PreviewReplyFrom constructs a [PreviewReply] from an unsafe.Pointer.
//
// The class you create when providing a data-based Quick Look preview extension.
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
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/initForPDFWithPageSize:documentCreationBlock:
func NewPreviewReplyForPDFWithPageSizeDocumentCreationBlock(defaultPageSize coregraphics.CGSize, documentCreationBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initForPDFWithPageSize:documentCreationBlock:"), defaultPageSize, documentCreationBlock)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/initWithContextSize:isBitmap:drawingBlock:
func NewPreviewReplyWithContextSizeIsBitmapDrawingBlock(contextSize coregraphics.CGSize, isBitmap bool, drawingBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithContextSize:isBitmap:drawingBlock:"), contextSize, isBitmap, drawingBlock)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/initWithDataOfContentType:contentSize:dataCreationBlock:
func NewPreviewReplyWithDataOfContentTypeContentSizeDataCreationBlock(contentType uniformtypeidentifiers.UTType, contentSize coregraphics.CGSize, dataCreationBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithDataOfContentType:contentSize:dataCreationBlock:"), contentType, contentSize, dataCreationBlock)
	rv.Autorelease()
	return rv
}

// Creates a preview reply from an existing file URL.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/init(fileURL:)
func NewPreviewReplyWithFileURL(fileURL unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithFileURL:"), fileURL)
	rv.Autorelease()
	return rv
}


// The attachments for a preview reply that provide additional data for the system to display the preview.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/attachments
func (p_ PreviewReply) Attachments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("attachments"))
	return rv
}


// SetAttachments sets the value of the attachments property.
// The attachments for a preview reply that provide additional data for the system to display the preview.

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/attachments
func (p_ PreviewReply) SetAttachments(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAttachments:"), value)
}
// String encoding for text or html based previews. Defaults to NSUTF8StringEncoding.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/stringEncoding-1k9kb
func (p_ PreviewReply) StringEncoding() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("stringEncoding"))
	return rv
}


// SetStringEncoding sets the value of the stringEncoding property.
// String encoding for text or html based previews. Defaults to NSUTF8StringEncoding.

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/stringEncoding-1k9kb
func (p_ PreviewReply) SetStringEncoding(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStringEncoding:"), value)
}

