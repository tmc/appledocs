// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNInstanceNormalizationGradient] class.
var (
	CNNInstanceNormalizationGradientClass     _CNNInstanceNormalizationGradientClass
	CNNInstanceNormalizationGradientClassOnce sync.Once
)

func getCNNInstanceNormalizationGradientClass() _CNNInstanceNormalizationGradientClass {
	CNNInstanceNormalizationGradientClassOnce.Do(func() {
		CNNInstanceNormalizationGradientClass = _CNNInstanceNormalizationGradientClass{objc.GetClass("MPSCNNInstanceNormalizationGradient")}
	})
	return CNNInstanceNormalizationGradientClass
}

type _CNNInstanceNormalizationGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNInstanceNormalizationGradient] class.
type ICNNInstanceNormalizationGradient interface {
	ICNNGradientKernel
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNInstanceNormalizationGradientClass) Alloc() CNNInstanceNormalizationGradient {
	rv := objc.Send[CNNInstanceNormalizationGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNInstanceNormalizationGradientClass) New() CNNInstanceNormalizationGradient {
	rv := objc.Send[CNNInstanceNormalizationGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNInstanceNormalizationGradient) Init() CNNInstanceNormalizationGradient {
	rv := objc.Send[CNNInstanceNormalizationGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNInstanceNormalizationGradient) Autorelease() CNNInstanceNormalizationGradient {
	rv := objc.Send[CNNInstanceNormalizationGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNInstanceNormalizationGradient creates a new CNNInstanceNormalizationGradient instance.
func NewCNNInstanceNormalizationGradient() CNNInstanceNormalizationGradient {
	return getCNNInstanceNormalizationGradientClass().New()
}





// A gradient instance normalization kernel.


// A gradient instance normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradient
type CNNInstanceNormalizationGradient struct {
	CNNGradientKernel
}

// CNNInstanceNormalizationGradientFrom constructs a [CNNInstanceNormalizationGradient] from an unsafe.Pointer.
//
// A gradient instance normalization kernel.
func CNNInstanceNormalizationGradientFrom(ptr unsafe.Pointer) CNNInstanceNormalizationGradient {
	return CNNInstanceNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}































