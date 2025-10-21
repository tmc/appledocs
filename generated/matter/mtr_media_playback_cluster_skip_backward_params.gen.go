// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRMediaPlaybackClusterSkipBackwardParams] class.
var (
	MTRMediaPlaybackClusterSkipBackwardParamsClass     _MTRMediaPlaybackClusterSkipBackwardParamsClass
	MTRMediaPlaybackClusterSkipBackwardParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterSkipBackwardParamsClass() _MTRMediaPlaybackClusterSkipBackwardParamsClass {
	MTRMediaPlaybackClusterSkipBackwardParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterSkipBackwardParamsClass = _MTRMediaPlaybackClusterSkipBackwardParamsClass{objc.GetClass("MTRMediaPlaybackClusterSkipBackwardParams")}
	})
	return MTRMediaPlaybackClusterSkipBackwardParamsClass
}

type _MTRMediaPlaybackClusterSkipBackwardParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterSkipBackwardParams] class.
type IMTRMediaPlaybackClusterSkipBackwardParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterSkipBackwardParams
type MTRMediaPlaybackClusterSkipBackwardParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterSkipBackwardParamsFrom constructs a [MTRMediaPlaybackClusterSkipBackwardParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterSkipBackwardParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterSkipBackwardParams {
	return MTRMediaPlaybackClusterSkipBackwardParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterSkipBackwardParamsClass) Alloc() MTRMediaPlaybackClusterSkipBackwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipBackwardParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterSkipBackwardParamsClass) New() MTRMediaPlaybackClusterSkipBackwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipBackwardParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterSkipBackwardParams) Init() MTRMediaPlaybackClusterSkipBackwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipBackwardParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterSkipBackwardParams) Autorelease() MTRMediaPlaybackClusterSkipBackwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipBackwardParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterSkipBackwardParams creates a new MTRMediaPlaybackClusterSkipBackwardParams instance.
func NewMTRMediaPlaybackClusterSkipBackwardParams() MTRMediaPlaybackClusterSkipBackwardParams {
	return getMTRMediaPlaybackClusterSkipBackwardParamsClass().New()
}




