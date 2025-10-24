// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNUpsamplingNearestNode */


/* debug [class_header]: Header for MPSCNNUpsamplingNearestNode */
// The class instance for the [CNNUpsamplingNearestNode] class.
var (
	CNNUpsamplingNearestNodeClass     _CNNUpsamplingNearestNodeClass
	CNNUpsamplingNearestNodeClassOnce sync.Once
)

func getCNNUpsamplingNearestNodeClass() _CNNUpsamplingNearestNodeClass {
	CNNUpsamplingNearestNodeClassOnce.Do(func() {
		CNNUpsamplingNearestNodeClass = _CNNUpsamplingNearestNodeClass{objc.GetClass("MPSCNNUpsamplingNearestNode")}
	})
	return CNNUpsamplingNearestNodeClass
}

type _CNNUpsamplingNearestNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNUpsamplingNearestNode */
// An interface definition for the [CNNUpsamplingNearestNode] class.
type ICNNUpsamplingNearestNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNUpsamplingNearestNode */
	// properties:
	ScaleFactorY() objectivec.IObject
	SetScaleFactorY(value objectivec.IObject)
	ScaleFactorX() objectivec.IObject
	SetScaleFactorX(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNUpsamplingNearestNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNUpsamplingNearestNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingNearestNodeClass) Alloc() CNNUpsamplingNearestNode {
	rv := objc.Send[CNNUpsamplingNearestNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingNearestNodeClass) New() CNNUpsamplingNearestNode {
	rv := objc.Send[CNNUpsamplingNearestNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingNearestNode) Init() CNNUpsamplingNearestNode {
	rv := objc.Send[CNNUpsamplingNearestNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingNearestNode) Autorelease() CNNUpsamplingNearestNode {
	rv := objc.Send[CNNUpsamplingNearestNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingNearestNode creates a new CNNUpsamplingNearestNode instance.
func NewCNNUpsamplingNearestNode() CNNUpsamplingNearestNode {
	return getCNNUpsamplingNearestNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNUpsamplingNearestNode */
// A representation of a nearest spatial upsampling filter.


// A representation of a nearest spatial upsampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestNode
type CNNUpsamplingNearestNode struct {
	FilterNode
}

// CNNUpsamplingNearestNodeFrom constructs a [CNNUpsamplingNearestNode] from an unsafe.Pointer.
//
// A representation of a nearest spatial upsampling filter.
func CNNUpsamplingNearestNodeFrom(ptr unsafe.Pointer) CNNUpsamplingNearestNode {
	return CNNUpsamplingNearestNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNUpsamplingNearestNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestnode/2875222-initwithsource
func NewCNNUpsamplingNearestNodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IImageNode, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearestNode {
	instance := getCNNUpsamplingNearestNodeClass().Alloc()
	rv := objc.Send[CNNUpsamplingNearestNode](instance.ID, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNUpsamplingNearestNodeWithSourceIntegerScaleFactorXIntegerScaleFactorY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNUpsamplingNearestNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestnode/2875985-nodewithsource
func (cc _CNNUpsamplingNearestNodeClass) NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IImageNode, integerScaleFactorX uint, integerScaleFactorY uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNUpsamplingNearestNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNUpsamplingNearestNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNUpsamplingNearestNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestnode/2875155-scalefactory
func (c_ CNNUpsamplingNearestNode) ScaleFactorY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorY"))
	return rv
}/* debug [instance_properties/getter]: scaleFactorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestnode/2875155-scalefactory
func (c_ CNNUpsamplingNearestNode) SetScaleFactorY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorY:"), value)
}/* debug [instance_properties/setter]: scaleFactorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestnode/2875209-scalefactorx
func (c_ CNNUpsamplingNearestNode) ScaleFactorX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorX"))
	return rv
}/* debug [instance_properties/getter]: scaleFactorX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestnode/2875209-scalefactorx
func (c_ CNNUpsamplingNearestNode) SetScaleFactorX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorX:"), value)
}/* debug [instance_properties/setter]: scaleFactorX */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNUpsamplingNearestNode */


