// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CalculateImageAestheticsScoresRequest] class.
var (
	CalculateImageAestheticsScoresRequestClass     _CalculateImageAestheticsScoresRequestClass
	CalculateImageAestheticsScoresRequestClassOnce sync.Once
)

func getCalculateImageAestheticsScoresRequestClass() _CalculateImageAestheticsScoresRequestClass {
	CalculateImageAestheticsScoresRequestClassOnce.Do(func() {
		CalculateImageAestheticsScoresRequestClass = _CalculateImageAestheticsScoresRequestClass{objc.GetClass("VNCalculateImageAestheticsScoresRequest")}
	})
	return CalculateImageAestheticsScoresRequestClass
}

type _CalculateImageAestheticsScoresRequestClass struct {
	class objc.Class
}





// An interface definition for the [CalculateImageAestheticsScoresRequest] class.
type ICalculateImageAestheticsScoresRequest interface {
	IImageBasedRequest
	

	// properties:
	Results() []ImageAestheticsScoresObservation


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CalculateImageAestheticsScoresRequestClass) Alloc() CalculateImageAestheticsScoresRequest {
	rv := objc.Send[CalculateImageAestheticsScoresRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CalculateImageAestheticsScoresRequestClass) New() CalculateImageAestheticsScoresRequest {
	rv := objc.Send[CalculateImageAestheticsScoresRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CalculateImageAestheticsScoresRequest) Init() CalculateImageAestheticsScoresRequest {
	rv := objc.Send[CalculateImageAestheticsScoresRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CalculateImageAestheticsScoresRequest) Autorelease() CalculateImageAestheticsScoresRequest {
	rv := objc.Send[CalculateImageAestheticsScoresRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCalculateImageAestheticsScoresRequest creates a new CalculateImageAestheticsScoresRequest instance.
func NewCalculateImageAestheticsScoresRequest() CalculateImageAestheticsScoresRequest {
	return getCalculateImageAestheticsScoresRequestClass().New()
}





// An object that analyzes an image for aesthetically pleasing attributes.


// An object that analyzes an image for aesthetically pleasing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCalculateImageAestheticsScoresRequest
type CalculateImageAestheticsScoresRequest struct {
	ImageBasedRequest
}

// CalculateImageAestheticsScoresRequestFrom constructs a [CalculateImageAestheticsScoresRequest] from an unsafe.Pointer.
//
// An object that analyzes an image for aesthetically pleasing attributes.
func CalculateImageAestheticsScoresRequestFrom(ptr unsafe.Pointer) CalculateImageAestheticsScoresRequest {
	return CalculateImageAestheticsScoresRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

























// The results of the aesthetics request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCalculateImageAestheticsScoresRequest/results
func (c_ CalculateImageAestheticsScoresRequest) Results() []ImageAestheticsScoresObservation {
	rv := objc.Send[[]ImageAestheticsScoresObservation](c_.ID, objc.Sel("results"))
	return rv
}








