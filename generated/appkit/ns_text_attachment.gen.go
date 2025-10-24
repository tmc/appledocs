// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextAttachment */


/* debug [class_header]: Header for NSTextAttachment */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextAttachment */
// An interface definition for the [TextAttachment] class.
type ITextAttachment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextAttachment */
	// properties:
	AllowsTextAttachmentView() bool
	SetAllowsTextAttachmentView(value bool)
	AttachmentCell() unsafe.Pointer
	SetAttachmentCell(value unsafe.Pointer)
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	Contents() objc.IObject /* cross-framework: NSData */
	SetContents(value objc.IObject /* cross-framework: NSData */)
	FileType() objc.IObject /* cross-framework: NSString */
	SetFileType(value objc.IObject /* cross-framework: NSString */)
	FileWrapper() foundation.FileWrapper
	SetFileWrapper(value foundation.FileWrapper)
	Image() IImage
	SetImage(value IImage)
	LineLayoutPadding() float64
	SetLineLayoutPadding(value float64)
	UsesTextAttachmentView() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextAttachment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextAttachment */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextAttachment */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextAttachment */

// Creates a text attachment object with the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/init(data:ofType:)
func NewTextAttachmentWithDataOfType(contentData objc.IObject /* cross-framework: NSData */, uti objc.IObject /* cross-framework: NSString */) TextAttachment {
	instance := getTextAttachmentClass().Alloc()
	rv := objc.Send[TextAttachment](instance.ID, objc.Sel("initWithData:ofType:"), contentData, uti)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextAttachmentWithDataOfType */


// Creates a text attachment object to contain the specified file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/init(fileWrapper:)
func NewTextAttachmentWithFileWrapper(fileWrapper foundation.FileWrapper) TextAttachment {
	instance := getTextAttachmentClass().Alloc()
	rv := objc.Send[TextAttachment](instance.ID, objc.Sel("initWithFileWrapper:"), fileWrapper)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextAttachmentWithFileWrapper */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextAttachment */

// Registers a specific file type with the attachment view provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/registerViewProviderClass(_:forFileType:)
func (tc _TextAttachmentClass) RegisterTextAttachmentViewProviderClassForFileType(textAttachmentViewProviderClass objc.Class, fileType objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("registerTextAttachmentViewProviderClass:forFileType:"), textAttachmentViewProviderClass, fileType)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterTextAttachmentViewProviderClassForFileType) */


// Returns the text attachment view provider class, if any, for the file type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/textAttachmentViewProviderClass(forFileType:)
func (tc _TextAttachmentClass) TextAttachmentViewProviderClassForFileType(fileType objc.IObject /* cross-framework: NSString */) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(tc.class), objc.Sel("textAttachmentViewProviderClassForFileType:"), fileType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TextAttachmentViewProviderClassForFileType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextAttachment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextAttachment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextAttachment */

// A Boolean value that determines whether the text attachment uses text attachment views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/allowsTextAttachmentView
func (t_ TextAttachment) AllowsTextAttachmentView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsTextAttachmentView"))
	return rv
}/* debug [instance_properties/getter]: allowsTextAttachmentView */


// A Boolean value that determines whether the text attachment uses text attachment views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/allowsTextAttachmentView
func (t_ TextAttachment) SetAllowsTextAttachmentView(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsTextAttachmentView:"), value)
}/* debug [instance_properties/setter]: allowsTextAttachmentView */


// The object that draws the icon for the text attachment and handles mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/attachmentCell
func (t_ TextAttachment) AttachmentCell() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("attachmentCell"))
	return rv
}/* debug [instance_properties/getter]: attachmentCell */


// The object that draws the icon for the text attachment and handles mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/attachmentCell
func (t_ TextAttachment) SetAttachmentCell(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttachmentCell:"), value)
}/* debug [instance_properties/setter]: attachmentCell */


// The layout bounds of the text attachment’s graphical representation in the text coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/bounds
func (t_ TextAttachment) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// The layout bounds of the text attachment’s graphical representation in the text coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/bounds
func (t_ TextAttachment) SetBounds(value corefoundation.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBounds:"), value)
}/* debug [instance_properties/setter]: bounds */


// The contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/contents
func (t_ TextAttachment) Contents() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("contents"))
	return rv
}/* debug [instance_properties/getter]: contents */


// The contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/contents
func (t_ TextAttachment) SetContents(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContents:"), value)
}/* debug [instance_properties/setter]: contents */


// The file type of the contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileType
func (t_ TextAttachment) FileType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("fileType"))
	return rv
}/* debug [instance_properties/getter]: fileType */


// The file type of the contents for the text attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileType
func (t_ TextAttachment) SetFileType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFileType:"), value)
}/* debug [instance_properties/setter]: fileType */


// The text attachment’s file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileWrapper
func (t_ TextAttachment) FileWrapper() foundation.FileWrapper {
	rv := objc.Send[foundation.FileWrapper](t_.ID, objc.Sel("fileWrapper"))
	return rv
}/* debug [instance_properties/getter]: fileWrapper */


// The text attachment’s file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/fileWrapper
func (t_ TextAttachment) SetFileWrapper(value foundation.FileWrapper) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFileWrapper:"), value)
}/* debug [instance_properties/setter]: fileWrapper */


// An instance of the relevant image class that represents the contents of the text attachment object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/image
func (t_ TextAttachment) Image() IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// An instance of the relevant image class that represents the contents of the text attachment object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/image
func (t_ TextAttachment) SetImage(value IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImage:"), value)
}/* debug [instance_properties/setter]: image */


// The layout padding before and after the text attachment bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/lineLayoutPadding
func (t_ TextAttachment) LineLayoutPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineLayoutPadding"))
	return rv
}/* debug [instance_properties/getter]: lineLayoutPadding */


// The layout padding before and after the text attachment bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/lineLayoutPadding
func (t_ TextAttachment) SetLineLayoutPadding(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineLayoutPadding:"), value)
}/* debug [instance_properties/setter]: lineLayoutPadding */


// A Boolean value that indicates whether the text attachment uses text attachment views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment/usesTextAttachmentView
func (t_ TextAttachment) UsesTextAttachmentView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesTextAttachmentView"))
	return rv
}/* debug [instance_properties/getter]: usesTextAttachmentView */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextAttachment */


