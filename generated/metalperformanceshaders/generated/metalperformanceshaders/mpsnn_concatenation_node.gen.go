// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNConcatenationNode */


/* debug [class_header]: Header for MPSNNConcatenationNode */
// The class instance for the [ConcatenationNode] class.
var (
	ConcatenationNodeClass     _ConcatenationNodeClass
	ConcatenationNodeClassOnce sync.Once
)

func getConcatenationNodeClass() _ConcatenationNodeClass {
	ConcatenationNodeClassOnce.Do(func() {
		ConcatenationNodeClass = _ConcatenationNodeClass{objc.GetClass("MPSNNConcatenationNode")}
	})
	return ConcatenationNodeClass
}

type _ConcatenationNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ConcatenationNode */
// An interface definition for the [ConcatenationNode] class.
type IConcatenationNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for ConcatenationNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ConcatenationNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ConcatenationNode */
// Alloc allocates a new instance without initialization.
func (cc _ConcatenationNodeClass) Alloc() ConcatenationNode {
	rv := objc.Send[ConcatenationNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ConcatenationNodeClass) New() ConcatenationNode {
	rv := objc.Send[ConcatenationNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConcatenationNode) Init() ConcatenationNode {
	rv := objc.Send[ConcatenationNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConcatenationNode) Autorelease() ConcatenationNode {
	rv := objc.Send[ConcatenationNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConcatenationNode creates a new ConcatenationNode instance.
func NewConcatenationNode() ConcatenationNode {
	return getConcatenationNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ConcatenationNode */
// A representation of the results from one or more kernels.


// A representation of the results from one or more kernels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationNode
type ConcatenationNode struct {
	FilterNode
}

// ConcatenationNodeFrom constructs a [ConcatenationNode] from an unsafe.Pointer.
//
// A representation of the results from one or more kernels.
func ConcatenationNodeFrom(ptr unsafe.Pointer) ConcatenationNode {
	return ConcatenationNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ConcatenationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationnode/2866423-initwithsources
func NewConcatenationNodeWithSources(sourceNodes unsafe.Pointer) ConcatenationNode {
	instance := getConcatenationNodeClass().Alloc()
	rv := objc.Send[ConcatenationNode](instance.ID, objc.Sel("initWithSources:"), sourceNodes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewConcatenationNodeWithSources */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ConcatenationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationnode/2866432-nodewithsources
func (cc _ConcatenationNodeClass) NodeWithSources(sourceNodes unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSources:"), sourceNodes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSources) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ConcatenationNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ConcatenationNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ConcatenationNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNConcatenationNode */


