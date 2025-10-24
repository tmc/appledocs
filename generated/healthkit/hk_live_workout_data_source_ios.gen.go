//go:build darwin && ios

// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for HKLiveWorkoutDataSource

// Stops automatically calculating statistics for the quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutDataSource/disableCollection(for:)
func (h_ HKLiveWorkoutDataSource) DisableCollectionForType(quantityType IHKQuantityType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("disableCollectionForType:"), quantityType)
}

// iOS-only properties
