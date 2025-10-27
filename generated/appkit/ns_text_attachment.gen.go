// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	AllowsTextAttachmentView() bool
	SetAllowsTextAttachmentView(value bool)
	AttachmentCell() unsafe.Pointer
	SetAttachmentCell(value unsafe.Pointer)
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	Contents() foundation.foundation.INSData
	SetContents(value foundation.foundation.INSData)
	FileType() foundation.foundation.INSString
	SetFileType(value foundation.foundation.INSString)
	FileWrapper() foundation.FileWrapper
	SetFileWrapper(value foundation.FileWrapper)
	Image() IImage
	SetImage(value IImage)
	LineLayoutPadding() float64
	SetLineLayoutPadding(value float64)
	UsesTextAttachmentView() bool


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TextAttachmentClass) Alloc() TextAttachment {
	rv := objc.Send[TextAttachment](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates a text attachment object with the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/init(data:ofType:)
func NewTextAttachmentWithDataOfType(contentData foundation.foundation.INSData, uti foundation.foundation.INSString) TextAttachment {
	instance := getTextAttachmentClass().Alloc()
	rv := objc.Send[TextAttachment](instance.ID, objc.Sel("initWithData:ofType:"), contentData, uti)
	rv.Autorelease()
	return rv
}


// Creates a text attachment object to contain the specified file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/init(fileWrapper:)
func NewTextAttachmentWithFileWrapper(fileWrapper foundation.FileWrapper) TextAttachment {
	instance := getTextAttachmentClass().Alloc()
	rv := objc.Send[TextAttachment](instance.ID, objc.Sel("initWithFileWrapper:"), fileWrapper)
	rv.Autorelease()
	return rv
}







// Registers a specific file type with the attachment view provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/registerViewProviderClass(_:forFileType:)
func (tc _TextAttachmentClass) RegisterTextAttachmentViewProviderClassForFileType(textAttachmentViewProviderClass objc.Class, fileType foundation.foundation.INSString) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("registerTextAttachmentViewProviderClass:forFileType:"), textAttachmentViewProviderClass, fileType)
}


// Returns the text attachment view provider class, if any, for the file type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/textAttachmentViewProviderClass(forFileType:)
func (tc _TextAttachmentClass) TextAttachmentViewProviderClassForFileType(fileType foundation.foundation.INSString) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(tc.class), objc.Sel("textAttachmentViewProviderClassForFileType:"), fileType)
	return rv
}

















// A Boolean value that determines whether the text attachment uses text attachment views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/allowsTextAttachmentView
func (t_ TextAttachment) AllowsTextAttachmentView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsTextAttachmentView"))
	return rv
}


// A Boolean value that determines whether the text attachment uses text attachment views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/allowsTextAttachmentView
func (t_ TextAttachment) SetAllowsTextAttachmentView(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsTextAttachmentView:"), value)
}


// The object that draws the icon for the text attachment and handles mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/attachmentCell
func (t_ TextAttachment) AttachmentCell() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("attachmentCell"))
	return rv
}


// The object that draws the icon for the text attachment and handles mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/attachmentCell
func (t_ TextAttachment) SetAttachmentCell(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttachmentCell:"), value)
}


// The layout bounds of the text attachment’s graphical representation in the text coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/bounds
func (t_ TextAttachment) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("bounds"))
	return rv
}


// The layout bounds of the text attachment’s graphical representation in the text coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/bounds
func (t_ TextAttachment) SetBounds(value corefoundation.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBounds:"), value)
}


// The contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/contents
func (t_ TextAttachment) Contents() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("contents"))
	return rv
}


// The contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/contents
func (t_ TextAttachment) SetContents(value foundation.foundation.INSData) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContents:"), value)
}


// The file type of the contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileType
func (t_ TextAttachment) FileType() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("fileType"))
	return rv
}


// The file type of the contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileType
func (t_ TextAttachment) SetFileType(value foundation.foundation.INSString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFileType:"), value)
}


// The text attachment’s file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileWrapper
func (t_ TextAttachment) FileWrapper() foundation.FileWrapper {
	rv := objc.Send[foundation.FileWrapper](t_.ID, objc.Sel("fileWrapper"))
	return rv
}


// The text attachment’s file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileWrapper
func (t_ TextAttachment) SetFileWrapper(value foundation.FileWrapper) {
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
func (t_ TextAttachment) LineLayoutPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineLayoutPadding"))
	return rv
}


// The layout padding before and after the text attachment bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/lineLayoutPadding
func (t_ TextAttachment) SetLineLayoutPadding(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineLayoutPadding:"), value)
}


// A Boolean value that indicates whether the text attachment uses text attachment views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/usesTextAttachmentView
func (t_ TextAttachment) UsesTextAttachmentView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesTextAttachmentView"))
	return rv
}







