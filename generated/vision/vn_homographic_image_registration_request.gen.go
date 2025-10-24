// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HomographicImageRegistrationRequest] class.
var (
	HomographicImageRegistrationRequestClass     _HomographicImageRegistrationRequestClass
	HomographicImageRegistrationRequestClassOnce sync.Once
)

func getHomographicImageRegistrationRequestClass() _HomographicImageRegistrationRequestClass {
	HomographicImageRegistrationRequestClassOnce.Do(func() {
		HomographicImageRegistrationRequestClass = _HomographicImageRegistrationRequestClass{objc.GetClass("VNHomographicImageRegistrationRequest")}
	})
	return HomographicImageRegistrationRequestClass
}

type _HomographicImageRegistrationRequestClass struct {
	class objc.Class
}

// An interface definition for the [HomographicImageRegistrationRequest] class.
type IHomographicImageRegistrationRequest interface {
	IImageRegistrationRequest
	// properties:
	Results() IVNImageHomographicAlignmentObservation
	SetResults(value IVNImageHomographicAlignmentObservation)
	VNHomographicImageRegistrationRequestRevision1() int
	// methods:
}

// An image-analysis request that determines the perspective warp matrix necessary to align the content of two images.
//
// Create and perform a homographic image registration request to align content in two images through a homography. A is an isomorphism of projected spaces, a bijection that maps lines to lines.


// An image-analysis request that determines the perspective warp matrix necessary to align the content of two images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHomographicImageRegistrationRequest
type HomographicImageRegistrationRequest struct {
	ImageRegistrationRequest
}

// HomographicImageRegistrationRequestFrom constructs a [HomographicImageRegistrationRequest] from an unsafe.Pointer.
//
// An image-analysis request that determines the perspective warp matrix necessary to align the content of two images.
func HomographicImageRegistrationRequestFrom(ptr unsafe.Pointer) HomographicImageRegistrationRequest {
	return HomographicImageRegistrationRequest{
		ImageRegistrationRequest: ImageRegistrationRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HomographicImageRegistrationRequestClass) Alloc() HomographicImageRegistrationRequest {
	rv := objc.Send[HomographicImageRegistrationRequest](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HomographicImageRegistrationRequestClass) New() HomographicImageRegistrationRequest {
	rv := objc.Send[HomographicImageRegistrationRequest](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HomographicImageRegistrationRequest) Init() HomographicImageRegistrationRequest {
	rv := objc.Send[HomographicImageRegistrationRequest](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HomographicImageRegistrationRequest) Autorelease() HomographicImageRegistrationRequest {
	rv := objc.Send[HomographicImageRegistrationRequest](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHomographicImageRegistrationRequest creates a new HomographicImageRegistrationRequest instance.
func NewHomographicImageRegistrationRequest() HomographicImageRegistrationRequest {
	return getHomographicImageRegistrationRequestClass().New()
}



// The results of the image registration request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhomographicimageregistrationrequest/results
func (h_ HomographicImageRegistrationRequest) Results() IVNImageHomographicAlignmentObservation {
	rv := objc.Send[ImageHomographicAlignmentObservation](h_.ID, objc.Sel("results"))
	return rv
}


// The results of the image registration request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhomographicimageregistrationrequest/results
func (h_ HomographicImageRegistrationRequest) SetResults(value IVNImageHomographicAlignmentObservation) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setResults:"), value)
}


// A constant for specifying revision 1 of the homographic image registration request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhomographicimageregistrationrequestrevision1
func (h_ HomographicImageRegistrationRequest) VNHomographicImageRegistrationRequestRevision1() int {
	rv := objc.Send[int](h_.ID, objc.Sel("VNHomographicImageRegistrationRequestRevision1"))
	return rv
}



