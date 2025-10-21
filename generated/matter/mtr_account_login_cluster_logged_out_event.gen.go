// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccountLoginClusterLoggedOutEvent] class.
var (
	MTRAccountLoginClusterLoggedOutEventClass     _MTRAccountLoginClusterLoggedOutEventClass
	MTRAccountLoginClusterLoggedOutEventClassOnce sync.Once
)

func getMTRAccountLoginClusterLoggedOutEventClass() _MTRAccountLoginClusterLoggedOutEventClass {
	MTRAccountLoginClusterLoggedOutEventClassOnce.Do(func() {
		MTRAccountLoginClusterLoggedOutEventClass = _MTRAccountLoginClusterLoggedOutEventClass{objc.GetClass("MTRAccountLoginClusterLoggedOutEvent")}
	})
	return MTRAccountLoginClusterLoggedOutEventClass
}

type _MTRAccountLoginClusterLoggedOutEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccountLoginClusterLoggedOutEvent] class.
type IMTRAccountLoginClusterLoggedOutEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterLoggedOutEvent
type MTRAccountLoginClusterLoggedOutEvent struct {
	objectivec.Object
}

// MTRAccountLoginClusterLoggedOutEventFrom constructs a [MTRAccountLoginClusterLoggedOutEvent] from an unsafe.Pointer.
func MTRAccountLoginClusterLoggedOutEventFrom(ptr unsafe.Pointer) MTRAccountLoginClusterLoggedOutEvent {
	return MTRAccountLoginClusterLoggedOutEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccountLoginClusterLoggedOutEventClass) Alloc() MTRAccountLoginClusterLoggedOutEvent {
	rv := objc.Send[MTRAccountLoginClusterLoggedOutEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccountLoginClusterLoggedOutEventClass) New() MTRAccountLoginClusterLoggedOutEvent {
	rv := objc.Send[MTRAccountLoginClusterLoggedOutEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccountLoginClusterLoggedOutEvent) Init() MTRAccountLoginClusterLoggedOutEvent {
	rv := objc.Send[MTRAccountLoginClusterLoggedOutEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccountLoginClusterLoggedOutEvent) Autorelease() MTRAccountLoginClusterLoggedOutEvent {
	rv := objc.Send[MTRAccountLoginClusterLoggedOutEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccountLoginClusterLoggedOutEvent creates a new MTRAccountLoginClusterLoggedOutEvent instance.
func NewMTRAccountLoginClusterLoggedOutEvent() MTRAccountLoginClusterLoggedOutEvent {
	return getMTRAccountLoginClusterLoggedOutEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterLoggedOutEvent/node
func (m_ MTRAccountLoginClusterLoggedOutEvent) Node() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("node"))
	return rv
}


// SetNode sets the value of the node property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterLoggedOutEvent/node
func (m_ MTRAccountLoginClusterLoggedOutEvent) SetNode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNode:"), value)
}



