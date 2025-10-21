// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTargetNavigatorClusterNavigateTargetResponseParams] class.
var (
	MTRTargetNavigatorClusterNavigateTargetResponseParamsClass     _MTRTargetNavigatorClusterNavigateTargetResponseParamsClass
	MTRTargetNavigatorClusterNavigateTargetResponseParamsClassOnce sync.Once
)

func getMTRTargetNavigatorClusterNavigateTargetResponseParamsClass() _MTRTargetNavigatorClusterNavigateTargetResponseParamsClass {
	MTRTargetNavigatorClusterNavigateTargetResponseParamsClassOnce.Do(func() {
		MTRTargetNavigatorClusterNavigateTargetResponseParamsClass = _MTRTargetNavigatorClusterNavigateTargetResponseParamsClass{objc.GetClass("MTRTargetNavigatorClusterNavigateTargetResponseParams")}
	})
	return MTRTargetNavigatorClusterNavigateTargetResponseParamsClass
}

type _MTRTargetNavigatorClusterNavigateTargetResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTargetNavigatorClusterNavigateTargetResponseParams] class.
type IMTRTargetNavigatorClusterNavigateTargetResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterNavigateTargetResponseParams
type MTRTargetNavigatorClusterNavigateTargetResponseParams struct {
	objectivec.Object
}

// MTRTargetNavigatorClusterNavigateTargetResponseParamsFrom constructs a [MTRTargetNavigatorClusterNavigateTargetResponseParams] from an unsafe.Pointer.
func MTRTargetNavigatorClusterNavigateTargetResponseParamsFrom(ptr unsafe.Pointer) MTRTargetNavigatorClusterNavigateTargetResponseParams {
	return MTRTargetNavigatorClusterNavigateTargetResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTargetNavigatorClusterNavigateTargetResponseParamsClass) Alloc() MTRTargetNavigatorClusterNavigateTargetResponseParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTargetNavigatorClusterNavigateTargetResponseParamsClass) New() MTRTargetNavigatorClusterNavigateTargetResponseParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTargetNavigatorClusterNavigateTargetResponseParams) Init() MTRTargetNavigatorClusterNavigateTargetResponseParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTargetNavigatorClusterNavigateTargetResponseParams) Autorelease() MTRTargetNavigatorClusterNavigateTargetResponseParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTargetNavigatorClusterNavigateTargetResponseParams creates a new MTRTargetNavigatorClusterNavigateTargetResponseParams instance.
func NewMTRTargetNavigatorClusterNavigateTargetResponseParams() MTRTargetNavigatorClusterNavigateTargetResponseParams {
	return getMTRTargetNavigatorClusterNavigateTargetResponseParamsClass().New()
}




