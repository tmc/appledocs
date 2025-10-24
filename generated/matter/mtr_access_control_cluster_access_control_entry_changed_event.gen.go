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
	// properties:
	AdminNodeID() objc.IObject /* cross-framework: NSNumber */
	SetAdminNodeID(value objc.IObject /* cross-framework: NSNumber */)
	AdminPasscodeID() objc.IObject /* cross-framework: NSNumber */
	SetAdminPasscodeID(value objc.IObject /* cross-framework: NSNumber */)
	ChangeType() objc.IObject /* cross-framework: NSNumber */
	SetChangeType(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	LatestValue() IMTRAccessControlClusterAccessControlEntryStruct
	SetLatestValue(value IMTRAccessControlClusterAccessControlEntryStruct)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/adminnodeid
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) AdminNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("adminNodeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/adminnodeid
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) SetAdminNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminNodeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/adminpasscodeid
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) AdminPasscodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("adminPasscodeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/adminpasscodeid
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) SetAdminPasscodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminPasscodeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/changetype
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) ChangeType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("changeType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/changetype
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) SetChangeType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChangeType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/fabricindex
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/fabricindex
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/latestvalue
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) LatestValue() IMTRAccessControlClusterAccessControlEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryStruct](m_.ID, objc.Sel("latestValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrychangedevent/latestvalue
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) SetLatestValue(value IMTRAccessControlClusterAccessControlEntryStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLatestValue:"), value)
}



