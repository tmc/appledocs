// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PGDeviceDescriptor] class.
var (
	PGDeviceDescriptorClass     _PGDeviceDescriptorClass
	PGDeviceDescriptorClassOnce sync.Once
)

func getPGDeviceDescriptorClass() _PGDeviceDescriptorClass {
	PGDeviceDescriptorClassOnce.Do(func() {
		PGDeviceDescriptorClass = _PGDeviceDescriptorClass{objc.GetClass("PGDeviceDescriptor")}
	})
	return PGDeviceDescriptorClass
}

type _PGDeviceDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [PGDeviceDescriptor] class.
type IPGDeviceDescriptor interface {
	objectivec.IObject
}

// A description of the paravirtualized graphics device to create.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor
type PGDeviceDescriptor struct {
	objectivec.Object
}

// PGDeviceDescriptorFrom constructs a [PGDeviceDescriptor] from an unsafe.Pointer.
//
// A description of the paravirtualized graphics device to create.
func PGDeviceDescriptorFrom(ptr unsafe.Pointer) PGDeviceDescriptor {
	return PGDeviceDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PGDeviceDescriptorClass) Alloc() PGDeviceDescriptor {
	rv := objc.Send[PGDeviceDescriptor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PGDeviceDescriptorClass) New() PGDeviceDescriptor {
	rv := objc.Send[PGDeviceDescriptor](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PGDeviceDescriptor) Init() PGDeviceDescriptor {
	rv := objc.Send[PGDeviceDescriptor](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PGDeviceDescriptor) Autorelease() PGDeviceDescriptor {
	rv := objc.Send[PGDeviceDescriptor](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPGDeviceDescriptor creates a new PGDeviceDescriptor instance.
func NewPGDeviceDescriptor() PGDeviceDescriptor {
	return getPGDeviceDescriptorClass().New()
}


// A handler that the framework calls to add a trace range.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/addTraceRange
func (p_ PGDeviceDescriptor) AddTraceRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("addTraceRange"))
	return rv
}


// SetAddTraceRange sets the value of the addTraceRange property.
// A handler that the framework calls to add a trace range.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/addTraceRange
func (p_ PGDeviceDescriptor) SetAddTraceRange(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAddTraceRange:"), value)
}

// A handler that the framework calls to create a task object.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/createTask
func (p_ PGDeviceDescriptor) CreateTask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("createTask"))
	return rv
}


// SetCreateTask sets the value of the createTask property.
// A handler that the framework calls to create a task object.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/createTask
func (p_ PGDeviceDescriptor) SetCreateTask(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCreateTask:"), value)
}

// A handler that the framework calls to destroy a task object.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/destroyTask
func (p_ PGDeviceDescriptor) DestroyTask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("destroyTask"))
	return rv
}


// SetDestroyTask sets the value of the destroyTask property.
// A handler that the framework calls to destroy a task object.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/destroyTask
func (p_ PGDeviceDescriptor) SetDestroyTask(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDestroyTask:"), value)
}

// The Metal device object to use to back the virtual graphics device.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/device
func (p_ PGDeviceDescriptor) Device() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("device"))
	return rv
}


// SetDevice sets the value of the device property.
// The Metal device object to use to back the virtual graphics device.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/device
func (p_ PGDeviceDescriptor) SetDevice(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDevice:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/displayPortCount
func (p_ PGDeviceDescriptor) DisplayPortCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("displayPortCount"))
	return rv
}


// SetDisplayPortCount sets the value of the displayPortCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/displayPortCount
func (p_ PGDeviceDescriptor) SetDisplayPortCount(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayPortCount:"), value)
}

// A handler that the framework calls to map memory into the virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/mapMemory
func (p_ PGDeviceDescriptor) MapMemory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mapMemory"))
	return rv
}


// SetMapMemory sets the value of the mapMemory property.
// A handler that the framework calls to map memory into the virtual machine.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/mapMemory
func (p_ PGDeviceDescriptor) SetMapMemory(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMapMemory:"), value)
}

// The length in bytes of the memory-mapped IO section.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/mmioLength
func (p_ PGDeviceDescriptor) MmioLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mmioLength"))
	return rv
}


// SetMmioLength sets the value of the mmioLength property.
// The length in bytes of the memory-mapped IO section.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/mmioLength
func (p_ PGDeviceDescriptor) SetMmioLength(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMmioLength:"), value)
}

// A handler that the system calls to raise an interrupt in the guest environment.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/raiseInterrupt
func (p_ PGDeviceDescriptor) RaiseInterrupt() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("raiseInterrupt"))
	return rv
}


// SetRaiseInterrupt sets the value of the raiseInterrupt property.
// A handler that the system calls to raise an interrupt in the guest environment.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/raiseInterrupt
func (p_ PGDeviceDescriptor) SetRaiseInterrupt(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRaiseInterrupt:"), value)
}

// A handler that the framework calls to read data from the guest’s memory.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/readMemory
func (p_ PGDeviceDescriptor) ReadMemory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("readMemory"))
	return rv
}


// SetReadMemory sets the value of the readMemory property.
// A handler that the framework calls to read data from the guest’s memory.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/readMemory
func (p_ PGDeviceDescriptor) SetReadMemory(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReadMemory:"), value)
}

// A handler that the framework calls to remove a trace range.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/removeTraceRange
func (p_ PGDeviceDescriptor) RemoveTraceRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("removeTraceRange"))
	return rv
}


// SetRemoveTraceRange sets the value of the removeTraceRange property.
// A handler that the framework calls to remove a trace range.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/removeTraceRange
func (p_ PGDeviceDescriptor) SetRemoveTraceRange(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRemoveTraceRange:"), value)
}

// A handler that the framework calls to unmap memory from the virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/unmapMemory
func (p_ PGDeviceDescriptor) UnmapMemory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("unmapMemory"))
	return rv
}


// SetUnmapMemory sets the value of the unmapMemory property.
// A handler that the framework calls to unmap memory from the virtual machine.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/unmapMemory
func (p_ PGDeviceDescriptor) SetUnmapMemory(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUnmapMemory:"), value)
}



