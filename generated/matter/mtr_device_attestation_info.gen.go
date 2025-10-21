// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDeviceAttestationInfo] class.
var (
	MTRDeviceAttestationInfoClass     _MTRDeviceAttestationInfoClass
	MTRDeviceAttestationInfoClassOnce sync.Once
)

func getMTRDeviceAttestationInfoClass() _MTRDeviceAttestationInfoClass {
	MTRDeviceAttestationInfoClassOnce.Do(func() {
		MTRDeviceAttestationInfoClass = _MTRDeviceAttestationInfoClass{objc.GetClass("MTRDeviceAttestationInfo")}
	})
	return MTRDeviceAttestationInfoClass
}

type _MTRDeviceAttestationInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceAttestationInfo] class.
type IMTRDeviceAttestationInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationInfo
type MTRDeviceAttestationInfo struct {
	objectivec.Object
}

// MTRDeviceAttestationInfoFrom constructs a [MTRDeviceAttestationInfo] from an unsafe.Pointer.
func MTRDeviceAttestationInfoFrom(ptr unsafe.Pointer) MTRDeviceAttestationInfo {
	return MTRDeviceAttestationInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceAttestationInfoClass) Alloc() MTRDeviceAttestationInfo {
	rv := objc.Send[MTRDeviceAttestationInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceAttestationInfoClass) New() MTRDeviceAttestationInfo {
	rv := objc.Send[MTRDeviceAttestationInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceAttestationInfo) Init() MTRDeviceAttestationInfo {
	rv := objc.Send[MTRDeviceAttestationInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceAttestationInfo) Autorelease() MTRDeviceAttestationInfo {
	rv := objc.Send[MTRDeviceAttestationInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceAttestationInfo creates a new MTRDeviceAttestationInfo instance.
func NewMTRDeviceAttestationInfo() MTRDeviceAttestationInfo {
	return getMTRDeviceAttestationInfoClass().New()
}




