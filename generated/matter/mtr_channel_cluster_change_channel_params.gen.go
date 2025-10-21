// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterChangeChannelParams] class.
var (
	MTRChannelClusterChangeChannelParamsClass     _MTRChannelClusterChangeChannelParamsClass
	MTRChannelClusterChangeChannelParamsClassOnce sync.Once
)

func getMTRChannelClusterChangeChannelParamsClass() _MTRChannelClusterChangeChannelParamsClass {
	MTRChannelClusterChangeChannelParamsClassOnce.Do(func() {
		MTRChannelClusterChangeChannelParamsClass = _MTRChannelClusterChangeChannelParamsClass{objc.GetClass("MTRChannelClusterChangeChannelParams")}
	})
	return MTRChannelClusterChangeChannelParamsClass
}

type _MTRChannelClusterChangeChannelParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterChangeChannelParams] class.
type IMTRChannelClusterChangeChannelParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChangeChannelParams
type MTRChannelClusterChangeChannelParams struct {
	objectivec.Object
}

// MTRChannelClusterChangeChannelParamsFrom constructs a [MTRChannelClusterChangeChannelParams] from an unsafe.Pointer.
func MTRChannelClusterChangeChannelParamsFrom(ptr unsafe.Pointer) MTRChannelClusterChangeChannelParams {
	return MTRChannelClusterChangeChannelParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterChangeChannelParamsClass) Alloc() MTRChannelClusterChangeChannelParams {
	rv := objc.Send[MTRChannelClusterChangeChannelParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterChangeChannelParamsClass) New() MTRChannelClusterChangeChannelParams {
	rv := objc.Send[MTRChannelClusterChangeChannelParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterChangeChannelParams) Init() MTRChannelClusterChangeChannelParams {
	rv := objc.Send[MTRChannelClusterChangeChannelParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterChangeChannelParams) Autorelease() MTRChannelClusterChangeChannelParams {
	rv := objc.Send[MTRChannelClusterChangeChannelParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterChangeChannelParams creates a new MTRChannelClusterChangeChannelParams instance.
func NewMTRChannelClusterChangeChannelParams() MTRChannelClusterChangeChannelParams {
	return getMTRChannelClusterChangeChannelParamsClass().New()
}




