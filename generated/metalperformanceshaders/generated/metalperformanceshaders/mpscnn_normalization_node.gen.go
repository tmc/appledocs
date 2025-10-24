// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNormalizationNode */


/* debug [class_header]: Header for MPSCNNNormalizationNode */
// The class instance for the [CNNNormalizationNode] class.
var (
	CNNNormalizationNodeClass     _CNNNormalizationNodeClass
	CNNNormalizationNodeClassOnce sync.Once
)

func getCNNNormalizationNodeClass() _CNNNormalizationNodeClass {
	CNNNormalizationNodeClassOnce.Do(func() {
		CNNNormalizationNodeClass = _CNNNormalizationNodeClass{objc.GetClass("MPSCNNNormalizationNode")}
	})
	return CNNNormalizationNodeClass
}

type _CNNNormalizationNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNormalizationNode */
// An interface definition for the [CNNNormalizationNode] class.
type ICNNNormalizationNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNNormalizationNode */
	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNormalizationNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNormalizationNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNNormalizationNodeClass) Alloc() CNNNormalizationNode {
	rv := objc.Send[CNNNormalizationNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNormalizationNodeClass) New() CNNNormalizationNode {
	rv := objc.Send[CNNNormalizationNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNormalizationNode) Init() CNNNormalizationNode {
	rv := objc.Send[CNNNormalizationNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNormalizationNode) Autorelease() CNNNormalizationNode {
	rv := objc.Send[CNNNormalizationNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNormalizationNode creates a new CNNNormalizationNode instance.
func NewCNNNormalizationNode() CNNNormalizationNode {
	return getCNNNormalizationNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNormalizationNode */
// Virtual base class for CNN normalization nodes.


// Virtual base class for CNN normalization nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationNode
type CNNNormalizationNode struct {
	FilterNode
}

// CNNNormalizationNodeFrom constructs a [CNNNormalizationNode] from an unsafe.Pointer.
//
// Virtual base class for CNN normalization nodes.
func CNNNormalizationNodeFrom(ptr unsafe.Pointer) CNNNormalizationNode {
	return CNNNormalizationNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866425-initwithsource
func NewCNNNormalizationNodeWithSource(sourceNode IImageNode) CNNNormalizationNode {
	instance := getCNNNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNNormalizationNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNormalizationNodeWithSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866460-nodewithsource
func (cc _CNNNormalizationNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNormalizationNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNormalizationNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866474-alpha
func (c_ CNNNormalizationNode) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866474-alpha
func (c_ CNNNormalizationNode) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866482-delta
func (c_ CNNNormalizationNode) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866482-delta
func (c_ CNNNormalizationNode) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}/* debug [instance_properties/setter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866497-beta
func (c_ CNNNormalizationNode) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866497-beta
func (c_ CNNNormalizationNode) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}/* debug [instance_properties/setter]: beta */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNormalizationNode */


