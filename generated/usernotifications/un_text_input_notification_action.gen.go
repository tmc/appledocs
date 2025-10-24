// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class UNTextInputNotificationAction */

/* debug [class_header]: Header for UNTextInputNotificationAction */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for UNTextInputNotificationAction */
// An interface definition for the [UNTextInputNotificationAction] class.
type IUNTextInputNotificationAction interface {
	IUNNotificationAction

	/* debug [class_interface_properties]: Properties for UNTextInputNotificationAction */
	// properties:
	TextInputButtonTitle() objc.IObject /* cross-framework: NSString */
	TextInputPlaceholder() objc.IObject /* cross-framework: NSString */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for UNTextInputNotificationAction */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for UNTextInputNotificationAction */
// Alloc allocates a new instance without initialization.
func (uc _UNTextInputNotificationActionClass) Alloc() UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for UNTextInputNotificationAction */
// An action that accepts user-typed text.
//
// Use objects to define an action that allows the user to provide a custom text-based response. When the user selects an action of this type, the system displays controls for the user to enter or dictate the text content. That text is then included in the response object that’s delivered to your app. For information on how to define actions and categories, see .

// An action that accepts user-typed text.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for UNTextInputNotificationAction */

// Creates an action object with an icon that accepts text input from the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/init(identifier:title:options:icon:textInputButtonTitle:textInputPlaceholder:)
func NewUNTextInputNotificationActionWithIdentifierTitleOptionsIconTextInputButtonTitleTextInputPlaceholder(identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, options UNNotificationActionOptions, icon IUNNotificationActionIcon, textInputButtonTitle objc.IObject /* cross-framework: NSString */, textInputPlaceholder objc.IObject /* cross-framework: NSString */) UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](objc.ID(getUNTextInputNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:icon:textInputButtonTitle:textInputPlaceholder:"), identifier, title, options, icon, textInputButtonTitle, textInputPlaceholder)
	return rv
} /* debug [class_init_methods/constructor]: NewUNTextInputNotificationActionWithIdentifierTitleOptionsIconTextInputButtonTitleTextInputPlaceholder */

// Creates an action object that accepts text input from the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/init(identifier:title:options:textInputButtonTitle:textInputPlaceholder:)
func NewUNTextInputNotificationActionWithIdentifierTitleOptionsTextInputButtonTitleTextInputPlaceholder(identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, options UNNotificationActionOptions, textInputButtonTitle objc.IObject /* cross-framework: NSString */, textInputPlaceholder objc.IObject /* cross-framework: NSString */) UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](objc.ID(getUNTextInputNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:textInputButtonTitle:textInputPlaceholder:"), identifier, title, options, textInputButtonTitle, textInputPlaceholder)
	return rv
} /* debug [class_init_methods/constructor]: NewUNTextInputNotificationActionWithIdentifierTitleOptionsTextInputButtonTitleTextInputPlaceholder */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for UNTextInputNotificationAction */

// Creates an action object with an icon that accepts text input from the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/init(identifier:title:options:icon:textInputButtonTitle:textInputPlaceholder:)
func (uc _UNTextInputNotificationActionClass) ActionWithIdentifierTitleOptionsIconTextInputButtonTitleTextInputPlaceholder(identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, options UNNotificationActionOptions, icon IUNNotificationActionIcon, textInputButtonTitle objc.IObject /* cross-framework: NSString */, textInputPlaceholder objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("actionWithIdentifier:title:options:icon:textInputButtonTitle:textInputPlaceholder:"), identifier, title, options, icon, textInputButtonTitle, textInputPlaceholder)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=ActionWithIdentifierTitleOptionsIconTextInputButtonTitleTextInputPlaceholder) */

// Creates an action object that accepts text input from the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/init(identifier:title:options:textInputButtonTitle:textInputPlaceholder:)
func (uc _UNTextInputNotificationActionClass) ActionWithIdentifierTitleOptionsTextInputButtonTitleTextInputPlaceholder(identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, options UNNotificationActionOptions, textInputButtonTitle objc.IObject /* cross-framework: NSString */, textInputPlaceholder objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("actionWithIdentifier:title:options:textInputButtonTitle:textInputPlaceholder:"), identifier, title, options, textInputButtonTitle, textInputPlaceholder)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=ActionWithIdentifierTitleOptionsTextInputButtonTitleTextInputPlaceholder) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for UNTextInputNotificationAction */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for UNTextInputNotificationAction */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for UNTextInputNotificationAction */

// The localized title of the text input button that the system displays to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/textInputButtonTitle
func (u_ UNTextInputNotificationAction) TextInputButtonTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("textInputButtonTitle"))
	return rv
} /* debug [instance_properties/getter]: textInputButtonTitle */

// The placeholder text that the system localizes and displays in the text input field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/textInputPlaceholder
func (u_ UNTextInputNotificationAction) TextInputPlaceholder() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("textInputPlaceholder"))
	return rv
} /* debug [instance_properties/getter]: textInputPlaceholder */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class UNTextInputNotificationAction */
