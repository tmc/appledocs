// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextAttachment] class.
var (
	TextAttachmentClass     _TextAttachmentClass
	TextAttachmentClassOnce sync.Once
)

func getTextAttachmentClass() _TextAttachmentClass {
	TextAttachmentClassOnce.Do(func() {
		TextAttachmentClass = _TextAttachmentClass{objc.GetClass("NSTextAttachment")}
	})
	return TextAttachmentClass
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


// A Boolean value that determines whether the text attachment uses text attachment views.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/allowstextattachmentview
func (t_ TextAttachment) AllowsTextAttachmentView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsTextAttachmentView"))
	return rv
}


// SetAllowsTextAttachmentView sets the value of the allowsTextAttachmentView property.
// A Boolean value that determines whether the text attachment uses text attachment views.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/allowstextattachmentview
func (t_ TextAttachment) SetAllowsTextAttachmentView(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsTextAttachmentView:"), value)
}

// The object that draws the icon for the text attachment and handles mouse events.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/attachmentcell
func (t_ TextAttachment) AttachmentCell() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("attachmentCell"))
	return rv
}


// SetAttachmentCell sets the value of the attachmentCell property.
// The object that draws the icon for the text attachment and handles mouse events.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/attachmentcell
func (t_ TextAttachment) SetAttachmentCell(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttachmentCell:"), value)
}

// The layout bounds of the text attachment’s graphical representation in the text coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/bounds
func (t_ TextAttachment) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("bounds"))
	return rv
}


// SetBounds sets the value of the bounds property.
// The layout bounds of the text attachment’s graphical representation in the text coordinate system.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/bounds
func (t_ TextAttachment) SetBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBounds:"), value)
}

// The contents for the text attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/contents
func (t_ TextAttachment) Contents() foundation.Data {
	rv := objc.Send[foundation.Data](t_.ID, objc.Sel("contents"))
	return rv
}


// SetContents sets the value of the contents property.
// The contents for the text attachment.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/contents
func (t_ TextAttachment) SetContents(value foundation.IData) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContents:"), value)
}

// The file type of the contents for the text attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/filetype
func (t_ TextAttachment) FileType() string {
	rv := objc.Send[string](t_.ID, objc.Sel("fileType"))
	return rv
}


// SetFileType sets the value of the fileType property.
// The file type of the contents for the text attachment.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/filetype
func (t_ TextAttachment) SetFileType(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFileType:"), objc.String(value))
}

// The text attachment’s file wrapper.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/filewrapper
func (t_ TextAttachment) FileWrapper() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("fileWrapper"))
	return rv
}


// SetFileWrapper sets the value of the fileWrapper property.
// The text attachment’s file wrapper.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/filewrapper
func (t_ TextAttachment) SetFileWrapper(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFileWrapper:"), value)
}

// An instance of the relevant image class that represents the contents of the text attachment object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/image
func (t_ TextAttachment) Image() Image {
	rv := objc.Send[Image](t_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// An instance of the relevant image class that represents the contents of the text attachment object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/image
func (t_ TextAttachment) SetImage(value IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImage:"), value)
}

// The layout padding before and after the text attachment bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/linelayoutpadding
func (t_ TextAttachment) LineLayoutPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineLayoutPadding"))
	return rv
}


// SetLineLayoutPadding sets the value of the lineLayoutPadding property.
// The layout padding before and after the text attachment bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/linelayoutpadding
func (t_ TextAttachment) SetLineLayoutPadding(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineLayoutPadding:"), value)
}

// A Boolean value that indicates whether the text attachment uses text attachment views.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/usestextattachmentview
func (t_ TextAttachment) UsesTextAttachmentView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesTextAttachmentView"))
	return rv
}


// SetUsesTextAttachmentView sets the value of the usesTextAttachmentView property.
// A Boolean value that indicates whether the text attachment uses text attachment views.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/usestextattachmentview
func (t_ TextAttachment) SetUsesTextAttachmentView(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesTextAttachmentView:"), value)
}



