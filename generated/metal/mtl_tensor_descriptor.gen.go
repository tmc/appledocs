// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTLTensorDescriptor] class.
var mTLTensorDescriptorClass = _MTLTensorDescriptorClass{objc.GetClass("MTLTensorDescriptor")}

type _MTLTensorDescriptorClass struct {
	class objc.Class
}

// A configuration type for creating new tensor instances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor

type MTLTensorDescriptor struct {
	objectivec.Object
}

// MTLTensorDescriptorFrom constructs a [MTLTensorDescriptor] from an unsafe.Pointer.
//
// A configuration type for creating new tensor instances.
func MTLTensorDescriptorFrom(ptr unsafe.Pointer) MTLTensorDescriptor {
	return MTLTensorDescriptor{objectivec.Object{objc.ID(ptr)}}
}



