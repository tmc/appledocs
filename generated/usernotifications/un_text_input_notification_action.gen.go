// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UNTextInputNotificationAction] class.
var (
	UNTextInputNotificationActionClass     _UNTextInputNotificationActionClass
	UNTextInputNotificationActionClassOnce sync.Once
)

func getUNTextInputNotificationActionClass() _UNTextInputNotificationActionClass {
	UNTextInputNotificationActionClassOnce.Do(func() {
		UNTextInputNotificationActionClass = _UNTextInputNotificationActionClass{objc.GetClass("UNTextInputNotificationAction")}
	})
	return UNTextInputNotificationActionClass
}

type _UNTextInputNotificationActionClass struct {
	class objc.Class
}

// An interface definition for the [UNTextInputNotificationAction] class.
type IUNTextInputNotificationAction interface {
	IUNNotificationAction
}

// An action that accepts user-typed text.
//
// Use objects to define an action that allows the user to provide a custom text-based response. When the user selects an action of this type, the system displays controls for the user to enter or dictate the text content. That text is then included in the response object that’s delivered to your app. For information on how to define actions and categories, see .
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction
type UNTextInputNotificationAction struct {
	UNNotificationAction
}

// UNTextInputNotificationActionFrom constructs a [UNTextInputNotificationAction] from an unsafe.Pointer.
//
// An action that accepts user-typed text.
func UNTextInputNotificationActionFrom(ptr unsafe.Pointer) UNTextInputNotificationAction {
	return UNTextInputNotificationAction{
		UNNotificationAction: UNNotificationActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UNTextInputNotificationActionClass) Alloc() UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UNTextInputNotificationActionClass) New() UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNTextInputNotificationAction) Init() UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNTextInputNotificationAction) Autorelease() UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNTextInputNotificationAction creates a new UNTextInputNotificationAction instance.
func NewUNTextInputNotificationAction() UNTextInputNotificationAction {
	return getUNTextInputNotificationActionClass().New()
}


// Creates an action object that accepts text input from the user.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/init(identifier:title:options:textInputButtonTitle:textInputPlaceholder:)
func NewUNTextInputNotificationActionWithIdentifierTitleOptionsTextInputButtonTitleTextInputPlaceholder(identifier string, title string, options unsafe.Pointer, textInputButtonTitle string, textInputPlaceholder string) UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](objc.ID(getUNTextInputNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:textInputButtonTitle:textInputPlaceholder:"), objc.String(identifier), objc.String(title), options, objc.String(textInputButtonTitle), objc.String(textInputPlaceholder))
	return rv
}

// Creates an action object with an icon that accepts text input from the user.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/init(identifier:title:options:icon:textInputButtonTitle:textInputPlaceholder:)
func NewUNTextInputNotificationActionWithIdentifierTitleOptionsIconTextInputButtonTitleTextInputPlaceholder(identifier string, title string, options unsafe.Pointer, icon unsafe.Pointer, textInputButtonTitle string, textInputPlaceholder string) UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](objc.ID(getUNTextInputNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:icon:textInputButtonTitle:textInputPlaceholder:"), objc.String(identifier), objc.String(title), options, icon, objc.String(textInputButtonTitle), objc.String(textInputPlaceholder))
	return rv
}


// Creates an action object with an icon that accepts text input from the user.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/init(identifier:title:options:icon:textInputButtonTitle:textInputPlaceholder:)
func (uc _UNTextInputNotificationActionClass) ActionWithIdentifierTitleOptionsIconTextInputButtonTitleTextInputPlaceholder(identifier string, title string, options unsafe.Pointer, icon unsafe.Pointer, textInputButtonTitle string, textInputPlaceholder string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("actionWithIdentifier:title:options:icon:textInputButtonTitle:textInputPlaceholder:"), objc.String(identifier), objc.String(title), options, icon, objc.String(textInputButtonTitle), objc.String(textInputPlaceholder))
	return rv
}

// Creates an action object that accepts text input from the user.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/init(identifier:title:options:textInputButtonTitle:textInputPlaceholder:)
func (uc _UNTextInputNotificationActionClass) ActionWithIdentifierTitleOptionsTextInputButtonTitleTextInputPlaceholder(identifier string, title string, options unsafe.Pointer, textInputButtonTitle string, textInputPlaceholder string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("actionWithIdentifier:title:options:textInputButtonTitle:textInputPlaceholder:"), objc.String(identifier), objc.String(title), options, objc.String(textInputButtonTitle), objc.String(textInputPlaceholder))
	return rv
}

// The localized title of the text input button that the system displays to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/textInputButtonTitle
func (u_ UNTextInputNotificationAction) TextInputButtonTitle() string {
	rv := objc.Send[string](u_.ID, objc.Sel("textInputButtonTitle"))
	return rv
}

// The placeholder text that the system localizes and displays in the text input field.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/textInputPlaceholder
func (u_ UNTextInputNotificationAction) TextInputPlaceholder() string {
	rv := objc.Send[string](u_.ID, objc.Sel("textInputPlaceholder"))
	return rv
}


