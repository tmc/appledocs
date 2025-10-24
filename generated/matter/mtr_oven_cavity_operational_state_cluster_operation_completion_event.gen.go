// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROvenCavityOperationalStateClusterOperationCompletionEvent */


/* debug [class_header]: Header for MTROvenCavityOperationalStateClusterOperationCompletionEvent */
// The class instance for the [MTROvenCavityOperationalStateClusterOperationCompletionEvent] class.
var (
	MTROvenCavityOperationalStateClusterOperationCompletionEventClass     _MTROvenCavityOperationalStateClusterOperationCompletionEventClass
	MTROvenCavityOperationalStateClusterOperationCompletionEventClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterOperationCompletionEventClass() _MTROvenCavityOperationalStateClusterOperationCompletionEventClass {
	MTROvenCavityOperationalStateClusterOperationCompletionEventClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterOperationCompletionEventClass = _MTROvenCavityOperationalStateClusterOperationCompletionEventClass{objc.GetClass("MTROvenCavityOperationalStateClusterOperationCompletionEvent")}
	})
	return MTROvenCavityOperationalStateClusterOperationCompletionEventClass
}

type _MTROvenCavityOperationalStateClusterOperationCompletionEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROvenCavityOperationalStateClusterOperationCompletionEvent */
// An interface definition for the [MTROvenCavityOperationalStateClusterOperationCompletionEvent] class.
type IMTROvenCavityOperationalStateClusterOperationCompletionEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROvenCavityOperationalStateClusterOperationCompletionEvent */
	// properties:
	CompletionErrorCode() objc.IObject /* cross-framework: NSNumber */
	SetCompletionErrorCode(value objc.IObject /* cross-framework: NSNumber */)
	PausedTime() objc.IObject /* cross-framework: NSNumber */
	SetPausedTime(value objc.IObject /* cross-framework: NSNumber */)
	TotalOperationalTime() objc.IObject /* cross-framework: NSNumber */
	SetTotalOperationalTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROvenCavityOperationalStateClusterOperationCompletionEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROvenCavityOperationalStateClusterOperationCompletionEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterOperationCompletionEventClass) Alloc() MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationCompletionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROvenCavityOperationalStateClusterOperationCompletionEventClass) New() MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationCompletionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) Init() MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationCompletionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) Autorelease() MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationCompletionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterOperationCompletionEvent creates a new MTROvenCavityOperationalStateClusterOperationCompletionEvent instance.
func NewMTROvenCavityOperationalStateClusterOperationCompletionEvent() MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	return getMTROvenCavityOperationalStateClusterOperationCompletionEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROvenCavityOperationalStateClusterOperationCompletionEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationCompletionEvent
type MTROvenCavityOperationalStateClusterOperationCompletionEvent struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterOperationCompletionEventFrom constructs a [MTROvenCavityOperationalStateClusterOperationCompletionEvent] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterOperationCompletionEventFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	return MTROvenCavityOperationalStateClusterOperationCompletionEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROvenCavityOperationalStateClusterOperationCompletionEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROvenCavityOperationalStateClusterOperationCompletionEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROvenCavityOperationalStateClusterOperationCompletionEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROvenCavityOperationalStateClusterOperationCompletionEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROvenCavityOperationalStateClusterOperationCompletionEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationCompletionEvent/completionErrorCode
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) CompletionErrorCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("completionErrorCode"))
	return rv
}/* debug [instance_properties/getter]: completionErrorCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationCompletionEvent/completionErrorCode
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) SetCompletionErrorCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCompletionErrorCode:"), value)
}/* debug [instance_properties/setter]: completionErrorCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclusteroperationcompletionevent/pausedtime
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) PausedTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("pausedTime"))
	return rv
}/* debug [instance_properties/getter]: pausedTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclusteroperationcompletionevent/pausedtime
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) SetPausedTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPausedTime:"), value)
}/* debug [instance_properties/setter]: pausedTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclusteroperationcompletionevent/totaloperationaltime
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) TotalOperationalTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("totalOperationalTime"))
	return rv
}/* debug [instance_properties/getter]: totalOperationalTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclusteroperationcompletionevent/totaloperationaltime
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) SetTotalOperationalTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTotalOperationalTime:"), value)
}/* debug [instance_properties/setter]: totalOperationalTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROvenCavityOperationalStateClusterOperationCompletionEvent */



