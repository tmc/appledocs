// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerparameters/certificationdeclarationcertificates
func (m_ MTRDeviceControllerParameters) CertificationDeclarationCertificates() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("certificationDeclarationCertificates"))
	return rv
}


// SetCertificationDeclarationCertificates sets the value of the certificationDeclarationCertificates property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerparameters/certificationdeclarationcertificates
func (m_ MTRDeviceControllerParameters) SetCertificationDeclarationCertificates(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificationDeclarationCertificates:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerparameters/concurrentsubscriptionestablishmentsallowedonthread
func (m_ MTRDeviceControllerParameters) ConcurrentSubscriptionEstablishmentsAllowedOnThread() int {
	rv := objc.Send[int](m_.ID, objc.Sel("concurrentSubscriptionEstablishmentsAllowedOnThread"))
	return rv
}


// SetConcurrentSubscriptionEstablishmentsAllowedOnThread sets the value of the concurrentSubscriptionEstablishmentsAllowedOnThread property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerparameters/concurrentsubscriptionestablishmentsallowedonthread
func (m_ MTRDeviceControllerParameters) SetConcurrentSubscriptionEstablishmentsAllowedOnThread(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConcurrentSubscriptionEstablishmentsAllowedOnThread:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerparameters/productattestationauthoritycertificates
func (m_ MTRDeviceControllerParameters) ProductAttestationAuthorityCertificates() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("productAttestationAuthorityCertificates"))
	return rv
}


// SetProductAttestationAuthorityCertificates sets the value of the productAttestationAuthorityCertificates property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerparameters/productattestationauthoritycertificates
func (m_ MTRDeviceControllerParameters) SetProductAttestationAuthorityCertificates(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductAttestationAuthorityCertificates:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerparameters/shouldadvertiseoperational
func (m_ MTRDeviceControllerParameters) ShouldAdvertiseOperational() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldAdvertiseOperational"))
	return rv
}


// SetShouldAdvertiseOperational sets the value of the shouldAdvertiseOperational property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerparameters/shouldadvertiseoperational
func (m_ MTRDeviceControllerParameters) SetShouldAdvertiseOperational(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldAdvertiseOperational:"), value)
}

// Sets the storage behavior configuration - see MTRDeviceStorageBehaviorConfiguration.h for details
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerparameters/storagebehaviorconfiguration
func (m_ MTRDeviceControllerParameters) StorageBehaviorConfiguration() MTRDeviceStorageBehaviorConfiguration {
	rv := objc.Send[MTRDeviceStorageBehaviorConfiguration](m_.ID, objc.Sel("storageBehaviorConfiguration"))
	return rv
}


// SetStorageBehaviorConfiguration sets the value of the storageBehaviorConfiguration property.
// Sets the storage behavior configuration - see MTRDeviceStorageBehaviorConfiguration.h for details

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerparameters/storagebehaviorconfiguration
func (m_ MTRDeviceControllerParameters) SetStorageBehaviorConfiguration(value IMTRDeviceStorageBehaviorConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStorageBehaviorConfiguration:"), value)
}



