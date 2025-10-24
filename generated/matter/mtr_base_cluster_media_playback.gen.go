// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterMediaPlayback] class.
var (
	MTRBaseClusterMediaPlaybackClass     _MTRBaseClusterMediaPlaybackClass
	MTRBaseClusterMediaPlaybackClassOnce sync.Once
)

func getMTRBaseClusterMediaPlaybackClass() _MTRBaseClusterMediaPlaybackClass {
	MTRBaseClusterMediaPlaybackClassOnce.Do(func() {
		MTRBaseClusterMediaPlaybackClass = _MTRBaseClusterMediaPlaybackClass{objc.GetClass("MTRBaseClusterMediaPlayback")}
	})
	return MTRBaseClusterMediaPlaybackClass
}

type _MTRBaseClusterMediaPlaybackClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterMediaPlayback] class.
type IMTRBaseClusterMediaPlayback interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMediaPlayback
type MTRBaseClusterMediaPlayback struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterMediaPlaybackFrom constructs a [MTRBaseClusterMediaPlayback] from an unsafe.Pointer.
func MTRBaseClusterMediaPlaybackFrom(ptr unsafe.Pointer) MTRBaseClusterMediaPlayback {
	return MTRBaseClusterMediaPlayback{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterMediaPlaybackClass) Alloc() MTRBaseClusterMediaPlayback {
	rv := objc.Send[MTRBaseClusterMediaPlayback](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterMediaPlaybackClass) New() MTRBaseClusterMediaPlayback {
	rv := objc.Send[MTRBaseClusterMediaPlayback](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterMediaPlayback) Init() MTRBaseClusterMediaPlayback {
	rv := objc.Send[MTRBaseClusterMediaPlayback](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterMediaPlayback) Autorelease() MTRBaseClusterMediaPlayback {
	rv := objc.Send[MTRBaseClusterMediaPlayback](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterMediaPlayback creates a new MTRBaseClusterMediaPlayback instance.
func NewMTRBaseClusterMediaPlayback() MTRBaseClusterMediaPlayback {
	return getMTRBaseClusterMediaPlaybackClass().New()
}
