//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MovementDisorderManager


// Returns the date of the most recently calculated results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/lastProcessedDate()
func (m_ MovementDisorderManager) LastProcessedDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("lastProcessedDate"))
	return rv
}

// Calculate and store tremor and dyskinetic symptom results for the duration of the specified time interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/monitorKinesias(forDuration:)
func (m_ MovementDisorderManager) MonitorKinesiasForDuration(duration float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("monitorKinesiasForDuration:"), duration)
}

// Returns the expiration date for the most recent monitoring period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/monitorKinesiasExpirationDate()
func (m_ MovementDisorderManager) MonitorKinesiasExpirationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("monitorKinesiasExpirationDate"))
	return rv
}

// Query for dyskinetic symptoms from the provided time interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/queryDyskineticSymptom(from:to:withHandler:)
func (m_ MovementDisorderManager) QueryDyskineticSymptomFromDateToDateWithHandler(fromDate objc.IObject /* cross-framework: NSDate */, toDate objc.IObject /* cross-framework: NSDate */, handler DyskineticSymptomResultHandler /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("queryDyskineticSymptomFromDate:toDate:withHandler:"), fromDate, toDate, handler)
}

// Query for tremor results from the provided time interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/queryTremor(from:to:withHandler:)
func (m_ MovementDisorderManager) QueryTremorFromDateToDateWithHandler(fromDate objc.IObject /* cross-framework: NSDate */, toDate objc.IObject /* cross-framework: NSDate */, handler TremorResultHandler /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("queryTremorFromDate:toDate:withHandler:"), fromDate, toDate, handler)
}

// iOS-only properties





