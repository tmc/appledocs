// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTL4ArgumentTableDescriptor] class.
var (
	MTL4ArgumentTableDescriptorClass     _MTL4ArgumentTableDescriptorClass
	MTL4ArgumentTableDescriptorClassOnce sync.Once
)

func getMTL4ArgumentTableDescriptorClass() _MTL4ArgumentTableDescriptorClass {
	MTL4ArgumentTableDescriptorClassOnce.Do(func() {
		MTL4ArgumentTableDescriptorClass = _MTL4ArgumentTableDescriptorClass{objc.GetClass("MTL4ArgumentTableDescriptor")}
	})
	return MTL4ArgumentTableDescriptorClass
}

type _MTL4ArgumentTableDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4ArgumentTableDescriptor] class.
type IMTL4ArgumentTableDescriptor interface {
	objectivec.IObject
}

// Groups parameters for the creation of a Metal argument table.
//
// Argument tables provide resource bindings to your Metal pipeline states.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor
type MTL4ArgumentTableDescriptor struct {
	objectivec.Object
}

// MTL4ArgumentTableDescriptorFrom constructs a [MTL4ArgumentTableDescriptor] from an unsafe.Pointer.
//
// Groups parameters for the creation of a Metal argument table.
func MTL4ArgumentTableDescriptorFrom(ptr unsafe.Pointer) MTL4ArgumentTableDescriptor {
	return MTL4ArgumentTableDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4ArgumentTableDescriptorClass) Alloc() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4ArgumentTableDescriptorClass) New() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4ArgumentTableDescriptor) Init() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4ArgumentTableDescriptor) Autorelease() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4ArgumentTableDescriptor creates a new MTL4ArgumentTableDescriptor instance.
func NewMTL4ArgumentTableDescriptor() MTL4ArgumentTableDescriptor {
	return getMTL4ArgumentTableDescriptorClass().New()
}




