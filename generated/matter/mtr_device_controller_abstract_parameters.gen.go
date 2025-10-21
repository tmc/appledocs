// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceControllerAbstractParameters] class.
var (
	MTRDeviceControllerAbstractParametersClass     _MTRDeviceControllerAbstractParametersClass
	MTRDeviceControllerAbstractParametersClassOnce sync.Once
)

func getMTRDeviceControllerAbstractParametersClass() _MTRDeviceControllerAbstractParametersClass {
	MTRDeviceControllerAbstractParametersClassOnce.Do(func() {
		MTRDeviceControllerAbstractParametersClass = _MTRDeviceControllerAbstractParametersClass{objc.GetClass("MTRDeviceControllerAbstractParameters")}
	})
	return MTRDeviceControllerAbstractParametersClass
}

type _MTRDeviceControllerAbstractParametersClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceControllerAbstractParameters] class.
type IMTRDeviceControllerAbstractParameters interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerAbstractParameters
type MTRDeviceControllerAbstractParameters struct {
	objectivec.Object
}

// MTRDeviceControllerAbstractParametersFrom constructs a [MTRDeviceControllerAbstractParameters] from an unsafe.Pointer.
func MTRDeviceControllerAbstractParametersFrom(ptr unsafe.Pointer) MTRDeviceControllerAbstractParameters {
	return MTRDeviceControllerAbstractParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerAbstractParametersClass) Alloc() MTRDeviceControllerAbstractParameters {
	rv := objc.Send[MTRDeviceControllerAbstractParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceControllerAbstractParametersClass) New() MTRDeviceControllerAbstractParameters {
	rv := objc.Send[MTRDeviceControllerAbstractParameters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceControllerAbstractParameters) Init() MTRDeviceControllerAbstractParameters {
	rv := objc.Send[MTRDeviceControllerAbstractParameters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceControllerAbstractParameters) Autorelease() MTRDeviceControllerAbstractParameters {
	rv := objc.Send[MTRDeviceControllerAbstractParameters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceControllerAbstractParameters creates a new MTRDeviceControllerAbstractParameters instance.
func NewMTRDeviceControllerAbstractParameters() MTRDeviceControllerAbstractParameters {
	return getMTRDeviceControllerAbstractParametersClass().New()
}


// Whether the controller should start out suspended.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerabstractparameters/startsuspended
func (m_ MTRDeviceControllerAbstractParameters) StartSuspended() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("startSuspended"))
	return rv
}


// SetStartSuspended sets the value of the startSuspended property.
// Whether the controller should start out suspended.

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerabstractparameters/startsuspended
func (m_ MTRDeviceControllerAbstractParameters) SetStartSuspended(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartSuspended:"), value)
}



