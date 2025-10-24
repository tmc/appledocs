// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ForwardLossNode] class.
var (
	ForwardLossNodeClass     _ForwardLossNodeClass
	ForwardLossNodeClassOnce sync.Once
)

func getForwardLossNodeClass() _ForwardLossNodeClass {
	ForwardLossNodeClassOnce.Do(func() {
		ForwardLossNodeClass = _ForwardLossNodeClass{objc.GetClass("MPSNNForwardLossNode")}
	})
	return ForwardLossNodeClass
}

type _ForwardLossNodeClass struct {
	class objc.Class
}

// An interface definition for the [ForwardLossNode] class.
type IForwardLossNode interface {
	IFilterNode
	// properties:
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
	PropertyCallBack() LossCallback /* not a class type */
	SetPropertyCallBack(value LossCallback /* not a class type */)
	ReduceAcrossBatch() bool
	SetReduceAcrossBatch(value bool)
	ReductionType() CNNReductionType /* not a class type */
	SetReductionType(value CNNReductionType /* not a class type */)
	Weight() float32
	SetWeight(value float32)
	// methods:
	GradientFiltersWithSources(sourceGradient []IImageNode) []ILossGradientNode
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode
type ForwardLossNode struct {
	FilterNode
}

// ForwardLossNodeFrom constructs a [ForwardLossNode] from an unsafe.Pointer.
func ForwardLossNodeFrom(ptr unsafe.Pointer) ForwardLossNode {
	return ForwardLossNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _ForwardLossNodeClass) Alloc() ForwardLossNode {
	rv := objc.Send[ForwardLossNode](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _ForwardLossNodeClass) New() ForwardLossNode {
	rv := objc.Send[ForwardLossNode](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ ForwardLossNode) Init() ForwardLossNode {
	rv := objc.Send[ForwardLossNode](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ ForwardLossNode) Autorelease() ForwardLossNode {
	rv := objc.Send[ForwardLossNode](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewForwardLossNode creates a new ForwardLossNode instance.
func NewForwardLossNode() ForwardLossNode {
	return getForwardLossNodeClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/gradientFilters(withSources:)
func (f_ ForwardLossNode) GradientFiltersWithSources(sourceGradient []IImageNode) []ILossGradientNode {
	rv := objc.Send[[]LossGradientNode](f_.ID, objc.Sel("gradientFiltersWithSources:"), sourceGradient)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/delta
func (f_ ForwardLossNode) Delta() float32 {
	rv := objc.Send[float32](f_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/delta
func (f_ ForwardLossNode) SetDelta(value float32) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/epsilon
func (f_ ForwardLossNode) Epsilon() float32 {
	rv := objc.Send[float32](f_.ID, objc.Sel("epsilon"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/epsilon
func (f_ ForwardLossNode) SetEpsilon(value float32) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEpsilon:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/labelsmoothing
func (f_ ForwardLossNode) LabelSmoothing() float32 {
	rv := objc.Send[float32](f_.ID, objc.Sel("labelSmoothing"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/labelsmoothing
func (f_ ForwardLossNode) SetLabelSmoothing(value float32) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLabelSmoothing:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/losstype
func (f_ ForwardLossNode) LossType() CNNLossType /* not a class type */ {
	rv := objc.Send[CNNLossType](f_.ID, objc.Sel("lossType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/losstype
func (f_ ForwardLossNode) SetLossType(value CNNLossType /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLossType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/numberofclasses
func (f_ ForwardLossNode) NumberOfClasses() int {
	rv := objc.Send[int](f_.ID, objc.Sel("numberOfClasses"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/numberofclasses
func (f_ ForwardLossNode) SetNumberOfClasses(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNumberOfClasses:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/propertycallback
func (f_ ForwardLossNode) PropertyCallBack() LossCallback /* not a class type */ {
	rv := objc.Send[LossCallback](f_.ID, objc.Sel("propertyCallBack"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/propertycallback
func (f_ ForwardLossNode) SetPropertyCallBack(value LossCallback /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPropertyCallBack:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/reduceacrossbatch
func (f_ ForwardLossNode) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/reduceacrossbatch
func (f_ ForwardLossNode) SetReduceAcrossBatch(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/reductiontype
func (f_ ForwardLossNode) ReductionType() CNNReductionType /* not a class type */ {
	rv := objc.Send[CNNReductionType](f_.ID, objc.Sel("reductionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/reductiontype
func (f_ ForwardLossNode) SetReductionType(value CNNReductionType /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReductionType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/weight
func (f_ ForwardLossNode) Weight() float32 {
	rv := objc.Send[float32](f_.ID, objc.Sel("weight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/weight
func (f_ ForwardLossNode) SetWeight(value float32) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWeight:"), value)
}



