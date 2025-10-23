// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [HKUserAnnotatedMedicationQuery] class.
type IHKUserAnnotatedMedicationQuery interface {
	IHKQuery
}



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

// Alloc allocates a new instance without initialization.
func (hc _HKUserAnnotatedMedicationQueryClass) Alloc() HKUserAnnotatedMedicationQuery {
	rv := objc.Send[HKUserAnnotatedMedicationQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedicationQuery/init(predicate:limit:resultsHandler:)
func NewHKUserAnnotatedMedicationQueryWithPredicateLimitResultsHandler(predicate foundation.IPredicate, limit uint, resultsHandler unsafe.Pointer) HKUserAnnotatedMedicationQuery {
	instance := getHKUserAnnotatedMedicationQueryClass().Alloc()
	rv := objc.Send[HKUserAnnotatedMedicationQuery](instance.ID, objc.Sel("initWithPredicate:limit:resultsHandler:"), predicate, limit, resultsHandler)
	rv.Autorelease()
	return rv
}



