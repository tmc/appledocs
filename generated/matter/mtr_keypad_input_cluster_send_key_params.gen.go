// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRKeypadInputClusterSendKeyParams] class.
var (
	MTRKeypadInputClusterSendKeyParamsClass     _MTRKeypadInputClusterSendKeyParamsClass
	MTRKeypadInputClusterSendKeyParamsClassOnce sync.Once
)

func getMTRKeypadInputClusterSendKeyParamsClass() _MTRKeypadInputClusterSendKeyParamsClass {
	MTRKeypadInputClusterSendKeyParamsClassOnce.Do(func() {
		MTRKeypadInputClusterSendKeyParamsClass = _MTRKeypadInputClusterSendKeyParamsClass{objc.GetClass("MTRKeypadInputClusterSendKeyParams")}
	})
	return MTRKeypadInputClusterSendKeyParamsClass
}

type _MTRKeypadInputClusterSendKeyParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRKeypadInputClusterSendKeyParams] class.
type IMTRKeypadInputClusterSendKeyParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRKeypadInputClusterSendKeyParams
type MTRKeypadInputClusterSendKeyParams struct {
	objectivec.Object
}

// MTRKeypadInputClusterSendKeyParamsFrom constructs a [MTRKeypadInputClusterSendKeyParams] from an unsafe.Pointer.
func MTRKeypadInputClusterSendKeyParamsFrom(ptr unsafe.Pointer) MTRKeypadInputClusterSendKeyParams {
	return MTRKeypadInputClusterSendKeyParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRKeypadInputClusterSendKeyParamsClass) Alloc() MTRKeypadInputClusterSendKeyParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRKeypadInputClusterSendKeyParamsClass) New() MTRKeypadInputClusterSendKeyParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRKeypadInputClusterSendKeyParams) Init() MTRKeypadInputClusterSendKeyParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRKeypadInputClusterSendKeyParams) Autorelease() MTRKeypadInputClusterSendKeyParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRKeypadInputClusterSendKeyParams creates a new MTRKeypadInputClusterSendKeyParams instance.
func NewMTRKeypadInputClusterSendKeyParams() MTRKeypadInputClusterSendKeyParams {
	return getMTRKeypadInputClusterSendKeyParamsClass().New()
}




