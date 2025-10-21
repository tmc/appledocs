// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessControlClusterAccessControlEntryChangedEvent] class.
var (
	MTRAccessControlClusterAccessControlEntryChangedEventClass     _MTRAccessControlClusterAccessControlEntryChangedEventClass
	MTRAccessControlClusterAccessControlEntryChangedEventClassOnce sync.Once
)

func getMTRAccessControlClusterAccessControlEntryChangedEventClass() _MTRAccessControlClusterAccessControlEntryChangedEventClass {
	MTRAccessControlClusterAccessControlEntryChangedEventClassOnce.Do(func() {
		MTRAccessControlClusterAccessControlEntryChangedEventClass = _MTRAccessControlClusterAccessControlEntryChangedEventClass{objc.GetClass("MTRAccessControlClusterAccessControlEntryChangedEvent")}
	})
	return MTRAccessControlClusterAccessControlEntryChangedEventClass
}

type _MTRAccessControlClusterAccessControlEntryChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessControlEntryChangedEvent] class.
type IMTRAccessControlClusterAccessControlEntryChangedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessControlEntryChangedEvent
type MTRAccessControlClusterAccessControlEntryChangedEvent struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessControlEntryChangedEventFrom constructs a [MTRAccessControlClusterAccessControlEntryChangedEvent] from an unsafe.Pointer.
func MTRAccessControlClusterAccessControlEntryChangedEventFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessControlEntryChangedEvent {
	return MTRAccessControlClusterAccessControlEntryChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessControlEntryChangedEventClass) Alloc() MTRAccessControlClusterAccessControlEntryChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessControlEntryChangedEventClass) New() MTRAccessControlClusterAccessControlEntryChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) Init() MTRAccessControlClusterAccessControlEntryChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) Autorelease() MTRAccessControlClusterAccessControlEntryChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessControlEntryChangedEvent creates a new MTRAccessControlClusterAccessControlEntryChangedEvent instance.
func NewMTRAccessControlClusterAccessControlEntryChangedEvent() MTRAccessControlClusterAccessControlEntryChangedEvent {
	return getMTRAccessControlClusterAccessControlEntryChangedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/adminnodeid
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) AdminNodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("adminNodeID"))
	return rv
}


// SetAdminNodeID sets the value of the adminNodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/adminnodeid
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) SetAdminNodeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminNodeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/adminpasscodeid
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) AdminPasscodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("adminPasscodeID"))
	return rv
}


// SetAdminPasscodeID sets the value of the adminPasscodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/adminpasscodeid
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) SetAdminPasscodeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminPasscodeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/changetype
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) ChangeType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("changeType"))
	return rv
}


// SetChangeType sets the value of the changeType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/changetype
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) SetChangeType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChangeType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/fabricindex
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/fabricindex
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) SetFabricIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/latestvalue
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) LatestValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("latestValue"))
	return rv
}


// SetLatestValue sets the value of the latestValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/latestvalue
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) SetLatestValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLatestValue:"), value)
}



