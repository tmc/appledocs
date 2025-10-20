// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptCommandDescription] class.
var (
	scriptCommandDescriptionClass     _ScriptCommandDescriptionClass
	scriptCommandDescriptionClassOnce sync.Once
)

func getScriptCommandDescriptionClass() _ScriptCommandDescriptionClass {
	scriptCommandDescriptionClassOnce.Do(func() {
		scriptCommandDescriptionClass = _ScriptCommandDescriptionClass{objc.GetClass("NSScriptCommandDescription")}
	})
	return scriptCommandDescriptionClass
}

type _ScriptCommandDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [ScriptCommandDescription] class.
type IScriptCommandDescription interface {
	objectivec.IObject
}

// A script command that a macOS app supports.
//
// A scriptable application provides scriptability information that describes the commands and objects scripters can use in scripts that target the application. An application’s scripting information is collected automatically by an instance of , which creates an for each command it finds, caches these objects in memory, and installs a command handler for each command. A script command instance stores the name, class, argument types, and return type of a command. For example, commands in AppleScript’s Core suite include , , , , , and . The public methods of are used primarily by Cocoa’s built-in scripting support in responding to Apple events that target the application. Although you can subclass the class, it is unlikely that you would need to do so, or to create instances of it.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription
type ScriptCommandDescription struct {
	objectivec.Object
}

// ScriptCommandDescriptionFrom constructs a [ScriptCommandDescription] from an unsafe.Pointer.
//
// A script command that a macOS app supports.
func ScriptCommandDescriptionFrom(ptr unsafe.Pointer) ScriptCommandDescription {
	return ScriptCommandDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScriptCommandDescriptionClass) Alloc() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScriptCommandDescriptionClass) New() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptCommandDescription) Init() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptCommandDescription) Autorelease() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptCommandDescription creates a new ScriptCommandDescription instance.
func NewScriptCommandDescription() ScriptCommandDescription {
	return getScriptCommandDescriptionClass().New()
}




