// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNUpsamplingBilinearNode] class.
var (
	CNNUpsamplingBilinearNodeClass     _CNNUpsamplingBilinearNodeClass
	CNNUpsamplingBilinearNodeClassOnce sync.Once
)

func getCNNUpsamplingBilinearNodeClass() _CNNUpsamplingBilinearNodeClass {
	CNNUpsamplingBilinearNodeClassOnce.Do(func() {
		CNNUpsamplingBilinearNodeClass = _CNNUpsamplingBilinearNodeClass{objc.GetClass("MPSCNNUpsamplingBilinearNode")}
	})
	return CNNUpsamplingBilinearNodeClass
}

type _CNNUpsamplingBilinearNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNUpsamplingBilinearNode] class.
type ICNNUpsamplingBilinearNode interface {
	IFilterNode
	

	// properties:
	ScaleFactorY() objectivec.IObject
	SetScaleFactorY(value objectivec.IObject)
	ScaleFactorX() objectivec.IObject
	SetScaleFactorX(value objectivec.IObject)
	AlignCorners() objectivec.IObject
	SetAlignCorners(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingBilinearNodeClass) Alloc() CNNUpsamplingBilinearNode {
	rv := objc.Send[CNNUpsamplingBilinearNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingBilinearNodeClass) New() CNNUpsamplingBilinearNode {
	rv := objc.Send[CNNUpsamplingBilinearNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingBilinearNode) Init() CNNUpsamplingBilinearNode {
	rv := objc.Send[CNNUpsamplingBilinearNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingBilinearNode) Autorelease() CNNUpsamplingBilinearNode {
	rv := objc.Send[CNNUpsamplingBilinearNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingBilinearNode creates a new CNNUpsamplingBilinearNode instance.
func NewCNNUpsamplingBilinearNode() CNNUpsamplingBilinearNode {
	return getCNNUpsamplingBilinearNodeClass().New()
}





// A representation of a bilinear spatial upsampling filter.


// A representation of a bilinear spatial upsampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode
type CNNUpsamplingBilinearNode struct {
	FilterNode
}

// CNNUpsamplingBilinearNodeFrom constructs a [CNNUpsamplingBilinearNode] from an unsafe.Pointer.
//
// A representation of a bilinear spatial upsampling filter.
func CNNUpsamplingBilinearNodeFrom(ptr unsafe.Pointer) CNNUpsamplingBilinearNode {
	return CNNUpsamplingBilinearNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875152-initwithsource
func NewCNNUpsamplingBilinearNodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IImageNode, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinearNode {
	instance := getCNNUpsamplingBilinearNodeClass().Alloc()
	rv := objc.Send[CNNUpsamplingBilinearNode](instance.ID, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2966688-initwithsource
func NewCNNUpsamplingBilinearNodeWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(sourceNode IImageNode, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) CNNUpsamplingBilinearNode {
	instance := getCNNUpsamplingBilinearNodeClass().Alloc()
	rv := objc.Send[CNNUpsamplingBilinearNode](instance.ID, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:alignCorners:"), sourceNode, integerScaleFactorX, integerScaleFactorY, alignCorners)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875987-nodewithsource
func (cc _CNNUpsamplingBilinearNodeClass) NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IImageNode, integerScaleFactorX uint, integerScaleFactorY uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2966689-nodewithsource
func (cc _CNNUpsamplingBilinearNodeClass) NodeWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(sourceNode IImageNode, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:integerScaleFactorX:integerScaleFactorY:alignCorners:"), sourceNode, integerScaleFactorX, integerScaleFactorY, alignCorners)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875150-scalefactory
func (c_ CNNUpsamplingBilinearNode) ScaleFactorY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875150-scalefactory
func (c_ CNNUpsamplingBilinearNode) SetScaleFactorY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875153-scalefactorx
func (c_ CNNUpsamplingBilinearNode) ScaleFactorX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875153-scalefactorx
func (c_ CNNUpsamplingBilinearNode) SetScaleFactorX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2966687-aligncorners
func (c_ CNNUpsamplingBilinearNode) AlignCorners() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alignCorners"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2966687-aligncorners
func (c_ CNNUpsamplingBilinearNode) SetAlignCorners(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlignCorners:"), value)
}







