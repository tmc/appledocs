// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [GraphTensor] class.
type IGraphTensor interface {
	IGraphObject
	

	// properties:
	DataType() DataType /* not a class type */
	Operation() IMPSGraphOperation
	Shape() Shape /* not a class type */


	

	// methods:


}





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

























// The data type of the tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor/dataType
func (g_ GraphTensor) DataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](g_.ID, objc.Sel("dataType"))
	return rv
}


// The operation responsible for creating this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor/operation
func (g_ GraphTensor) Operation() IMPSGraphOperation {
	rv := objc.Send[GraphOperation](g_.ID, objc.Sel("operation"))
	return rv
}


// The shape of the tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor/shape
func (g_ GraphTensor) Shape() Shape /* not a class type */ {
	rv := objc.Send[Shape](g_.ID, objc.Sel("shape"))
	return rv
}








