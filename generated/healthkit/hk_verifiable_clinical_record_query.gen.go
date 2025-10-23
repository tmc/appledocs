// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [HKVerifiableClinicalRecordQuery] class.
type IHKVerifiableClinicalRecordQuery interface {
	IHKQuery
	RecordTypes() string
	SetRecordTypes(value string)
	SourceTypes() HKVerifiableClinicalRecordSourceType
	SetSourceTypes(value HKVerifiableClinicalRecordSourceType)
}

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

// Alloc allocates a new instance without initialization.
func (hc _HKVerifiableClinicalRecordQueryClass) Alloc() HKVerifiableClinicalRecordQuery {
	rv := objc.Send[HKVerifiableClinicalRecordQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a query for one-time access to a verifiable clinical record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecordQuery/init(recordTypes:sourceTypes:predicate:resultsHandler:)
func NewHKVerifiableClinicalRecordQueryWithRecordTypesSourceTypesPredicateResultsHandler(recordTypes []string, sourceTypes []string, predicate foundation.IPredicate, resultsHandler unsafe.Pointer) HKVerifiableClinicalRecordQuery {
	instance := getHKVerifiableClinicalRecordQueryClass().Alloc()
	rv := objc.Send[HKVerifiableClinicalRecordQuery](instance.ID, objc.Sel("initWithRecordTypes:sourceTypes:predicate:resultsHandler:"), recordTypes, sourceTypes, predicate, resultsHandler)
	rv.Autorelease()
	return rv
}



// The type of records that this query returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecordquery/recordtypes
func (h_ HKVerifiableClinicalRecordQuery) RecordTypes() string {
	rv := objc.Send[string](h_.ID, objc.Sel("recordTypes"))
	return rv
}


// The type of records that this query returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecordquery/recordtypes
func (h_ HKVerifiableClinicalRecordQuery) SetRecordTypes(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setRecordTypes:"), objc.String(value))
}


// The format of the verifiable clinical record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecordquery/sourcetypes
func (h_ HKVerifiableClinicalRecordQuery) SourceTypes() HKVerifiableClinicalRecordSourceType {
	rv := objc.Send[HKVerifiableClinicalRecordSourceType](h_.ID, objc.Sel("sourceTypes"))
	return rv
}


// The format of the verifiable clinical record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecordquery/sourcetypes
func (h_ HKVerifiableClinicalRecordQuery) SetSourceTypes(value HKVerifiableClinicalRecordSourceType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSourceTypes:"), value)
}


