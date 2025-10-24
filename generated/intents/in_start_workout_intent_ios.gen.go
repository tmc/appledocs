//go:build darwin && ios

// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for INStartWorkoutIntent


// iOS-only properties

// The numerical goal of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartWorkoutIntent/goalValue-1dzvb
func (i_ INStartWorkoutIntent) GoalValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](i_.ID, objc.Sel("goalValue"))
	return rv
}

// The units associated with the workout goal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartWorkoutIntent/workoutGoalUnitType
func (i_ INStartWorkoutIntent) WorkoutGoalUnitType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("workoutGoalUnitType"))
	return rv
}





