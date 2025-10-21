// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphRandomOpDescriptor] class.
var (
	GraphRandomOpDescriptorClass     _GraphRandomOpDescriptorClass
	GraphRandomOpDescriptorClassOnce sync.Once
)

func getGraphRandomOpDescriptorClass() _GraphRandomOpDescriptorClass {
	GraphRandomOpDescriptorClassOnce.Do(func() {
		GraphRandomOpDescriptorClass = _GraphRandomOpDescriptorClass{objc.GetClass("MPSGraphRandomOpDescriptor")}
	})
	return GraphRandomOpDescriptorClass
}

type _GraphRandomOpDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphRandomOpDescriptor] class.
type IGraphRandomOpDescriptor interface {
	IGraphObject
}

// A class that describes the random operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor
type GraphRandomOpDescriptor struct {
	GraphObject
}

// GraphRandomOpDescriptorFrom constructs a [GraphRandomOpDescriptor] from an unsafe.Pointer.
//
// A class that describes the random operation.
func GraphRandomOpDescriptorFrom(ptr unsafe.Pointer) GraphRandomOpDescriptor {
	return GraphRandomOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphRandomOpDescriptorClass) Alloc() GraphRandomOpDescriptor {
	rv := objc.Send[GraphRandomOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphRandomOpDescriptorClass) New() GraphRandomOpDescriptor {
	rv := objc.Send[GraphRandomOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphRandomOpDescriptor) Init() GraphRandomOpDescriptor {
	rv := objc.Send[GraphRandomOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphRandomOpDescriptor) Autorelease() GraphRandomOpDescriptor {
	rv := objc.Send[GraphRandomOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphRandomOpDescriptor creates a new GraphRandomOpDescriptor instance.
func NewGraphRandomOpDescriptor() GraphRandomOpDescriptor {
	return getGraphRandomOpDescriptorClass().New()
}


// The data type of the generated result values.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/dataType
func (g_ GraphRandomOpDescriptor) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("dataType"))
	return rv
}


// SetDataType sets the value of the dataType property.
// The data type of the generated result values.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/dataType
func (g_ GraphRandomOpDescriptor) SetDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}

// The type of distribution to draw samples from. See MPSGraphRandomDistribution.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/distribution
func (g_ GraphRandomOpDescriptor) Distribution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("distribution"))
	return rv
}


// SetDistribution sets the value of the distribution property.
// The type of distribution to draw samples from. See MPSGraphRandomDistribution.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/distribution
func (g_ GraphRandomOpDescriptor) SetDistribution(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDistribution:"), value)
}



