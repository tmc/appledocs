// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSwitchClusterLongPressEvent] class.
var (
	MTRSwitchClusterLongPressEventClass     _MTRSwitchClusterLongPressEventClass
	MTRSwitchClusterLongPressEventClassOnce sync.Once
)

func getMTRSwitchClusterLongPressEventClass() _MTRSwitchClusterLongPressEventClass {
	MTRSwitchClusterLongPressEventClassOnce.Do(func() {
		MTRSwitchClusterLongPressEventClass = _MTRSwitchClusterLongPressEventClass{objc.GetClass("MTRSwitchClusterLongPressEvent")}
	})
	return MTRSwitchClusterLongPressEventClass
}

type _MTRSwitchClusterLongPressEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSwitchClusterLongPressEvent] class.
type IMTRSwitchClusterLongPressEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterLongPressEvent
type MTRSwitchClusterLongPressEvent struct {
	objectivec.Object
}

// MTRSwitchClusterLongPressEventFrom constructs a [MTRSwitchClusterLongPressEvent] from an unsafe.Pointer.
func MTRSwitchClusterLongPressEventFrom(ptr unsafe.Pointer) MTRSwitchClusterLongPressEvent {
	return MTRSwitchClusterLongPressEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterLongPressEventClass) Alloc() MTRSwitchClusterLongPressEvent {
	rv := objc.Send[MTRSwitchClusterLongPressEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSwitchClusterLongPressEventClass) New() MTRSwitchClusterLongPressEvent {
	rv := objc.Send[MTRSwitchClusterLongPressEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterLongPressEvent) Init() MTRSwitchClusterLongPressEvent {
	rv := objc.Send[MTRSwitchClusterLongPressEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterLongPressEvent) Autorelease() MTRSwitchClusterLongPressEvent {
	rv := objc.Send[MTRSwitchClusterLongPressEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterLongPressEvent creates a new MTRSwitchClusterLongPressEvent instance.
func NewMTRSwitchClusterLongPressEvent() MTRSwitchClusterLongPressEvent {
	return getMTRSwitchClusterLongPressEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclusterlongpressevent/newposition
func (m_ MTRSwitchClusterLongPressEvent) NewPosition() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("newPosition"))
	return rv
}


// SetNewPosition sets the value of the newPosition property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclusterlongpressevent/newposition
func (m_ MTRSwitchClusterLongPressEvent) SetNewPosition(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewPosition:"), value)
}



