// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRSubscribeParams] class.
var (
	MTRSubscribeParamsClass     _MTRSubscribeParamsClass
	MTRSubscribeParamsClassOnce sync.Once
)

func getMTRSubscribeParamsClass() _MTRSubscribeParamsClass {
	MTRSubscribeParamsClassOnce.Do(func() {
		MTRSubscribeParamsClass = _MTRSubscribeParamsClass{objc.GetClass("MTRSubscribeParams")}
	})
	return MTRSubscribeParamsClass
}

type _MTRSubscribeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRSubscribeParams] class.
type IMTRSubscribeParams interface {
	IMTRReadParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams
type MTRSubscribeParams struct {
	MTRReadParams
}

// MTRSubscribeParamsFrom constructs a [MTRSubscribeParams] from an unsafe.Pointer.
func MTRSubscribeParamsFrom(ptr unsafe.Pointer) MTRSubscribeParams {
	return MTRSubscribeParams{
		MTRReadParams: MTRReadParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSubscribeParamsClass) Alloc() MTRSubscribeParams {
	rv := objc.Send[MTRSubscribeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSubscribeParamsClass) New() MTRSubscribeParams {
	rv := objc.Send[MTRSubscribeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSubscribeParams) Init() MTRSubscribeParams {
	rv := objc.Send[MTRSubscribeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSubscribeParams) Autorelease() MTRSubscribeParams {
	rv := objc.Send[MTRSubscribeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSubscribeParams creates a new MTRSubscribeParams instance.
func NewMTRSubscribeParams() MTRSubscribeParams {
	return getMTRSubscribeParamsClass().New()
}




