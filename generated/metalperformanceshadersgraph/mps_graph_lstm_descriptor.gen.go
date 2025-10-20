// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphLSTMDescriptor] class.
var (
	GraphLSTMDescriptorClass     _GraphLSTMDescriptorClass
	GraphLSTMDescriptorClassOnce sync.Once
)

func getGraphLSTMDescriptorClass() _GraphLSTMDescriptorClass {
	GraphLSTMDescriptorClassOnce.Do(func() {
		GraphLSTMDescriptorClass = _GraphLSTMDescriptorClass{objc.GetClass("MPSGraphLSTMDescriptor")}
	})
	return GraphLSTMDescriptorClass
}

type _GraphLSTMDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphLSTMDescriptor] class.
type IGraphLSTMDescriptor interface {
	IGraphObject
}

// The class that defines the parameters for a long short-term memory (LSTM) operation.
//
// Use this descriptor with the following methods:
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor
type GraphLSTMDescriptor struct {
	GraphObject
}

// GraphLSTMDescriptorFrom constructs a [GraphLSTMDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a long short-term memory (LSTM) operation.
func GraphLSTMDescriptorFrom(ptr unsafe.Pointer) GraphLSTMDescriptor {
	return GraphLSTMDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphLSTMDescriptorClass) Alloc() GraphLSTMDescriptor {
	rv := objc.Send[GraphLSTMDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphLSTMDescriptorClass) New() GraphLSTMDescriptor {
	rv := objc.Send[GraphLSTMDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphLSTMDescriptor) Init() GraphLSTMDescriptor {
	rv := objc.Send[GraphLSTMDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphLSTMDescriptor) Autorelease() GraphLSTMDescriptor {
	rv := objc.Send[GraphLSTMDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphLSTMDescriptor creates a new GraphLSTMDescriptor instance.
func NewGraphLSTMDescriptor() GraphLSTMDescriptor {
	return getGraphLSTMDescriptorClass().New()
}


// A parameter that defines the activation function used with the cell gate of the LSTM operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/cellGateActivation
func (g_ GraphLSTMDescriptor) CellGateActivation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("cellGateActivation"))
	return rv
}


// SetCellGateActivation sets the value of the cellGateActivation property.
// A parameter that defines the activation function used with the cell gate of the LSTM operation.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphLSTMDescriptor/cellGateActivation
func (g_ GraphLSTMDescriptor) SetCellGateActivation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCellGateActivation:"), value)
}


