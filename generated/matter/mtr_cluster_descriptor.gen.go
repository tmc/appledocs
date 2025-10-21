// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterDescriptor] class.
var (
	MTRClusterDescriptorClass     _MTRClusterDescriptorClass
	MTRClusterDescriptorClassOnce sync.Once
)

func getMTRClusterDescriptorClass() _MTRClusterDescriptorClass {
	MTRClusterDescriptorClassOnce.Do(func() {
		MTRClusterDescriptorClass = _MTRClusterDescriptorClass{objc.GetClass("MTRClusterDescriptor")}
	})
	return MTRClusterDescriptorClass
}

type _MTRClusterDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterDescriptor] class.
type IMTRClusterDescriptor interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDescriptor
type MTRClusterDescriptor struct {
	MTRGenericCluster
}

// MTRClusterDescriptorFrom constructs a [MTRClusterDescriptor] from an unsafe.Pointer.
func MTRClusterDescriptorFrom(ptr unsafe.Pointer) MTRClusterDescriptor {
	return MTRClusterDescriptor{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterDescriptorClass) Alloc() MTRClusterDescriptor {
	rv := objc.Send[MTRClusterDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterDescriptorClass) New() MTRClusterDescriptor {
	rv := objc.Send[MTRClusterDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterDescriptor) Init() MTRClusterDescriptor {
	rv := objc.Send[MTRClusterDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterDescriptor) Autorelease() MTRClusterDescriptor {
	rv := objc.Send[MTRClusterDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterDescriptor creates a new MTRClusterDescriptor instance.
func NewMTRClusterDescriptor() MTRClusterDescriptor {
	return getMTRClusterDescriptorClass().New()
}




