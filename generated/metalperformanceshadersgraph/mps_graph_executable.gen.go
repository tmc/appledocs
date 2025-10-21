// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphExecutable] class.
var (
	GraphExecutableClass     _GraphExecutableClass
	GraphExecutableClassOnce sync.Once
)

func getGraphExecutableClass() _GraphExecutableClass {
	GraphExecutableClassOnce.Do(func() {
		GraphExecutableClass = _GraphExecutableClass{objc.GetClass("MPSGraphExecutable")}
	})
	return GraphExecutableClass
}

type _GraphExecutableClass struct {
	class objc.Class
}

// An interface definition for the [GraphExecutable] class.
type IGraphExecutable interface {
	IGraphObject
	GetOutputTypesWithDeviceInputTypesCompilationDescriptor(device unsafe.Pointer, inputTypes unsafe.Pointer, compilationDescriptor unsafe.Pointer) []GraphShapedType
}

// The compiled representation of a compute graph executable.
//
// An is a compiled graph for specific feeds for specific target tensors and target operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable
type GraphExecutable struct {
	GraphObject
}

// GraphExecutableFrom constructs a [GraphExecutable] from an unsafe.Pointer.
//
// The compiled representation of a compute graph executable.
func GraphExecutableFrom(ptr unsafe.Pointer) GraphExecutable {
	return GraphExecutable{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphExecutableClass) Alloc() GraphExecutable {
	rv := objc.Send[GraphExecutable](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphExecutableClass) New() GraphExecutable {
	rv := objc.Send[GraphExecutable](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphExecutable) Init() GraphExecutable {
	rv := objc.Send[GraphExecutable](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphExecutable) Autorelease() GraphExecutable {
	rv := objc.Send[GraphExecutable](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphExecutable creates a new GraphExecutable instance.
func NewGraphExecutable() GraphExecutable {
	return getGraphExecutableClass().New()
}


// Get output shapes for a specialized executable.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/getOutputTypes(with:inputTypes:compilationDescriptor:)
func (g_ GraphExecutable) GetOutputTypesWithDeviceInputTypesCompilationDescriptor(device unsafe.Pointer, inputTypes unsafe.Pointer, compilationDescriptor unsafe.Pointer) []GraphShapedType {
	rv := objc.Send[[]GraphShapedType](g_.ID, objc.Sel("getOutputTypesWithDevice:inputTypes:compilationDescriptor:"), device, inputTypes, compilationDescriptor)
	return rv
}

// Tensors fed to the graph, can be used to order the inputs when executable is created with a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/feedtensors
func (g_ GraphExecutable) FeedTensors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("feedTensors"))
	return rv
}


// SetFeedTensors sets the value of the feedTensors property.
// Tensors fed to the graph, can be used to order the inputs when executable is created with a graph.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/feedtensors
func (g_ GraphExecutable) SetFeedTensors(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setFeedTensors:"), value)
}

// Options for the graph executable.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/options
func (g_ GraphExecutable) Options() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("options"))
	return rv
}


// SetOptions sets the value of the options property.
// Options for the graph executable.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/options
func (g_ GraphExecutable) SetOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOptions:"), value)
}

// Tensors targeted by the graph, can be used to order the outputs when executable was created with a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable/targetTensors
func (g_ GraphExecutable) TargetTensors() []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("targetTensors"))
	return rv
}



