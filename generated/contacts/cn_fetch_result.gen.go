// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNFetchResult */


/* debug [class_header]: Header for CNFetchResult */
// The class instance for the [CNFetchResult] class.
var (
	CNFetchResultClass     _CNFetchResultClass
	CNFetchResultClassOnce sync.Once
)

func getCNFetchResultClass() _CNFetchResultClass {
	CNFetchResultClassOnce.Do(func() {
		CNFetchResultClass = _CNFetchResultClass{objc.GetClass("CNFetchResult")}
	})
	return CNFetchResultClass
}

type _CNFetchResultClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNFetchResult */
// An interface definition for the [CNFetchResult] class.
type ICNFetchResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNFetchResult */
	// properties:
	CurrentHistoryToken() objc.IObject /* cross-framework: NSData */
	Value() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNFetchResult */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNFetchResult */
// Alloc allocates a new instance without initialization.
func (cc _CNFetchResultClass) Alloc() CNFetchResult {
	rv := objc.Send[CNFetchResult](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNFetchResultClass) New() CNFetchResult {
	rv := objc.Send[CNFetchResult](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNFetchResult) Init() CNFetchResult {
	rv := objc.Send[CNFetchResult](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNFetchResult) Autorelease() CNFetchResult {
	rv := objc.Send[CNFetchResult](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNFetchResult creates a new CNFetchResult instance.
func NewCNFetchResult() CNFetchResult {
	return getCNFetchResultClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNFetchResult */
// An object that represents the result of a change-history fetch request.


// An object that represents the result of a change-history fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNFetchResult
type CNFetchResult struct {
	objectivec.Object
}

// CNFetchResultFrom constructs a [CNFetchResult] from an unsafe.Pointer.
//
// An object that represents the result of a change-history fetch request.
func CNFetchResultFrom(ptr unsafe.Pointer) CNFetchResult {
	return CNFetchResult{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNFetchResult *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNFetchResult */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNFetchResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNFetchResult */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNFetchResult */

// An opaque token that indicates a point in history in the user’s Contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNFetchResult/currentHistoryToken
func (c_ CNFetchResult) CurrentHistoryToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("currentHistoryToken"))
	return rv
}/* debug [instance_properties/getter]: currentHistoryToken */


// The result of the fetch request, expressed as the value type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNFetchResult/value
func (c_ CNFetchResult) Value() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNFetchResult */



