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
func (h_ HKHealthStore) AuthorizationViewControllerPresenter() appkit.ViewController {
	rv := objc.Send[appkit.ViewController](h_.ID, objc.Sel("authorizationViewControllerPresenter"))
	return rv
}
func (h_ HKHealthStore) SetAuthorizationViewControllerPresenter(value appkit.ViewController) {
	h_.ID.Send(objc.RegisterName("setAuthorizationViewControllerPresenter:"), value)
}





