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
	ICNNKernel
	

	// properties:
	NumberOfRegions() objectivec.IObject
	SetNumberOfRegions(value objectivec.IObject)
	Regions() objectivec.IObject
	SetRegions(value objectivec.IObject)
	ResizeHeight() objectivec.IObject
	SetResizeHeight(value objectivec.IObject)
	ResizeWidth() objectivec.IObject
	SetResizeWidth(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CropAndResizeBilinearClass) Alloc() CropAndResizeBilinear {
	rv := objc.Send[CropAndResizeBilinear](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A cropping and bilinear resizing filter.


// A cropping and bilinear resizing filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear
type CropAndResizeBilinear struct {
	CNNKernel
}

// CropAndResizeBilinearFrom constructs a [CropAndResizeBilinear] from an unsafe.Pointer.
//
// A cropping and bilinear resizing filter.
func CropAndResizeBilinearFrom(ptr unsafe.Pointer) CropAndResizeBilinear {
	return CropAndResizeBilinear{
		CNNKernel: CNNKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013788-initwithcoder
func NewCropAndResizeBilinearWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CropAndResizeBilinear {
	instance := getCropAndResizeBilinearClass().Alloc()
	rv := objc.Send[CropAndResizeBilinear](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013789-initwithdevice
func NewCropAndResizeBilinearWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions(device unsafe.Pointer, resizeWidth uint, resizeHeight uint, numberOfRegions uint, regions objc.IObject /* cross-framework: MPSRegion */) CropAndResizeBilinear {
	instance := getCropAndResizeBilinearClass().Alloc()
	rv := objc.Send[CropAndResizeBilinear](instance.ID, objc.Sel("initWithDevice:resizeWidth:resizeHeight:numberOfRegions:regions:"), device, resizeWidth, resizeHeight, numberOfRegions, regions)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013790-numberofregions
func (c_ CropAndResizeBilinear) NumberOfRegions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("numberOfRegions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013790-numberofregions
func (c_ CropAndResizeBilinear) SetNumberOfRegions(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfRegions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013791-regions
func (c_ CropAndResizeBilinear) Regions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("regions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013791-regions
func (c_ CropAndResizeBilinear) SetRegions(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRegions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013792-resizeheight
func (c_ CropAndResizeBilinear) ResizeHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("resizeHeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013792-resizeheight
func (c_ CropAndResizeBilinear) SetResizeHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResizeHeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013793-resizewidth
func (c_ CropAndResizeBilinear) ResizeWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("resizeWidth"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013793-resizewidth
func (c_ CropAndResizeBilinear) SetResizeWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResizeWidth:"), value)
}







