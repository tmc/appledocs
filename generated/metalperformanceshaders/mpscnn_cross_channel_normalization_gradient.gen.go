// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNCrossChannelNormalizationGradient] class.
var (
	CNNCrossChannelNormalizationGradientClass     _CNNCrossChannelNormalizationGradientClass
	CNNCrossChannelNormalizationGradientClassOnce sync.Once
)

func getCNNCrossChannelNormalizationGradientClass() _CNNCrossChannelNormalizationGradientClass {
	CNNCrossChannelNormalizationGradientClassOnce.Do(func() {
		CNNCrossChannelNormalizationGradientClass = _CNNCrossChannelNormalizationGradientClass{objc.GetClass("MPSCNNCrossChannelNormalizationGradient")}
	})
	return CNNCrossChannelNormalizationGradientClass
}

type _CNNCrossChannelNormalizationGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNCrossChannelNormalizationGradient] class.
type ICNNCrossChannelNormalizationGradient interface {
	ICNNGradientKernel
	

	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	KernelSize() objectivec.IObject
	SetKernelSize(value objectivec.IObject)
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNCrossChannelNormalizationGradientClass) Alloc() CNNCrossChannelNormalizationGradient {
	rv := objc.Send[CNNCrossChannelNormalizationGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNCrossChannelNormalizationGradientClass) New() CNNCrossChannelNormalizationGradient {
	rv := objc.Send[CNNCrossChannelNormalizationGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNCrossChannelNormalizationGradient) Init() CNNCrossChannelNormalizationGradient {
	rv := objc.Send[CNNCrossChannelNormalizationGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNCrossChannelNormalizationGradient) Autorelease() CNNCrossChannelNormalizationGradient {
	rv := objc.Send[CNNCrossChannelNormalizationGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNCrossChannelNormalizationGradient creates a new CNNCrossChannelNormalizationGradient instance.
func NewCNNCrossChannelNormalizationGradient() CNNCrossChannelNormalizationGradient {
	return getCNNCrossChannelNormalizationGradientClass().New()
}





// A gradient normalization kernel applied across feature channels.


// A gradient normalization kernel applied across feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradient
type CNNCrossChannelNormalizationGradient struct {
	CNNGradientKernel
}

// CNNCrossChannelNormalizationGradientFrom constructs a [CNNCrossChannelNormalizationGradient] from an unsafe.Pointer.
//
// A gradient normalization kernel applied across feature channels.
func CNNCrossChannelNormalizationGradientFrom(ptr unsafe.Pointer) CNNCrossChannelNormalizationGradient {
	return CNNCrossChannelNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942476-initwithcoder
func NewCNNCrossChannelNormalizationGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNCrossChannelNormalizationGradient {
	instance := getCNNCrossChannelNormalizationGradientClass().Alloc()
	rv := objc.Send[CNNCrossChannelNormalizationGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942463-initwithdevice
func NewCNNCrossChannelNormalizationGradientWithDeviceKernelSize(device unsafe.Pointer, kernelSize uint) CNNCrossChannelNormalizationGradient {
	instance := getCNNCrossChannelNormalizationGradientClass().Alloc()
	rv := objc.Send[CNNCrossChannelNormalizationGradient](instance.ID, objc.Sel("initWithDevice:kernelSize:"), device, kernelSize)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942464-alpha
func (c_ CNNCrossChannelNormalizationGradient) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942464-alpha
func (c_ CNNCrossChannelNormalizationGradient) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942465-delta
func (c_ CNNCrossChannelNormalizationGradient) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942465-delta
func (c_ CNNCrossChannelNormalizationGradient) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942468-kernelsize
func (c_ CNNCrossChannelNormalizationGradient) KernelSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942468-kernelsize
func (c_ CNNCrossChannelNormalizationGradient) SetKernelSize(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942477-beta
func (c_ CNNCrossChannelNormalizationGradient) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("beta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942477-beta
func (c_ CNNCrossChannelNormalizationGradient) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}







