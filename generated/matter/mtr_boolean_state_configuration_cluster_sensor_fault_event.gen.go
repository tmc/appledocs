// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBooleanStateConfigurationClusterSensorFaultEvent] class.
var (
	MTRBooleanStateConfigurationClusterSensorFaultEventClass     _MTRBooleanStateConfigurationClusterSensorFaultEventClass
	MTRBooleanStateConfigurationClusterSensorFaultEventClassOnce sync.Once
)

func getMTRBooleanStateConfigurationClusterSensorFaultEventClass() _MTRBooleanStateConfigurationClusterSensorFaultEventClass {
	MTRBooleanStateConfigurationClusterSensorFaultEventClassOnce.Do(func() {
		MTRBooleanStateConfigurationClusterSensorFaultEventClass = _MTRBooleanStateConfigurationClusterSensorFaultEventClass{objc.GetClass("MTRBooleanStateConfigurationClusterSensorFaultEvent")}
	})
	return MTRBooleanStateConfigurationClusterSensorFaultEventClass
}

type _MTRBooleanStateConfigurationClusterSensorFaultEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBooleanStateConfigurationClusterSensorFaultEvent] class.
type IMTRBooleanStateConfigurationClusterSensorFaultEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBooleanStateConfigurationClusterSensorFaultEvent
type MTRBooleanStateConfigurationClusterSensorFaultEvent struct {
	objectivec.Object
}

// MTRBooleanStateConfigurationClusterSensorFaultEventFrom constructs a [MTRBooleanStateConfigurationClusterSensorFaultEvent] from an unsafe.Pointer.
func MTRBooleanStateConfigurationClusterSensorFaultEventFrom(ptr unsafe.Pointer) MTRBooleanStateConfigurationClusterSensorFaultEvent {
	return MTRBooleanStateConfigurationClusterSensorFaultEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBooleanStateConfigurationClusterSensorFaultEventClass) Alloc() MTRBooleanStateConfigurationClusterSensorFaultEvent {
	rv := objc.Send[MTRBooleanStateConfigurationClusterSensorFaultEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBooleanStateConfigurationClusterSensorFaultEventClass) New() MTRBooleanStateConfigurationClusterSensorFaultEvent {
	rv := objc.Send[MTRBooleanStateConfigurationClusterSensorFaultEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBooleanStateConfigurationClusterSensorFaultEvent) Init() MTRBooleanStateConfigurationClusterSensorFaultEvent {
	rv := objc.Send[MTRBooleanStateConfigurationClusterSensorFaultEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBooleanStateConfigurationClusterSensorFaultEvent) Autorelease() MTRBooleanStateConfigurationClusterSensorFaultEvent {
	rv := objc.Send[MTRBooleanStateConfigurationClusterSensorFaultEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBooleanStateConfigurationClusterSensorFaultEvent creates a new MTRBooleanStateConfigurationClusterSensorFaultEvent instance.
func NewMTRBooleanStateConfigurationClusterSensorFaultEvent() MTRBooleanStateConfigurationClusterSensorFaultEvent {
	return getMTRBooleanStateConfigurationClusterSensorFaultEventClass().New()
}




