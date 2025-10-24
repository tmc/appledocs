// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNConcatenationGradientNode */


/* debug [class_header]: Header for MPSNNConcatenationGradientNode */
// The class instance for the [ConcatenationGradientNode] class.
var (
	ConcatenationGradientNodeClass     _ConcatenationGradientNodeClass
	ConcatenationGradientNodeClassOnce sync.Once
)

func getConcatenationGradientNodeClass() _ConcatenationGradientNodeClass {
	ConcatenationGradientNodeClassOnce.Do(func() {
		ConcatenationGradientNodeClass = _ConcatenationGradientNodeClass{objc.GetClass("MPSNNConcatenationGradientNode")}
	})
	return ConcatenationGradientNodeClass
}

type _ConcatenationGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ConcatenationGradientNode */
// An interface definition for the [ConcatenationGradientNode] class.
type IConcatenationGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for ConcatenationGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ConcatenationGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ConcatenationGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _ConcatenationGradientNodeClass) Alloc() ConcatenationGradientNode {
	rv := objc.Send[ConcatenationGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ConcatenationGradientNodeClass) New() ConcatenationGradientNode {
	rv := objc.Send[ConcatenationGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConcatenationGradientNode) Init() ConcatenationGradientNode {
	rv := objc.Send[ConcatenationGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConcatenationGradientNode) Autorelease() ConcatenationGradientNode {
	rv := objc.Send[ConcatenationGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConcatenationGradientNode creates a new ConcatenationGradientNode instance.
func NewConcatenationGradientNode() ConcatenationGradientNode {
	return getConcatenationGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ConcatenationGradientNode */
// A representation of the results from one or more gradient kernels.


// A representation of the results from one or more gradient kernels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationGradientNode
type ConcatenationGradientNode struct {
	GradientFilterNode
}

// ConcatenationGradientNodeFrom constructs a [ConcatenationGradientNode] from an unsafe.Pointer.
//
// A representation of the results from one or more gradient kernels.
func ConcatenationGradientNodeFrom(ptr unsafe.Pointer) ConcatenationGradientNode {
	return ConcatenationGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ConcatenationGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationgradientnode/2951934-initwithsourcegradient
func NewConcatenationGradientNodeWithSourceGradientSourceImageGradientState(gradientSourceNode IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) ConcatenationGradientNode {
	instance := getConcatenationGradientNodeClass().Alloc()
	rv := objc.Send[ConcatenationGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradientSourceNode, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewConcatenationGradientNodeWithSourceGradientSourceImageGradientState */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ConcatenationGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationgradientnode/2951948-nodewithsourcegradient
func (cc _ConcatenationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(gradientSourceNode IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), gradientSourceNode, sourceImage, gradientState)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientState) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ConcatenationGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ConcatenationGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ConcatenationGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNConcatenationGradientNode */


