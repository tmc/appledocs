// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNUpsamplingBilinearGradientNode] class.
var (
	CNNUpsamplingBilinearGradientNodeClass     _CNNUpsamplingBilinearGradientNodeClass
	CNNUpsamplingBilinearGradientNodeClassOnce sync.Once
)

func getCNNUpsamplingBilinearGradientNodeClass() _CNNUpsamplingBilinearGradientNodeClass {
	CNNUpsamplingBilinearGradientNodeClassOnce.Do(func() {
		CNNUpsamplingBilinearGradientNodeClass = _CNNUpsamplingBilinearGradientNodeClass{objc.GetClass("MPSCNNUpsamplingBilinearGradientNode")}
	})
	return CNNUpsamplingBilinearGradientNodeClass
}

type _CNNUpsamplingBilinearGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNUpsamplingBilinearGradientNode] class.
type ICNNUpsamplingBilinearGradientNode interface {
	IGradientFilterNode
	

	// properties:
	ScaleFactorX() objectivec.IObject
	SetScaleFactorX(value objectivec.IObject)
	ScaleFactorY() objectivec.IObject
	SetScaleFactorY(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingBilinearGradientNodeClass) Alloc() CNNUpsamplingBilinearGradientNode {
	rv := objc.Send[CNNUpsamplingBilinearGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingBilinearGradientNodeClass) New() CNNUpsamplingBilinearGradientNode {
	rv := objc.Send[CNNUpsamplingBilinearGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingBilinearGradientNode) Init() CNNUpsamplingBilinearGradientNode {
	rv := objc.Send[CNNUpsamplingBilinearGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingBilinearGradientNode) Autorelease() CNNUpsamplingBilinearGradientNode {
	rv := objc.Send[CNNUpsamplingBilinearGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingBilinearGradientNode creates a new CNNUpsamplingBilinearGradientNode instance.
func NewCNNUpsamplingBilinearGradientNode() CNNUpsamplingBilinearGradientNode {
	return getCNNUpsamplingBilinearGradientNodeClass().New()
}





// A representation of a gradient bilinear spatial upsampling filter.


// A representation of a gradient bilinear spatial upsampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradientNode
type CNNUpsamplingBilinearGradientNode struct {
	GradientFilterNode
}

// CNNUpsamplingBilinearGradientNodeFrom constructs a [CNNUpsamplingBilinearGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient bilinear spatial upsampling filter.
func CNNUpsamplingBilinearGradientNodeFrom(ptr unsafe.Pointer) CNNUpsamplingBilinearGradientNode {
	return CNNUpsamplingBilinearGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradientnode/2947991-initwithsourcegradient
func NewCNNUpsamplingBilinearGradientNodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, scaleFactorX float64, scaleFactorY float64) CNNUpsamplingBilinearGradientNode {
	instance := getCNNUpsamplingBilinearGradientNodeClass().Alloc()
	rv := objc.Send[CNNUpsamplingBilinearGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradientnode/2948025-nodewithsourcegradient
func (cc _CNNUpsamplingBilinearGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, scaleFactorX float64, scaleFactorY float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradientnode/2948051-scalefactorx
func (c_ CNNUpsamplingBilinearGradientNode) ScaleFactorX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradientnode/2948051-scalefactorx
func (c_ CNNUpsamplingBilinearGradientNode) SetScaleFactorX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradientnode/2948054-scalefactory
func (c_ CNNUpsamplingBilinearGradientNode) ScaleFactorY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradientnode/2948054-scalefactory
func (c_ CNNUpsamplingBilinearGradientNode) SetScaleFactorY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorY:"), value)
}







