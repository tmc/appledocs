// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTimeSynchronizationClusterTimeZoneType] class.
var (
	MTRTimeSynchronizationClusterTimeZoneTypeClass     _MTRTimeSynchronizationClusterTimeZoneTypeClass
	MTRTimeSynchronizationClusterTimeZoneTypeClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterTimeZoneTypeClass() _MTRTimeSynchronizationClusterTimeZoneTypeClass {
	MTRTimeSynchronizationClusterTimeZoneTypeClassOnce.Do(func() {
		MTRTimeSynchronizationClusterTimeZoneTypeClass = _MTRTimeSynchronizationClusterTimeZoneTypeClass{objc.GetClass("MTRTimeSynchronizationClusterTimeZoneType")}
	})
	return MTRTimeSynchronizationClusterTimeZoneTypeClass
}

type _MTRTimeSynchronizationClusterTimeZoneTypeClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterTimeZoneType] class.
type IMTRTimeSynchronizationClusterTimeZoneType interface {
	IMTRTimeSynchronizationClusterTimeZoneStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneType
type MTRTimeSynchronizationClusterTimeZoneType struct {
	MTRTimeSynchronizationClusterTimeZoneStruct
}

// MTRTimeSynchronizationClusterTimeZoneTypeFrom constructs a [MTRTimeSynchronizationClusterTimeZoneType] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterTimeZoneTypeFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterTimeZoneType {
	return MTRTimeSynchronizationClusterTimeZoneType{
		MTRTimeSynchronizationClusterTimeZoneStruct: MTRTimeSynchronizationClusterTimeZoneStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterTimeZoneTypeClass) Alloc() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterTimeZoneTypeClass) New() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Init() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Autorelease() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterTimeZoneType creates a new MTRTimeSynchronizationClusterTimeZoneType instance.
func NewMTRTimeSynchronizationClusterTimeZoneType() MTRTimeSynchronizationClusterTimeZoneType {
	return getMTRTimeSynchronizationClusterTimeZoneTypeClass().New()
}




