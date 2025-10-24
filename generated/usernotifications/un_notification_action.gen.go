// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UNNotificationAction */


/* debug [class_header]: Header for UNNotificationAction */
// The class instance for the [UNNotificationAction] class.
var (
	UNNotificationActionClass     _UNNotificationActionClass
	UNNotificationActionClassOnce sync.Once
)

func getUNNotificationActionClass() _UNNotificationActionClass {
	UNNotificationActionClassOnce.Do(func() {
		UNNotificationActionClass = _UNNotificationActionClass{objc.GetClass("UNNotificationAction")}
	})
	return UNNotificationActionClass
}

type _UNNotificationActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UNNotificationAction */
// An interface definition for the [UNNotificationAction] class.
type IUNNotificationAction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UNNotificationAction */
	// properties:
	Icon() IUNNotificationActionIcon
	Identifier() objc.IObject /* cross-framework: NSString */
	Options() UNNotificationActionOptions
	Title() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UNNotificationAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UNNotificationAction */
// Alloc allocates a new instance without initialization.
func (uc _UNNotificationActionClass) Alloc() UNNotificationAction {
	rv := objc.Send[UNNotificationAction](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNNotificationActionClass) New() UNNotificationAction {
	rv := objc.Send[UNNotificationAction](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationAction) Init() UNNotificationAction {
	rv := objc.Send[UNNotificationAction](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationAction) Autorelease() UNNotificationAction {
	rv := objc.Send[UNNotificationAction](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationAction creates a new UNNotificationAction instance.
func NewUNNotificationAction() UNNotificationAction {
	return getUNNotificationActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UNNotificationAction */
// A task your app performs in response to a notification that the system delivers.
//
// Use objects to define the actions that your app can perform in response to a delivered notification. You define the actions that your app supports. For example, a meeting app might define actions for accepting or rejecting a meeting invitation. The action object itself contains the title to display in an action button and the button’s appearance. After creating action objects, add them to a object and register your categories with the system. For information on how to define actions and categories, see .


// A task your app performs in response to a notification that the system delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction
type UNNotificationAction struct {
	objectivec.Object
}

// UNNotificationActionFrom constructs a [UNNotificationAction] from an unsafe.Pointer.
//
// A task your app performs in response to a notification that the system delivers.
func UNNotificationActionFrom(ptr unsafe.Pointer) UNNotificationAction {
	return UNNotificationAction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UNNotificationAction */

// Creates an action object by using the specified title and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/init(identifier:title:options:)
func NewUNNotificationActionWithIdentifierTitleOptions(identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, options UNNotificationActionOptions) UNNotificationAction {
	rv := objc.Send[UNNotificationAction](objc.ID(getUNNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:"), identifier, title, options)
	return rv
}/* debug [class_init_methods/constructor]: NewUNNotificationActionWithIdentifierTitleOptions */


// Creates an action object by using the specified title, options, and icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/init(identifier:title:options:icon:)
func NewUNNotificationActionWithIdentifierTitleOptionsIcon(identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, options UNNotificationActionOptions, icon IUNNotificationActionIcon) UNNotificationAction {
	rv := objc.Send[UNNotificationAction](objc.ID(getUNNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:icon:"), identifier, title, options, icon)
	return rv
}/* debug [class_init_methods/constructor]: NewUNNotificationActionWithIdentifierTitleOptionsIcon */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UNNotificationAction */

// Creates an action object by using the specified title and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/init(identifier:title:options:)
func (uc _UNNotificationActionClass) ActionWithIdentifierTitleOptions(identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, options UNNotificationActionOptions) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("actionWithIdentifier:title:options:"), identifier, title, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ActionWithIdentifierTitleOptions) */


// Creates an action object by using the specified title, options, and icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/init(identifier:title:options:icon:)
func (uc _UNNotificationActionClass) ActionWithIdentifierTitleOptionsIcon(identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, options UNNotificationActionOptions, icon IUNNotificationActionIcon) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("actionWithIdentifier:title:options:icon:"), identifier, title, options, icon)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ActionWithIdentifierTitleOptionsIcon) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UNNotificationAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UNNotificationAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UNNotificationAction */

// The icon associated with the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/icon
func (u_ UNNotificationAction) Icon() IUNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](u_.ID, objc.Sel("icon"))
	return rv
}/* debug [instance_properties/getter]: icon */


// The unique string that your app uses to identify the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/identifier
func (u_ UNNotificationAction) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The behaviors associated with the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/options
func (u_ UNNotificationAction) Options() UNNotificationActionOptions {
	rv := objc.Send[UNNotificationActionOptions](u_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// The localized string to use as the title of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/title
func (u_ UNNotificationAction) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UNNotificationAction */


