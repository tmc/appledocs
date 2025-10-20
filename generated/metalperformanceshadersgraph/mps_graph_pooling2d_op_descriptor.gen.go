// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphPooling2DOpDescriptor] class.
var (
	GraphPooling2DOpDescriptorClass     _GraphPooling2DOpDescriptorClass
	GraphPooling2DOpDescriptorClassOnce sync.Once
)

func getGraphPooling2DOpDescriptorClass() _GraphPooling2DOpDescriptorClass {
	GraphPooling2DOpDescriptorClassOnce.Do(func() {
		GraphPooling2DOpDescriptorClass = _GraphPooling2DOpDescriptorClass{objc.GetClass("MPSGraphPooling2DOpDescriptor")}
	})
	return GraphPooling2DOpDescriptorClass
}

type _GraphPooling2DOpDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphPooling2DOpDescriptor] class.
type IGraphPooling2DOpDescriptor interface {
	IGraphObject
}

// The class that defines the parameters for a 2D pooling operation.
//
// Use this descriptor with the following methods:
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor
type GraphPooling2DOpDescriptor struct {
	GraphObject
}

// GraphPooling2DOpDescriptorFrom constructs a [GraphPooling2DOpDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a 2D pooling operation.
func GraphPooling2DOpDescriptorFrom(ptr unsafe.Pointer) GraphPooling2DOpDescriptor {
	return GraphPooling2DOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphPooling2DOpDescriptorClass) Alloc() GraphPooling2DOpDescriptor {
	rv := objc.Send[GraphPooling2DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphPooling2DOpDescriptorClass) New() GraphPooling2DOpDescriptor {
	rv := objc.Send[GraphPooling2DOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphPooling2DOpDescriptor) Init() GraphPooling2DOpDescriptor {
	rv := objc.Send[GraphPooling2DOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphPooling2DOpDescriptor) Autorelease() GraphPooling2DOpDescriptor {
	rv := objc.Send[GraphPooling2DOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphPooling2DOpDescriptor creates a new GraphPooling2DOpDescriptor instance.
func NewGraphPooling2DOpDescriptor() GraphPooling2DOpDescriptor {
	return getGraphPooling2DOpDescriptorClass().New()
}


// Defines the explicit padding value for the width dimension to add before the data.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingLeft
func (g_ GraphPooling2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingLeft"))
	return rv
}


// SetPaddingLeft sets the value of the paddingLeft property.
// Defines the explicit padding value for the width dimension to add before the data.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingLeft
func (g_ GraphPooling2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}


