// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CNContactFetchRequest] class.
type ICNContactFetchRequest interface {
	ICNFetchRequest
}

// An object that defines the options to use when fetching contacts.
//
// You need at least one contact property key to fetch a contact’s properties. Use this class with the method to execute the contact fetch request.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CNContactFetchRequestClass) Alloc() CNContactFetchRequest {
	rv := objc.Send[CNContactFetchRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a fetch request for the specified keys.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/init(keysToFetch:)
func NewCNContactFetchRequestWithKeysToFetch(keysToFetch unsafe.Pointer) CNContactFetchRequest {
	instance := getCNContactFetchRequestClass().Alloc()
	rv := objc.Send[CNContactFetchRequest](instance.ID, objc.Sel("initWithKeysToFetch:"), keysToFetch)
	rv.Autorelease()
	return rv
}


// The properties to fetch in the returned contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/keysToFetch
func (c_ CNContactFetchRequest) KeysToFetch() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("keysToFetch"))
	return rv
}


// SetKeysToFetch sets the value of the keysToFetch property.
// The properties to fetch in the returned contacts.

//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/keysToFetch
func (c_ CNContactFetchRequest) SetKeysToFetch(value []objc.ID) {
	// Convert Go slice to NSArray
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
}
// A Boolean value that indicates whether to return mutable contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/mutableObjects
func (c_ CNContactFetchRequest) MutableObjects() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("mutableObjects"))
	return rv
}


// SetMutableObjects sets the value of the mutableObjects property.
// A Boolean value that indicates whether to return mutable contacts.

//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/mutableObjects
func (c_ CNContactFetchRequest) SetMutableObjects(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMutableObjects:"), value)
}
// The predicate to match contacts against.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/predicate
func (c_ CNContactFetchRequest) Predicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("predicate"))
	return rv
}


// SetPredicate sets the value of the predicate property.
// The predicate to match contacts against.

//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/predicate
func (c_ CNContactFetchRequest) SetPredicate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredicate:"), value)
}
// The sort order for contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/sortOrder
func (c_ CNContactFetchRequest) SortOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sortOrder"))
	return rv
}


// SetSortOrder sets the value of the sortOrder property.
// The sort order for contacts.

//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/sortOrder
func (c_ CNContactFetchRequest) SetSortOrder(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSortOrder:"), value)
}
// A Boolean value that indicates whether to return linked contacts as unified contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/unifyResults
func (c_ CNContactFetchRequest) UnifyResults() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("unifyResults"))
	return rv
}


// SetUnifyResults sets the value of the unifyResults property.
// A Boolean value that indicates whether to return linked contacts as unified contacts.

//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactFetchRequest/unifyResults
func (c_ CNContactFetchRequest) SetUnifyResults(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUnifyResults:"), value)
}

