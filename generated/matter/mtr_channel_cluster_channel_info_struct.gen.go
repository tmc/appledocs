// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterChannelInfoStruct] class.
var (
	MTRChannelClusterChannelInfoStructClass     _MTRChannelClusterChannelInfoStructClass
	MTRChannelClusterChannelInfoStructClassOnce sync.Once
)

func getMTRChannelClusterChannelInfoStructClass() _MTRChannelClusterChannelInfoStructClass {
	MTRChannelClusterChannelInfoStructClassOnce.Do(func() {
		MTRChannelClusterChannelInfoStructClass = _MTRChannelClusterChannelInfoStructClass{objc.GetClass("MTRChannelClusterChannelInfoStruct")}
	})
	return MTRChannelClusterChannelInfoStructClass
}

type _MTRChannelClusterChannelInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterChannelInfoStruct] class.
type IMTRChannelClusterChannelInfoStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelInfoStruct
type MTRChannelClusterChannelInfoStruct struct {
	objectivec.Object
}

// MTRChannelClusterChannelInfoStructFrom constructs a [MTRChannelClusterChannelInfoStruct] from an unsafe.Pointer.
func MTRChannelClusterChannelInfoStructFrom(ptr unsafe.Pointer) MTRChannelClusterChannelInfoStruct {
	return MTRChannelClusterChannelInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterChannelInfoStructClass) Alloc() MTRChannelClusterChannelInfoStruct {
	rv := objc.Send[MTRChannelClusterChannelInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterChannelInfoStructClass) New() MTRChannelClusterChannelInfoStruct {
	rv := objc.Send[MTRChannelClusterChannelInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterChannelInfoStruct) Init() MTRChannelClusterChannelInfoStruct {
	rv := objc.Send[MTRChannelClusterChannelInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterChannelInfoStruct) Autorelease() MTRChannelClusterChannelInfoStruct {
	rv := objc.Send[MTRChannelClusterChannelInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterChannelInfoStruct creates a new MTRChannelClusterChannelInfoStruct instance.
func NewMTRChannelClusterChannelInfoStruct() MTRChannelClusterChannelInfoStruct {
	return getMTRChannelClusterChannelInfoStructClass().New()
}




