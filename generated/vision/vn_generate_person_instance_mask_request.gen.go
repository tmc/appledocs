// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GeneratePersonInstanceMaskRequest] class.
var (
	GeneratePersonInstanceMaskRequestClass     _GeneratePersonInstanceMaskRequestClass
	GeneratePersonInstanceMaskRequestClassOnce sync.Once
)

func getGeneratePersonInstanceMaskRequestClass() _GeneratePersonInstanceMaskRequestClass {
	GeneratePersonInstanceMaskRequestClassOnce.Do(func() {
		GeneratePersonInstanceMaskRequestClass = _GeneratePersonInstanceMaskRequestClass{objc.GetClass("VNGeneratePersonInstanceMaskRequest")}
	})
	return GeneratePersonInstanceMaskRequestClass
}

type _GeneratePersonInstanceMaskRequestClass struct {
	class objc.Class
}

// An interface definition for the [GeneratePersonInstanceMaskRequest] class.
type IGeneratePersonInstanceMaskRequest interface {
	IImageBasedRequest
}

// An object that produces a mask of individual people it finds in the input image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonInstanceMaskRequest
type GeneratePersonInstanceMaskRequest struct {
	ImageBasedRequest
}

// GeneratePersonInstanceMaskRequestFrom constructs a [GeneratePersonInstanceMaskRequest] from an unsafe.Pointer.
//
// An object that produces a mask of individual people it finds in the input image.
func GeneratePersonInstanceMaskRequestFrom(ptr unsafe.Pointer) GeneratePersonInstanceMaskRequest {
	return GeneratePersonInstanceMaskRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GeneratePersonInstanceMaskRequestClass) Alloc() GeneratePersonInstanceMaskRequest {
	rv := objc.Send[GeneratePersonInstanceMaskRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GeneratePersonInstanceMaskRequestClass) New() GeneratePersonInstanceMaskRequest {
	rv := objc.Send[GeneratePersonInstanceMaskRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GeneratePersonInstanceMaskRequest) Init() GeneratePersonInstanceMaskRequest {
	rv := objc.Send[GeneratePersonInstanceMaskRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GeneratePersonInstanceMaskRequest) Autorelease() GeneratePersonInstanceMaskRequest {
	rv := objc.Send[GeneratePersonInstanceMaskRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGeneratePersonInstanceMaskRequest creates a new GeneratePersonInstanceMaskRequest instance.
func NewGeneratePersonInstanceMaskRequest() GeneratePersonInstanceMaskRequest {
	return getGeneratePersonInstanceMaskRequestClass().New()
}


// The results of the instance mask request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersoninstancemaskrequest/results
func (g_ GeneratePersonInstanceMaskRequest) Results() VNInstanceMaskObservation {
	rv := objc.Send[VNInstanceMaskObservation](g_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of the instance mask request.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersoninstancemaskrequest/results
func (g_ GeneratePersonInstanceMaskRequest) SetResults(value IVNInstanceMaskObservation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResults:"), value)
}

// A constant for specifying revision 1 of the person instance mask request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersoninstancemaskrequestrevision1
func (g_ GeneratePersonInstanceMaskRequest) VNGeneratePersonInstanceMaskRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGeneratePersonInstanceMaskRequestRevision1"))
	return rv
}



