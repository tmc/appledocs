// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INRideStatus] class.
var (
	INRideStatusClass     _INRideStatusClass
	INRideStatusClassOnce sync.Once
)

func getINRideStatusClass() _INRideStatusClass {
	INRideStatusClassOnce.Do(func() {
		INRideStatusClass = _INRideStatusClass{objc.GetClass("INRideStatus")}
	})
	return INRideStatusClass
}

type _INRideStatusClass struct {
	class objc.Class
}

// An interface definition for the [INRideStatus] class.
type IINRideStatus interface {
	objectivec.IObject
	CompletionStatus() unsafe.Pointer
	SetCompletionStatus(value unsafe.Pointer)
	Driver() unsafe.Pointer
	SetDriver(value unsafe.Pointer)
	Phase() INRidePhase
	SetPhase(value INRidePhase)
	RideIdentifier() string
	SetRideIdentifier(value string)
	RideOption() unsafe.Pointer
	SetRideOption(value unsafe.Pointer)
	UserActivityForCancelingInApplication() foundation.UserActivity
	SetUserActivityForCancelingInApplication(value foundation.IUserActivity)
	Vehicle() INRideVehicle
	SetVehicle(value INRideVehicle)
	AdditionalActionActivities() foundation.UserActivity
	SetAdditionalActionActivities(value foundation.IUserActivity)
	DropOffLocation() corelocation.Placemark
	SetDropOffLocation(value corelocation.IPlacemark)
	EstimatedDropOffDate() foundation.Date
	SetEstimatedDropOffDate(value foundation.IDate)
	EstimatedPickupDate() foundation.Date
	SetEstimatedPickupDate(value foundation.IDate)
	EstimatedPickupEndDate() foundation.Date
	SetEstimatedPickupEndDate(value foundation.IDate)
	PickupLocation() corelocation.Placemark
	SetPickupLocation(value corelocation.IPlacemark)
	ScheduledPickupTime() INDateComponentsRange
	SetScheduledPickupTime(value INDateComponentsRange)
	Waypoints() corelocation.Placemark
	SetWaypoints(value corelocation.IPlacemark)
}

// The status of a ride booked through a ride-booking service.
//
// When the user books a ride or requests the status of a ride, you create an object as part of your response and fill it with the relevant information. A ride status object conveys information about the current status of a ride, such as whether the ride is ongoing, completed, or has yet to occur. This object also contains details about the ride such as the pickup location, drop-off location, and information about the driver and vehicle. Siri and Maps display the information in this object to the user at appropriate times. When configuring ride status objects, always provide values for as many properties as possible. Siri and Maps display almost all of the information that you provide to the user, so it is good to offer as many details as you can.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus
type INRideStatus struct {
	objectivec.Object
}

// INRideStatusFrom constructs a [INRideStatus] from an unsafe.Pointer.
//
// The status of a ride booked through a ride-booking service.
func INRideStatusFrom(ptr unsafe.Pointer) INRideStatus {
	return INRideStatus{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INRideStatusClass) Alloc() INRideStatus {
	rv := objc.Send[INRideStatus](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INRideStatusClass) New() INRideStatus {
	rv := objc.Send[INRideStatus](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INRideStatus) Init() INRideStatus {
	rv := objc.Send[INRideStatus](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INRideStatus) Autorelease() INRideStatus {
	rv := objc.Send[INRideStatus](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINRideStatus creates a new INRideStatus instance.
func NewINRideStatus() INRideStatus {
	return getINRideStatusClass().New()
}


// Information about how the ride ended.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/completionStatus
func (i_ INRideStatus) CompletionStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("completionStatus"))
	return rv
}


// SetCompletionStatus sets the value of the completionStatus property.
// Information about how the ride ended.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/completionStatus
func (i_ INRideStatus) SetCompletionStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCompletionStatus:"), value)
}

// The driver providing the ride.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/driver
func (i_ INRideStatus) Driver() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("driver"))
	return rv
}


// SetDriver sets the value of the driver property.
// The driver providing the ride.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/driver
func (i_ INRideStatus) SetDriver(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDriver:"), value)
}

// The current status of the ride.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/phase
func (i_ INRideStatus) Phase() INRidePhase {
	rv := objc.Send[INRidePhase](i_.ID, objc.Sel("phase"))
	return rv
}


// SetPhase sets the value of the phase property.
// The current status of the ride.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/phase
func (i_ INRideStatus) SetPhase(value INRidePhase) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPhase:"), value)
}

// The unique string that you use to identify the ride.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/rideIdentifier
func (i_ INRideStatus) RideIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("rideIdentifier"))
	return rv
}


// SetRideIdentifier sets the value of the rideIdentifier property.
// The unique string that you use to identify the ride.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/rideIdentifier
func (i_ INRideStatus) SetRideIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRideIdentifier:"), objc.String(value))
}

// Information about the type of ride that you are offering to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/rideOption
func (i_ INRideStatus) RideOption() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("rideOption"))
	return rv
}


// SetRideOption sets the value of the rideOption property.
// Information about the type of ride that you are offering to the user.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/rideOption
func (i_ INRideStatus) SetRideOption(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRideOption:"), value)
}

// A user activity object for canceling the ride request.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/userActivityForCancelingInApplication
func (i_ INRideStatus) UserActivityForCancelingInApplication() foundation.UserActivity {
	rv := objc.Send[foundation.UserActivity](i_.ID, objc.Sel("userActivityForCancelingInApplication"))
	return rv
}


// SetUserActivityForCancelingInApplication sets the value of the userActivityForCancelingInApplication property.
// A user activity object for canceling the ride request.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/userActivityForCancelingInApplication
func (i_ INRideStatus) SetUserActivityForCancelingInApplication(value foundation.IUserActivity) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUserActivityForCancelingInApplication:"), value)
}

// The vehicle assigned to pick up the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/vehicle
func (i_ INRideStatus) Vehicle() INRideVehicle {
	rv := objc.Send[INRideVehicle](i_.ID, objc.Sel("vehicle"))
	return rv
}


// SetVehicle sets the value of the vehicle property.
// The vehicle assigned to pick up the user.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideStatus/vehicle
func (i_ INRideStatus) SetVehicle(value INRideVehicle) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setVehicle:"), value)
}

// Additional activities that the user may choose to perform while the ride request is active.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/additionalactionactivities
func (i_ INRideStatus) AdditionalActionActivities() foundation.UserActivity {
	rv := objc.Send[foundation.UserActivity](i_.ID, objc.Sel("additionalActionActivities"))
	return rv
}


// SetAdditionalActionActivities sets the value of the additionalActionActivities property.
// Additional activities that the user may choose to perform while the ride request is active.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/additionalactionactivities
func (i_ INRideStatus) SetAdditionalActionActivities(value foundation.IUserActivity) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAdditionalActionActivities:"), value)
}

// The destination of the user.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/dropofflocation
func (i_ INRideStatus) DropOffLocation() corelocation.Placemark {
	rv := objc.Send[corelocation.Placemark](i_.ID, objc.Sel("dropOffLocation"))
	return rv
}


// SetDropOffLocation sets the value of the dropOffLocation property.
// The destination of the user.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/dropofflocation
func (i_ INRideStatus) SetDropOffLocation(value corelocation.IPlacemark) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDropOffLocation:"), value)
}

// The estimated drop-off time for the ride.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/estimateddropoffdate
func (i_ INRideStatus) EstimatedDropOffDate() foundation.Date {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("estimatedDropOffDate"))
	return rv
}


// SetEstimatedDropOffDate sets the value of the estimatedDropOffDate property.
// The estimated drop-off time for the ride.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/estimateddropoffdate
func (i_ INRideStatus) SetEstimatedDropOffDate(value foundation.IDate) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEstimatedDropOffDate:"), value)
}

// The estimated pickup time for the ride.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/estimatedpickupdate
func (i_ INRideStatus) EstimatedPickupDate() foundation.Date {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("estimatedPickupDate"))
	return rv
}


// SetEstimatedPickupDate sets the value of the estimatedPickupDate property.
// The estimated pickup time for the ride.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/estimatedpickupdate
func (i_ INRideStatus) SetEstimatedPickupDate(value foundation.IDate) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEstimatedPickupDate:"), value)
}

// The date after which the driver may stop waiting for the user.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/estimatedpickupenddate
func (i_ INRideStatus) EstimatedPickupEndDate() foundation.Date {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("estimatedPickupEndDate"))
	return rv
}


// SetEstimatedPickupEndDate sets the value of the estimatedPickupEndDate property.
// The date after which the driver may stop waiting for the user.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/estimatedpickupenddate
func (i_ INRideStatus) SetEstimatedPickupEndDate(value foundation.IDate) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEstimatedPickupEndDate:"), value)
}

// The starting location for the ride.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/pickuplocation
func (i_ INRideStatus) PickupLocation() corelocation.Placemark {
	rv := objc.Send[corelocation.Placemark](i_.ID, objc.Sel("pickupLocation"))
	return rv
}


// SetPickupLocation sets the value of the pickupLocation property.
// The starting location for the ride.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/pickuplocation
func (i_ INRideStatus) SetPickupLocation(value corelocation.IPlacemark) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPickupLocation:"), value)
}

// The scheduled pickup time for the user.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/scheduledpickuptime
func (i_ INRideStatus) ScheduledPickupTime() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("scheduledPickupTime"))
	return rv
}


// SetScheduledPickupTime sets the value of the scheduledPickupTime property.
// The scheduled pickup time for the user.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/scheduledpickuptime
func (i_ INRideStatus) SetScheduledPickupTime(value INDateComponentsRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScheduledPickupTime:"), value)
}

// The intermediate waypoints for the journey.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/waypoints
func (i_ INRideStatus) Waypoints() corelocation.Placemark {
	rv := objc.Send[corelocation.Placemark](i_.ID, objc.Sel("waypoints"))
	return rv
}


// SetWaypoints sets the value of the waypoints property.
// The intermediate waypoints for the journey.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/waypoints
func (i_ INRideStatus) SetWaypoints(value corelocation.IPlacemark) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWaypoints:"), value)
}



