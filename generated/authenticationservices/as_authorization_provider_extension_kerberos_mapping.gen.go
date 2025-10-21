// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [AuthorizationProviderExtensionKerberosMapping] class.
type IAuthorizationProviderExtensionKerberosMapping interface {
	objectivec.IObject
}

// A set of Kerberos mappings that the system login process uses.
//
// This class contains a set of mappings for the sign-on token when importing the Kerberos ticket.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationProviderExtensionKerberosMappingClass) Alloc() AuthorizationProviderExtensionKerberosMapping {
	rv := objc.Send[AuthorizationProviderExtensionKerberosMapping](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The key name of the Kerberos client name string.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/clientNameKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) ClientNameKeyName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("clientNameKeyName"))
	return rv
}


// SetClientNameKeyName sets the value of the clientNameKeyName property.
// The key name of the Kerberos client name string.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/clientNameKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetClientNameKeyName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setClientNameKeyName:"), objc.String(value))
}

// The key name of the Kerberos session key type number.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/encryptionKeyTypeKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) EncryptionKeyTypeKeyName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("encryptionKeyTypeKeyName"))
	return rv
}


// SetEncryptionKeyTypeKeyName sets the value of the encryptionKeyTypeKeyName property.
// The key name of the Kerberos session key type number.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/encryptionKeyTypeKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetEncryptionKeyTypeKeyName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEncryptionKeyTypeKeyName:"), objc.String(value))
}

// The key name of the Base 64-encoded Kerberos AS-REP string.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/messageBufferKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) MessageBufferKeyName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("messageBufferKeyName"))
	return rv
}


// SetMessageBufferKeyName sets the value of the messageBufferKeyName property.
// The key name of the Base 64-encoded Kerberos AS-REP string.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/messageBufferKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetMessageBufferKeyName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMessageBufferKeyName:"), objc.String(value))
}

// The key name of the Kerberos realm string.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/realmKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) RealmKeyName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("realmKeyName"))
	return rv
}


// SetRealmKeyName sets the value of the realmKeyName property.
// The key name of the Kerberos realm string.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/realmKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetRealmKeyName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRealmKeyName:"), objc.String(value))
}

// The key name of the Kerberos service name string.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/serviceNameKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) ServiceNameKeyName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("serviceNameKeyName"))
	return rv
}


// SetServiceNameKeyName sets the value of the serviceNameKeyName property.
// The key name of the Kerberos service name string.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/serviceNameKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetServiceNameKeyName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setServiceNameKeyName:"), objc.String(value))
}

// The key name of the Kerberos session key.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/sessionKeyKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SessionKeyKeyName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("sessionKeyKeyName"))
	return rv
}


// SetSessionKeyKeyName sets the value of the sessionKeyKeyName property.
// The key name of the Kerberos session key.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/sessionKeyKeyName
func (a_ AuthorizationProviderExtensionKerberosMapping) SetSessionKeyKeyName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSessionKeyKeyName:"), objc.String(value))
}

// The keypath in the response JSON that uses this set of mappings.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/ticketKeyPath
func (a_ AuthorizationProviderExtensionKerberosMapping) TicketKeyPath() string {
	rv := objc.Send[string](a_.ID, objc.Sel("ticketKeyPath"))
	return rv
}


// SetTicketKeyPath sets the value of the ticketKeyPath property.
// The keypath in the response JSON that uses this set of mappings.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKerberosMapping/ticketKeyPath
func (a_ AuthorizationProviderExtensionKerberosMapping) SetTicketKeyPath(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTicketKeyPath:"), objc.String(value))
}



