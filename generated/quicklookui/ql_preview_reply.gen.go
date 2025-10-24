// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	StringEncoding() StringEncoding /* not a class type */
	SetStringEncoding(value StringEncoding /* not a class type */)
	Attachments() IQLPreviewReplyAttachment
	SetAttachments(value IQLPreviewReplyAttachment)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// The class you create when providing a data-based Quick Look preview extension.
//
// Create an instance of from the method in your subclass of . Create an instance to return data; for example, an image, PDF, or HTML; that the system displays as the preview for the content that the system indicates with .


// The class you create when providing a data-based Quick Look preview extension.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/initForPDFWithPageSize:documentCreationBlock:
func NewPreviewReplyForPDFWithPageSizeDocumentCreationBlock(defaultPageSize objc.IObject /* cross-framework: Size */, documentCreationBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initForPDFWithPageSize:documentCreationBlock:"), defaultPageSize, documentCreationBlock)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/initWithContextSize:isBitmap:drawingBlock:
func NewPreviewReplyWithContextSizeIsBitmapDrawingBlock(contextSize objc.IObject /* cross-framework: Size */, isBitmap bool, drawingBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithContextSize:isBitmap:drawingBlock:"), contextSize, isBitmap, drawingBlock)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/initWithDataOfContentType:contentSize:dataCreationBlock:
func NewPreviewReplyWithDataOfContentTypeContentSizeDataCreationBlock(contentType objc.IObject /* cross-framework: UTType */, contentSize objc.IObject /* cross-framework: Size */, dataCreationBlock Data  * (^)( QLPreviewReply  *  reply ,  NSError  * *  error /* not a class type */) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithDataOfContentType:contentSize:dataCreationBlock:"), contentType, contentSize, dataCreationBlock)
	rv.Autorelease()
	return rv
}


// Creates a preview reply from an existing file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/init(fileURL:)
func NewPreviewReplyWithFileURL(fileURL objc.IObject /* cross-framework: NSURL */) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithFileURL:"), fileURL)
	rv.Autorelease()
	return rv
}



// String encoding for text or html based previews. Defaults to NSUTF8StringEncoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/stringEncoding-1k9kb
func (p_ PreviewReply) StringEncoding() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](p_.ID, objc.Sel("stringEncoding"))
	return rv
}


// String encoding for text or html based previews. Defaults to NSUTF8StringEncoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/stringEncoding-1k9kb
func (p_ PreviewReply) SetStringEncoding(value StringEncoding /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStringEncoding:"), value)
}


// The attachments for a preview reply that provide additional data for the system to display the preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewreply/attachments
func (p_ PreviewReply) Attachments() IQLPreviewReplyAttachment {
	rv := objc.Send[PreviewReplyAttachment](p_.ID, objc.Sel("attachments"))
	return rv
}


// The attachments for a preview reply that provide additional data for the system to display the preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewreply/attachments
func (p_ PreviewReply) SetAttachments(value IQLPreviewReplyAttachment) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAttachments:"), value)
}


// The title for the system to display with the preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewreply/title
func (p_ PreviewReply) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("title"))
	return rv
}


// The title for the system to display with the preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewreply/title
func (p_ PreviewReply) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), value)
}


