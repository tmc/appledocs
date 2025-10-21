// Code generated from Apple documentation for IntentsUI. DO NOT EDIT.

package intentsui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [INUIAddVoiceShortcutViewController] class.
var (
	INUIAddVoiceShortcutViewControllerClass     _INUIAddVoiceShortcutViewControllerClass
	INUIAddVoiceShortcutViewControllerClassOnce sync.Once
)

func getINUIAddVoiceShortcutViewControllerClass() _INUIAddVoiceShortcutViewControllerClass {
	INUIAddVoiceShortcutViewControllerClassOnce.Do(func() {
		INUIAddVoiceShortcutViewControllerClass = _INUIAddVoiceShortcutViewControllerClass{objc.GetClass("INUIAddVoiceShortcutViewController")}
	})
	return INUIAddVoiceShortcutViewControllerClass
}

type _INUIAddVoiceShortcutViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [INUIAddVoiceShortcutViewController] class.
type IINUIAddVoiceShortcutViewController interface {
	appkit.IViewController
}

// A view controller that guides the user through the steps for adding a shortcut to Siri.
//
// When the user performs an action such as placing an order for tomato soup, the app should provide the option to add the action to Siri as a shortcut. To present this option in your app, use to display an button. Using this button makes your app consistent with other apps that support Siri Shortcuts. After creating the button, assign its action to a method that displays . This controller steps the user through the process of adding the shortcut to Siri. To receive notifications of events from the view controller, set the delegate to an object that conforms to the protocol. The listing below adds an button to a view and let the user record an invocation phrase.
//
// [Full Topic]: https://developer.apple.com/documentation/IntentsUI/INUIAddVoiceShortcutViewController
type INUIAddVoiceShortcutViewController struct {
	appkit.ViewController
}

// INUIAddVoiceShortcutViewControllerFrom constructs a [INUIAddVoiceShortcutViewController] from an unsafe.Pointer.
//
// A view controller that guides the user through the steps for adding a shortcut to Siri.
func INUIAddVoiceShortcutViewControllerFrom(ptr unsafe.Pointer) INUIAddVoiceShortcutViewController {
	return INUIAddVoiceShortcutViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INUIAddVoiceShortcutViewControllerClass) Alloc() INUIAddVoiceShortcutViewController {
	rv := objc.Send[INUIAddVoiceShortcutViewController](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INUIAddVoiceShortcutViewControllerClass) New() INUIAddVoiceShortcutViewController {
	rv := objc.Send[INUIAddVoiceShortcutViewController](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INUIAddVoiceShortcutViewController) Init() INUIAddVoiceShortcutViewController {
	rv := objc.Send[INUIAddVoiceShortcutViewController](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INUIAddVoiceShortcutViewController) Autorelease() INUIAddVoiceShortcutViewController {
	rv := objc.Send[INUIAddVoiceShortcutViewController](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINUIAddVoiceShortcutViewController creates a new INUIAddVoiceShortcutViewController instance.
func NewINUIAddVoiceShortcutViewController() INUIAddVoiceShortcutViewController {
	return getINUIAddVoiceShortcutViewControllerClass().New()
}


// The object that retrieves notifications from the view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/intentsui/inuiaddvoiceshortcutviewcontroller/delegate
func (i_ INUIAddVoiceShortcutViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The object that retrieves notifications from the view controller.

//
// [Full Topic]: https://developer.apple.com/documentation/intentsui/inuiaddvoiceshortcutviewcontroller/delegate
func (i_ INUIAddVoiceShortcutViewController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}



