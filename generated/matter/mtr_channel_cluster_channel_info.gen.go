// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRChannelClusterChannelInfo] class.
var (
	MTRChannelClusterChannelInfoClass     _MTRChannelClusterChannelInfoClass
	MTRChannelClusterChannelInfoClassOnce sync.Once
)

func getMTRChannelClusterChannelInfoClass() _MTRChannelClusterChannelInfoClass {
	MTRChannelClusterChannelInfoClassOnce.Do(func() {
		MTRChannelClusterChannelInfoClass = _MTRChannelClusterChannelInfoClass{objc.GetClass("MTRChannelClusterChannelInfo")}
	})
	return MTRChannelClusterChannelInfoClass
}

type _MTRChannelClusterChannelInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterChannelInfo] class.
type IMTRChannelClusterChannelInfo interface {
	IMTRChannelClusterChannelInfoStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelInfo
type MTRChannelClusterChannelInfo struct {
	MTRChannelClusterChannelInfoStruct
}

// MTRChannelClusterChannelInfoFrom constructs a [MTRChannelClusterChannelInfo] from an unsafe.Pointer.
func MTRChannelClusterChannelInfoFrom(ptr unsafe.Pointer) MTRChannelClusterChannelInfo {
	return MTRChannelClusterChannelInfo{
		MTRChannelClusterChannelInfoStruct: MTRChannelClusterChannelInfoStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterChannelInfoClass) Alloc() MTRChannelClusterChannelInfo {
	rv := objc.Send[MTRChannelClusterChannelInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterChannelInfoClass) New() MTRChannelClusterChannelInfo {
	rv := objc.Send[MTRChannelClusterChannelInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterChannelInfo) Init() MTRChannelClusterChannelInfo {
	rv := objc.Send[MTRChannelClusterChannelInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterChannelInfo) Autorelease() MTRChannelClusterChannelInfo {
	rv := objc.Send[MTRChannelClusterChannelInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterChannelInfo creates a new MTRChannelClusterChannelInfo instance.
func NewMTRChannelClusterChannelInfo() MTRChannelClusterChannelInfo {
	return getMTRChannelClusterChannelInfoClass().New()
}




