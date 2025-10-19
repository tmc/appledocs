// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTLRenderPassDescriptor] class.
var mTLRenderPassDescriptorClass = _MTLRenderPassDescriptorClass{objc.GetClass("MTLRenderPassDescriptor")}

type _MTLRenderPassDescriptorClass struct {
	class objc.Class
}

// A group of render targets that hold the results of a render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassDescriptor

type MTLRenderPassDescriptor struct {
	objectivec.Object
}

// MTLRenderPassDescriptorFrom constructs a [MTLRenderPassDescriptor] from an unsafe.Pointer.
//
// A group of render targets that hold the results of a render pass.
func MTLRenderPassDescriptorFrom(ptr unsafe.Pointer) MTLRenderPassDescriptor {
	return MTLRenderPassDescriptor{objectivec.Object{objc.ID(ptr)}}
}



