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
}

// The symbolic representation of a compute data type.
//
// will take a refrence, this is so can work with the tensor. All tensors are created, owned and destroyed by the MPSGraph
//
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

// Alloc allocates a new instance without initialization.
func (gc _GraphTensorClass) Alloc() GraphTensor {
	rv := objc.Send[GraphTensor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The operation responsible for creating this tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensor/operation
func (g_ GraphTensor) Operation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("operation"))
	return rv
}


// SetOperation sets the value of the operation property.
// The operation responsible for creating this tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensor/operation
func (g_ GraphTensor) SetOperation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOperation:"), value)
}

// The data type of the tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor/dataType
func (g_ GraphTensor) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("dataType"))
	return rv
}

// The shape of the tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor/shape
func (g_ GraphTensor) Shape() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("shape"))
	return rv
}



