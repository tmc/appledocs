// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INReservationAction] class.
var (
	INReservationActionClass     _INReservationActionClass
	INReservationActionClassOnce sync.Once
)

func getINReservationActionClass() _INReservationActionClass {
	INReservationActionClassOnce.Do(func() {
		INReservationActionClass = _INReservationActionClass{objc.GetClass("INReservationAction")}
	})
	return INReservationActionClass
}

type _INReservationActionClass struct {
	class objc.Class
}

// An interface definition for the [INReservationAction] class.
type IINReservationAction interface {
	objectivec.IObject
}

// An action a user can perform that’s relevant to a reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INReservationAction
type INReservationAction struct {
	objectivec.Object
}

// INReservationActionFrom constructs a [INReservationAction] from an unsafe.Pointer.
//
// An action a user can perform that’s relevant to a reservation.
func INReservationActionFrom(ptr unsafe.Pointer) INReservationAction {
	return INReservationAction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INReservationActionClass) Alloc() INReservationAction {
	rv := objc.Send[INReservationAction](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INReservationActionClass) New() INReservationAction {
	rv := objc.Send[INReservationAction](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INReservationAction) Init() INReservationAction {
	rv := objc.Send[INReservationAction](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INReservationAction) Autorelease() INReservationAction {
	rv := objc.Send[INReservationAction](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINReservationAction creates a new INReservationAction instance.
func NewINReservationAction() INReservationAction {
	return getINReservationActionClass().New()
}


// The date and time range that the action is valid.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inreservationaction/validduration
func (i_ INReservationAction) ValidDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("validDuration"))
	return rv
}


// SetValidDuration sets the value of the validDuration property.
// The date and time range that the action is valid.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inreservationaction/validduration
func (i_ INReservationAction) SetValidDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setValidDuration:"), value)
}

// The user activity object used when launching your app.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inreservationaction/useractivity
func (i_ INReservationAction) UserActivity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("userActivity"))
	return rv
}


// SetUserActivity sets the value of the userActivity property.
// The user activity object used when launching your app.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inreservationaction/useractivity
func (i_ INReservationAction) SetUserActivity(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUserActivity:"), value)
}

// The type of action for the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inreservationaction/type
func (i_ INReservationAction) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The type of action for the reservation.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inreservationaction/type
func (i_ INReservationAction) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setType:"), value)
}



