// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [UNNotificationAction] class.
type IUNNotificationAction interface {
	objectivec.IObject
}

// A task your app performs in response to a notification that the system delivers.
//
// Use objects to define the actions that your app can perform in response to a delivered notification. You define the actions that your app supports. For example, a meeting app might define actions for accepting or rejecting a meeting invitation. The action object itself contains the title to display in an action button and the button’s appearance. After creating action objects, add them to a object and register your categories with the system. For information on how to define actions and categories, see .
//
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

// Alloc allocates a new instance without initialization.
func (uc _UNNotificationActionClass) Alloc() UNNotificationAction {
	rv := objc.Send[UNNotificationAction](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates an action object by using the specified title and options.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/init(identifier:title:options:)
func NewUNNotificationActionWithIdentifierTitleOptions(identifier appkit.string, title appkit.string, options UNNotificationActionOptions) UNNotificationAction {
	rv := objc.Send[UNNotificationAction](objc.ID(getUNNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:"), identifier, title, options)
	return rv
}



// Creates an action object by using the specified title, options, and icon.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/init(identifier:title:options:icon:)
func NewUNNotificationActionWithIdentifierTitleOptionsIcon(identifier appkit.string, title appkit.string, options UNNotificationActionOptions, icon IUNNotificationActionIcon) UNNotificationAction {
	rv := objc.Send[UNNotificationAction](objc.ID(getUNNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:icon:"), identifier, title, options, icon)
	return rv
}


// Creates an action object by using the specified title and options.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/init(identifier:title:options:)
func (uc _UNNotificationActionClass) ActionWithIdentifierTitleOptions(identifier appkit.string, title appkit.string, options UNNotificationActionOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("actionWithIdentifier:title:options:"), identifier, title, options)
	return rv
}

// Creates an action object by using the specified title, options, and icon.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/init(identifier:title:options:icon:)
func (uc _UNNotificationActionClass) ActionWithIdentifierTitleOptionsIcon(identifier appkit.string, title appkit.string, options UNNotificationActionOptions, icon IUNNotificationActionIcon) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("actionWithIdentifier:title:options:icon:"), identifier, title, options, icon)
	return rv
}

// The icon associated with the action.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/icon
func (u_ UNNotificationAction) Icon() UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](u_.ID, objc.Sel("icon"))
	return rv
}

// The unique string that your app uses to identify the action.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/identifier
func (u_ UNNotificationAction) Identifier() appkit.string {
	rv := objc.Send[appkit.string](u_.ID, objc.Sel("identifier"))
	return rv
}

// The behaviors associated with the action.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/options
func (u_ UNNotificationAction) Options() UNNotificationActionOptions {
	rv := objc.Send[UNNotificationActionOptions](u_.ID, objc.Sel("options"))
	return rv
}

// The localized string to use as the title of the action.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/title
func (u_ UNNotificationAction) Title() appkit.string {
	rv := objc.Send[appkit.string](u_.ID, objc.Sel("title"))
	return rv
}


