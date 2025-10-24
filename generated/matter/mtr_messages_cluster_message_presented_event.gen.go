// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMessagesClusterMessagePresentedEvent */


/* debug [class_header]: Header for MTRMessagesClusterMessagePresentedEvent */
// The class instance for the [MTRMessagesClusterMessagePresentedEvent] class.
var (
	MTRMessagesClusterMessagePresentedEventClass     _MTRMessagesClusterMessagePresentedEventClass
	MTRMessagesClusterMessagePresentedEventClassOnce sync.Once
)

func getMTRMessagesClusterMessagePresentedEventClass() _MTRMessagesClusterMessagePresentedEventClass {
	MTRMessagesClusterMessagePresentedEventClassOnce.Do(func() {
		MTRMessagesClusterMessagePresentedEventClass = _MTRMessagesClusterMessagePresentedEventClass{objc.GetClass("MTRMessagesClusterMessagePresentedEvent")}
	})
	return MTRMessagesClusterMessagePresentedEventClass
}

type _MTRMessagesClusterMessagePresentedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMessagesClusterMessagePresentedEvent */
// An interface definition for the [MTRMessagesClusterMessagePresentedEvent] class.
type IMTRMessagesClusterMessagePresentedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMessagesClusterMessagePresentedEvent */
	// properties:
	MessageID() objc.IObject /* cross-framework: NSData */
	SetMessageID(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMessagesClusterMessagePresentedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMessagesClusterMessagePresentedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterMessagePresentedEventClass) Alloc() MTRMessagesClusterMessagePresentedEvent {
	rv := objc.Send[MTRMessagesClusterMessagePresentedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMessagesClusterMessagePresentedEventClass) New() MTRMessagesClusterMessagePresentedEvent {
	rv := objc.Send[MTRMessagesClusterMessagePresentedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterMessagePresentedEvent) Init() MTRMessagesClusterMessagePresentedEvent {
	rv := objc.Send[MTRMessagesClusterMessagePresentedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterMessagePresentedEvent) Autorelease() MTRMessagesClusterMessagePresentedEvent {
	rv := objc.Send[MTRMessagesClusterMessagePresentedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterMessagePresentedEvent creates a new MTRMessagesClusterMessagePresentedEvent instance.
func NewMTRMessagesClusterMessagePresentedEvent() MTRMessagesClusterMessagePresentedEvent {
	return getMTRMessagesClusterMessagePresentedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMessagesClusterMessagePresentedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessagePresentedEvent
type MTRMessagesClusterMessagePresentedEvent struct {
	objectivec.Object
}

// MTRMessagesClusterMessagePresentedEventFrom constructs a [MTRMessagesClusterMessagePresentedEvent] from an unsafe.Pointer.
func MTRMessagesClusterMessagePresentedEventFrom(ptr unsafe.Pointer) MTRMessagesClusterMessagePresentedEvent {
	return MTRMessagesClusterMessagePresentedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMessagesClusterMessagePresentedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMessagesClusterMessagePresentedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMessagesClusterMessagePresentedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMessagesClusterMessagePresentedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMessagesClusterMessagePresentedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessagePresentedEvent/messageID
func (m_ MTRMessagesClusterMessagePresentedEvent) MessageID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("messageID"))
	return rv
}/* debug [instance_properties/getter]: messageID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessagePresentedEvent/messageID
func (m_ MTRMessagesClusterMessagePresentedEvent) SetMessageID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}/* debug [instance_properties/setter]: messageID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMessagesClusterMessagePresentedEvent */



