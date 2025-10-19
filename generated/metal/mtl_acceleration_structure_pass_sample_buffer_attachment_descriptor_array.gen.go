// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray] class.
var mTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass = _MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray")}

type _MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray

type MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayFrom constructs a [MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray] from an unsafe.Pointer.
func MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	return MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}



