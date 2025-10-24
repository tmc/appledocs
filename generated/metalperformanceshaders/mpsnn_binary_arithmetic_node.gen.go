// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	Bias() objectivec.IObject
	SetBias(value objectivec.IObject)
	PrimaryScale() objectivec.IObject
	SetPrimaryScale(value objectivec.IObject)
	MinimumValue() objectivec.IObject
	SetMinimumValue(value objectivec.IObject)
	SecondaryStrideInPixelsX() objectivec.IObject
	SetSecondaryStrideInPixelsX(value objectivec.IObject)
	PrimaryStrideInPixelsX() objectivec.IObject
	SetPrimaryStrideInPixelsX(value objectivec.IObject)
	SecondaryStrideInFeatureChannels() objectivec.IObject
	SetSecondaryStrideInFeatureChannels(value objectivec.IObject)
	SecondaryScale() objectivec.IObject
	SetSecondaryScale(value objectivec.IObject)
	MaximumValue() objectivec.IObject
	SetMaximumValue(value objectivec.IObject)
	PrimaryStrideInFeatureChannels() objectivec.IObject
	SetPrimaryStrideInFeatureChannels(value objectivec.IObject)
	SecondaryStrideInPixelsY() objectivec.IObject
	SetSecondaryStrideInPixelsY(value objectivec.IObject)
	PrimaryStrideInPixelsY() objectivec.IObject
	SetPrimaryStrideInPixelsY(value objectivec.IObject)


	

	// methods:
	GradientFilters()
	GradientFiltersWithSources(gradientImages unsafe.Pointer) unsafe.Pointer
	GradientClass()


}





// Alloc allocates a new instance without initialization.
func (bc _BinaryArithmeticNodeClass) Alloc() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// Virtual base class for basic arithmetic nodes.


// Virtual base class for basic arithmetic nodes.
//
// [Full Topic]
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890825-initwithleftsource
func NewBinaryArithmeticNodeWithLeftSourceRightSource(left IImageNode, right IImageNode) BinaryArithmeticNode {
	instance := getBinaryArithmeticNodeClass().Alloc()
	rv := objc.Send[BinaryArithmeticNode](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), left, right)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890820-initwithsources
func NewBinaryArithmeticNodeWithSources(sourceNodes unsafe.Pointer) BinaryArithmeticNode {
	instance := getBinaryArithmeticNodeClass().Alloc()
	rv := objc.Send[BinaryArithmeticNode](instance.ID, objc.Sel("initWithSources:"), sourceNodes)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890829-nodewithsources
func (bc _BinaryArithmeticNodeClass) NodeWithSources(sourceNodes unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("nodeWithSources:"), sourceNodes)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890830-nodewithleftsource
func (bc _BinaryArithmeticNodeClass) NodeWithLeftSourceRightSource(left IImageNode, right IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("nodeWithLeftSource:rightSource:"), left, right)
	return rv
}












// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952967-gradientfilters
func (b_ BinaryArithmeticNode) GradientFilters() {
	objc.Send[objc.ID](b_.ID, objc.Sel("gradientFilters"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952967-gradientfilterswithsources
func (b_ BinaryArithmeticNode) GradientFiltersWithSources(gradientImages unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("gradientFiltersWithSources:"), gradientImages)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952978-gradientclass
func (b_ BinaryArithmeticNode) GradientClass() {
	objc.Send[objc.ID](b_.ID, objc.Sel("gradientClass"))
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952964-bias
func (b_ BinaryArithmeticNode) Bias() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("bias"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952964-bias
func (b_ BinaryArithmeticNode) SetBias(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBias:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952966-primaryscale
func (b_ BinaryArithmeticNode) PrimaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("primaryScale"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952966-primaryscale
func (b_ BinaryArithmeticNode) SetPrimaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryScale:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952970-minimumvalue
func (b_ BinaryArithmeticNode) MinimumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("minimumValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952970-minimumvalue
func (b_ BinaryArithmeticNode) SetMinimumValue(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMinimumValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952972-secondarystrideinpixelsx
func (b_ BinaryArithmeticNode) SecondaryStrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952972-secondarystrideinpixelsx
func (b_ BinaryArithmeticNode) SetSecondaryStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952973-primarystrideinpixelsx
func (b_ BinaryArithmeticNode) PrimaryStrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("primaryStrideInPixelsX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952973-primarystrideinpixelsx
func (b_ BinaryArithmeticNode) SetPrimaryStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952974-secondarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) SecondaryStrideInFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952974-secondarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) SetSecondaryStrideInFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952976-secondaryscale
func (b_ BinaryArithmeticNode) SecondaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("secondaryScale"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952976-secondaryscale
func (b_ BinaryArithmeticNode) SetSecondaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryScale:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952979-maximumvalue
func (b_ BinaryArithmeticNode) MaximumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("maximumValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952979-maximumvalue
func (b_ BinaryArithmeticNode) SetMaximumValue(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMaximumValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952983-primarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) PrimaryStrideInFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("primaryStrideInFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952983-primarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) SetPrimaryStrideInFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryStrideInFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952985-secondarystrideinpixelsy
func (b_ BinaryArithmeticNode) SecondaryStrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952985-secondarystrideinpixelsy
func (b_ BinaryArithmeticNode) SetSecondaryStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952996-primarystrideinpixelsy
func (b_ BinaryArithmeticNode) PrimaryStrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("primaryStrideInPixelsY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952996-primarystrideinpixelsy
func (b_ BinaryArithmeticNode) SetPrimaryStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}







