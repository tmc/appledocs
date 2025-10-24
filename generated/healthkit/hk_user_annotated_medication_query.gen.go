// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKUserAnnotatedMedicationQuery */


/* debug [class_header]: Header for HKUserAnnotatedMedicationQuery */
// The class instance for the [HKUserAnnotatedMedicationQuery] class.
var (
	HKUserAnnotatedMedicationQueryClass     _HKUserAnnotatedMedicationQueryClass
	HKUserAnnotatedMedicationQueryClassOnce sync.Once
)

func getHKUserAnnotatedMedicationQueryClass() _HKUserAnnotatedMedicationQueryClass {
	HKUserAnnotatedMedicationQueryClassOnce.Do(func() {
		HKUserAnnotatedMedicationQueryClass = _HKUserAnnotatedMedicationQueryClass{objc.GetClass("HKUserAnnotatedMedicationQuery")}
	})
	return HKUserAnnotatedMedicationQueryClass
}

type _HKUserAnnotatedMedicationQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKUserAnnotatedMedicationQuery */
// An interface definition for the [HKUserAnnotatedMedicationQuery] class.
type IHKUserAnnotatedMedicationQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKUserAnnotatedMedicationQuery */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKUserAnnotatedMedicationQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKUserAnnotatedMedicationQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKUserAnnotatedMedicationQueryClass) Alloc() HKUserAnnotatedMedicationQuery {
	rv := objc.Send[HKUserAnnotatedMedicationQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKUserAnnotatedMedicationQueryClass) New() HKUserAnnotatedMedicationQuery {
	rv := objc.Send[HKUserAnnotatedMedicationQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKUserAnnotatedMedicationQuery) Init() HKUserAnnotatedMedicationQuery {
	rv := objc.Send[HKUserAnnotatedMedicationQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKUserAnnotatedMedicationQuery) Autorelease() HKUserAnnotatedMedicationQuery {
	rv := objc.Send[HKUserAnnotatedMedicationQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKUserAnnotatedMedicationQuery creates a new HKUserAnnotatedMedicationQuery instance.
func NewHKUserAnnotatedMedicationQuery() HKUserAnnotatedMedicationQuery {
	return getHKUserAnnotatedMedicationQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKUserAnnotatedMedicationQuery */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedicationQuery
type HKUserAnnotatedMedicationQuery struct {
	HKQuery
}

// HKUserAnnotatedMedicationQueryFrom constructs a [HKUserAnnotatedMedicationQuery] from an unsafe.Pointer.
func HKUserAnnotatedMedicationQueryFrom(ptr unsafe.Pointer) HKUserAnnotatedMedicationQuery {
	return HKUserAnnotatedMedicationQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKUserAnnotatedMedicationQuery */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedicationQuery/init(predicate:limit:resultsHandler:)
func NewHKUserAnnotatedMedicationQueryWithPredicateLimitResultsHandler(predicate foundation.Predicate, limit uint, resultsHandler unsafe.Pointer) HKUserAnnotatedMedicationQuery {
	instance := getHKUserAnnotatedMedicationQueryClass().Alloc()
	rv := objc.Send[HKUserAnnotatedMedicationQuery](instance.ID, objc.Sel("initWithPredicate:limit:resultsHandler:"), predicate, limit, resultsHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKUserAnnotatedMedicationQueryWithPredicateLimitResultsHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKUserAnnotatedMedicationQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKUserAnnotatedMedicationQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKUserAnnotatedMedicationQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKUserAnnotatedMedicationQuery */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKUserAnnotatedMedicationQuery */


