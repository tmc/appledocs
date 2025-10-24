// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LossGradient] class.
var (
	LossGradientClass     _LossGradientClass
	LossGradientClassOnce sync.Once
)

func getLossGradientClass() _LossGradientClass {
	LossGradientClassOnce.Do(func() {
		LossGradientClass = _LossGradientClass{objc.GetClass("MPSNNLossGradient")}
	})
	return LossGradientClass
}

type _LossGradientClass struct {
	class objc.Class
}

// An interface definition for the [LossGradient] class.
type ILossGradient interface {
	ICNNBinaryKernel
	// properties:
	ComputeLabelGradients() bool
	SetComputeLabelGradients(value bool)
	Delta() float32
	SetDelta(value float32)
	Epsilon() float32
	SetEpsilon(value float32)
	LabelSmoothing() float32
	SetLabelSmoothing(value float32)
	LossType() CNNLossType /* not a class type */
	SetLossType(value CNNLossType /* not a class type */)
	NumberOfClasses() int
	SetNumberOfClasses(value int)
	ReduceAcrossBatch() bool
	SetReduceAcrossBatch(value bool)
	ReductionType() CNNReductionType /* not a class type */
	SetReductionType(value CNNReductionType /* not a class type */)
	Weight() float32
	SetWeight(value float32)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient
type LossGradient struct {
	CNNBinaryKernel
}

// LossGradientFrom constructs a [LossGradient] from an unsafe.Pointer.
func LossGradientFrom(ptr unsafe.Pointer) LossGradient {
	return LossGradient{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (lc _LossGradientClass) Alloc() LossGradient {
	rv := objc.Send[LossGradient](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LossGradientClass) New() LossGradient {
	rv := objc.Send[LossGradient](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LossGradient) Init() LossGradient {
	rv := objc.Send[LossGradient](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LossGradient) Autorelease() LossGradient {
	rv := objc.Send[LossGradient](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLossGradient creates a new LossGradient instance.
func NewLossGradient() LossGradient {
	return getLossGradientClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/init(device:lossDescriptor:)
func NewLossGradientWithDeviceLossDescriptor(device objectivec.IObject, lossDescriptor CNNLossDescriptor /* not a class type */) LossGradient {
	instance := getLossGradientClass().Alloc()
	rv := objc.Send[LossGradient](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/computelabelgradients
func (l_ LossGradient) ComputeLabelGradients() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("computeLabelGradients"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/computelabelgradients
func (l_ LossGradient) SetComputeLabelGradients(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setComputeLabelGradients:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/delta
func (l_ LossGradient) Delta() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/delta
func (l_ LossGradient) SetDelta(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/epsilon
func (l_ LossGradient) Epsilon() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("epsilon"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/epsilon
func (l_ LossGradient) SetEpsilon(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEpsilon:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/labelsmoothing
func (l_ LossGradient) LabelSmoothing() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("labelSmoothing"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/labelsmoothing
func (l_ LossGradient) SetLabelSmoothing(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLabelSmoothing:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/losstype
func (l_ LossGradient) LossType() CNNLossType /* not a class type */ {
	rv := objc.Send[CNNLossType](l_.ID, objc.Sel("lossType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/losstype
func (l_ LossGradient) SetLossType(value CNNLossType /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLossType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/numberofclasses
func (l_ LossGradient) NumberOfClasses() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfClasses"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/numberofclasses
func (l_ LossGradient) SetNumberOfClasses(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfClasses:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/reduceacrossbatch
func (l_ LossGradient) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/reduceacrossbatch
func (l_ LossGradient) SetReduceAcrossBatch(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/reductiontype
func (l_ LossGradient) ReductionType() CNNReductionType /* not a class type */ {
	rv := objc.Send[CNNReductionType](l_.ID, objc.Sel("reductionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/reductiontype
func (l_ LossGradient) SetReductionType(value CNNReductionType /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReductionType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/weight
func (l_ LossGradient) Weight() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("weight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/weight
func (l_ LossGradient) SetWeight(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWeight:"), value)
}


