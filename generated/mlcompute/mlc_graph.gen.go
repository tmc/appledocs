// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCGraph */


/* debug [class_header]: Header for MLCGraph */
// The class instance for the [CGraph] class.
var (
	CGraphClass     _CGraphClass
	CGraphClassOnce sync.Once
)

func getCGraphClass() _CGraphClass {
	CGraphClassOnce.Do(func() {
		CGraphClass = _CGraphClass{objc.GetClass("MLCGraph")}
	})
	return CGraphClass
}

type _CGraphClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CGraph */
// An interface definition for the [CGraph] class.
type ICGraph interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CGraph */
	// properties:
	Device() IMLCDevice
	Layers() []CLayer
	SummarizedDOTDescription() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CGraph */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CGraph */
// Alloc allocates a new instance without initialization.
func (cc _CGraphClass) Alloc() CGraph {
	rv := objc.Send[CGraph](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CGraphClass) New() CGraph {
	rv := objc.Send[CGraph](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CGraph) Init() CGraph {
	rv := objc.Send[CGraph](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CGraph) Autorelease() CGraph {
	rv := objc.Send[CGraph](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCGraph creates a new CGraph instance.
func NewCGraph() CGraph {
	return getCGraphClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CGraph */
// A graph of layers you use to build a training or inference graph.


// A graph of layers you use to build a training or inference graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph
type CGraph struct {
	objectivec.Object
}

// CGraphFrom constructs a [CGraph] from an unsafe.Pointer.
//
// A graph of layers you use to build a training or inference graph.
func CGraphFrom(ptr unsafe.Pointer) CGraph {
	return CGraph{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CGraph *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CGraph */

// Creates a new graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/graph
func (cc _CGraphClass) Graph() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("graph"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Graph) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CGraph */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CGraph */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CGraph */

// The device you’ll use for compiling and executing a graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/device
func (c_ CGraph) Device() IMLCDevice {
	rv := objc.Send[CDevice](c_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// An array that contains the layers in the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/layers
func (c_ CGraph) Layers() []CLayer {
	rv := objc.Send[[]CLayer](c_.ID, objc.Sel("layers"))
	return rv
}/* debug [instance_properties/getter]: layers */


// A DOT representation of the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/summarizedDOTDescription
func (c_ CGraph) SummarizedDOTDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("summarizedDOTDescription"))
	return rv
}/* debug [instance_properties/getter]: summarizedDOTDescription */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCGraph */



