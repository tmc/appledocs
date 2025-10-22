// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	objectivec.IObject
	Alpha() float32
	SetAlpha(value float32)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient
type GramMatrixCalculationGradient struct {
	objectivec.Object
}

// GramMatrixCalculationGradientFrom constructs a [GramMatrixCalculationGradient] from an unsafe.Pointer.
func GramMatrixCalculationGradientFrom(ptr unsafe.Pointer) GramMatrixCalculationGradient {
	return GramMatrixCalculationGradient{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GramMatrixCalculationGradientClass) Alloc() GramMatrixCalculationGradient {
	rv := objc.Send[GramMatrixCalculationGradient](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient/init(coder:device:)
func NewGramMatrixCalculationGradientWithCoderDevice(aDecoder foundation.ICoder, device objectivec.IObject) GramMatrixCalculationGradient {
	instance := getGramMatrixCalculationGradientClass().Alloc()
	rv := objc.Send[GramMatrixCalculationGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient/alpha
func (g_ GramMatrixCalculationGradient) Alpha() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("alpha"))
	return rv
}


// SetAlpha sets the value of the alpha property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient/alpha
func (g_ GramMatrixCalculationGradient) SetAlpha(value float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAlpha:"), value)
}


