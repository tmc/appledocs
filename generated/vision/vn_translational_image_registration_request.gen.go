// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TranslationalImageRegistrationRequest] class.
var (
	TranslationalImageRegistrationRequestClass     _TranslationalImageRegistrationRequestClass
	TranslationalImageRegistrationRequestClassOnce sync.Once
)

func getTranslationalImageRegistrationRequestClass() _TranslationalImageRegistrationRequestClass {
	TranslationalImageRegistrationRequestClassOnce.Do(func() {
		TranslationalImageRegistrationRequestClass = _TranslationalImageRegistrationRequestClass{objc.GetClass("VNTranslationalImageRegistrationRequest")}
	})
	return TranslationalImageRegistrationRequestClass
}

type _TranslationalImageRegistrationRequestClass struct {
	class objc.Class
}

// An interface definition for the [TranslationalImageRegistrationRequest] class.
type ITranslationalImageRegistrationRequest interface {
	IImageRegistrationRequest
}

// An image-analysis request that determines the affine transform necessary to align the content of two images.
//
// Create and perform a translational image registration request to align content in two images through translation.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTranslationalImageRegistrationRequest
type TranslationalImageRegistrationRequest struct {
	ImageRegistrationRequest
}

// TranslationalImageRegistrationRequestFrom constructs a [TranslationalImageRegistrationRequest] from an unsafe.Pointer.
//
// An image-analysis request that determines the affine transform necessary to align the content of two images.
func TranslationalImageRegistrationRequestFrom(ptr unsafe.Pointer) TranslationalImageRegistrationRequest {
	return TranslationalImageRegistrationRequest{
		ImageRegistrationRequest: ImageRegistrationRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TranslationalImageRegistrationRequestClass) Alloc() TranslationalImageRegistrationRequest {
	rv := objc.Send[TranslationalImageRegistrationRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TranslationalImageRegistrationRequestClass) New() TranslationalImageRegistrationRequest {
	rv := objc.Send[TranslationalImageRegistrationRequest](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TranslationalImageRegistrationRequest) Init() TranslationalImageRegistrationRequest {
	rv := objc.Send[TranslationalImageRegistrationRequest](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TranslationalImageRegistrationRequest) Autorelease() TranslationalImageRegistrationRequest {
	rv := objc.Send[TranslationalImageRegistrationRequest](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTranslationalImageRegistrationRequest creates a new TranslationalImageRegistrationRequest instance.
func NewTranslationalImageRegistrationRequest() TranslationalImageRegistrationRequest {
	return getTranslationalImageRegistrationRequestClass().New()
}


// The results of a translational image alignment request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntranslationalimageregistrationrequest/results
func (t_ TranslationalImageRegistrationRequest) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of a translational image alignment request.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntranslationalimageregistrationrequest/results
func (t_ TranslationalImageRegistrationRequest) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResults:"), value)
}



