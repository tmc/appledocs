// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GenerateOpticalFlowRequest] class.
var (
	GenerateOpticalFlowRequestClass     _GenerateOpticalFlowRequestClass
	GenerateOpticalFlowRequestClassOnce sync.Once
)

func getGenerateOpticalFlowRequestClass() _GenerateOpticalFlowRequestClass {
	GenerateOpticalFlowRequestClassOnce.Do(func() {
		GenerateOpticalFlowRequestClass = _GenerateOpticalFlowRequestClass{objc.GetClass("VNGenerateOpticalFlowRequest")}
	})
	return GenerateOpticalFlowRequestClass
}

type _GenerateOpticalFlowRequestClass struct {
	class objc.Class
}

// An interface definition for the [GenerateOpticalFlowRequest] class.
type IGenerateOpticalFlowRequest interface {
	ITargetedImageRequest
}

// An object that generates directional change vectors for each pixel in the targeted image.
//
// This request operates at a pixel level, so both images need to have the same dimensions to successfully perform the analysis. Setting a region of interest limits the region in which the analysis occurs. However, the system reports the resulting observation at full resolution. Optical flow requests are resource-intensive, so create only one request at a time, and release it immediately after generating optical flows.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest
type GenerateOpticalFlowRequest struct {
	TargetedImageRequest
}

// GenerateOpticalFlowRequestFrom constructs a [GenerateOpticalFlowRequest] from an unsafe.Pointer.
//
// An object that generates directional change vectors for each pixel in the targeted image.
func GenerateOpticalFlowRequestFrom(ptr unsafe.Pointer) GenerateOpticalFlowRequest {
	return GenerateOpticalFlowRequest{
		TargetedImageRequest: TargetedImageRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GenerateOpticalFlowRequestClass) Alloc() GenerateOpticalFlowRequest {
	rv := objc.Send[GenerateOpticalFlowRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GenerateOpticalFlowRequestClass) New() GenerateOpticalFlowRequest {
	rv := objc.Send[GenerateOpticalFlowRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenerateOpticalFlowRequest) Init() GenerateOpticalFlowRequest {
	rv := objc.Send[GenerateOpticalFlowRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenerateOpticalFlowRequest) Autorelease() GenerateOpticalFlowRequest {
	rv := objc.Send[GenerateOpticalFlowRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenerateOpticalFlowRequest creates a new GenerateOpticalFlowRequest instance.
func NewGenerateOpticalFlowRequest() GenerateOpticalFlowRequest {
	return getGenerateOpticalFlowRequestClass().New()
}




