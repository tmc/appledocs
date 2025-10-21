// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRDeviceControllerParameters] class.
var (
	MTRDeviceControllerParametersClass     _MTRDeviceControllerParametersClass
	MTRDeviceControllerParametersClassOnce sync.Once
)

func getMTRDeviceControllerParametersClass() _MTRDeviceControllerParametersClass {
	MTRDeviceControllerParametersClassOnce.Do(func() {
		MTRDeviceControllerParametersClass = _MTRDeviceControllerParametersClass{objc.GetClass("MTRDeviceControllerParameters")}
	})
	return MTRDeviceControllerParametersClass
}

type _MTRDeviceControllerParametersClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceControllerParameters] class.
type IMTRDeviceControllerParameters interface {
	IMTRDeviceControllerAbstractParameters
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerParameters
type MTRDeviceControllerParameters struct {
	MTRDeviceControllerAbstractParameters
}

// MTRDeviceControllerParametersFrom constructs a [MTRDeviceControllerParameters] from an unsafe.Pointer.
func MTRDeviceControllerParametersFrom(ptr unsafe.Pointer) MTRDeviceControllerParameters {
	return MTRDeviceControllerParameters{
		MTRDeviceControllerAbstractParameters: MTRDeviceControllerAbstractParametersFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerParametersClass) Alloc() MTRDeviceControllerParameters {
	rv := objc.Send[MTRDeviceControllerParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceControllerParametersClass) New() MTRDeviceControllerParameters {
	rv := objc.Send[MTRDeviceControllerParameters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceControllerParameters) Init() MTRDeviceControllerParameters {
	rv := objc.Send[MTRDeviceControllerParameters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceControllerParameters) Autorelease() MTRDeviceControllerParameters {
	rv := objc.Send[MTRDeviceControllerParameters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceControllerParameters creates a new MTRDeviceControllerParameters instance.
func NewMTRDeviceControllerParameters() MTRDeviceControllerParameters {
	return getMTRDeviceControllerParametersClass().New()
}




