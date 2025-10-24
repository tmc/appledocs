// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class QLThumbnailGenerationRequest */


/* debug [class_header]: Header for QLThumbnailGenerationRequest */
// The class instance for the [ThumbnailGenerationRequest] class.
var (
	ThumbnailGenerationRequestClass     _ThumbnailGenerationRequestClass
	ThumbnailGenerationRequestClassOnce sync.Once
)

func getThumbnailGenerationRequestClass() _ThumbnailGenerationRequestClass {
	ThumbnailGenerationRequestClassOnce.Do(func() {
		ThumbnailGenerationRequestClass = _ThumbnailGenerationRequestClass{objc.GetClass("QLThumbnailGenerationRequest")}
	})
	return ThumbnailGenerationRequestClass
}

type _ThumbnailGenerationRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ThumbnailGenerationRequest */
// An interface definition for the [ThumbnailGenerationRequest] class.
type IThumbnailGenerationRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ThumbnailGenerationRequest */
	// properties:
	ContentType() uniformtypeidentifiers.UTType
	SetContentType(value uniformtypeidentifiers.UTType)
	IconMode() bool
	SetIconMode(value bool)
	MinimumDimension() float64
	SetMinimumDimension(value float64)
	RepresentationTypes() ThumbnailGenerationRequestRepresentationTypes
	Scale() float64
	Size() corefoundation.CGSize
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ThumbnailGenerationRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ThumbnailGenerationRequest */
// Alloc allocates a new instance without initialization.
func (tc _ThumbnailGenerationRequestClass) Alloc() ThumbnailGenerationRequest {
	rv := objc.Send[ThumbnailGenerationRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _ThumbnailGenerationRequestClass) New() ThumbnailGenerationRequest {
	rv := objc.Send[ThumbnailGenerationRequest](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ThumbnailGenerationRequest) Init() ThumbnailGenerationRequest {
	rv := objc.Send[ThumbnailGenerationRequest](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ThumbnailGenerationRequest) Autorelease() ThumbnailGenerationRequest {
	rv := objc.Send[ThumbnailGenerationRequest](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewThumbnailGenerationRequest creates a new ThumbnailGenerationRequest instance.
func NewThumbnailGenerationRequest() ThumbnailGenerationRequest {
	return getThumbnailGenerationRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ThumbnailGenerationRequest */
// A request to generate a thumbnail for a file.


// A request to generate a thumbnail for a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request
type ThumbnailGenerationRequest struct {
	objectivec.Object
}

// ThumbnailGenerationRequestFrom constructs a [ThumbnailGenerationRequest] from an unsafe.Pointer.
//
// A request to generate a thumbnail for a file.
func ThumbnailGenerationRequestFrom(ptr unsafe.Pointer) ThumbnailGenerationRequest {
	return ThumbnailGenerationRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ThumbnailGenerationRequest */

// Creates a new request for a thumbnail with the specified parameters for a file at a provided URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/init(fileAt:size:scale:representationTypes:)
func NewThumbnailGenerationRequestWithFileAtURLSizeScaleRepresentationTypes(url objc.IObject /* cross-framework: NSURL */, size corefoundation.CGSize, scale float64, representationTypes ThumbnailGenerationRequestRepresentationTypes) ThumbnailGenerationRequest {
	instance := getThumbnailGenerationRequestClass().Alloc()
	rv := objc.Send[ThumbnailGenerationRequest](instance.ID, objc.Sel("initWithFileAtURL:size:scale:representationTypes:"), url, size, scale, representationTypes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewThumbnailGenerationRequestWithFileAtURLSizeScaleRepresentationTypes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ThumbnailGenerationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ThumbnailGenerationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ThumbnailGenerationRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ThumbnailGenerationRequest */

// The content type of the source data for the thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/contentType
func (t_ ThumbnailGenerationRequest) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](t_.ID, objc.Sel("contentType"))
	return rv
}/* debug [instance_properties/getter]: contentType */


// The content type of the source data for the thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/contentType
func (t_ ThumbnailGenerationRequest) SetContentType(value uniformtypeidentifiers.UTType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContentType:"), value)
}/* debug [instance_properties/setter]: contentType */


// A Boolean value indicating whether the generated thumbnail request should include icon decorations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/iconMode
func (t_ ThumbnailGenerationRequest) IconMode() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("iconMode"))
	return rv
}/* debug [instance_properties/getter]: iconMode */


// A Boolean value indicating whether the generated thumbnail request should include icon decorations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/iconMode
func (t_ ThumbnailGenerationRequest) SetIconMode(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIconMode:"), value)
}/* debug [instance_properties/setter]: iconMode */


// The minimum height or width for a generated thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/minimumDimension
func (t_ ThumbnailGenerationRequest) MinimumDimension() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("minimumDimension"))
	return rv
}/* debug [instance_properties/getter]: minimumDimension */


// The minimum height or width for a generated thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/minimumDimension
func (t_ ThumbnailGenerationRequest) SetMinimumDimension(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinimumDimension:"), value)
}/* debug [instance_properties/setter]: minimumDimension */


// The thumbnail sizes that you provide for a thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/representationTypes-swift.property
func (t_ ThumbnailGenerationRequest) RepresentationTypes() ThumbnailGenerationRequestRepresentationTypes {
	rv := objc.Send[ThumbnailGenerationRequestRepresentationTypes](t_.ID, objc.Sel("representationTypes"))
	return rv
}/* debug [instance_properties/getter]: representationTypes */


// The pixel density of the display on the intended device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/scale
func (t_ ThumbnailGenerationRequest) Scale() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("scale"))
	return rv
}/* debug [instance_properties/getter]: scale */


// The size of the thumbnails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/size
func (t_ ThumbnailGenerationRequest) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLThumbnailGenerationRequest */


