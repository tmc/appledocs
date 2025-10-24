// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationProviderExtensionKerberosMapping */


/* debug [class_header]: Header for ASAuthorizationProviderExtensionKerberosMapping */
// The class instance for the [AuthorizationProviderExtensionKerberosMapping] class.
var (
	AuthorizationProviderExtensionKerberosMappingClass     _AuthorizationProviderExtensionKerberosMappingClass
	AuthorizationProviderExtensionKerberosMappingClassOnce sync.Once
)

func getAuthorizationProviderExtensionKerberosMappingClass() _AuthorizationProviderExtensionKerberosMappingClass {
	AuthorizationProviderExtensionKerberosMappingClassOnce.Do(func() {
		AuthorizationProviderExtensionKerberosMappingClass = _AuthorizationProviderExtensionKerberosMappingClass{objc.GetClass("ASAuthorizationProviderExtensionKerberosMapping")}
	})
	return AuthorizationProviderExtensionKerberosMappingClass
}

type _AuthorizationProviderExtensionKerberosMappingClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationProviderExtensionKerberosMapping */
// An interface definition for the [AuthorizationProviderExtensionKerberosMapping] class.
type IAuthorizationProviderExtensionKerberosMapping interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationProviderExtensionKerberosMapping */
	// properties:
	ClientNameKeyName() objc.IObject /* cross-framework: NSString */
	SetClientNameKeyName(value objc.IObject /* cross-framework: NSString */)
	EncryptionKeyTypeKeyName() objc.IObject /* cross-framework: NSString */
	SetEncryptionKeyTypeKeyName(value objc.IObject /* cross-framework: NSString */)
	MessageBufferKeyName() objc.IObject /* cross-framework: NSString */
	SetMessageBufferKeyName(value objc.IObject /* cross-framework: NSString */)
	RealmKeyName() objc.IObject /* cross-framework: NSString */
	SetRealmKeyName(value objc.IObject /* cross-framework: NSString */)
	ServiceNameKeyName() objc.IObject /* cross-framework: NSString */
	SetServiceNameKeyName(value objc.IObject /* cross-framework: NSString */)
	SessionKeyKeyName() objc.IObject /* cross-framework: NSString */
	SetSessionKeyKeyName(value objc.IObject /* cross-framework: NSString */)
	TicketKeyPath() objc.IObject /* cross-framework: NSString */
	SetTicketKeyPath(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationProviderExtensionKerberosMapping */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationProviderExtensionKerberosMapping */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationProviderExtensionKerberosMappingClass) Alloc() AuthorizationProviderExtensionKerberosMapping {
	rv := objc.Send[AuthorizationProviderExtensionKerberosMapping](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationProviderExtensionKerberosMappingClass) New() AuthorizationProviderExtensionKerberosMapping {
	rv := objc.Send[AuthorizationProviderExtensionKerberosMapping](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationProviderExtensionKerberosMapping) Init() AuthorizationProviderExtensionKerberosMapping {
	rv := objc.Send[AuthorizationProviderExtensionKerberosMapping](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationProviderExtensionKerberosMapping) Autorelease() AuthorizationProviderExtensionKerberosMapping {
	rv := objc.Send[AuthorizationProviderExtensionKerberosMapping](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationProviderExtensionKerberosMapping creates a new AuthorizationProviderExtensionKerberosMapping instance.
func NewAuthorizationProviderExtensionKerberosMapping() AuthorizationProviderExtensionKerberosMapping {
	return getAuthorizationProviderExtensionKerberosMappingClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationProviderExtensionKerberosMapping */
// A set of Kerberos mappings that the system login process uses.
//
// This class contains a set of mappings for the sign-on token when importing the Kerberos ticket.


// A set of Kerberos mappings that the system login process uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping
type AuthorizationProviderExtensionKerberosMapping struct {
	objectivec.Object
}

// AuthorizationProviderExtensionKerberosMappingFrom constructs a [AuthorizationProviderExtensionKerberosMapping] from an unsafe.Pointer.
//
// A set of Kerberos mappings that the system login process uses.
func AuthorizationProviderExtensionKerberosMappingFrom(ptr unsafe.Pointer) AuthorizationProviderExtensionKerberosMapping {
	return AuthorizationProviderExtensionKerberosMapping{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationProviderExtensionKerberosMapping *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationProviderExtensionKerberosMapping */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationProviderExtensionKerberosMapping */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationProviderExtensionKerberosMapping */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationProviderExtensionKerberosMapping */

// The key name of the Kerberos client name string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/clientNameKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) ClientNameKeyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("clientNameKeyName"))
	return rv
}/* debug [instance_properties/getter]: clientNameKeyName */


// The key name of the Kerberos client name string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/clientNameKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetClientNameKeyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setClientNameKeyName:"), value)
}/* debug [instance_properties/setter]: clientNameKeyName */


// The key name of the Kerberos session key type number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/encryptionKeyTypeKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) EncryptionKeyTypeKeyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("encryptionKeyTypeKeyName"))
	return rv
}/* debug [instance_properties/getter]: encryptionKeyTypeKeyName */


// The key name of the Kerberos session key type number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/encryptionKeyTypeKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetEncryptionKeyTypeKeyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEncryptionKeyTypeKeyName:"), value)
}/* debug [instance_properties/setter]: encryptionKeyTypeKeyName */


// The key name of the Base 64-encoded Kerberos AS-REP string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/messageBufferKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) MessageBufferKeyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("messageBufferKeyName"))
	return rv
}/* debug [instance_properties/getter]: messageBufferKeyName */


// The key name of the Base 64-encoded Kerberos AS-REP string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/messageBufferKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetMessageBufferKeyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMessageBufferKeyName:"), value)
}/* debug [instance_properties/setter]: messageBufferKeyName */


// The key name of the Kerberos realm string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/realmKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) RealmKeyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("realmKeyName"))
	return rv
}/* debug [instance_properties/getter]: realmKeyName */


// The key name of the Kerberos realm string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/realmKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetRealmKeyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRealmKeyName:"), value)
}/* debug [instance_properties/setter]: realmKeyName */


// The key name of the Kerberos service name string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/serviceNameKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) ServiceNameKeyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("serviceNameKeyName"))
	return rv
}/* debug [instance_properties/getter]: serviceNameKeyName */


// The key name of the Kerberos service name string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/serviceNameKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetServiceNameKeyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setServiceNameKeyName:"), value)
}/* debug [instance_properties/setter]: serviceNameKeyName */


// The key name of the Kerberos session key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/sessionKeyKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SessionKeyKeyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("sessionKeyKeyName"))
	return rv
}/* debug [instance_properties/getter]: sessionKeyKeyName */


// The key name of the Kerberos session key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/sessionKeyKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetSessionKeyKeyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSessionKeyKeyName:"), value)
}/* debug [instance_properties/setter]: sessionKeyKeyName */


// The keypath in the response JSON that uses this set of mappings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/ticketKeyPath
func (a_ AuthorizationProviderExtensionKerberosMapping) TicketKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("ticketKeyPath"))
	return rv
}/* debug [instance_properties/getter]: ticketKeyPath */


// The keypath in the response JSON that uses this set of mappings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/ticketKeyPath
func (a_ AuthorizationProviderExtensionKerberosMapping) SetTicketKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTicketKeyPath:"), value)
}/* debug [instance_properties/setter]: ticketKeyPath */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationProviderExtensionKerberosMapping */



