// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryFetchRequest */


/* debug [class_header]: Header for CNChangeHistoryFetchRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryFetchRequest */
// An interface definition for the [CNChangeHistoryFetchRequest] class.
type ICNChangeHistoryFetchRequest interface {
	ICNFetchRequest
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryFetchRequest */
	// properties:
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
	StartingToken() objc.IObject /* cross-framework: NSData */
	SetStartingToken(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryFetchRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryFetchRequest */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryFetchRequestClass) Alloc() CNChangeHistoryFetchRequest {
	rv := objc.Send[CNChangeHistoryFetchRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryFetchRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryFetchRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryFetchRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryFetchRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryFetchRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryFetchRequest */

// An array of contact property keys or key descriptors from contact objects to fetch in the returned contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/additionalContactKeyDescriptors
func (c_ CNChangeHistoryFetchRequest) AdditionalContactKeyDescriptors() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("additionalContactKeyDescriptors"))
	return rv
}/* debug [instance_properties/getter]: additionalContactKeyDescriptors */


// An array of contact property keys or key descriptors from contact objects to fetch in the returned contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/additionalContactKeyDescriptors
func (c_ CNChangeHistoryFetchRequest) SetAdditionalContactKeyDescriptors(value []objc.ID) {
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
}/* debug [instance_properties/setter]: additionalContactKeyDescriptors */


// An array of strings that identify transaction authors to exclude from the fetch results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/excludedTransactionAuthors
func (c_ CNChangeHistoryFetchRequest) ExcludedTransactionAuthors() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("excludedTransactionAuthors"))
	return rv
}/* debug [instance_properties/getter]: excludedTransactionAuthors */


// An array of strings that identify transaction authors to exclude from the fetch results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/excludedTransactionAuthors
func (c_ CNChangeHistoryFetchRequest) SetExcludedTransactionAuthors(value []string) {
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
}/* debug [instance_properties/setter]: excludedTransactionAuthors */


// A Boolean value that indicates whether the fetch should also return group changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/includeGroupChanges
func (c_ CNChangeHistoryFetchRequest) IncludeGroupChanges() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("includeGroupChanges"))
	return rv
}/* debug [instance_properties/getter]: includeGroupChanges */


// A Boolean value that indicates whether the fetch should also return group changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/includeGroupChanges
func (c_ CNChangeHistoryFetchRequest) SetIncludeGroupChanges(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIncludeGroupChanges:"), value)
}/* debug [instance_properties/setter]: includeGroupChanges */


// A Boolean value that indicates whether the fetch should return mutable contacts and groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/mutableObjects
func (c_ CNChangeHistoryFetchRequest) MutableObjects() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("mutableObjects"))
	return rv
}/* debug [instance_properties/getter]: mutableObjects */


// A Boolean value that indicates whether the fetch should return mutable contacts and groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/mutableObjects
func (c_ CNChangeHistoryFetchRequest) SetMutableObjects(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMutableObjects:"), value)
}/* debug [instance_properties/setter]: mutableObjects */


// A Boolean value that indicates whether the fetch should return contact changes as unified contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/shouldUnifyResults
func (c_ CNChangeHistoryFetchRequest) ShouldUnifyResults() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldUnifyResults"))
	return rv
}/* debug [instance_properties/getter]: shouldUnifyResults */


// A Boolean value that indicates whether the fetch should return contact changes as unified contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/shouldUnifyResults
func (c_ CNChangeHistoryFetchRequest) SetShouldUnifyResults(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldUnifyResults:"), value)
}/* debug [instance_properties/setter]: shouldUnifyResults */


// An opaque token that indicates a point in history in the user’s Contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/startingToken
func (c_ CNChangeHistoryFetchRequest) StartingToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("startingToken"))
	return rv
}/* debug [instance_properties/getter]: startingToken */


// An opaque token that indicates a point in history in the user’s Contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryFetchRequest/startingToken
func (c_ CNChangeHistoryFetchRequest) SetStartingToken(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStartingToken:"), value)
}/* debug [instance_properties/setter]: startingToken */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryFetchRequest */



