// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Authorization] class.
var (
	AuthorizationClass     _AuthorizationClass
	AuthorizationClassOnce sync.Once
)

func getAuthorizationClass() _AuthorizationClass {
	AuthorizationClassOnce.Do(func() {
		AuthorizationClass = _AuthorizationClass{objc.GetClass("ASAuthorization")}
	})
	return AuthorizationClass
}

type _AuthorizationClass struct {
	class objc.Class
}

// An interface definition for the [Authorization] class.
type IAuthorization interface {
	objectivec.IObject
	// properties:
	Credential() AuthorizationCredential /* not a class type */
	SetCredential(value AuthorizationCredential /* not a class type */)
	Provider() AuthorizationProvider /* not a class type */
	SetProvider(value AuthorizationProvider /* not a class type */)
	// methods:
}

// The encapsulation of a successful authorization by a controller.


// The encapsulation of a successful authorization by a controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorization
type Authorization struct {
	objectivec.Object
}

// AuthorizationFrom constructs a [Authorization] from an unsafe.Pointer.
//
// The encapsulation of a successful authorization by a controller.
func AuthorizationFrom(ptr unsafe.Pointer) Authorization {
	return Authorization{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationClass) Alloc() Authorization {
	rv := objc.Send[Authorization](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationClass) New() Authorization {
	rv := objc.Send[Authorization](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Authorization) Init() Authorization {
	rv := objc.Send[Authorization](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Authorization) Autorelease() Authorization {
	rv := objc.Send[Authorization](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorization creates a new Authorization instance.
func NewAuthorization() Authorization {
	return getAuthorizationClass().New()
}



// Information provided about a user after successful authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorization/credential
func (a_ Authorization) Credential() AuthorizationCredential /* not a class type */ {
	rv := objc.Send[AuthorizationCredential](a_.ID, objc.Sel("credential"))
	return rv
}


// Information provided about a user after successful authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorization/credential
func (a_ Authorization) SetCredential(value AuthorizationCredential /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCredential:"), value)
}


// The provider that created the request that resulted in the successful authorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorization/provider
func (a_ Authorization) Provider() AuthorizationProvider /* not a class type */ {
	rv := objc.Send[AuthorizationProvider](a_.ID, objc.Sel("provider"))
	return rv
}


// The provider that created the request that resulted in the successful authorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorization/provider
func (a_ Authorization) SetProvider(value AuthorizationProvider /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProvider:"), value)
}



