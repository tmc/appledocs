// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INShortcut] class.
var (
	INShortcutClass     _INShortcutClass
	INShortcutClassOnce sync.Once
)

func getINShortcutClass() _INShortcutClass {
	INShortcutClassOnce.Do(func() {
		INShortcutClass = _INShortcutClass{objc.GetClass("INShortcut")}
	})
	return INShortcutClass
}

type _INShortcutClass struct {
	class objc.Class
}

// An interface definition for the [INShortcut] class.
type IINShortcut interface {
	objectivec.IObject
	// properties:
	UserActivity() UserActivity /* not a class type */
	Intent() INIntent /* already interface */
	SetIntent(value INIntent /* already interface */)
	// methods:
}

// An object representing an action available in your app that the system may suggest to a user or a user may add to Siri.


// An object representing an action available in your app that the system may suggest to a user or a user may add to Siri.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INShortcutReference
type INShortcut struct {
	objectivec.Object
}

// INShortcutFrom constructs a [INShortcut] from an unsafe.Pointer.
//
// An object representing an action available in your app that the system may suggest to a user or a user may add to Siri.
func INShortcutFrom(ptr unsafe.Pointer) INShortcut {
	return INShortcut{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INShortcutClass) Alloc() INShortcut {
	rv := objc.Send[INShortcut](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INShortcutClass) New() INShortcut {
	rv := objc.Send[INShortcut](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INShortcut) Init() INShortcut {
	rv := objc.Send[INShortcut](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INShortcut) Autorelease() INShortcut {
	rv := objc.Send[INShortcut](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINShortcut creates a new INShortcut instance.
func NewINShortcut() INShortcut {
	return getINShortcutClass().New()
}



// Creates a shortcut with the specified intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INShortcutReference/init(intent:)
func NewINShortcutWithIntent(intent INIntent /* already interface */) INShortcut {
	instance := getINShortcutClass().Alloc()
	rv := objc.Send[INShortcut](instance.ID, objc.Sel("initWithIntent:"), intent)
	rv.Autorelease()
	return rv
}



// The user activity that defines the action to perform when invoking the shortcut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INShortcutReference/userActivity
func (i_ INShortcut) UserActivity() UserActivity /* not a class type */ {
	rv := objc.Send[UserActivity](i_.ID, objc.Sel("userActivity"))
	return rv
}


// The intent that performs the action when invoking the shortcut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inshortcutreference/intent
func (i_ INShortcut) Intent() INIntent /* already interface */ {
	rv := objc.Send[INIntent](i_.ID, objc.Sel("intent"))
	return rv
}


// The intent that performs the action when invoking the shortcut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inshortcutreference/intent
func (i_ INShortcut) SetIntent(value INIntent /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIntent:"), value)
}


