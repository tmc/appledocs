// Code generated from Apple documentation for IntentsUI. DO NOT EDIT.

package intentsui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/intents"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INUIAddVoiceShortcutButton] class.
var (
	INUIAddVoiceShortcutButtonClass     _INUIAddVoiceShortcutButtonClass
	INUIAddVoiceShortcutButtonClassOnce sync.Once
)

func getINUIAddVoiceShortcutButtonClass() _INUIAddVoiceShortcutButtonClass {
	INUIAddVoiceShortcutButtonClassOnce.Do(func() {
		INUIAddVoiceShortcutButtonClass = _INUIAddVoiceShortcutButtonClass{objc.GetClass("INUIAddVoiceShortcutButton")}
	})
	return INUIAddVoiceShortcutButtonClass
}

type _INUIAddVoiceShortcutButtonClass struct {
	class objc.Class
}

// An interface definition for the [INUIAddVoiceShortcutButton] class.
type IINUIAddVoiceShortcutButton interface {
	appkit.IButton
	// properties:
	CornerRadius() float64
	SetCornerRadius(value float64)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Shortcut() intents.INShortcut
	SetShortcut(value intents.INShortcut)
	Style() unsafe.Pointer
	SetStyle(value unsafe.Pointer)
	// methods:
}

// A button that allows the user to add or edit a shortcut.
//
// When the user performs an action such as placing an order for tomato soup, the app should provide the option to add the action to Siri as a shortcut. To present this option in your app, use to display an “Add to Siri” button. Using this button makes your app consistent with other apps that support Siri Shortcuts. Set the property on the button to have it automatically update the status of the shortcut. If the user has already added the shortcut to Siri, the button displays “Added” instead of “Add” and includes the phrase that the user chose when adding the shortcut. The methods in aren’t called unless the property is set. After creating the button, assign its action to a method that displays . This controller guides the user through the process of adding the shortcut to Siri. The code listing below adds an “Add to Siri” button to a view and lets the user record an invocation phrase

// A button that allows the user to add or edit a shortcut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IntentsUI/INUIAddVoiceShortcutButton
type INUIAddVoiceShortcutButton struct {
	appkit.Button
}

// INUIAddVoiceShortcutButtonFrom constructs a [INUIAddVoiceShortcutButton] from an unsafe.Pointer.
//
// A button that allows the user to add or edit a shortcut.
func INUIAddVoiceShortcutButtonFrom(ptr unsafe.Pointer) INUIAddVoiceShortcutButton {
	return INUIAddVoiceShortcutButton{
		Button: appkit.ButtonFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INUIAddVoiceShortcutButtonClass) Alloc() INUIAddVoiceShortcutButton {
	rv := objc.Send[INUIAddVoiceShortcutButton](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INUIAddVoiceShortcutButtonClass) New() INUIAddVoiceShortcutButton {
	rv := objc.Send[INUIAddVoiceShortcutButton](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INUIAddVoiceShortcutButton) Init() INUIAddVoiceShortcutButton {
	rv := objc.Send[INUIAddVoiceShortcutButton](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INUIAddVoiceShortcutButton) Autorelease() INUIAddVoiceShortcutButton {
	rv := objc.Send[INUIAddVoiceShortcutButton](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINUIAddVoiceShortcutButton creates a new INUIAddVoiceShortcutButton instance.
func NewINUIAddVoiceShortcutButton() INUIAddVoiceShortcutButton {
	return getINUIAddVoiceShortcutButtonClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intentsui/inuiaddvoiceshortcutbutton/cornerradius
func (i_ INUIAddVoiceShortcutButton) CornerRadius() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("cornerRadius"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intentsui/inuiaddvoiceshortcutbutton/cornerradius
func (i_ INUIAddVoiceShortcutButton) SetCornerRadius(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCornerRadius:"), value)
}

// The object that receives presentation requests from the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intentsui/inuiaddvoiceshortcutbutton/delegate
func (i_ INUIAddVoiceShortcutButton) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}

// The object that receives presentation requests from the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intentsui/inuiaddvoiceshortcutbutton/delegate
func (i_ INUIAddVoiceShortcutButton) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}

// The shortcut Siri invokes when the user speaks the invocation phrase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intentsui/inuiaddvoiceshortcutbutton/shortcut
func (i_ INUIAddVoiceShortcutButton) Shortcut() intents.INShortcut {
	rv := objc.Send[intents.INShortcut](i_.ID, objc.Sel("shortcut"))
	return rv
}

// The shortcut Siri invokes when the user speaks the invocation phrase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intentsui/inuiaddvoiceshortcutbutton/shortcut
func (i_ INUIAddVoiceShortcutButton) SetShortcut(value intents.INShortcut) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setShortcut:"), value)
}

// The button style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intentsui/inuiaddvoiceshortcutbutton/style
func (i_ INUIAddVoiceShortcutButton) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("style"))
	return rv
}

// The button style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intentsui/inuiaddvoiceshortcutbutton/style
func (i_ INUIAddVoiceShortcutButton) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setStyle:"), value)
}
