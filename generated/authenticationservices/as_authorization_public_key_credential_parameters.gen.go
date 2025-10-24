// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPublicKeyCredentialParameters */


/* debug [class_header]: Header for ASAuthorizationPublicKeyCredentialParameters */
// The class instance for the [AuthorizationPublicKeyCredentialParameters] class.
var (
	AuthorizationPublicKeyCredentialParametersClass     _AuthorizationPublicKeyCredentialParametersClass
	AuthorizationPublicKeyCredentialParametersClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialParametersClass() _AuthorizationPublicKeyCredentialParametersClass {
	AuthorizationPublicKeyCredentialParametersClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialParametersClass = _AuthorizationPublicKeyCredentialParametersClass{objc.GetClass("ASAuthorizationPublicKeyCredentialParameters")}
	})
	return AuthorizationPublicKeyCredentialParametersClass
}

type _AuthorizationPublicKeyCredentialParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPublicKeyCredentialParameters */
// An interface definition for the [AuthorizationPublicKeyCredentialParameters] class.
type IAuthorizationPublicKeyCredentialParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPublicKeyCredentialParameters */
	// properties:
	Algorithm() COSEAlgorithmIdentifier /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPublicKeyCredentialParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPublicKeyCredentialParameters */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialParametersClass) Alloc() AuthorizationPublicKeyCredentialParameters {
	rv := objc.Send[AuthorizationPublicKeyCredentialParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPublicKeyCredentialParametersClass) New() AuthorizationPublicKeyCredentialParameters {
	rv := objc.Send[AuthorizationPublicKeyCredentialParameters](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialParameters) Init() AuthorizationPublicKeyCredentialParameters {
	rv := objc.Send[AuthorizationPublicKeyCredentialParameters](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialParameters) Autorelease() AuthorizationPublicKeyCredentialParameters {
	rv := objc.Send[AuthorizationPublicKeyCredentialParameters](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialParameters creates a new AuthorizationPublicKeyCredentialParameters instance.
func NewAuthorizationPublicKeyCredentialParameters() AuthorizationPublicKeyCredentialParameters {
	return getAuthorizationPublicKeyCredentialParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPublicKeyCredentialParameters */
// An object that provides required parameters for the credential during registration.
//
// This object is mainly for signing algorithm negotiation, and is only relevant for physical security keys.


// An object that provides required parameters for the credential during registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialParameters
type AuthorizationPublicKeyCredentialParameters struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialParametersFrom constructs a [AuthorizationPublicKeyCredentialParameters] from an unsafe.Pointer.
//
// An object that provides required parameters for the credential during registration.
func AuthorizationPublicKeyCredentialParametersFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialParameters {
	return AuthorizationPublicKeyCredentialParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPublicKeyCredentialParameters */

// Creates the object with an algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialParameters/init(algorithm:)
func NewAuthorizationPublicKeyCredentialParametersWithAlgorithm(algorithm COSEAlgorithmIdentifier /* typedef */) AuthorizationPublicKeyCredentialParameters {
	instance := getAuthorizationPublicKeyCredentialParametersClass().Alloc()
	rv := objc.Send[AuthorizationPublicKeyCredentialParameters](instance.ID, objc.Sel("initWithAlgorithm:"), algorithm)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationPublicKeyCredentialParametersWithAlgorithm */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPublicKeyCredentialParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPublicKeyCredentialParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPublicKeyCredentialParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPublicKeyCredentialParameters */

// The algorithm to use for negitation between the authenticator and the relying party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialParameters/algorithm
func (a_ AuthorizationPublicKeyCredentialParameters) Algorithm() COSEAlgorithmIdentifier /* typedef */ {
	rv := objc.Send[int](a_.ID, objc.Sel("algorithm"))
	return rv
}/* debug [instance_properties/getter]: algorithm */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPublicKeyCredentialParameters */


