// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

package collaboration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CBIdentityAuthority] class.
var (
	CBIdentityAuthorityClass     _CBIdentityAuthorityClass
	CBIdentityAuthorityClassOnce sync.Once
)

func getCBIdentityAuthorityClass() _CBIdentityAuthorityClass {
	CBIdentityAuthorityClassOnce.Do(func() {
		CBIdentityAuthorityClass = _CBIdentityAuthorityClass{objc.GetClass("CBIdentityAuthority")}
	})
	return CBIdentityAuthorityClass
}

type _CBIdentityAuthorityClass struct {
	class objc.Class
}

// An interface definition for the [CBIdentityAuthority] class.
type ICBIdentityAuthority interface {
	objectivec.IObject
	// properties:
	LocalizedName() objc.IObject /* cross-framework: NSString */
	SetLocalizedName(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// An identity authority is a database that stores information about identities. The class defines one or more identity authorities. You can search this database for identities in conjunction with the class factory methods.


// An identity authority is a database that stores information about identities. The class defines one or more identity authorities. You can search this database for identities in conjunction with the class factory methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityAuthority
type CBIdentityAuthority struct {
	objectivec.Object
}

// CBIdentityAuthorityFrom constructs a [CBIdentityAuthority] from an unsafe.Pointer.
//
// An identity authority is a database that stores information about identities. The class defines one or more identity authorities. You can search this database for identities in conjunction with the class factory methods.
func CBIdentityAuthorityFrom(ptr unsafe.Pointer) CBIdentityAuthority {
	return CBIdentityAuthority{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CBIdentityAuthorityClass) Alloc() CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBIdentityAuthorityClass) New() CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBIdentityAuthority) Init() CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBIdentityAuthority) Autorelease() CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBIdentityAuthority creates a new CBIdentityAuthority instance.
func NewCBIdentityAuthority() CBIdentityAuthority {
	return getCBIdentityAuthorityClass().New()
}



// Returns an identity authority specified by a given Core Services Identity authority object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityAuthority/identityAuthorityWithCSIdentityAuthority:
func (cc _CBIdentityAuthorityClass) IdentityAuthorityWithCSIdentityAuthority(CSIdentityAuthority unsafe.Pointer) CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](objc.ID(cc.class), objc.Sel("identityAuthorityWithCSIdentityAuthority:"), CSIdentityAuthority)
	return rv
}


// Returns the localized name of the identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentityauthority/localizedname
func (c_ CBIdentityAuthority) LocalizedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedName"))
	return rv
}


// Returns the localized name of the identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbidentityauthority/localizedname
func (c_ CBIdentityAuthority) SetLocalizedName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedName:"), value)
}



