// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INReservation] class.
var (
	INReservationClass     _INReservationClass
	INReservationClassOnce sync.Once
)

func getINReservationClass() _INReservationClass {
	INReservationClassOnce.Do(func() {
		INReservationClass = _INReservationClass{objc.GetClass("INReservation")}
	})
	return INReservationClass
}

type _INReservationClass struct {
	class objc.Class
}

// An interface definition for the [INReservation] class.
type IINReservation interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other Intents classes.

// A parent class referenced by other Intents classes. [Full Topic]
type INReservation struct {
	objectivec.Object
}

// INReservationFrom constructs a [INReservation] from an unsafe.Pointer.
//
// A parent class referenced by other Intents classes.
func INReservationFrom(ptr unsafe.Pointer) INReservation {
	return INReservation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INReservationClass) Alloc() INReservation {
	rv := objc.Send[INReservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INReservationClass) New() INReservation {
	rv := objc.Send[INReservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INReservation) Init() INReservation {
	rv := objc.Send[INReservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INReservation) Autorelease() INReservation {
	rv := objc.Send[INReservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINReservation creates a new INReservation instance.
func NewINReservation() INReservation {
	return getINReservationClass().New()
}
