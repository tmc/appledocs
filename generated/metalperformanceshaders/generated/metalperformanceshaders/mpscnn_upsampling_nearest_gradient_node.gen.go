// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNUpsamplingNearestGradientNode */


/* debug [class_header]: Header for MPSCNNUpsamplingNearestGradientNode */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNUpsamplingNearestGradientNode */
// An interface definition for the [CNNUpsamplingNearestGradientNode] class.
type ICNNUpsamplingNearestGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for CNNUpsamplingNearestGradientNode */
	// properties:
	ScaleFactorX() objectivec.IObject
	SetScaleFactorX(value objectivec.IObject)
	ScaleFactorY() objectivec.IObject
	SetScaleFactorY(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNUpsamplingNearestGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNUpsamplingNearestGradientNode */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNUpsamplingNearestGradientNode */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNUpsamplingNearestGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2947983-initwithsourcegradient
func NewCNNUpsamplingNearestGradientNodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, scaleFactorX float64, scaleFactorY float64) CNNUpsamplingNearestGradientNode {
	instance := getCNNUpsamplingNearestGradientNodeClass().Alloc()
	rv := objc.Send[CNNUpsamplingNearestGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNUpsamplingNearestGradientNodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNUpsamplingNearestGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948021-nodewithsourcegradient
func (cc _CNNUpsamplingNearestGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, scaleFactorX float64, scaleFactorY float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNUpsamplingNearestGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNUpsamplingNearestGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNUpsamplingNearestGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948024-scalefactorx
func (c_ CNNUpsamplingNearestGradientNode) ScaleFactorX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorX"))
	return rv
}/* debug [instance_properties/getter]: scaleFactorX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948024-scalefactorx
func (c_ CNNUpsamplingNearestGradientNode) SetScaleFactorX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorX:"), value)
}/* debug [instance_properties/setter]: scaleFactorX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948035-scalefactory
func (c_ CNNUpsamplingNearestGradientNode) ScaleFactorY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorY"))
	return rv
}/* debug [instance_properties/getter]: scaleFactorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948035-scalefactory
func (c_ CNNUpsamplingNearestGradientNode) SetScaleFactorY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorY:"), value)
}/* debug [instance_properties/setter]: scaleFactorY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNUpsamplingNearestGradientNode */


