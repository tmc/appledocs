// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TremorResult] class.
var (
	TremorResultClass     _TremorResultClass
	TremorResultClassOnce sync.Once
)

func getTremorResultClass() _TremorResultClass {
	TremorResultClassOnce.Do(func() {
		TremorResultClass = _TremorResultClass{objc.GetClass("CMTremorResult")}
	})
	return TremorResultClass
}

type _TremorResultClass struct {
	class objc.Class
}

// An interface definition for the [TremorResult] class.
type ITremorResult interface {
	objectivec.IObject
}

// A result object that contains data about the presence and strength of tremors during a one-minute interval.
//
// The following equation is always true: .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult
type TremorResult struct {
	objectivec.Object
}

// TremorResultFrom constructs a [TremorResult] from an unsafe.Pointer.
//
// A result object that contains data about the presence and strength of tremors during a one-minute interval.
func TremorResultFrom(ptr unsafe.Pointer) TremorResult {
	return TremorResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TremorResultClass) Alloc() TremorResult {
	rv := objc.Send[TremorResult](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TremorResultClass) New() TremorResult {
	rv := objc.Send[TremorResult](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TremorResult) Init() TremorResult {
	rv := objc.Send[TremorResult](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TremorResult) Autorelease() TremorResult {
	rv := objc.Send[TremorResult](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTremorResult creates a new TremorResult instance.
func NewTremorResult() TremorResult {
	return getTremorResultClass().New()
}


// The result’s end time and date.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/endDate
func (t_ TremorResult) EndDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("endDate"))
	return rv
}

// The percentage of time when a tremor was likely, and the displacement amplitude was mild.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentMild
func (t_ TremorResult) PercentMild() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("percentMild"))
	return rv
}

// The percentage of time when a tremor was likely, and the displacement amplitude was moderate.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentModerate
func (t_ TremorResult) PercentModerate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("percentModerate"))
	return rv
}

// The percentage of time when no tremor was detected.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentNone
func (t_ TremorResult) PercentNone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("percentNone"))
	return rv
}

// The percentage of time when a tremor was likely, and the displacement amplitude was slight.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentSlight
func (t_ TremorResult) PercentSlight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("percentSlight"))
	return rv
}

// The percentage of time when a tremor was likely, and the displacement amplitude was strong.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentStrong
func (t_ TremorResult) PercentStrong() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("percentStrong"))
	return rv
}

// The percentage of time when the algorithm couldn’t make a determination.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentUnknown
func (t_ TremorResult) PercentUnknown() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("percentUnknown"))
	return rv
}

// The result’s start time and date.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/startDate
func (t_ TremorResult) StartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("startDate"))
	return rv
}



