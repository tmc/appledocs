//go:build darwin && ios

// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for HKWorkoutSession


// Returns the live workout builder associated with the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/associatedWorkoutBuilder()
func (h_ HKWorkoutSession) AssociatedWorkoutBuilder() IHKLiveWorkoutBuilder {
	rv := objc.Send[HKLiveWorkoutBuilder](h_.ID, objc.Sel("associatedWorkoutBuilder"))
	return rv
}

// Sends the provided data to the remote workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/sendToRemoteWorkoutSession(data:completion:)
func (h_ HKWorkoutSession) SendDataToRemoteWorkoutSessionCompletion(data objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("sendDataToRemoteWorkoutSession:completion:"), data, completion)
}

// Starts mirroring the workout session to the companion iOS device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/startMirroringToCompanionDevice(completion:)
func (h_ HKWorkoutSession) StartMirroringToCompanionDeviceWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("startMirroringToCompanionDeviceWithCompletion:"), completion)
}

// Stops mirroring the workout session to the companion iOS device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/stopMirroringToCompanionDevice(completion:)
func (h_ HKWorkoutSession) StopMirroringToCompanionDeviceWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("stopMirroringToCompanionDeviceWithCompletion:"), completion)
}

// iOS-only properties




