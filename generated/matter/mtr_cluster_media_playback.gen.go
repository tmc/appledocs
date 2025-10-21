// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterMediaPlayback] class.
var (
	MTRClusterMediaPlaybackClass     _MTRClusterMediaPlaybackClass
	MTRClusterMediaPlaybackClassOnce sync.Once
)

func getMTRClusterMediaPlaybackClass() _MTRClusterMediaPlaybackClass {
	MTRClusterMediaPlaybackClassOnce.Do(func() {
		MTRClusterMediaPlaybackClass = _MTRClusterMediaPlaybackClass{objc.GetClass("MTRClusterMediaPlayback")}
	})
	return MTRClusterMediaPlaybackClass
}

type _MTRClusterMediaPlaybackClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterMediaPlayback] class.
type IMTRClusterMediaPlayback interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMediaPlayback
type MTRClusterMediaPlayback struct {
	MTRGenericCluster
}

// MTRClusterMediaPlaybackFrom constructs a [MTRClusterMediaPlayback] from an unsafe.Pointer.
func MTRClusterMediaPlaybackFrom(ptr unsafe.Pointer) MTRClusterMediaPlayback {
	return MTRClusterMediaPlayback{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterMediaPlaybackClass) Alloc() MTRClusterMediaPlayback {
	rv := objc.Send[MTRClusterMediaPlayback](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterMediaPlaybackClass) New() MTRClusterMediaPlayback {
	rv := objc.Send[MTRClusterMediaPlayback](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterMediaPlayback) Init() MTRClusterMediaPlayback {
	rv := objc.Send[MTRClusterMediaPlayback](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterMediaPlayback) Autorelease() MTRClusterMediaPlayback {
	rv := objc.Send[MTRClusterMediaPlayback](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterMediaPlayback creates a new MTRClusterMediaPlayback instance.
func NewMTRClusterMediaPlayback() MTRClusterMediaPlayback {
	return getMTRClusterMediaPlaybackClass().New()
}




