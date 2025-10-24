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

/* debug [class.gen.go]: Generating class QLPreviewReply */


/* debug [class_header]: Header for QLPreviewReply */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PreviewReply */
// An interface definition for the [PreviewReply] class.
type IPreviewReply interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PreviewReply */
	// properties:
	Attachments() foundation.IDictionary
	SetAttachments(value foundation.IDictionary)
	StringEncoding() StringEncoding /* not a class type */
	SetStringEncoding(value StringEncoding /* not a class type */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PreviewReply */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PreviewReply */
// Alloc allocates a new instance without initialization.
func (pc _PreviewReplyClass) Alloc() PreviewReply {
	rv := objc.Send[PreviewReply](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PreviewReply */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreviewReply */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/initForPDFWithPageSize:documentCreationBlock:
func NewPreviewReplyForPDFWithPageSizeDocumentCreationBlock(defaultPageSize corefoundation.CGSize, documentCreationBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initForPDFWithPageSize:documentCreationBlock:"), defaultPageSize, documentCreationBlock)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewReplyForPDFWithPageSizeDocumentCreationBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/initWithContextSize:isBitmap:drawingBlock:
func NewPreviewReplyWithContextSizeIsBitmapDrawingBlock(contextSize corefoundation.CGSize, isBitmap bool, drawingBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithContextSize:isBitmap:drawingBlock:"), contextSize, isBitmap, drawingBlock)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewReplyWithContextSizeIsBitmapDrawingBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/initWithDataOfContentType:contentSize:dataCreationBlock:
func NewPreviewReplyWithDataOfContentTypeContentSizeDataCreationBlock(contentType uniformtypeidentifiers.UTType, contentSize corefoundation.CGSize, dataCreationBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithDataOfContentType:contentSize:dataCreationBlock:"), contentType, contentSize, dataCreationBlock)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewReplyWithDataOfContentTypeContentSizeDataCreationBlock */


// Creates a preview reply from an existing file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/init(fileURL:)
func NewPreviewReplyWithFileURL(fileURL objc.IObject /* cross-framework: NSURL */) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithFileURL:"), fileURL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewReplyWithFileURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PreviewReply */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PreviewReply */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PreviewReply */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PreviewReply */

// The attachments for a preview reply that provide additional data for the system to display the preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/attachments
func (p_ PreviewReply) Attachments() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("attachments"))
	return rv
}/* debug [instance_properties/getter]: attachments */


// The attachments for a preview reply that provide additional data for the system to display the preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/attachments
func (p_ PreviewReply) SetAttachments(value foundation.IDictionary) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAttachments:"), value)
}/* debug [instance_properties/setter]: attachments */


// String encoding for text or html based previews. Defaults to NSUTF8StringEncoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/stringEncoding-1k9kb
func (p_ PreviewReply) StringEncoding() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](p_.ID, objc.Sel("stringEncoding"))
	return rv
}/* debug [instance_properties/getter]: stringEncoding */


// String encoding for text or html based previews. Defaults to NSUTF8StringEncoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/stringEncoding-1k9kb
func (p_ PreviewReply) SetStringEncoding(value StringEncoding /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStringEncoding:"), value)
}/* debug [instance_properties/setter]: stringEncoding */


// The title for the system to display with the preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/title
func (p_ PreviewReply) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title for the system to display with the preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewReply/title
func (p_ PreviewReply) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLPreviewReply */


