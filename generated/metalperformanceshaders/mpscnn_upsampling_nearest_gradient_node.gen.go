// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNUpsamplingNearestGradientNode] class.
var (
	CNNUpsamplingNearestGradientNodeClass     _CNNUpsamplingNearestGradientNodeClass
	CNNUpsamplingNearestGradientNodeClassOnce sync.Once
)

func getCNNUpsamplingNearestGradientNodeClass() _CNNUpsamplingNearestGradientNodeClass {
	CNNUpsamplingNearestGradientNodeClassOnce.Do(func() {
		CNNUpsamplingNearestGradientNodeClass = _CNNUpsamplingNearestGradientNodeClass{objc.GetClass("MPSCNNUpsamplingNearestGradientNode")}
	})
	return CNNUpsamplingNearestGradientNodeClass
}

type _CNNUpsamplingNearestGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNUpsamplingNearestGradientNode] class.
type ICNNUpsamplingNearestGradientNode interface {
	IGradientFilterNode
	

	// properties:
	ScaleFactorX() objectivec.IObject
	SetScaleFactorX(value objectivec.IObject)
	ScaleFactorY() objectivec.IObject
	SetScaleFactorY(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingNearestGradientNodeClass) Alloc() CNNUpsamplingNearestGradientNode {
	rv := objc.Send[CNNUpsamplingNearestGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingNearestGradientNodeClass) New() CNNUpsamplingNearestGradientNode {
	rv := objc.Send[CNNUpsamplingNearestGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingNearestGradientNode) Init() CNNUpsamplingNearestGradientNode {
	rv := objc.Send[CNNUpsamplingNearestGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingNearestGradientNode) Autorelease() CNNUpsamplingNearestGradientNode {
	rv := objc.Send[CNNUpsamplingNearestGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingNearestGradientNode creates a new CNNUpsamplingNearestGradientNode instance.
func NewCNNUpsamplingNearestGradientNode() CNNUpsamplingNearestGradientNode {
	return getCNNUpsamplingNearestGradientNodeClass().New()
}





// A representation of a gradient nearest spatial upsampling filter.


// A representation of a gradient nearest spatial upsampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradientNode
type CNNUpsamplingNearestGradientNode struct {
	GradientFilterNode
}

// CNNUpsamplingNearestGradientNodeFrom constructs a [CNNUpsamplingNearestGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient nearest spatial upsampling filter.
func CNNUpsamplingNearestGradientNodeFrom(ptr unsafe.Pointer) CNNUpsamplingNearestGradientNode {
	return CNNUpsamplingNearestGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2947983-initwithsourcegradient
func NewCNNUpsamplingNearestGradientNodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, scaleFactorX float64, scaleFactorY float64) CNNUpsamplingNearestGradientNode {
	instance := getCNNUpsamplingNearestGradientNodeClass().Alloc()
	rv := objc.Send[CNNUpsamplingNearestGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948021-nodewithsourcegradient
func (cc _CNNUpsamplingNearestGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, scaleFactorX float64, scaleFactorY float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948024-scalefactorx
func (c_ CNNUpsamplingNearestGradientNode) ScaleFactorX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948024-scalefactorx
func (c_ CNNUpsamplingNearestGradientNode) SetScaleFactorX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948035-scalefactory
func (c_ CNNUpsamplingNearestGradientNode) ScaleFactorY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948035-scalefactory
func (c_ CNNUpsamplingNearestGradientNode) SetScaleFactorY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorY:"), value)
}







