// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	ValveLevel() objc.IObject /* cross-framework: NSNumber */
	SetValveLevel(value objc.IObject /* cross-framework: NSNumber */)
	ValveState() objc.IObject /* cross-framework: NSNumber */
	SetValveState(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustervalvestatechangedevent/valvelevel
func (m_ MTRValveConfigurationAndControlClusterValveStateChangedEvent) ValveLevel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("valveLevel"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustervalvestatechangedevent/valvelevel
func (m_ MTRValveConfigurationAndControlClusterValveStateChangedEvent) SetValveLevel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValveLevel:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustervalvestatechangedevent/valvestate
func (m_ MTRValveConfigurationAndControlClusterValveStateChangedEvent) ValveState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("valveState"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustervalvestatechangedevent/valvestate
func (m_ MTRValveConfigurationAndControlClusterValveStateChangedEvent) SetValveState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValveState:"), value)
}
