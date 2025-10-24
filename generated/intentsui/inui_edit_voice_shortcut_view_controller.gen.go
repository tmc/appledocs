// Code generated from Apple documentation for IntentsUI. DO NOT EDIT.

package intentsui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/intents"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INUIEditVoiceShortcutViewController] class.
var (
	INUIEditVoiceShortcutViewControllerClass     _INUIEditVoiceShortcutViewControllerClass
	INUIEditVoiceShortcutViewControllerClassOnce sync.Once
)

func getINUIEditVoiceShortcutViewControllerClass() _INUIEditVoiceShortcutViewControllerClass {
	INUIEditVoiceShortcutViewControllerClassOnce.Do(func() {
		INUIEditVoiceShortcutViewControllerClass = _INUIEditVoiceShortcutViewControllerClass{objc.GetClass("INUIEditVoiceShortcutViewController")}
	})
	return INUIEditVoiceShortcutViewControllerClass
}

type _INUIEditVoiceShortcutViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [INUIEditVoiceShortcutViewController] class.
type IINUIEditVoiceShortcutViewController interface {
	appkit.IViewController
	Delegate() objc.ID
	SetDelegate(value objc.ID)
}

// A view controller that lets the user edit or remove an existing shortcut.
//
// To let the user edit or remove a shortcut, create an instance of with the . Then present the view controller to the user. To receive notifications of changes made to the shortcut, set the controller’s to an object that conforms to the protocol.

// A view controller that lets the user edit or remove an existing shortcut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IntentsUI/INUIEditVoiceShortcutViewController
type INUIEditVoiceShortcutViewController struct {
	appkit.ViewController
}

// INUIEditVoiceShortcutViewControllerFrom constructs a [INUIEditVoiceShortcutViewController] from an unsafe.Pointer.
//
// A view controller that lets the user edit or remove an existing shortcut.
func INUIEditVoiceShortcutViewControllerFrom(ptr unsafe.Pointer) INUIEditVoiceShortcutViewController {
	return INUIEditVoiceShortcutViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INUIEditVoiceShortcutViewControllerClass) Alloc() INUIEditVoiceShortcutViewController {
	rv := objc.Send[INUIEditVoiceShortcutViewController](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INUIEditVoiceShortcutViewControllerClass) New() INUIEditVoiceShortcutViewController {
	rv := objc.Send[INUIEditVoiceShortcutViewController](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INUIEditVoiceShortcutViewController) Init() INUIEditVoiceShortcutViewController {
	rv := objc.Send[INUIEditVoiceShortcutViewController](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INUIEditVoiceShortcutViewController) Autorelease() INUIEditVoiceShortcutViewController {
	rv := objc.Send[INUIEditVoiceShortcutViewController](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINUIEditVoiceShortcutViewController creates a new INUIEditVoiceShortcutViewController instance.
func NewINUIEditVoiceShortcutViewController() INUIEditVoiceShortcutViewController {
	return getINUIEditVoiceShortcutViewControllerClass().New()
}

// Creates a view controller with the shortcut to edit or remove.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IntentsUI/INUIEditVoiceShortcutViewController/init(voiceShortcut:)
func NewINUIEditVoiceShortcutViewControllerWithVoiceShortcut(voiceShortcut intents.INVoiceShortcut) INUIEditVoiceShortcutViewController {
	instance := getINUIEditVoiceShortcutViewControllerClass().Alloc()
	rv := objc.Send[INUIEditVoiceShortcutViewController](instance.ID, objc.Sel("initWithVoiceShortcut:"), voiceShortcut)
	rv.Autorelease()
	return rv
}

// The object that retrieves notifications from the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IntentsUI/INUIEditVoiceShortcutViewController/delegate
func (i_ INUIEditVoiceShortcutViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("delegate"))
	return rv
}

// The object that retrieves notifications from the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IntentsUI/INUIEditVoiceShortcutViewController/delegate
func (i_ INUIEditVoiceShortcutViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}
