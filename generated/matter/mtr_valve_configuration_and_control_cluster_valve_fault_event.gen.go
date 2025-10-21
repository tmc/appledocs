// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRValveConfigurationAndControlClusterValveFaultEvent] class.
var (
	MTRValveConfigurationAndControlClusterValveFaultEventClass     _MTRValveConfigurationAndControlClusterValveFaultEventClass
	MTRValveConfigurationAndControlClusterValveFaultEventClassOnce sync.Once
)

func getMTRValveConfigurationAndControlClusterValveFaultEventClass() _MTRValveConfigurationAndControlClusterValveFaultEventClass {
	MTRValveConfigurationAndControlClusterValveFaultEventClassOnce.Do(func() {
		MTRValveConfigurationAndControlClusterValveFaultEventClass = _MTRValveConfigurationAndControlClusterValveFaultEventClass{objc.GetClass("MTRValveConfigurationAndControlClusterValveFaultEvent")}
	})
	return MTRValveConfigurationAndControlClusterValveFaultEventClass
}

type _MTRValveConfigurationAndControlClusterValveFaultEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRValveConfigurationAndControlClusterValveFaultEvent] class.
type IMTRValveConfigurationAndControlClusterValveFaultEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRValveConfigurationAndControlClusterValveFaultEvent
type MTRValveConfigurationAndControlClusterValveFaultEvent struct {
	objectivec.Object
}

// MTRValveConfigurationAndControlClusterValveFaultEventFrom constructs a [MTRValveConfigurationAndControlClusterValveFaultEvent] from an unsafe.Pointer.
func MTRValveConfigurationAndControlClusterValveFaultEventFrom(ptr unsafe.Pointer) MTRValveConfigurationAndControlClusterValveFaultEvent {
	return MTRValveConfigurationAndControlClusterValveFaultEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRValveConfigurationAndControlClusterValveFaultEventClass) Alloc() MTRValveConfigurationAndControlClusterValveFaultEvent {
	rv := objc.Send[MTRValveConfigurationAndControlClusterValveFaultEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRValveConfigurationAndControlClusterValveFaultEventClass) New() MTRValveConfigurationAndControlClusterValveFaultEvent {
	rv := objc.Send[MTRValveConfigurationAndControlClusterValveFaultEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRValveConfigurationAndControlClusterValveFaultEvent) Init() MTRValveConfigurationAndControlClusterValveFaultEvent {
	rv := objc.Send[MTRValveConfigurationAndControlClusterValveFaultEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRValveConfigurationAndControlClusterValveFaultEvent) Autorelease() MTRValveConfigurationAndControlClusterValveFaultEvent {
	rv := objc.Send[MTRValveConfigurationAndControlClusterValveFaultEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRValveConfigurationAndControlClusterValveFaultEvent creates a new MTRValveConfigurationAndControlClusterValveFaultEvent instance.
func NewMTRValveConfigurationAndControlClusterValveFaultEvent() MTRValveConfigurationAndControlClusterValveFaultEvent {
	return getMTRValveConfigurationAndControlClusterValveFaultEventClass().New()
}




