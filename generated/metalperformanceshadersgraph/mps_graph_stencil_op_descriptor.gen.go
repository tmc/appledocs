// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphStencilOpDescriptor] class.
var (
	GraphStencilOpDescriptorClass     _GraphStencilOpDescriptorClass
	GraphStencilOpDescriptorClassOnce sync.Once
)

func getGraphStencilOpDescriptorClass() _GraphStencilOpDescriptorClass {
	GraphStencilOpDescriptorClassOnce.Do(func() {
		GraphStencilOpDescriptorClass = _GraphStencilOpDescriptorClass{objc.GetClass("MPSGraphStencilOpDescriptor")}
	})
	return GraphStencilOpDescriptorClass
}

type _GraphStencilOpDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphStencilOpDescriptor] class.
type IGraphStencilOpDescriptor interface {
	IGraphObject
}

// The class that defines the parameters for a stencil operation.
//
// Use this descriptor with the following method:
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor
type GraphStencilOpDescriptor struct {
	GraphObject
}

// GraphStencilOpDescriptorFrom constructs a [GraphStencilOpDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a stencil operation.
func GraphStencilOpDescriptorFrom(ptr unsafe.Pointer) GraphStencilOpDescriptor {
	return GraphStencilOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphStencilOpDescriptorClass) Alloc() GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphStencilOpDescriptorClass) New() GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphStencilOpDescriptor) Init() GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphStencilOpDescriptor) Autorelease() GraphStencilOpDescriptor {
	rv := objc.Send[GraphStencilOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphStencilOpDescriptor creates a new GraphStencilOpDescriptor instance.
func NewGraphStencilOpDescriptor() GraphStencilOpDescriptor {
	return getGraphStencilOpDescriptorClass().New()
}


// The property that determines which values to use for padding the input tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/boundaryMode
func (g_ GraphStencilOpDescriptor) BoundaryMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("boundaryMode"))
	return rv
}


// SetBoundaryMode sets the value of the boundaryMode property.
// The property that determines which values to use for padding the input tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/boundaryMode
func (g_ GraphStencilOpDescriptor) SetBoundaryMode(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBoundaryMode:"), value)
}
// The padding value for .
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/paddingConstant
func (g_ GraphStencilOpDescriptor) PaddingConstant() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("paddingConstant"))
	return rv
}


// SetPaddingConstant sets the value of the paddingConstant property.
// The padding value for .

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/paddingConstant
func (g_ GraphStencilOpDescriptor) SetPaddingConstant(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingConstant:"), value)
}
// The property that defines strides for spatial dimensions.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/strides
func (g_ GraphStencilOpDescriptor) Strides() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("strides"))
	return rv
}


// SetStrides sets the value of the strides property.
// The property that defines strides for spatial dimensions.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/strides
func (g_ GraphStencilOpDescriptor) SetStrides(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrides:"), value)
}


