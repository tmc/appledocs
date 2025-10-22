// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LocationManager] class.
var (
	LocationManagerClass     _LocationManagerClass
	LocationManagerClassOnce sync.Once
)

func getLocationManagerClass() _LocationManagerClass {
	LocationManagerClassOnce.Do(func() {
		LocationManagerClass = _LocationManagerClass{objc.GetClass("CLLocationManager")}
	})
	return LocationManagerClass
}

type _LocationManagerClass struct {
	class objc.Class
}

// An interface definition for the [LocationManager] class.
type ILocationManager interface {
	objectivec.IObject
	AllowDeferredLocationUpdatesUntilTraveledTimeout(distance unsafe.Pointer, timeout foundation.ITimeInterval)
	DisallowDeferredLocationUpdates()
	DismissHeadingCalibrationDisplay()
	RequestAlwaysAuthorization()
	RequestHistoricalLocationsWithPurposeKeySampleCountCompletionHandler(purposeKey string, sampleCount int, handler unsafe.Pointer)
	RequestLocation()
	RequestStateForRegion(region ICLRegion)
	RequestTemporaryFullAccuracyAuthorizationWithPurposeKey(purposeKey string)
	RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion(purposeKey string, completion unsafe.Pointer)
	RequestWhenInUseAuthorization()
	StartMonitoringForRegion(region ICLRegion)
	StartMonitoringForRegionDesiredAccuracy(region ICLRegion, accuracy unsafe.Pointer)
	StartMonitoringLocationPushesWithCompletion(completion unsafe.Pointer)
	StartMonitoringSignificantLocationChanges()
	StartMonitoringVisits()
	StartRangingBeaconsInRegion(region ICLBeaconRegion)
	StartRangingBeaconsSatisfyingConstraint(constraint ICLBeaconIdentityConstraint)
	StartUpdatingHeading()
	StartUpdatingLocation()
	StopMonitoringForRegion(region ICLRegion)
	StopMonitoringLocationPushes()
	StopMonitoringSignificantLocationChanges()
	StopMonitoringVisits()
	StopRangingBeaconsInRegion(region ICLBeaconRegion)
	StopRangingBeaconsSatisfyingConstraint(constraint ICLBeaconIdentityConstraint)
	StopUpdatingHeading()
	StopUpdatingLocation()
	AccuracyAuthorization() AccuracyAuthorization
	ActivityType() ActivityType
	SetActivityType(value ActivityType)
	AllowsBackgroundLocationUpdates() bool
	SetAllowsBackgroundLocationUpdates(value bool)
	AuthorizationStatus() AuthorizationStatus
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	DesiredAccuracy() unsafe.Pointer
	SetDesiredAccuracy(value unsafe.Pointer)
	DistanceFilter() unsafe.Pointer
	SetDistanceFilter(value unsafe.Pointer)
	Heading() CLHeading
	HeadingAvailable() bool
	HeadingFilter() unsafe.Pointer
	SetHeadingFilter(value unsafe.Pointer)
	HeadingOrientation() DeviceOrientation
	SetHeadingOrientation(value DeviceOrientation)
	AuthorizedForWidgetUpdates() bool
	Location() CLLocation
	MaximumRegionMonitoringDistance() unsafe.Pointer
	MonitoredRegions() unsafe.Pointer
	PausesLocationUpdatesAutomatically() bool
	SetPausesLocationUpdatesAutomatically(value bool)
	RangedBeaconConstraints() unsafe.Pointer
	ShowsBackgroundLocationIndicator() bool
	SetShowsBackgroundLocationIndicator(value bool)
	CLLocationDistanceMax() unsafe.Pointer
	IsAuthorizedForWidgetUpdates() bool
	SetIsAuthorizedForWidgetUpdates(value bool)
	CLTimeIntervalMax() unsafe.Pointer
	KCLDistanceFilterNone() unsafe.Pointer
	KCLHeadingFilterNone() unsafe.Pointer
}

// The object you use to start and stop the delivery of location-related events to your app.
//
// A object is the central place to manage your app’s location-related behaviors. Use a location-manager object to configure, start, and stop location services. You might use these services to: Track large or small changes in the user’s current location with a configurable degree of accuracy. Report heading changes from the onboard compass. Monitor geographical regions of interest and generate events when someone enters or leaves those regions. Report the range to nearby Bluetooth beacons. Create one or more location-manager objects in your app and use them where you need location data. After you create a location-manager object, configure it so that Core Location knows how often to report location changes. In particular, configure the and properties with values that reflect your app’s needs. A object reports all location-related updates to its object, which is an object that conforms to the protocol. Assign the delegate immediately when you configure your location manager, because the system reports the app’s authorization status to the delegate’s method after the location manager finishes initializing itself. Core Location calls the methods of your delegate object using the of the thread on which you initialized the object. That thread must itself have an active , like the one found in your app’s main thread. For more information, see .
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


// Returns the app’s authorization status for using location services.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/authorizationStatus()
func (lc _LocationManagerClass) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(lc.class), objc.Sel("authorizationStatus"))
	return rv
}

// Returns a Boolean value indicating whether the device supports deferred location updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/deferredLocationUpdatesAvailable()
func (lc _LocationManagerClass) DeferredLocationUpdatesAvailable() bool {
	rv := objc.Send[bool](objc.ID(lc.class), objc.Sel("deferredLocationUpdatesAvailable"))
	return rv
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

// Returns a Boolean value indicating whether region monitoring is supported on the current device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/regionMonitoringAvailable()
func (lc _LocationManagerClass) RegionMonitoringAvailable() bool {
	rv := objc.Send[bool](objc.ID(lc.class), objc.Sel("regionMonitoringAvailable"))
	return rv
}

// Returns a Boolean value indicating whether region monitoring is currently enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/regionMonitoringEnabled()
func (lc _LocationManagerClass) RegionMonitoringEnabled() bool {
	rv := objc.Send[bool](objc.ID(lc.class), objc.Sel("regionMonitoringEnabled"))
	return rv
}

// Returns a Boolean value indicating whether the significant-change location service is available on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/significantLocationChangeMonitoringAvailable()
func (lc _LocationManagerClass) SignificantLocationChangeMonitoringAvailable() bool {
	rv := objc.Send[bool](objc.ID(lc.class), objc.Sel("significantLocationChangeMonitoringAvailable"))
	return rv
}

// Asks the location manager to defer the delivery of location updates until the specified criteria are met.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/allowDeferredLocationUpdates(untilTraveled:timeout:)
func (l_ LocationManager) AllowDeferredLocationUpdatesUntilTraveledTimeout(distance unsafe.Pointer, timeout foundation.ITimeInterval) {
	objc.Send[objc.ID](l_.ID, objc.Sel("allowDeferredLocationUpdatesUntilTraveled:timeout:"), distance, timeout)
}

// Cancels the deferral of location updates for this app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/disallowDeferredLocationUpdates()
func (l_ LocationManager) DisallowDeferredLocationUpdates() {
	objc.Send[objc.ID](l_.ID, objc.Sel("disallowDeferredLocationUpdates"))
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

// Retrieves the state of a region asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestState(for:)
func (l_ LocationManager) RequestStateForRegion(region ICLRegion) {
	objc.Send[objc.ID](l_.ID, objc.Sel("requestStateForRegion:"), region)
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
func (l_ LocationManager) StartMonitoringForRegion(region ICLRegion) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startMonitoringForRegion:"), region)
}

// Starts monitoring the specified region for boundary crossings.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startMonitoring(for:desiredAccuracy:)
func (l_ LocationManager) StartMonitoringForRegionDesiredAccuracy(region ICLRegion, accuracy unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startMonitoringForRegion:desiredAccuracy:"), region, accuracy)
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
func (l_ LocationManager) StartRangingBeaconsInRegion(region ICLBeaconRegion) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startRangingBeaconsInRegion:"), region)
}

// Starts the delivery of notifications for the specified beacon constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startRangingBeacons(satisfying:)
func (l_ LocationManager) StartRangingBeaconsSatisfyingConstraint(constraint ICLBeaconIdentityConstraint) {
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

// Stops monitoring the specified region.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopMonitoring(for:)
func (l_ LocationManager) StopMonitoringForRegion(region ICLRegion) {
	objc.Send[objc.ID](l_.ID, objc.Sel("stopMonitoringForRegion:"), region)
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

// Stops the delivery of notifications for the specified beacon region.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopRangingBeacons(in:)
func (l_ LocationManager) StopRangingBeaconsInRegion(region ICLBeaconRegion) {
	objc.Send[objc.ID](l_.ID, objc.Sel("stopRangingBeaconsInRegion:"), region)
}

// Stops the delivery of notifications for the specified beacon constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopRangingBeacons(satisfying:)
func (l_ LocationManager) StopRangingBeaconsSatisfyingConstraint(constraint ICLBeaconIdentityConstraint) {
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

// A value that indicates the level of location accuracy the app has permission to use.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/accuracyAuthorization
func (l_ LocationManager) AccuracyAuthorization() AccuracyAuthorization {
	rv := objc.Send[AccuracyAuthorization](l_.ID, objc.Sel("accuracyAuthorization"))
	return rv
}

// The type of activity the app expects the user to typically perform while in the app’s location session.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/activityType
func (l_ LocationManager) ActivityType() ActivityType {
	rv := objc.Send[ActivityType](l_.ID, objc.Sel("activityType"))
	return rv
}


// SetActivityType sets the value of the activityType property.
// The type of activity the app expects the user to typically perform while in the app’s location session.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/activityType
func (l_ LocationManager) SetActivityType(value ActivityType) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setActivityType:"), value)
}

// A Boolean value that indicates whether the app receives location updates when running in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/allowsBackgroundLocationUpdates
func (l_ LocationManager) AllowsBackgroundLocationUpdates() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("allowsBackgroundLocationUpdates"))
	return rv
}


// SetAllowsBackgroundLocationUpdates sets the value of the allowsBackgroundLocationUpdates property.
// A Boolean value that indicates whether the app receives location updates when running in the background.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/allowsBackgroundLocationUpdates
func (l_ LocationManager) SetAllowsBackgroundLocationUpdates(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAllowsBackgroundLocationUpdates:"), value)
}

// The current authorization status for the app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/authorizationStatus-swift.property
func (l_ LocationManager) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](l_.ID, objc.Sel("authorizationStatus"))
	return rv
}

// The delegate object to receive update events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/delegate
func (l_ LocationManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object to receive update events.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/delegate
func (l_ LocationManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelegate:"), value)
}

// The accuracy of the location data that your app wants to receive.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/desiredAccuracy
func (l_ LocationManager) DesiredAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("desiredAccuracy"))
	return rv
}


// SetDesiredAccuracy sets the value of the desiredAccuracy property.
// The accuracy of the location data that your app wants to receive.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/desiredAccuracy
func (l_ LocationManager) SetDesiredAccuracy(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDesiredAccuracy:"), value)
}

// The minimum distance in meters the device must move horizontally before an update event is generated.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/distanceFilter
func (l_ LocationManager) DistanceFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("distanceFilter"))
	return rv
}


// SetDistanceFilter sets the value of the distanceFilter property.
// The minimum distance in meters the device must move horizontally before an update event is generated.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/distanceFilter
func (l_ LocationManager) SetDistanceFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDistanceFilter:"), value)
}

// The most recently reported heading.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/heading
func (l_ LocationManager) Heading() CLHeading {
	rv := objc.Send[CLHeading](l_.ID, objc.Sel("heading"))
	return rv
}

// A Boolean value indicating whether the location manager is able to generate heading-related events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/headingAvailable-swift.property
func (l_ LocationManager) HeadingAvailable() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("headingAvailable"))
	return rv
}

// The minimum angular change in degrees required to generate new heading events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/headingFilter
func (l_ LocationManager) HeadingFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("headingFilter"))
	return rv
}


// SetHeadingFilter sets the value of the headingFilter property.
// The minimum angular change in degrees required to generate new heading events.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/headingFilter
func (l_ LocationManager) SetHeadingFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setHeadingFilter:"), value)
}

// The device orientation to use when computing heading values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/headingOrientation
func (l_ LocationManager) HeadingOrientation() DeviceOrientation {
	rv := objc.Send[DeviceOrientation](l_.ID, objc.Sel("headingOrientation"))
	return rv
}


// SetHeadingOrientation sets the value of the headingOrientation property.
// The device orientation to use when computing heading values.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/headingOrientation
func (l_ LocationManager) SetHeadingOrientation(value DeviceOrientation) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setHeadingOrientation:"), value)
}

// A Boolean value that indicates whether a widget is eligible to receive location updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/isAuthorizedForWidgetUpdates
func (l_ LocationManager) AuthorizedForWidgetUpdates() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("authorizedForWidgetUpdates"))
	return rv
}

// The most recently retrieved user location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/location
func (l_ LocationManager) Location() CLLocation {
	rv := objc.Send[CLLocation](l_.ID, objc.Sel("location"))
	return rv
}

// The largest boundary distance that can be assigned to a region.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/maximumRegionMonitoringDistance
func (l_ LocationManager) MaximumRegionMonitoringDistance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("maximumRegionMonitoringDistance"))
	return rv
}

// The set of shared regions monitored by all location-manager objects.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/monitoredRegions
func (l_ LocationManager) MonitoredRegions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("monitoredRegions"))
	return rv
}

// A Boolean value that indicates whether the location-manager object may pause location updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/pausesLocationUpdatesAutomatically
func (l_ LocationManager) PausesLocationUpdatesAutomatically() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("pausesLocationUpdatesAutomatically"))
	return rv
}


// SetPausesLocationUpdatesAutomatically sets the value of the pausesLocationUpdatesAutomatically property.
// A Boolean value that indicates whether the location-manager object may pause location updates.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/pausesLocationUpdatesAutomatically
func (l_ LocationManager) SetPausesLocationUpdatesAutomatically(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPausesLocationUpdatesAutomatically:"), value)
}

// The set of beacon constraints currently being tracked using ranging.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/rangedBeaconConstraints
func (l_ LocationManager) RangedBeaconConstraints() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("rangedBeaconConstraints"))
	return rv
}

// A Boolean value that indicates whether the status bar changes its appearance when an app uses location services in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/showsBackgroundLocationIndicator
func (l_ LocationManager) ShowsBackgroundLocationIndicator() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("showsBackgroundLocationIndicator"))
	return rv
}


// SetShowsBackgroundLocationIndicator sets the value of the showsBackgroundLocationIndicator property.
// A Boolean value that indicates whether the status bar changes its appearance when an app uses location services in the background.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/showsBackgroundLocationIndicator
func (l_ LocationManager) SetShowsBackgroundLocationIndicator(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShowsBackgroundLocationIndicator:"), value)
}

// A constant indicating the maximum distance.
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationdistancemax
func (l_ LocationManager) CLLocationDistanceMax() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("CLLocationDistanceMax"))
	return rv
}

// A Boolean value that indicates whether a widget is eligible to receive location updates.
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/isauthorizedforwidgetupdates
func (l_ LocationManager) IsAuthorizedForWidgetUpdates() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isAuthorizedForWidgetUpdates"))
	return rv
}


// SetIsAuthorizedForWidgetUpdates sets the value of the isAuthorizedForWidgetUpdates property.
// A Boolean value that indicates whether a widget is eligible to receive location updates.

//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationmanager/isauthorizedforwidgetupdates
func (l_ LocationManager) SetIsAuthorizedForWidgetUpdates(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsAuthorizedForWidgetUpdates:"), value)
}

// A value representing an unlimited amount of time.
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cltimeintervalmax
func (l_ LocationManager) CLTimeIntervalMax() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("CLTimeIntervalMax"))
	return rv
}

// A constant indicating that all movement should be reported.
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/kcldistancefilternone
func (l_ LocationManager) KCLDistanceFilterNone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("kCLDistanceFilterNone"))
	return rv
}

// A constant indicating that all header values should be reported.
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/kclheadingfilternone
func (l_ LocationManager) KCLHeadingFilterNone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("kCLHeadingFilterNone"))
	return rv
}



