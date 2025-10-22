// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ThumbnailReply] class.
type IThumbnailReply interface {
	objectivec.IObject
	ExtensionBadge() string
	SetExtensionBadge(value string)
}

// The object that provides a thumbnail for a custom file type.
//
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

// Alloc allocates a new instance without initialization.
func (tc _ThumbnailReplyClass) Alloc() ThumbnailReply {
	rv := objc.Send[ThumbnailReply](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a new thumbnail for a custom file type in the current context.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/init(contextSize:currentContextDrawing:)
func NewThumbnailReplyWithContextSizeCurrentContextDrawingBlock(contextSize coregraphics.CGSize, drawingBlock unsafe.Pointer) ThumbnailReply {
	rv := objc.Send[ThumbnailReply](objc.ID(getThumbnailReplyClass().class), objc.Sel("replyWithContextSize:currentContextDrawingBlock:"), contextSize, drawingBlock)
	return rv
}



// Creates a new thumbnail for a custom file type in the given context.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/init(contextSize:drawing:)
func NewThumbnailReplyWithContextSizeDrawingBlock(contextSize coregraphics.CGSize, drawingBlock unsafe.Pointer) ThumbnailReply {
	rv := objc.Send[ThumbnailReply](objc.ID(getThumbnailReplyClass().class), objc.Sel("replyWithContextSize:drawingBlock:"), contextSize, drawingBlock)
	return rv
}


// Creates a new thumbnail for a custom file type in the current context.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/init(contextSize:currentContextDrawing:)
func (tc _ThumbnailReplyClass) ReplyWithContextSizeCurrentContextDrawingBlock(contextSize coregraphics.CGSize, drawingBlock unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("replyWithContextSize:currentContextDrawingBlock:"), contextSize, drawingBlock)
	return rv
}

// Creates a new thumbnail for a custom file type in the given context.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/init(contextSize:drawing:)
func (tc _ThumbnailReplyClass) ReplyWithContextSizeDrawingBlock(contextSize coregraphics.CGSize, drawingBlock unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("replyWithContextSize:drawingBlock:"), contextSize, drawingBlock)
	return rv
}

// A short string that identifies the file type that the system uses as a badge when producing an icon thumbnail.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/extensionBadge
func (t_ ThumbnailReply) ExtensionBadge() string {
	rv := objc.Send[string](t_.ID, objc.Sel("extensionBadge"))
	return rv
}


// SetExtensionBadge sets the value of the extensionBadge property.
// A short string that identifies the file type that the system uses as a badge when producing an icon thumbnail.

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailReply/extensionBadge
func (t_ ThumbnailReply) SetExtensionBadge(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExtensionBadge:"), objc.String(value))
}


