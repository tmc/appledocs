// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSmokeCOAlarmClusterAlarmMutedEvent] class.
var (
	MTRSmokeCOAlarmClusterAlarmMutedEventClass     _MTRSmokeCOAlarmClusterAlarmMutedEventClass
	MTRSmokeCOAlarmClusterAlarmMutedEventClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterAlarmMutedEventClass() _MTRSmokeCOAlarmClusterAlarmMutedEventClass {
	MTRSmokeCOAlarmClusterAlarmMutedEventClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterAlarmMutedEventClass = _MTRSmokeCOAlarmClusterAlarmMutedEventClass{objc.GetClass("MTRSmokeCOAlarmClusterAlarmMutedEvent")}
	})
	return MTRSmokeCOAlarmClusterAlarmMutedEventClass
}

type _MTRSmokeCOAlarmClusterAlarmMutedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterAlarmMutedEvent] class.
type IMTRSmokeCOAlarmClusterAlarmMutedEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterAlarmMutedEvent
type MTRSmokeCOAlarmClusterAlarmMutedEvent struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterAlarmMutedEventFrom constructs a [MTRSmokeCOAlarmClusterAlarmMutedEvent] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterAlarmMutedEventFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterAlarmMutedEvent {
	return MTRSmokeCOAlarmClusterAlarmMutedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterAlarmMutedEventClass) Alloc() MTRSmokeCOAlarmClusterAlarmMutedEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterAlarmMutedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterAlarmMutedEventClass) New() MTRSmokeCOAlarmClusterAlarmMutedEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterAlarmMutedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterAlarmMutedEvent) Init() MTRSmokeCOAlarmClusterAlarmMutedEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterAlarmMutedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterAlarmMutedEvent) Autorelease() MTRSmokeCOAlarmClusterAlarmMutedEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterAlarmMutedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterAlarmMutedEvent creates a new MTRSmokeCOAlarmClusterAlarmMutedEvent instance.
func NewMTRSmokeCOAlarmClusterAlarmMutedEvent() MTRSmokeCOAlarmClusterAlarmMutedEvent {
	return getMTRSmokeCOAlarmClusterAlarmMutedEventClass().New()
}
