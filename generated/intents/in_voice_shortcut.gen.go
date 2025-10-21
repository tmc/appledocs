// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INVoiceShortcut] class.
var (
	INVoiceShortcutClass     _INVoiceShortcutClass
	INVoiceShortcutClassOnce sync.Once
)

func getINVoiceShortcutClass() _INVoiceShortcutClass {
	INVoiceShortcutClassOnce.Do(func() {
		INVoiceShortcutClass = _INVoiceShortcutClass{objc.GetClass("INVoiceShortcut")}
	})
	return INVoiceShortcutClass
}

type _INVoiceShortcutClass struct {
	class objc.Class
}

// An interface definition for the [INVoiceShortcut] class.
type IINVoiceShortcut interface {
	objectivec.IObject
}

// A shortcut the user added to Siri.
//
// To add a shortcut to Siri, create an object and add it using . You don’t create an instance of ; the system creates the instance for you when an instance is needed.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INVoiceShortcut
type INVoiceShortcut struct {
	objectivec.Object
}

// INVoiceShortcutFrom constructs a [INVoiceShortcut] from an unsafe.Pointer.
//
// A shortcut the user added to Siri.
func INVoiceShortcutFrom(ptr unsafe.Pointer) INVoiceShortcut {
	return INVoiceShortcut{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INVoiceShortcutClass) Alloc() INVoiceShortcut {
	rv := objc.Send[INVoiceShortcut](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INVoiceShortcutClass) New() INVoiceShortcut {
	rv := objc.Send[INVoiceShortcut](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INVoiceShortcut) Init() INVoiceShortcut {
	rv := objc.Send[INVoiceShortcut](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INVoiceShortcut) Autorelease() INVoiceShortcut {
	rv := objc.Send[INVoiceShortcut](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINVoiceShortcut creates a new INVoiceShortcut instance.
func NewINVoiceShortcut() INVoiceShortcut {
	return getINVoiceShortcutClass().New()
}




