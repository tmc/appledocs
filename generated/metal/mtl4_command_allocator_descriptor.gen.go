// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTL4CommandAllocatorDescriptor] class.
var (
	MTL4CommandAllocatorDescriptorClass     _MTL4CommandAllocatorDescriptorClass
	MTL4CommandAllocatorDescriptorClassOnce sync.Once
)

func getMTL4CommandAllocatorDescriptorClass() _MTL4CommandAllocatorDescriptorClass {
	MTL4CommandAllocatorDescriptorClassOnce.Do(func() {
		MTL4CommandAllocatorDescriptorClass = _MTL4CommandAllocatorDescriptorClass{objc.GetClass("MTL4CommandAllocatorDescriptor")}
	})
	return MTL4CommandAllocatorDescriptorClass
}

type _MTL4CommandAllocatorDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4CommandAllocatorDescriptor] class.
type IMTL4CommandAllocatorDescriptor interface {
	objectivec.IObject
	Label() string
	SetLabel(value string)
	MTL4CommandQueueErrorDomain() string
}

// Groups together parameters for creating a command allocator.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandAllocatorDescriptor
type MTL4CommandAllocatorDescriptor struct {
	objectivec.Object
}

// MTL4CommandAllocatorDescriptorFrom constructs a [MTL4CommandAllocatorDescriptor] from an unsafe.Pointer.
//
// Groups together parameters for creating a command allocator.
func MTL4CommandAllocatorDescriptorFrom(ptr unsafe.Pointer) MTL4CommandAllocatorDescriptor {
	return MTL4CommandAllocatorDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4CommandAllocatorDescriptorClass) Alloc() MTL4CommandAllocatorDescriptor {
	rv := objc.Send[MTL4CommandAllocatorDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4CommandAllocatorDescriptorClass) New() MTL4CommandAllocatorDescriptor {
	rv := objc.Send[MTL4CommandAllocatorDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4CommandAllocatorDescriptor) Init() MTL4CommandAllocatorDescriptor {
	rv := objc.Send[MTL4CommandAllocatorDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4CommandAllocatorDescriptor) Autorelease() MTL4CommandAllocatorDescriptor {
	rv := objc.Send[MTL4CommandAllocatorDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CommandAllocatorDescriptor creates a new MTL4CommandAllocatorDescriptor instance.
func NewMTL4CommandAllocatorDescriptor() MTL4CommandAllocatorDescriptor {
	return getMTL4CommandAllocatorDescriptorClass().New()
}


// An optional label you can assign to the command allocator to aid debugging.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandallocatordescriptor/label
func (m_ MTL4CommandAllocatorDescriptor) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// An optional label you can assign to the command allocator to aid debugging.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandallocatordescriptor/label
func (m_ MTL4CommandAllocatorDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m_ MTL4CommandAllocatorDescriptor) MTL4CommandQueueErrorDomain() string {
	rv := objc.Send[string](m_.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return rv
}



