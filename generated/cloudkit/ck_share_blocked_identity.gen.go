// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/contacts"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKShareBlockedIdentity] class.
var (
	CKShareBlockedIdentityClass     _CKShareBlockedIdentityClass
	CKShareBlockedIdentityClassOnce sync.Once
)

func getCKShareBlockedIdentityClass() _CKShareBlockedIdentityClass {
	CKShareBlockedIdentityClassOnce.Do(func() {
		CKShareBlockedIdentityClass = _CKShareBlockedIdentityClass{objc.GetClass("CKShareBlockedIdentity")}
	})
	return CKShareBlockedIdentityClass
}

type _CKShareBlockedIdentityClass struct {
	class objc.Class
}

// An interface definition for the [CKShareBlockedIdentity] class.
type ICKShareBlockedIdentity interface {
	objectivec.IObject
	// properties:
	Contact() contacts.objc.IObject /* cross-framework: CNContact */
	UserIdentity() ICKUserIdentity
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/BlockedIdentity
type CKShareBlockedIdentity struct {
	objectivec.Object
}

// CKShareBlockedIdentityFrom constructs a [CKShareBlockedIdentity] from an unsafe.Pointer.
func CKShareBlockedIdentityFrom(ptr unsafe.Pointer) CKShareBlockedIdentity {
	return CKShareBlockedIdentity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKShareBlockedIdentityClass) Alloc() CKShareBlockedIdentity {
	rv := objc.Send[CKShareBlockedIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKShareBlockedIdentityClass) New() CKShareBlockedIdentity {
	rv := objc.Send[CKShareBlockedIdentity](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKShareBlockedIdentity) Init() CKShareBlockedIdentity {
	rv := objc.Send[CKShareBlockedIdentity](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKShareBlockedIdentity) Autorelease() CKShareBlockedIdentity {
	rv := objc.Send[CKShareBlockedIdentity](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareBlockedIdentity creates a new CKShareBlockedIdentity instance.
func NewCKShareBlockedIdentity() CKShareBlockedIdentity {
	return getCKShareBlockedIdentityClass().New()
}



// A displayable representing the blocked user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/BlockedIdentity/contact
func (c_ CKShareBlockedIdentity) Contact() contacts.objc.IObject /* cross-framework: CNContact */ {
	rv := objc.Send[contacts.CNContact](c_.ID, objc.Sel("contact"))
	return rv
}


// The identity of the user who has been blocked from requesting access to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/BlockedIdentity/userIdentity
func (c_ CKShareBlockedIdentity) UserIdentity() ICKUserIdentity {
	rv := objc.Send[CKUserIdentity](c_.ID, objc.Sel("userIdentity"))
	return rv
}



