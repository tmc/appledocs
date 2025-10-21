// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRWindowCoveringClusterDownOrCloseParams] class.
var (
	MTRWindowCoveringClusterDownOrCloseParamsClass     _MTRWindowCoveringClusterDownOrCloseParamsClass
	MTRWindowCoveringClusterDownOrCloseParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterDownOrCloseParamsClass() _MTRWindowCoveringClusterDownOrCloseParamsClass {
	MTRWindowCoveringClusterDownOrCloseParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterDownOrCloseParamsClass = _MTRWindowCoveringClusterDownOrCloseParamsClass{objc.GetClass("MTRWindowCoveringClusterDownOrCloseParams")}
	})
	return MTRWindowCoveringClusterDownOrCloseParamsClass
}

type _MTRWindowCoveringClusterDownOrCloseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWindowCoveringClusterDownOrCloseParams] class.
type IMTRWindowCoveringClusterDownOrCloseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterDownOrCloseParams
type MTRWindowCoveringClusterDownOrCloseParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterDownOrCloseParamsFrom constructs a [MTRWindowCoveringClusterDownOrCloseParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterDownOrCloseParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterDownOrCloseParams {
	return MTRWindowCoveringClusterDownOrCloseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterDownOrCloseParamsClass) Alloc() MTRWindowCoveringClusterDownOrCloseParams {
	rv := objc.Send[MTRWindowCoveringClusterDownOrCloseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWindowCoveringClusterDownOrCloseParamsClass) New() MTRWindowCoveringClusterDownOrCloseParams {
	rv := objc.Send[MTRWindowCoveringClusterDownOrCloseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterDownOrCloseParams) Init() MTRWindowCoveringClusterDownOrCloseParams {
	rv := objc.Send[MTRWindowCoveringClusterDownOrCloseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterDownOrCloseParams) Autorelease() MTRWindowCoveringClusterDownOrCloseParams {
	rv := objc.Send[MTRWindowCoveringClusterDownOrCloseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterDownOrCloseParams creates a new MTRWindowCoveringClusterDownOrCloseParams instance.
func NewMTRWindowCoveringClusterDownOrCloseParams() MTRWindowCoveringClusterDownOrCloseParams {
	return getMTRWindowCoveringClusterDownOrCloseParamsClass().New()
}




