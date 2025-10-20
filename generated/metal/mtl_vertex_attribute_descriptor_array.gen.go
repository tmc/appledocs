// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VertexAttributeDescriptorArray] class.
var (
	VertexAttributeDescriptorArrayClass     _VertexAttributeDescriptorArrayClass
	VertexAttributeDescriptorArrayClassOnce sync.Once
)

func getVertexAttributeDescriptorArrayClass() _VertexAttributeDescriptorArrayClass {
	VertexAttributeDescriptorArrayClassOnce.Do(func() {
		VertexAttributeDescriptorArrayClass = _VertexAttributeDescriptorArrayClass{objc.GetClass("MTLVertexAttributeDescriptorArray")}
	})
	return VertexAttributeDescriptorArrayClass
}

type _VertexAttributeDescriptorArrayClass struct {
	class objc.Class
}

// An interface definition for the [VertexAttributeDescriptorArray] class.
type IVertexAttributeDescriptorArray interface {
	objectivec.IObject
	SetObjectAtIndexedSubscript(attributeDesc unsafe.Pointer, index uint)
	ObjectAtIndexedSubscript(index uint) unsafe.Pointer
}

// An array of vertex attribute descriptor instances.
//
// An instance is an array of instances that defines how vertex attribute data is formatted and assigned to an index in the attribute argument table. The methods of set or retrieve the attribute formatting information from the array.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptorArray
type VertexAttributeDescriptorArray struct {
	objectivec.Object
}

// VertexAttributeDescriptorArrayFrom constructs a [VertexAttributeDescriptorArray] from an unsafe.Pointer.
//
// An array of vertex attribute descriptor instances.
func VertexAttributeDescriptorArrayFrom(ptr unsafe.Pointer) VertexAttributeDescriptorArray {
	return VertexAttributeDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VertexAttributeDescriptorArrayClass) Alloc() VertexAttributeDescriptorArray {
	rv := objc.Send[VertexAttributeDescriptorArray](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VertexAttributeDescriptorArrayClass) New() VertexAttributeDescriptorArray {
	rv := objc.Send[VertexAttributeDescriptorArray](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VertexAttributeDescriptorArray) Init() VertexAttributeDescriptorArray {
	rv := objc.Send[VertexAttributeDescriptorArray](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VertexAttributeDescriptorArray) Autorelease() VertexAttributeDescriptorArray {
	rv := objc.Send[VertexAttributeDescriptorArray](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVertexAttributeDescriptorArray creates a new VertexAttributeDescriptorArray instance.
func NewVertexAttributeDescriptorArray() VertexAttributeDescriptorArray {
	return getVertexAttributeDescriptorArrayClass().New()
}


// Sets state for the specified vertex attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptorArray/setObject:atIndexedSubscript:
func (v_ VertexAttributeDescriptorArray) SetObjectAtIndexedSubscript(attributeDesc unsafe.Pointer, index uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setObject:atIndexedSubscript:"), attributeDesc, index)
}

// Returns the state of the specified vertex attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptorArray/subscript(_:)
func (v_ VertexAttributeDescriptorArray) ObjectAtIndexedSubscript(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}



