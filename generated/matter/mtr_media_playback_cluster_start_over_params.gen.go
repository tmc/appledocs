// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRMediaPlaybackClusterStartOverParams] class.
var (
	MTRMediaPlaybackClusterStartOverParamsClass     _MTRMediaPlaybackClusterStartOverParamsClass
	MTRMediaPlaybackClusterStartOverParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterStartOverParamsClass() _MTRMediaPlaybackClusterStartOverParamsClass {
	MTRMediaPlaybackClusterStartOverParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterStartOverParamsClass = _MTRMediaPlaybackClusterStartOverParamsClass{objc.GetClass("MTRMediaPlaybackClusterStartOverParams")}
	})
	return MTRMediaPlaybackClusterStartOverParamsClass
}

type _MTRMediaPlaybackClusterStartOverParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterStartOverParams] class.
type IMTRMediaPlaybackClusterStartOverParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStartOverParams
type MTRMediaPlaybackClusterStartOverParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterStartOverParamsFrom constructs a [MTRMediaPlaybackClusterStartOverParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterStartOverParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterStartOverParams {
	return MTRMediaPlaybackClusterStartOverParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterStartOverParamsClass) Alloc() MTRMediaPlaybackClusterStartOverParams {
	rv := objc.Send[MTRMediaPlaybackClusterStartOverParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterStartOverParamsClass) New() MTRMediaPlaybackClusterStartOverParams {
	rv := objc.Send[MTRMediaPlaybackClusterStartOverParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterStartOverParams) Init() MTRMediaPlaybackClusterStartOverParams {
	rv := objc.Send[MTRMediaPlaybackClusterStartOverParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterStartOverParams) Autorelease() MTRMediaPlaybackClusterStartOverParams {
	rv := objc.Send[MTRMediaPlaybackClusterStartOverParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterStartOverParams creates a new MTRMediaPlaybackClusterStartOverParams instance.
func NewMTRMediaPlaybackClusterStartOverParams() MTRMediaPlaybackClusterStartOverParams {
	return getMTRMediaPlaybackClusterStartOverParamsClass().New()
}




