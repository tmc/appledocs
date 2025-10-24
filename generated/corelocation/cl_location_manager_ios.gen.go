//go:build darwin && ios

// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for LocationManager

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestHistoricalLocations(purposeKey:sampleCount:completionHandler:)
func (l_ LocationManager) RequestHistoricalLocationsWithPurposeKeySampleCountCompletionHandler(purposeKey objc.IObject /* cross-framework: NSString */, sampleCount int, handler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("requestHistoricalLocationsWithPurposeKey:sampleCount:completionHandler:"), purposeKey, sampleCount, handler)
}

// Starts monitoring for the delivery of Apple Push Notification service (APNs) location pushes, and provides a device-specific token for sending pushes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startMonitoringLocationPushes(completion:)
func (l_ LocationManager) StartMonitoringLocationPushesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startMonitoringLocationPushesWithCompletion:"), completion)
}

// Stops monitoring for Apple Push Notification service (APNs) location pushes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopMonitoringLocationPushes()
func (l_ LocationManager) StopMonitoringLocationPushes() {
	objc.Send[objc.ID](l_.ID, objc.Sel("stopMonitoringLocationPushes"))
}

// Stops the generation of heading updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopUpdatingHeading()
func (l_ LocationManager) StopUpdatingHeading() {
	objc.Send[objc.ID](l_.ID, objc.Sel("stopUpdatingHeading"))
}

// iOS-only properties

// A Boolean value that indicates whether the status bar changes its appearance when an app uses location services in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/showsBackgroundLocationIndicator
func (l_ LocationManager) ShowsBackgroundLocationIndicator() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("showsBackgroundLocationIndicator"))
	return rv
}
func (l_ LocationManager) SetShowsBackgroundLocationIndicator(value bool) {
	l_.ID.Send(objc.RegisterName("setShowsBackgroundLocationIndicator:"), value)
}
