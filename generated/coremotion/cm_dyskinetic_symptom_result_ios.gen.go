//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for DyskineticSymptomResult


// iOS-only properties

// The result’s end time and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDyskineticSymptomResult/endDate
func (d_ DyskineticSymptomResult) EndDate() objc.IObject /* cross-framework: NSDate */ {
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
func (d_ DyskineticSymptomResult) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("startDate"))
	return rv
}





