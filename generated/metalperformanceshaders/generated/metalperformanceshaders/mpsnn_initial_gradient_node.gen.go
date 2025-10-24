// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNInitialGradientNode */


/* debug [class_header]: Header for MPSNNInitialGradientNode */
// The class instance for the [InitialGradientNode] class.
var (
	InitialGradientNodeClass     _InitialGradientNodeClass
	InitialGradientNodeClassOnce sync.Once
)

func getInitialGradientNodeClass() _InitialGradientNodeClass {
	InitialGradientNodeClassOnce.Do(func() {
		InitialGradientNodeClass = _InitialGradientNodeClass{objc.GetClass("MPSNNInitialGradientNode")}
	})
	return InitialGradientNodeClass
}

type _InitialGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InitialGradientNode */
// An interface definition for the [InitialGradientNode] class.
type IInitialGradientNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for InitialGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InitialGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InitialGradientNode */
// Alloc allocates a new instance without initialization.
func (ic _InitialGradientNodeClass) Alloc() InitialGradientNode {
	rv := objc.Send[InitialGradientNode](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InitialGradientNodeClass) New() InitialGradientNode {
	rv := objc.Send[InitialGradientNode](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InitialGradientNode) Init() InitialGradientNode {
	rv := objc.Send[InitialGradientNode](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InitialGradientNode) Autorelease() InitialGradientNode {
	rv := objc.Send[InitialGradientNode](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInitialGradientNode creates a new InitialGradientNode instance.
func NewInitialGradientNode() InitialGradientNode {
	return getInitialGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InitialGradientNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradientNode
type InitialGradientNode struct {
	FilterNode
}

// InitialGradientNodeFrom constructs a [InitialGradientNode] from an unsafe.Pointer.
func InitialGradientNodeFrom(ptr unsafe.Pointer) InitialGradientNode {
	return InitialGradientNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InitialGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnninitialgradientnode/3131848-initwithsource
func NewInitialGradientNodeWithSource(source IImageNode) InitialGradientNode {
	instance := getInitialGradientNodeClass().Alloc()
	rv := objc.Send[InitialGradientNode](instance.ID, objc.Sel("initWithSource:"), source)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewInitialGradientNodeWithSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InitialGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnninitialgradientnode/3131849-nodewithsource
func (ic _InitialGradientNodeClass) NodeWithSource(source IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("nodeWithSource:"), source)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InitialGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InitialGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InitialGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNInitialGradientNode */


