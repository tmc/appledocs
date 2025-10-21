// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

// Enum types and constants
// CLAccuracyAuthorization - Constants that indicate the level of location accuracy the app has authorization to use.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAccuracyAuthorization
type AccuracyAuthorization uint

const (
// AccuracyAuthorizationFullAccuracy - The user authorized the app to access location data with full accuracy.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAccuracyAuthorization/fullAccuracy
AccuracyAuthorizationFullAccuracy AccuracyAuthorization = 0
// AccuracyAuthorizationReducedAccuracy - The user authorized the app to access location data with reduced accuracy.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAccuracyAuthorization/reducedAccuracy
AccuracyAuthorizationReducedAccuracy AccuracyAuthorization = 0
)

// CLActivityType - Constants that indicate the type of activity associated with location updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType
type ActivityType uint

const (
// ActivityTypeAirborne - The value that indicates activities in the air.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType/airborne
ActivityTypeAirborne ActivityType = 0
// ActivityTypeAutomotiveNavigation - The value that indicates positioning in an automobile following a road network.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType/automotiveNavigation
ActivityTypeAutomotiveNavigation ActivityType = 0
// ActivityTypeFitness - The value that indicates positioning during dedicated fitness sessions, such as walking workouts, running workouts, cycling workouts, and so on.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType/fitness
ActivityTypeFitness ActivityType = 0
// ActivityTypeOther - The value that indicates the app is using location manager for an unspecified activity.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType/other
ActivityTypeOther ActivityType = 0
// ActivityTypeOtherNavigation - The value that indicates positioning for activities that don’t or may not adhere to roads such as cycling, scooters, trains, boats and off-road vehicles.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType/otherNavigation
ActivityTypeOtherNavigation ActivityType = 0
)

// CLAuthorizationStatus - Constants that indicate the app’s authorization to use location services.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus
type AuthorizationStatus uint

const (
// kCLAuthorizationStatusAuthorized - The user authorized the app to use location services.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/authorized
kCLAuthorizationStatusAuthorized AuthorizationStatus = 0
// kCLAuthorizationStatusAuthorizedAlways - The user authorized the app to start location services at any time.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/authorizedAlways
kCLAuthorizationStatusAuthorizedAlways AuthorizationStatus = 0
// kCLAuthorizationStatusAuthorizedWhenInUse - The user authorized the app to start location services while it is in use.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/authorizedWhenInUse
kCLAuthorizationStatusAuthorizedWhenInUse AuthorizationStatus = 0
// kCLAuthorizationStatusDenied - The user denied the use of location services for the app or they are disabled globally in Settings.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/denied
kCLAuthorizationStatusDenied AuthorizationStatus = 0
// kCLAuthorizationStatusNotDetermined - The user has not chosen whether the app can use location services.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/notDetermined
kCLAuthorizationStatusNotDetermined AuthorizationStatus = 0
// kCLAuthorizationStatusRestricted - The app is not authorized to use location services.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus/restricted
kCLAuthorizationStatusRestricted AuthorizationStatus = 0
)

// CLDeviceOrientation - Constants indicating the physical orientation of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLDeviceOrientation
type DeviceOrientation uint

// CLError - Error codes returned by the location manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code
type Error uint

const (
// kCLErrorDeferredAccuracyTooLow - A constant that indicates deferred mode isn’t supported for the requested accuracy.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/deferredAccuracyTooLow
kCLErrorDeferredAccuracyTooLow Error = 0
// kCLErrorDeferredCanceled - A constant that indicates your app or the location manager canceled the request for deferred updates.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/deferredCanceled
kCLErrorDeferredCanceled Error = 0
// kCLErrorDeferredDistanceFiltered - A constant that indicates deferred mode doesn’t support distance filters.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/deferredDistanceFiltered
kCLErrorDeferredDistanceFiltered Error = 0
// kCLErrorDeferredFailed - A constant that indicates the location manager didn’t enter deferred mode for an unknown reason.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/deferredFailed
kCLErrorDeferredFailed Error = 0
// kCLErrorDenied - A constant that indicates the user denied access to the location service.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/denied
kCLErrorDenied Error = 0
// kCLErrorHeadingFailure - A constant that indicates the location manager can’t determine the heading.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/headingFailure
kCLErrorHeadingFailure Error = 0
// kCLErrorLocationUnknown - A constant that indicates the location manager was unable to obtain a location value right now.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/locationUnknown
kCLErrorLocationUnknown Error = 0
// kCLErrorNetwork - A constant that indicates the network was unavailable or a network error occurred.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/network
kCLErrorNetwork Error = 0
// kCLErrorPromptDeclined - A constant that indicates the user didn’t grant the requested temporary authorization.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/promptDeclined
kCLErrorPromptDeclined Error = 0
// kCLErrorRangingFailure - A constant that indicates a general ranging error occurred.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/rangingFailure
kCLErrorRangingFailure Error = 0
// kCLErrorRangingUnavailable - A constant that indicates ranging is disabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/rangingUnavailable
kCLErrorRangingUnavailable Error = 0
)

// CLLiveUpdateConfiguration - Specifies the types of locations that a location updater generates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLiveUpdateConfiguration
type LiveUpdateConfiguration uint

const (
// LiveUpdateConfigurationDefault - The default configuration.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLiveUpdateConfiguration/CLLiveUpdateConfigurationDefault
LiveUpdateConfigurationDefault LiveUpdateConfiguration = 0
)

// CLLocationPushServiceError - Error codes the location manager returns if starting to monitor for location push notifications fails.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code
type LocationPushServiceError uint

const (
// LocationPushServiceErrorMissingEntitlement - An error code that indicates the app is missing the entitlement it needs to use the location push service.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code/missingEntitlement
LocationPushServiceErrorMissingEntitlement LocationPushServiceError = 0
// LocationPushServiceErrorMissingPushExtension - An error code that indicates the app is missing a Location Push Service Extension.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code/missingPushExtension
LocationPushServiceErrorMissingPushExtension LocationPushServiceError = 0
// LocationPushServiceErrorMissingPushServerEnvironment - An error code that indicates the app is missing an Apple Push Notification service (APNs) environment entitlement.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code/missingPushServerEnvironment
LocationPushServiceErrorMissingPushServerEnvironment LocationPushServiceError = 0
// LocationPushServiceErrorUnknown - An error code that indicates the app was unable to start the location push service for an unknown reason.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code/unknown
LocationPushServiceErrorUnknown LocationPushServiceError = 0
// LocationPushServiceErrorUnsupportedPlatform - An error code that indicates the location push service isn’t available on this platform.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code/unsupportedPlatform
LocationPushServiceErrorUnsupportedPlatform LocationPushServiceError = 0
)

// CLMonitoringState - Values that represent the current state of a monitoring condition.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState
type MonitoringState uint

const (
// MonitoringStateSatisfied - The condition is in a satisfied state.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState/CLMonitoringStateSatisfied
MonitoringStateSatisfied MonitoringState = 0
// MonitoringStateUnknown - The condition is in an unknown state.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState/CLMonitoringStateUnknown
MonitoringStateUnknown MonitoringState = 0
// MonitoringStateUnmonitored - The condition is in an unmonitored state.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState/CLMonitoringStateUnmonitored
MonitoringStateUnmonitored MonitoringState = 0
// MonitoringStateUnsatisfied - The condition is in an unsatisfied state.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState/CLMonitoringStateUnsatisfied
MonitoringStateUnsatisfied MonitoringState = 0
)

// CLProximity - Constants that reflect the relative distance to a beacon.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity
type Proximity uint

const (
// ProximityFar - The beacon is far away.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity/far
ProximityFar Proximity = 0
// ProximityImmediate - The beacon is in the user’s immediate vicinity.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity/immediate
ProximityImmediate Proximity = 0
// ProximityNear - The beacon is relatively close to the user.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity/near
ProximityNear Proximity = 0
// ProximityUnknown - The proximity of the beacon could not be determined.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity/unknown
ProximityUnknown Proximity = 0
)

// CLRegionState - Constants that reflect the relationship of the current location to the region boundaries.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegionState
type RegionState uint

const (
// RegionStateInside - The location is inside of the given region.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegionState/inside
RegionStateInside RegionState = 0
// RegionStateOutside - The location is outside of the given region.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegionState/outside
RegionStateOutside RegionState = 0
// RegionStateUnknown - It is unknown whether the location is inside or outside of the region.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegionState/unknown
RegionStateUnknown RegionState = 0
)

// CLServiceSessionAuthorizationRequirement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionAuthorizationRequirement
type ServiceSessionAuthorizationRequirement uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionAuthorizationRequirement/CLServiceSessionAuthorizationRequirementAlways
ServiceSessionAuthorizationRequirementAlways ServiceSessionAuthorizationRequirement = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionAuthorizationRequirement/CLServiceSessionAuthorizationRequirementNone
ServiceSessionAuthorizationRequirementNone ServiceSessionAuthorizationRequirement = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionAuthorizationRequirement/CLServiceSessionAuthorizationRequirementWhenInUse
ServiceSessionAuthorizationRequirementWhenInUse ServiceSessionAuthorizationRequirement = 0
)


