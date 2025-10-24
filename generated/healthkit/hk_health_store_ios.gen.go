//go:build darwin && ios

// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for HKHealthStore


// Recovers an active workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/recoverActiveWorkoutSession(completion:)
func (h_ HKHealthStore) RecoverActiveWorkoutSessionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("recoverActiveWorkoutSessionWithCompletion:"), completion)
}

// iOS-only properties

// The view controller that presents HealthKit authorization sheets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/authorizationViewControllerPresenter
func (h_ HKHealthStore) AuthorizationViewControllerPresenter() objc.IObject /* cross-framework: ViewController */ {
	rv := objc.Send[appkit.ViewController](h_.ID, objc.Sel("authorizationViewControllerPresenter"))
	return rv
}
func (h_ HKHealthStore) SetAuthorizationViewControllerPresenter(value objc.IObject /* cross-framework: ViewController */) {
	h_.ID.Send(objc.RegisterName("setAuthorizationViewControllerPresenter:"), value)
}

// A block that the system calls when it starts a mirrored workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/workoutSessionMirroringStartHandler
func (h_ HKHealthStore) WorkoutSessionMirroringStartHandler() func(unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer)](h_.ID, objc.Sel("workoutSessionMirroringStartHandler"))
	return rv
}
func (h_ HKHealthStore) SetWorkoutSessionMirroringStartHandler(value func(unsafe.Pointer)) {
	h_.ID.Send(objc.RegisterName("setWorkoutSessionMirroringStartHandler:"), value)
}





