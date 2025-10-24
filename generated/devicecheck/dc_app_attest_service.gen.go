// Code generated from Apple documentation for DeviceCheck. DO NOT EDIT.

package devicecheck

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DCAppAttestService */


/* debug [class_header]: Header for DCAppAttestService */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DCAppAttestService */
// An interface definition for the [DCAppAttestService] class.
type IDCAppAttestService interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DCAppAttestService */
	// properties:
	Supported() bool
	IsSupported() bool
	SetIsSupported(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DCAppAttestService */
	// methods:
	AttestKeyClientDataHashCompletionHandler(keyId objc.IObject /* cross-framework: NSString */, clientDataHash objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)
	GenerateAssertionClientDataHashCompletionHandler(keyId objc.IObject /* cross-framework: NSString */, clientDataHash objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)
	GenerateKeyWithCompletionHandler(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DCAppAttestService */
// Alloc allocates a new instance without initialization.
func (dc _DCAppAttestServiceClass) Alloc() DCAppAttestService {
	rv := objc.Send[DCAppAttestService](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DCAppAttestService */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DCAppAttestService *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DCAppAttestService */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DCAppAttestService */

// The shared App Attest service that you use to validate your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/shared
func (dc _DCAppAttestServiceClass) SharedService() DCAppAttestService {
	rv := objc.Send[DCAppAttestService](objc.ID(dc.class), objc.Sel("sharedService"))
	return rv
}/* debug [class_properties_class/property]: sharedService */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DCAppAttestService */

// Asks Apple to attest to the validity of a generated cryptographic key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/attestKey(_:clientDataHash:completionHandler:)
func (d_ DCAppAttestService) AttestKeyClientDataHashCompletionHandler(keyId objc.IObject /* cross-framework: NSString */, clientDataHash objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("attestKey:clientDataHash:completionHandler:"), keyId, clientDataHash, completionHandler)
}/* debug [instance_methods/method]: AttestKeyClientDataHashCompletionHandler */


// Creates a block of data that demonstrates the legitimacy of an instance of your app running on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/generateAssertion(_:clientDataHash:completionHandler:)
func (d_ DCAppAttestService) GenerateAssertionClientDataHashCompletionHandler(keyId objc.IObject /* cross-framework: NSString */, clientDataHash objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("generateAssertion:clientDataHash:completionHandler:"), keyId, clientDataHash, completionHandler)
}/* debug [instance_methods/method]: GenerateAssertionClientDataHashCompletionHandler */


// Creates a new cryptographic key for use with the App Attest service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/generateKey(completionHandler:)
func (d_ DCAppAttestService) GenerateKeyWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("generateKeyWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GenerateKeyWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DCAppAttestService */

// A Boolean value that indicates whether a particular device provides the App Attest service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/isSupported
func (d_ DCAppAttestService) Supported() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("supported"))
	return rv
}/* debug [instance_properties/getter]: supported */


// The shared App Attest service that you use to validate your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/shared
func (d_ DCAppAttestService) SharedService() IDCAppAttestService {
	rv := objc.Send[DCAppAttestService](d_.ID, objc.Sel("sharedService"))
	return rv
}/* debug [instance_properties/getter]: sharedService */


// A Boolean value that indicates whether a particular device provides the App
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/devicecheck/dcappattestservice/issupported
func (d_ DCAppAttestService) IsSupported() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isSupported"))
	return rv
}/* debug [instance_properties/getter]: isSupported */


// A Boolean value that indicates whether a particular device provides the App
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/devicecheck/dcappattestservice/issupported
func (d_ DCAppAttestService) SetIsSupported(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsSupported:"), value)
}/* debug [instance_properties/setter]: isSupported */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DCAppAttestService */



