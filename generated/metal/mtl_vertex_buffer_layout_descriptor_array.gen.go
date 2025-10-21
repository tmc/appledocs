// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VertexBufferLayoutDescriptorArray] class.
var (
	VertexBufferLayoutDescriptorArrayClass     _VertexBufferLayoutDescriptorArrayClass
	VertexBufferLayoutDescriptorArrayClassOnce sync.Once
)

func getVertexBufferLayoutDescriptorArrayClass() _VertexBufferLayoutDescriptorArrayClass {
	VertexBufferLayoutDescriptorArrayClassOnce.Do(func() {
		VertexBufferLayoutDescriptorArrayClass = _VertexBufferLayoutDescriptorArrayClass{objc.GetClass("MTLVertexBufferLayoutDescriptorArray")}
	})
	return VertexBufferLayoutDescriptorArrayClass
}

type _VertexBufferLayoutDescriptorArrayClass struct {
	class objc.Class
}

// An interface definition for the [VertexBufferLayoutDescriptorArray] class.
type IVertexBufferLayoutDescriptorArray interface {
	objectivec.IObject
	SetObjectAtIndexedSubscript(bufferDesc unsafe.Pointer, index uint)
	ObjectAtIndexedSubscript(index uint) unsafe.Pointer
}

// An array of vertex buffer layout descriptor instances.
//
// An holds an array of vertex buffer layout states. The methods of set the vertex buffer layout state in the array or retrieve the state from the array.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptorArray
type VertexBufferLayoutDescriptorArray struct {
	objectivec.Object
}

// VertexBufferLayoutDescriptorArrayFrom constructs a [VertexBufferLayoutDescriptorArray] from an unsafe.Pointer.
//
// An array of vertex buffer layout descriptor instances.
func VertexBufferLayoutDescriptorArrayFrom(ptr unsafe.Pointer) VertexBufferLayoutDescriptorArray {
	return VertexBufferLayoutDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VertexBufferLayoutDescriptorArrayClass) Alloc() VertexBufferLayoutDescriptorArray {
	rv := objc.Send[VertexBufferLayoutDescriptorArray](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VertexBufferLayoutDescriptorArrayClass) New() VertexBufferLayoutDescriptorArray {
	rv := objc.Send[VertexBufferLayoutDescriptorArray](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VertexBufferLayoutDescriptorArray) Init() VertexBufferLayoutDescriptorArray {
	rv := objc.Send[VertexBufferLayoutDescriptorArray](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VertexBufferLayoutDescriptorArray) Autorelease() VertexBufferLayoutDescriptorArray {
	rv := objc.Send[VertexBufferLayoutDescriptorArray](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVertexBufferLayoutDescriptorArray creates a new VertexBufferLayoutDescriptorArray instance.
func NewVertexBufferLayoutDescriptorArray() VertexBufferLayoutDescriptorArray {
	return getVertexBufferLayoutDescriptorArrayClass().New()
}


// Sets the state of the specified vertex buffer layout.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptorArray/setObject:atIndexedSubscript:
func (v_ VertexBufferLayoutDescriptorArray) SetObjectAtIndexedSubscript(bufferDesc unsafe.Pointer, index uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setObject:atIndexedSubscript:"), bufferDesc, index)
}

// Returns the state of the specified vertex buffer layout.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptorArray/subscript(_:)
func (v_ VertexBufferLayoutDescriptorArray) ObjectAtIndexedSubscript(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v_ VertexBufferLayoutDescriptorArray) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v_.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}



