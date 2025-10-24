// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DraggingSession] class.
var (
	DraggingSessionClass     _DraggingSessionClass
	DraggingSessionClassOnce sync.Once
)

func getDraggingSessionClass() _DraggingSessionClass {
	DraggingSessionClassOnce.Do(func() {
		DraggingSessionClass = _DraggingSessionClass{objc.GetClass("NSDraggingSession")}
	})
	return DraggingSessionClass
}

type _DraggingSessionClass struct {
	class objc.Class
}

// An interface definition for the [DraggingSession] class.
type IDraggingSession interface {
	objectivec.IObject
	// properties:
	AnimatesToStartingPositionsOnCancelOrFail() bool
	SetAnimatesToStartingPositionsOnCancelOrFail(value bool)
	DraggingFormation() DraggingFormation
	SetDraggingFormation(value DraggingFormation)
	DraggingLeaderIndex() int
	SetDraggingLeaderIndex(value int)
	DraggingLocation() objc.IObject /* cross-framework: Point */
	DraggingPasteboard() IPasteboard
	DraggingSequenceNumber() int
	// methods:
	EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.Class, searchOptions foundation.IDictionary, block unsafe.Pointer)
}

// The encapsulation of a drag-and-drop action that supports modification of the drag while in progress.
//
// You start a new dragging session by calling the method method. This method immediately returns and you can further modify the properties of the dragging session. The actual drag begins at the next turn of the run loop.


// The encapsulation of a drag-and-drop action that supports modification of the drag while in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession
type DraggingSession struct {
	objectivec.Object
}

// DraggingSessionFrom constructs a [DraggingSession] from an unsafe.Pointer.
//
// The encapsulation of a drag-and-drop action that supports modification of the drag while in progress.
func DraggingSessionFrom(ptr unsafe.Pointer) DraggingSession {
	return DraggingSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DraggingSessionClass) Alloc() DraggingSession {
	rv := objc.Send[DraggingSession](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DraggingSessionClass) New() DraggingSession {
	rv := objc.Send[DraggingSession](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DraggingSession) Init() DraggingSession {
	rv := objc.Send[DraggingSession](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DraggingSession) Autorelease() DraggingSession {
	rv := objc.Send[DraggingSession](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDraggingSession creates a new DraggingSession instance.
func NewDraggingSession() DraggingSession {
	return getDraggingSessionClass().New()
}



// Enumerates through each dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/enumerateDraggingItems(options:for:classes:searchOptions:using:)
func (d_ DraggingSession) EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.Class, searchOptions foundation.IDictionary, block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:"), enumOpts, view, classArray, searchOptions, block)
}


// Controls whether the dragging image animates back to its starting point on a cancelled or failed drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/animatesToStartingPositionsOnCancelOrFail
func (d_ DraggingSession) AnimatesToStartingPositionsOnCancelOrFail() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("animatesToStartingPositionsOnCancelOrFail"))
	return rv
}


// Controls whether the dragging image animates back to its starting point on a cancelled or failed drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/animatesToStartingPositionsOnCancelOrFail
func (d_ DraggingSession) SetAnimatesToStartingPositionsOnCancelOrFail(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAnimatesToStartingPositionsOnCancelOrFail:"), value)
}


// Controls the dragging formation when the drag is not over the source or a valid destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingFormation
func (d_ DraggingSession) DraggingFormation() DraggingFormation {
	rv := objc.Send[DraggingFormation](d_.ID, objc.Sel("draggingFormation"))
	return rv
}


// Controls the dragging formation when the drag is not over the source or a valid destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingFormation
func (d_ DraggingSession) SetDraggingFormation(value DraggingFormation) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDraggingFormation:"), value)
}


// The index of the dragging item under the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingLeaderIndex
func (d_ DraggingSession) DraggingLeaderIndex() int {
	rv := objc.Send[int](d_.ID, objc.Sel("draggingLeaderIndex"))
	return rv
}


// The index of the dragging item under the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingLeaderIndex
func (d_ DraggingSession) SetDraggingLeaderIndex(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDraggingLeaderIndex:"), value)
}


// The current cursor location of the drag in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingLocation
func (d_ DraggingSession) DraggingLocation() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](d_.ID, objc.Sel("draggingLocation"))
	return rv
}


// Returns the pasteboard object that contains the data being dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingPasteboard
func (d_ DraggingSession) DraggingPasteboard() IPasteboard {
	rv := objc.Send[Pasteboard](d_.ID, objc.Sel("draggingPasteboard"))
	return rv
}


// Returns a number that uniquely identifies the dragging session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingSequenceNumber
func (d_ DraggingSession) DraggingSequenceNumber() int {
	rv := objc.Send[int](d_.ID, objc.Sel("draggingSequenceNumber"))
	return rv
}



