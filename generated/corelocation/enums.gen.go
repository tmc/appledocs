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

// CLActivityType - Constants that indicate the type of activity associated with location updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType
type CLActivityType uint

const (
	// CLActivityTypeAutomotiveNavigation - The value that indicates positioning in an automobile following a road network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLActivityType/automotiveNavigation
	CLActivityTypeAutomotiveNavigation CLActivityType = 0
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

// CLDeviceOrientation - Constants indicating the physical orientation of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLDeviceOrientation
type CLDeviceOrientation uint

// CLError - Error codes returned by the location manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code
type CLError uint

const (
	// kCLErrorNetwork - A constant that indicates the network was unavailable or a network error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code/network
	kCLErrorNetwork CLError = 0
)

// CLLiveUpdateConfiguration - Specifies the types of locations that a location updater generates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLiveUpdateConfiguration
type CLLiveUpdateConfiguration int

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

// CLProximity - Constants that reflect the relative distance to a beacon.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity
type CLProximity uint

const (
	// CLProximityUnknown - The proximity of the beacon could not be determined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLProximity/unknown
	CLProximityUnknown CLProximity = 0
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
