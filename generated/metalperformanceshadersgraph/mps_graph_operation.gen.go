// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MPSGraphOperation */


/* debug [class_header]: Header for MPSGraphOperation */
// The class instance for the [GraphOperation] class.
var (
	GraphOperationClass     _GraphOperationClass
	GraphOperationClassOnce sync.Once
)

func getGraphOperationClass() _GraphOperationClass {
	GraphOperationClassOnce.Do(func() {
		GraphOperationClass = _GraphOperationClass{objc.GetClass("MPSGraphOperation")}
	})
	return GraphOperationClass
}

type _GraphOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphOperation */
// An interface definition for the [GraphOperation] class.
type IGraphOperation interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphOperation */
	// properties:
	ControlDependencies() []GraphOperation
	Graph() IMPSGraph
	InputTensors() []GraphTensor
	Name() objc.IObject /* cross-framework: NSString */
	OutputTensors() []GraphTensor
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphOperation */
// Alloc allocates a new instance without initialization.
func (gc _GraphOperationClass) Alloc() GraphOperation {
	rv := objc.Send[GraphOperation](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphOperationClass) New() GraphOperation {
	rv := objc.Send[GraphOperation](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphOperation) Init() GraphOperation {
	rv := objc.Send[GraphOperation](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphOperation) Autorelease() GraphOperation {
	rv := objc.Send[GraphOperation](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphOperation creates a new GraphOperation instance.
func NewGraphOperation() GraphOperation {
	return getGraphOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphOperation */
// A symbolic representation of a compute operation.
//
// will take a refrence, this is so can work with the tensor. All operations are created, owned and destroyed by the graph.


// A symbolic representation of a compute operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation
type GraphOperation struct {
	GraphObject
}

// GraphOperationFrom constructs a [GraphOperation] from an unsafe.Pointer.
//
// A symbolic representation of a compute operation.
func GraphOperationFrom(ptr unsafe.Pointer) GraphOperation {
	return GraphOperation{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphOperation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphOperation */

// The set of operations guaranteed to execute before this operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/controlDependencies
func (g_ GraphOperation) ControlDependencies() []GraphOperation {
	rv := objc.Send[[]GraphOperation](g_.ID, objc.Sel("controlDependencies"))
	return rv
}/* debug [instance_properties/getter]: controlDependencies */


// The graph on which the operation is defined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/graph
func (g_ GraphOperation) Graph() IMPSGraph {
	rv := objc.Send[Graph](g_.ID, objc.Sel("graph"))
	return rv
}/* debug [instance_properties/getter]: graph */


// The input tensors of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/inputTensors
func (g_ GraphOperation) InputTensors() []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("inputTensors"))
	return rv
}/* debug [instance_properties/getter]: inputTensors */


// Name of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/name
func (g_ GraphOperation) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The output tensors of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/outputTensors
func (g_ GraphOperation) OutputTensors() []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("outputTensors"))
	return rv
}/* debug [instance_properties/getter]: outputTensors */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphOperation */



