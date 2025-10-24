//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for TremorResult


// iOS-only properties

// The result’s end time and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/endDate
func (t_ TremorResult) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("endDate"))
	return rv
}

// The percentage of time when a tremor was likely, and the displacement amplitude was mild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentMild
func (t_ TremorResult) PercentMild() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("percentMild"))
	return rv
}

// The percentage of time when a tremor was likely, and the displacement amplitude was moderate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentModerate
func (t_ TremorResult) PercentModerate() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("percentModerate"))
	return rv
}

// The percentage of time when no tremor was detected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentNone
func (t_ TremorResult) PercentNone() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("percentNone"))
	return rv
}

// The percentage of time when a tremor was likely, and the displacement amplitude was slight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentSlight
func (t_ TremorResult) PercentSlight() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("percentSlight"))
	return rv
}

// The percentage of time when a tremor was likely, and the displacement amplitude was strong.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentStrong
func (t_ TremorResult) PercentStrong() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("percentStrong"))
	return rv
}

// The percentage of time when the algorithm couldn’t make a determination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/percentUnknown
func (t_ TremorResult) PercentUnknown() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("percentUnknown"))
	return rv
}

// The result’s start time and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult/startDate
func (t_ TremorResult) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("startDate"))
	return rv
}





