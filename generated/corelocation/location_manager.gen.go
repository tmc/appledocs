// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LocationManager] class.
var (
	locationManagerClass     _LocationManagerClass
	locationManagerClassOnce sync.Once
)

func getLocationManagerClass() _LocationManagerClass {
	locationManagerClassOnce.Do(func() {
		locationManagerClass = _LocationManagerClass{objc.GetClass("CLLocationManager")}
	})
	return locationManagerClass
}

type _LocationManagerClass struct {
	class objc.Class
}

// An interface definition for the [LocationManager] class.
type ILocationManager interface {
	objectivec.IObject
	DismissHeadingCalibrationDisplay()
	RequestAlwaysAuthorization()
	RequestHistoricalLocationsWithPurposeKeySampleCountCompletionHandler(purposeKey string, sampleCount int, handler unsafe.Pointer)
	RequestLocation()
	RequestTemporaryFullAccuracyAuthorizationWithPurposeKey(purposeKey string)
	RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion(purposeKey string, completion unsafe.Pointer)
	RequestWhenInUseAuthorization()
	StartMonitoringForRegion(region unsafe.Pointer)
	StartMonitoringLocationPushesWithCompletion(completion unsafe.Pointer)
	StartMonitoringSignificantLocationChanges()
	StartMonitoringVisits()
	StartRangingBeaconsInRegion(region unsafe.Pointer)
	StartRangingBeaconsSatisfyingConstraint(constraint unsafe.Pointer)
	StartUpdatingHeading()
	StartUpdatingLocation()
	StopMonitoringLocationPushes()
	StopMonitoringSignificantLocationChanges()
	StopMonitoringVisits()
	StopRangingBeaconsSatisfyingConstraint(constraint unsafe.Pointer)
	StopUpdatingHeading()
	StopUpdatingLocation()
}

// The object you use to start and stop the delivery of location-related events to your app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager
type LocationManager struct {
	objectivec.Object
}

