// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRAccessControlClusterFabricRestrictionReviewUpdateEvent] class.
var (
	MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass     _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass
	MTRAccessControlClusterFabricRestrictionReviewUpdateEventClassOnce sync.Once
)

func getMTRAccessControlClusterFabricRestrictionReviewUpdateEventClass() _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass {
	MTRAccessControlClusterFabricRestrictionReviewUpdateEventClassOnce.Do(func() {
		MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass = _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass{objc.GetClass("MTRAccessControlClusterFabricRestrictionReviewUpdateEvent")}
	})
	return MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass
}

type _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterFabricRestrictionReviewUpdateEvent] class.
type IMTRAccessControlClusterFabricRestrictionReviewUpdateEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterFabricRestrictionReviewUpdateEvent
type MTRAccessControlClusterFabricRestrictionReviewUpdateEvent struct {
	objectivec.Object
}

// MTRAccessControlClusterFabricRestrictionReviewUpdateEventFrom constructs a [MTRAccessControlClusterFabricRestrictionReviewUpdateEvent] from an unsafe.Pointer.
func MTRAccessControlClusterFabricRestrictionReviewUpdateEventFrom(ptr unsafe.Pointer) MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	return MTRAccessControlClusterFabricRestrictionReviewUpdateEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass) Alloc() MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	rv := objc.Send[MTRAccessControlClusterFabricRestrictionReviewUpdateEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterFabricRestrictionReviewUpdateEventClass) New() MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	rv := objc.Send[MTRAccessControlClusterFabricRestrictionReviewUpdateEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) Init() MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	rv := objc.Send[MTRAccessControlClusterFabricRestrictionReviewUpdateEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) Autorelease() MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	rv := objc.Send[MTRAccessControlClusterFabricRestrictionReviewUpdateEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterFabricRestrictionReviewUpdateEvent creates a new MTRAccessControlClusterFabricRestrictionReviewUpdateEvent instance.
func NewMTRAccessControlClusterFabricRestrictionReviewUpdateEvent() MTRAccessControlClusterFabricRestrictionReviewUpdateEvent {
	return getMTRAccessControlClusterFabricRestrictionReviewUpdateEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterFabricRestrictionReviewUpdateEvent/fabricIndex
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) FabricIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterFabricRestrictionReviewUpdateEvent/fabricIndex
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) SetFabricIndex(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterFabricRestrictionReviewUpdateEvent/instruction
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) Instruction() string {
	rv := objc.Send[string](m_.ID, objc.Sel("instruction"))
	return rv
}


// SetInstruction sets the value of the instruction property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterFabricRestrictionReviewUpdateEvent/instruction
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) SetInstruction(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstruction:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterFabricRestrictionReviewUpdateEvent/token
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) Token() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("token"))
	return rv
}


// SetToken sets the value of the token property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterFabricRestrictionReviewUpdateEvent/token
func (m_ MTRAccessControlClusterFabricRestrictionReviewUpdateEvent) SetToken(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setToken:"), value)
}


