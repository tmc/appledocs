// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEMessageAction */


/* debug [class_header]: Header for MEMessageAction */
// The class instance for the [MEMessageAction] class.
var (
	MEMessageActionClass     _MEMessageActionClass
	MEMessageActionClassOnce sync.Once
)

func getMEMessageActionClass() _MEMessageActionClass {
	MEMessageActionClassOnce.Do(func() {
		MEMessageActionClass = _MEMessageActionClass{objc.GetClass("MEMessageAction")}
	})
	return MEMessageActionClass
}

type _MEMessageActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEMessageAction */
// An interface definition for the [MEMessageAction] class.
type IMEMessageAction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEMessageAction */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEMessageAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEMessageAction */
// Alloc allocates a new instance without initialization.
func (mc _MEMessageActionClass) Alloc() MEMessageAction {
	rv := objc.Send[MEMessageAction](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEMessageActionClass) New() MEMessageAction {
	rv := objc.Send[MEMessageAction](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEMessageAction) Init() MEMessageAction {
	rv := objc.Send[MEMessageAction](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEMessageAction) Autorelease() MEMessageAction {
	rv := objc.Send[MEMessageAction](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEMessageAction creates a new MEMessageAction instance.
func NewMEMessageAction() MEMessageAction {
	return getMEMessageActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEMessageAction */
// An action the system performs on a message, such as setting a color or archiving it.


// An action the system performs on a message, such as setting a color or archiving it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction
type MEMessageAction struct {
	objectivec.Object
}

// MEMessageActionFrom constructs a [MEMessageAction] from an unsafe.Pointer.
//
// An action the system performs on a message, such as setting a color or archiving it.
func MEMessageActionFrom(ptr unsafe.Pointer) MEMessageAction {
	return MEMessageAction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEMessageAction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEMessageAction */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/flag(_:)
func (mc _MEMessageActionClass) FlagActionWithFlag(flag MEMessageActionFlag) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("flagActionWithFlag:"), flag)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FlagActionWithFlag) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/setBackgroundColor(_:)
func (mc _MEMessageActionClass) SetBackgroundColorActionWithColor(color MEMessageActionMessageColor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("setBackgroundColorActionWithColor:"), color)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetBackgroundColorActionWithColor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEMessageAction */

// An object that marks the message as read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/markAsRead
func (mc _MEMessageActionClass) MarkAsReadAction() MEMessageAction {
	rv := objc.Send[MEMessageAction](objc.ID(mc.class), objc.Sel("markAsReadAction"))
	return rv
}/* debug [class_properties_class/property]: markAsReadAction */

// An object that marks the message as unread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/markAsUnread
func (mc _MEMessageActionClass) MarkAsUnreadAction() MEMessageAction {
	rv := objc.Send[MEMessageAction](objc.ID(mc.class), objc.Sel("markAsUnreadAction"))
	return rv
}/* debug [class_properties_class/property]: markAsUnreadAction */

// An object that moves the message to the account’s Archive mailbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/moveToArchive
func (mc _MEMessageActionClass) MoveToArchiveAction() MEMessageAction {
	rv := objc.Send[MEMessageAction](objc.ID(mc.class), objc.Sel("moveToArchiveAction"))
	return rv
}/* debug [class_properties_class/property]: moveToArchiveAction */

// An object that moves the message to the account’s Junk mailbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/moveToJunk
func (mc _MEMessageActionClass) MoveToJunkAction() MEMessageAction {
	rv := objc.Send[MEMessageAction](objc.ID(mc.class), objc.Sel("moveToJunkAction"))
	return rv
}/* debug [class_properties_class/property]: moveToJunkAction */

// An object that moves the message to the account’s Trash mailbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/moveToTrash
func (mc _MEMessageActionClass) MoveToTrashAction() MEMessageAction {
	rv := objc.Send[MEMessageAction](objc.ID(mc.class), objc.Sel("moveToTrashAction"))
	return rv
}/* debug [class_properties_class/property]: moveToTrashAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEMessageAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEMessageAction */

// An object that marks the message as read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/markAsRead
func (m_ MEMessageAction) MarkAsReadAction() IMEMessageAction {
	rv := objc.Send[MEMessageAction](m_.ID, objc.Sel("markAsReadAction"))
	return rv
}/* debug [instance_properties/getter]: markAsReadAction */


// An object that marks the message as unread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/markAsUnread
func (m_ MEMessageAction) MarkAsUnreadAction() IMEMessageAction {
	rv := objc.Send[MEMessageAction](m_.ID, objc.Sel("markAsUnreadAction"))
	return rv
}/* debug [instance_properties/getter]: markAsUnreadAction */


// An object that moves the message to the account’s Archive mailbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/moveToArchive
func (m_ MEMessageAction) MoveToArchiveAction() IMEMessageAction {
	rv := objc.Send[MEMessageAction](m_.ID, objc.Sel("moveToArchiveAction"))
	return rv
}/* debug [instance_properties/getter]: moveToArchiveAction */


// An object that moves the message to the account’s Junk mailbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/moveToJunk
func (m_ MEMessageAction) MoveToJunkAction() IMEMessageAction {
	rv := objc.Send[MEMessageAction](m_.ID, objc.Sel("moveToJunkAction"))
	return rv
}/* debug [instance_properties/getter]: moveToJunkAction */


// An object that moves the message to the account’s Trash mailbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageAction/moveToTrash
func (m_ MEMessageAction) MoveToTrashAction() IMEMessageAction {
	rv := objc.Send[MEMessageAction](m_.ID, objc.Sel("moveToTrashAction"))
	return rv
}/* debug [instance_properties/getter]: moveToTrashAction */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEMessageAction */



