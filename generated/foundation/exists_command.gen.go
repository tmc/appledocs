// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ExistsCommand] class.
var existsCommandClass = _ExistsCommandClass{objc.GetClass("NSExistsCommand")}

type _ExistsCommandClass struct {
	class objc.Class
}

// An interface definition for the [ExistsCommand] class.
type IExistsCommand interface {
	IScriptCommand
}

// A command that determines whether a scriptable object exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExistsCommand

type ExistsCommand struct {
	ScriptCommand
}

// ExistsCommandFrom constructs a [ExistsCommand] from an unsafe.Pointer.
//
// A command that determines whether a scriptable object exists.
func ExistsCommandFrom(ptr unsafe.Pointer) ExistsCommand {
	return ExistsCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (ec _ExistsCommandClass) Alloc() ExistsCommand {
	rv := objc.Send[ExistsCommand](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ec _ExistsCommandClass) New() ExistsCommand {
	rv := objc.Send[ExistsCommand](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExistsCommand) Init() ExistsCommand {
	rv := objc.Send[ExistsCommand](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExistsCommand) Autorelease() ExistsCommand {
	rv := objc.Send[ExistsCommand](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExistsCommand creates a new ExistsCommand instance.
func NewExistsCommand() ExistsCommand {
	return existsCommandClass.New()
}




