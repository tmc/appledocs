// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CNChangeHistoryFetchRequest] class.
var (
	CNChangeHistoryFetchRequestClass     _CNChangeHistoryFetchRequestClass
	CNChangeHistoryFetchRequestClassOnce sync.Once
)

func getCNChangeHistoryFetchRequestClass() _CNChangeHistoryFetchRequestClass {
	CNChangeHistoryFetchRequestClassOnce.Do(func() {
		CNChangeHistoryFetchRequestClass = _CNChangeHistoryFetchRequestClass{objc.GetClass("CNChangeHistoryFetchRequest")}
	})
	return CNChangeHistoryFetchRequestClass
}

type _CNChangeHistoryFetchRequestClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryFetchRequest] class.
type ICNChangeHistoryFetchRequest interface {
	ICNFetchRequest
	AdditionalContactKeyDescriptors() []objc.ID
	SetAdditionalContactKeyDescriptors(value []objc.ID)
	ExcludedTransactionAuthors() []string
	SetExcludedTransactionAuthors(value []string)
	IncludeGroupChanges() bool
	SetIncludeGroupChanges(value bool)
	MutableObjects() bool
	SetMutableObjects(value bool)
	ShouldUnifyResults() bool
	SetShouldUnifyResults(value bool)
	StartingToken() foundation.NSData
	SetStartingToken(value foundation.IData)
}

// An object that specifies the criteria for fetching change history.
//
// The system always returns changes to contacts. The system coalesces changes to remove redundant adds, updates, and deletes. Create and configure a fetch request, then call to process changes.


// An object that specifies the criteria for fetching change history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest
type CNChangeHistoryFetchRequest struct {
	CNFetchRequest
}

// CNChangeHistoryFetchRequestFrom constructs a [CNChangeHistoryFetchRequest] from an unsafe.Pointer.
//
// An object that specifies the criteria for fetching change history.
func CNChangeHistoryFetchRequestFrom(ptr unsafe.Pointer) CNChangeHistoryFetchRequest {
	return CNChangeHistoryFetchRequest{
		CNFetchRequest: CNFetchRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryFetchRequestClass) Alloc() CNChangeHistoryFetchRequest {
	rv := objc.Send[CNChangeHistoryFetchRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryFetchRequestClass) New() CNChangeHistoryFetchRequest {
	rv := objc.Send[CNChangeHistoryFetchRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryFetchRequest) Init() CNChangeHistoryFetchRequest {
	rv := objc.Send[CNChangeHistoryFetchRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryFetchRequest) Autorelease() CNChangeHistoryFetchRequest {
	rv := objc.Send[CNChangeHistoryFetchRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryFetchRequest creates a new CNChangeHistoryFetchRequest instance.
func NewCNChangeHistoryFetchRequest() CNChangeHistoryFetchRequest {
	return getCNChangeHistoryFetchRequestClass().New()
}



// An array of contact property keys or key descriptors from contact objects to fetch in the returned contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/additionalContactKeyDescriptors
func (c_ CNChangeHistoryFetchRequest) AdditionalContactKeyDescriptors() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("additionalContactKeyDescriptors"))
	return rv
}


// An array of contact property keys or key descriptors from contact objects to fetch in the returned contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/additionalContactKeyDescriptors
func (c_ CNChangeHistoryFetchRequest) SetAdditionalContactKeyDescriptors(value []objc.ID) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setAdditionalContactKeyDescriptors:"), nsArray)
}


// An array of strings that identify transaction authors to exclude from the fetch results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/excludedTransactionAuthors
func (c_ CNChangeHistoryFetchRequest) ExcludedTransactionAuthors() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("excludedTransactionAuthors"))
	return rv
}


// An array of strings that identify transaction authors to exclude from the fetch results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/excludedTransactionAuthors
func (c_ CNChangeHistoryFetchRequest) SetExcludedTransactionAuthors(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setExcludedTransactionAuthors:"), nsArray)
}


// A Boolean value that indicates whether the fetch should also return group changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/includeGroupChanges
func (c_ CNChangeHistoryFetchRequest) IncludeGroupChanges() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("includeGroupChanges"))
	return rv
}


// A Boolean value that indicates whether the fetch should also return group changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/includeGroupChanges
func (c_ CNChangeHistoryFetchRequest) SetIncludeGroupChanges(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIncludeGroupChanges:"), value)
}


// A Boolean value that indicates whether the fetch should return mutable contacts and groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/mutableObjects
func (c_ CNChangeHistoryFetchRequest) MutableObjects() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("mutableObjects"))
	return rv
}


// A Boolean value that indicates whether the fetch should return mutable contacts and groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/mutableObjects
func (c_ CNChangeHistoryFetchRequest) SetMutableObjects(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMutableObjects:"), value)
}


// A Boolean value that indicates whether the fetch should return contact changes as unified contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/shouldUnifyResults
func (c_ CNChangeHistoryFetchRequest) ShouldUnifyResults() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldUnifyResults"))
	return rv
}


// A Boolean value that indicates whether the fetch should return contact changes as unified contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/shouldUnifyResults
func (c_ CNChangeHistoryFetchRequest) SetShouldUnifyResults(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldUnifyResults:"), value)
}


// An opaque token that indicates a point in history in the user’s Contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/startingToken
func (c_ CNChangeHistoryFetchRequest) StartingToken() foundation.NSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("startingToken"))
	return rv
}


// An opaque token that indicates a point in history in the user’s Contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/startingToken
func (c_ CNChangeHistoryFetchRequest) SetStartingToken(value foundation.IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStartingToken:"), value)
}



