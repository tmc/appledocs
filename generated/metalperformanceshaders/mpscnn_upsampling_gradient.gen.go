// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNUpsamplingGradient] class.
var (
	CNNUpsamplingGradientClass     _CNNUpsamplingGradientClass
	CNNUpsamplingGradientClassOnce sync.Once
)

func getCNNUpsamplingGradientClass() _CNNUpsamplingGradientClass {
	CNNUpsamplingGradientClassOnce.Do(func() {
		CNNUpsamplingGradientClass = _CNNUpsamplingGradientClass{objc.GetClass("MPSCNNUpsamplingGradient")}
	})
	return CNNUpsamplingGradientClass
}

type _CNNUpsamplingGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNUpsamplingGradient] class.
type ICNNUpsamplingGradient interface {
	ICNNGradientKernel
	

	// properties:
	ScaleFactorY() objectivec.IObject
	SetScaleFactorY(value objectivec.IObject)
	ScaleFactorX() objectivec.IObject
	SetScaleFactorX(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingGradientClass) Alloc() CNNUpsamplingGradient {
	rv := objc.Send[CNNUpsamplingGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingGradientClass) New() CNNUpsamplingGradient {
	rv := objc.Send[CNNUpsamplingGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingGradient) Init() CNNUpsamplingGradient {
	rv := objc.Send[CNNUpsamplingGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingGradient) Autorelease() CNNUpsamplingGradient {
	rv := objc.Send[CNNUpsamplingGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingGradient creates a new CNNUpsamplingGradient instance.
func NewCNNUpsamplingGradient() CNNUpsamplingGradient {
	return getCNNUpsamplingGradientClass().New()
}





// A gradient filter that upsamples an existing Metal Performance Shaders image.


// A gradient filter that upsamples an existing Metal Performance Shaders image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingGradient
type CNNUpsamplingGradient struct {
	CNNGradientKernel
}

// CNNUpsamplingGradientFrom constructs a [CNNUpsamplingGradient] from an unsafe.Pointer.
//
// A gradient filter that upsamples an existing Metal Performance Shaders image.
func CNNUpsamplingGradientFrom(ptr unsafe.Pointer) CNNUpsamplingGradient {
	return CNNUpsamplingGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplinggradient/2942628-scalefactory
func (c_ CNNUpsamplingGradient) ScaleFactorY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplinggradient/2942628-scalefactory
func (c_ CNNUpsamplingGradient) SetScaleFactorY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplinggradient/2942630-scalefactorx
func (c_ CNNUpsamplingGradient) ScaleFactorX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplinggradient/2942630-scalefactorx
func (c_ CNNUpsamplingGradient) SetScaleFactorX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorX:"), value)
}








