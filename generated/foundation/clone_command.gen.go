// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CloneCommand] class.
var (
	cloneCommandClass     _CloneCommandClass
	cloneCommandClassOnce sync.Once
)

func getCloneCommandClass() _CloneCommandClass {
	cloneCommandClassOnce.Do(func() {
		cloneCommandClass = _CloneCommandClass{objc.GetClass("NSCloneCommand")}
	})
	return cloneCommandClass
}

type _CloneCommandClass struct {
	class objc.Class
}

// An interface definition for the [CloneCommand] class.
type ICloneCommand interface {
	IScriptCommand
}

// A command that clones one or more scriptable objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCloneCommand
type CloneCommand struct {
	ScriptCommand
}

// CloneCommandFrom constructs a [CloneCommand] from an unsafe.Pointer.
//
// A command that clones one or more scriptable objects.
func CloneCommandFrom(ptr unsafe.Pointer) CloneCommand {
	return CloneCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CloneCommandClass) Alloc() CloneCommand {
	rv := objc.Send[CloneCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CloneCommandClass) New() CloneCommand {
	rv := objc.Send[CloneCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CloneCommand) Init() CloneCommand {
	rv := objc.Send[CloneCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CloneCommand) Autorelease() CloneCommand {
	rv := objc.Send[CloneCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCloneCommand creates a new CloneCommand instance.
func NewCloneCommand() CloneCommand {
	return getCloneCommandClass().New()
}




