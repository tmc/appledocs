// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BinaryArithmeticNode] class.
var (
	BinaryArithmeticNodeClass     _BinaryArithmeticNodeClass
	BinaryArithmeticNodeClassOnce sync.Once
)

func getBinaryArithmeticNodeClass() _BinaryArithmeticNodeClass {
	BinaryArithmeticNodeClassOnce.Do(func() {
		BinaryArithmeticNodeClass = _BinaryArithmeticNodeClass{objc.GetClass("MPSNNBinaryArithmeticNode")}
	})
	return BinaryArithmeticNodeClass
}

type _BinaryArithmeticNodeClass struct {
	class objc.Class
}

// An interface definition for the [BinaryArithmeticNode] class.
type IBinaryArithmeticNode interface {
	IFilterNode
	Bias() float32
	SetBias(value float32)
	MaximumValue() float32
	SetMaximumValue(value float32)
	MinimumValue() float32
	SetMinimumValue(value float32)
	PrimaryScale() float32
	SetPrimaryScale(value float32)
	PrimaryStrideInFeatureChannels() int
	SetPrimaryStrideInFeatureChannels(value int)
	PrimaryStrideInPixelsX() int
	SetPrimaryStrideInPixelsX(value int)
	PrimaryStrideInPixelsY() int
	SetPrimaryStrideInPixelsY(value int)
	SecondaryScale() float32
	SetSecondaryScale(value float32)
	SecondaryStrideInFeatureChannels() int
	SetSecondaryStrideInFeatureChannels(value int)
	SecondaryStrideInPixelsX() int
	SetSecondaryStrideInPixelsX(value int)
	SecondaryStrideInPixelsY() int
	SetSecondaryStrideInPixelsY(value int)
}

// Virtual base class for basic arithmetic nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode
type BinaryArithmeticNode struct {
	FilterNode
}

// BinaryArithmeticNodeFrom constructs a [BinaryArithmeticNode] from an unsafe.Pointer.
//
// Virtual base class for basic arithmetic nodes.
func BinaryArithmeticNodeFrom(ptr unsafe.Pointer) BinaryArithmeticNode {
	return BinaryArithmeticNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BinaryArithmeticNodeClass) Alloc() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BinaryArithmeticNodeClass) New() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BinaryArithmeticNode) Init() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BinaryArithmeticNode) Autorelease() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBinaryArithmeticNode creates a new BinaryArithmeticNode instance.
func NewBinaryArithmeticNode() BinaryArithmeticNode {
	return getBinaryArithmeticNodeClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/bias
func (b_ BinaryArithmeticNode) Bias() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("bias"))
	return rv
}


// SetBias sets the value of the bias property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/bias
func (b_ BinaryArithmeticNode) SetBias(value float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBias:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/maximumvalue
func (b_ BinaryArithmeticNode) MaximumValue() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("maximumValue"))
	return rv
}


// SetMaximumValue sets the value of the maximumValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/maximumvalue
func (b_ BinaryArithmeticNode) SetMaximumValue(value float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMaximumValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/minimumvalue
func (b_ BinaryArithmeticNode) MinimumValue() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("minimumValue"))
	return rv
}


// SetMinimumValue sets the value of the minimumValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/minimumvalue
func (b_ BinaryArithmeticNode) SetMinimumValue(value float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMinimumValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/primaryscale
func (b_ BinaryArithmeticNode) PrimaryScale() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("primaryScale"))
	return rv
}


// SetPrimaryScale sets the value of the primaryScale property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/primaryscale
func (b_ BinaryArithmeticNode) SetPrimaryScale(value float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryScale:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/primarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) PrimaryStrideInFeatureChannels() int {
	rv := objc.Send[int](b_.ID, objc.Sel("primaryStrideInFeatureChannels"))
	return rv
}


// SetPrimaryStrideInFeatureChannels sets the value of the primaryStrideInFeatureChannels property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/primarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) SetPrimaryStrideInFeatureChannels(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryStrideInFeatureChannels:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/primarystrideinpixelsx
func (b_ BinaryArithmeticNode) PrimaryStrideInPixelsX() int {
	rv := objc.Send[int](b_.ID, objc.Sel("primaryStrideInPixelsX"))
	return rv
}


// SetPrimaryStrideInPixelsX sets the value of the primaryStrideInPixelsX property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/primarystrideinpixelsx
func (b_ BinaryArithmeticNode) SetPrimaryStrideInPixelsX(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/primarystrideinpixelsy
func (b_ BinaryArithmeticNode) PrimaryStrideInPixelsY() int {
	rv := objc.Send[int](b_.ID, objc.Sel("primaryStrideInPixelsY"))
	return rv
}


// SetPrimaryStrideInPixelsY sets the value of the primaryStrideInPixelsY property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/primarystrideinpixelsy
func (b_ BinaryArithmeticNode) SetPrimaryStrideInPixelsY(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/secondaryscale
func (b_ BinaryArithmeticNode) SecondaryScale() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("secondaryScale"))
	return rv
}


// SetSecondaryScale sets the value of the secondaryScale property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/secondaryscale
func (b_ BinaryArithmeticNode) SetSecondaryScale(value float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryScale:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/secondarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) SecondaryStrideInFeatureChannels() int {
	rv := objc.Send[int](b_.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}


// SetSecondaryStrideInFeatureChannels sets the value of the secondaryStrideInFeatureChannels property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/secondarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) SetSecondaryStrideInFeatureChannels(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/secondarystrideinpixelsx
func (b_ BinaryArithmeticNode) SecondaryStrideInPixelsX() int {
	rv := objc.Send[int](b_.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}


// SetSecondaryStrideInPixelsX sets the value of the secondaryStrideInPixelsX property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/secondarystrideinpixelsx
func (b_ BinaryArithmeticNode) SetSecondaryStrideInPixelsX(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/secondarystrideinpixelsy
func (b_ BinaryArithmeticNode) SecondaryStrideInPixelsY() int {
	rv := objc.Send[int](b_.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}


// SetSecondaryStrideInPixelsY sets the value of the secondaryStrideInPixelsY property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/secondarystrideinpixelsy
func (b_ BinaryArithmeticNode) SetSecondaryStrideInPixelsY(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}



