// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [GramMatrixCalculationGradient] class.
var (
	GramMatrixCalculationGradientClass     _GramMatrixCalculationGradientClass
	GramMatrixCalculationGradientClassOnce sync.Once
)

func getGramMatrixCalculationGradientClass() _GramMatrixCalculationGradientClass {
	GramMatrixCalculationGradientClassOnce.Do(func() {
		GramMatrixCalculationGradientClass = _GramMatrixCalculationGradientClass{objc.GetClass("MPSNNGramMatrixCalculationGradient")}
	})
	return GramMatrixCalculationGradientClass
}

type _GramMatrixCalculationGradientClass struct {
	class objc.Class
}





// An interface definition for the [GramMatrixCalculationGradient] class.
type IGramMatrixCalculationGradient interface {
	ICNNGradientKernel
	

	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (gc _GramMatrixCalculationGradientClass) Alloc() GramMatrixCalculationGradient {
	rv := objc.Send[GramMatrixCalculationGradient](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GramMatrixCalculationGradientClass) New() GramMatrixCalculationGradient {
	rv := objc.Send[GramMatrixCalculationGradient](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GramMatrixCalculationGradient) Init() GramMatrixCalculationGradient {
	rv := objc.Send[GramMatrixCalculationGradient](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GramMatrixCalculationGradient) Autorelease() GramMatrixCalculationGradient {
	rv := objc.Send[GramMatrixCalculationGradient](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGramMatrixCalculationGradient creates a new GramMatrixCalculationGradient instance.
func NewGramMatrixCalculationGradient() GramMatrixCalculationGradient {
	return getGramMatrixCalculationGradientClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient
type GramMatrixCalculationGradient struct {
	CNNGradientKernel
}

// GramMatrixCalculationGradientFrom constructs a [GramMatrixCalculationGradient] from an unsafe.Pointer.
func GramMatrixCalculationGradientFrom(ptr unsafe.Pointer) GramMatrixCalculationGradient {
	return GramMatrixCalculationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradient/3114083-initwithcoder
func NewGramMatrixCalculationGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) GramMatrixCalculationGradient {
	instance := getGramMatrixCalculationGradientClass().Alloc()
	rv := objc.Send[GramMatrixCalculationGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradient/3114084-initwithdevice
func NewGramMatrixCalculationGradientWithDevice(device unsafe.Pointer) GramMatrixCalculationGradient {
	instance := getGramMatrixCalculationGradientClass().Alloc()
	rv := objc.Send[GramMatrixCalculationGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradient/3114085-initwithdevice
func NewGramMatrixCalculationGradientWithDeviceAlpha(device unsafe.Pointer, alpha float32) GramMatrixCalculationGradient {
	instance := getGramMatrixCalculationGradientClass().Alloc()
	rv := objc.Send[GramMatrixCalculationGradient](instance.ID, objc.Sel("initWithDevice:alpha:"), device, alpha)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradient/3114082-alpha
func (g_ GramMatrixCalculationGradient) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradient/3114082-alpha
func (g_ GramMatrixCalculationGradient) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAlpha:"), value)
}







