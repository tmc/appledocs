// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASOneTimeCodeCredentialRequest */


/* debug [class_header]: Header for ASOneTimeCodeCredentialRequest */
// The class instance for the [OneTimeCodeCredentialRequest] class.
var (
	OneTimeCodeCredentialRequestClass     _OneTimeCodeCredentialRequestClass
	OneTimeCodeCredentialRequestClassOnce sync.Once
)

func getOneTimeCodeCredentialRequestClass() _OneTimeCodeCredentialRequestClass {
	OneTimeCodeCredentialRequestClassOnce.Do(func() {
		OneTimeCodeCredentialRequestClass = _OneTimeCodeCredentialRequestClass{objc.GetClass("ASOneTimeCodeCredentialRequest")}
	})
	return OneTimeCodeCredentialRequestClass
}

type _OneTimeCodeCredentialRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OneTimeCodeCredentialRequest */
// An interface definition for the [OneTimeCodeCredentialRequest] class.
type IOneTimeCodeCredentialRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OneTimeCodeCredentialRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OneTimeCodeCredentialRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OneTimeCodeCredentialRequest */
// Alloc allocates a new instance without initialization.
func (oc _OneTimeCodeCredentialRequestClass) Alloc() OneTimeCodeCredentialRequest {
	rv := objc.Send[OneTimeCodeCredentialRequest](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OneTimeCodeCredentialRequestClass) New() OneTimeCodeCredentialRequest {
	rv := objc.Send[OneTimeCodeCredentialRequest](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OneTimeCodeCredentialRequest) Init() OneTimeCodeCredentialRequest {
	rv := objc.Send[OneTimeCodeCredentialRequest](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OneTimeCodeCredentialRequest) Autorelease() OneTimeCodeCredentialRequest {
	rv := objc.Send[OneTimeCodeCredentialRequest](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOneTimeCodeCredentialRequest creates a new OneTimeCodeCredentialRequest instance.
func NewOneTimeCodeCredentialRequest() OneTimeCodeCredentialRequest {
	return getOneTimeCodeCredentialRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OneTimeCodeCredentialRequest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredentialRequest
type OneTimeCodeCredentialRequest struct {
	objectivec.Object
}

// OneTimeCodeCredentialRequestFrom constructs a [OneTimeCodeCredentialRequest] from an unsafe.Pointer.
func OneTimeCodeCredentialRequestFrom(ptr unsafe.Pointer) OneTimeCodeCredentialRequest {
	return OneTimeCodeCredentialRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OneTimeCodeCredentialRequest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredentialRequest/init(credentialIdentity:)
func NewOneTimeCodeCredentialRequestWithCredentialIdentity(credentialIdentity IASOneTimeCodeCredentialIdentity) OneTimeCodeCredentialRequest {
	instance := getOneTimeCodeCredentialRequestClass().Alloc()
	rv := objc.Send[OneTimeCodeCredentialRequest](instance.ID, objc.Sel("initWithCredentialIdentity:"), credentialIdentity)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOneTimeCodeCredentialRequestWithCredentialIdentity */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OneTimeCodeCredentialRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OneTimeCodeCredentialRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OneTimeCodeCredentialRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OneTimeCodeCredentialRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASOneTimeCodeCredentialRequest */


