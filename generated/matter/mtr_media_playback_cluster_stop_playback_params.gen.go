// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRMediaPlaybackClusterStopPlaybackParams] class.
var (
	MTRMediaPlaybackClusterStopPlaybackParamsClass     _MTRMediaPlaybackClusterStopPlaybackParamsClass
	MTRMediaPlaybackClusterStopPlaybackParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterStopPlaybackParamsClass() _MTRMediaPlaybackClusterStopPlaybackParamsClass {
	MTRMediaPlaybackClusterStopPlaybackParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterStopPlaybackParamsClass = _MTRMediaPlaybackClusterStopPlaybackParamsClass{objc.GetClass("MTRMediaPlaybackClusterStopPlaybackParams")}
	})
	return MTRMediaPlaybackClusterStopPlaybackParamsClass
}

type _MTRMediaPlaybackClusterStopPlaybackParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterStopPlaybackParams] class.
type IMTRMediaPlaybackClusterStopPlaybackParams interface {
	IMTRMediaPlaybackClusterStopParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStopPlaybackParams
type MTRMediaPlaybackClusterStopPlaybackParams struct {
	MTRMediaPlaybackClusterStopParams
}

// MTRMediaPlaybackClusterStopPlaybackParamsFrom constructs a [MTRMediaPlaybackClusterStopPlaybackParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterStopPlaybackParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterStopPlaybackParams {
	return MTRMediaPlaybackClusterStopPlaybackParams{
		MTRMediaPlaybackClusterStopParams: MTRMediaPlaybackClusterStopParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterStopPlaybackParamsClass) Alloc() MTRMediaPlaybackClusterStopPlaybackParams {
	rv := objc.Send[MTRMediaPlaybackClusterStopPlaybackParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterStopPlaybackParamsClass) New() MTRMediaPlaybackClusterStopPlaybackParams {
	rv := objc.Send[MTRMediaPlaybackClusterStopPlaybackParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterStopPlaybackParams) Init() MTRMediaPlaybackClusterStopPlaybackParams {
	rv := objc.Send[MTRMediaPlaybackClusterStopPlaybackParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterStopPlaybackParams) Autorelease() MTRMediaPlaybackClusterStopPlaybackParams {
	rv := objc.Send[MTRMediaPlaybackClusterStopPlaybackParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterStopPlaybackParams creates a new MTRMediaPlaybackClusterStopPlaybackParams instance.
func NewMTRMediaPlaybackClusterStopPlaybackParams() MTRMediaPlaybackClusterStopPlaybackParams {
	return getMTRMediaPlaybackClusterStopPlaybackParamsClass().New()
}




