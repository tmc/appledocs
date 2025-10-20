// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ExistsCommand] class.
var (
	existsCommandClass     _ExistsCommandClass
	existsCommandClassOnce sync.Once
)

func getExistsCommandClass() _ExistsCommandClass {
	existsCommandClassOnce.Do(func() {
		existsCommandClass = _ExistsCommandClass{objc.GetClass("NSExistsCommand")}
	})
	return existsCommandClass
}

type _ExistsCommandClass struct {
	class objc.Class
}

// An interface definition for the [ExistsCommand] class.
type IExistsCommand interface {
	IScriptCommand
}

// A command that determines whether a scriptable object exists.
//
// An instance of determines whether a specified scriptable object, such as a word, paragraph, or image, exists. When an instance of is executed, it evaluates the receiver specifier for the command to determine if it specifies any objects. is part of Cocoa’s built-in scripting support. Most applications don’t need to subclass .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getExistsCommandClass().New()
}




