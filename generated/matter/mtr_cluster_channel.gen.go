// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterChannel] class.
var (
	MTRClusterChannelClass     _MTRClusterChannelClass
	MTRClusterChannelClassOnce sync.Once
)

func getMTRClusterChannelClass() _MTRClusterChannelClass {
	MTRClusterChannelClassOnce.Do(func() {
		MTRClusterChannelClass = _MTRClusterChannelClass{objc.GetClass("MTRClusterChannel")}
	})
	return MTRClusterChannelClass
}

type _MTRClusterChannelClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterChannel] class.
type IMTRClusterChannel interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterChannel
type MTRClusterChannel struct {
	MTRGenericCluster
}

// MTRClusterChannelFrom constructs a [MTRClusterChannel] from an unsafe.Pointer.
func MTRClusterChannelFrom(ptr unsafe.Pointer) MTRClusterChannel {
	return MTRClusterChannel{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterChannelClass) Alloc() MTRClusterChannel {
	rv := objc.Send[MTRClusterChannel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterChannelClass) New() MTRClusterChannel {
	rv := objc.Send[MTRClusterChannel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterChannel) Init() MTRClusterChannel {
	rv := objc.Send[MTRClusterChannel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterChannel) Autorelease() MTRClusterChannel {
	rv := objc.Send[MTRClusterChannel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterChannel creates a new MTRClusterChannel instance.
func NewMTRClusterChannel() MTRClusterChannel {
	return getMTRClusterChannelClass().New()
}
