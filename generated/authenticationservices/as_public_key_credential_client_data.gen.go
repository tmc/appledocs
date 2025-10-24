// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPublicKeyCredentialClientData */


/* debug [class_header]: Header for ASPublicKeyCredentialClientData */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PublicKeyCredentialClientData */
// An interface definition for the [PublicKeyCredentialClientData] class.
type IPublicKeyCredentialClientData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PublicKeyCredentialClientData */
	// properties:
	Challenge() objc.IObject /* cross-framework: NSData */
	SetChallenge(value objc.IObject /* cross-framework: NSData */)
	CrossOrigin() PublicKeyCredentialClientDataCrossOriginValue
	SetCrossOrigin(value PublicKeyCredentialClientDataCrossOriginValue)
	Origin() objc.IObject /* cross-framework: NSString */
	SetOrigin(value objc.IObject /* cross-framework: NSString */)
	TopOrigin() objc.IObject /* cross-framework: NSString */
	SetTopOrigin(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PublicKeyCredentialClientData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PublicKeyCredentialClientData */
// Alloc allocates a new instance without initialization.
func (pc _PublicKeyCredentialClientDataClass) Alloc() PublicKeyCredentialClientData {
	rv := objc.Send[PublicKeyCredentialClientData](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PublicKeyCredentialClientData */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PublicKeyCredentialClientData */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class/initWithChallenge:origin:
func NewPublicKeyCredentialClientDataWithChallengeOrigin(challenge objc.IObject /* cross-framework: NSData */, origin objc.IObject /* cross-framework: NSString */) PublicKeyCredentialClientData {
	instance := getPublicKeyCredentialClientDataClass().Alloc()
	rv := objc.Send[PublicKeyCredentialClientData](instance.ID, objc.Sel("initWithChallenge:origin:"), challenge, origin)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPublicKeyCredentialClientDataWithChallengeOrigin */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PublicKeyCredentialClientData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PublicKeyCredentialClientData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PublicKeyCredentialClientData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PublicKeyCredentialClientData */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class/challenge
func (p_ PublicKeyCredentialClientData) Challenge() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("challenge"))
	return rv
}/* debug [instance_properties/getter]: challenge */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class/challenge
func (p_ PublicKeyCredentialClientData) SetChallenge(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setChallenge:"), value)
}/* debug [instance_properties/setter]: challenge */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class/crossOrigin
func (p_ PublicKeyCredentialClientData) CrossOrigin() PublicKeyCredentialClientDataCrossOriginValue {
	rv := objc.Send[PublicKeyCredentialClientDataCrossOriginValue](p_.ID, objc.Sel("crossOrigin"))
	return rv
}/* debug [instance_properties/getter]: crossOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class/crossOrigin
func (p_ PublicKeyCredentialClientData) SetCrossOrigin(value PublicKeyCredentialClientDataCrossOriginValue) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCrossOrigin:"), value)
}/* debug [instance_properties/setter]: crossOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class/origin
func (p_ PublicKeyCredentialClientData) Origin() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("origin"))
	return rv
}/* debug [instance_properties/getter]: origin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class/origin
func (p_ PublicKeyCredentialClientData) SetOrigin(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOrigin:"), value)
}/* debug [instance_properties/setter]: origin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class/topOrigin
func (p_ PublicKeyCredentialClientData) TopOrigin() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("topOrigin"))
	return rv
}/* debug [instance_properties/getter]: topOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientData-c.class/topOrigin
func (p_ PublicKeyCredentialClientData) SetTopOrigin(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTopOrigin:"), value)
}/* debug [instance_properties/setter]: topOrigin */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPublicKeyCredentialClientData */


