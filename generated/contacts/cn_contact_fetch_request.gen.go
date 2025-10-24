// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNContactFetchRequest */


/* debug [class_header]: Header for CNContactFetchRequest */
// The class instance for the [CNContactFetchRequest] class.
var (
	CNContactFetchRequestClass     _CNContactFetchRequestClass
	CNContactFetchRequestClassOnce sync.Once
)

func getCNContactFetchRequestClass() _CNContactFetchRequestClass {
	CNContactFetchRequestClassOnce.Do(func() {
		CNContactFetchRequestClass = _CNContactFetchRequestClass{objc.GetClass("CNContactFetchRequest")}
	})
	return CNContactFetchRequestClass
}

type _CNContactFetchRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNContactFetchRequest */
// An interface definition for the [CNContactFetchRequest] class.
type ICNContactFetchRequest interface {
	ICNFetchRequest
	
/* debug [class_interface_properties]: Properties for CNContactFetchRequest */
	// properties:
	KeysToFetch() []objc.ID
	SetKeysToFetch(value []objc.ID)
	MutableObjects() bool
	SetMutableObjects(value bool)
	Predicate() foundation.Predicate
	SetPredicate(value foundation.Predicate)
	SortOrder() CNContactSortOrder
	SetSortOrder(value CNContactSortOrder)
	UnifyResults() bool
	SetUnifyResults(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNContactFetchRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNContactFetchRequest */
// Alloc allocates a new instance without initialization.
func (cc _CNContactFetchRequestClass) Alloc() CNContactFetchRequest {
	rv := objc.Send[CNContactFetchRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNContactFetchRequestClass) New() CNContactFetchRequest {
	rv := objc.Send[CNContactFetchRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactFetchRequest) Init() CNContactFetchRequest {
	rv := objc.Send[CNContactFetchRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactFetchRequest) Autorelease() CNContactFetchRequest {
	rv := objc.Send[CNContactFetchRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactFetchRequest creates a new CNContactFetchRequest instance.
func NewCNContactFetchRequest() CNContactFetchRequest {
	return getCNContactFetchRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNContactFetchRequest */
// An object that defines the options to use when fetching contacts.
//
// You need at least one contact property key to fetch a contact’s properties. Use this class with the method to execute the contact fetch request.


// An object that defines the options to use when fetching contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest
type CNContactFetchRequest struct {
	CNFetchRequest
}

// CNContactFetchRequestFrom constructs a [CNContactFetchRequest] from an unsafe.Pointer.
//
// An object that defines the options to use when fetching contacts.
func CNContactFetchRequestFrom(ptr unsafe.Pointer) CNContactFetchRequest {
	return CNContactFetchRequest{
		CNFetchRequest: CNFetchRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNContactFetchRequest */

// Creates a fetch request for the specified keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/init(keysToFetch:)
func NewCNContactFetchRequestWithKeysToFetch(keysToFetch []objc.ID) CNContactFetchRequest {
	instance := getCNContactFetchRequestClass().Alloc()
	rv := objc.Send[CNContactFetchRequest](instance.ID, objc.Sel("initWithKeysToFetch:"), keysToFetch)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNContactFetchRequestWithKeysToFetch */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNContactFetchRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNContactFetchRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNContactFetchRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNContactFetchRequest */

// The properties to fetch in the returned contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/keysToFetch
func (c_ CNContactFetchRequest) KeysToFetch() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("keysToFetch"))
	return rv
}/* debug [instance_properties/getter]: keysToFetch */


// The properties to fetch in the returned contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/keysToFetch
func (c_ CNContactFetchRequest) SetKeysToFetch(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeysToFetch:"), nsArray)
}/* debug [instance_properties/setter]: keysToFetch */


// A Boolean value that indicates whether to return mutable contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/mutableObjects
func (c_ CNContactFetchRequest) MutableObjects() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("mutableObjects"))
	return rv
}/* debug [instance_properties/getter]: mutableObjects */


// A Boolean value that indicates whether to return mutable contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/mutableObjects
func (c_ CNContactFetchRequest) SetMutableObjects(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMutableObjects:"), value)
}/* debug [instance_properties/setter]: mutableObjects */


// The predicate to match contacts against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/predicate
func (c_ CNContactFetchRequest) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](c_.ID, objc.Sel("predicate"))
	return rv
}/* debug [instance_properties/getter]: predicate */


// The predicate to match contacts against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/predicate
func (c_ CNContactFetchRequest) SetPredicate(value foundation.Predicate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredicate:"), value)
}/* debug [instance_properties/setter]: predicate */


// The sort order for contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/sortOrder
func (c_ CNContactFetchRequest) SortOrder() CNContactSortOrder {
	rv := objc.Send[CNContactSortOrder](c_.ID, objc.Sel("sortOrder"))
	return rv
}/* debug [instance_properties/getter]: sortOrder */


// The sort order for contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/sortOrder
func (c_ CNContactFetchRequest) SetSortOrder(value CNContactSortOrder) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSortOrder:"), value)
}/* debug [instance_properties/setter]: sortOrder */


// A Boolean value that indicates whether to return linked contacts as unified contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/unifyResults
func (c_ CNContactFetchRequest) UnifyResults() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("unifyResults"))
	return rv
}/* debug [instance_properties/getter]: unifyResults */


// A Boolean value that indicates whether to return linked contacts as unified contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/unifyResults
func (c_ CNContactFetchRequest) SetUnifyResults(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUnifyResults:"), value)
}/* debug [instance_properties/setter]: unifyResults */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNContactFetchRequest */


