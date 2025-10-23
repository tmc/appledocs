// Code generated from Apple documentation for DeviceCheck. DO NOT EDIT.

package devicecheck

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DCAppAttestService] class.
var (
	DCAppAttestServiceClass     _DCAppAttestServiceClass
	DCAppAttestServiceClassOnce sync.Once
)

func getDCAppAttestServiceClass() _DCAppAttestServiceClass {
	DCAppAttestServiceClassOnce.Do(func() {
		DCAppAttestServiceClass = _DCAppAttestServiceClass{objc.GetClass("DCAppAttestService")}
	})
	return DCAppAttestServiceClass
}

type _DCAppAttestServiceClass struct {
	class objc.Class
}

// An interface definition for the [DCAppAttestService] class.
type IDCAppAttestService interface {
	objectivec.IObject
	Supported() bool
	IsSupported() bool
	SetIsSupported(value bool)
	AttestKeyClientDataHashCompletionHandler(keyId string, clientDataHash foundation.NSData, completionHandler unsafe.Pointer)
	GenerateAssertionClientDataHashCompletionHandler(keyId string, clientDataHash foundation.NSData, completionHandler unsafe.Pointer)
	GenerateKeyWithCompletionHandler(completionHandler unsafe.Pointer)
}

// A service that you use to validate the instance of your app running on a device.
//
// Use the instance of the class to assert the legitimacy of a particular instance of your app to your server. After ensuring service availability by reading the property, you use the service to: Create a cryptographic key in the Secure Enclave by calling the method. Ask Apple to certify the key by calling the method. - Prepare an assertion of your app’s integrity to accompany any or all server requests using the method. For more information about how to support App Attest in your app, see . For information about the complementary procedures you implement on your server, see .


// A service that you use to validate the instance of your app running on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService
type DCAppAttestService struct {
	objectivec.Object
}

// DCAppAttestServiceFrom constructs a [DCAppAttestService] from an unsafe.Pointer.
//
// A service that you use to validate the instance of your app running on a device.
func DCAppAttestServiceFrom(ptr unsafe.Pointer) DCAppAttestService {
	return DCAppAttestService{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DCAppAttestServiceClass) Alloc() DCAppAttestService {
	rv := objc.Send[DCAppAttestService](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DCAppAttestServiceClass) New() DCAppAttestService {
	rv := objc.Send[DCAppAttestService](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DCAppAttestService) Init() DCAppAttestService {
	rv := objc.Send[DCAppAttestService](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DCAppAttestService) Autorelease() DCAppAttestService {
	rv := objc.Send[DCAppAttestService](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDCAppAttestService creates a new DCAppAttestService instance.
func NewDCAppAttestService() DCAppAttestService {
	return getDCAppAttestServiceClass().New()
}



// The shared App Attest service that you use to validate your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/shared
func (dc _DCAppAttestServiceClass) SharedService() DCAppAttestService {
	rv := objc.Send[DCAppAttestService](objc.ID(dc.class), objc.Sel("sharedService"))
	return rv
}

// Asks Apple to attest to the validity of a generated cryptographic key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/attestKey(_:clientDataHash:completionHandler:)
func (d_ DCAppAttestService) AttestKeyClientDataHashCompletionHandler(keyId string, clientDataHash foundation.NSData, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("attestKey:clientDataHash:completionHandler:"), objc.String(keyId), clientDataHash, completionHandler)
}


// Creates a block of data that demonstrates the legitimacy of an instance of your app running on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/generateAssertion(_:clientDataHash:completionHandler:)
func (d_ DCAppAttestService) GenerateAssertionClientDataHashCompletionHandler(keyId string, clientDataHash foundation.NSData, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("generateAssertion:clientDataHash:completionHandler:"), objc.String(keyId), clientDataHash, completionHandler)
}


// Creates a new cryptographic key for use with the App Attest service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/generateKey(completionHandler:)
func (d_ DCAppAttestService) GenerateKeyWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("generateKeyWithCompletionHandler:"), completionHandler)
}


// A Boolean value that indicates whether a particular device provides the App Attest service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/isSupported
func (d_ DCAppAttestService) Supported() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("supported"))
	return rv
}


// The shared App Attest service that you use to validate your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/shared
func (d_ DCAppAttestService) SharedService() IDCAppAttestService {
	rv := objc.Send[DCAppAttestService](d_.ID, objc.Sel("sharedService"))
	return rv
}


// A Boolean value that indicates whether a particular device provides the App
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/devicecheck/dcappattestservice/issupported
func (d_ DCAppAttestService) IsSupported() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isSupported"))
	return rv
}


// A Boolean value that indicates whether a particular device provides the App
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/devicecheck/dcappattestservice/issupported
func (d_ DCAppAttestService) SetIsSupported(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsSupported:"), value)
}



