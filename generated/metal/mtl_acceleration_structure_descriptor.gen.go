// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccelerationStructureDescriptor] class.
var (
	AccelerationStructureDescriptorClass     _AccelerationStructureDescriptorClass
	AccelerationStructureDescriptorClassOnce sync.Once
)

func getAccelerationStructureDescriptorClass() _AccelerationStructureDescriptorClass {
	AccelerationStructureDescriptorClassOnce.Do(func() {
		AccelerationStructureDescriptorClass = _AccelerationStructureDescriptorClass{objc.GetClass("MTLAccelerationStructureDescriptor")}
	})
	return AccelerationStructureDescriptorClass
}

type _AccelerationStructureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructureDescriptor] class.
type IAccelerationStructureDescriptor interface {
	objectivec.IObject
	Usage() AccelerationStructureUsage
	SetUsage(value IAccelerationStructureUsage)
}

// A base class for classes that define the configuration for a new acceleration structure.
//
// This is the base class for other acceleration structure descriptors. Don’t use this class directly. Use one of the derived classes instead, as describes.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureDescriptor
type AccelerationStructureDescriptor struct {
	objectivec.Object
}

// AccelerationStructureDescriptorFrom constructs a [AccelerationStructureDescriptor] from an unsafe.Pointer.
//
// A base class for classes that define the configuration for a new acceleration structure.
func AccelerationStructureDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureDescriptor {
	return AccelerationStructureDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureDescriptorClass) Alloc() AccelerationStructureDescriptor {
	rv := objc.Send[AccelerationStructureDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureDescriptorClass) New() AccelerationStructureDescriptor {
	rv := objc.Send[AccelerationStructureDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureDescriptor) Init() AccelerationStructureDescriptor {
	rv := objc.Send[AccelerationStructureDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureDescriptor) Autorelease() AccelerationStructureDescriptor {
	rv := objc.Send[AccelerationStructureDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureDescriptor creates a new AccelerationStructureDescriptor instance.
func NewAccelerationStructureDescriptor() AccelerationStructureDescriptor {
	return getAccelerationStructureDescriptorClass().New()
}


// The options that describe how you intend to use the acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuredescriptor/usage
func (a_ AccelerationStructureDescriptor) Usage() AccelerationStructureUsage {
	rv := objc.Send[AccelerationStructureUsage](a_.ID, objc.Sel("usage"))
	return rv
}


// SetUsage sets the value of the usage property.
// The options that describe how you intend to use the acceleration structure.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuredescriptor/usage
func (a_ AccelerationStructureDescriptor) SetUsage(value IAccelerationStructureUsage) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUsage:"), value)
}



