// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNDropoutNode */


/* debug [class_header]: Header for MPSCNNDropoutNode */
// The class instance for the [CNNDropoutNode] class.
var (
	CNNDropoutNodeClass     _CNNDropoutNodeClass
	CNNDropoutNodeClassOnce sync.Once
)

func getCNNDropoutNodeClass() _CNNDropoutNodeClass {
	CNNDropoutNodeClassOnce.Do(func() {
		CNNDropoutNodeClass = _CNNDropoutNodeClass{objc.GetClass("MPSCNNDropoutNode")}
	})
	return CNNDropoutNodeClass
}

type _CNNDropoutNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNDropoutNode */
// An interface definition for the [CNNDropoutNode] class.
type ICNNDropoutNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNDropoutNode */
	// properties:
	KeepProbability() objectivec.IObject
	SetKeepProbability(value objectivec.IObject)
	MaskStrideInPixels() Size get /* not a class type */
	SetMaskStrideInPixels(value Size get /* not a class type */)
	Seed() objectivec.IObject
	SetSeed(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNDropoutNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNDropoutNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNDropoutNodeClass) Alloc() CNNDropoutNode {
	rv := objc.Send[CNNDropoutNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDropoutNodeClass) New() CNNDropoutNode {
	rv := objc.Send[CNNDropoutNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDropoutNode) Init() CNNDropoutNode {
	rv := objc.Send[CNNDropoutNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDropoutNode) Autorelease() CNNDropoutNode {
	rv := objc.Send[CNNDropoutNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDropoutNode creates a new CNNDropoutNode instance.
func NewCNNDropoutNode() CNNDropoutNode {
	return getCNNDropoutNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNDropoutNode */
// A representation of a dropout filter.


// A representation of a dropout filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode
type CNNDropoutNode struct {
	FilterNode
}

// CNNDropoutNodeFrom constructs a [CNNDropoutNode] from an unsafe.Pointer.
//
// A representation of a dropout filter.
func CNNDropoutNodeFrom(ptr unsafe.Pointer) CNNDropoutNode {
	return CNNDropoutNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNDropoutNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947969-initwithsource
func NewCNNDropoutNodeWithSource(source IImageNode) CNNDropoutNode {
	instance := getCNNDropoutNodeClass().Alloc()
	rv := objc.Send[CNNDropoutNode](instance.ID, objc.Sel("initWithSource:"), source)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDropoutNodeWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2948000-initwithsource
func NewCNNDropoutNodeWithSourceKeepProbability(source IImageNode, keepProbability float32) CNNDropoutNode {
	instance := getCNNDropoutNodeClass().Alloc()
	rv := objc.Send[CNNDropoutNode](instance.ID, objc.Sel("initWithSource:keepProbability:"), source, keepProbability)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDropoutNodeWithSourceKeepProbability */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947990-initwithsource
func NewCNNDropoutNodeWithSourceKeepProbabilitySeedMaskStrideInPixels(source IImageNode, keepProbability float32, seed uint, maskStrideInPixels Size /* not a class type */) CNNDropoutNode {
	instance := getCNNDropoutNodeClass().Alloc()
	rv := objc.Send[CNNDropoutNode](instance.ID, objc.Sel("initWithSource:keepProbability:seed:maskStrideInPixels:"), source, keepProbability, seed, maskStrideInPixels)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDropoutNodeWithSourceKeepProbabilitySeedMaskStrideInPixels */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNDropoutNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947975-nodewithsource
func (cc _CNNDropoutNodeClass) NodeWithSourceKeepProbability(source IImageNode, keepProbability float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:keepProbability:"), source, keepProbability)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceKeepProbability) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947989-nodewithsource
func (cc _CNNDropoutNodeClass) NodeWithSourceKeepProbabilitySeedMaskStrideInPixels(source IImageNode, keepProbability float32, seed uint, maskStrideInPixels Size /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:keepProbability:seed:maskStrideInPixels:"), source, keepProbability, seed, maskStrideInPixels)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceKeepProbabilitySeedMaskStrideInPixels) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2948007-nodewithsource
func (cc _CNNDropoutNodeClass) NodeWithSource(source IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), source)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNDropoutNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNDropoutNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNDropoutNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947982-keepprobability
func (c_ CNNDropoutNode) KeepProbability() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("keepProbability"))
	return rv
}/* debug [instance_properties/getter]: keepProbability */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947982-keepprobability
func (c_ CNNDropoutNode) SetKeepProbability(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeepProbability:"), value)
}/* debug [instance_properties/setter]: keepProbability */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947998-maskstrideinpixels
func (c_ CNNDropoutNode) MaskStrideInPixels() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("maskStrideInPixels"))
	return rv
}/* debug [instance_properties/getter]: maskStrideInPixels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2947998-maskstrideinpixels
func (c_ CNNDropoutNode) SetMaskStrideInPixels(value Size get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaskStrideInPixels:"), value)
}/* debug [instance_properties/setter]: maskStrideInPixels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2948030-seed
func (c_ CNNDropoutNode) Seed() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("seed"))
	return rv
}/* debug [instance_properties/getter]: seed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutnode/2948030-seed
func (c_ CNNDropoutNode) SetSeed(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSeed:"), value)
}/* debug [instance_properties/setter]: seed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNDropoutNode */


