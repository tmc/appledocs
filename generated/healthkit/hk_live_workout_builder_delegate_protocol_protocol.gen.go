// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PHKLiveWorkoutBuilderDelegate is the HKLiveWorkoutBuilderDelegate protocol interface.
//
// A protocol for monitoring live workout builders.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS +
//   - watchOS 5.0+
//
// See: doc://com.apple.healthkit/documentation/HealthKit/HKLiveWorkoutBuilderDelegate
type PHKLiveWorkoutBuilderDelegate interface {
	// Required methods
	WorkoutBuilderDidCollectDataOfTypes(workoutBuilder IHKLiveWorkoutBuilder, collectedTypes unsafe.Pointer)/* debug [protocol_interface/required_method]: WorkoutBuilderDidCollectDataOfTypes */
	WorkoutBuilderDidCollectEvent(workoutBuilder IHKLiveWorkoutBuilder)/* debug [protocol_interface/required_method]: WorkoutBuilderDidCollectEvent */
	// Optional methods
	WorkoutBuilderDidBeginActivity(workoutBuilder IHKLiveWorkoutBuilder, workoutActivity IHKWorkoutActivity)
	HasWorkoutBuilderDidBeginActivity() bool
	WorkoutBuilderDidEndActivity(workoutBuilder IHKLiveWorkoutBuilder, workoutActivity IHKWorkoutActivity)
	HasWorkoutBuilderDidEndActivity() bool
}

// HKLiveWorkoutBuilderDelegate is a delegate implementation builder for the PHKLiveWorkoutBuilderDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type HKLiveWorkoutBuilderDelegate struct {
	_WorkoutBuilderDidBeginActivity func(workoutBuilder IHKLiveWorkoutBuilder, workoutActivity IHKWorkoutActivity)
	_WorkoutBuilderDidEndActivity func(workoutBuilder IHKLiveWorkoutBuilder, workoutActivity IHKWorkoutActivity)
	_WorkoutBuilderDidCollectDataOfTypes func(workoutBuilder IHKLiveWorkoutBuilder, collectedTypes unsafe.Pointer)
	_WorkoutBuilderDidCollectEvent func(workoutBuilder IHKLiveWorkoutBuilder)
}

// SetWorkoutBuilderDidBeginActivity sets the handler for the WorkoutBuilderDidBeginActivity delegate method.
//
// Tells the delegate that a new workout activity has started.
func (d *HKLiveWorkoutBuilderDelegate) SetWorkoutBuilderDidBeginActivity(f func(workoutBuilder IHKLiveWorkoutBuilder, workoutActivity IHKWorkoutActivity)) {
	d._WorkoutBuilderDidBeginActivity = f
}

// SetWorkoutBuilderDidEndActivity sets the handler for the WorkoutBuilderDidEndActivity delegate method.
//
// Tells the delegate that the current workout activity has ended.
func (d *HKLiveWorkoutBuilderDelegate) SetWorkoutBuilderDidEndActivity(f func(workoutBuilder IHKLiveWorkoutBuilder, workoutActivity IHKWorkoutActivity)) {
	d._WorkoutBuilderDidEndActivity = f
}

// SetWorkoutBuilderDidCollectDataOfTypes sets the handler for the WorkoutBuilderDidCollectDataOfTypes delegate method.
//
// Tells the delegate that new data has been added to the builder.
func (d *HKLiveWorkoutBuilderDelegate) SetWorkoutBuilderDidCollectDataOfTypes(f func(workoutBuilder IHKLiveWorkoutBuilder, collectedTypes unsafe.Pointer)) {
	d._WorkoutBuilderDidCollectDataOfTypes = f
}

// SetWorkoutBuilderDidCollectEvent sets the handler for the WorkoutBuilderDidCollectEvent delegate method.
//
// Tells the delegate that a new event has been added to the builder.
func (d *HKLiveWorkoutBuilderDelegate) SetWorkoutBuilderDidCollectEvent(f func(workoutBuilder IHKLiveWorkoutBuilder)) {
	d._WorkoutBuilderDidCollectEvent = f
}

// WorkoutBuilderDidBeginActivity implements the PHKLiveWorkoutBuilderDelegate interface.
func (d *HKLiveWorkoutBuilderDelegate) WorkoutBuilderDidBeginActivity(workoutBuilder IHKLiveWorkoutBuilder, workoutActivity IHKWorkoutActivity) {
	if d._WorkoutBuilderDidBeginActivity != nil {
		d._WorkoutBuilderDidBeginActivity(workoutBuilder, workoutActivity)
	}
}

// HasWorkoutBuilderDidBeginActivity returns true if a handler for WorkoutBuilderDidBeginActivity has been set.
func (d *HKLiveWorkoutBuilderDelegate) HasWorkoutBuilderDidBeginActivity() bool {
	return d._WorkoutBuilderDidBeginActivity != nil
}

// WorkoutBuilderDidEndActivity implements the PHKLiveWorkoutBuilderDelegate interface.
func (d *HKLiveWorkoutBuilderDelegate) WorkoutBuilderDidEndActivity(workoutBuilder IHKLiveWorkoutBuilder, workoutActivity IHKWorkoutActivity) {
	if d._WorkoutBuilderDidEndActivity != nil {
		d._WorkoutBuilderDidEndActivity(workoutBuilder, workoutActivity)
	}
}

// HasWorkoutBuilderDidEndActivity returns true if a handler for WorkoutBuilderDidEndActivity has been set.
func (d *HKLiveWorkoutBuilderDelegate) HasWorkoutBuilderDidEndActivity() bool {
	return d._WorkoutBuilderDidEndActivity != nil
}

// WorkoutBuilderDidCollectDataOfTypes implements the PHKLiveWorkoutBuilderDelegate interface.
func (d *HKLiveWorkoutBuilderDelegate) WorkoutBuilderDidCollectDataOfTypes(workoutBuilder IHKLiveWorkoutBuilder, collectedTypes unsafe.Pointer) {
	if d._WorkoutBuilderDidCollectDataOfTypes != nil {
		d._WorkoutBuilderDidCollectDataOfTypes(workoutBuilder, collectedTypes)
	}
}

// HasWorkoutBuilderDidCollectDataOfTypes returns true if a handler for WorkoutBuilderDidCollectDataOfTypes has been set.
func (d *HKLiveWorkoutBuilderDelegate) HasWorkoutBuilderDidCollectDataOfTypes() bool {
	return d._WorkoutBuilderDidCollectDataOfTypes != nil
}

// WorkoutBuilderDidCollectEvent implements the PHKLiveWorkoutBuilderDelegate interface.
func (d *HKLiveWorkoutBuilderDelegate) WorkoutBuilderDidCollectEvent(workoutBuilder IHKLiveWorkoutBuilder) {
	if d._WorkoutBuilderDidCollectEvent != nil {
		d._WorkoutBuilderDidCollectEvent(workoutBuilder)
	}
}

// HasWorkoutBuilderDidCollectEvent returns true if a handler for WorkoutBuilderDidCollectEvent has been set.
func (d *HKLiveWorkoutBuilderDelegate) HasWorkoutBuilderDidCollectEvent() bool {
	return d._WorkoutBuilderDidCollectEvent != nil
}
