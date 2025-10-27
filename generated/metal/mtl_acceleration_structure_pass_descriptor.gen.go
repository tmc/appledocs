// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AccelerationStructurePassDescriptor] class.
var (
	AccelerationStructurePassDescriptorClass     _AccelerationStructurePassDescriptorClass
	AccelerationStructurePassDescriptorClassOnce sync.Once
)

func getAccelerationStructurePassDescriptorClass() _AccelerationStructurePassDescriptorClass {
	AccelerationStructurePassDescriptorClassOnce.Do(func() {
		AccelerationStructurePassDescriptorClass = _AccelerationStructurePassDescriptorClass{objc.GetClass("MTLAccelerationStructurePassDescriptor")}
	})
	return AccelerationStructurePassDescriptorClass
}

type _AccelerationStructurePassDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [AccelerationStructurePassDescriptor] class.
type IAccelerationStructurePassDescriptor interface {
	objectivec.IObject
	

	// properties:
	SampleBufferAttachments() IMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructurePassDescriptorClass) Alloc() AccelerationStructurePassDescriptor {
	rv := objc.Send[AccelerationStructurePassDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccelerationStructurePassDescriptorClass) New() AccelerationStructurePassDescriptor {
	rv := objc.Send[AccelerationStructurePassDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructurePassDescriptor) Init() AccelerationStructurePassDescriptor {
	rv := objc.Send[AccelerationStructurePassDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructurePassDescriptor) Autorelease() AccelerationStructurePassDescriptor {
	rv := objc.Send[AccelerationStructurePassDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructurePassDescriptor creates a new AccelerationStructurePassDescriptor instance.
func NewAccelerationStructurePassDescriptor() AccelerationStructurePassDescriptor {
	return getAccelerationStructurePassDescriptorClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassDescriptor
type AccelerationStructurePassDescriptor struct {
	objectivec.Object
}

// AccelerationStructurePassDescriptorFrom constructs a [AccelerationStructurePassDescriptor] from an unsafe.Pointer.
func AccelerationStructurePassDescriptorFrom(ptr unsafe.Pointer) AccelerationStructurePassDescriptor {
	return AccelerationStructurePassDescriptor{objectivec.Object{objc.ID(ptr)}}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassDescriptor/accelerationStructurePassDescriptor
func (ac _AccelerationStructurePassDescriptorClass) AccelerationStructurePassDescriptor() IAccelerationStructurePassDescriptor {
	rv := objc.Send[AccelerationStructurePassDescriptor](objc.ID(ac.class), objc.Sel("accelerationStructurePassDescriptor"))
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassDescriptor/sampleBufferAttachments
func (a_ AccelerationStructurePassDescriptor) SampleBufferAttachments() IMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptorArray](a_.ID, objc.Sel("sampleBufferAttachments"))
	return rv
}








