// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRValveConfigurationAndControlClusterValveStateChangedEvent] class.
var (
	MTRValveConfigurationAndControlClusterValveStateChangedEventClass     _MTRValveConfigurationAndControlClusterValveStateChangedEventClass
	MTRValveConfigurationAndControlClusterValveStateChangedEventClassOnce sync.Once
)

func getMTRValveConfigurationAndControlClusterValveStateChangedEventClass() _MTRValveConfigurationAndControlClusterValveStateChangedEventClass {
	MTRValveConfigurationAndControlClusterValveStateChangedEventClassOnce.Do(func() {
		MTRValveConfigurationAndControlClusterValveStateChangedEventClass = _MTRValveConfigurationAndControlClusterValveStateChangedEventClass{objc.GetClass("MTRValveConfigurationAndControlClusterValveStateChangedEvent")}
	})
	return MTRValveConfigurationAndControlClusterValveStateChangedEventClass
}

type _MTRValveConfigurationAndControlClusterValveStateChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRValveConfigurationAndControlClusterValveStateChangedEvent] class.
type IMTRValveConfigurationAndControlClusterValveStateChangedEvent interface {
	objectivec.IObject
	ValveLevel() foundation.Number
	SetValveLevel(value foundation.INumber)
	ValveState() foundation.Number
	SetValveState(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRValveConfigurationAndControlClusterValveStateChangedEvent
type MTRValveConfigurationAndControlClusterValveStateChangedEvent struct {
	objectivec.Object
}

// MTRValveConfigurationAndControlClusterValveStateChangedEventFrom constructs a [MTRValveConfigurationAndControlClusterValveStateChangedEvent] from an unsafe.Pointer.
func MTRValveConfigurationAndControlClusterValveStateChangedEventFrom(ptr unsafe.Pointer) MTRValveConfigurationAndControlClusterValveStateChangedEvent {
	return MTRValveConfigurationAndControlClusterValveStateChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRValveConfigurationAndControlClusterValveStateChangedEventClass) Alloc() MTRValveConfigurationAndControlClusterValveStateChangedEvent {
	rv := objc.Send[MTRValveConfigurationAndControlClusterValveStateChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRValveConfigurationAndControlClusterValveStateChangedEventClass) New() MTRValveConfigurationAndControlClusterValveStateChangedEvent {
	rv := objc.Send[MTRValveConfigurationAndControlClusterValveStateChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRValveConfigurationAndControlClusterValveStateChangedEvent) Init() MTRValveConfigurationAndControlClusterValveStateChangedEvent {
	rv := objc.Send[MTRValveConfigurationAndControlClusterValveStateChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRValveConfigurationAndControlClusterValveStateChangedEvent) Autorelease() MTRValveConfigurationAndControlClusterValveStateChangedEvent {
	rv := objc.Send[MTRValveConfigurationAndControlClusterValveStateChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRValveConfigurationAndControlClusterValveStateChangedEvent creates a new MTRValveConfigurationAndControlClusterValveStateChangedEvent instance.
func NewMTRValveConfigurationAndControlClusterValveStateChangedEvent() MTRValveConfigurationAndControlClusterValveStateChangedEvent {
	return getMTRValveConfigurationAndControlClusterValveStateChangedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustervalvestatechangedevent/valvelevel
func (m_ MTRValveConfigurationAndControlClusterValveStateChangedEvent) ValveLevel() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("valveLevel"))
	return rv
}


// SetValveLevel sets the value of the valveLevel property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustervalvestatechangedevent/valvelevel
func (m_ MTRValveConfigurationAndControlClusterValveStateChangedEvent) SetValveLevel(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValveLevel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustervalvestatechangedevent/valvestate
func (m_ MTRValveConfigurationAndControlClusterValveStateChangedEvent) ValveState() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("valveState"))
	return rv
}


// SetValveState sets the value of the valveState property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustervalvestatechangedevent/valvestate
func (m_ MTRValveConfigurationAndControlClusterValveStateChangedEvent) SetValveState(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValveState:"), value)
}



