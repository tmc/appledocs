// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralCommissioningClusterBasicCommissioningInfo] class.
var (
	MTRGeneralCommissioningClusterBasicCommissioningInfoClass     _MTRGeneralCommissioningClusterBasicCommissioningInfoClass
	MTRGeneralCommissioningClusterBasicCommissioningInfoClassOnce sync.Once
)

func getMTRGeneralCommissioningClusterBasicCommissioningInfoClass() _MTRGeneralCommissioningClusterBasicCommissioningInfoClass {
	MTRGeneralCommissioningClusterBasicCommissioningInfoClassOnce.Do(func() {
		MTRGeneralCommissioningClusterBasicCommissioningInfoClass = _MTRGeneralCommissioningClusterBasicCommissioningInfoClass{objc.GetClass("MTRGeneralCommissioningClusterBasicCommissioningInfo")}
	})
	return MTRGeneralCommissioningClusterBasicCommissioningInfoClass
}

type _MTRGeneralCommissioningClusterBasicCommissioningInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralCommissioningClusterBasicCommissioningInfo] class.
type IMTRGeneralCommissioningClusterBasicCommissioningInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralCommissioningClusterBasicCommissioningInfo
type MTRGeneralCommissioningClusterBasicCommissioningInfo struct {
	objectivec.Object
}

// MTRGeneralCommissioningClusterBasicCommissioningInfoFrom constructs a [MTRGeneralCommissioningClusterBasicCommissioningInfo] from an unsafe.Pointer.
func MTRGeneralCommissioningClusterBasicCommissioningInfoFrom(ptr unsafe.Pointer) MTRGeneralCommissioningClusterBasicCommissioningInfo {
	return MTRGeneralCommissioningClusterBasicCommissioningInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralCommissioningClusterBasicCommissioningInfoClass) Alloc() MTRGeneralCommissioningClusterBasicCommissioningInfo {
	rv := objc.Send[MTRGeneralCommissioningClusterBasicCommissioningInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralCommissioningClusterBasicCommissioningInfoClass) New() MTRGeneralCommissioningClusterBasicCommissioningInfo {
	rv := objc.Send[MTRGeneralCommissioningClusterBasicCommissioningInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralCommissioningClusterBasicCommissioningInfo) Init() MTRGeneralCommissioningClusterBasicCommissioningInfo {
	rv := objc.Send[MTRGeneralCommissioningClusterBasicCommissioningInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralCommissioningClusterBasicCommissioningInfo) Autorelease() MTRGeneralCommissioningClusterBasicCommissioningInfo {
	rv := objc.Send[MTRGeneralCommissioningClusterBasicCommissioningInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralCommissioningClusterBasicCommissioningInfo creates a new MTRGeneralCommissioningClusterBasicCommissioningInfo instance.
func NewMTRGeneralCommissioningClusterBasicCommissioningInfo() MTRGeneralCommissioningClusterBasicCommissioningInfo {
	return getMTRGeneralCommissioningClusterBasicCommissioningInfoClass().New()
}




