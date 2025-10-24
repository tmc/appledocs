// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSDraggingSession */


/* debug [class_header]: Header for NSDraggingSession */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DraggingSession */
// An interface definition for the [DraggingSession] class.
type IDraggingSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DraggingSession */
	// properties:
	AnimatesToStartingPositionsOnCancelOrFail() bool
	SetAnimatesToStartingPositionsOnCancelOrFail(value bool)
	DraggingFormation() DraggingFormation
	SetDraggingFormation(value DraggingFormation)
	DraggingLeaderIndex() int
	SetDraggingLeaderIndex(value int)
	DraggingLocation() vision.Point
	DraggingPasteboard() IPasteboard
	DraggingSequenceNumber() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DraggingSession */
	// methods:
	EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.Class, searchOptions foundation.IDictionary, block unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DraggingSession */
// Alloc allocates a new instance without initialization.
func (dc _DraggingSessionClass) Alloc() DraggingSession {
	rv := objc.Send[DraggingSession](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DraggingSession */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DraggingSession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DraggingSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DraggingSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DraggingSession */

// Enumerates through each dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/enumerateDraggingItems(options:for:classes:searchOptions:using:)
func (d_ DraggingSession) EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.Class, searchOptions foundation.IDictionary, block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:"), enumOpts, view, classArray, searchOptions, block)
}/* debug [instance_methods/method]: EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DraggingSession */

// Controls whether the dragging image animates back to its starting point on a cancelled or failed drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/animatesToStartingPositionsOnCancelOrFail
func (d_ DraggingSession) AnimatesToStartingPositionsOnCancelOrFail() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("animatesToStartingPositionsOnCancelOrFail"))
	return rv
}/* debug [instance_properties/getter]: animatesToStartingPositionsOnCancelOrFail */


// Controls whether the dragging image animates back to its starting point on a cancelled or failed drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/animatesToStartingPositionsOnCancelOrFail
func (d_ DraggingSession) SetAnimatesToStartingPositionsOnCancelOrFail(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAnimatesToStartingPositionsOnCancelOrFail:"), value)
}/* debug [instance_properties/setter]: animatesToStartingPositionsOnCancelOrFail */


// Controls the dragging formation when the drag is not over the source or a valid destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingFormation
func (d_ DraggingSession) DraggingFormation() DraggingFormation {
	rv := objc.Send[DraggingFormation](d_.ID, objc.Sel("draggingFormation"))
	return rv
}/* debug [instance_properties/getter]: draggingFormation */


// Controls the dragging formation when the drag is not over the source or a valid destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingFormation
func (d_ DraggingSession) SetDraggingFormation(value DraggingFormation) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDraggingFormation:"), value)
}/* debug [instance_properties/setter]: draggingFormation */


// The index of the dragging item under the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingLeaderIndex
func (d_ DraggingSession) DraggingLeaderIndex() int {
	rv := objc.Send[int](d_.ID, objc.Sel("draggingLeaderIndex"))
	return rv
}/* debug [instance_properties/getter]: draggingLeaderIndex */


// The index of the dragging item under the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingLeaderIndex
func (d_ DraggingSession) SetDraggingLeaderIndex(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDraggingLeaderIndex:"), value)
}/* debug [instance_properties/setter]: draggingLeaderIndex */


// The current cursor location of the drag in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingLocation
func (d_ DraggingSession) DraggingLocation() vision.Point {
	rv := objc.Send[vision.Point](d_.ID, objc.Sel("draggingLocation"))
	return rv
}/* debug [instance_properties/getter]: draggingLocation */


// Returns the pasteboard object that contains the data being dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingPasteboard
func (d_ DraggingSession) DraggingPasteboard() IPasteboard {
	rv := objc.Send[Pasteboard](d_.ID, objc.Sel("draggingPasteboard"))
	return rv
}/* debug [instance_properties/getter]: draggingPasteboard */


// Returns a number that uniquely identifies the dragging session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingSequenceNumber
func (d_ DraggingSession) DraggingSequenceNumber() int {
	rv := objc.Send[int](d_.ID, objc.Sel("draggingSequenceNumber"))
	return rv
}/* debug [instance_properties/getter]: draggingSequenceNumber */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDraggingSession */



