// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRControllerFactoryParams] class.
var (
	MTRControllerFactoryParamsClass     _MTRControllerFactoryParamsClass
	MTRControllerFactoryParamsClassOnce sync.Once
)

func getMTRControllerFactoryParamsClass() _MTRControllerFactoryParamsClass {
	MTRControllerFactoryParamsClassOnce.Do(func() {
		MTRControllerFactoryParamsClass = _MTRControllerFactoryParamsClass{objc.GetClass("MTRControllerFactoryParams")}
	})
	return MTRControllerFactoryParamsClass
}

type _MTRControllerFactoryParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRControllerFactoryParams] class.
type IMTRControllerFactoryParams interface {
	IMTRDeviceControllerFactoryParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactoryParams
type MTRControllerFactoryParams struct {
	MTRDeviceControllerFactoryParams
}

// MTRControllerFactoryParamsFrom constructs a [MTRControllerFactoryParams] from an unsafe.Pointer.
func MTRControllerFactoryParamsFrom(ptr unsafe.Pointer) MTRControllerFactoryParams {
	return MTRControllerFactoryParams{
		MTRDeviceControllerFactoryParams: MTRDeviceControllerFactoryParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRControllerFactoryParamsClass) Alloc() MTRControllerFactoryParams {
	rv := objc.Send[MTRControllerFactoryParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRControllerFactoryParamsClass) New() MTRControllerFactoryParams {
	rv := objc.Send[MTRControllerFactoryParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRControllerFactoryParams) Init() MTRControllerFactoryParams {
	rv := objc.Send[MTRControllerFactoryParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRControllerFactoryParams) Autorelease() MTRControllerFactoryParams {
	rv := objc.Send[MTRControllerFactoryParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRControllerFactoryParams creates a new MTRControllerFactoryParams instance.
func NewMTRControllerFactoryParams() MTRControllerFactoryParams {
	return getMTRControllerFactoryParamsClass().New()
}




