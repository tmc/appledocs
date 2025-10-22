// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Update] class.
var (
	UpdateClass     _UpdateClass
	UpdateClassOnce sync.Once
)

func getUpdateClass() _UpdateClass {
	UpdateClassOnce.Do(func() {
		UpdateClass = _UpdateClass{objc.GetClass("CLUpdate")}
	})
	return UpdateClass
}

type _UpdateClass struct {
	class objc.Class
}

// An interface definition for the [Update] class.
type IUpdate interface {
	objectivec.IObject
	AccuracyLimited() bool
	AuthorizationDenied() bool
	AuthorizationDeniedGlobally() bool
	AuthorizationRequestInProgress() bool
	AuthorizationRestricted() bool
	InsufficientlyInUse() bool
	IsStationary() bool
	Location() CLLocation
	LocationUnavailable() bool
	ServiceSessionRequired() bool
	Stationary() bool
}

// An object that represents a location update.


// An object that represents a location update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate

type Update struct {
	objectivec.Object
}

// UpdateFrom constructs a [Update] from an unsafe.Pointer.
//
// An object that represents a location update.
func UpdateFrom(ptr unsafe.Pointer) Update {
	return Update{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UpdateClass) Alloc() Update {
	rv := objc.Send[Update](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UpdateClass) New() Update {
	rv := objc.Send[Update](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ Update) Init() Update {
	rv := objc.Send[Update](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ Update) Autorelease() Update {
	rv := objc.Send[Update](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUpdate creates a new Update instance.
func NewUpdate() Update {
	return getUpdateClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate/accuracyLimited

func (u_ Update) AccuracyLimited() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("accuracyLimited"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate/authorizationDenied

func (u_ Update) AuthorizationDenied() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("authorizationDenied"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate/authorizationDeniedGlobally

func (u_ Update) AuthorizationDeniedGlobally() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("authorizationDeniedGlobally"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate/authorizationRequestInProgress

func (u_ Update) AuthorizationRequestInProgress() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("authorizationRequestInProgress"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate/authorizationRestricted

func (u_ Update) AuthorizationRestricted() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("authorizationRestricted"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate/insufficientlyInUse

func (u_ Update) InsufficientlyInUse() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("insufficientlyInUse"))
	return rv
}


// A Boolean value that indicates whether the device is stationary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate/isStationary

func (u_ Update) IsStationary() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isStationary"))
	return rv
}


// A person’s location, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate/location

func (u_ Update) Location() CLLocation {
	rv := objc.Send[CLLocation](u_.ID, objc.Sel("location"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate/locationUnavailable

func (u_ Update) LocationUnavailable() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("locationUnavailable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate/serviceSessionRequired

func (u_ Update) ServiceSessionRequired() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("serviceSessionRequired"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate/stationary

func (u_ Update) Stationary() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("stationary"))
	return rv
}



