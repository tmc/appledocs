// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDeviceControllerStartupParams] class.
var (
	MTRDeviceControllerStartupParamsClass     _MTRDeviceControllerStartupParamsClass
	MTRDeviceControllerStartupParamsClassOnce sync.Once
)

func getMTRDeviceControllerStartupParamsClass() _MTRDeviceControllerStartupParamsClass {
	MTRDeviceControllerStartupParamsClassOnce.Do(func() {
		MTRDeviceControllerStartupParamsClass = _MTRDeviceControllerStartupParamsClass{objc.GetClass("MTRDeviceControllerStartupParams")}
	})
	return MTRDeviceControllerStartupParamsClass
}

type _MTRDeviceControllerStartupParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceControllerStartupParams] class.
type IMTRDeviceControllerStartupParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams
type MTRDeviceControllerStartupParams struct {
	objectivec.Object
}

// MTRDeviceControllerStartupParamsFrom constructs a [MTRDeviceControllerStartupParams] from an unsafe.Pointer.
func MTRDeviceControllerStartupParamsFrom(ptr unsafe.Pointer) MTRDeviceControllerStartupParams {
	return MTRDeviceControllerStartupParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerStartupParamsClass) Alloc() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceControllerStartupParamsClass) New() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceControllerStartupParams) Init() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceControllerStartupParams) Autorelease() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceControllerStartupParams creates a new MTRDeviceControllerStartupParams instance.
func NewMTRDeviceControllerStartupParams() MTRDeviceControllerStartupParams {
	return getMTRDeviceControllerStartupParamsClass().New()
}




