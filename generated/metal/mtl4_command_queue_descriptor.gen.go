// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTL4CommandQueueDescriptor] class.
var (
	MTL4CommandQueueDescriptorClass     _MTL4CommandQueueDescriptorClass
	MTL4CommandQueueDescriptorClassOnce sync.Once
)

func getMTL4CommandQueueDescriptorClass() _MTL4CommandQueueDescriptorClass {
	MTL4CommandQueueDescriptorClassOnce.Do(func() {
		MTL4CommandQueueDescriptorClass = _MTL4CommandQueueDescriptorClass{objc.GetClass("MTL4CommandQueueDescriptor")}
	})
	return MTL4CommandQueueDescriptorClass
}

type _MTL4CommandQueueDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4CommandQueueDescriptor] class.
type IMTL4CommandQueueDescriptor interface {
	objectivec.IObject
}

// Groups together parameters for the creation of a new command queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueDescriptor
type MTL4CommandQueueDescriptor struct {
	objectivec.Object
}

// MTL4CommandQueueDescriptorFrom constructs a [MTL4CommandQueueDescriptor] from an unsafe.Pointer.
//
// Groups together parameters for the creation of a new command queue.
func MTL4CommandQueueDescriptorFrom(ptr unsafe.Pointer) MTL4CommandQueueDescriptor {
	return MTL4CommandQueueDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4CommandQueueDescriptorClass) Alloc() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4CommandQueueDescriptorClass) New() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4CommandQueueDescriptor) Init() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4CommandQueueDescriptor) Autorelease() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CommandQueueDescriptor creates a new MTL4CommandQueueDescriptor instance.
func NewMTL4CommandQueueDescriptor() MTL4CommandQueueDescriptor {
	return getMTL4CommandQueueDescriptorClass().New()
}




