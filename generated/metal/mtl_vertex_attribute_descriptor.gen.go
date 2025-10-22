// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VertexAttributeDescriptor] class.
var (
	VertexAttributeDescriptorClass     _VertexAttributeDescriptorClass
	VertexAttributeDescriptorClassOnce sync.Once
)

func getVertexAttributeDescriptorClass() _VertexAttributeDescriptorClass {
	VertexAttributeDescriptorClassOnce.Do(func() {
		VertexAttributeDescriptorClass = _VertexAttributeDescriptorClass{objc.GetClass("MTLVertexAttributeDescriptor")}
	})
	return VertexAttributeDescriptorClass
}

type _VertexAttributeDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [VertexAttributeDescriptor] class.
type IVertexAttributeDescriptor interface {
	objectivec.IObject
	BufferIndex() uint
	SetBufferIndex(value uint)
	Format() VertexFormat
	SetFormat(value VertexFormat)
	Offset() uint
	SetOffset(value uint)
	MTLBufferLayoutStrideDynamic() int
}

// An object that determines how to store attribute data in memory and map it to the arguments of a vertex function.
//
// A vertex attribute descriptor provides organization information so a vertex shader function can locate and load data into its arguments. The descriptor maps memory locations to attribute locations. It supports access to multiple attributes (such as vertex coordinates, surface normals, and texture coordinates) that are interleaved within the same buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor
type VertexAttributeDescriptor struct {
	objectivec.Object
}

// VertexAttributeDescriptorFrom constructs a [VertexAttributeDescriptor] from an unsafe.Pointer.
//
// An object that determines how to store attribute data in memory and map it to the arguments of a vertex function.
func VertexAttributeDescriptorFrom(ptr unsafe.Pointer) VertexAttributeDescriptor {
	return VertexAttributeDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VertexAttributeDescriptorClass) Alloc() VertexAttributeDescriptor {
	rv := objc.Send[VertexAttributeDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VertexAttributeDescriptorClass) New() VertexAttributeDescriptor {
	rv := objc.Send[VertexAttributeDescriptor](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VertexAttributeDescriptor) Init() VertexAttributeDescriptor {
	rv := objc.Send[VertexAttributeDescriptor](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VertexAttributeDescriptor) Autorelease() VertexAttributeDescriptor {
	rv := objc.Send[VertexAttributeDescriptor](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVertexAttributeDescriptor creates a new VertexAttributeDescriptor instance.
func NewVertexAttributeDescriptor() VertexAttributeDescriptor {
	return getVertexAttributeDescriptorClass().New()
}


// The index in the argument table for the associated vertex buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/bufferIndex
func (v_ VertexAttributeDescriptor) BufferIndex() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("bufferIndex"))
	return rv
}


// SetBufferIndex sets the value of the bufferIndex property.
// The index in the argument table for the associated vertex buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/bufferIndex
func (v_ VertexAttributeDescriptor) SetBufferIndex(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBufferIndex:"), value)
}

// The format of the vertex attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/format
func (v_ VertexAttributeDescriptor) Format() VertexFormat {
	rv := objc.Send[VertexFormat](v_.ID, objc.Sel("format"))
	return rv
}


// SetFormat sets the value of the format property.
// The format of the vertex attribute.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/format
func (v_ VertexAttributeDescriptor) SetFormat(value VertexFormat) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFormat:"), value)
}

// The location of an attribute in vertex data, determined by the byte offset from the start of the vertex data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/offset
func (v_ VertexAttributeDescriptor) Offset() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("offset"))
	return rv
}


// SetOffset sets the value of the offset property.
// The location of an attribute in vertex data, determined by the byte offset from the start of the vertex data.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/offset
func (v_ VertexAttributeDescriptor) SetOffset(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v_ VertexAttributeDescriptor) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v_.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}



