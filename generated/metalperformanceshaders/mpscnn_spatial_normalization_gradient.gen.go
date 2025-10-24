// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNSpatialNormalizationGradient] class.
var (
	CNNSpatialNormalizationGradientClass     _CNNSpatialNormalizationGradientClass
	CNNSpatialNormalizationGradientClassOnce sync.Once
)

func getCNNSpatialNormalizationGradientClass() _CNNSpatialNormalizationGradientClass {
	CNNSpatialNormalizationGradientClassOnce.Do(func() {
		CNNSpatialNormalizationGradientClass = _CNNSpatialNormalizationGradientClass{objc.GetClass("MPSCNNSpatialNormalizationGradient")}
	})
	return CNNSpatialNormalizationGradientClass
}

type _CNNSpatialNormalizationGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNSpatialNormalizationGradient] class.
type ICNNSpatialNormalizationGradient interface {
	ICNNGradientKernel
	

	// properties:
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNSpatialNormalizationGradientClass) Alloc() CNNSpatialNormalizationGradient {
	rv := objc.Send[CNNSpatialNormalizationGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSpatialNormalizationGradientClass) New() CNNSpatialNormalizationGradient {
	rv := objc.Send[CNNSpatialNormalizationGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSpatialNormalizationGradient) Init() CNNSpatialNormalizationGradient {
	rv := objc.Send[CNNSpatialNormalizationGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSpatialNormalizationGradient) Autorelease() CNNSpatialNormalizationGradient {
	rv := objc.Send[CNNSpatialNormalizationGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSpatialNormalizationGradient creates a new CNNSpatialNormalizationGradient instance.
func NewCNNSpatialNormalizationGradient() CNNSpatialNormalizationGradient {
	return getCNNSpatialNormalizationGradientClass().New()
}





// A gradient spatial normalization kernel.


// A gradient spatial normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradient
type CNNSpatialNormalizationGradient struct {
	CNNGradientKernel
}

// CNNSpatialNormalizationGradientFrom constructs a [CNNSpatialNormalizationGradient] from an unsafe.Pointer.
//
// A gradient spatial normalization kernel.
func CNNSpatialNormalizationGradientFrom(ptr unsafe.Pointer) CNNSpatialNormalizationGradient {
	return CNNSpatialNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942473-initwithcoder
func NewCNNSpatialNormalizationGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNSpatialNormalizationGradient {
	instance := getCNNSpatialNormalizationGradientClass().Alloc()
	rv := objc.Send[CNNSpatialNormalizationGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942461-initwithdevice
func NewCNNSpatialNormalizationGradientWithDeviceKernelWidthKernelHeight(device unsafe.Pointer, kernelWidth uint, kernelHeight uint) CNNSpatialNormalizationGradient {
	instance := getCNNSpatialNormalizationGradientClass().Alloc()
	rv := objc.Send[CNNSpatialNormalizationGradient](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942470-beta
func (c_ CNNSpatialNormalizationGradient) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("beta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942470-beta
func (c_ CNNSpatialNormalizationGradient) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942478-alpha
func (c_ CNNSpatialNormalizationGradient) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942478-alpha
func (c_ CNNSpatialNormalizationGradient) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942486-delta
func (c_ CNNSpatialNormalizationGradient) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942486-delta
func (c_ CNNSpatialNormalizationGradient) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}







