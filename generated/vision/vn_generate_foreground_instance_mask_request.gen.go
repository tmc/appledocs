// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [GenerateForegroundInstanceMaskRequest] class.
var (
	GenerateForegroundInstanceMaskRequestClass     _GenerateForegroundInstanceMaskRequestClass
	GenerateForegroundInstanceMaskRequestClassOnce sync.Once
)

func getGenerateForegroundInstanceMaskRequestClass() _GenerateForegroundInstanceMaskRequestClass {
	GenerateForegroundInstanceMaskRequestClassOnce.Do(func() {
		GenerateForegroundInstanceMaskRequestClass = _GenerateForegroundInstanceMaskRequestClass{objc.GetClass("VNGenerateForegroundInstanceMaskRequest")}
	})
	return GenerateForegroundInstanceMaskRequestClass
}

type _GenerateForegroundInstanceMaskRequestClass struct {
	class objc.Class
}





// An interface definition for the [GenerateForegroundInstanceMaskRequest] class.
type IGenerateForegroundInstanceMaskRequest interface {
	IImageBasedRequest
	

	// properties:
	Results() []InstanceMaskObservation
	VNGenerateForegroundInstanceMaskRequestRevision1() int


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (gc _GenerateForegroundInstanceMaskRequestClass) Alloc() GenerateForegroundInstanceMaskRequest {
	rv := objc.Send[GenerateForegroundInstanceMaskRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GenerateForegroundInstanceMaskRequestClass) New() GenerateForegroundInstanceMaskRequest {
	rv := objc.Send[GenerateForegroundInstanceMaskRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenerateForegroundInstanceMaskRequest) Init() GenerateForegroundInstanceMaskRequest {
	rv := objc.Send[GenerateForegroundInstanceMaskRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenerateForegroundInstanceMaskRequest) Autorelease() GenerateForegroundInstanceMaskRequest {
	rv := objc.Send[GenerateForegroundInstanceMaskRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenerateForegroundInstanceMaskRequest creates a new GenerateForegroundInstanceMaskRequest instance.
func NewGenerateForegroundInstanceMaskRequest() GenerateForegroundInstanceMaskRequest {
	return getGenerateForegroundInstanceMaskRequestClass().New()
}





// A request that generates an instance mask of noticable objects to separate from the background.


// A request that generates an instance mask of noticable objects to separate from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateForegroundInstanceMaskRequest
type GenerateForegroundInstanceMaskRequest struct {
	ImageBasedRequest
}

// GenerateForegroundInstanceMaskRequestFrom constructs a [GenerateForegroundInstanceMaskRequest] from an unsafe.Pointer.
//
// A request that generates an instance mask of noticable objects to separate from the background.
func GenerateForegroundInstanceMaskRequestFrom(ptr unsafe.Pointer) GenerateForegroundInstanceMaskRequest {
	return GenerateForegroundInstanceMaskRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

























// The instance masks the request observes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateForegroundInstanceMaskRequest/results
func (g_ GenerateForegroundInstanceMaskRequest) Results() []InstanceMaskObservation {
	rv := objc.Send[[]InstanceMaskObservation](g_.ID, objc.Sel("results"))
	return rv
}


// A constant for specifying the first revision of the foreground instance mask request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateforegroundinstancemaskrequestrevision1
func (g_ GenerateForegroundInstanceMaskRequest) VNGenerateForegroundInstanceMaskRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGenerateForegroundInstanceMaskRequestRevision1"))
	return rv
}








