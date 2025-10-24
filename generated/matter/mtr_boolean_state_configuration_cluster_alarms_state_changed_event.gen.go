// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent] class.
var (
	MTRBooleanStateConfigurationClusterAlarmsStateChangedEventClass     _MTRBooleanStateConfigurationClusterAlarmsStateChangedEventClass
	MTRBooleanStateConfigurationClusterAlarmsStateChangedEventClassOnce sync.Once
)

func getMTRBooleanStateConfigurationClusterAlarmsStateChangedEventClass() _MTRBooleanStateConfigurationClusterAlarmsStateChangedEventClass {
	MTRBooleanStateConfigurationClusterAlarmsStateChangedEventClassOnce.Do(func() {
		MTRBooleanStateConfigurationClusterAlarmsStateChangedEventClass = _MTRBooleanStateConfigurationClusterAlarmsStateChangedEventClass{objc.GetClass("MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent")}
	})
	return MTRBooleanStateConfigurationClusterAlarmsStateChangedEventClass
}

type _MTRBooleanStateConfigurationClusterAlarmsStateChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent] class.
type IMTRBooleanStateConfigurationClusterAlarmsStateChangedEvent interface {
	objectivec.IObject
	// properties:
	AlarmsActive() objc.IObject /* cross-framework: NSNumber */
	SetAlarmsActive(value objc.IObject /* cross-framework: NSNumber */)
	AlarmsSuppressed() objc.IObject /* cross-framework: NSNumber */
	SetAlarmsSuppressed(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent
type MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent struct {
	objectivec.Object
}

// MTRBooleanStateConfigurationClusterAlarmsStateChangedEventFrom constructs a [MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent] from an unsafe.Pointer.
func MTRBooleanStateConfigurationClusterAlarmsStateChangedEventFrom(ptr unsafe.Pointer) MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent {
	return MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBooleanStateConfigurationClusterAlarmsStateChangedEventClass) Alloc() MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent {
	rv := objc.Send[MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBooleanStateConfigurationClusterAlarmsStateChangedEventClass) New() MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent {
	rv := objc.Send[MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent) Init() MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent {
	rv := objc.Send[MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent) Autorelease() MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent {
	rv := objc.Send[MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBooleanStateConfigurationClusterAlarmsStateChangedEvent creates a new MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent instance.
func NewMTRBooleanStateConfigurationClusterAlarmsStateChangedEvent() MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent {
	return getMTRBooleanStateConfigurationClusterAlarmsStateChangedEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclusteralarmsstatechangedevent/alarmsactive
func (m_ MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent) AlarmsActive() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("alarmsActive"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclusteralarmsstatechangedevent/alarmsactive
func (m_ MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent) SetAlarmsActive(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarmsActive:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclusteralarmsstatechangedevent/alarmssuppressed
func (m_ MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent) AlarmsSuppressed() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("alarmsSuppressed"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclusteralarmsstatechangedevent/alarmssuppressed
func (m_ MTRBooleanStateConfigurationClusterAlarmsStateChangedEvent) SetAlarmsSuppressed(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarmsSuppressed:"), value)
}



