// Code generated from Apple documentation for SecurityFoundation. DO NOT EDIT.

package securityfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFAuthorization */

/* debug [class_header]: Header for SFAuthorization */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SFAuthorization */
// An interface definition for the [SFAuthorization] class.
type ISFAuthorization interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for SFAuthorization */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SFAuthorization */
	// methods:
	AuthorizationRef() unsafe.Pointer
	InvalidateCredentials()
	ObtainWithRightFlagsError(rightName unsafe.Pointer, flags unsafe.Pointer, error_ unsafe.Pointer) bool
	ObtainWithRightsFlagsEnvironmentAuthorizedRightsError(rights unsafe.Pointer, flags unsafe.Pointer, environment unsafe.Pointer, authorizedRights unsafe.Pointer, error_ unsafe.Pointer) bool
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SFAuthorization */
// Alloc allocates a new instance without initialization.
func (sc _SFAuthorizationClass) Alloc() SFAuthorization {
	rv := objc.Send[SFAuthorization](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SFAuthorization */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SFAuthorization */

// Initializes an authorization object with the specified flags, rights, and environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/init(flags:rights:environment:)
func NewSFAuthorizationWithFlagsRightsEnvironment(flags unsafe.Pointer, rights unsafe.Pointer, environment unsafe.Pointer) SFAuthorization {
	instance := getSFAuthorizationClass().Alloc()
	rv := objc.Send[SFAuthorization](instance.ID, objc.Sel("initWithFlags:rights:environment:"), flags, rights, environment)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewSFAuthorizationWithFlagsRightsEnvironment */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SFAuthorization */

// Returns an authorization object initialized with a default environment, flags, and rights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/authorization()
func (sc _SFAuthorizationClass) Authorization() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("authorization"))
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=Authorization) */

// Returns an authorization object initialized with the specified flags, rights and environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/authorization(with:rights:environment:)
func (sc _SFAuthorizationClass) AuthorizationWithFlagsRightsEnvironment(flags unsafe.Pointer, rights unsafe.Pointer, environment unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("authorizationWithFlags:rights:environment:"), flags, rights, environment)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=AuthorizationWithFlagsRightsEnvironment) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SFAuthorization */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SFAuthorization */

// Returns the authorization reference for this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/authorizationRef()
func (s_ SFAuthorization) AuthorizationRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("authorizationRef"))
	return rv
} /* debug [instance_methods/method]: AuthorizationRef */

// Prevents any rights that were obtained by this object from being preserved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/invalidateCredentials()
func (s_ SFAuthorization) InvalidateCredentials() {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidateCredentials"))
} /* debug [instance_methods/method]: InvalidateCredentials */

// Authorizes and preauthorizes one specific right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/obtain(withRight:flags:)
func (s_ SFAuthorization) ObtainWithRightFlagsError(rightName unsafe.Pointer, flags unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("obtainWithRight:flags:error:"), rightName, flags, error_)
	return rv
} /* debug [instance_methods/method]: ObtainWithRightFlagsError */

// Authorizes and preauthorizes rights to access a privileged operation and returns the granted rights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation/SFAuthorization/obtain(withRights:flags:environment:authorizedRights:)
func (s_ SFAuthorization) ObtainWithRightsFlagsEnvironmentAuthorizedRightsError(rights unsafe.Pointer, flags unsafe.Pointer, environment unsafe.Pointer, authorizedRights unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("obtainWithRights:flags:environment:authorizedRights:error:"), rights, flags, environment, authorizedRights, error_)
	return rv
} /* debug [instance_methods/method]: ObtainWithRightsFlagsEnvironmentAuthorizedRightsError */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SFAuthorization */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SFAuthorization */
