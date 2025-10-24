// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphTensor */


/* debug [class_header]: Header for MPSGraphTensor */
// The class instance for the [GraphTensor] class.
var (
	GraphTensorClass     _GraphTensorClass
	GraphTensorClassOnce sync.Once
)

func getGraphTensorClass() _GraphTensorClass {
	GraphTensorClassOnce.Do(func() {
		GraphTensorClass = _GraphTensorClass{objc.GetClass("MPSGraphTensor")}
	})
	return GraphTensorClass
}

type _GraphTensorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphTensor */
// An interface definition for the [GraphTensor] class.
type IGraphTensor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphTensor */
	// properties:
	DataType() objc.IObject /* cross-framework: DataType */
	Operation() IMPSGraphOperation
	Shape() Shape /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphTensor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphTensor */
// Alloc allocates a new instance without initialization.
func (gc _GraphTensorClass) Alloc() GraphTensor {
	rv := objc.Send[GraphTensor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphTensorClass) New() GraphTensor {
	rv := objc.Send[GraphTensor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphTensor) Init() GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphTensor) Autorelease() GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphTensor creates a new GraphTensor instance.
func NewGraphTensor() GraphTensor {
	return getGraphTensorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphTensor */
// The symbolic representation of a compute data type.
//
// will take a refrence, this is so can work with the tensor. All tensors are created, owned and destroyed by the MPSGraph


// The symbolic representation of a compute data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor
type GraphTensor struct {
	GraphObject
}

// GraphTensorFrom constructs a [GraphTensor] from an unsafe.Pointer.
//
// The symbolic representation of a compute data type.
func GraphTensorFrom(ptr unsafe.Pointer) GraphTensor {
	return GraphTensor{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphTensor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphTensor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphTensor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphTensor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphTensor */

// The data type of the tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor/dataType
func (g_ GraphTensor) DataType() objc.IObject /* cross-framework: DataType */ {
	rv := objc.Send[metalperformanceshaders.DataType](g_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The operation responsible for creating this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor/operation
func (g_ GraphTensor) Operation() IMPSGraphOperation {
	rv := objc.Send[GraphOperation](g_.ID, objc.Sel("operation"))
	return rv
}/* debug [instance_properties/getter]: operation */


// The shape of the tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor/shape
func (g_ GraphTensor) Shape() Shape /* not a class type */ {
	rv := objc.Send[Shape](g_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphTensor */



