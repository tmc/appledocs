// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DetectAnimalBodyPoseRequest] class.
var (
	DetectAnimalBodyPoseRequestClass     _DetectAnimalBodyPoseRequestClass
	DetectAnimalBodyPoseRequestClassOnce sync.Once
)

func getDetectAnimalBodyPoseRequestClass() _DetectAnimalBodyPoseRequestClass {
	DetectAnimalBodyPoseRequestClassOnce.Do(func() {
		DetectAnimalBodyPoseRequestClass = _DetectAnimalBodyPoseRequestClass{objc.GetClass("VNDetectAnimalBodyPoseRequest")}
	})
	return DetectAnimalBodyPoseRequestClass
}

type _DetectAnimalBodyPoseRequestClass struct {
	class objc.Class
}

// An interface definition for the [DetectAnimalBodyPoseRequest] class.
type IDetectAnimalBodyPoseRequest interface {
	IImageBasedRequest
}

// A request that detects an animal body pose.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest
type DetectAnimalBodyPoseRequest struct {
	ImageBasedRequest
}

// DetectAnimalBodyPoseRequestFrom constructs a [DetectAnimalBodyPoseRequest] from an unsafe.Pointer.
//
// A request that detects an animal body pose.
func DetectAnimalBodyPoseRequestFrom(ptr unsafe.Pointer) DetectAnimalBodyPoseRequest {
	return DetectAnimalBodyPoseRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DetectAnimalBodyPoseRequestClass) Alloc() DetectAnimalBodyPoseRequest {
	rv := objc.Send[DetectAnimalBodyPoseRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DetectAnimalBodyPoseRequestClass) New() DetectAnimalBodyPoseRequest {
	rv := objc.Send[DetectAnimalBodyPoseRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectAnimalBodyPoseRequest) Init() DetectAnimalBodyPoseRequest {
	rv := objc.Send[DetectAnimalBodyPoseRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectAnimalBodyPoseRequest) Autorelease() DetectAnimalBodyPoseRequest {
	rv := objc.Send[DetectAnimalBodyPoseRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectAnimalBodyPoseRequest creates a new DetectAnimalBodyPoseRequest instance.
func NewDetectAnimalBodyPoseRequest() DetectAnimalBodyPoseRequest {
	return getDetectAnimalBodyPoseRequestClass().New()
}




