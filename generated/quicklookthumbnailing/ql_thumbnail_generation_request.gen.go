// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

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

// An interface definition for the [ThumbnailGenerationRequest] class.
type IThumbnailGenerationRequest interface {
	objectivec.IObject
	// properties:
	ContentType() objc.IObject /* cross-framework: UTType */
	SetContentType(value objc.IObject /* cross-framework: UTType */)
	IconMode() bool
	SetIconMode(value bool)
	MinimumDimension() float64
	SetMinimumDimension(value float64)
	RepresentationTypes() unsafe.Pointer
	SetRepresentationTypes(value unsafe.Pointer)
	Scale() float64
	SetScale(value float64)
	Size() objc.IObject /* cross-framework: Size */
	SetSize(value objc.IObject /* cross-framework: Size */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (tc _ThumbnailGenerationRequestClass) Alloc() ThumbnailGenerationRequest {
	rv := objc.Send[ThumbnailGenerationRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The content type of the source data for the thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/contenttype
func (t_ ThumbnailGenerationRequest) ContentType() objc.IObject /* cross-framework: UTType */ {
	rv := objc.Send[uniformtypeidentifiers.UTType](t_.ID, objc.Sel("contentType"))
	return rv
}


// The content type of the source data for the thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/contenttype
func (t_ ThumbnailGenerationRequest) SetContentType(value objc.IObject /* cross-framework: UTType */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContentType:"), value)
}


// A Boolean value indicating whether the generated thumbnail request should include icon decorations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/iconmode
func (t_ ThumbnailGenerationRequest) IconMode() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("iconMode"))
	return rv
}


// A Boolean value indicating whether the generated thumbnail request should include icon decorations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/iconmode
func (t_ ThumbnailGenerationRequest) SetIconMode(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIconMode:"), value)
}


// The minimum height or width for a generated thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/minimumdimension
func (t_ ThumbnailGenerationRequest) MinimumDimension() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("minimumDimension"))
	return rv
}


// The minimum height or width for a generated thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/minimumdimension
func (t_ ThumbnailGenerationRequest) SetMinimumDimension(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinimumDimension:"), value)
}


// The thumbnail sizes that you provide for a thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/representationtypes-swift.property
func (t_ ThumbnailGenerationRequest) RepresentationTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("representationTypes"))
	return rv
}


// The thumbnail sizes that you provide for a thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/representationtypes-swift.property
func (t_ ThumbnailGenerationRequest) SetRepresentationTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRepresentationTypes:"), value)
}


// The pixel density of the display on the intended device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/scale
func (t_ ThumbnailGenerationRequest) Scale() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("scale"))
	return rv
}


// The pixel density of the display on the intended device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/scale
func (t_ ThumbnailGenerationRequest) SetScale(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setScale:"), value)
}


// The size of the thumbnails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/size
func (t_ ThumbnailGenerationRequest) Size() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](t_.ID, objc.Sel("size"))
	return rv
}


// The size of the thumbnails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/size
func (t_ ThumbnailGenerationRequest) SetSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSize:"), value)
}