// LocationManagerFrom constructs a [LocationManager] from an unsafe.Pointer.
//
// The object you use to start and stop the delivery of location-related events to your app.
func LocationManagerFrom(ptr unsafe.Pointer) LocationManager {
	return LocationManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LocationManagerClass) Alloc() LocationManager {
	rv := objc.Send[LocationManager](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LocationManagerClass) New() LocationManager {
	rv := objc.Send[LocationManager](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LocationManager) Init() LocationManager {
	rv := objc.Send[LocationManager](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LocationManager) Autorelease() LocationManager {
	rv := objc.Send[LocationManager](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocationManager creates a new LocationManager instance.
func NewLocationManager() LocationManager {
	return getLocationManagerClass().New()
}


// Returns a Boolean value indicating whether the location manager is able to generate heading-related events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/headingAvailable()
func (lc _LocationManagerClass) HeadingAvailable() bool {
	rv := objc.Send[bool](objc.ID(lc.class), objc.Sel("headingAvailable"))
	return rv
}
// Returns a Boolean value indicating whether the device supports region monitoring using the specified class.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/isMonitoringAvailable(for:)
func (lc _LocationManagerClass) IsMonitoringAvailableForClass(regionClass objc.Class) bool {
	rv := objc.Send[bool](objc.ID(lc.class), objc.Sel("isMonitoringAvailableForClass:"), regionClass)
	return rv
}
// Returns a Boolean value indicating whether the device supports ranging of beacons that use the iBeacon protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/isRangingAvailable()
func (lc _LocationManagerClass) IsRangingAvailable() bool {
	rv := objc.Send[bool](objc.ID(lc.class), objc.Sel("isRangingAvailable"))
	return rv
}
// Returns a Boolean value indicating whether location services are enabled on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/locationServicesEnabled()
func (lc _LocationManagerClass) LocationServicesEnabled() bool {
	rv := objc.Send[bool](objc.ID(lc.class), objc.Sel("locationServicesEnabled"))
	return rv
}
// Returns a Boolean value indicating whether the significant-change location service is available on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/significantLocationChangeMonitoringAvailable()
func (lc _LocationManagerClass) SignificantLocationChangeMonitoringAvailable() bool {
	rv := objc.Send[bool](objc.ID(lc.class), objc.Sel("significantLocationChangeMonitoringAvailable"))
	return rv
}
// Dismisses the heading calibration view from the screen immediately.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/dismissHeadingCalibrationDisplay()
func (l_ LocationManager) DismissHeadingCalibrationDisplay() {
	objc.Send[objc.ID](l_.ID, objc.Sel("dismissHeadingCalibrationDisplay"))
}
// Requests the user’s permission to use location services regardless of whether the app is in use.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestAlwaysAuthorization()
func (l_ LocationManager) RequestAlwaysAuthorization() {
	objc.Send[objc.ID](l_.ID, objc.Sel("requestAlwaysAuthorization"))
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestHistoricalLocations(purposeKey:sampleCount:completionHandler:)
func (l_ LocationManager) RequestHistoricalLocationsWithPurposeKeySampleCountCompletionHandler(purposeKey string, sampleCount int, handler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("requestHistoricalLocationsWithPurposeKey:sampleCount:completionHandler:"), objc.String(purposeKey), sampleCount, handler)
}
// Requests the one-time delivery of the user’s current location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestLocation()
func (l_ LocationManager) RequestLocation() {
	objc.Send[objc.ID](l_.ID, objc.Sel("requestLocation"))
}
// Requests permission to temporarily use location services with full accuracy.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestTemporaryFullAccuracyAuthorization(withPurposeKey:)
func (l_ LocationManager) RequestTemporaryFullAccuracyAuthorizationWithPurposeKey(purposeKey string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("requestTemporaryFullAccuracyAuthorizationWithPurposeKey:"), objc.String(purposeKey))
}
// Requests permission to temporarily use location services with full accuracy and reports the results to the provided completion handler.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestTemporaryFullAccuracyAuthorization(withPurposeKey:completion:)
func (l_ LocationManager) RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion(purposeKey string, completion unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("requestTemporaryFullAccuracyAuthorizationWithPurposeKey:completion:"), objc.String(purposeKey), completion)
}
// Requests the user’s permission to use location services while the app is in use.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestWhenInUseAuthorization()
func (l_ LocationManager) RequestWhenInUseAuthorization() {
	objc.Send[objc.ID](l_.ID, objc.Sel("requestWhenInUseAuthorization"))
}
// Starts monitoring the specified region.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startMonitoring(for:)
func (l_ LocationManager) StartMonitoringForRegion(region unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startMonitoringForRegion:"), region)
}
// Starts monitoring for the delivery of Apple Push Notification service (APNs) location pushes, and provides a device-specific token for sending pushes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startMonitoringLocationPushes(completion:)
func (l_ LocationManager) StartMonitoringLocationPushesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startMonitoringLocationPushesWithCompletion:"), completion)
}
// Starts the generation of updates based on significant location changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startMonitoringSignificantLocationChanges()
func (l_ LocationManager) StartMonitoringSignificantLocationChanges() {
	objc.Send[objc.ID](l_.ID, objc.Sel("startMonitoringSignificantLocationChanges"))
}
// Starts the delivery of visit-related events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startMonitoringVisits()
func (l_ LocationManager) StartMonitoringVisits() {
	objc.Send[objc.ID](l_.ID, objc.Sel("startMonitoringVisits"))
}
// Starts the delivery of notifications for the specified beacon region.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startRangingBeacons(in:)
func (l_ LocationManager) StartRangingBeaconsInRegion(region unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startRangingBeaconsInRegion:"), region)
}
// Starts the delivery of notifications for the specified beacon constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startRangingBeacons(satisfying:)
func (l_ LocationManager) StartRangingBeaconsSatisfyingConstraint(constraint unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startRangingBeaconsSatisfyingConstraint:"), constraint)
}
// Starts the generation of updates that report the user’s current heading.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startUpdatingHeading()
func (l_ LocationManager) StartUpdatingHeading() {
	objc.Send[objc.ID](l_.ID, objc.Sel("startUpdatingHeading"))
}
// Starts the generation of updates that report the user’s current location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startUpdatingLocation()
func (l_ LocationManager) StartUpdatingLocation() {
	objc.Send[objc.ID](l_.ID, objc.Sel("startUpdatingLocation"))
}
// Stops monitoring for Apple Push Notification service (APNs) location pushes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopMonitoringLocationPushes()
func (l_ LocationManager) StopMonitoringLocationPushes() {
	objc.Send[objc.ID](l_.ID, objc.Sel("stopMonitoringLocationPushes"))
}
// Stops the delivery of location events based on significant location changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopMonitoringSignificantLocationChanges()
func (l_ LocationManager) StopMonitoringSignificantLocationChanges() {
	objc.Send[objc.ID](l_.ID, objc.Sel("stopMonitoringSignificantLocationChanges"))
}
// Stops the delivery of visit-related events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopMonitoringVisits()
func (l_ LocationManager) StopMonitoringVisits() {
	objc.Send[objc.ID](l_.ID, objc.Sel("stopMonitoringVisits"))
}
// Stops the delivery of notifications for the specified beacon constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopRangingBeacons(satisfying:)
func (l_ LocationManager) StopRangingBeaconsSatisfyingConstraint(constraint unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("stopRangingBeaconsSatisfyingConstraint:"), constraint)
}
// Stops the generation of heading updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopUpdatingHeading()
func (l_ LocationManager) StopUpdatingHeading() {
	objc.Send[objc.ID](l_.ID, objc.Sel("stopUpdatingHeading"))
}
// Stops the generation of location updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopUpdatingLocation()
func (l_ LocationManager) StopUpdatingLocation() {
	objc.Send[objc.ID](l_.ID, objc.Sel("stopUpdatingLocation"))
}


