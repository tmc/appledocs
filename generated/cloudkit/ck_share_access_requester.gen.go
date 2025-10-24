// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/contacts"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKShareAccessRequester */


/* debug [class_header]: Header for CKShareAccessRequester */
// The class instance for the [CKShareAccessRequester] class.
var (
	CKShareAccessRequesterClass     _CKShareAccessRequesterClass
	CKShareAccessRequesterClassOnce sync.Once
)

func getCKShareAccessRequesterClass() _CKShareAccessRequesterClass {
	CKShareAccessRequesterClassOnce.Do(func() {
		CKShareAccessRequesterClass = _CKShareAccessRequesterClass{objc.GetClass("CKShareAccessRequester")}
	})
	return CKShareAccessRequesterClass
}

type _CKShareAccessRequesterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKShareAccessRequester */
// An interface definition for the [CKShareAccessRequester] class.
type ICKShareAccessRequester interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKShareAccessRequester */
	// properties:
	Contact() contacts.CNContact
	ParticipantLookupInfo() ICKUserIdentityLookupInfo
	UserIdentity() ICKUserIdentity
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKShareAccessRequester */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKShareAccessRequester */
// Alloc allocates a new instance without initialization.
func (cc _CKShareAccessRequesterClass) Alloc() CKShareAccessRequester {
	rv := objc.Send[CKShareAccessRequester](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKShareAccessRequesterClass) New() CKShareAccessRequester {
	rv := objc.Send[CKShareAccessRequester](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKShareAccessRequester) Init() CKShareAccessRequester {
	rv := objc.Send[CKShareAccessRequester](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKShareAccessRequester) Autorelease() CKShareAccessRequester {
	rv := objc.Send[CKShareAccessRequester](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareAccessRequester creates a new CKShareAccessRequester instance.
func NewCKShareAccessRequester() CKShareAccessRequester {
	return getCKShareAccessRequesterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKShareAccessRequester */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester
type CKShareAccessRequester struct {
	objectivec.Object
}

// CKShareAccessRequesterFrom constructs a [CKShareAccessRequester] from an unsafe.Pointer.
func CKShareAccessRequesterFrom(ptr unsafe.Pointer) CKShareAccessRequester {
	return CKShareAccessRequester{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKShareAccessRequester *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKShareAccessRequester */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKShareAccessRequester */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKShareAccessRequester */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKShareAccessRequester */

// A displayable representing the requester.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester/contact
func (c_ CKShareAccessRequester) Contact() contacts.CNContact {
	rv := objc.Send[contacts.CNContact](c_.ID, objc.Sel("contact"))
	return rv
}/* debug [instance_properties/getter]: contact */


// Lookup information for the requester.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester/participantLookupInfo
func (c_ CKShareAccessRequester) ParticipantLookupInfo() ICKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](c_.ID, objc.Sel("participantLookupInfo"))
	return rv
}/* debug [instance_properties/getter]: participantLookupInfo */


// The identity of the user requesting access to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester/userIdentity
func (c_ CKShareAccessRequester) UserIdentity() ICKUserIdentity {
	rv := objc.Send[CKUserIdentity](c_.ID, objc.Sel("userIdentity"))
	return rv
}/* debug [instance_properties/getter]: userIdentity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKShareAccessRequester */



