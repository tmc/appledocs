// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

package collaboration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CBGroupIdentity] class.
var (
	CBGroupIdentityClass     _CBGroupIdentityClass
	CBGroupIdentityClassOnce sync.Once
)

func getCBGroupIdentityClass() _CBGroupIdentityClass {
	CBGroupIdentityClassOnce.Do(func() {
		CBGroupIdentityClass = _CBGroupIdentityClass{objc.GetClass("CBGroupIdentity")}
	})
	return CBGroupIdentityClass
}

type _CBGroupIdentityClass struct {
	class objc.Class
}

// An interface definition for the [CBGroupIdentity] class.
type ICBGroupIdentity interface {
	ICBIdentity
}

// An object of the class represents a group identity and is used for viewing the attributes of group identities from an identity authority. The principal attributes of a object are a POSIX group identifier (GID) and a list of members.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity
type CBGroupIdentity struct {
	CBIdentity
}

// CBGroupIdentityFrom constructs a [CBGroupIdentity] from an unsafe.Pointer.
//
// An object of the class represents a group identity and is used for viewing the attributes of group identities from an identity authority. The principal attributes of a object are a POSIX group identifier (GID) and a list of members.
func CBGroupIdentityFrom(ptr unsafe.Pointer) CBGroupIdentity {
	return CBGroupIdentity{
		CBIdentity: CBIdentityFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CBGroupIdentityClass) Alloc() CBGroupIdentity {
	rv := objc.Send[CBGroupIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBGroupIdentityClass) New() CBGroupIdentity {
	rv := objc.Send[CBGroupIdentity](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBGroupIdentity) Init() CBGroupIdentity {
	rv := objc.Send[CBGroupIdentity](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBGroupIdentity) Autorelease() CBGroupIdentity {
	rv := objc.Send[CBGroupIdentity](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBGroupIdentity creates a new CBGroupIdentity instance.
func NewCBGroupIdentity() CBGroupIdentity {
	return getCBGroupIdentityClass().New()
}


// Returns the group identity with the given POSIX GID in the specified identity authority.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity/init(posixGID:authority:)
func NewCBGroupIdentityWithPosixGIDAuthority(gid unsafe.Pointer, authority unsafe.Pointer) CBGroupIdentity {
	rv := objc.Send[CBGroupIdentity](objc.ID(getCBGroupIdentityClass().class), objc.Sel("groupIdentityWithPosixGID:authority:"), gid, authority)
	return rv
}


// Returns the group identity with the given POSIX GID in the specified identity authority.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity/init(posixGID:authority:)
func (cc _CBGroupIdentityClass) GroupIdentityWithPosixGIDAuthority(gid unsafe.Pointer, authority unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("groupIdentityWithPosixGID:authority:"), gid, authority)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity/memberIdentities
func (c_ CBGroupIdentity) MemberIdentities() []CBIdentity {
	rv := objc.Send[[]CBIdentity](c_.ID, objc.Sel("memberIdentities"))
	return rv
}

// Returns the members of the group.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity/members
func (c_ CBGroupIdentity) Members() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("members"))
	return rv
}

// Returns the POSIX GID of the identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity/posixGID
func (c_ CBGroupIdentity) PosixGID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("posixGID"))
	return rv
}


