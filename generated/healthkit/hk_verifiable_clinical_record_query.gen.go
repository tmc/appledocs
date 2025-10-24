// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKVerifiableClinicalRecordQuery */


/* debug [class_header]: Header for HKVerifiableClinicalRecordQuery */
// The class instance for the [HKVerifiableClinicalRecordQuery] class.
var (
	HKVerifiableClinicalRecordQueryClass     _HKVerifiableClinicalRecordQueryClass
	HKVerifiableClinicalRecordQueryClassOnce sync.Once
)

func getHKVerifiableClinicalRecordQueryClass() _HKVerifiableClinicalRecordQueryClass {
	HKVerifiableClinicalRecordQueryClassOnce.Do(func() {
		HKVerifiableClinicalRecordQueryClass = _HKVerifiableClinicalRecordQueryClass{objc.GetClass("HKVerifiableClinicalRecordQuery")}
	})
	return HKVerifiableClinicalRecordQueryClass
}

type _HKVerifiableClinicalRecordQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKVerifiableClinicalRecordQuery */
// An interface definition for the [HKVerifiableClinicalRecordQuery] class.
type IHKVerifiableClinicalRecordQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKVerifiableClinicalRecordQuery */
	// properties:
	RecordTypes() []string
	SourceTypes() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKVerifiableClinicalRecordQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKVerifiableClinicalRecordQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKVerifiableClinicalRecordQueryClass) Alloc() HKVerifiableClinicalRecordQuery {
	rv := objc.Send[HKVerifiableClinicalRecordQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKVerifiableClinicalRecordQueryClass) New() HKVerifiableClinicalRecordQuery {
	rv := objc.Send[HKVerifiableClinicalRecordQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKVerifiableClinicalRecordQuery) Init() HKVerifiableClinicalRecordQuery {
	rv := objc.Send[HKVerifiableClinicalRecordQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKVerifiableClinicalRecordQuery) Autorelease() HKVerifiableClinicalRecordQuery {
	rv := objc.Send[HKVerifiableClinicalRecordQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKVerifiableClinicalRecordQuery creates a new HKVerifiableClinicalRecordQuery instance.
func NewHKVerifiableClinicalRecordQuery() HKVerifiableClinicalRecordQuery {
	return getHKVerifiableClinicalRecordQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKVerifiableClinicalRecordQuery */
// A query for one-time access to a SMART Health Card or EU Digital COVID Certificate.
//
// Use an object to request one-time access to a SMART Health Card or EU Digital COVID Certificate. For example, the following code requests cards that represent immunizations within the last six months. Unlike other HealthKit queries, you don’t need to request permission to read verifiable health records before running this query. HealthKit prompts the user for permission to read the records each time you run the query.


// A query for one-time access to a SMART Health Card or EU Digital COVID Certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecordQuery
type HKVerifiableClinicalRecordQuery struct {
	HKQuery
}

// HKVerifiableClinicalRecordQueryFrom constructs a [HKVerifiableClinicalRecordQuery] from an unsafe.Pointer.
//
// A query for one-time access to a SMART Health Card or EU Digital COVID Certificate.
func HKVerifiableClinicalRecordQueryFrom(ptr unsafe.Pointer) HKVerifiableClinicalRecordQuery {
	return HKVerifiableClinicalRecordQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKVerifiableClinicalRecordQuery */

// Creates a query for one-time access to a SMART Health Card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecordQuery/init(recordTypes:predicate:resultsHandler:)
func NewHKVerifiableClinicalRecordQueryWithRecordTypesPredicateResultsHandler(recordTypes []string, predicate foundation.Predicate, resultsHandler unsafe.Pointer) HKVerifiableClinicalRecordQuery {
	instance := getHKVerifiableClinicalRecordQueryClass().Alloc()
	rv := objc.Send[HKVerifiableClinicalRecordQuery](instance.ID, objc.Sel("initWithRecordTypes:predicate:resultsHandler:"), recordTypes, predicate, resultsHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKVerifiableClinicalRecordQueryWithRecordTypesPredicateResultsHandler */


// Creates a query for one-time access to a verifiable clinical record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecordQuery/init(recordTypes:sourceTypes:predicate:resultsHandler:)
func NewHKVerifiableClinicalRecordQueryWithRecordTypesSourceTypesPredicateResultsHandler(recordTypes []string, sourceTypes []string, predicate foundation.Predicate, resultsHandler unsafe.Pointer) HKVerifiableClinicalRecordQuery {
	instance := getHKVerifiableClinicalRecordQueryClass().Alloc()
	rv := objc.Send[HKVerifiableClinicalRecordQuery](instance.ID, objc.Sel("initWithRecordTypes:sourceTypes:predicate:resultsHandler:"), recordTypes, sourceTypes, predicate, resultsHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKVerifiableClinicalRecordQueryWithRecordTypesSourceTypesPredicateResultsHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKVerifiableClinicalRecordQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKVerifiableClinicalRecordQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKVerifiableClinicalRecordQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKVerifiableClinicalRecordQuery */

// The type of records that this query returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecordQuery/recordTypes
func (h_ HKVerifiableClinicalRecordQuery) RecordTypes() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("recordTypes"))
	return rv
}/* debug [instance_properties/getter]: recordTypes */


// The format of the verifiable clinical record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecordQuery/sourceTypes
func (h_ HKVerifiableClinicalRecordQuery) SourceTypes() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("sourceTypes"))
	return rv
}/* debug [instance_properties/getter]: sourceTypes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKVerifiableClinicalRecordQuery */


