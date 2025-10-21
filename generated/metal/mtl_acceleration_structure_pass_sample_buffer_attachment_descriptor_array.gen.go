// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccelerationStructurePassSampleBufferAttachmentDescriptorArray] class.
var (
	AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass     _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass
	AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass() _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass {
	AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass = _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray")}
	})
	return AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass
}

type _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructurePassSampleBufferAttachmentDescriptorArray] class.
type IAccelerationStructurePassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray
type AccelerationStructurePassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// AccelerationStructurePassSampleBufferAttachmentDescriptorArrayFrom constructs a [AccelerationStructurePassSampleBufferAttachmentDescriptorArray] from an unsafe.Pointer.
func AccelerationStructurePassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	return AccelerationStructurePassSampleBufferAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass) Alloc() AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptorArray](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass) New() AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptorArray](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptorArray) Init() AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptorArray](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptorArray) Autorelease() AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptorArray](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructurePassSampleBufferAttachmentDescriptorArray creates a new AccelerationStructurePassSampleBufferAttachmentDescriptorArray instance.
func NewAccelerationStructurePassSampleBufferAttachmentDescriptorArray() AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	return getAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass().New()
}




