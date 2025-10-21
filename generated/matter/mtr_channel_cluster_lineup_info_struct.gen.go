// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRChannelClusterLineupInfoStruct] class.
var (
	MTRChannelClusterLineupInfoStructClass     _MTRChannelClusterLineupInfoStructClass
	MTRChannelClusterLineupInfoStructClassOnce sync.Once
)

func getMTRChannelClusterLineupInfoStructClass() _MTRChannelClusterLineupInfoStructClass {
	MTRChannelClusterLineupInfoStructClassOnce.Do(func() {
		MTRChannelClusterLineupInfoStructClass = _MTRChannelClusterLineupInfoStructClass{objc.GetClass("MTRChannelClusterLineupInfoStruct")}
	})
	return MTRChannelClusterLineupInfoStructClass
}

type _MTRChannelClusterLineupInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterLineupInfoStruct] class.
type IMTRChannelClusterLineupInfoStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterLineupInfoStruct
type MTRChannelClusterLineupInfoStruct struct {
	objectivec.Object
}

// MTRChannelClusterLineupInfoStructFrom constructs a [MTRChannelClusterLineupInfoStruct] from an unsafe.Pointer.
func MTRChannelClusterLineupInfoStructFrom(ptr unsafe.Pointer) MTRChannelClusterLineupInfoStruct {
	return MTRChannelClusterLineupInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterLineupInfoStructClass) Alloc() MTRChannelClusterLineupInfoStruct {
	rv := objc.Send[MTRChannelClusterLineupInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterLineupInfoStructClass) New() MTRChannelClusterLineupInfoStruct {
	rv := objc.Send[MTRChannelClusterLineupInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterLineupInfoStruct) Init() MTRChannelClusterLineupInfoStruct {
	rv := objc.Send[MTRChannelClusterLineupInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterLineupInfoStruct) Autorelease() MTRChannelClusterLineupInfoStruct {
	rv := objc.Send[MTRChannelClusterLineupInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterLineupInfoStruct creates a new MTRChannelClusterLineupInfoStruct instance.
func NewMTRChannelClusterLineupInfoStruct() MTRChannelClusterLineupInfoStruct {
	return getMTRChannelClusterLineupInfoStructClass().New()
}




