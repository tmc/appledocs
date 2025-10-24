// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextPreview */


/* debug [class_header]: Header for NSTextPreview */
// The class instance for the [TextPreview] class.
var (
	TextPreviewClass     _TextPreviewClass
	TextPreviewClassOnce sync.Once
)

func getTextPreviewClass() _TextPreviewClass {
	TextPreviewClassOnce.Do(func() {
		TextPreviewClass = _TextPreviewClass{objc.GetClass("NSTextPreview")}
	})
	return TextPreviewClass
}

type _TextPreviewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextPreview */
// An interface definition for the [TextPreview] class.
type ITextPreview interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextPreview */
	// properties:
	CandidateRects() []foundation.Value
	PresentationFrame() Rect /* not a class type */
	PreviewImage() ImageRef /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextPreview */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextPreview */
// Alloc allocates a new instance without initialization.
func (tc _TextPreviewClass) Alloc() TextPreview {
	rv := objc.Send[TextPreview](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextPreviewClass) New() TextPreview {
	rv := objc.Send[TextPreview](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextPreview) Init() TextPreview {
	rv := objc.Send[TextPreview](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextPreview) Autorelease() TextPreview {
	rv := objc.Send[TextPreview](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextPreview creates a new TextPreview instance.
func NewTextPreview() TextPreview {
	return getTextPreviewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextPreview */
// A snapshot of the text in your view, which the system uses to create user-visible effects.
//
// An object provides a static image of your view’s text content that the system can use to create animations. You provide preview objects in response to system requests, such as ones from Writing Tools. In addition to creating an image of your view’s text, you also specify the location of that text in your view’s frame rectangle. When creating animations, the system places the image on top of your view’s content and animates changes to the image instead of to your view. Create an object in response to specific system requests. Create an image with a transparent background and render your view’s text into the image using the current text attributes. Construct your object with both the image and the frame rectangle that represents the location of the rendered text in your view’s coordinate system. To highlight specific portions of text, instead of all the text in the image, provide a set of candidate rectangles with the locations of the text you want to highlight.


// A snapshot of the text in your view, which the system uses to create user-visible effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextPreview
type TextPreview struct {
	objectivec.Object
}

// TextPreviewFrom constructs a [TextPreview] from an unsafe.Pointer.
//
// A snapshot of the text in your view, which the system uses to create user-visible effects.
func TextPreviewFrom(ptr unsafe.Pointer) TextPreview {
	return TextPreview{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextPreview */

// Creates a text preview using the specified image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextPreview/init(snapshotImage:presentationFrame:)
func NewTextPreviewWithSnapshotImagePresentationFrame(snapshotImage ImageRef /* not a class type */, presentationFrame Rect /* not a class type */) TextPreview {
	instance := getTextPreviewClass().Alloc()
	rv := objc.Send[TextPreview](instance.ID, objc.Sel("initWithSnapshotImage:presentationFrame:"), snapshotImage, presentationFrame)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextPreviewWithSnapshotImagePresentationFrame */


// Creates a text preview using the specified image and rectangles that indicate the portions of text to highlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextPreview/init(snapshotImage:presentationFrame:candidateRects:)
func NewTextPreviewWithSnapshotImagePresentationFrameCandidateRects(snapshotImage ImageRef /* not a class type */, presentationFrame Rect /* not a class type */, candidateRects []foundation.Value) TextPreview {
	instance := getTextPreviewClass().Alloc()
	rv := objc.Send[TextPreview](instance.ID, objc.Sel("initWithSnapshotImage:presentationFrame:candidateRects:"), snapshotImage, presentationFrame, candidateRects)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextPreviewWithSnapshotImagePresentationFrameCandidateRects */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextPreview */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextPreview */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextPreview */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextPreview */

// Rectangles that define the specific portions of text to highlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextPreview/candidateRects
func (t_ TextPreview) CandidateRects() []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("candidateRects"))
	return rv
}/* debug [instance_properties/getter]: candidateRects */


// The frame rectangle that places the preview image directly over the matching text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextPreview/presentationFrame
func (t_ TextPreview) PresentationFrame() Rect /* not a class type */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("presentationFrame"))
	return rv
}/* debug [instance_properties/getter]: presentationFrame */


// The image that contains the requested text from your view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextPreview/previewImage
func (t_ TextPreview) PreviewImage() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](t_.ID, objc.Sel("previewImage"))
	return rv
}/* debug [instance_properties/getter]: previewImage */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextPreview */


