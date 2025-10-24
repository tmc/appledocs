// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SetCommand] class.
var (
	SetCommandClass     _SetCommandClass
	SetCommandClassOnce sync.Once
)

func getSetCommandClass() _SetCommandClass {
	SetCommandClassOnce.Do(func() {
		SetCommandClass = _SetCommandClass{objc.GetClass("NSSetCommand")}
	})
	return SetCommandClass
}

type _SetCommandClass struct {
	class objc.Class
}

// An interface definition for the [SetCommand] class.
type ISetCommand interface {
	IScriptCommand
	// properties:
	KeySpecifier() IScriptObjectSpecifier
	SetKeySpecifier(value IScriptObjectSpecifier)
	// methods:
	SetReceiversSpecifier(receiversRef IScriptObjectSpecifier)
}

// A command that sets one or more attributes or relationships to one or more values.
//
// An instance of sets one or more attributes or relationships to one or more values; for example, it may set the (x, y) coordinates for a window’s position or set the name of a document. is part of Cocoa’s built-in scripting support. It works automatically to support the command through key-value coding. Most applications don’t need to subclass or call its methods. uses available scripting class descriptions to determine whether it should set a value for an attribute (or property), or set a value for all elements (to-many objects). For the latter, it invokes ; for the former, it invokes (or, if the receiver overrides , it invokes that method, to support backward binary compatibility.) For information on working with commands, see in .


// A command that sets one or more attributes or relationships to one or more values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSetCommand
type SetCommand struct {
	ScriptCommand
}

// SetCommandFrom constructs a [SetCommand] from an unsafe.Pointer.
//
// A command that sets one or more attributes or relationships to one or more values.
func SetCommandFrom(ptr unsafe.Pointer) SetCommand {
	return SetCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SetCommandClass) Alloc() SetCommand {
	rv := objc.Send[SetCommand](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SetCommandClass) New() SetCommand {
	rv := objc.Send[SetCommand](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SetCommand) Init() SetCommand {
	rv := objc.Send[SetCommand](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SetCommand) Autorelease() SetCommand {
	rv := objc.Send[SetCommand](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSetCommand creates a new SetCommand instance.
func NewSetCommand() SetCommand {
	return getSetCommandClass().New()
}



// Sets the receiver’s object specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSetCommand/setReceiversSpecifier(_:)
func (s_ SetCommand) SetReceiversSpecifier(receiversRef IScriptObjectSpecifier) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReceiversSpecifier:"), receiversRef)
}


// Returns a specifier that identifies the attribute or relationship that is to be set for the receiver of the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssetcommand/keyspecifier
func (s_ SetCommand) KeySpecifier() IScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](s_.ID, objc.Sel("keySpecifier"))
	return rv
}


// Returns a specifier that identifies the attribute or relationship that is to be set for the receiver of the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssetcommand/keyspecifier
func (s_ SetCommand) SetKeySpecifier(value IScriptObjectSpecifier) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKeySpecifier:"), value)
}



