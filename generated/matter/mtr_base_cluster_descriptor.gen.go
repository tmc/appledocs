// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterDescriptor] class.
var (
	MTRBaseClusterDescriptorClass     _MTRBaseClusterDescriptorClass
	MTRBaseClusterDescriptorClassOnce sync.Once
)

func getMTRBaseClusterDescriptorClass() _MTRBaseClusterDescriptorClass {
	MTRBaseClusterDescriptorClassOnce.Do(func() {
		MTRBaseClusterDescriptorClass = _MTRBaseClusterDescriptorClass{objc.GetClass("MTRBaseClusterDescriptor")}
	})
	return MTRBaseClusterDescriptorClass
}

type _MTRBaseClusterDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterDescriptor] class.
type IMTRBaseClusterDescriptor interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDescriptor
type MTRBaseClusterDescriptor struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterDescriptorFrom constructs a [MTRBaseClusterDescriptor] from an unsafe.Pointer.
func MTRBaseClusterDescriptorFrom(ptr unsafe.Pointer) MTRBaseClusterDescriptor {
	return MTRBaseClusterDescriptor{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterDescriptorClass) Alloc() MTRBaseClusterDescriptor {
	rv := objc.Send[MTRBaseClusterDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterDescriptorClass) New() MTRBaseClusterDescriptor {
	rv := objc.Send[MTRBaseClusterDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterDescriptor) Init() MTRBaseClusterDescriptor {
	rv := objc.Send[MTRBaseClusterDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterDescriptor) Autorelease() MTRBaseClusterDescriptor {
	rv := objc.Send[MTRBaseClusterDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterDescriptor creates a new MTRBaseClusterDescriptor instance.
func NewMTRBaseClusterDescriptor() MTRBaseClusterDescriptor {
	return getMTRBaseClusterDescriptorClass().New()
}
