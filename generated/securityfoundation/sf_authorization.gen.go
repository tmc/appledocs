// Code generated from Apple documentation for SecurityFoundation. DO NOT EDIT.

package securityfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFAuthorization] class.
var (
	SFAuthorizationClass     _SFAuthorizationClass
	SFAuthorizationClassOnce sync.Once
)

func getSFAuthorizationClass() _SFAuthorizationClass {
	SFAuthorizationClassOnce.Do(func() {
		SFAuthorizationClass = _SFAuthorizationClass{objc.GetClass("SFAuthorization")}
	})
	return SFAuthorizationClass
}

type _SFAuthorizationClass struct {
	class objc.Class
}

// An interface definition for the [SFAuthorization] class.
type ISFAuthorization interface {
	objectivec.IObject
	// properties:
	// methods:
	AuthorizationRef() unsafe.Pointer
	InvalidateCredentials()
	ObtainWithRightFlagsError(rightName unsafe.Pointer, flags unsafe.Pointer, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	ObtainWithRightsFlagsEnvironmentAuthorizedRightsError(rights unsafe.Pointer, flags unsafe.Pointer, environment unsafe.Pointer, authorizedRights unsafe.Pointer, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
}

// A class that allows you to restrict a user’s access to particular features in your Mac app or daemon.
//
// The class is an interface for some of the functions in the Authorization Services API. You can use the method to obtain an authorization reference, used in other calls to Authorization Services functions. The Authorization Services API is documented in .


// A class that allows you to restrict a user’s access to particular features in your Mac app or daemon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization
type SFAuthorization struct {
	objectivec.Object
}

// SFAuthorizationFrom constructs a [SFAuthorization] from an unsafe.Pointer.
//
// A class that allows you to restrict a user’s access to particular features in your Mac app or daemon.
func SFAuthorizationFrom(ptr unsafe.Pointer) SFAuthorization {
	return SFAuthorization{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFAuthorizationClass) Alloc() SFAuthorization {
	rv := objc.Send[SFAuthorization](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFAuthorizationClass) New() SFAuthorization {
	rv := objc.Send[SFAuthorization](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFAuthorization) Init() SFAuthorization {
	rv := objc.Send[SFAuthorization](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFAuthorization) Autorelease() SFAuthorization {
	rv := objc.Send[SFAuthorization](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFAuthorization creates a new SFAuthorization instance.
func NewSFAuthorization() SFAuthorization {
	return getSFAuthorizationClass().New()
}



// Initializes an authorization object with the specified flags, rights, and environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/init(flags:rights:environment:)
func NewSFAuthorizationWithFlagsRightsEnvironment(flags unsafe.Pointer, rights unsafe.Pointer, environment unsafe.Pointer) SFAuthorization {
	instance := getSFAuthorizationClass().Alloc()
	rv := objc.Send[SFAuthorization](instance.ID, objc.Sel("initWithFlags:rights:environment:"), flags, rights, environment)
	rv.Autorelease()
	return rv
}



// Returns an authorization object initialized with a default environment, flags, and rights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/authorization()
func (sc _SFAuthorizationClass) Authorization() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("authorization"))
	return rv
}


// Returns an authorization object initialized with the specified flags, rights and environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/authorization(with:rights:environment:)
func (sc _SFAuthorizationClass) AuthorizationWithFlagsRightsEnvironment(flags unsafe.Pointer, rights unsafe.Pointer, environment unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("authorizationWithFlags:rights:environment:"), flags, rights, environment)
	return rv
}


// Returns the authorization reference for this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/authorizationRef()
func (s_ SFAuthorization) AuthorizationRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("authorizationRef"))
	return rv
}


// Prevents any rights that were obtained by this object from being preserved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/invalidateCredentials()
func (s_ SFAuthorization) InvalidateCredentials() {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidateCredentials"))
}


// Authorizes and preauthorizes one specific right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/obtain(withRight:flags:)
func (s_ SFAuthorization) ObtainWithRightFlagsError(rightName unsafe.Pointer, flags unsafe.Pointer, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("obtainWithRight:flags:error:"), rightName, flags, error_)
	return rv
}


// Authorizes and preauthorizes rights to access a privileged operation and returns the granted rights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/obtain(withRights:flags:environment:authorizedRights:)
func (s_ SFAuthorization) ObtainWithRightsFlagsEnvironmentAuthorizedRightsError(rights unsafe.Pointer, flags unsafe.Pointer, environment unsafe.Pointer, authorizedRights unsafe.Pointer, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("obtainWithRights:flags:environment:authorizedRights:error:"), rights, flags, environment, authorizedRights, error_)
	return rv
}


