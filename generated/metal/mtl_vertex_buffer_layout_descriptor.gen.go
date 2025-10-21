// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VertexBufferLayoutDescriptor] class.
var (
	VertexBufferLayoutDescriptorClass     _VertexBufferLayoutDescriptorClass
	VertexBufferLayoutDescriptorClassOnce sync.Once
)

func getVertexBufferLayoutDescriptorClass() _VertexBufferLayoutDescriptorClass {
	VertexBufferLayoutDescriptorClassOnce.Do(func() {
		VertexBufferLayoutDescriptorClass = _VertexBufferLayoutDescriptorClass{objc.GetClass("MTLVertexBufferLayoutDescriptor")}
	})
	return VertexBufferLayoutDescriptorClass
}

type _VertexBufferLayoutDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [VertexBufferLayoutDescriptor] class.
type IVertexBufferLayoutDescriptor interface {
	objectivec.IObject
}

// An object that configures how a render pipeline fetches data to send to the vertex function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor
type VertexBufferLayoutDescriptor struct {
	objectivec.Object
}

// VertexBufferLayoutDescriptorFrom constructs a [VertexBufferLayoutDescriptor] from an unsafe.Pointer.
//
// An object that configures how a render pipeline fetches data to send to the vertex function.
func VertexBufferLayoutDescriptorFrom(ptr unsafe.Pointer) VertexBufferLayoutDescriptor {
	return VertexBufferLayoutDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VertexBufferLayoutDescriptorClass) Alloc() VertexBufferLayoutDescriptor {
	rv := objc.Send[VertexBufferLayoutDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VertexBufferLayoutDescriptorClass) New() VertexBufferLayoutDescriptor {
	rv := objc.Send[VertexBufferLayoutDescriptor](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VertexBufferLayoutDescriptor) Init() VertexBufferLayoutDescriptor {
	rv := objc.Send[VertexBufferLayoutDescriptor](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VertexBufferLayoutDescriptor) Autorelease() VertexBufferLayoutDescriptor {
	rv := objc.Send[VertexBufferLayoutDescriptor](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVertexBufferLayoutDescriptor creates a new VertexBufferLayoutDescriptor instance.
func NewVertexBufferLayoutDescriptor() VertexBufferLayoutDescriptor {
	return getVertexBufferLayoutDescriptorClass().New()
}


// The circumstances under which the vertex and its attributes are presented to the vertex function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stepFunction
func (v_ VertexBufferLayoutDescriptor) StepFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("stepFunction"))
	return rv
}


// SetStepFunction sets the value of the stepFunction property.
// The circumstances under which the vertex and its attributes are presented to the vertex function.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stepFunction
func (v_ VertexBufferLayoutDescriptor) SetStepFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStepFunction:"), value)
}

// The interval at which the vertex and its attributes are presented to the vertex function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stepRate
func (v_ VertexBufferLayoutDescriptor) StepRate() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("stepRate"))
	return rv
}


// SetStepRate sets the value of the stepRate property.
// The interval at which the vertex and its attributes are presented to the vertex function.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stepRate
func (v_ VertexBufferLayoutDescriptor) SetStepRate(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStepRate:"), value)
}

// The number of bytes between the first byte of two consecutive vertices in a buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stride
func (v_ VertexBufferLayoutDescriptor) Stride() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("stride"))
	return rv
}


// SetStride sets the value of the stride property.
// The number of bytes between the first byte of two consecutive vertices in a buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stride
func (v_ VertexBufferLayoutDescriptor) SetStride(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStride:"), value)
}



