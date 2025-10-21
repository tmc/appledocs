// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
}

// A request to generate a thumbnail for a file.
//
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




// Creates a new request for a thumbnail with the specified parameters for a file at a provided URL.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/init(fileAt:size:scale:representationTypes:)
func NewThumbnailGenerationRequestWithFileAtURLSizeScaleRepresentationTypes(url foundation.URL, size coregraphics.CGSize, scale float64, representationTypes unsafe.Pointer) ThumbnailGenerationRequest {
	instance := getThumbnailGenerationRequestClass().Alloc()
	rv := objc.Send[ThumbnailGenerationRequest](instance.ID, objc.Sel("initWithFileAtURL:size:scale:representationTypes:"), url, size, scale, representationTypes)
	rv.Autorelease()
	return rv
}


// A Boolean value indicating whether the generated thumbnail request should include icon decorations.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/iconmode
func (t_ ThumbnailGenerationRequest) IconMode() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("iconMode"))
	return rv
}


// SetIconMode sets the value of the iconMode property.
// A Boolean value indicating whether the generated thumbnail request should include icon decorations.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/iconmode
func (t_ ThumbnailGenerationRequest) SetIconMode(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIconMode:"), value)
}

// The size of the thumbnails.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/size
func (t_ ThumbnailGenerationRequest) Size() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("size"))
	return rv
}


// SetSize sets the value of the size property.
// The size of the thumbnails.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/size
func (t_ ThumbnailGenerationRequest) SetSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSize:"), value)
}

// The content type of the source data for the thumbnail request.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/contenttype
func (t_ ThumbnailGenerationRequest) ContentType() UTType {
	rv := objc.Send[UTType](t_.ID, objc.Sel("contentType"))
	return rv
}


// SetContentType sets the value of the contentType property.
// The content type of the source data for the thumbnail request.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailgenerator/request/contenttype
func (t_ ThumbnailGenerationRequest) SetContentType(value UTType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContentType:"), value)
}

// The minimum height or width for a generated thumbnail.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/minimumDimension
func (t_ ThumbnailGenerationRequest) MinimumDimension() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("minimumDimension"))
	return rv
}


// SetMinimumDimension sets the value of the minimumDimension property.
// The minimum height or width for a generated thumbnail.

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/minimumDimension
func (t_ ThumbnailGenerationRequest) SetMinimumDimension(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinimumDimension:"), value)
}

// The thumbnail sizes that you provide for a thumbnail request.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/representationTypes-swift.property
func (t_ ThumbnailGenerationRequest) RepresentationTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("representationTypes"))
	return rv
}

// The pixel density of the display on the intended device.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/scale
func (t_ ThumbnailGenerationRequest) Scale() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("scale"))
	return rv
}


