// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDescriptorClusterDeviceTypeStruct] class.
var (
	MTRDescriptorClusterDeviceTypeStructClass     _MTRDescriptorClusterDeviceTypeStructClass
	MTRDescriptorClusterDeviceTypeStructClassOnce sync.Once
)

func getMTRDescriptorClusterDeviceTypeStructClass() _MTRDescriptorClusterDeviceTypeStructClass {
	MTRDescriptorClusterDeviceTypeStructClassOnce.Do(func() {
		MTRDescriptorClusterDeviceTypeStructClass = _MTRDescriptorClusterDeviceTypeStructClass{objc.GetClass("MTRDescriptorClusterDeviceTypeStruct")}
	})
	return MTRDescriptorClusterDeviceTypeStructClass
}

type _MTRDescriptorClusterDeviceTypeStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRDescriptorClusterDeviceTypeStruct] class.
type IMTRDescriptorClusterDeviceTypeStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDescriptorClusterDeviceTypeStruct
type MTRDescriptorClusterDeviceTypeStruct struct {
	objectivec.Object
}

// MTRDescriptorClusterDeviceTypeStructFrom constructs a [MTRDescriptorClusterDeviceTypeStruct] from an unsafe.Pointer.
func MTRDescriptorClusterDeviceTypeStructFrom(ptr unsafe.Pointer) MTRDescriptorClusterDeviceTypeStruct {
	return MTRDescriptorClusterDeviceTypeStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDescriptorClusterDeviceTypeStructClass) Alloc() MTRDescriptorClusterDeviceTypeStruct {
	rv := objc.Send[MTRDescriptorClusterDeviceTypeStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDescriptorClusterDeviceTypeStructClass) New() MTRDescriptorClusterDeviceTypeStruct {
	rv := objc.Send[MTRDescriptorClusterDeviceTypeStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDescriptorClusterDeviceTypeStruct) Init() MTRDescriptorClusterDeviceTypeStruct {
	rv := objc.Send[MTRDescriptorClusterDeviceTypeStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDescriptorClusterDeviceTypeStruct) Autorelease() MTRDescriptorClusterDeviceTypeStruct {
	rv := objc.Send[MTRDescriptorClusterDeviceTypeStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDescriptorClusterDeviceTypeStruct creates a new MTRDescriptorClusterDeviceTypeStruct instance.
func NewMTRDescriptorClusterDeviceTypeStruct() MTRDescriptorClusterDeviceTypeStruct {
	return getMTRDescriptorClusterDeviceTypeStructClass().New()
}




