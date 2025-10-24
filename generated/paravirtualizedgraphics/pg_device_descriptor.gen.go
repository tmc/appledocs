// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PGDeviceDescriptor */


/* debug [class_header]: Header for PGDeviceDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PGDeviceDescriptor */
// An interface definition for the [PGDeviceDescriptor] class.
type IPGDeviceDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PGDeviceDescriptor */
	// properties:
	AddTraceRange() unsafe.Pointer
	SetAddTraceRange(value unsafe.Pointer)
	CreateTask() unsafe.Pointer
	SetCreateTask(value unsafe.Pointer)
	DestroyTask() unsafe.Pointer
	SetDestroyTask(value unsafe.Pointer)
	Device() unsafe.Pointer
	SetDevice(value unsafe.Pointer)
	DisplayPortCount() uint32 /* not a class type */
	SetDisplayPortCount(value uint32 /* not a class type */)
	MapMemory() unsafe.Pointer
	SetMapMemory(value unsafe.Pointer)
	MmioLength() uintptr /* not a class type */
	SetMmioLength(value uintptr /* not a class type */)
	RaiseInterrupt() unsafe.Pointer
	SetRaiseInterrupt(value unsafe.Pointer)
	ReadMemory() unsafe.Pointer
	SetReadMemory(value unsafe.Pointer)
	RemoveTraceRange() unsafe.Pointer
	SetRemoveTraceRange(value unsafe.Pointer)
	UnmapMemory() unsafe.Pointer
	SetUnmapMemory(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PGDeviceDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PGDeviceDescriptor */
// Alloc allocates a new instance without initialization.
func (pc _PGDeviceDescriptorClass) Alloc() PGDeviceDescriptor {
	rv := objc.Send[PGDeviceDescriptor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PGDeviceDescriptor */
// A description of the paravirtualized graphics device to create.


// A description of the paravirtualized graphics device to create.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PGDeviceDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PGDeviceDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PGDeviceDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PGDeviceDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PGDeviceDescriptor */

// A handler that the framework calls to add a trace range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/addTraceRange
func (p_ PGDeviceDescriptor) AddTraceRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("addTraceRange"))
	return rv
}/* debug [instance_properties/getter]: addTraceRange */


// A handler that the framework calls to add a trace range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/addTraceRange
func (p_ PGDeviceDescriptor) SetAddTraceRange(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAddTraceRange:"), value)
}/* debug [instance_properties/setter]: addTraceRange */


// A handler that the framework calls to create a task object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/createTask
func (p_ PGDeviceDescriptor) CreateTask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("createTask"))
	return rv
}/* debug [instance_properties/getter]: createTask */


// A handler that the framework calls to create a task object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/createTask
func (p_ PGDeviceDescriptor) SetCreateTask(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCreateTask:"), value)
}/* debug [instance_properties/setter]: createTask */


// A handler that the framework calls to destroy a task object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/destroyTask
func (p_ PGDeviceDescriptor) DestroyTask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("destroyTask"))
	return rv
}/* debug [instance_properties/getter]: destroyTask */


// A handler that the framework calls to destroy a task object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/destroyTask
func (p_ PGDeviceDescriptor) SetDestroyTask(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDestroyTask:"), value)
}/* debug [instance_properties/setter]: destroyTask */


// The Metal device object to use to back the virtual graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/device
func (p_ PGDeviceDescriptor) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The Metal device object to use to back the virtual graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/device
func (p_ PGDeviceDescriptor) SetDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/displayPortCount
func (p_ PGDeviceDescriptor) DisplayPortCount() uint32 /* not a class type */ {
	rv := objc.Send[uint32](p_.ID, objc.Sel("displayPortCount"))
	return rv
}/* debug [instance_properties/getter]: displayPortCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/displayPortCount
func (p_ PGDeviceDescriptor) SetDisplayPortCount(value uint32 /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayPortCount:"), value)
}/* debug [instance_properties/setter]: displayPortCount */


// A handler that the framework calls to map memory into the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/mapMemory
func (p_ PGDeviceDescriptor) MapMemory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mapMemory"))
	return rv
}/* debug [instance_properties/getter]: mapMemory */


// A handler that the framework calls to map memory into the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/mapMemory
func (p_ PGDeviceDescriptor) SetMapMemory(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMapMemory:"), value)
}/* debug [instance_properties/setter]: mapMemory */


// The length in bytes of the memory-mapped IO section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/mmioLength
func (p_ PGDeviceDescriptor) MmioLength() uintptr /* not a class type */ {
	rv := objc.Send[uintptr](p_.ID, objc.Sel("mmioLength"))
	return rv
}/* debug [instance_properties/getter]: mmioLength */


// The length in bytes of the memory-mapped IO section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/mmioLength
func (p_ PGDeviceDescriptor) SetMmioLength(value uintptr /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMmioLength:"), value)
}/* debug [instance_properties/setter]: mmioLength */


// A handler that the system calls to raise an interrupt in the guest environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/raiseInterrupt
func (p_ PGDeviceDescriptor) RaiseInterrupt() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("raiseInterrupt"))
	return rv
}/* debug [instance_properties/getter]: raiseInterrupt */


// A handler that the system calls to raise an interrupt in the guest environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/raiseInterrupt
func (p_ PGDeviceDescriptor) SetRaiseInterrupt(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRaiseInterrupt:"), value)
}/* debug [instance_properties/setter]: raiseInterrupt */


// A handler that the framework calls to read data from the guest’s memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/readMemory
func (p_ PGDeviceDescriptor) ReadMemory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("readMemory"))
	return rv
}/* debug [instance_properties/getter]: readMemory */


// A handler that the framework calls to read data from the guest’s memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/readMemory
func (p_ PGDeviceDescriptor) SetReadMemory(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReadMemory:"), value)
}/* debug [instance_properties/setter]: readMemory */


// A handler that the framework calls to remove a trace range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/removeTraceRange
func (p_ PGDeviceDescriptor) RemoveTraceRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("removeTraceRange"))
	return rv
}/* debug [instance_properties/getter]: removeTraceRange */


// A handler that the framework calls to remove a trace range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/removeTraceRange
func (p_ PGDeviceDescriptor) SetRemoveTraceRange(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRemoveTraceRange:"), value)
}/* debug [instance_properties/setter]: removeTraceRange */


// A handler that the framework calls to unmap memory from the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/unmapMemory
func (p_ PGDeviceDescriptor) UnmapMemory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("unmapMemory"))
	return rv
}/* debug [instance_properties/getter]: unmapMemory */


// A handler that the framework calls to unmap memory from the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDeviceDescriptor/unmapMemory
func (p_ PGDeviceDescriptor) SetUnmapMemory(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUnmapMemory:"), value)
}/* debug [instance_properties/setter]: unmapMemory */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PGDeviceDescriptor */



