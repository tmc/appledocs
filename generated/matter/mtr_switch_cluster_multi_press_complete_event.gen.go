// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSwitchClusterMultiPressCompleteEvent] class.
var (
	MTRSwitchClusterMultiPressCompleteEventClass     _MTRSwitchClusterMultiPressCompleteEventClass
	MTRSwitchClusterMultiPressCompleteEventClassOnce sync.Once
)

func getMTRSwitchClusterMultiPressCompleteEventClass() _MTRSwitchClusterMultiPressCompleteEventClass {
	MTRSwitchClusterMultiPressCompleteEventClassOnce.Do(func() {
		MTRSwitchClusterMultiPressCompleteEventClass = _MTRSwitchClusterMultiPressCompleteEventClass{objc.GetClass("MTRSwitchClusterMultiPressCompleteEvent")}
	})
	return MTRSwitchClusterMultiPressCompleteEventClass
}

type _MTRSwitchClusterMultiPressCompleteEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSwitchClusterMultiPressCompleteEvent] class.
type IMTRSwitchClusterMultiPressCompleteEvent interface {
	objectivec.IObject
	// properties:
	NewPosition() objc.IObject /* cross-framework: NSNumber */
	SetNewPosition(value objc.IObject /* cross-framework: NSNumber */)
	PreviousPosition() objc.IObject /* cross-framework: NSNumber */
	SetPreviousPosition(value objc.IObject /* cross-framework: NSNumber */)
	TotalNumberOfPressesCounted() objc.IObject /* cross-framework: NSNumber */
	SetTotalNumberOfPressesCounted(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressCompleteEvent
type MTRSwitchClusterMultiPressCompleteEvent struct {
	objectivec.Object
}

// MTRSwitchClusterMultiPressCompleteEventFrom constructs a [MTRSwitchClusterMultiPressCompleteEvent] from an unsafe.Pointer.
func MTRSwitchClusterMultiPressCompleteEventFrom(ptr unsafe.Pointer) MTRSwitchClusterMultiPressCompleteEvent {
	return MTRSwitchClusterMultiPressCompleteEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterMultiPressCompleteEventClass) Alloc() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSwitchClusterMultiPressCompleteEventClass) New() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterMultiPressCompleteEvent) Init() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterMultiPressCompleteEvent) Autorelease() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterMultiPressCompleteEvent creates a new MTRSwitchClusterMultiPressCompleteEvent instance.
func NewMTRSwitchClusterMultiPressCompleteEvent() MTRSwitchClusterMultiPressCompleteEvent {
	return getMTRSwitchClusterMultiPressCompleteEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclustermultipresscompleteevent/newposition
func (m_ MTRSwitchClusterMultiPressCompleteEvent) NewPosition() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newPosition"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclustermultipresscompleteevent/newposition
func (m_ MTRSwitchClusterMultiPressCompleteEvent) SetNewPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewPosition:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclustermultipresscompleteevent/previousposition
func (m_ MTRSwitchClusterMultiPressCompleteEvent) PreviousPosition() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("previousPosition"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclustermultipresscompleteevent/previousposition
func (m_ MTRSwitchClusterMultiPressCompleteEvent) SetPreviousPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousPosition:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclustermultipresscompleteevent/totalnumberofpressescounted
func (m_ MTRSwitchClusterMultiPressCompleteEvent) TotalNumberOfPressesCounted() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("totalNumberOfPressesCounted"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclustermultipresscompleteevent/totalnumberofpressescounted
func (m_ MTRSwitchClusterMultiPressCompleteEvent) SetTotalNumberOfPressesCounted(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTotalNumberOfPressesCounted:"), value)
}



