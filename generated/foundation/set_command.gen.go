// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SetCommand] class.
var (
	setCommandClass     _SetCommandClass
	setCommandClassOnce sync.Once
)

func getSetCommandClass() _SetCommandClass {
	setCommandClassOnce.Do(func() {
		setCommandClass = _SetCommandClass{objc.GetClass("NSSetCommand")}
	})
	return setCommandClass
}

type _SetCommandClass struct {
	class objc.Class
}

// An interface definition for the [SetCommand] class.
type ISetCommand interface {
	IScriptCommand
}

// A command that sets one or more attributes or relationships to one or more values. [Full Topic]
//
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




