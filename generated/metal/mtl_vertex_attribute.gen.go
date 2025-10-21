// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VertexAttribute] class.
var (
	VertexAttributeClass     _VertexAttributeClass
	VertexAttributeClassOnce sync.Once
)

func getVertexAttributeClass() _VertexAttributeClass {
	VertexAttributeClassOnce.Do(func() {
		VertexAttributeClass = _VertexAttributeClass{objc.GetClass("MTLVertexAttribute")}
	})
	return VertexAttributeClass
}

type _VertexAttributeClass struct {
	class objc.Class
}

// An interface definition for the [VertexAttribute] class.
type IVertexAttribute interface {
	objectivec.IObject
}

// An instance that represents an attribute of a vertex function.
//
// An instance represents an attribute for per-vertex input in a vertex function. You use vertex attribute instances to inspect the inputs of a vertex function by examining the property of the corresponding instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttribute
type VertexAttribute struct {
	objectivec.Object
}

// VertexAttributeFrom constructs a [VertexAttribute] from an unsafe.Pointer.
//
// An instance that represents an attribute of a vertex function.
func VertexAttributeFrom(ptr unsafe.Pointer) VertexAttribute {
	return VertexAttribute{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VertexAttributeClass) Alloc() VertexAttribute {
	rv := objc.Send[VertexAttribute](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VertexAttributeClass) New() VertexAttribute {
	rv := objc.Send[VertexAttribute](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VertexAttribute) Init() VertexAttribute {
	rv := objc.Send[VertexAttribute](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VertexAttribute) Autorelease() VertexAttribute {
	rv := objc.Send[VertexAttribute](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVertexAttribute creates a new VertexAttribute instance.
func NewVertexAttribute() VertexAttribute {
	return getVertexAttributeClass().New()
}




