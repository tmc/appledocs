// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [GraphOperation] class.
type IGraphOperation interface {
	IGraphObject
	

	// properties:
	ControlDependencies() []GraphOperation
	Graph() IMPSGraph
	InputTensors() []GraphTensor
	Name() foundation.foundation.INSString
	OutputTensors() []GraphTensor


	

	// methods:


}





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

























// The set of operations guaranteed to execute before this operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/controlDependencies
func (g_ GraphOperation) ControlDependencies() []GraphOperation {
	rv := objc.Send[[]GraphOperation](g_.ID, objc.Sel("controlDependencies"))
	return rv
}


// The graph on which the operation is defined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/graph
func (g_ GraphOperation) Graph() IMPSGraph {
	rv := objc.Send[Graph](g_.ID, objc.Sel("graph"))
	return rv
}


// The input tensors of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/inputTensors
func (g_ GraphOperation) InputTensors() []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("inputTensors"))
	return rv
}


// Name of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/name
func (g_ GraphOperation) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("name"))
	return rv
}


// The output tensors of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/outputTensors
func (g_ GraphOperation) OutputTensors() []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("outputTensors"))
	return rv
}








