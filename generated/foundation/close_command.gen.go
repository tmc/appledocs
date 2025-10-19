// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CloseCommand] class.
var closeCommandClass = _CloseCommandClass{objc.GetClass("NSCloseCommand")}

type _CloseCommandClass struct {
	class objc.Class
}

// An interface definition for the [CloseCommand] class.
type ICloseCommand interface {
	IScriptCommand
}

// A command that closes one or more scriptable objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCloseCommand

type CloseCommand struct {
	ScriptCommand
}

// CloseCommandFrom constructs a [CloseCommand] from an unsafe.Pointer.
//
// A command that closes one or more scriptable objects.
func CloseCommandFrom(ptr unsafe.Pointer) CloseCommand {
	return CloseCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (cc _CloseCommandClass) Alloc() CloseCommand {
	rv := objc.Send[CloseCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _CloseCommandClass) New() CloseCommand {
	rv := objc.Send[CloseCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CloseCommand) Init() CloseCommand {
	rv := objc.Send[CloseCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CloseCommand) Autorelease() CloseCommand {
	rv := objc.Send[CloseCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCloseCommand creates a new CloseCommand instance.
func NewCloseCommand() CloseCommand {
	return closeCommandClass.New()
}




