// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBasicClusterLeaveEvent] class.
var (
	MTRBasicClusterLeaveEventClass     _MTRBasicClusterLeaveEventClass
	MTRBasicClusterLeaveEventClassOnce sync.Once
)

func getMTRBasicClusterLeaveEventClass() _MTRBasicClusterLeaveEventClass {
	MTRBasicClusterLeaveEventClassOnce.Do(func() {
		MTRBasicClusterLeaveEventClass = _MTRBasicClusterLeaveEventClass{objc.GetClass("MTRBasicClusterLeaveEvent")}
	})
	return MTRBasicClusterLeaveEventClass
}

type _MTRBasicClusterLeaveEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicClusterLeaveEvent] class.
type IMTRBasicClusterLeaveEvent interface {
	IMTRBasicInformationClusterLeaveEvent
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicClusterLeaveEvent
type MTRBasicClusterLeaveEvent struct {
	MTRBasicInformationClusterLeaveEvent
}

// MTRBasicClusterLeaveEventFrom constructs a [MTRBasicClusterLeaveEvent] from an unsafe.Pointer.
func MTRBasicClusterLeaveEventFrom(ptr unsafe.Pointer) MTRBasicClusterLeaveEvent {
	return MTRBasicClusterLeaveEvent{
		MTRBasicInformationClusterLeaveEvent: MTRBasicInformationClusterLeaveEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicClusterLeaveEventClass) Alloc() MTRBasicClusterLeaveEvent {
	rv := objc.Send[MTRBasicClusterLeaveEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicClusterLeaveEventClass) New() MTRBasicClusterLeaveEvent {
	rv := objc.Send[MTRBasicClusterLeaveEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicClusterLeaveEvent) Init() MTRBasicClusterLeaveEvent {
	rv := objc.Send[MTRBasicClusterLeaveEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicClusterLeaveEvent) Autorelease() MTRBasicClusterLeaveEvent {
	rv := objc.Send[MTRBasicClusterLeaveEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicClusterLeaveEvent creates a new MTRBasicClusterLeaveEvent instance.
func NewMTRBasicClusterLeaveEvent() MTRBasicClusterLeaveEvent {
	return getMTRBasicClusterLeaveEventClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasicclusterleaveevent/fabricindex
func (m_ MTRBasicClusterLeaveEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasicclusterleaveevent/fabricindex
func (m_ MTRBasicClusterLeaveEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}
