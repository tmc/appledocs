// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMessagesClusterMessageStruct */


/* debug [class_header]: Header for MTRMessagesClusterMessageStruct */
// The class instance for the [MTRMessagesClusterMessageStruct] class.
var (
	MTRMessagesClusterMessageStructClass     _MTRMessagesClusterMessageStructClass
	MTRMessagesClusterMessageStructClassOnce sync.Once
)

func getMTRMessagesClusterMessageStructClass() _MTRMessagesClusterMessageStructClass {
	MTRMessagesClusterMessageStructClassOnce.Do(func() {
		MTRMessagesClusterMessageStructClass = _MTRMessagesClusterMessageStructClass{objc.GetClass("MTRMessagesClusterMessageStruct")}
	})
	return MTRMessagesClusterMessageStructClass
}

type _MTRMessagesClusterMessageStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMessagesClusterMessageStruct */
// An interface definition for the [MTRMessagesClusterMessageStruct] class.
type IMTRMessagesClusterMessageStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMessagesClusterMessageStruct */
	// properties:
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	MessageControl() objc.IObject /* cross-framework: NSNumber */
	SetMessageControl(value objc.IObject /* cross-framework: NSNumber */)
	MessageID() foundation.Data
	SetMessageID(value foundation.Data)
	MessageText() objc.IObject /* cross-framework: NSString */
	SetMessageText(value objc.IObject /* cross-framework: NSString */)
	Priority() objc.IObject /* cross-framework: NSNumber */
	SetPriority(value objc.IObject /* cross-framework: NSNumber */)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMessagesClusterMessageStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMessagesClusterMessageStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterMessageStructClass) Alloc() MTRMessagesClusterMessageStruct {
	rv := objc.Send[MTRMessagesClusterMessageStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMessagesClusterMessageStructClass) New() MTRMessagesClusterMessageStruct {
	rv := objc.Send[MTRMessagesClusterMessageStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterMessageStruct) Init() MTRMessagesClusterMessageStruct {
	rv := objc.Send[MTRMessagesClusterMessageStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterMessageStruct) Autorelease() MTRMessagesClusterMessageStruct {
	rv := objc.Send[MTRMessagesClusterMessageStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterMessageStruct creates a new MTRMessagesClusterMessageStruct instance.
func NewMTRMessagesClusterMessageStruct() MTRMessagesClusterMessageStruct {
	return getMTRMessagesClusterMessageStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMessagesClusterMessageStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct
type MTRMessagesClusterMessageStruct struct {
	objectivec.Object
}

// MTRMessagesClusterMessageStructFrom constructs a [MTRMessagesClusterMessageStruct] from an unsafe.Pointer.
func MTRMessagesClusterMessageStructFrom(ptr unsafe.Pointer) MTRMessagesClusterMessageStruct {
	return MTRMessagesClusterMessageStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMessagesClusterMessageStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMessagesClusterMessageStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMessagesClusterMessageStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMessagesClusterMessageStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMessagesClusterMessageStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/duration
func (m_ MTRMessagesClusterMessageStruct) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageStruct/duration
func (m_ MTRMessagesClusterMessageStruct) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagestruct/messagecontrol
func (m_ MTRMessagesClusterMessageStruct) MessageControl() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("messageControl"))
	return rv
}/* debug [instance_properties/getter]: messageControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagestruct/messagecontrol
func (m_ MTRMessagesClusterMessageStruct) SetMessageControl(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageControl:"), value)
}/* debug [instance_properties/setter]: messageControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagestruct/messageid
func (m_ MTRMessagesClusterMessageStruct) MessageID() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("messageID"))
	return rv
}/* debug [instance_properties/getter]: messageID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagestruct/messageid
func (m_ MTRMessagesClusterMessageStruct) SetMessageID(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}/* debug [instance_properties/setter]: messageID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagestruct/messagetext
func (m_ MTRMessagesClusterMessageStruct) MessageText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("messageText"))
	return rv
}/* debug [instance_properties/getter]: messageText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagestruct/messagetext
func (m_ MTRMessagesClusterMessageStruct) SetMessageText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageText:"), value)
}/* debug [instance_properties/setter]: messageText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagestruct/priority
func (m_ MTRMessagesClusterMessageStruct) Priority() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("priority"))
	return rv
}/* debug [instance_properties/getter]: priority */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagestruct/priority
func (m_ MTRMessagesClusterMessageStruct) SetPriority(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPriority:"), value)
}/* debug [instance_properties/setter]: priority */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagestruct/starttime
func (m_ MTRMessagesClusterMessageStruct) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}/* debug [instance_properties/getter]: startTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagestruct/starttime
func (m_ MTRMessagesClusterMessageStruct) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}/* debug [instance_properties/setter]: startTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMessagesClusterMessageStruct */



