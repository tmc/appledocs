// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation


// Enum types and constants

// CLAccuracyAuthorization - Constants that indicate the level of location accuracy the app has authorization to use.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAccuracyAuthorization
type CLAccuracyAuthorization uint

const (
	// CLAccuracyAuthorizationFullAccuracy - The user authorized the app to access location data with full accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAccuracyAuthorization/fullAccuracy
	CLAccuracyAuthorizationFullAccuracy CLAccuracyAuthorization = 0
	// CLAccuracyAuthorizationReducedAccuracy - The user authorized the app to access location data with reduced accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAccuracyAuthorization/reducedAccuracy
	CLAccuracyAuthorizationReducedAccuracy CLAccuracyAuthorization = 0
)


// CLAuthorizationStatus - Constants that indicate the app’s authorization to use location services.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus
type CLAuthorizationStatus uint

const (
	// kCLAuthorizationStatusAuthorized - The user authorized the app to use location services.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/authorized
	kCLAuthorizationStatusAuthorized CLAuthorizationStatus = 0
	// kCLAuthorizationStatusAuthorizedAlways - The user authorized the app to start location services at any time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/authorizedAlways
	kCLAuthorizationStatusAuthorizedAlways CLAuthorizationStatus = 0
	// kCLAuthorizationStatusAuthorizedWhenInUse - The user authorized the app to start location services while it is in use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/authorizedWhenInUse
	kCLAuthorizationStatusAuthorizedWhenInUse CLAuthorizationStatus = 0
	// kCLAuthorizationStatusDenied - The user denied the use of location services for the app or they are disabled globally in Settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/denied
	kCLAuthorizationStatusDenied CLAuthorizationStatus = 0
	// kCLAuthorizationStatusNotDetermined - The user has not chosen whether the app can use location services.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/notDetermined
	kCLAuthorizationStatusNotDetermined CLAuthorizationStatus = 0
	// kCLAuthorizationStatusRestricted - The app is not authorized to use location services.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/restricted
	kCLAuthorizationStatusRestricted CLAuthorizationStatus = 0
)


// CLError - Error codes returned by the location manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code
type CLError uint

const (
	// kCLErrorDeferredAccuracyTooLow - A constant that indicates deferred mode isn’t supported for the requested accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/deferredAccuracyTooLow
	kCLErrorDeferredAccuracyTooLow CLError = 0
	// kCLErrorDeferredCanceled - A constant that indicates your app or the location manager canceled the request for deferred updates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/deferredCanceled
	kCLErrorDeferredCanceled CLError = 0
	// kCLErrorDeferredDistanceFiltered - A constant that indicates deferred mode doesn’t support distance filters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/deferredDistanceFiltered
	kCLErrorDeferredDistanceFiltered CLError = 0
	// kCLErrorDeferredFailed - A constant that indicates the location manager didn’t enter deferred mode for an unknown reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/deferredFailed
	kCLErrorDeferredFailed CLError = 0
	// kCLErrorDeferredNotUpdatingLocation - A constant that indicates the location manager didn’t enter deferred mode because location updates were already disabled or paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/deferredNotUpdatingLocation
	kCLErrorDeferredNotUpdatingLocation CLError = 0
	// kCLErrorDenied - A constant that indicates the user denied access to the location service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/denied
	kCLErrorDenied CLError = 0
	// kCLErrorGeocodeCanceled - A constant that indicates the geocode request was canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/geocodeCanceled
	kCLErrorGeocodeCanceled CLError = 0
	// kCLErrorGeocodeFoundNoResult - A constant that indicates the geocode request yielded no result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/geocodeFoundNoResult
	kCLErrorGeocodeFoundNoResult CLError = 0
	// kCLErrorGeocodeFoundPartialResult - A constant that indicates the geocode request yielded a partial result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/geocodeFoundPartialResult
	kCLErrorGeocodeFoundPartialResult CLError = 0
	// kCLErrorHeadingFailure - A constant that indicates the location manager can’t determine the heading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/headingFailure
	kCLErrorHeadingFailure CLError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/historicalLocationError
	kCLErrorHistoricalLocationError CLError = 0
	// kCLErrorLocationUnknown - A constant that indicates the location manager was unable to obtain a location value right now.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/locationUnknown
	kCLErrorLocationUnknown CLError = 0
	// kCLErrorNetwork - A constant that indicates the network was unavailable or a network error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/network
	kCLErrorNetwork CLError = 0
	// kCLErrorPromptDeclined - A constant that indicates the user didn’t grant the requested temporary authorization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/promptDeclined
	kCLErrorPromptDeclined CLError = 0
	// kCLErrorRangingFailure - A constant that indicates a general ranging error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/rangingFailure
	kCLErrorRangingFailure CLError = 0
	// kCLErrorRangingUnavailable - A constant that indicates ranging is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/rangingUnavailable
	kCLErrorRangingUnavailable CLError = 0
	// kCLErrorRegionMonitoringDenied - A constant that indicates the user denied access to the region monitoring service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/regionMonitoringDenied
	kCLErrorRegionMonitoringDenied CLError = 0
	// kCLErrorRegionMonitoringFailure - A constant that indicates the location manager failed to monitor a registered region.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/regionMonitoringFailure
	kCLErrorRegionMonitoringFailure CLError = 0
	// kCLErrorRegionMonitoringResponseDelayed - A constant that indicates Core Location will deliver events but they may be delayed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/regionMonitoringResponseDelayed
	kCLErrorRegionMonitoringResponseDelayed CLError = 0
	// kCLErrorRegionMonitoringSetupDelayed - A constant that indicates Core Location failed to initialize the region monitoring feature.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/regionMonitoringSetupDelayed
	kCLErrorRegionMonitoringSetupDelayed CLError = 0
)


// CLLocationPushServiceError - Error codes the location manager returns if starting to monitor for location push notifications fails.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code
type CLLocationPushServiceError uint

const (
	// CLLocationPushServiceErrorMissingEntitlement - An error code that indicates the app is missing the entitlement it needs to use the location push service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code/missingEntitlement
	CLLocationPushServiceErrorMissingEntitlement CLLocationPushServiceError = 0
	// CLLocationPushServiceErrorMissingPushExtension - An error code that indicates the app is missing a Location Push Service Extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code/missingPushExtension
	CLLocationPushServiceErrorMissingPushExtension CLLocationPushServiceError = 0
	// CLLocationPushServiceErrorMissingPushServerEnvironment - An error code that indicates the app is missing an Apple Push Notification service (APNs) environment entitlement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code/missingPushServerEnvironment
	CLLocationPushServiceErrorMissingPushServerEnvironment CLLocationPushServiceError = 0
	// CLLocationPushServiceErrorUnknown - An error code that indicates the app was unable to start the location push service for an unknown reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code/unknown
	CLLocationPushServiceErrorUnknown CLLocationPushServiceError = 0
	// CLLocationPushServiceErrorUnsupportedPlatform - An error code that indicates the location push service isn’t available on this platform.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code/unsupportedPlatform
	CLLocationPushServiceErrorUnsupportedPlatform CLLocationPushServiceError = 0
)


// CLServiceSessionAuthorizationRequirement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionAuthorizationRequirement
type CLServiceSessionAuthorizationRequirement int

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionAuthorizationRequirement/CLServiceSessionAuthorizationRequirementAlways
	CLServiceSessionAuthorizationRequirementAlways CLServiceSessionAuthorizationRequirement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionAuthorizationRequirement/CLServiceSessionAuthorizationRequirementNone
	CLServiceSessionAuthorizationRequirementNone CLServiceSessionAuthorizationRequirement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionAuthorizationRequirement/CLServiceSessionAuthorizationRequirementWhenInUse
	CLServiceSessionAuthorizationRequirementWhenInUse CLServiceSessionAuthorizationRequirement = 0
)


// CLActivityType - Constants that indicate the type of activity associated with location updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType
type CLActivityType uint

const (
	// CLActivityTypeAirborne - The value that indicates activities in the air.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType/airborne
	CLActivityTypeAirborne CLActivityType = 0
	// CLActivityTypeAutomotiveNavigation - The value that indicates positioning in an automobile following a road network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType/automotiveNavigation
	CLActivityTypeAutomotiveNavigation CLActivityType = 0
	// CLActivityTypeFitness - The value that indicates positioning during dedicated fitness sessions, such as walking workouts, running workouts, cycling workouts, and so on.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType/fitness
	CLActivityTypeFitness CLActivityType = 0
	// CLActivityTypeOther - The value that indicates the app is using location manager for an unspecified activity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType/other
	CLActivityTypeOther CLActivityType = 0
	// CLActivityTypeOtherNavigation - The value that indicates positioning for activities that don’t or may not adhere to roads such as cycling, scooters, trains, boats and off-road vehicles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType/otherNavigation
	CLActivityTypeOtherNavigation CLActivityType = 0
)


// CLDeviceOrientation - Constants indicating the physical orientation of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLDeviceOrientation
type CLDeviceOrientation uint

const (
	// CLDeviceOrientationFaceDown - The device is held parallel to the ground with the screen facing downwards.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLDeviceOrientation/faceDown
	CLDeviceOrientationFaceDown CLDeviceOrientation = 0
	// CLDeviceOrientationFaceUp - The device is held parallel to the ground with the screen facing upwards.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLDeviceOrientation/faceUp
	CLDeviceOrientationFaceUp CLDeviceOrientation = 0
	// CLDeviceOrientationLandscapeLeft - The device is in landscape mode, with the device held upright and the home button on the right side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLDeviceOrientation/landscapeLeft
	CLDeviceOrientationLandscapeLeft CLDeviceOrientation = 0
	// CLDeviceOrientationLandscapeRight - The device is in landscape mode, with the device held upright and the home button on the left side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLDeviceOrientation/landscapeRight
	CLDeviceOrientationLandscapeRight CLDeviceOrientation = 0
	// CLDeviceOrientationPortrait - The device is in portrait mode, with the device held upright and the home button at the bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLDeviceOrientation/portrait
	CLDeviceOrientationPortrait CLDeviceOrientation = 0
	// CLDeviceOrientationPortraitUpsideDown - The device is in portrait mode but upside down, with the device held upright and the home button at the top.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLDeviceOrientation/portraitUpsideDown
	CLDeviceOrientationPortraitUpsideDown CLDeviceOrientation = 0
	// CLDeviceOrientationUnknown - The orientation is currently not known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLDeviceOrientation/unknown
	CLDeviceOrientationUnknown CLDeviceOrientation = 0
)


// CLLiveUpdateConfiguration - Specifies the types of locations that a location updater generates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLiveUpdateConfiguration
type CLLiveUpdateConfiguration int

const (
	// CLLiveUpdateConfigurationAirborne - A configuration for airborne use cases.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLiveUpdateConfiguration/CLLiveUpdateConfigurationAirborne
	CLLiveUpdateConfigurationAirborne CLLiveUpdateConfiguration = 0
	// CLLiveUpdateConfigurationAutomotiveNavigation - A configuration for automotive navigation use cases.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLiveUpdateConfiguration/CLLiveUpdateConfigurationAutomotiveNavigation
	CLLiveUpdateConfigurationAutomotiveNavigation CLLiveUpdateConfiguration = 0
	// CLLiveUpdateConfigurationDefault - The default configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLiveUpdateConfiguration/CLLiveUpdateConfigurationDefault
	CLLiveUpdateConfigurationDefault CLLiveUpdateConfiguration = 0
	// CLLiveUpdateConfigurationFitness - A configuration for fitness use cases.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLiveUpdateConfiguration/CLLiveUpdateConfigurationFitness
	CLLiveUpdateConfigurationFitness CLLiveUpdateConfiguration = 0
	// CLLiveUpdateConfigurationOtherNavigation - A configuration for other navigation use cases.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLiveUpdateConfiguration/CLLiveUpdateConfigurationOtherNavigation
	CLLiveUpdateConfigurationOtherNavigation CLLiveUpdateConfiguration = 0
)


// CLMonitoringState - Values that represent the current state of a monitoring condition.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState
type CLMonitoringState uint

const (
	// CLMonitoringStateSatisfied - The condition is in a satisfied state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState/CLMonitoringStateSatisfied
	CLMonitoringStateSatisfied CLMonitoringState = 0
	// CLMonitoringStateUnknown - The condition is in an unknown state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState/CLMonitoringStateUnknown
	CLMonitoringStateUnknown CLMonitoringState = 0
	// CLMonitoringStateUnmonitored - The condition is in an unmonitored state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState/CLMonitoringStateUnmonitored
	CLMonitoringStateUnmonitored CLMonitoringState = 0
	// CLMonitoringStateUnsatisfied - The condition is in an unsatisfied state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState/CLMonitoringStateUnsatisfied
	CLMonitoringStateUnsatisfied CLMonitoringState = 0
)


// CLProximity - Constants that reflect the relative distance to a beacon.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity
type CLProximity uint

const (
	// CLProximityFar - The beacon is far away.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity/far
	CLProximityFar CLProximity = 0
	// CLProximityImmediate - The beacon is in the user’s immediate vicinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity/immediate
	CLProximityImmediate CLProximity = 0
	// CLProximityNear - The beacon is relatively close to the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity/near
	CLProximityNear CLProximity = 0
	// CLProximityUnknown - The proximity of the beacon could not be determined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity/unknown
	CLProximityUnknown CLProximity = 0
)


// CLRegionState - Constants that reflect the relationship of the current location to the region boundaries.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegionState
type CLRegionState uint

const (
	// CLRegionStateInside - The location is inside of the given region.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegionState/inside
	CLRegionStateInside CLRegionState = 0
	// CLRegionStateOutside - The location is outside of the given region.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegionState/outside
	CLRegionStateOutside CLRegionState = 0
	// CLRegionStateUnknown - It is unknown whether the location is inside or outside of the region.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegionState/unknown
	CLRegionStateUnknown CLRegionState = 0
)


