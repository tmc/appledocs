// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKWorkoutEffortRelationshipQuery */


/* debug [class_header]: Header for HKWorkoutEffortRelationshipQuery */
// The class instance for the [HKWorkoutEffortRelationshipQuery] class.
var (
	HKWorkoutEffortRelationshipQueryClass     _HKWorkoutEffortRelationshipQueryClass
	HKWorkoutEffortRelationshipQueryClassOnce sync.Once
)

func getHKWorkoutEffortRelationshipQueryClass() _HKWorkoutEffortRelationshipQueryClass {
	HKWorkoutEffortRelationshipQueryClassOnce.Do(func() {
		HKWorkoutEffortRelationshipQueryClass = _HKWorkoutEffortRelationshipQueryClass{objc.GetClass("HKWorkoutEffortRelationshipQuery")}
	})
	return HKWorkoutEffortRelationshipQueryClass
}

type _HKWorkoutEffortRelationshipQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkoutEffortRelationshipQuery */
// An interface definition for the [HKWorkoutEffortRelationshipQuery] class.
type IHKWorkoutEffortRelationshipQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKWorkoutEffortRelationshipQuery */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkoutEffortRelationshipQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkoutEffortRelationshipQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutEffortRelationshipQueryClass) Alloc() HKWorkoutEffortRelationshipQuery {
	rv := objc.Send[HKWorkoutEffortRelationshipQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKWorkoutEffortRelationshipQueryClass) New() HKWorkoutEffortRelationshipQuery {
	rv := objc.Send[HKWorkoutEffortRelationshipQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutEffortRelationshipQuery) Init() HKWorkoutEffortRelationshipQuery {
	rv := objc.Send[HKWorkoutEffortRelationshipQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutEffortRelationshipQuery) Autorelease() HKWorkoutEffortRelationshipQuery {
	rv := objc.Send[HKWorkoutEffortRelationshipQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutEffortRelationshipQuery creates a new HKWorkoutEffortRelationshipQuery instance.
func NewHKWorkoutEffortRelationshipQuery() HKWorkoutEffortRelationshipQuery {
	return getHKWorkoutEffortRelationshipQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkoutEffortRelationshipQuery */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationshipQuery
type HKWorkoutEffortRelationshipQuery struct {
	HKQuery
}

// HKWorkoutEffortRelationshipQueryFrom constructs a [HKWorkoutEffortRelationshipQuery] from an unsafe.Pointer.
func HKWorkoutEffortRelationshipQueryFrom(ptr unsafe.Pointer) HKWorkoutEffortRelationshipQuery {
	return HKWorkoutEffortRelationshipQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkoutEffortRelationshipQuery */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationshipQuery/init(predicate:anchor:options:resultsHandler:)
func NewHKWorkoutEffortRelationshipQueryWithPredicateAnchorOptionsResultsHandler(predicate foundation.Predicate, anchor IHKQueryAnchor, options HKWorkoutEffortRelationshipQueryOptions, resultsHandler unsafe.Pointer) HKWorkoutEffortRelationshipQuery {
	instance := getHKWorkoutEffortRelationshipQueryClass().Alloc()
	rv := objc.Send[HKWorkoutEffortRelationshipQuery](instance.ID, objc.Sel("initWithPredicate:anchor:options:resultsHandler:"), predicate, anchor, options, resultsHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutEffortRelationshipQueryWithPredicateAnchorOptionsResultsHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkoutEffortRelationshipQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkoutEffortRelationshipQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkoutEffortRelationshipQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkoutEffortRelationshipQuery */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkoutEffortRelationshipQuery */


