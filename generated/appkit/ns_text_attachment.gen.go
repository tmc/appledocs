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
	// properties:
	AllowsTextAttachmentView() bool /* primitive/slice/pointer. */
	SetAllowsTextAttachmentView(value bool /* primitive/slice/pointer. */)
	AttachmentCell() objc.ID
	SetAttachmentCell(value objc.ID)
	Bounds() coregraphics.CGRect
	SetBounds(value coregraphics.CGRect)
	Contents() foundation.objc.IObject /* cross-framework: NSData */
	SetContents(value foundation.objc.IObject /* cross-framework: NSData */)
	FileType() string /* primitive/slice/pointer. */
	SetFileType(value string /* primitive/slice/pointer. */)
	FileWrapper() FileWrapper /* not a class type */
	SetFileWrapper(value FileWrapper /* not a class type */)
	Image() IImage
	SetImage(value IImage)
	LineLayoutPadding() float64 /* primitive/slice/pointer. */
	SetLineLayoutPadding(value float64 /* primitive/slice/pointer. */)
	UsesTextAttachmentView() bool /* primitive/slice/pointer. */
	// methods:
}

// The values for the attachment characteristics of attributed strings and related objects.
//
// The class uses text attachment objects as the values for attachment attributes (stored in the attributed string under the key). A text attachment object contains either an object or an object, which in turn holds the contents of the attached file. The properties of this class configure the appearance of the text attachment in your interface. In macOS, the text attachment also uses a cell object that conforms to the protocol to draw the image that represents the text and handles mouse events. For more information about text attachments, see the and . In macOS 12 and iOS 15 and later, and provide additional capabilities to represent document locations in terms of an  or an  , and provide support for view-based text attachments.


// The values for the attachment characteristics of attributed strings and related objects.
//
// [Full Topic]
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



// Creates a text attachment object with the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/init(data:ofType:)
func NewTextAttachmentWithDataOfType(contentData foundation.objc.IObject /* cross-framework NSData */, uti string /* primitive/slice/pointer. */) TextAttachment {
	instance := getTextAttachmentClass().Alloc()
	rv := objc.Send[TextAttachment](instance.ID, objc.Sel("initWithData:ofType:"), contentData, objc.String(uti))
	rv.Autorelease()
	return rv
}


// Creates a text attachment object to contain the specified file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/init(fileWrapper:)
func NewTextAttachmentWithFileWrapper(fileWrapper FileWrapper /* not a class type */) TextAttachment {
	instance := getTextAttachmentClass().Alloc()
	rv := objc.Send[TextAttachment](instance.ID, objc.Sel("initWithFileWrapper:"), fileWrapper)
	rv.Autorelease()
	return rv
}



// Registers a specific file type with the attachment view provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/registerViewProviderClass(_:forFileType:)
func (tc _TextAttachmentClass) RegisterTextAttachmentViewProviderClassForFileType(textAttachmentViewProviderClass objc.Class, fileType string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("registerTextAttachmentViewProviderClass:forFileType:"), textAttachmentViewProviderClass, objc.String(fileType))
}


// Returns the text attachment view provider class, if any, for the file type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/textAttachmentViewProviderClass(forFileType:)
func (tc _TextAttachmentClass) TextAttachmentViewProviderClassForFileType(fileType string /* primitive/slice/pointer. */) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(tc.class), objc.Sel("textAttachmentViewProviderClassForFileType:"), objc.String(fileType))
	return rv
}


// A Boolean value that determines whether the text attachment uses text attachment views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/allowsTextAttachmentView
func (t_ TextAttachment) AllowsTextAttachmentView() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsTextAttachmentView"))
	return rv
}


// A Boolean value that determines whether the text attachment uses text attachment views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/allowsTextAttachmentView
func (t_ TextAttachment) SetAllowsTextAttachmentView(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsTextAttachmentView:"), value)
}


// The object that draws the icon for the text attachment and handles mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/attachmentCell
func (t_ TextAttachment) AttachmentCell() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("attachmentCell"))
	return rv
}


// The object that draws the icon for the text attachment and handles mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/attachmentCell
func (t_ TextAttachment) SetAttachmentCell(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttachmentCell:"), value)
}


// The layout bounds of the text attachment’s graphical representation in the text coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/bounds
func (t_ TextAttachment) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("bounds"))
	return rv
}


// The layout bounds of the text attachment’s graphical representation in the text coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/bounds
func (t_ TextAttachment) SetBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBounds:"), value)
}


// The contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/contents
func (t_ TextAttachment) Contents() foundation.objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("contents"))
	return rv
}


// The contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/contents
func (t_ TextAttachment) SetContents(value foundation.objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContents:"), value)
}


// The file type of the contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileType
func (t_ TextAttachment) FileType() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](t_.ID, objc.Sel("fileType"))
	return rv
}


// The file type of the contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileType
func (t_ TextAttachment) SetFileType(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFileType:"), objc.String(value))
}


// The text attachment’s file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileWrapper
func (t_ TextAttachment) FileWrapper() FileWrapper /* not a class type */ {
	rv := objc.Send[FileWrapper](t_.ID, objc.Sel("fileWrapper"))
	return rv
}


// The text attachment’s file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileWrapper
func (t_ TextAttachment) SetFileWrapper(value FileWrapper /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFileWrapper:"), value)
}


// An instance of the relevant image class that represents the contents of the text attachment object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/image
func (t_ TextAttachment) Image() IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("image"))
	return rv
}


// An instance of the relevant image class that represents the contents of the text attachment object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/image
func (t_ TextAttachment) SetImage(value IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImage:"), value)
}


// The layout padding before and after the text attachment bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/lineLayoutPadding
func (t_ TextAttachment) LineLayoutPadding() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineLayoutPadding"))
	return rv
}


// The layout padding before and after the text attachment bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/lineLayoutPadding
func (t_ TextAttachment) SetLineLayoutPadding(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineLayoutPadding:"), value)
}


// A Boolean value that indicates whether the text attachment uses text attachment views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/usesTextAttachmentView
func (t_ TextAttachment) UsesTextAttachmentView() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesTextAttachmentView"))
	return rv
}


