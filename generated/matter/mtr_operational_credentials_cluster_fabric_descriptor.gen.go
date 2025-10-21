// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTROperationalCredentialsClusterFabricDescriptor] class.
var (
	MTROperationalCredentialsClusterFabricDescriptorClass     _MTROperationalCredentialsClusterFabricDescriptorClass
	MTROperationalCredentialsClusterFabricDescriptorClassOnce sync.Once
)

func getMTROperationalCredentialsClusterFabricDescriptorClass() _MTROperationalCredentialsClusterFabricDescriptorClass {
	MTROperationalCredentialsClusterFabricDescriptorClassOnce.Do(func() {
		MTROperationalCredentialsClusterFabricDescriptorClass = _MTROperationalCredentialsClusterFabricDescriptorClass{objc.GetClass("MTROperationalCredentialsClusterFabricDescriptor")}
	})
	return MTROperationalCredentialsClusterFabricDescriptorClass
}

type _MTROperationalCredentialsClusterFabricDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterFabricDescriptor] class.
type IMTROperationalCredentialsClusterFabricDescriptor interface {
	IMTROperationalCredentialsClusterFabricDescriptorStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptor
type MTROperationalCredentialsClusterFabricDescriptor struct {
	MTROperationalCredentialsClusterFabricDescriptorStruct
}

// MTROperationalCredentialsClusterFabricDescriptorFrom constructs a [MTROperationalCredentialsClusterFabricDescriptor] from an unsafe.Pointer.
func MTROperationalCredentialsClusterFabricDescriptorFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterFabricDescriptor {
	return MTROperationalCredentialsClusterFabricDescriptor{
		MTROperationalCredentialsClusterFabricDescriptorStruct: MTROperationalCredentialsClusterFabricDescriptorStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterFabricDescriptorClass) Alloc() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterFabricDescriptorClass) New() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterFabricDescriptor) Init() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterFabricDescriptor) Autorelease() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterFabricDescriptor creates a new MTROperationalCredentialsClusterFabricDescriptor instance.
func NewMTROperationalCredentialsClusterFabricDescriptor() MTROperationalCredentialsClusterFabricDescriptor {
	return getMTROperationalCredentialsClusterFabricDescriptorClass().New()
}




