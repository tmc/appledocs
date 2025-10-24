// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INVoiceShortcutCenter] class.
var (
	INVoiceShortcutCenterClass     _INVoiceShortcutCenterClass
	INVoiceShortcutCenterClassOnce sync.Once
)

func getINVoiceShortcutCenterClass() _INVoiceShortcutCenterClass {
	INVoiceShortcutCenterClassOnce.Do(func() {
		INVoiceShortcutCenterClass = _INVoiceShortcutCenterClass{objc.GetClass("INVoiceShortcutCenter")}
	})
	return INVoiceShortcutCenterClass
}

type _INVoiceShortcutCenterClass struct {
	class objc.Class
}

// An interface definition for the [INVoiceShortcutCenter] class.
type IINVoiceShortcutCenter interface {
	objectivec.IObject
	SetShortcutSuggestions(suggestions []INShortcut)
}

// Retrieve the user’s shortcuts and make shortcut suggestions.
//
// With Shortcut Center, your app can: Retrieve shortcuts associated with your app that the user added to Siri. Suggest shortcuts the user may want to add to Siri. Before you can retrieve or suggest shortcuts, get a reference to the Shortcut Center from the class property. To retrieve all shortcuts associated with your app, call . To retrieve a particular shortcut, use the method, passing in the shortcut’s identifier. These methods return shortcuts associated with your app that the user added to Siri using your app or the Settings app. To suggest shortcuts for actions that the user hasn’t performed in your app but may want to add to Siri, call , passing in a list of suggested shortcuts. The user views the suggestions in the Gallery of the Shortcuts app. For more information, see .

// Retrieve the user’s shortcuts and make shortcut suggestions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INVoiceShortcutCenter
type INVoiceShortcutCenter struct {
	objectivec.Object
}

// INVoiceShortcutCenterFrom constructs a [INVoiceShortcutCenter] from an unsafe.Pointer.
//
// Retrieve the user’s shortcuts and make shortcut suggestions.
func INVoiceShortcutCenterFrom(ptr unsafe.Pointer) INVoiceShortcutCenter {
	return INVoiceShortcutCenter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INVoiceShortcutCenterClass) Alloc() INVoiceShortcutCenter {
	rv := objc.Send[INVoiceShortcutCenter](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INVoiceShortcutCenterClass) New() INVoiceShortcutCenter {
	rv := objc.Send[INVoiceShortcutCenter](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INVoiceShortcutCenter) Init() INVoiceShortcutCenter {
	rv := objc.Send[INVoiceShortcutCenter](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INVoiceShortcutCenter) Autorelease() INVoiceShortcutCenter {
	rv := objc.Send[INVoiceShortcutCenter](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINVoiceShortcutCenter creates a new INVoiceShortcutCenter instance.
func NewINVoiceShortcutCenter() INVoiceShortcutCenter {
	return getINVoiceShortcutCenterClass().New()
}

// The shared shortcut center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INVoiceShortcutCenter/shared
func (ic _INVoiceShortcutCenterClass) SharedCenter() INVoiceShortcutCenter {
	rv := objc.Send[INVoiceShortcutCenter](objc.ID(ic.class), objc.Sel("sharedCenter"))
	return rv
}

// Suggests shortcuts the user may want to add to Siri.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INVoiceShortcutCenter/setShortcutSuggestions(_:)
func (i_ INVoiceShortcutCenter) SetShortcutSuggestions(suggestions []INShortcut) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setShortcutSuggestions:"), suggestions)
}

// The shared shortcut center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INVoiceShortcutCenter/shared
func (i_ INVoiceShortcutCenter) SharedCenter() INVoiceShortcutCenter {
	rv := objc.Send[INVoiceShortcutCenter](i_.ID, objc.Sel("sharedCenter"))
	return rv
}
