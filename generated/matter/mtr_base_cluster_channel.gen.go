// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterChannel] class.
var (
	MTRBaseClusterChannelClass     _MTRBaseClusterChannelClass
	MTRBaseClusterChannelClassOnce sync.Once
)

func getMTRBaseClusterChannelClass() _MTRBaseClusterChannelClass {
	MTRBaseClusterChannelClassOnce.Do(func() {
		MTRBaseClusterChannelClass = _MTRBaseClusterChannelClass{objc.GetClass("MTRBaseClusterChannel")}
	})
	return MTRBaseClusterChannelClass
}

type _MTRBaseClusterChannelClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterChannel] class.
type IMTRBaseClusterChannel interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterChannel
type MTRBaseClusterChannel struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterChannelFrom constructs a [MTRBaseClusterChannel] from an unsafe.Pointer.
func MTRBaseClusterChannelFrom(ptr unsafe.Pointer) MTRBaseClusterChannel {
	return MTRBaseClusterChannel{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterChannelClass) Alloc() MTRBaseClusterChannel {
	rv := objc.Send[MTRBaseClusterChannel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterChannelClass) New() MTRBaseClusterChannel {
	rv := objc.Send[MTRBaseClusterChannel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterChannel) Init() MTRBaseClusterChannel {
	rv := objc.Send[MTRBaseClusterChannel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterChannel) Autorelease() MTRBaseClusterChannel {
	rv := objc.Send[MTRBaseClusterChannel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterChannel creates a new MTRBaseClusterChannel instance.
func NewMTRBaseClusterChannel() MTRBaseClusterChannel {
	return getMTRBaseClusterChannelClass().New()
}




