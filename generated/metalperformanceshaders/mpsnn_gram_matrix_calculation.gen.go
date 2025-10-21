// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GramMatrixCalculation] class.
var (
	GramMatrixCalculationClass     _GramMatrixCalculationClass
	GramMatrixCalculationClassOnce sync.Once
)

func getGramMatrixCalculationClass() _GramMatrixCalculationClass {
	GramMatrixCalculationClassOnce.Do(func() {
		GramMatrixCalculationClass = _GramMatrixCalculationClass{objc.GetClass("MPSNNGramMatrixCalculation")}
	})
	return GramMatrixCalculationClass
}

type _GramMatrixCalculationClass struct {
	class objc.Class
}

// An interface definition for the [GramMatrixCalculation] class.
type IGramMatrixCalculation interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculation
type GramMatrixCalculation struct {
	objectivec.Object
}

// GramMatrixCalculationFrom constructs a [GramMatrixCalculation] from an unsafe.Pointer.
func GramMatrixCalculationFrom(ptr unsafe.Pointer) GramMatrixCalculation {
	return GramMatrixCalculation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GramMatrixCalculationClass) Alloc() GramMatrixCalculation {
	rv := objc.Send[GramMatrixCalculation](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GramMatrixCalculationClass) New() GramMatrixCalculation {
	rv := objc.Send[GramMatrixCalculation](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GramMatrixCalculation) Init() GramMatrixCalculation {
	rv := objc.Send[GramMatrixCalculation](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GramMatrixCalculation) Autorelease() GramMatrixCalculation {
	rv := objc.Send[GramMatrixCalculation](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGramMatrixCalculation creates a new GramMatrixCalculation instance.
func NewGramMatrixCalculation() GramMatrixCalculation {
	return getGramMatrixCalculationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculation/init(coder:device:)
func NewGramMatrixCalculationWithCoderDevice(aDecoder foundation.ICoder, device objectivec.IObject) GramMatrixCalculation {
	instance := getGramMatrixCalculationClass().Alloc()
	rv := objc.Send[GramMatrixCalculation](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculation/alpha
func (g_ GramMatrixCalculation) Alpha() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("alpha"))
	return rv
}


// SetAlpha sets the value of the alpha property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculation/alpha
func (g_ GramMatrixCalculation) SetAlpha(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAlpha:"), value)
}


