// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKShareBlockedIdentity */


/* debug [class_header]: Header for CKShareBlockedIdentity */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKShareBlockedIdentity */
// An interface definition for the [CKShareBlockedIdentity] class.
type ICKShareBlockedIdentity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKShareBlockedIdentity */
	// properties:
	Contact() contacts.CNContact
	UserIdentity() ICKUserIdentity
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKShareBlockedIdentity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKShareBlockedIdentity */
// Alloc allocates a new instance without initialization.
func (cc _CKShareBlockedIdentityClass) Alloc() CKShareBlockedIdentity {
	rv := objc.Send[CKShareBlockedIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKShareBlockedIdentity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/BlockedIdentity
type CKShareBlockedIdentity struct {
	objectivec.Object
}

// CKShareBlockedIdentityFrom constructs a [CKShareBlockedIdentity] from an unsafe.Pointer.
func CKShareBlockedIdentityFrom(ptr unsafe.Pointer) CKShareBlockedIdentity {
	return CKShareBlockedIdentity{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKShareBlockedIdentity *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKShareBlockedIdentity */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKShareBlockedIdentity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKShareBlockedIdentity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKShareBlockedIdentity */

// A displayable representing the blocked user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/BlockedIdentity/contact
func (c_ CKShareBlockedIdentity) Contact() contacts.CNContact {
	rv := objc.Send[contacts.CNContact](c_.ID, objc.Sel("contact"))
	return rv
}/* debug [instance_properties/getter]: contact */


// The identity of the user who has been blocked from requesting access to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/BlockedIdentity/userIdentity
func (c_ CKShareBlockedIdentity) UserIdentity() ICKUserIdentity {
	rv := objc.Send[CKUserIdentity](c_.ID, objc.Sel("userIdentity"))
	return rv
}/* debug [instance_properties/getter]: userIdentity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKShareBlockedIdentity */



