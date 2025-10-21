// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRDeviceControllerExternalCertificateParameters] class.
var (
	MTRDeviceControllerExternalCertificateParametersClass     _MTRDeviceControllerExternalCertificateParametersClass
	MTRDeviceControllerExternalCertificateParametersClassOnce sync.Once
)

func getMTRDeviceControllerExternalCertificateParametersClass() _MTRDeviceControllerExternalCertificateParametersClass {
	MTRDeviceControllerExternalCertificateParametersClassOnce.Do(func() {
		MTRDeviceControllerExternalCertificateParametersClass = _MTRDeviceControllerExternalCertificateParametersClass{objc.GetClass("MTRDeviceControllerExternalCertificateParameters")}
	})
	return MTRDeviceControllerExternalCertificateParametersClass
}

type _MTRDeviceControllerExternalCertificateParametersClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceControllerExternalCertificateParameters] class.
type IMTRDeviceControllerExternalCertificateParameters interface {
	IMTRDeviceControllerParameters
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerExternalCertificateParameters
type MTRDeviceControllerExternalCertificateParameters struct {
	MTRDeviceControllerParameters
}

// MTRDeviceControllerExternalCertificateParametersFrom constructs a [MTRDeviceControllerExternalCertificateParameters] from an unsafe.Pointer.
func MTRDeviceControllerExternalCertificateParametersFrom(ptr unsafe.Pointer) MTRDeviceControllerExternalCertificateParameters {
	return MTRDeviceControllerExternalCertificateParameters{
		MTRDeviceControllerParameters: MTRDeviceControllerParametersFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerExternalCertificateParametersClass) Alloc() MTRDeviceControllerExternalCertificateParameters {
	rv := objc.Send[MTRDeviceControllerExternalCertificateParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceControllerExternalCertificateParametersClass) New() MTRDeviceControllerExternalCertificateParameters {
	rv := objc.Send[MTRDeviceControllerExternalCertificateParameters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceControllerExternalCertificateParameters) Init() MTRDeviceControllerExternalCertificateParameters {
	rv := objc.Send[MTRDeviceControllerExternalCertificateParameters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceControllerExternalCertificateParameters) Autorelease() MTRDeviceControllerExternalCertificateParameters {
	rv := objc.Send[MTRDeviceControllerExternalCertificateParameters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceControllerExternalCertificateParameters creates a new MTRDeviceControllerExternalCertificateParameters instance.
func NewMTRDeviceControllerExternalCertificateParameters() MTRDeviceControllerExternalCertificateParameters {
	return getMTRDeviceControllerExternalCertificateParametersClass().New()
}




