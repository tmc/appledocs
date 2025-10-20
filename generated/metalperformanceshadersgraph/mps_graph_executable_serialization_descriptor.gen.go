// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphExecutableSerializationDescriptor] class.
var (
	GraphExecutableSerializationDescriptorClass     _GraphExecutableSerializationDescriptorClass
	GraphExecutableSerializationDescriptorClassOnce sync.Once
)

func getGraphExecutableSerializationDescriptorClass() _GraphExecutableSerializationDescriptorClass {
	GraphExecutableSerializationDescriptorClassOnce.Do(func() {
		GraphExecutableSerializationDescriptorClass = _GraphExecutableSerializationDescriptorClass{objc.GetClass("MPSGraphExecutableSerializationDescriptor")}
	})
	return GraphExecutableSerializationDescriptorClass
}

type _GraphExecutableSerializationDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphExecutableSerializationDescriptor] class.
type IGraphExecutableSerializationDescriptor interface {
	IGraphObject
}

// A class that consists of all the levers to serialize an executable.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor
type GraphExecutableSerializationDescriptor struct {
	GraphObject
}

// GraphExecutableSerializationDescriptorFrom constructs a [GraphExecutableSerializationDescriptor] from an unsafe.Pointer.
//
// A class that consists of all the levers to serialize an executable.
func GraphExecutableSerializationDescriptorFrom(ptr unsafe.Pointer) GraphExecutableSerializationDescriptor {
	return GraphExecutableSerializationDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphExecutableSerializationDescriptorClass) Alloc() GraphExecutableSerializationDescriptor {
	rv := objc.Send[GraphExecutableSerializationDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphExecutableSerializationDescriptorClass) New() GraphExecutableSerializationDescriptor {
	rv := objc.Send[GraphExecutableSerializationDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphExecutableSerializationDescriptor) Init() GraphExecutableSerializationDescriptor {
	rv := objc.Send[GraphExecutableSerializationDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphExecutableSerializationDescriptor) Autorelease() GraphExecutableSerializationDescriptor {
	rv := objc.Send[GraphExecutableSerializationDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphExecutableSerializationDescriptor creates a new GraphExecutableSerializationDescriptor instance.
func NewGraphExecutableSerializationDescriptor() GraphExecutableSerializationDescriptor {
	return getGraphExecutableSerializationDescriptorClass().New()
}


// Flag to append to an existing .mpsgraphpackage if found at provided url.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor/append
func (g_ GraphExecutableSerializationDescriptor) Append() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("append"))
	return rv
}


// SetAppend sets the value of the append property.
// Flag to append to an existing .mpsgraphpackage if found at provided url.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor/append
func (g_ GraphExecutableSerializationDescriptor) SetAppend(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAppend:"), value)
}


