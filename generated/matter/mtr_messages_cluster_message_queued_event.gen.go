// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMessagesClusterMessageQueuedEvent */


/* debug [class_header]: Header for MTRMessagesClusterMessageQueuedEvent */
// The class instance for the [MTRMessagesClusterMessageQueuedEvent] class.
var (
	MTRMessagesClusterMessageQueuedEventClass     _MTRMessagesClusterMessageQueuedEventClass
	MTRMessagesClusterMessageQueuedEventClassOnce sync.Once
)

func getMTRMessagesClusterMessageQueuedEventClass() _MTRMessagesClusterMessageQueuedEventClass {
	MTRMessagesClusterMessageQueuedEventClassOnce.Do(func() {
		MTRMessagesClusterMessageQueuedEventClass = _MTRMessagesClusterMessageQueuedEventClass{objc.GetClass("MTRMessagesClusterMessageQueuedEvent")}
	})
	return MTRMessagesClusterMessageQueuedEventClass
}

type _MTRMessagesClusterMessageQueuedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMessagesClusterMessageQueuedEvent */
// An interface definition for the [MTRMessagesClusterMessageQueuedEvent] class.
type IMTRMessagesClusterMessageQueuedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMessagesClusterMessageQueuedEvent */
	// properties:
	MessageID() objc.IObject /* cross-framework: NSData */
	SetMessageID(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMessagesClusterMessageQueuedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMessagesClusterMessageQueuedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterMessageQueuedEventClass) Alloc() MTRMessagesClusterMessageQueuedEvent {
	rv := objc.Send[MTRMessagesClusterMessageQueuedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMessagesClusterMessageQueuedEventClass) New() MTRMessagesClusterMessageQueuedEvent {
	rv := objc.Send[MTRMessagesClusterMessageQueuedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterMessageQueuedEvent) Init() MTRMessagesClusterMessageQueuedEvent {
	rv := objc.Send[MTRMessagesClusterMessageQueuedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterMessageQueuedEvent) Autorelease() MTRMessagesClusterMessageQueuedEvent {
	rv := objc.Send[MTRMessagesClusterMessageQueuedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterMessageQueuedEvent creates a new MTRMessagesClusterMessageQueuedEvent instance.
func NewMTRMessagesClusterMessageQueuedEvent() MTRMessagesClusterMessageQueuedEvent {
	return getMTRMessagesClusterMessageQueuedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMessagesClusterMessageQueuedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageQueuedEvent
type MTRMessagesClusterMessageQueuedEvent struct {
	objectivec.Object
}

// MTRMessagesClusterMessageQueuedEventFrom constructs a [MTRMessagesClusterMessageQueuedEvent] from an unsafe.Pointer.
func MTRMessagesClusterMessageQueuedEventFrom(ptr unsafe.Pointer) MTRMessagesClusterMessageQueuedEvent {
	return MTRMessagesClusterMessageQueuedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMessagesClusterMessageQueuedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMessagesClusterMessageQueuedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMessagesClusterMessageQueuedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMessagesClusterMessageQueuedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMessagesClusterMessageQueuedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageQueuedEvent/messageID
func (m_ MTRMessagesClusterMessageQueuedEvent) MessageID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("messageID"))
	return rv
}/* debug [instance_properties/getter]: messageID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageQueuedEvent/messageID
func (m_ MTRMessagesClusterMessageQueuedEvent) SetMessageID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}/* debug [instance_properties/setter]: messageID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMessagesClusterMessageQueuedEvent */



