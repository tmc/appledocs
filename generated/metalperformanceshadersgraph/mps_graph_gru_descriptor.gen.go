// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphGRUDescriptor] class.
var (
	GraphGRUDescriptorClass     _GraphGRUDescriptorClass
	GraphGRUDescriptorClassOnce sync.Once
)

func getGraphGRUDescriptorClass() _GraphGRUDescriptorClass {
	GraphGRUDescriptorClassOnce.Do(func() {
		GraphGRUDescriptorClass = _GraphGRUDescriptorClass{objc.GetClass("MPSGraphGRUDescriptor")}
	})
	return GraphGRUDescriptorClass
}

type _GraphGRUDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphGRUDescriptor] class.
type IGraphGRUDescriptor interface {
	IGraphObject
}

// The class that defines the parameters for a gated recurrent unit (GRU) operation.
//
// Use this descriptor with the following methods:
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor
type GraphGRUDescriptor struct {
	GraphObject
}

// GraphGRUDescriptorFrom constructs a [GraphGRUDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a gated recurrent unit (GRU) operation.
func GraphGRUDescriptorFrom(ptr unsafe.Pointer) GraphGRUDescriptor {
	return GraphGRUDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphGRUDescriptorClass) Alloc() GraphGRUDescriptor {
	rv := objc.Send[GraphGRUDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphGRUDescriptorClass) New() GraphGRUDescriptor {
	rv := objc.Send[GraphGRUDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphGRUDescriptor) Init() GraphGRUDescriptor {
	rv := objc.Send[GraphGRUDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphGRUDescriptor) Autorelease() GraphGRUDescriptor {
	rv := objc.Send[GraphGRUDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphGRUDescriptor creates a new GraphGRUDescriptor instance.
func NewGraphGRUDescriptor() GraphGRUDescriptor {
	return getGraphGRUDescriptorClass().New()
}


// A parameter that enables the GRU layer to support training.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/training
func (g_ GraphGRUDescriptor) Training() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("training"))
	return rv
}


// SetTraining sets the value of the training property.
// A parameter that enables the GRU layer to support training.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/training
func (g_ GraphGRUDescriptor) SetTraining(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTraining:"), value)
}

// A parameter that defines the activation function to use with the update-gate of the GRU operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/updateGateActivation
func (g_ GraphGRUDescriptor) UpdateGateActivation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("updateGateActivation"))
	return rv
}


// SetUpdateGateActivation sets the value of the updateGateActivation property.
// A parameter that defines the activation function to use with the update-gate of the GRU operation.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphGRUDescriptor/updateGateActivation
func (g_ GraphGRUDescriptor) SetUpdateGateActivation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUpdateGateActivation:"), value)
}



