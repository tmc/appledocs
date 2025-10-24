// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QLThumbnailReply */


/* debug [class_header]: Header for QLThumbnailReply */
// The class instance for the [ThumbnailReply] class.
var (
	ThumbnailReplyClass     _ThumbnailReplyClass
	ThumbnailReplyClassOnce sync.Once
)

func getThumbnailReplyClass() _ThumbnailReplyClass {
	ThumbnailReplyClassOnce.Do(func() {
		ThumbnailReplyClass = _ThumbnailReplyClass{objc.GetClass("QLThumbnailReply")}
	})
	return ThumbnailReplyClass
}

type _ThumbnailReplyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ThumbnailReply */
// An interface definition for the [ThumbnailReply] class.
type IThumbnailReply interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ThumbnailReply */
	// properties:
	ExtensionBadge() objc.IObject /* cross-framework: NSString */
	SetExtensionBadge(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ThumbnailReply */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ThumbnailReply */
// Alloc allocates a new instance without initialization.
func (tc _ThumbnailReplyClass) Alloc() ThumbnailReply {
	rv := objc.Send[ThumbnailReply](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _ThumbnailReplyClass) New() ThumbnailReply {
	rv := objc.Send[ThumbnailReply](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ThumbnailReply) Init() ThumbnailReply {
	rv := objc.Send[ThumbnailReply](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ThumbnailReply) Autorelease() ThumbnailReply {
	rv := objc.Send[ThumbnailReply](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewThumbnailReply creates a new ThumbnailReply instance.
func NewThumbnailReply() ThumbnailReply {
	return getThumbnailReplyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ThumbnailReply */
// The object that provides a thumbnail for a custom file type.


// The object that provides a thumbnail for a custom file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply
type ThumbnailReply struct {
	objectivec.Object
}

// ThumbnailReplyFrom constructs a [ThumbnailReply] from an unsafe.Pointer.
//
// The object that provides a thumbnail for a custom file type.
func ThumbnailReplyFrom(ptr unsafe.Pointer) ThumbnailReply {
	return ThumbnailReply{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ThumbnailReply */

// Creates a new thumbnail for a custom file type in the current context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/init(contextSize:currentContextDrawing:)
func NewThumbnailReplyWithContextSizeCurrentContextDrawingBlock(contextSize corefoundation.CGSize, drawingBlock unsafe.Pointer) ThumbnailReply {
	rv := objc.Send[ThumbnailReply](objc.ID(getThumbnailReplyClass().class), objc.Sel("replyWithContextSize:currentContextDrawingBlock:"), contextSize, drawingBlock)
	return rv
}/* debug [class_init_methods/constructor]: NewThumbnailReplyWithContextSizeCurrentContextDrawingBlock */


// Creates a new thumbnail for a custom file type in the given context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/init(contextSize:drawing:)
func NewThumbnailReplyWithContextSizeDrawingBlock(contextSize corefoundation.CGSize, drawingBlock unsafe.Pointer) ThumbnailReply {
	rv := objc.Send[ThumbnailReply](objc.ID(getThumbnailReplyClass().class), objc.Sel("replyWithContextSize:drawingBlock:"), contextSize, drawingBlock)
	return rv
}/* debug [class_init_methods/constructor]: NewThumbnailReplyWithContextSizeDrawingBlock */


// Creates a new thumbnail for a custom file type using a file at the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/init(imageFileURL:)
func NewThumbnailReplyWithImageFileURL(fileURL objc.IObject /* cross-framework: NSURL */) ThumbnailReply {
	rv := objc.Send[ThumbnailReply](objc.ID(getThumbnailReplyClass().class), objc.Sel("replyWithImageFileURL:"), fileURL)
	return rv
}/* debug [class_init_methods/constructor]: NewThumbnailReplyWithImageFileURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ThumbnailReply */

// Creates a new thumbnail for a custom file type in the current context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/init(contextSize:currentContextDrawing:)
func (tc _ThumbnailReplyClass) ReplyWithContextSizeCurrentContextDrawingBlock(contextSize corefoundation.CGSize, drawingBlock unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("replyWithContextSize:currentContextDrawingBlock:"), contextSize, drawingBlock)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReplyWithContextSizeCurrentContextDrawingBlock) */


// Creates a new thumbnail for a custom file type in the given context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/init(contextSize:drawing:)
func (tc _ThumbnailReplyClass) ReplyWithContextSizeDrawingBlock(contextSize corefoundation.CGSize, drawingBlock unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("replyWithContextSize:drawingBlock:"), contextSize, drawingBlock)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReplyWithContextSizeDrawingBlock) */


// Creates a new thumbnail for a custom file type using a file at the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/init(imageFileURL:)
func (tc _ThumbnailReplyClass) ReplyWithImageFileURL(fileURL objc.IObject /* cross-framework: NSURL */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("replyWithImageFileURL:"), fileURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReplyWithImageFileURL) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ThumbnailReply */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ThumbnailReply */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ThumbnailReply */

// A short string that identifies the file type that the system uses as a badge when producing an icon thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/extensionBadge
func (t_ ThumbnailReply) ExtensionBadge() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("extensionBadge"))
	return rv
}/* debug [instance_properties/getter]: extensionBadge */


// A short string that identifies the file type that the system uses as a badge when producing an icon thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/extensionBadge
func (t_ ThumbnailReply) SetExtensionBadge(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExtensionBadge:"), value)
}/* debug [instance_properties/setter]: extensionBadge */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLThumbnailReply */


