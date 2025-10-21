// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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




