// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PublicKeyCredentialClientData] class.
var (
	PublicKeyCredentialClientDataClass     _PublicKeyCredentialClientDataClass
	PublicKeyCredentialClientDataClassOnce sync.Once
)

func getPublicKeyCredentialClientDataClass() _PublicKeyCredentialClientDataClass {
	PublicKeyCredentialClientDataClassOnce.Do(func() {
		PublicKeyCredentialClientDataClass = _PublicKeyCredentialClientDataClass{objc.GetClass("ASPublicKeyCredentialClientData")}
	})
	return PublicKeyCredentialClientDataClass
}

type _PublicKeyCredentialClientDataClass struct {
	class objc.Class
}

// An interface definition for the [PublicKeyCredentialClientData] class.
type IPublicKeyCredentialClientData interface {
	objectivec.IObject
	// properties:
	Origin() string /* primitive/slice/pointer. */
	SetOrigin(value string /* primitive/slice/pointer. */)
	// methods:
}

// This object represents the client data for a public key credential request, as defined in the WebAuthentication standard.


// This object represents the client data for a public key credential request, as defined in the WebAuthentication standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class
type PublicKeyCredentialClientData struct {
	objectivec.Object
}

// PublicKeyCredentialClientDataFrom constructs a [PublicKeyCredentialClientData] from an unsafe.Pointer.
//
// This object represents the client data for a public key credential request, as defined in the WebAuthentication standard.
func PublicKeyCredentialClientDataFrom(ptr unsafe.Pointer) PublicKeyCredentialClientData {
	return PublicKeyCredentialClientData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PublicKeyCredentialClientDataClass) Alloc() PublicKeyCredentialClientData {
	rv := objc.Send[PublicKeyCredentialClientData](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PublicKeyCredentialClientDataClass) New() PublicKeyCredentialClientData {
	rv := objc.Send[PublicKeyCredentialClientData](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PublicKeyCredentialClientData) Init() PublicKeyCredentialClientData {
	rv := objc.Send[PublicKeyCredentialClientData](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PublicKeyCredentialClientData) Autorelease() PublicKeyCredentialClientData {
	rv := objc.Send[PublicKeyCredentialClientData](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPublicKeyCredentialClientData creates a new PublicKeyCredentialClientData instance.
func NewPublicKeyCredentialClientData() PublicKeyCredentialClientData {
	return getPublicKeyCredentialClientDataClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class/origin
func (p_ PublicKeyCredentialClientData) Origin() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("origin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class/origin
func (p_ PublicKeyCredentialClientData) SetOrigin(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOrigin:"), objc.String(value))
}



