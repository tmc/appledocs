// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRDescriptorClusterDeviceType] class.
var (
	MTRDescriptorClusterDeviceTypeClass     _MTRDescriptorClusterDeviceTypeClass
	MTRDescriptorClusterDeviceTypeClassOnce sync.Once
)

func getMTRDescriptorClusterDeviceTypeClass() _MTRDescriptorClusterDeviceTypeClass {
	MTRDescriptorClusterDeviceTypeClassOnce.Do(func() {
		MTRDescriptorClusterDeviceTypeClass = _MTRDescriptorClusterDeviceTypeClass{objc.GetClass("MTRDescriptorClusterDeviceType")}
	})
	return MTRDescriptorClusterDeviceTypeClass
}

type _MTRDescriptorClusterDeviceTypeClass struct {
	class objc.Class
}

// An interface definition for the [MTRDescriptorClusterDeviceType] class.
type IMTRDescriptorClusterDeviceType interface {
	IMTRDescriptorClusterDeviceTypeStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDescriptorClusterDeviceType
type MTRDescriptorClusterDeviceType struct {
	MTRDescriptorClusterDeviceTypeStruct
}

// MTRDescriptorClusterDeviceTypeFrom constructs a [MTRDescriptorClusterDeviceType] from an unsafe.Pointer.
func MTRDescriptorClusterDeviceTypeFrom(ptr unsafe.Pointer) MTRDescriptorClusterDeviceType {
	return MTRDescriptorClusterDeviceType{
		MTRDescriptorClusterDeviceTypeStruct: MTRDescriptorClusterDeviceTypeStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDescriptorClusterDeviceTypeClass) Alloc() MTRDescriptorClusterDeviceType {
	rv := objc.Send[MTRDescriptorClusterDeviceType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDescriptorClusterDeviceTypeClass) New() MTRDescriptorClusterDeviceType {
	rv := objc.Send[MTRDescriptorClusterDeviceType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDescriptorClusterDeviceType) Init() MTRDescriptorClusterDeviceType {
	rv := objc.Send[MTRDescriptorClusterDeviceType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDescriptorClusterDeviceType) Autorelease() MTRDescriptorClusterDeviceType {
	rv := objc.Send[MTRDescriptorClusterDeviceType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDescriptorClusterDeviceType creates a new MTRDescriptorClusterDeviceType instance.
func NewMTRDescriptorClusterDeviceType() MTRDescriptorClusterDeviceType {
	return getMTRDescriptorClusterDeviceTypeClass().New()
}




