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

/* debug [class.gen.go]: Generating class CBIdentity */


/* debug [class_header]: Header for CBIdentity */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBIdentity */
// An interface definition for the [CBIdentity] class.
type ICBIdentity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CBIdentity */
	// properties:
	Aliases() []string
	Authority() ICBIdentityAuthority
	CSIdentity() unsafe.Pointer
	EmailAddress() objc.IObject /* cross-framework: NSString */
	FullName() objc.IObject /* cross-framework: NSString */
	Image() appkit.Image
	Hidden() bool
	PersistentReference() objc.IObject /* cross-framework: NSData */
	PosixName() objc.IObject /* cross-framework: NSString */
	UniqueIdentifier() foundation.UUID
	UUIDString() objc.IObject /* cross-framework: NSString */
	IsHidden() bool
	SetIsHidden(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBIdentity */
	// methods:
	IsMemberOfGroup(group ICBGroupIdentity) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBIdentity */
// Alloc allocates a new instance without initialization.
func (cc _CBIdentityClass) Alloc() CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBIdentity */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBIdentity */

// Returns the identity object with the given name from the specified identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(name:authority:)
func NewCBIdentityWithNameAuthority(name objc.IObject /* cross-framework: NSString */, authority ICBIdentityAuthority) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(getCBIdentityClass().class), objc.Sel("identityWithName:authority:"), name, authority)
	return rv
}/* debug [class_init_methods/constructor]: NewCBIdentityWithNameAuthority */


// Returns the identity object matching the persistent reference data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(persistentReference:)
func NewCBIdentityWithPersistentReference(data objc.IObject /* cross-framework: NSData */) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(getCBIdentityClass().class), objc.Sel("identityWithPersistentReference:"), data)
	return rv
}/* debug [class_init_methods/constructor]: NewCBIdentityWithPersistentReference */


// Returns the identity object with the given UUID from the specified identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(uuidString:authority:)
func NewCBIdentityWithUUIDStringAuthority(uuid objc.IObject /* cross-framework: NSString */, authority ICBIdentityAuthority) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(getCBIdentityClass().class), objc.Sel("identityWithUUIDString:authority:"), uuid, authority)
	return rv
}/* debug [class_init_methods/constructor]: NewCBIdentityWithUUIDStringAuthority */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(uniqueIdentifier:authority:)
func NewCBIdentityWithUniqueIdentifierAuthority(uuid foundation.UUID, authority ICBIdentityAuthority) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(getCBIdentityClass().class), objc.Sel("identityWithUniqueIdentifier:authority:"), uuid, authority)
	return rv
}/* debug [class_init_methods/constructor]: NewCBIdentityWithUniqueIdentifierAuthority */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBIdentity */

// Returns an identity object created from the specified Core Services Identity opaque object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/identityWithCSIdentity:
func (cc _CBIdentityClass) IdentityWithCSIdentity(csIdentity unsafe.Pointer) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(cc.class), objc.Sel("identityWithCSIdentity:"), csIdentity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IdentityWithCSIdentity) */


// Returns the identity object with the given name from the specified identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(name:authority:)
func (cc _CBIdentityClass) IdentityWithNameAuthority(name objc.IObject /* cross-framework: NSString */, authority ICBIdentityAuthority) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(cc.class), objc.Sel("identityWithName:authority:"), name, authority)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IdentityWithNameAuthority) */


// Returns the identity object matching the persistent reference data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(persistentReference:)
func (cc _CBIdentityClass) IdentityWithPersistentReference(data objc.IObject /* cross-framework: NSData */) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(cc.class), objc.Sel("identityWithPersistentReference:"), data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IdentityWithPersistentReference) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(uniqueIdentifier:authority:)
func (cc _CBIdentityClass) IdentityWithUniqueIdentifierAuthority(uuid foundation.UUID, authority ICBIdentityAuthority) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(cc.class), objc.Sel("identityWithUniqueIdentifier:authority:"), uuid, authority)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IdentityWithUniqueIdentifierAuthority) */


// Returns the identity object with the given UUID from the specified identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/init(uuidString:authority:)
func (cc _CBIdentityClass) IdentityWithUUIDStringAuthority(uuid objc.IObject /* cross-framework: NSString */, authority ICBIdentityAuthority) CBIdentity {
	rv := objc.Send[CBIdentity](objc.ID(cc.class), objc.Sel("identityWithUUIDString:authority:"), uuid, authority)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IdentityWithUUIDStringAuthority) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBIdentity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBIdentity */

// Returns a Boolean value indicating whether the identity is a member of the specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/isMember(ofGroup:)
func (c_ CBIdentity) IsMemberOfGroup(group ICBGroupIdentity) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMemberOfGroup:"), group)
	return rv
}/* debug [instance_methods/method]: IsMemberOfGroup */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBIdentity */

// Returns an array of aliases (alternate names) for the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/aliases
func (c_ CBIdentity) Aliases() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("aliases"))
	return rv
}/* debug [instance_properties/getter]: aliases */


// Returns the identity authority where the identity is stored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/authority
func (c_ CBIdentity) Authority() ICBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](c_.ID, objc.Sel("authority"))
	return rv
}/* debug [instance_properties/getter]: authority */


// Returns an opaque object for use with the Core Services Identity API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/CSIdentity
func (c_ CBIdentity) CSIdentity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("CSIdentity"))
	return rv
}/* debug [instance_properties/getter]: CSIdentity */


// Returns the email address of an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/emailAddress
func (c_ CBIdentity) EmailAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("emailAddress"))
	return rv
}/* debug [instance_properties/getter]: emailAddress */


// Returns the full name of the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/fullName
func (c_ CBIdentity) FullName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("fullName"))
	return rv
}/* debug [instance_properties/getter]: fullName */


// Returns the image associated with an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/image
func (c_ CBIdentity) Image() appkit.Image {
	rv := objc.Send[appkit.Image](c_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// Returns a Boolean value indicating the state of the identity’s hidden property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/isHidden
func (c_ CBIdentity) Hidden() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hidden"))
	return rv
}/* debug [instance_properties/getter]: hidden */


// Returns a persistent reference to store a reference to an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/persistentReference
func (c_ CBIdentity) PersistentReference() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("persistentReference"))
	return rv
}/* debug [instance_properties/getter]: persistentReference */


// Returns the POSIX name of the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/posixName
func (c_ CBIdentity) PosixName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("posixName"))
	return rv
}/* debug [instance_properties/getter]: posixName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/uniqueIdentifier
func (c_ CBIdentity) UniqueIdentifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}/* debug [instance_properties/getter]: uniqueIdentifier */


// Returns the UUID of the identity as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentity/uuidString
func (c_ CBIdentity) UUIDString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("UUIDString"))
	return rv
}/* debug [instance_properties/getter]: UUIDString */


// Returns a Boolean value indicating the state of the identity’s hidden property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/ishidden
func (c_ CBIdentity) IsHidden() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// Returns a Boolean value indicating the state of the identity’s hidden property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentity/ishidden
func (c_ CBIdentity) SetIsHidden(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBIdentity */


