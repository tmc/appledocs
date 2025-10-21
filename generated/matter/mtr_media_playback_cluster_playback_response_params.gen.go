// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterPlaybackResponseParams] class.
var (
	MTRMediaPlaybackClusterPlaybackResponseParamsClass     _MTRMediaPlaybackClusterPlaybackResponseParamsClass
	MTRMediaPlaybackClusterPlaybackResponseParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterPlaybackResponseParamsClass() _MTRMediaPlaybackClusterPlaybackResponseParamsClass {
	MTRMediaPlaybackClusterPlaybackResponseParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterPlaybackResponseParamsClass = _MTRMediaPlaybackClusterPlaybackResponseParamsClass{objc.GetClass("MTRMediaPlaybackClusterPlaybackResponseParams")}
	})
	return MTRMediaPlaybackClusterPlaybackResponseParamsClass
}

type _MTRMediaPlaybackClusterPlaybackResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterPlaybackResponseParams] class.
type IMTRMediaPlaybackClusterPlaybackResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterPlaybackResponseParams
type MTRMediaPlaybackClusterPlaybackResponseParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterPlaybackResponseParamsFrom constructs a [MTRMediaPlaybackClusterPlaybackResponseParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterPlaybackResponseParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterPlaybackResponseParams {
	return MTRMediaPlaybackClusterPlaybackResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterPlaybackResponseParamsClass) Alloc() MTRMediaPlaybackClusterPlaybackResponseParams {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterPlaybackResponseParamsClass) New() MTRMediaPlaybackClusterPlaybackResponseParams {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterPlaybackResponseParams) Init() MTRMediaPlaybackClusterPlaybackResponseParams {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterPlaybackResponseParams) Autorelease() MTRMediaPlaybackClusterPlaybackResponseParams {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterPlaybackResponseParams creates a new MTRMediaPlaybackClusterPlaybackResponseParams instance.
func NewMTRMediaPlaybackClusterPlaybackResponseParams() MTRMediaPlaybackClusterPlaybackResponseParams {
	return getMTRMediaPlaybackClusterPlaybackResponseParamsClass().New()
}




