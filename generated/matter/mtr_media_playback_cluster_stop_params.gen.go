// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRMediaPlaybackClusterStopParams] class.
var (
	MTRMediaPlaybackClusterStopParamsClass     _MTRMediaPlaybackClusterStopParamsClass
	MTRMediaPlaybackClusterStopParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterStopParamsClass() _MTRMediaPlaybackClusterStopParamsClass {
	MTRMediaPlaybackClusterStopParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterStopParamsClass = _MTRMediaPlaybackClusterStopParamsClass{objc.GetClass("MTRMediaPlaybackClusterStopParams")}
	})
	return MTRMediaPlaybackClusterStopParamsClass
}

type _MTRMediaPlaybackClusterStopParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterStopParams] class.
type IMTRMediaPlaybackClusterStopParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStopParams
type MTRMediaPlaybackClusterStopParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterStopParamsFrom constructs a [MTRMediaPlaybackClusterStopParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterStopParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterStopParams {
	return MTRMediaPlaybackClusterStopParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterStopParamsClass) Alloc() MTRMediaPlaybackClusterStopParams {
	rv := objc.Send[MTRMediaPlaybackClusterStopParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterStopParamsClass) New() MTRMediaPlaybackClusterStopParams {
	rv := objc.Send[MTRMediaPlaybackClusterStopParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterStopParams) Init() MTRMediaPlaybackClusterStopParams {
	rv := objc.Send[MTRMediaPlaybackClusterStopParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterStopParams) Autorelease() MTRMediaPlaybackClusterStopParams {
	rv := objc.Send[MTRMediaPlaybackClusterStopParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterStopParams creates a new MTRMediaPlaybackClusterStopParams instance.
func NewMTRMediaPlaybackClusterStopParams() MTRMediaPlaybackClusterStopParams {
	return getMTRMediaPlaybackClusterStopParamsClass().New()
}




