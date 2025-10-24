// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessControlClusterAccessControlExtensionChangedEvent] class.
var (
	MTRAccessControlClusterAccessControlExtensionChangedEventClass     _MTRAccessControlClusterAccessControlExtensionChangedEventClass
	MTRAccessControlClusterAccessControlExtensionChangedEventClassOnce sync.Once
)

func getMTRAccessControlClusterAccessControlExtensionChangedEventClass() _MTRAccessControlClusterAccessControlExtensionChangedEventClass {
	MTRAccessControlClusterAccessControlExtensionChangedEventClassOnce.Do(func() {
		MTRAccessControlClusterAccessControlExtensionChangedEventClass = _MTRAccessControlClusterAccessControlExtensionChangedEventClass{objc.GetClass("MTRAccessControlClusterAccessControlExtensionChangedEvent")}
	})
	return MTRAccessControlClusterAccessControlExtensionChangedEventClass
}

type _MTRAccessControlClusterAccessControlExtensionChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessControlExtensionChangedEvent] class.
type IMTRAccessControlClusterAccessControlExtensionChangedEvent interface {
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
	LatestValue() IMTRAccessControlClusterAccessControlExtensionStruct
	SetLatestValue(value IMTRAccessControlClusterAccessControlExtensionStruct)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessControlExtensionChangedEvent
type MTRAccessControlClusterAccessControlExtensionChangedEvent struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessControlExtensionChangedEventFrom constructs a [MTRAccessControlClusterAccessControlExtensionChangedEvent] from an unsafe.Pointer.
func MTRAccessControlClusterAccessControlExtensionChangedEventFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessControlExtensionChangedEvent {
	return MTRAccessControlClusterAccessControlExtensionChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessControlExtensionChangedEventClass) Alloc() MTRAccessControlClusterAccessControlExtensionChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessControlExtensionChangedEventClass) New() MTRAccessControlClusterAccessControlExtensionChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) Init() MTRAccessControlClusterAccessControlExtensionChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) Autorelease() MTRAccessControlClusterAccessControlExtensionChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessControlExtensionChangedEvent creates a new MTRAccessControlClusterAccessControlExtensionChangedEvent instance.
func NewMTRAccessControlClusterAccessControlExtensionChangedEvent() MTRAccessControlClusterAccessControlExtensionChangedEvent {
	return getMTRAccessControlClusterAccessControlExtensionChangedEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionchangedevent/adminnodeid
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) AdminNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("adminNodeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionchangedevent/adminnodeid
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) SetAdminNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminNodeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionchangedevent/adminpasscodeid
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) AdminPasscodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("adminPasscodeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionchangedevent/adminpasscodeid
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) SetAdminPasscodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminPasscodeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionchangedevent/changetype
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) ChangeType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("changeType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionchangedevent/changetype
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) SetChangeType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChangeType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionchangedevent/fabricindex
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionchangedevent/fabricindex
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionchangedevent/latestvalue
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) LatestValue() IMTRAccessControlClusterAccessControlExtensionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionStruct](m_.ID, objc.Sel("latestValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionchangedevent/latestvalue
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) SetLatestValue(value IMTRAccessControlClusterAccessControlExtensionStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLatestValue:"), value)
}



