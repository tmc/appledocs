// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNSpatialNormalization] class.
var (
	CNNSpatialNormalizationClass     _CNNSpatialNormalizationClass
	CNNSpatialNormalizationClassOnce sync.Once
)

func getCNNSpatialNormalizationClass() _CNNSpatialNormalizationClass {
	CNNSpatialNormalizationClassOnce.Do(func() {
		CNNSpatialNormalizationClass = _CNNSpatialNormalizationClass{objc.GetClass("MPSCNNSpatialNormalization")}
	})
	return CNNSpatialNormalizationClass
}

type _CNNSpatialNormalizationClass struct {
	class objc.Class
}





// An interface definition for the [CNNSpatialNormalization] class.
type ICNNSpatialNormalization interface {
	ICNNKernel
	

	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNSpatialNormalizationClass) Alloc() CNNSpatialNormalization {
	rv := objc.Send[CNNSpatialNormalization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSpatialNormalizationClass) New() CNNSpatialNormalization {
	rv := objc.Send[CNNSpatialNormalization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSpatialNormalization) Init() CNNSpatialNormalization {
	rv := objc.Send[CNNSpatialNormalization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSpatialNormalization) Autorelease() CNNSpatialNormalization {
	rv := objc.Send[CNNSpatialNormalization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSpatialNormalization creates a new CNNSpatialNormalization instance.
func NewCNNSpatialNormalization() CNNSpatialNormalization {
	return getCNNSpatialNormalizationClass().New()
}





// A spatial normalization kernel.
//
// The spatial normalization for a feature channel applies the kernel over local regions which extend spatially, but are in separate feature channels (i.e., they have the shape ). For each feature channel, the function computes the sum of squares of inside each rectangle, . It then divides each element of as follows: Where and are the values of the and properties, respectively. It is your responsibility to ensure that the combination of the values of the and properties does not result in a situation where the denominator becomes zero (in such situations the resulting pixel-value is undefined).


// A spatial normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalization
type CNNSpatialNormalization struct {
	CNNKernel
}

// CNNSpatialNormalizationFrom constructs a [CNNSpatialNormalization] from an unsafe.Pointer.
//
// A spatial normalization kernel.
func CNNSpatialNormalizationFrom(ptr unsafe.Pointer) CNNSpatialNormalization {
	return CNNSpatialNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}






// Initializes a spatial normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalization/init(coder:device:)
func NewCNNSpatialNormalizationWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNSpatialNormalization {
	instance := getCNNSpatialNormalizationClass().Alloc()
	rv := objc.Send[CNNSpatialNormalization](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes a spatial normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648831-initwithdevice
func NewCNNSpatialNormalizationWithDeviceKernelWidthKernelHeight(device unsafe.Pointer, kernelWidth uint, kernelHeight uint) CNNSpatialNormalization {
	instance := getCNNSpatialNormalizationClass().Alloc()
	rv := objc.Send[CNNSpatialNormalization](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	rv.Autorelease()
	return rv
}






















// The "alpha" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648825-alpha
func (c_ CNNSpatialNormalization) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alpha"))
	return rv
}


// The "alpha" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648825-alpha
func (c_ CNNSpatialNormalization) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}


// The "delta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648933-delta
func (c_ CNNSpatialNormalization) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}


// The "delta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648933-delta
func (c_ CNNSpatialNormalization) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}


// The "beta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648936-beta
func (c_ CNNSpatialNormalization) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("beta"))
	return rv
}


// The "beta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648936-beta
func (c_ CNNSpatialNormalization) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}







