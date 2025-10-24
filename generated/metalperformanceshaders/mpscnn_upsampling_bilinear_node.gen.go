// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNUpsamplingBilinearNode */


/* debug [class_header]: Header for MPSCNNUpsamplingBilinearNode */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNUpsamplingBilinearNode */
// An interface definition for the [CNNUpsamplingBilinearNode] class.
type ICNNUpsamplingBilinearNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNUpsamplingBilinearNode */
	// properties:
	ScaleFactorY() objectivec.IObject
	SetScaleFactorY(value objectivec.IObject)
	ScaleFactorX() objectivec.IObject
	SetScaleFactorX(value objectivec.IObject)
	AlignCorners() objectivec.IObject
	SetAlignCorners(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNUpsamplingBilinearNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNUpsamplingBilinearNode */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNUpsamplingBilinearNode */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNUpsamplingBilinearNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875152-initwithsource
func NewCNNUpsamplingBilinearNodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IImageNode, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinearNode {
	instance := getCNNUpsamplingBilinearNodeClass().Alloc()
	rv := objc.Send[CNNUpsamplingBilinearNode](instance.ID, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNUpsamplingBilinearNodeWithSourceIntegerScaleFactorXIntegerScaleFactorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2966688-initwithsource
func NewCNNUpsamplingBilinearNodeWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(sourceNode IImageNode, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) CNNUpsamplingBilinearNode {
	instance := getCNNUpsamplingBilinearNodeClass().Alloc()
	rv := objc.Send[CNNUpsamplingBilinearNode](instance.ID, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:alignCorners:"), sourceNode, integerScaleFactorX, integerScaleFactorY, alignCorners)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNUpsamplingBilinearNodeWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNUpsamplingBilinearNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875987-nodewithsource
func (cc _CNNUpsamplingBilinearNodeClass) NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IImageNode, integerScaleFactorX uint, integerScaleFactorY uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2966689-nodewithsource
func (cc _CNNUpsamplingBilinearNodeClass) NodeWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(sourceNode IImageNode, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:integerScaleFactorX:integerScaleFactorY:alignCorners:"), sourceNode, integerScaleFactorX, integerScaleFactorY, alignCorners)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNUpsamplingBilinearNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNUpsamplingBilinearNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNUpsamplingBilinearNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875150-scalefactory
func (c_ CNNUpsamplingBilinearNode) ScaleFactorY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorY"))
	return rv
}/* debug [instance_properties/getter]: scaleFactorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875150-scalefactory
func (c_ CNNUpsamplingBilinearNode) SetScaleFactorY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorY:"), value)
}/* debug [instance_properties/setter]: scaleFactorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875153-scalefactorx
func (c_ CNNUpsamplingBilinearNode) ScaleFactorX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorX"))
	return rv
}/* debug [instance_properties/getter]: scaleFactorX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875153-scalefactorx
func (c_ CNNUpsamplingBilinearNode) SetScaleFactorX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorX:"), value)
}/* debug [instance_properties/setter]: scaleFactorX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2966687-aligncorners
func (c_ CNNUpsamplingBilinearNode) AlignCorners() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alignCorners"))
	return rv
}/* debug [instance_properties/getter]: alignCorners */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2966687-aligncorners
func (c_ CNNUpsamplingBilinearNode) SetAlignCorners(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlignCorners:"), value)
}/* debug [instance_properties/setter]: alignCorners */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNUpsamplingBilinearNode */


