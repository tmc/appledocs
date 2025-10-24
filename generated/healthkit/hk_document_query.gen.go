// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKDocumentQuery */


/* debug [class_header]: Header for HKDocumentQuery */
// The class instance for the [HKDocumentQuery] class.
var (
	HKDocumentQueryClass     _HKDocumentQueryClass
	HKDocumentQueryClassOnce sync.Once
)

func getHKDocumentQueryClass() _HKDocumentQueryClass {
	HKDocumentQueryClassOnce.Do(func() {
		HKDocumentQueryClass = _HKDocumentQueryClass{objc.GetClass("HKDocumentQuery")}
	})
	return HKDocumentQueryClass
}

type _HKDocumentQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKDocumentQuery */
// An interface definition for the [HKDocumentQuery] class.
type IHKDocumentQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKDocumentQuery */
	// properties:
	IncludeDocumentData() bool
	Limit() uint
	SortDescriptors() []cloudkit.SortDescriptor
	HKObjectQueryNoLimit() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKDocumentQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKDocumentQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKDocumentQueryClass) Alloc() HKDocumentQuery {
	rv := objc.Send[HKDocumentQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKDocumentQueryClass) New() HKDocumentQuery {
	rv := objc.Send[HKDocumentQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKDocumentQuery) Init() HKDocumentQuery {
	rv := objc.Send[HKDocumentQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKDocumentQuery) Autorelease() HKDocumentQuery {
	rv := objc.Send[HKDocumentQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKDocumentQuery creates a new HKDocumentQuery instance.
func NewHKDocumentQuery() HKDocumentQuery {
	return getHKDocumentQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKDocumentQuery */
// A query that returns a snapshot of all matching documents currently saved in the HealthKit store.
//
// Use an object to search for documents in the HealthKit store. You can provide a predicate to filter the search results, a sort order for the returned samples, or even a limit to the number of samples returned. Document queries are immutable: The query’s properties are set when the query is first created. They cannot change.


// A query that returns a snapshot of all matching documents currently saved in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDocumentQuery
type HKDocumentQuery struct {
	HKQuery
}

// HKDocumentQueryFrom constructs a [HKDocumentQuery] from an unsafe.Pointer.
//
// A query that returns a snapshot of all matching documents currently saved in the HealthKit store.
func HKDocumentQueryFrom(ptr unsafe.Pointer) HKDocumentQuery {
	return HKDocumentQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKDocumentQuery */

// Instantiates and returns a document query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDocumentQuery/init(documentType:predicate:limit:sortDescriptors:includeDocumentData:resultsHandler:)
func NewHKDocumentQueryWithDocumentTypePredicateLimitSortDescriptorsIncludeDocumentDataResultsHandler(documentType IHKDocumentType, predicate foundation.Predicate, limit uint, sortDescriptors []cloudkit.SortDescriptor, includeDocumentData bool, resultsHandler unsafe.Pointer) HKDocumentQuery {
	instance := getHKDocumentQueryClass().Alloc()
	rv := objc.Send[HKDocumentQuery](instance.ID, objc.Sel("initWithDocumentType:predicate:limit:sortDescriptors:includeDocumentData:resultsHandler:"), documentType, predicate, limit, sortDescriptors, includeDocumentData, resultsHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKDocumentQueryWithDocumentTypePredicateLimitSortDescriptorsIncludeDocumentDataResultsHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKDocumentQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKDocumentQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKDocumentQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKDocumentQuery */

// A Boolean value that indicates whether the sample includes the full document’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDocumentQuery/includeDocumentData
func (h_ HKDocumentQuery) IncludeDocumentData() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("includeDocumentData"))
	return rv
}/* debug [instance_properties/getter]: includeDocumentData */


// The maximum number of documents the receiver will return upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDocumentQuery/limit
func (h_ HKDocumentQuery) Limit() uint {
	rv := objc.Send[uint](h_.ID, objc.Sel("limit"))
	return rv
}/* debug [instance_properties/getter]: limit */


// An array of sort descriptors that specify the order of the results returned by this query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDocumentQuery/sortDescriptors
func (h_ HKDocumentQuery) SortDescriptors() []cloudkit.SortDescriptor {
	rv := objc.Send[[]cloudkit.SortDescriptor](h_.ID, objc.Sel("sortDescriptors"))
	return rv
}/* debug [instance_properties/getter]: sortDescriptors */


// A value indicating that the query returns all the matching samples in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobjectquerynolimit
func (h_ HKDocumentQuery) HKObjectQueryNoLimit() int {
	rv := objc.Send[int](h_.ID, objc.Sel("HKObjectQueryNoLimit"))
	return rv
}/* debug [instance_properties/getter]: HKObjectQueryNoLimit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKDocumentQuery */


