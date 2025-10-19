// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTLVertexDescriptor] class.
var mTLVertexDescriptorClass = _MTLVertexDescriptorClass{objc.GetClass("MTLVertexDescriptor")}

type _MTLVertexDescriptorClass struct {
	class objc.Class
}

// An object that describes how to organize and map data to a vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor

type MTLVertexDescriptor struct {
	objectivec.Object
}

// MTLVertexDescriptorFrom constructs a [MTLVertexDescriptor] from an unsafe.Pointer.
//
// An object that describes how to organize and map data to a vertex function.
func MTLVertexDescriptorFrom(ptr unsafe.Pointer) MTLVertexDescriptor {
	return MTLVertexDescriptor{objectivec.Object{objc.ID(ptr)}}
}



