// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

package collaboration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CBIdentity] class.
var (
	CBIdentityClass     _CBIdentityClass
	CBIdentityClassOnce sync.Once
)

func getCBIdentityClass() _CBIdentityClass {
	CBIdentityClassOnce.Do(func() {
		CBIdentityClass = _CBIdentityClass{objc.GetClass("CBIdentity")}
	})
	return CBIdentityClass
}

type _CBIdentityClass struct {
	class objc.Class
}

// An interface definition for the [CBIdentity] class.
type ICBIdentity interface {
	objectivec.IObject
	IsMemberOfGroup(group unsafe.Pointer) bool
}

// A object is used for accessing the attributes of an identity stored in an identity authority. You can use an identity object for finding identities, and storing them in an access control list (ACL). If you need to edit these attributes, take advantage of the class in Core Services.
//
// You can obtain a object from one of the following class factory methods: , , , or . A object has methods to support for interoperability with the Core Services Identity API. Send to your object to return an opaque object for use in the Core Services Identity API. Similarly, call to use an Core Services Identity opaque object in the Collaboration framework. There are two subclasses of : and . If you are working specifically with a group identity, use . Similarly, if you are working with a user identity, use .
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity
type CBIdentity struct {
	objectivec.Object
}

// CBIdentityFrom constructs a [CBIdentity] from an unsafe.Pointer.
//
// A object is used for accessing the attributes of an identity stored in an identity authority. You can use an identity object for finding identities, and storing them in an access control list (ACL). If you need to edit these attributes, take advantage of the class in Core Services.
func CBIdentityFrom(ptr unsafe.Pointer) CBIdentity {
	return CBIdentity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CBIdentityClass) Alloc() CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBIdentityClass) New() CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBIdentity) Init() CBIdentity {
	rv := objc.Send[CBIdentity](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBIdentity) Autorelease() CBIdentity {
	rv := objc.Send[CBIdentity](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBIdentity creates a new CBIdentity instance.
func NewCBIdentity() CBIdentity {
	return getCBIdentityClass().New()
}


// Returns the identity object with the given name from the specified identity authority.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(name:authority:)
func NewCBIdentityWithNameAuthority(name string, authority unsafe.Pointer) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(getCBIdentityClass().class), objc.Sel("identityWithName:authority:"), objc.String(name), authority)
	return rv
}

// Returns the identity object matching the persistent reference data.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(persistentReference:)
func NewCBIdentityWithPersistentReference(data unsafe.Pointer) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(getCBIdentityClass().class), objc.Sel("identityWithPersistentReference:"), data)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(uniqueIdentifier:authority:)
func NewCBIdentityWithUniqueIdentifierAuthority(uuid unsafe.Pointer, authority unsafe.Pointer) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(getCBIdentityClass().class), objc.Sel("identityWithUniqueIdentifier:authority:"), uuid, authority)
	return rv
}

// Returns the identity object with the given UUID from the specified identity authority.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(uuidString:authority:)
func NewCBIdentityWithUUIDStringAuthority(uuid string, authority unsafe.Pointer) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(getCBIdentityClass().class), objc.Sel("identityWithUUIDString:authority:"), objc.String(uuid), authority)
	return rv
}


// Returns an identity object created from the specified Core Services Identity opaque object.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/identityWithCSIdentity:
func (cc _CBIdentityClass) IdentityWithCSIdentity(csIdentity unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("identityWithCSIdentity:"), csIdentity)
	return rv
}

// Returns the identity object with the given name from the specified identity authority.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(name:authority:)
func (cc _CBIdentityClass) IdentityWithNameAuthority(name string, authority unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("identityWithName:authority:"), objc.String(name), authority)
	return rv
}

// Returns the identity object matching the persistent reference data.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(persistentReference:)
func (cc _CBIdentityClass) IdentityWithPersistentReference(data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("identityWithPersistentReference:"), data)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(uniqueIdentifier:authority:)
func (cc _CBIdentityClass) IdentityWithUniqueIdentifierAuthority(uuid unsafe.Pointer, authority unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("identityWithUniqueIdentifier:authority:"), uuid, authority)
	return rv
}

// Returns the identity object with the given UUID from the specified identity authority.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(uuidString:authority:)
func (cc _CBIdentityClass) IdentityWithUUIDStringAuthority(uuid string, authority unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("identityWithUUIDString:authority:"), objc.String(uuid), authority)
	return rv
}

// Returns a Boolean value indicating whether the identity is a member of the specified group.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/isMember(ofGroup:)
func (c_ CBIdentity) IsMemberOfGroup(group unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMemberOfGroup:"), group)
	return rv
}

// Returns an opaque object for use with the Core Services Identity API.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/CSIdentity
func (c_ CBIdentity) CSIdentity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("CSIdentity"))
	return rv
}

// Returns an array of aliases (alternate names) for the identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/aliases
func (c_ CBIdentity) Aliases() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("aliases"))
	return rv
}

// Returns the identity authority where the identity is stored.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/authority
func (c_ CBIdentity) Authority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("authority"))
	return rv
}

// Returns the email address of an identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/emailAddress
func (c_ CBIdentity) EmailAddress() string {
	rv := objc.Send[string](c_.ID, objc.Sel("emailAddress"))
	return rv
}

// Returns the full name of the identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/fullName
func (c_ CBIdentity) FullName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("fullName"))
	return rv
}

// Returns the image associated with an identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/image
func (c_ CBIdentity) Image() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("image"))
	return rv
}

// Returns a Boolean value indicating the state of the identity’s hidden property.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/isHidden
func (c_ CBIdentity) Hidden() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hidden"))
	return rv
}

// Returns a persistent reference to store a reference to an identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/persistentReference
func (c_ CBIdentity) PersistentReference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("persistentReference"))
	return rv
}

// Returns the POSIX name of the identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/posixName
func (c_ CBIdentity) PosixName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("posixName"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/uniqueIdentifier
func (c_ CBIdentity) UniqueIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}

// Returns the UUID of the identity as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/uuidString
func (c_ CBIdentity) UUIDString() string {
	rv := objc.Send[string](c_.ID, objc.Sel("UUIDString"))
	return rv
}


