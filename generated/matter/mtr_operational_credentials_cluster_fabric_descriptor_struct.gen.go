// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterFabricDescriptorStruct] class.
var (
	MTROperationalCredentialsClusterFabricDescriptorStructClass     _MTROperationalCredentialsClusterFabricDescriptorStructClass
	MTROperationalCredentialsClusterFabricDescriptorStructClassOnce sync.Once
)

func getMTROperationalCredentialsClusterFabricDescriptorStructClass() _MTROperationalCredentialsClusterFabricDescriptorStructClass {
	MTROperationalCredentialsClusterFabricDescriptorStructClassOnce.Do(func() {
		MTROperationalCredentialsClusterFabricDescriptorStructClass = _MTROperationalCredentialsClusterFabricDescriptorStructClass{objc.GetClass("MTROperationalCredentialsClusterFabricDescriptorStruct")}
	})
	return MTROperationalCredentialsClusterFabricDescriptorStructClass
}

type _MTROperationalCredentialsClusterFabricDescriptorStructClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterFabricDescriptorStruct] class.
type IMTROperationalCredentialsClusterFabricDescriptorStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct
type MTROperationalCredentialsClusterFabricDescriptorStruct struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterFabricDescriptorStructFrom constructs a [MTROperationalCredentialsClusterFabricDescriptorStruct] from an unsafe.Pointer.
func MTROperationalCredentialsClusterFabricDescriptorStructFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterFabricDescriptorStruct {
	return MTROperationalCredentialsClusterFabricDescriptorStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterFabricDescriptorStructClass) Alloc() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterFabricDescriptorStructClass) New() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) Init() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) Autorelease() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterFabricDescriptorStruct creates a new MTROperationalCredentialsClusterFabricDescriptorStruct instance.
func NewMTROperationalCredentialsClusterFabricDescriptorStruct() MTROperationalCredentialsClusterFabricDescriptorStruct {
	return getMTROperationalCredentialsClusterFabricDescriptorStructClass().New()
}




