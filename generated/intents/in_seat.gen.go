// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INSeat] class.
var (
	INSeatClass     _INSeatClass
	INSeatClassOnce sync.Once
)

func getINSeatClass() _INSeatClass {
	INSeatClassOnce.Do(func() {
		INSeatClass = _INSeatClass{objc.GetClass("INSeat")}
	})
	return INSeatClass
}

type _INSeatClass struct {
	class objc.Class
}

// An interface definition for the [INSeat] class.
type IINSeat interface {
	objectivec.IObject
}

// An object containing seat information associated with a reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSeat
type INSeat struct {
	objectivec.Object
}

// INSeatFrom constructs a [INSeat] from an unsafe.Pointer.
//
// An object containing seat information associated with a reservation.
func INSeatFrom(ptr unsafe.Pointer) INSeat {
	return INSeat{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INSeatClass) Alloc() INSeat {
	rv := objc.Send[INSeat](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSeatClass) New() INSeat {
	rv := objc.Send[INSeat](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSeat) Init() INSeat {
	rv := objc.Send[INSeat](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSeat) Autorelease() INSeat {
	rv := objc.Send[INSeat](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSeat creates a new INSeat instance.
func NewINSeat() INSeat {
	return getINSeatClass().New()
}




