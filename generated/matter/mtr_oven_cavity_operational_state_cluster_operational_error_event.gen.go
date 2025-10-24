// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROvenCavityOperationalStateClusterOperationalErrorEvent */


/* debug [class_header]: Header for MTROvenCavityOperationalStateClusterOperationalErrorEvent */
// The class instance for the [MTROvenCavityOperationalStateClusterOperationalErrorEvent] class.
var (
	MTROvenCavityOperationalStateClusterOperationalErrorEventClass     _MTROvenCavityOperationalStateClusterOperationalErrorEventClass
	MTROvenCavityOperationalStateClusterOperationalErrorEventClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterOperationalErrorEventClass() _MTROvenCavityOperationalStateClusterOperationalErrorEventClass {
	MTROvenCavityOperationalStateClusterOperationalErrorEventClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterOperationalErrorEventClass = _MTROvenCavityOperationalStateClusterOperationalErrorEventClass{objc.GetClass("MTROvenCavityOperationalStateClusterOperationalErrorEvent")}
	})
	return MTROvenCavityOperationalStateClusterOperationalErrorEventClass
}

type _MTROvenCavityOperationalStateClusterOperationalErrorEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROvenCavityOperationalStateClusterOperationalErrorEvent */
// An interface definition for the [MTROvenCavityOperationalStateClusterOperationalErrorEvent] class.
type IMTROvenCavityOperationalStateClusterOperationalErrorEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROvenCavityOperationalStateClusterOperationalErrorEvent */
	// properties:
	ErrorState() IMTROvenCavityOperationalStateClusterErrorStateStruct
	SetErrorState(value IMTROvenCavityOperationalStateClusterErrorStateStruct)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROvenCavityOperationalStateClusterOperationalErrorEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROvenCavityOperationalStateClusterOperationalErrorEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterOperationalErrorEventClass) Alloc() MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROvenCavityOperationalStateClusterOperationalErrorEventClass) New() MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalErrorEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterOperationalErrorEvent) Init() MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalErrorEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterOperationalErrorEvent) Autorelease() MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalErrorEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterOperationalErrorEvent creates a new MTROvenCavityOperationalStateClusterOperationalErrorEvent instance.
func NewMTROvenCavityOperationalStateClusterOperationalErrorEvent() MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	return getMTROvenCavityOperationalStateClusterOperationalErrorEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROvenCavityOperationalStateClusterOperationalErrorEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalErrorEvent
type MTROvenCavityOperationalStateClusterOperationalErrorEvent struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterOperationalErrorEventFrom constructs a [MTROvenCavityOperationalStateClusterOperationalErrorEvent] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterOperationalErrorEventFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	return MTROvenCavityOperationalStateClusterOperationalErrorEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROvenCavityOperationalStateClusterOperationalErrorEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROvenCavityOperationalStateClusterOperationalErrorEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROvenCavityOperationalStateClusterOperationalErrorEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROvenCavityOperationalStateClusterOperationalErrorEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROvenCavityOperationalStateClusterOperationalErrorEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalErrorEvent/errorState
func (m_ MTROvenCavityOperationalStateClusterOperationalErrorEvent) ErrorState() IMTROvenCavityOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("errorState"))
	return rv
}/* debug [instance_properties/getter]: errorState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalErrorEvent/errorState
func (m_ MTROvenCavityOperationalStateClusterOperationalErrorEvent) SetErrorState(value IMTROvenCavityOperationalStateClusterErrorStateStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorState:"), value)
}/* debug [instance_properties/setter]: errorState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROvenCavityOperationalStateClusterOperationalErrorEvent */



