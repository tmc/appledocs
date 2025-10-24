// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CXProvider */


/* debug [class_header]: Header for CXProvider */
// The class instance for the [CXProvider] class.
var (
	CXProviderClass     _CXProviderClass
	CXProviderClassOnce sync.Once
)

func getCXProviderClass() _CXProviderClass {
	CXProviderClassOnce.Do(func() {
		CXProviderClass = _CXProviderClass{objc.GetClass("CXProvider")}
	})
	return CXProviderClass
}

type _CXProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXProvider */
// An interface definition for the [CXProvider] class.
type ICXProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CXProvider */
	// properties:
	CXErrorDomain() objc.IObject /* cross-framework: NSString */
	CXErrorDomainIncomingCall() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXProvider */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXProvider */
// Alloc allocates a new instance without initialization.
func (cc _CXProviderClass) Alloc() CXProvider {
	rv := objc.Send[CXProvider](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXProviderClass) New() CXProvider {
	rv := objc.Send[CXProvider](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXProvider) Init() CXProvider {
	rv := objc.Send[CXProvider](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXProvider) Autorelease() CXProvider {
	rv := objc.Send[CXProvider](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXProvider creates a new CXProvider instance.
func NewCXProvider() CXProvider {
	return getCXProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXProvider */
// An object that represents a telephony provider.
//
// A object is responsible for reporting out-of-band notifications that occur to the system. A VoIP app should create only one instance of and store it for use globally. A object is initialized with a object to specify the behavior and capabilities of calls. Each provider can specify an object conforming to the protocol to respond to events, such as the call starting, the call being put on hold, or the provider’s audio session being activated.


// An object that represents a telephony provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider
type CXProvider struct {
	objectivec.Object
}

// CXProviderFrom constructs a [CXProvider] from an unsafe.Pointer.
//
// An object that represents a telephony provider.
func CXProviderFrom(ptr unsafe.Pointer) CXProvider {
	return CXProvider{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXProvider */

// Initializes a new provider with the specified configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/init(configuration:)
func NewCXProviderWithConfiguration(configuration ICXProviderConfiguration) CXProvider {
	instance := getCXProviderClass().Alloc()
	rv := objc.Send[CXProvider](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXProviderWithConfiguration */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXProvider */

// Reports a new incoming call after your notification service extension decrypts a VoIP call request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportNewIncomingVoIPPushPayload(_:completion:)
func (cc _CXProviderClass) ReportNewIncomingVoIPPushPayloadCompletion(dictionaryPayload objc.IObject /* cross-framework: NSDictionary */, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("reportNewIncomingVoIPPushPayload:completion:"), dictionaryPayload, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReportNewIncomingVoIPPushPayloadCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXProvider */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXProvider */

// The domain for CallKit errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxerrordomain
func (c_ CXProvider) CXErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CXErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: CXErrorDomain */


// The domain for errors that occur during incoming calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxerrordomainincomingcall
func (c_ CXProvider) CXErrorDomainIncomingCall() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CXErrorDomainIncomingCall"))
	return rv
}/* debug [instance_properties/getter]: CXErrorDomainIncomingCall */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXProvider */


