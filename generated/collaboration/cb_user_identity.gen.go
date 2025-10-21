// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

package collaboration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CBUserIdentity] class.
var (
	CBUserIdentityClass     _CBUserIdentityClass
	CBUserIdentityClassOnce sync.Once
)

func getCBUserIdentityClass() _CBUserIdentityClass {
	CBUserIdentityClassOnce.Do(func() {
		CBUserIdentityClass = _CBUserIdentityClass{objc.GetClass("CBUserIdentity")}
	})
	return CBUserIdentityClass
}

type _CBUserIdentityClass struct {
	class objc.Class
}

// An interface definition for the [CBUserIdentity] class.
type ICBUserIdentity interface {
	ICBIdentity
	AuthenticateWithPassword(password string) bool
}

// An object of the class represents a user identity and is used for accessing the attributes of a user identity from an identity authority. The principal attributes of are a POSIX user identifier (UID), password, and certificate.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity
type CBUserIdentity struct {
	CBIdentity
}

// CBUserIdentityFrom constructs a [CBUserIdentity] from an unsafe.Pointer.
//
// An object of the class represents a user identity and is used for accessing the attributes of a user identity from an identity authority. The principal attributes of are a POSIX user identifier (UID), password, and certificate.
func CBUserIdentityFrom(ptr unsafe.Pointer) CBUserIdentity {
	return CBUserIdentity{
		CBIdentity: CBIdentityFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CBUserIdentityClass) Alloc() CBUserIdentity {
	rv := objc.Send[CBUserIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBUserIdentityClass) New() CBUserIdentity {
	rv := objc.Send[CBUserIdentity](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBUserIdentity) Init() CBUserIdentity {
	rv := objc.Send[CBUserIdentity](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBUserIdentity) Autorelease() CBUserIdentity {
	rv := objc.Send[CBUserIdentity](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBUserIdentity creates a new CBUserIdentity instance.
func NewCBUserIdentity() CBUserIdentity {
	return getCBUserIdentityClass().New()
}


// Returns a Boolean value indicating whether the given password is correct for the identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity/authenticate(withPassword:)
func (c_ CBUserIdentity) AuthenticateWithPassword(password string) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("authenticateWithPassword:"), objc.String(password))
	return rv
}

// Returns the public authentication certificate associated with a user identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity/certificate
func (c_ CBUserIdentity) Certificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("certificate"))
	return rv
}

// Returns a Boolean value indicating whether the identity is allowed to authenticate.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity/isEnabled
func (c_ CBUserIdentity) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}

// Returns the POSIX UID of the identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity/posixUID
func (c_ CBUserIdentity) PosixUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("posixUID"))
	return rv
}


