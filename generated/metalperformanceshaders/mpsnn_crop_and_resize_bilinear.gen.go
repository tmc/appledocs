// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CropAndResizeBilinear] class.
var (
	CropAndResizeBilinearClass     _CropAndResizeBilinearClass
	CropAndResizeBilinearClassOnce sync.Once
)

func getCropAndResizeBilinearClass() _CropAndResizeBilinearClass {
	CropAndResizeBilinearClassOnce.Do(func() {
		CropAndResizeBilinearClass = _CropAndResizeBilinearClass{objc.GetClass("MPSNNCropAndResizeBilinear")}
	})
	return CropAndResizeBilinearClass
}

type _CropAndResizeBilinearClass struct {
	class objc.Class
}

// An interface definition for the [CropAndResizeBilinear] class.
type ICropAndResizeBilinear interface {
	objectivec.IObject
}

// A cropping and bilinear resizing filter.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear
type CropAndResizeBilinear struct {
	objectivec.Object
}

// CropAndResizeBilinearFrom constructs a [CropAndResizeBilinear] from an unsafe.Pointer.
//
// A cropping and bilinear resizing filter.
func CropAndResizeBilinearFrom(ptr unsafe.Pointer) CropAndResizeBilinear {
	return CropAndResizeBilinear{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CropAndResizeBilinearClass) Alloc() CropAndResizeBilinear {
	rv := objc.Send[CropAndResizeBilinear](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CropAndResizeBilinearClass) New() CropAndResizeBilinear {
	rv := objc.Send[CropAndResizeBilinear](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CropAndResizeBilinear) Init() CropAndResizeBilinear {
	rv := objc.Send[CropAndResizeBilinear](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CropAndResizeBilinear) Autorelease() CropAndResizeBilinear {
	rv := objc.Send[CropAndResizeBilinear](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCropAndResizeBilinear creates a new CropAndResizeBilinear instance.
func NewCropAndResizeBilinear() CropAndResizeBilinear {
	return getCropAndResizeBilinearClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear/init(coder:device:)
func NewCropAndResizeBilinearWithCoderDevice(aDecoder unsafe.Pointer, device objc.ID) CropAndResizeBilinear {
	instance := getCropAndResizeBilinearClass().Alloc()
	rv := objc.Send[CropAndResizeBilinear](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear/regions
func (c_ CropAndResizeBilinear) Regions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("regions"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear/resizeHeight
func (c_ CropAndResizeBilinear) ResizeHeight() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("resizeHeight"))
	return rv
}


