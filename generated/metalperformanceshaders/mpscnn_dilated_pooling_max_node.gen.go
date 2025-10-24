// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNDilatedPoolingMaxNode */


/* debug [class_header]: Header for MPSCNNDilatedPoolingMaxNode */
// The class instance for the [CNNDilatedPoolingMaxNode] class.
var (
	CNNDilatedPoolingMaxNodeClass     _CNNDilatedPoolingMaxNodeClass
	CNNDilatedPoolingMaxNodeClassOnce sync.Once
)

func getCNNDilatedPoolingMaxNodeClass() _CNNDilatedPoolingMaxNodeClass {
	CNNDilatedPoolingMaxNodeClassOnce.Do(func() {
		CNNDilatedPoolingMaxNodeClass = _CNNDilatedPoolingMaxNodeClass{objc.GetClass("MPSCNNDilatedPoolingMaxNode")}
	})
	return CNNDilatedPoolingMaxNodeClass
}

type _CNNDilatedPoolingMaxNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNDilatedPoolingMaxNode */
// An interface definition for the [CNNDilatedPoolingMaxNode] class.
type ICNNDilatedPoolingMaxNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNDilatedPoolingMaxNode */
	// properties:
	DilationRateX() objectivec.IObject
	SetDilationRateX(value objectivec.IObject)
	DilationRateY() objectivec.IObject
	SetDilationRateY(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNDilatedPoolingMaxNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNDilatedPoolingMaxNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNDilatedPoolingMaxNodeClass) Alloc() CNNDilatedPoolingMaxNode {
	rv := objc.Send[CNNDilatedPoolingMaxNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDilatedPoolingMaxNodeClass) New() CNNDilatedPoolingMaxNode {
	rv := objc.Send[CNNDilatedPoolingMaxNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDilatedPoolingMaxNode) Init() CNNDilatedPoolingMaxNode {
	rv := objc.Send[CNNDilatedPoolingMaxNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDilatedPoolingMaxNode) Autorelease() CNNDilatedPoolingMaxNode {
	rv := objc.Send[CNNDilatedPoolingMaxNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDilatedPoolingMaxNode creates a new CNNDilatedPoolingMaxNode instance.
func NewCNNDilatedPoolingMaxNode() CNNDilatedPoolingMaxNode {
	return getCNNDilatedPoolingMaxNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNDilatedPoolingMaxNode */
// A representation of a dilated max pooling filter.


// A representation of a dilated max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode
type CNNDilatedPoolingMaxNode struct {
	FilterNode
}

// CNNDilatedPoolingMaxNodeFrom constructs a [CNNDilatedPoolingMaxNode] from an unsafe.Pointer.
//
// A representation of a dilated max pooling filter.
func CNNDilatedPoolingMaxNodeFrom(ptr unsafe.Pointer) CNNDilatedPoolingMaxNode {
	return CNNDilatedPoolingMaxNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNDilatedPoolingMaxNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2873240-initwithsource
func NewCNNDilatedPoolingMaxNodeWithSourceFilterSize(sourceNode IImageNode, size uint) CNNDilatedPoolingMaxNode {
	instance := getCNNDilatedPoolingMaxNodeClass().Alloc()
	rv := objc.Send[CNNDilatedPoolingMaxNode](instance.ID, objc.Sel("initWithSource:filterSize:"), sourceNode, size)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDilatedPoolingMaxNodeWithSourceFilterSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2887340-initwithsource
func NewCNNDilatedPoolingMaxNodeWithSourceFilterSizeStrideDilationRate(sourceNode IImageNode, size uint, stride uint, dilationRate uint) CNNDilatedPoolingMaxNode {
	instance := getCNNDilatedPoolingMaxNodeClass().Alloc()
	rv := objc.Send[CNNDilatedPoolingMaxNode](instance.ID, objc.Sel("initWithSource:filterSize:stride:dilationRate:"), sourceNode, size, stride, dilationRate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDilatedPoolingMaxNodeWithSourceFilterSizeStrideDilationRate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2887339-initwithsource
func NewCNNDilatedPoolingMaxNodeWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceNode IImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) CNNDilatedPoolingMaxNode {
	instance := getCNNDilatedPoolingMaxNodeClass().Alloc()
	rv := objc.Send[CNNDilatedPoolingMaxNode](instance.ID, objc.Sel("initWithSource:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:"), sourceNode, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDilatedPoolingMaxNodeWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNDilatedPoolingMaxNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2873227-nodewithsource
func (cc _CNNDilatedPoolingMaxNodeClass) NodeWithSourceFilterSize(sourceNode IImageNode, size uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:filterSize:"), sourceNode, size)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceFilterSize) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2919744-nodewithsource
func (cc _CNNDilatedPoolingMaxNodeClass) NodeWithSourceFilterSizeStrideDilationRate(sourceNode IImageNode, size uint, stride uint, dilationRate uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:filterSize:stride:dilationRate:"), sourceNode, size, stride, dilationRate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceFilterSizeStrideDilationRate) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNDilatedPoolingMaxNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNDilatedPoolingMaxNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNDilatedPoolingMaxNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2887341-dilationratex
func (c_ CNNDilatedPoolingMaxNode) DilationRateX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateX"))
	return rv
}/* debug [instance_properties/getter]: dilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2887341-dilationratex
func (c_ CNNDilatedPoolingMaxNode) SetDilationRateX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateX:"), value)
}/* debug [instance_properties/setter]: dilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2887342-dilationratey
func (c_ CNNDilatedPoolingMaxNode) DilationRateY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateY"))
	return rv
}/* debug [instance_properties/getter]: dilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2887342-dilationratey
func (c_ CNNDilatedPoolingMaxNode) SetDilationRateY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateY:"), value)
}/* debug [instance_properties/setter]: dilationRateY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNDilatedPoolingMaxNode */


