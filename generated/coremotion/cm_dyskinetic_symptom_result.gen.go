// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DyskineticSymptomResult] class.
var (
	DyskineticSymptomResultClass     _DyskineticSymptomResultClass
	DyskineticSymptomResultClassOnce sync.Once
)

func getDyskineticSymptomResultClass() _DyskineticSymptomResultClass {
	DyskineticSymptomResultClassOnce.Do(func() {
		DyskineticSymptomResultClass = _DyskineticSymptomResultClass{objc.GetClass("CMDyskineticSymptomResult")}
	})
	return DyskineticSymptomResultClass
}

type _DyskineticSymptomResultClass struct {
	class objc.Class
}

// An interface definition for the [DyskineticSymptomResult] class.
type IDyskineticSymptomResult interface {
	objectivec.IObject
	EndDate() foundation.NSDate
	PercentLikely() float32
	PercentUnlikely() float32
	StartDate() foundation.NSDate
}

// A result object that contains data about the likely presence of dyskinetic symptoms during a one-minute interval.
//
// Dyskinesias are uncontrolled, involuntary movements that occur as a side effect of taking Levadopa to control Parkinson’s disease. Dyskinesias can manifest in a single body part, such as the arm, leg, or head, or they can affect the entire body. Particular dyskinesias resemble actions like fidgeting, writhing, wriggling, head bobbing, or body swaying. These symptoms tend to occur during the drug’s peak dosage. Dyskinesias typically occur in patients with advanced Parkinson’s disease, who may require higher dosages of Levadopa. The following equation is always true: .


// A result object that contains data about the likely presence of dyskinetic symptoms during a one-minute interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDyskineticSymptomResult

type DyskineticSymptomResult struct {
	objectivec.Object
}

// DyskineticSymptomResultFrom constructs a [DyskineticSymptomResult] from an unsafe.Pointer.
//
// A result object that contains data about the likely presence of dyskinetic symptoms during a one-minute interval.
func DyskineticSymptomResultFrom(ptr unsafe.Pointer) DyskineticSymptomResult {
	return DyskineticSymptomResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DyskineticSymptomResultClass) Alloc() DyskineticSymptomResult {
	rv := objc.Send[DyskineticSymptomResult](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DyskineticSymptomResultClass) New() DyskineticSymptomResult {
	rv := objc.Send[DyskineticSymptomResult](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DyskineticSymptomResult) Init() DyskineticSymptomResult {
	rv := objc.Send[DyskineticSymptomResult](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DyskineticSymptomResult) Autorelease() DyskineticSymptomResult {
	rv := objc.Send[DyskineticSymptomResult](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDyskineticSymptomResult creates a new DyskineticSymptomResult instance.
func NewDyskineticSymptomResult() DyskineticSymptomResult {
	return getDyskineticSymptomResultClass().New()
}



// The result’s end time and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDyskineticSymptomResult/endDate

func (d_ DyskineticSymptomResult) EndDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("endDate"))
	return rv
}


// The percentage of time when dyskinetic symptoms were likely.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDyskineticSymptomResult/percentLikely

func (d_ DyskineticSymptomResult) PercentLikely() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("percentLikely"))
	return rv
}


// The percentage of time when dyskinetic symptoms were unlikely.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDyskineticSymptomResult/percentUnlikely

func (d_ DyskineticSymptomResult) PercentUnlikely() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("percentUnlikely"))
	return rv
}


// The result’s start time and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDyskineticSymptomResult/startDate

func (d_ DyskineticSymptomResult) StartDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("startDate"))
	return rv
}



