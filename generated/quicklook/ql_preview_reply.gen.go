// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

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


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply
type PreviewReply struct {
	objectivec.Object
}

// PreviewReplyFrom constructs a [PreviewReply] from an unsafe.Pointer.
func PreviewReplyFrom(ptr unsafe.Pointer) PreviewReply {
	return PreviewReply{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreviewReply */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/initForPDFWithPageSize:documentCreationBlock:
func NewPreviewReplyForPDFWithPageSizeDocumentCreationBlock(defaultPageSize corefoundation.CGSize, documentCreationBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initForPDFWithPageSize:documentCreationBlock:"), defaultPageSize, documentCreationBlock)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewReplyForPDFWithPageSizeDocumentCreationBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/initWithContextSize:isBitmap:drawingBlock:
func NewPreviewReplyWithContextSizeIsBitmapDrawingBlock(contextSize corefoundation.CGSize, isBitmap bool, drawingBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithContextSize:isBitmap:drawingBlock:"), contextSize, isBitmap, drawingBlock)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewReplyWithContextSizeIsBitmapDrawingBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/initWithDataOfContentType:contentSize:dataCreationBlock:
func NewPreviewReplyWithDataOfContentTypeContentSizeDataCreationBlock(contentType uniformtypeidentifiers.UTType, contentSize corefoundation.CGSize, dataCreationBlock unsafe.Pointer) PreviewReply {
	instance := getPreviewReplyClass().Alloc()
	rv := objc.Send[PreviewReply](instance.ID, objc.Sel("initWithDataOfContentType:contentSize:dataCreationBlock:"), contentType, contentSize, dataCreationBlock)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewReplyWithDataOfContentTypeContentSizeDataCreationBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/init(fileURL:)
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
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLPreviewReply */


