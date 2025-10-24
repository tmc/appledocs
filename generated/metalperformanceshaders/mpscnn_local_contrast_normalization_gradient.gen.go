// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNLocalContrastNormalizationGradient] class.
var (
	CNNLocalContrastNormalizationGradientClass     _CNNLocalContrastNormalizationGradientClass
	CNNLocalContrastNormalizationGradientClassOnce sync.Once
)

func getCNNLocalContrastNormalizationGradientClass() _CNNLocalContrastNormalizationGradientClass {
	CNNLocalContrastNormalizationGradientClassOnce.Do(func() {
		CNNLocalContrastNormalizationGradientClass = _CNNLocalContrastNormalizationGradientClass{objc.GetClass("MPSCNNLocalContrastNormalizationGradient")}
	})
	return CNNLocalContrastNormalizationGradientClass
}

type _CNNLocalContrastNormalizationGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNLocalContrastNormalizationGradient] class.
type ICNNLocalContrastNormalizationGradient interface {
	ICNNGradientKernel
	

	// properties:
	Ps() objectivec.IObject
	SetPs(value objectivec.IObject)
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	P0() objectivec.IObject
	SetP0(value objectivec.IObject)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)
	Pm() objectivec.IObject
	SetPm(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNLocalContrastNormalizationGradientClass) Alloc() CNNLocalContrastNormalizationGradient {
	rv := objc.Send[CNNLocalContrastNormalizationGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLocalContrastNormalizationGradientClass) New() CNNLocalContrastNormalizationGradient {
	rv := objc.Send[CNNLocalContrastNormalizationGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLocalContrastNormalizationGradient) Init() CNNLocalContrastNormalizationGradient {
	rv := objc.Send[CNNLocalContrastNormalizationGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLocalContrastNormalizationGradient) Autorelease() CNNLocalContrastNormalizationGradient {
	rv := objc.Send[CNNLocalContrastNormalizationGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLocalContrastNormalizationGradient creates a new CNNLocalContrastNormalizationGradient instance.
func NewCNNLocalContrastNormalizationGradient() CNNLocalContrastNormalizationGradient {
	return getCNNLocalContrastNormalizationGradientClass().New()
}





// A gradient local-contrast normalization kernel.


// A gradient local-contrast normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient
type CNNLocalContrastNormalizationGradient struct {
	CNNGradientKernel
}

// CNNLocalContrastNormalizationGradientFrom constructs a [CNNLocalContrastNormalizationGradient] from an unsafe.Pointer.
//
// A gradient local-contrast normalization kernel.
func CNNLocalContrastNormalizationGradientFrom(ptr unsafe.Pointer) CNNLocalContrastNormalizationGradient {
	return CNNLocalContrastNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942467-initwithcoder
func NewCNNLocalContrastNormalizationGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNLocalContrastNormalizationGradient {
	instance := getCNNLocalContrastNormalizationGradientClass().Alloc()
	rv := objc.Send[CNNLocalContrastNormalizationGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942462-initwithdevice
func NewCNNLocalContrastNormalizationGradientWithDeviceKernelWidthKernelHeight(device unsafe.Pointer, kernelWidth uint, kernelHeight uint) CNNLocalContrastNormalizationGradient {
	instance := getCNNLocalContrastNormalizationGradientClass().Alloc()
	rv := objc.Send[CNNLocalContrastNormalizationGradient](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942466-ps
func (c_ CNNLocalContrastNormalizationGradient) Ps() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("ps"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942466-ps
func (c_ CNNLocalContrastNormalizationGradient) SetPs(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942471-alpha
func (c_ CNNLocalContrastNormalizationGradient) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942471-alpha
func (c_ CNNLocalContrastNormalizationGradient) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942472-p0
func (c_ CNNLocalContrastNormalizationGradient) P0() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("p0"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942472-p0
func (c_ CNNLocalContrastNormalizationGradient) SetP0(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setP0:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942482-delta
func (c_ CNNLocalContrastNormalizationGradient) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942482-delta
func (c_ CNNLocalContrastNormalizationGradient) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942484-beta
func (c_ CNNLocalContrastNormalizationGradient) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("beta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942484-beta
func (c_ CNNLocalContrastNormalizationGradient) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942485-pm
func (c_ CNNLocalContrastNormalizationGradient) Pm() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("pm"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradient/2942485-pm
func (c_ CNNLocalContrastNormalizationGradient) SetPm(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPm:"), value)
}







