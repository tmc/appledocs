// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MLCInferenceGraph */


/* debug [class_header]: Header for MLCInferenceGraph */
// The class instance for the [CInferenceGraph] class.
var (
	CInferenceGraphClass     _CInferenceGraphClass
	CInferenceGraphClassOnce sync.Once
)

func getCInferenceGraphClass() _CInferenceGraphClass {
	CInferenceGraphClassOnce.Do(func() {
		CInferenceGraphClass = _CInferenceGraphClass{objc.GetClass("MLCInferenceGraph")}
	})
	return CInferenceGraphClass
}

type _CInferenceGraphClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CInferenceGraph */
// An interface definition for the [CInferenceGraph] class.
type ICInferenceGraph interface {
	ICGraph
	
/* debug [class_interface_properties]: Properties for CInferenceGraph */
	// properties:
	DeviceMemorySize() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CInferenceGraph */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CInferenceGraph */
// Alloc allocates a new instance without initialization.
func (cc _CInferenceGraphClass) Alloc() CInferenceGraph {
	rv := objc.Send[CInferenceGraph](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CInferenceGraphClass) New() CInferenceGraph {
	rv := objc.Send[CInferenceGraph](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CInferenceGraph) Init() CInferenceGraph {
	rv := objc.Send[CInferenceGraph](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CInferenceGraph) Autorelease() CInferenceGraph {
	rv := objc.Send[CInferenceGraph](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCInferenceGraph creates a new CInferenceGraph instance.
func NewCInferenceGraph() CInferenceGraph {
	return getCInferenceGraphClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CInferenceGraph */
// An inference graph created from one or more MLCGraph instances plus additional layers added directly to the inference graph.


// An inference graph created from one or more MLCGraph instances plus additional layers added directly to the inference graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInferenceGraph
type CInferenceGraph struct {
	CGraph
}

// CInferenceGraphFrom constructs a [CInferenceGraph] from an unsafe.Pointer.
//
// An inference graph created from one or more MLCGraph instances plus additional layers added directly to the inference graph.
func CInferenceGraphFrom(ptr unsafe.Pointer) CInferenceGraph {
	return CInferenceGraph{
		CGraph: CGraphFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CInferenceGraph */

// Creates an inference graph with the layers from the graph objects you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInferenceGraph/init(graphObjects:)
func NewCInferenceGraphWithGraphObjects(graphObjects []CGraph) CInferenceGraph {
	rv := objc.Send[CInferenceGraph](objc.ID(getCInferenceGraphClass().class), objc.Sel("graphWithGraphObjects:"), graphObjects)
	return rv
}/* debug [class_init_methods/constructor]: NewCInferenceGraphWithGraphObjects */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CInferenceGraph */

// Creates an inference graph with the layers from the graph objects you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInferenceGraph/init(graphObjects:)
func (cc _CInferenceGraphClass) GraphWithGraphObjects(graphObjects []CGraph) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("graphWithGraphObjects:"), graphObjects)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphWithGraphObjects) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CInferenceGraph */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CInferenceGraph */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CInferenceGraph */

// The device memory size in bytes for all intermediate tensors in the inference graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInferenceGraph/deviceMemorySize
func (c_ CInferenceGraph) DeviceMemorySize() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("deviceMemorySize"))
	return rv
}/* debug [instance_properties/getter]: deviceMemorySize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCInferenceGraph */


