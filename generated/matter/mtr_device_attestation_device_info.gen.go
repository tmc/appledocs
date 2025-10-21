// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDeviceAttestationDeviceInfo] class.
var (
	MTRDeviceAttestationDeviceInfoClass     _MTRDeviceAttestationDeviceInfoClass
	MTRDeviceAttestationDeviceInfoClassOnce sync.Once
)

func getMTRDeviceAttestationDeviceInfoClass() _MTRDeviceAttestationDeviceInfoClass {
	MTRDeviceAttestationDeviceInfoClassOnce.Do(func() {
		MTRDeviceAttestationDeviceInfoClass = _MTRDeviceAttestationDeviceInfoClass{objc.GetClass("MTRDeviceAttestationDeviceInfo")}
	})
	return MTRDeviceAttestationDeviceInfoClass
}

type _MTRDeviceAttestationDeviceInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceAttestationDeviceInfo] class.
type IMTRDeviceAttestationDeviceInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceAttestationDeviceInfo
type MTRDeviceAttestationDeviceInfo struct {
	objectivec.Object
}

// MTRDeviceAttestationDeviceInfoFrom constructs a [MTRDeviceAttestationDeviceInfo] from an unsafe.Pointer.
func MTRDeviceAttestationDeviceInfoFrom(ptr unsafe.Pointer) MTRDeviceAttestationDeviceInfo {
	return MTRDeviceAttestationDeviceInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceAttestationDeviceInfoClass) Alloc() MTRDeviceAttestationDeviceInfo {
	rv := objc.Send[MTRDeviceAttestationDeviceInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceAttestationDeviceInfoClass) New() MTRDeviceAttestationDeviceInfo {
	rv := objc.Send[MTRDeviceAttestationDeviceInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceAttestationDeviceInfo) Init() MTRDeviceAttestationDeviceInfo {
	rv := objc.Send[MTRDeviceAttestationDeviceInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceAttestationDeviceInfo) Autorelease() MTRDeviceAttestationDeviceInfo {
	rv := objc.Send[MTRDeviceAttestationDeviceInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceAttestationDeviceInfo creates a new MTRDeviceAttestationDeviceInfo instance.
func NewMTRDeviceAttestationDeviceInfo() MTRDeviceAttestationDeviceInfo {
	return getMTRDeviceAttestationDeviceInfoClass().New()
}




