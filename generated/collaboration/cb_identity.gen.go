// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

package collaboration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Authority() ICBIdentityAuthority
	Aliases() string /* primitive/slice/pointer. */
	SetAliases(value string /* primitive/slice/pointer. */)
	EmailAddress() string /* primitive/slice/pointer. */
	SetEmailAddress(value string /* primitive/slice/pointer. */)
	FullName() string /* primitive/slice/pointer. */
	SetFullName(value string /* primitive/slice/pointer. */)
	Image() appkit.objc.IObject /* cross-framework: Image */
	SetImage(value appkit.objc.IObject /* cross-framework: Image */)
	IsHidden() bool /* primitive/slice/pointer. */
	SetIsHidden(value bool /* primitive/slice/pointer. */)
	PersistentReference() foundation.objc.IObject /* cross-framework: Data */
	SetPersistentReference(value foundation.objc.IObject /* cross-framework: Data */)
	PosixName() string /* primitive/slice/pointer. */
	SetPosixName(value string /* primitive/slice/pointer. */)
	UniqueIdentifier() foundation.objc.IObject /* cross-framework: UUID */
	SetUniqueIdentifier(value foundation.objc.IObject /* cross-framework: UUID */)
	UuidString() string /* primitive/slice/pointer. */
	SetUuidString(value string /* primitive/slice/pointer. */)
	// methods:
}

// A object is used for accessing the attributes of an identity stored in an identity authority. You can use an identity object for finding identities, and storing them in an access control list (ACL). If you need to edit these attributes, take advantage of the class in Core Services.
//
// You can obtain a object from one of the following class factory methods: , , , or . A object has methods to support for interoperability with the Core Services Identity API. Send to your object to return an opaque object for use in the Core Services Identity API. Similarly, call to use an Core Services Identity opaque object in the Collaboration framework. There are two subclasses of : and . If you are working specifically with a group identity, use . Similarly, if you are working with a user identity, use .


// A object is used for accessing the attributes of an identity stored in an identity authority. You can use an identity object for finding identities, and storing them in an access control list (ACL). If you need to edit these attributes, take advantage of the class in Core Services.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(name:authority:)
func NewCBIdentityWithNameAuthority(name string /* primitive/slice/pointer. */, authority ICBIdentityAuthority) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(getCBIdentityClass().class), objc.Sel("identityWithName:authority:"), objc.String(name), authority)
	return rv
}



// Returns the identity object with the given name from the specified identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(name:authority:)
func (cc _CBIdentityClass) IdentityWithNameAuthority(name string /* primitive/slice/pointer. */, authority ICBIdentityAuthority) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(cc.class), objc.Sel("identityWithName:authority:"), objc.String(name), authority)
	return rv
}


// Returns the identity authority where the identity is stored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/authority
func (c_ CBIdentity) Authority() ICBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](c_.ID, objc.Sel("authority"))
	return rv
}


// Returns an array of aliases (alternate names) for the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/aliases
func (c_ CBIdentity) Aliases() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("aliases"))
	return rv
}


// Returns an array of aliases (alternate names) for the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/aliases
func (c_ CBIdentity) SetAliases(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAliases:"), objc.String(value))
}


// Returns the email address of an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/emailaddress
func (c_ CBIdentity) EmailAddress() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("emailAddress"))
	return rv
}


// Returns the email address of an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/emailaddress
func (c_ CBIdentity) SetEmailAddress(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailAddress:"), objc.String(value))
}


// Returns the full name of the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/fullname
func (c_ CBIdentity) FullName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("fullName"))
	return rv
}


// Returns the full name of the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/fullname
func (c_ CBIdentity) SetFullName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFullName:"), objc.String(value))
}


// Returns the image associated with an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/image
func (c_ CBIdentity) Image() appkit.objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](c_.ID, objc.Sel("image"))
	return rv
}


// Returns the image associated with an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/image
func (c_ CBIdentity) SetImage(value appkit.objc.IObject /* cross-framework: Image */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImage:"), value)
}


// Returns a Boolean value indicating the state of the identity’s hidden property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/ishidden
func (c_ CBIdentity) IsHidden() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHidden"))
	return rv
}


// Returns a Boolean value indicating the state of the identity’s hidden property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/ishidden
func (c_ CBIdentity) SetIsHidden(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHidden:"), value)
}


// Returns a persistent reference to store a reference to an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/persistentreference
func (c_ CBIdentity) PersistentReference() foundation.objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("persistentReference"))
	return rv
}


// Returns a persistent reference to store a reference to an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/persistentreference
func (c_ CBIdentity) SetPersistentReference(value foundation.objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPersistentReference:"), value)
}


// Returns the POSIX name of the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/posixname
func (c_ CBIdentity) PosixName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("posixName"))
	return rv
}


// Returns the POSIX name of the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/posixname
func (c_ CBIdentity) SetPosixName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPosixName:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/uniqueidentifier
func (c_ CBIdentity) UniqueIdentifier() foundation.objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/uniqueidentifier
func (c_ CBIdentity) SetUniqueIdentifier(value foundation.objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUniqueIdentifier:"), value)
}


// Returns the UUID of the identity as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/uuidstring
func (c_ CBIdentity) UuidString() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("uuidString"))
	return rv
}


// Returns the UUID of the identity as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/uuidstring
func (c_ CBIdentity) SetUuidString(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUuidString:"), objc.String(value))
}


