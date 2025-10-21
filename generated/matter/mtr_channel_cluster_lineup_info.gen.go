// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRChannelClusterLineupInfo] class.
var (
	MTRChannelClusterLineupInfoClass     _MTRChannelClusterLineupInfoClass
	MTRChannelClusterLineupInfoClassOnce sync.Once
)

func getMTRChannelClusterLineupInfoClass() _MTRChannelClusterLineupInfoClass {
	MTRChannelClusterLineupInfoClassOnce.Do(func() {
		MTRChannelClusterLineupInfoClass = _MTRChannelClusterLineupInfoClass{objc.GetClass("MTRChannelClusterLineupInfo")}
	})
	return MTRChannelClusterLineupInfoClass
}

type _MTRChannelClusterLineupInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterLineupInfo] class.
type IMTRChannelClusterLineupInfo interface {
	IMTRChannelClusterLineupInfoStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterLineupInfo
type MTRChannelClusterLineupInfo struct {
	MTRChannelClusterLineupInfoStruct
}

// MTRChannelClusterLineupInfoFrom constructs a [MTRChannelClusterLineupInfo] from an unsafe.Pointer.
func MTRChannelClusterLineupInfoFrom(ptr unsafe.Pointer) MTRChannelClusterLineupInfo {
	return MTRChannelClusterLineupInfo{
		MTRChannelClusterLineupInfoStruct: MTRChannelClusterLineupInfoStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterLineupInfoClass) Alloc() MTRChannelClusterLineupInfo {
	rv := objc.Send[MTRChannelClusterLineupInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterLineupInfoClass) New() MTRChannelClusterLineupInfo {
	rv := objc.Send[MTRChannelClusterLineupInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterLineupInfo) Init() MTRChannelClusterLineupInfo {
	rv := objc.Send[MTRChannelClusterLineupInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterLineupInfo) Autorelease() MTRChannelClusterLineupInfo {
	rv := objc.Send[MTRChannelClusterLineupInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterLineupInfo creates a new MTRChannelClusterLineupInfo instance.
func NewMTRChannelClusterLineupInfo() MTRChannelClusterLineupInfo {
	return getMTRChannelClusterLineupInfoClass().New()
}




